// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vm

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

// VmSafeAccountAccess is an auto generated low-level Go binding around an user-defined struct.
type VmSafeAccountAccess struct {
	ChainInfo       VmSafeChainInfo
	Kind            uint8
	Account         common.Address
	Accessor        common.Address
	Initialized     bool
	OldBalance      *big.Int
	NewBalance      *big.Int
	DeployedCode    []byte
	Value           *big.Int
	Data            []byte
	Reverted        bool
	StorageAccesses []VmSafeStorageAccess
	Depth           uint64
}

// VmSafeChainInfo is an auto generated low-level Go binding around an user-defined struct.
type VmSafeChainInfo struct {
	ForkId  *big.Int
	ChainId *big.Int
}

// VmSafeDirEntry is an auto generated low-level Go binding around an user-defined struct.
type VmSafeDirEntry struct {
	ErrorMessage string
	Path         string
	Depth        uint64
	IsDir        bool
	IsSymlink    bool
}

// VmSafeEthGetLogs is an auto generated low-level Go binding around an user-defined struct.
type VmSafeEthGetLogs struct {
	Emitter          common.Address
	Topics           [][32]byte
	Data             []byte
	BlockHash        [32]byte
	BlockNumber      uint64
	TransactionHash  [32]byte
	TransactionIndex uint64
	LogIndex         *big.Int
	Removed          bool
}

// VmSafeFfiResult is an auto generated low-level Go binding around an user-defined struct.
type VmSafeFfiResult struct {
	ExitCode int32
	Stdout   []byte
	Stderr   []byte
}

// VmSafeFsMetadata is an auto generated low-level Go binding around an user-defined struct.
type VmSafeFsMetadata struct {
	IsDir     bool
	IsSymlink bool
	Length    *big.Int
	ReadOnly  bool
	Modified  *big.Int
	Accessed  *big.Int
	Created   *big.Int
}

// VmSafeGas is an auto generated low-level Go binding around an user-defined struct.
type VmSafeGas struct {
	GasLimit      uint64
	GasTotalUsed  uint64
	GasMemoryUsed uint64
	GasRefunded   int64
	GasRemaining  uint64
}

// VmSafeLog is an auto generated low-level Go binding around an user-defined struct.
type VmSafeLog struct {
	Topics  [][32]byte
	Data    []byte
	Emitter common.Address
}

// VmSafeRpc is an auto generated low-level Go binding around an user-defined struct.
type VmSafeRpc struct {
	Key string
	Url string
}

// VmSafeStorageAccess is an auto generated low-level Go binding around an user-defined struct.
type VmSafeStorageAccess struct {
	Account       common.Address
	Slot          [32]byte
	IsWrite       bool
	PreviousValue [32]byte
	NewValue      [32]byte
	Reverted      bool
}

// VmSafeWallet is an auto generated low-level Go binding around an user-defined struct.
type VmSafeWallet struct {
	Addr       common.Address
	PublicKeyX *big.Int
	PublicKeyY *big.Int
	PrivateKey *big.Int
}

// VmMetaData contains all meta data concerning the Vm contract.
var VmMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"accesses\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"readSlots\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"writeSlots\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activeFork\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"forkId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"privateKey\",\"type\":\"uint256\"}],\"name\":\"addr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"keyAddr\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"allowCheatcodes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"left\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"right\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDelta\",\"type\":\"uint256\"}],\"name\":\"assertApproxEqAbs\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"left\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"right\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"maxDelta\",\"type\":\"uint256\"}],\"name\":\"assertApproxEqAbs\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"left\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"right\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"maxDelta\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertApproxEqAbs\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"left\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"right\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDelta\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertApproxEqAbs\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"left\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"right\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDelta\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"}],\"name\":\"assertApproxEqAbsDecimal\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"left\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"right\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"maxDelta\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"}],\"name\":\"assertApproxEqAbsDecimal\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"left\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"right\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDelta\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertApproxEqAbsDecimal\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"left\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"right\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"maxDelta\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertApproxEqAbsDecimal\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"left\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"right\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPercentDelta\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertApproxEqRel\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"left\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"right\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPercentDelta\",\"type\":\"uint256\"}],\"name\":\"assertApproxEqRel\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"left\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"right\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"maxPercentDelta\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertApproxEqRel\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"left\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"right\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"maxPercentDelta\",\"type\":\"uint256\"}],\"name\":\"assertApproxEqRel\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"left\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"right\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPercentDelta\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"}],\"name\":\"assertApproxEqRelDecimal\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"left\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"right\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPercentDelta\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertApproxEqRelDecimal\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"left\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"right\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"maxPercentDelta\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"}],\"name\":\"assertApproxEqRelDecimal\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"left\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"right\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"maxPercentDelta\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertApproxEqRelDecimal\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"left\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"right\",\"type\":\"bytes32[]\"}],\"name\":\"assertEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256[]\",\"name\":\"left\",\"type\":\"int256[]\"},{\"internalType\":\"int256[]\",\"name\":\"right\",\"type\":\"int256[]\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"left\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"right\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"left\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"right\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"left\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"right\",\"type\":\"address[]\"}],\"name\":\"assertEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"left\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"right\",\"type\":\"address[]\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"left\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"right\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"left\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"right\",\"type\":\"address\"}],\"name\":\"assertEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"left\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"right\",\"type\":\"uint256[]\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool[]\",\"name\":\"left\",\"type\":\"bool[]\"},{\"internalType\":\"bool[]\",\"name\":\"right\",\"type\":\"bool[]\"}],\"name\":\"assertEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256[]\",\"name\":\"left\",\"type\":\"int256[]\"},{\"internalType\":\"int256[]\",\"name\":\"right\",\"type\":\"int256[]\"}],\"name\":\"assertEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"left\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"right\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"left\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"right\",\"type\":\"bytes32\"}],\"name\":\"assertEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"left\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"right\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"left\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"right\",\"type\":\"uint256[]\"}],\"name\":\"assertEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"left\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"right\",\"type\":\"bytes\"}],\"name\":\"assertEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"left\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"right\",\"type\":\"uint256\"}],\"name\":\"assertEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"left\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"right\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"left\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"right\",\"type\":\"string[]\"}],\"name\":\"assertEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"left\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"right\",\"type\":\"bytes32[]\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"left\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"right\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool[]\",\"name\":\"left\",\"type\":\"bool[]\"},{\"internalType\":\"bool[]\",\"name\":\"right\",\"type\":\"bool[]\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"left\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"right\",\"type\":\"bytes[]\"}],\"name\":\"assertEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"left\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"right\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"left\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"right\",\"type\":\"string\"}],\"name\":\"assertEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"left\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"right\",\"type\":\"bytes[]\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"left\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"right\",\"type\":\"bool\"}],\"name\":\"assertEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"left\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"right\",\"type\":\"int256\"}],\"name\":\"assertEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"left\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"right\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"}],\"name\":\"assertEqDecimal\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"left\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"right\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"}],\"name\":\"assertEqDecimal\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"left\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"right\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertEqDecimal\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"left\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"right\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertEqDecimal\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"condition\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertFalse\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"condition\",\"type\":\"bool\"}],\"name\":\"assertFalse\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"left\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"right\",\"type\":\"int256\"}],\"name\":\"assertGe\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"left\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"right\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertGe\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"left\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"right\",\"type\":\"uint256\"}],\"name\":\"assertGe\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"left\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"right\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertGe\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"left\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"right\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"}],\"name\":\"assertGeDecimal\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"left\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"right\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertGeDecimal\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"left\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"right\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertGeDecimal\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"left\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"right\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"}],\"name\":\"assertGeDecimal\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"left\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"right\",\"type\":\"int256\"}],\"name\":\"assertGt\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"left\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"right\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertGt\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"left\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"right\",\"type\":\"uint256\"}],\"name\":\"assertGt\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"left\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"right\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertGt\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"left\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"right\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertGtDecimal\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"left\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"right\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertGtDecimal\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"left\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"right\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"}],\"name\":\"assertGtDecimal\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"left\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"right\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"}],\"name\":\"assertGtDecimal\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"left\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"right\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertLe\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"left\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"right\",\"type\":\"uint256\"}],\"name\":\"assertLe\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"left\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"right\",\"type\":\"int256\"}],\"name\":\"assertLe\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"left\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"right\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertLe\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"left\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"right\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"}],\"name\":\"assertLeDecimal\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"left\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"right\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertLeDecimal\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"left\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"right\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertLeDecimal\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"left\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"right\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"}],\"name\":\"assertLeDecimal\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"left\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"right\",\"type\":\"int256\"}],\"name\":\"assertLt\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"left\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"right\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertLt\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"left\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"right\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertLt\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"left\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"right\",\"type\":\"uint256\"}],\"name\":\"assertLt\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"left\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"right\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"}],\"name\":\"assertLtDecimal\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"left\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"right\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertLtDecimal\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"left\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"right\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertLtDecimal\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"left\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"right\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"}],\"name\":\"assertLtDecimal\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"left\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"right\",\"type\":\"bytes32[]\"}],\"name\":\"assertNotEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256[]\",\"name\":\"left\",\"type\":\"int256[]\"},{\"internalType\":\"int256[]\",\"name\":\"right\",\"type\":\"int256[]\"}],\"name\":\"assertNotEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"left\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"right\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertNotEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"left\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"right\",\"type\":\"bytes[]\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertNotEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"left\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"right\",\"type\":\"bool\"}],\"name\":\"assertNotEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool[]\",\"name\":\"left\",\"type\":\"bool[]\"},{\"internalType\":\"bool[]\",\"name\":\"right\",\"type\":\"bool[]\"}],\"name\":\"assertNotEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"left\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"right\",\"type\":\"bytes\"}],\"name\":\"assertNotEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"left\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"right\",\"type\":\"address[]\"}],\"name\":\"assertNotEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"left\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"right\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertNotEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"left\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"right\",\"type\":\"uint256[]\"}],\"name\":\"assertNotEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool[]\",\"name\":\"left\",\"type\":\"bool[]\"},{\"internalType\":\"bool[]\",\"name\":\"right\",\"type\":\"bool[]\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertNotEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"left\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"right\",\"type\":\"string\"}],\"name\":\"assertNotEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"left\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"right\",\"type\":\"address[]\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertNotEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"left\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"right\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertNotEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"left\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"right\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertNotEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"left\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"right\",\"type\":\"bytes32\"}],\"name\":\"assertNotEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"left\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"right\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertNotEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"left\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"right\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertNotEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"left\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"right\",\"type\":\"uint256[]\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertNotEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"left\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"right\",\"type\":\"address\"}],\"name\":\"assertNotEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"left\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"right\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertNotEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"left\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"right\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertNotEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"left\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"right\",\"type\":\"uint256\"}],\"name\":\"assertNotEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"left\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"right\",\"type\":\"bytes32[]\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertNotEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"left\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"right\",\"type\":\"string[]\"}],\"name\":\"assertNotEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256[]\",\"name\":\"left\",\"type\":\"int256[]\"},{\"internalType\":\"int256[]\",\"name\":\"right\",\"type\":\"int256[]\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertNotEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"left\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"right\",\"type\":\"bytes[]\"}],\"name\":\"assertNotEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"left\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"right\",\"type\":\"int256\"}],\"name\":\"assertNotEq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"left\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"right\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"}],\"name\":\"assertNotEqDecimal\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"left\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"right\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertNotEqDecimal\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"left\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"right\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"}],\"name\":\"assertNotEqDecimal\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"left\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"right\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertNotEqDecimal\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"condition\",\"type\":\"bool\"}],\"name\":\"assertTrue\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"condition\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"assertTrue\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"condition\",\"type\":\"bool\"}],\"name\":\"assume\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newBlobBaseFee\",\"type\":\"uint256\"}],\"name\":\"blobBaseFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"hashes\",\"type\":\"bytes32[]\"}],\"name\":\"blobhashes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"char\",\"type\":\"string\"}],\"name\":\"breakpoint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"char\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"breakpoint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"broadcast\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"broadcast\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"privateKey\",\"type\":\"uint256\"}],\"name\":\"broadcast\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newChainId\",\"type\":\"uint256\"}],\"name\":\"chainId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clearMockedCalls\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"}],\"name\":\"closeFile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newCoinbase\",\"type\":\"address\"}],\"name\":\"coinbase\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"initCodeHash\",\"type\":\"bytes32\"}],\"name\":\"computeCreate2Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"initCodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"deployer\",\"type\":\"address\"}],\"name\":\"computeCreate2Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"deployer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"computeCreateAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"from\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"}],\"name\":\"copyFile\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"copied\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"recursive\",\"type\":\"bool\"}],\"name\":\"createDir\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"urlOrAlias\",\"type\":\"string\"}],\"name\":\"createFork\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"forkId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"urlOrAlias\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"createFork\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"forkId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"urlOrAlias\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"name\":\"createFork\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"forkId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"urlOrAlias\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"createSelectFork\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"forkId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"urlOrAlias\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"name\":\"createSelectFork\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"forkId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"urlOrAlias\",\"type\":\"string\"}],\"name\":\"createSelectFork\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"forkId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"walletLabel\",\"type\":\"string\"}],\"name\":\"createWallet\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"publicKeyX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"publicKeyY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"privateKey\",\"type\":\"uint256\"}],\"internalType\":\"structVmSafe.Wallet\",\"name\":\"wallet\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"privateKey\",\"type\":\"uint256\"}],\"name\":\"createWallet\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"publicKeyX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"publicKeyY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"privateKey\",\"type\":\"uint256\"}],\"internalType\":\"structVmSafe.Wallet\",\"name\":\"wallet\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"privateKey\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"walletLabel\",\"type\":\"string\"}],\"name\":\"createWallet\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"publicKeyX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"publicKeyY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"privateKey\",\"type\":\"uint256\"}],\"internalType\":\"structVmSafe.Wallet\",\"name\":\"wallet\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newBalance\",\"type\":\"uint256\"}],\"name\":\"deal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"snapshotId\",\"type\":\"uint256\"}],\"name\":\"deleteSnapshot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deleteSnapshots\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mnemonic\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"derivationPath\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"language\",\"type\":\"string\"}],\"name\":\"deriveKey\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"privateKey\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mnemonic\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"language\",\"type\":\"string\"}],\"name\":\"deriveKey\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"privateKey\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mnemonic\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"name\":\"deriveKey\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"privateKey\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mnemonic\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"derivationPath\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"name\":\"deriveKey\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"privateKey\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newDifficulty\",\"type\":\"uint256\"}],\"name\":\"difficulty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"pathToStateJson\",\"type\":\"string\"}],\"name\":\"dumpState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"ensNamehash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"envAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"delim\",\"type\":\"string\"}],\"name\":\"envAddress\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"value\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"envBool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"delim\",\"type\":\"string\"}],\"name\":\"envBool\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"value\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"envBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"delim\",\"type\":\"string\"}],\"name\":\"envBytes\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"value\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"delim\",\"type\":\"string\"}],\"name\":\"envBytes32\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"value\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"envBytes32\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"envExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"delim\",\"type\":\"string\"}],\"name\":\"envInt\",\"outputs\":[{\"internalType\":\"int256[]\",\"name\":\"value\",\"type\":\"int256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"envInt\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"delim\",\"type\":\"string\"},{\"internalType\":\"bytes32[]\",\"name\":\"defaultValue\",\"type\":\"bytes32[]\"}],\"name\":\"envOr\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"value\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"delim\",\"type\":\"string\"},{\"internalType\":\"int256[]\",\"name\":\"defaultValue\",\"type\":\"int256[]\"}],\"name\":\"envOr\",\"outputs\":[{\"internalType\":\"int256[]\",\"name\":\"value\",\"type\":\"int256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"defaultValue\",\"type\":\"bool\"}],\"name\":\"envOr\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"defaultValue\",\"type\":\"address\"}],\"name\":\"envOr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"defaultValue\",\"type\":\"uint256\"}],\"name\":\"envOr\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"delim\",\"type\":\"string\"},{\"internalType\":\"bytes[]\",\"name\":\"defaultValue\",\"type\":\"bytes[]\"}],\"name\":\"envOr\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"value\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"delim\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"defaultValue\",\"type\":\"uint256[]\"}],\"name\":\"envOr\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"value\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"delim\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"defaultValue\",\"type\":\"string[]\"}],\"name\":\"envOr\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"value\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"defaultValue\",\"type\":\"bytes\"}],\"name\":\"envOr\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"defaultValue\",\"type\":\"bytes32\"}],\"name\":\"envOr\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"defaultValue\",\"type\":\"int256\"}],\"name\":\"envOr\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"delim\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"defaultValue\",\"type\":\"address[]\"}],\"name\":\"envOr\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"value\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"defaultValue\",\"type\":\"string\"}],\"name\":\"envOr\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"delim\",\"type\":\"string\"},{\"internalType\":\"bool[]\",\"name\":\"defaultValue\",\"type\":\"bool[]\"}],\"name\":\"envOr\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"value\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"delim\",\"type\":\"string\"}],\"name\":\"envString\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"value\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"envString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"envUint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"delim\",\"type\":\"string\"}],\"name\":\"envUint\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"value\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"newRuntimeBytecode\",\"type\":\"bytes\"}],\"name\":\"etch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fromBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toBlock\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"topics\",\"type\":\"bytes32[]\"}],\"name\":\"eth_getLogs\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"emitter\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"topics\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"transactionHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"transactionIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"logIndex\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"removed\",\"type\":\"bool\"}],\"internalType\":\"structVmSafe.EthGetLogs[]\",\"name\":\"logs\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"callee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"gas\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"expectCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"callee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"gas\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"count\",\"type\":\"uint64\"}],\"name\":\"expectCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"callee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"count\",\"type\":\"uint64\"}],\"name\":\"expectCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"callee\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"expectCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"callee\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"count\",\"type\":\"uint64\"}],\"name\":\"expectCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"callee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"expectCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"callee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"minGas\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"expectCallMinGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"callee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"minGas\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"count\",\"type\":\"uint64\"}],\"name\":\"expectCallMinGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expectEmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"checkTopic1\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"checkTopic2\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"checkTopic3\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"checkData\",\"type\":\"bool\"}],\"name\":\"expectEmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"checkTopic1\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"checkTopic2\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"checkTopic3\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"checkData\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"emitter\",\"type\":\"address\"}],\"name\":\"expectEmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"emitter\",\"type\":\"address\"}],\"name\":\"expectEmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"revertData\",\"type\":\"bytes4\"}],\"name\":\"expectRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"revertData\",\"type\":\"bytes\"}],\"name\":\"expectRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expectRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"min\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"max\",\"type\":\"uint64\"}],\"name\":\"expectSafeMemory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"min\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"max\",\"type\":\"uint64\"}],\"name\":\"expectSafeMemoryCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newBasefee\",\"type\":\"uint256\"}],\"name\":\"fee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"commandInput\",\"type\":\"string[]\"}],\"name\":\"ffi\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"}],\"name\":\"fsMetadata\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isDir\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isSymlink\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"readOnly\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"modified\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accessed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"created\",\"type\":\"uint256\"}],\"internalType\":\"structVmSafe.FsMetadata\",\"name\":\"metadata\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlobBaseFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blobBaseFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlobhashes\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"hashes\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"artifactPath\",\"type\":\"string\"}],\"name\":\"getCode\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"creationBytecode\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"artifactPath\",\"type\":\"string\"}],\"name\":\"getDeployedCode\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"runtimeBytecode\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getLabel\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"currentLabel\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"elementSlot\",\"type\":\"bytes32\"}],\"name\":\"getMappingKeyAndParentOf\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"found\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"parent\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"mappingSlot\",\"type\":\"bytes32\"}],\"name\":\"getMappingLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"mappingSlot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"getMappingSlotAt\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"publicKeyX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"publicKeyY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"privateKey\",\"type\":\"uint256\"}],\"internalType\":\"structVmSafe.Wallet\",\"name\":\"wallet\",\"type\":\"tuple\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRecordedLogs\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"topics\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"emitter\",\"type\":\"address\"}],\"internalType\":\"structVmSafe.Log[]\",\"name\":\"logs\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"input\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"indexOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumVmSafe.ForgeContext\",\"name\":\"context\",\"type\":\"uint8\"}],\"name\":\"isContext\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"}],\"name\":\"isDir\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"}],\"name\":\"isFile\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isPersistent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"persistent\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"keyExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"keyExistsJson\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"toml\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"keyExistsToml\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"newLabel\",\"type\":\"string\"}],\"name\":\"label\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastCallGas\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gasTotalUsed\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gasMemoryUsed\",\"type\":\"uint64\"},{\"internalType\":\"int64\",\"name\":\"gasRefunded\",\"type\":\"int64\"},{\"internalType\":\"uint64\",\"name\":\"gasRemaining\",\"type\":\"uint64\"}],\"internalType\":\"structVmSafe.Gas\",\"name\":\"gas\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"load\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"data\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"pathToAllocsJson\",\"type\":\"string\"}],\"name\":\"loadAllocs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"makePersistent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account1\",\"type\":\"address\"}],\"name\":\"makePersistent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"makePersistent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account2\",\"type\":\"address\"}],\"name\":\"makePersistent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"callee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"mockCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"callee\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"mockCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"callee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"revertData\",\"type\":\"bytes\"}],\"name\":\"mockCallRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"callee\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"revertData\",\"type\":\"bytes\"}],\"name\":\"mockCallRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"stringifiedValue\",\"type\":\"string\"}],\"name\":\"parseAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"parsedValue\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"stringifiedValue\",\"type\":\"string\"}],\"name\":\"parseBool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"parsedValue\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"stringifiedValue\",\"type\":\"string\"}],\"name\":\"parseBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"parsedValue\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"stringifiedValue\",\"type\":\"string\"}],\"name\":\"parseBytes32\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"parsedValue\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"stringifiedValue\",\"type\":\"string\"}],\"name\":\"parseInt\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"parsedValue\",\"type\":\"int256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"}],\"name\":\"parseJson\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"abiEncodedData\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"parseJson\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"abiEncodedData\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"parseJsonAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"parseJsonAddressArray\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"parseJsonBool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"parseJsonBoolArray\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"parseJsonBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"parseJsonBytes32\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"parseJsonBytes32Array\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"parseJsonBytesArray\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"parseJsonInt\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"parseJsonIntArray\",\"outputs\":[{\"internalType\":\"int256[]\",\"name\":\"\",\"type\":\"int256[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"parseJsonKeys\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"keys\",\"type\":\"string[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"parseJsonString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"parseJsonStringArray\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"parseJsonUint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"parseJsonUintArray\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"toml\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"parseToml\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"abiEncodedData\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"toml\",\"type\":\"string\"}],\"name\":\"parseToml\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"abiEncodedData\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"toml\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"parseTomlAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"toml\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"parseTomlAddressArray\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"toml\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"parseTomlBool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"toml\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"parseTomlBoolArray\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"toml\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"parseTomlBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"toml\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"parseTomlBytes32\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"toml\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"parseTomlBytes32Array\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"toml\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"parseTomlBytesArray\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"toml\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"parseTomlInt\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"toml\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"parseTomlIntArray\",\"outputs\":[{\"internalType\":\"int256[]\",\"name\":\"\",\"type\":\"int256[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"toml\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"parseTomlKeys\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"keys\",\"type\":\"string[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"toml\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"parseTomlString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"toml\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"parseTomlStringArray\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"toml\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"parseTomlUint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"toml\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"parseTomlUintArray\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"stringifiedValue\",\"type\":\"string\"}],\"name\":\"parseUint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"parsedValue\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseGasMetering\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"txOrigin\",\"type\":\"address\"}],\"name\":\"prank\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"}],\"name\":\"prank\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newPrevrandao\",\"type\":\"bytes32\"}],\"name\":\"prevrandao\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPrevrandao\",\"type\":\"uint256\"}],\"name\":\"prevrandao\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"projectRoot\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"promptText\",\"type\":\"string\"}],\"name\":\"prompt\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"input\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"promptText\",\"type\":\"string\"}],\"name\":\"promptAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"promptText\",\"type\":\"string\"}],\"name\":\"promptSecret\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"input\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"promptText\",\"type\":\"string\"}],\"name\":\"promptSecretUint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"promptText\",\"type\":\"string\"}],\"name\":\"promptUint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"randomAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"randomUint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"randomUint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"readCallers\",\"outputs\":[{\"internalType\":\"enumVmSafe.CallerMode\",\"name\":\"callerMode\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"txOrigin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"maxDepth\",\"type\":\"uint64\"}],\"name\":\"readDir\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"errorMessage\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"depth\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"isDir\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isSymlink\",\"type\":\"bool\"}],\"internalType\":\"structVmSafe.DirEntry[]\",\"name\":\"entries\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"maxDepth\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"followLinks\",\"type\":\"bool\"}],\"name\":\"readDir\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"errorMessage\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"depth\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"isDir\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isSymlink\",\"type\":\"bool\"}],\"internalType\":\"structVmSafe.DirEntry[]\",\"name\":\"entries\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"}],\"name\":\"readDir\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"errorMessage\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"depth\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"isDir\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isSymlink\",\"type\":\"bool\"}],\"internalType\":\"structVmSafe.DirEntry[]\",\"name\":\"entries\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"}],\"name\":\"readFile\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"}],\"name\":\"readFileBinary\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"}],\"name\":\"readLine\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"line\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"linkPath\",\"type\":\"string\"}],\"name\":\"readLink\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"targetPath\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"record\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"recordLogs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"privateKey\",\"type\":\"uint256\"}],\"name\":\"rememberKey\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"keyAddr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"recursive\",\"type\":\"bool\"}],\"name\":\"removeDir\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"}],\"name\":\"removeFile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"input\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"from\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"}],\"name\":\"replace\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"output\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"resetNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resumeGasMetering\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"snapshotId\",\"type\":\"uint256\"}],\"name\":\"revertTo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"snapshotId\",\"type\":\"uint256\"}],\"name\":\"revertToAndDelete\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"revokePersistent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokePersistent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newHeight\",\"type\":\"uint256\"}],\"name\":\"roll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"name\":\"rollFork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"forkId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"rollFork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"rollFork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"forkId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"name\":\"rollFork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"method\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"params\",\"type\":\"string\"}],\"name\":\"rpc\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"rpcAlias\",\"type\":\"string\"}],\"name\":\"rpcUrl\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rpcUrlStructs\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structVmSafe.Rpc[]\",\"name\":\"urls\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rpcUrls\",\"outputs\":[{\"internalType\":\"string[2][]\",\"name\":\"urls\",\"type\":\"string[2][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"forkId\",\"type\":\"uint256\"}],\"name\":\"selectFork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"objectKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"valueKey\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"values\",\"type\":\"address[]\"}],\"name\":\"serializeAddress\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"objectKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"valueKey\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"}],\"name\":\"serializeAddress\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"objectKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"valueKey\",\"type\":\"string\"},{\"internalType\":\"bool[]\",\"name\":\"values\",\"type\":\"bool[]\"}],\"name\":\"serializeBool\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"objectKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"valueKey\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"serializeBool\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"objectKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"valueKey\",\"type\":\"string\"},{\"internalType\":\"bytes[]\",\"name\":\"values\",\"type\":\"bytes[]\"}],\"name\":\"serializeBytes\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"objectKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"valueKey\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"serializeBytes\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"objectKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"valueKey\",\"type\":\"string\"},{\"internalType\":\"bytes32[]\",\"name\":\"values\",\"type\":\"bytes32[]\"}],\"name\":\"serializeBytes32\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"objectKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"valueKey\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"serializeBytes32\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"objectKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"valueKey\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"name\":\"serializeInt\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"objectKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"valueKey\",\"type\":\"string\"},{\"internalType\":\"int256[]\",\"name\":\"values\",\"type\":\"int256[]\"}],\"name\":\"serializeInt\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"objectKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"serializeJson\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"objectKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"valueKey\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"values\",\"type\":\"string[]\"}],\"name\":\"serializeString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"objectKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"valueKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"serializeString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"objectKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"valueKey\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"serializeUint\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"objectKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"valueKey\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"serializeUint\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"objectKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"valueKey\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"serializeUintToHex\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"setEnv\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"newNonce\",\"type\":\"uint64\"}],\"name\":\"setNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"newNonce\",\"type\":\"uint64\"}],\"name\":\"setNonceUnsafe\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"sign\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"sign\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"publicKeyX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"publicKeyY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"privateKey\",\"type\":\"uint256\"}],\"internalType\":\"structVmSafe.Wallet\",\"name\":\"wallet\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"sign\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"privateKey\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"sign\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"privateKey\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"signP256\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"skipTest\",\"type\":\"bool\"}],\"name\":\"skip\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"sleep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"snapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"snapshotId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"input\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"delimiter\",\"type\":\"string\"}],\"name\":\"split\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"outputs\",\"type\":\"string[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startBroadcast\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"startBroadcast\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"privateKey\",\"type\":\"uint256\"}],\"name\":\"startBroadcast\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startMappingRecording\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"}],\"name\":\"startPrank\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"txOrigin\",\"type\":\"address\"}],\"name\":\"startPrank\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startStateDiffRecording\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stopAndReturnStateDiff\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"forkId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structVmSafe.ChainInfo\",\"name\":\"chainInfo\",\"type\":\"tuple\"},{\"internalType\":\"enumVmSafe.AccountAccessKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"accessor\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"oldBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newBalance\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"deployedCode\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"reverted\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isWrite\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"previousValue\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newValue\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"reverted\",\"type\":\"bool\"}],\"internalType\":\"structVmSafe.StorageAccess[]\",\"name\":\"storageAccesses\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"depth\",\"type\":\"uint64\"}],\"internalType\":\"structVmSafe.AccountAccess[]\",\"name\":\"accountAccesses\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stopBroadcast\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stopExpectSafeMemory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stopMappingRecording\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stopPrank\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"store\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"toBase64\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"toBase64\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"toBase64URL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"toBase64URL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"input\",\"type\":\"string\"}],\"name\":\"toLowercase\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"output\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"}],\"name\":\"toString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"stringifiedValue\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"toString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"stringifiedValue\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"toString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"stringifiedValue\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"toString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"stringifiedValue\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"name\":\"toString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"stringifiedValue\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"toString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"stringifiedValue\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"input\",\"type\":\"string\"}],\"name\":\"toUppercase\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"output\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"forkId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"name\":\"transact\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"name\":\"transact\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"input\",\"type\":\"string\"}],\"name\":\"trim\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"output\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"commandInput\",\"type\":\"string[]\"}],\"name\":\"tryFfi\",\"outputs\":[{\"components\":[{\"internalType\":\"int32\",\"name\":\"exitCode\",\"type\":\"int32\"},{\"internalType\":\"bytes\",\"name\":\"stdout\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"stderr\",\"type\":\"bytes\"}],\"internalType\":\"structVmSafe.FfiResult\",\"name\":\"result\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newGasPrice\",\"type\":\"uint256\"}],\"name\":\"txGasPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unixTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"milliseconds\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTimestamp\",\"type\":\"uint256\"}],\"name\":\"warp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"writeFile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"writeFileBinary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"valueKey\",\"type\":\"string\"}],\"name\":\"writeJson\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"}],\"name\":\"writeJson\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"writeLine\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"valueKey\",\"type\":\"string\"}],\"name\":\"writeToml\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"}],\"name\":\"writeToml\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// VmABI is the input ABI used to generate the binding from.
// Deprecated: Use VmMetaData.ABI instead.
var VmABI = VmMetaData.ABI

// Vm is an auto generated Go binding around an Ethereum contract.
type Vm struct {
	VmCaller     // Read-only binding to the contract
	VmTransactor // Write-only binding to the contract
	VmFilterer   // Log filterer for contract events
}

// VmCaller is an auto generated read-only Go binding around an Ethereum contract.
type VmCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VmTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VmTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VmFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VmFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VmSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VmSession struct {
	Contract     *Vm               // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VmCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VmCallerSession struct {
	Contract *VmCaller     // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VmTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VmTransactorSession struct {
	Contract     *VmTransactor     // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VmRaw is an auto generated low-level Go binding around an Ethereum contract.
type VmRaw struct {
	Contract *Vm // Generic contract binding to access the raw methods on
}

// VmCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VmCallerRaw struct {
	Contract *VmCaller // Generic read-only contract binding to access the raw methods on
}

// VmTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VmTransactorRaw struct {
	Contract *VmTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVm creates a new instance of Vm, bound to a specific deployed contract.
func NewVm(address common.Address, backend bind.ContractBackend) (*Vm, error) {
	contract, err := bindVm(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vm{VmCaller: VmCaller{contract: contract}, VmTransactor: VmTransactor{contract: contract}, VmFilterer: VmFilterer{contract: contract}}, nil
}

// NewVmCaller creates a new read-only instance of Vm, bound to a specific deployed contract.
func NewVmCaller(address common.Address, caller bind.ContractCaller) (*VmCaller, error) {
	contract, err := bindVm(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VmCaller{contract: contract}, nil
}

// NewVmTransactor creates a new write-only instance of Vm, bound to a specific deployed contract.
func NewVmTransactor(address common.Address, transactor bind.ContractTransactor) (*VmTransactor, error) {
	contract, err := bindVm(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VmTransactor{contract: contract}, nil
}

// NewVmFilterer creates a new log filterer instance of Vm, bound to a specific deployed contract.
func NewVmFilterer(address common.Address, filterer bind.ContractFilterer) (*VmFilterer, error) {
	contract, err := bindVm(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VmFilterer{contract: contract}, nil
}

// bindVm binds a generic wrapper to an already deployed contract.
func bindVm(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VmMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vm *VmRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vm.Contract.VmCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vm *VmRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.Contract.VmTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vm *VmRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vm.Contract.VmTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vm *VmCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vm.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vm *VmTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vm *VmTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vm.Contract.contract.Transact(opts, method, params...)
}

// ActiveFork is a free data retrieval call binding the contract method 0x2f103f22.
//
// Solidity: function activeFork() view returns(uint256 forkId)
func (_Vm *VmCaller) ActiveFork(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "activeFork")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveFork is a free data retrieval call binding the contract method 0x2f103f22.
//
// Solidity: function activeFork() view returns(uint256 forkId)
func (_Vm *VmSession) ActiveFork() (*big.Int, error) {
	return _Vm.Contract.ActiveFork(&_Vm.CallOpts)
}

// ActiveFork is a free data retrieval call binding the contract method 0x2f103f22.
//
// Solidity: function activeFork() view returns(uint256 forkId)
func (_Vm *VmCallerSession) ActiveFork() (*big.Int, error) {
	return _Vm.Contract.ActiveFork(&_Vm.CallOpts)
}

// Addr is a free data retrieval call binding the contract method 0xffa18649.
//
// Solidity: function addr(uint256 privateKey) pure returns(address keyAddr)
func (_Vm *VmCaller) Addr(opts *bind.CallOpts, privateKey *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "addr", privateKey)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Addr is a free data retrieval call binding the contract method 0xffa18649.
//
// Solidity: function addr(uint256 privateKey) pure returns(address keyAddr)
func (_Vm *VmSession) Addr(privateKey *big.Int) (common.Address, error) {
	return _Vm.Contract.Addr(&_Vm.CallOpts, privateKey)
}

// Addr is a free data retrieval call binding the contract method 0xffa18649.
//
// Solidity: function addr(uint256 privateKey) pure returns(address keyAddr)
func (_Vm *VmCallerSession) Addr(privateKey *big.Int) (common.Address, error) {
	return _Vm.Contract.Addr(&_Vm.CallOpts, privateKey)
}

// AssertApproxEqAbs is a free data retrieval call binding the contract method 0x16d207c6.
//
// Solidity: function assertApproxEqAbs(uint256 left, uint256 right, uint256 maxDelta) pure returns()
func (_Vm *VmCaller) AssertApproxEqAbs(opts *bind.CallOpts, left *big.Int, right *big.Int, maxDelta *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertApproxEqAbs", left, right, maxDelta)

	if err != nil {
		return err
	}

	return err

}

// AssertApproxEqAbs is a free data retrieval call binding the contract method 0x16d207c6.
//
// Solidity: function assertApproxEqAbs(uint256 left, uint256 right, uint256 maxDelta) pure returns()
func (_Vm *VmSession) AssertApproxEqAbs(left *big.Int, right *big.Int, maxDelta *big.Int) error {
	return _Vm.Contract.AssertApproxEqAbs(&_Vm.CallOpts, left, right, maxDelta)
}

// AssertApproxEqAbs is a free data retrieval call binding the contract method 0x16d207c6.
//
// Solidity: function assertApproxEqAbs(uint256 left, uint256 right, uint256 maxDelta) pure returns()
func (_Vm *VmCallerSession) AssertApproxEqAbs(left *big.Int, right *big.Int, maxDelta *big.Int) error {
	return _Vm.Contract.AssertApproxEqAbs(&_Vm.CallOpts, left, right, maxDelta)
}

// AssertApproxEqAbs0 is a free data retrieval call binding the contract method 0x240f839d.
//
// Solidity: function assertApproxEqAbs(int256 left, int256 right, uint256 maxDelta) pure returns()
func (_Vm *VmCaller) AssertApproxEqAbs0(opts *bind.CallOpts, left *big.Int, right *big.Int, maxDelta *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertApproxEqAbs0", left, right, maxDelta)

	if err != nil {
		return err
	}

	return err

}

// AssertApproxEqAbs0 is a free data retrieval call binding the contract method 0x240f839d.
//
// Solidity: function assertApproxEqAbs(int256 left, int256 right, uint256 maxDelta) pure returns()
func (_Vm *VmSession) AssertApproxEqAbs0(left *big.Int, right *big.Int, maxDelta *big.Int) error {
	return _Vm.Contract.AssertApproxEqAbs0(&_Vm.CallOpts, left, right, maxDelta)
}

// AssertApproxEqAbs0 is a free data retrieval call binding the contract method 0x240f839d.
//
// Solidity: function assertApproxEqAbs(int256 left, int256 right, uint256 maxDelta) pure returns()
func (_Vm *VmCallerSession) AssertApproxEqAbs0(left *big.Int, right *big.Int, maxDelta *big.Int) error {
	return _Vm.Contract.AssertApproxEqAbs0(&_Vm.CallOpts, left, right, maxDelta)
}

// AssertApproxEqAbs1 is a free data retrieval call binding the contract method 0x8289e621.
//
// Solidity: function assertApproxEqAbs(int256 left, int256 right, uint256 maxDelta, string error) pure returns()
func (_Vm *VmCaller) AssertApproxEqAbs1(opts *bind.CallOpts, left *big.Int, right *big.Int, maxDelta *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertApproxEqAbs1", left, right, maxDelta, error)

	if err != nil {
		return err
	}

	return err

}

// AssertApproxEqAbs1 is a free data retrieval call binding the contract method 0x8289e621.
//
// Solidity: function assertApproxEqAbs(int256 left, int256 right, uint256 maxDelta, string error) pure returns()
func (_Vm *VmSession) AssertApproxEqAbs1(left *big.Int, right *big.Int, maxDelta *big.Int, error string) error {
	return _Vm.Contract.AssertApproxEqAbs1(&_Vm.CallOpts, left, right, maxDelta, error)
}

// AssertApproxEqAbs1 is a free data retrieval call binding the contract method 0x8289e621.
//
// Solidity: function assertApproxEqAbs(int256 left, int256 right, uint256 maxDelta, string error) pure returns()
func (_Vm *VmCallerSession) AssertApproxEqAbs1(left *big.Int, right *big.Int, maxDelta *big.Int, error string) error {
	return _Vm.Contract.AssertApproxEqAbs1(&_Vm.CallOpts, left, right, maxDelta, error)
}

// AssertApproxEqAbs2 is a free data retrieval call binding the contract method 0xf710b062.
//
// Solidity: function assertApproxEqAbs(uint256 left, uint256 right, uint256 maxDelta, string error) pure returns()
func (_Vm *VmCaller) AssertApproxEqAbs2(opts *bind.CallOpts, left *big.Int, right *big.Int, maxDelta *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertApproxEqAbs2", left, right, maxDelta, error)

	if err != nil {
		return err
	}

	return err

}

// AssertApproxEqAbs2 is a free data retrieval call binding the contract method 0xf710b062.
//
// Solidity: function assertApproxEqAbs(uint256 left, uint256 right, uint256 maxDelta, string error) pure returns()
func (_Vm *VmSession) AssertApproxEqAbs2(left *big.Int, right *big.Int, maxDelta *big.Int, error string) error {
	return _Vm.Contract.AssertApproxEqAbs2(&_Vm.CallOpts, left, right, maxDelta, error)
}

// AssertApproxEqAbs2 is a free data retrieval call binding the contract method 0xf710b062.
//
// Solidity: function assertApproxEqAbs(uint256 left, uint256 right, uint256 maxDelta, string error) pure returns()
func (_Vm *VmCallerSession) AssertApproxEqAbs2(left *big.Int, right *big.Int, maxDelta *big.Int, error string) error {
	return _Vm.Contract.AssertApproxEqAbs2(&_Vm.CallOpts, left, right, maxDelta, error)
}

// AssertApproxEqAbsDecimal is a free data retrieval call binding the contract method 0x045c55ce.
//
// Solidity: function assertApproxEqAbsDecimal(uint256 left, uint256 right, uint256 maxDelta, uint256 decimals) pure returns()
func (_Vm *VmCaller) AssertApproxEqAbsDecimal(opts *bind.CallOpts, left *big.Int, right *big.Int, maxDelta *big.Int, decimals *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertApproxEqAbsDecimal", left, right, maxDelta, decimals)

	if err != nil {
		return err
	}

	return err

}

// AssertApproxEqAbsDecimal is a free data retrieval call binding the contract method 0x045c55ce.
//
// Solidity: function assertApproxEqAbsDecimal(uint256 left, uint256 right, uint256 maxDelta, uint256 decimals) pure returns()
func (_Vm *VmSession) AssertApproxEqAbsDecimal(left *big.Int, right *big.Int, maxDelta *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertApproxEqAbsDecimal(&_Vm.CallOpts, left, right, maxDelta, decimals)
}

// AssertApproxEqAbsDecimal is a free data retrieval call binding the contract method 0x045c55ce.
//
// Solidity: function assertApproxEqAbsDecimal(uint256 left, uint256 right, uint256 maxDelta, uint256 decimals) pure returns()
func (_Vm *VmCallerSession) AssertApproxEqAbsDecimal(left *big.Int, right *big.Int, maxDelta *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertApproxEqAbsDecimal(&_Vm.CallOpts, left, right, maxDelta, decimals)
}

// AssertApproxEqAbsDecimal0 is a free data retrieval call binding the contract method 0x3d5bc8bc.
//
// Solidity: function assertApproxEqAbsDecimal(int256 left, int256 right, uint256 maxDelta, uint256 decimals) pure returns()
func (_Vm *VmCaller) AssertApproxEqAbsDecimal0(opts *bind.CallOpts, left *big.Int, right *big.Int, maxDelta *big.Int, decimals *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertApproxEqAbsDecimal0", left, right, maxDelta, decimals)

	if err != nil {
		return err
	}

	return err

}

// AssertApproxEqAbsDecimal0 is a free data retrieval call binding the contract method 0x3d5bc8bc.
//
// Solidity: function assertApproxEqAbsDecimal(int256 left, int256 right, uint256 maxDelta, uint256 decimals) pure returns()
func (_Vm *VmSession) AssertApproxEqAbsDecimal0(left *big.Int, right *big.Int, maxDelta *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertApproxEqAbsDecimal0(&_Vm.CallOpts, left, right, maxDelta, decimals)
}

// AssertApproxEqAbsDecimal0 is a free data retrieval call binding the contract method 0x3d5bc8bc.
//
// Solidity: function assertApproxEqAbsDecimal(int256 left, int256 right, uint256 maxDelta, uint256 decimals) pure returns()
func (_Vm *VmCallerSession) AssertApproxEqAbsDecimal0(left *big.Int, right *big.Int, maxDelta *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertApproxEqAbsDecimal0(&_Vm.CallOpts, left, right, maxDelta, decimals)
}

// AssertApproxEqAbsDecimal1 is a free data retrieval call binding the contract method 0x60429eb2.
//
// Solidity: function assertApproxEqAbsDecimal(uint256 left, uint256 right, uint256 maxDelta, uint256 decimals, string error) pure returns()
func (_Vm *VmCaller) AssertApproxEqAbsDecimal1(opts *bind.CallOpts, left *big.Int, right *big.Int, maxDelta *big.Int, decimals *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertApproxEqAbsDecimal1", left, right, maxDelta, decimals, error)

	if err != nil {
		return err
	}

	return err

}

// AssertApproxEqAbsDecimal1 is a free data retrieval call binding the contract method 0x60429eb2.
//
// Solidity: function assertApproxEqAbsDecimal(uint256 left, uint256 right, uint256 maxDelta, uint256 decimals, string error) pure returns()
func (_Vm *VmSession) AssertApproxEqAbsDecimal1(left *big.Int, right *big.Int, maxDelta *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertApproxEqAbsDecimal1(&_Vm.CallOpts, left, right, maxDelta, decimals, error)
}

// AssertApproxEqAbsDecimal1 is a free data retrieval call binding the contract method 0x60429eb2.
//
// Solidity: function assertApproxEqAbsDecimal(uint256 left, uint256 right, uint256 maxDelta, uint256 decimals, string error) pure returns()
func (_Vm *VmCallerSession) AssertApproxEqAbsDecimal1(left *big.Int, right *big.Int, maxDelta *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertApproxEqAbsDecimal1(&_Vm.CallOpts, left, right, maxDelta, decimals, error)
}

// AssertApproxEqAbsDecimal2 is a free data retrieval call binding the contract method 0x6a5066d4.
//
// Solidity: function assertApproxEqAbsDecimal(int256 left, int256 right, uint256 maxDelta, uint256 decimals, string error) pure returns()
func (_Vm *VmCaller) AssertApproxEqAbsDecimal2(opts *bind.CallOpts, left *big.Int, right *big.Int, maxDelta *big.Int, decimals *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertApproxEqAbsDecimal2", left, right, maxDelta, decimals, error)

	if err != nil {
		return err
	}

	return err

}

// AssertApproxEqAbsDecimal2 is a free data retrieval call binding the contract method 0x6a5066d4.
//
// Solidity: function assertApproxEqAbsDecimal(int256 left, int256 right, uint256 maxDelta, uint256 decimals, string error) pure returns()
func (_Vm *VmSession) AssertApproxEqAbsDecimal2(left *big.Int, right *big.Int, maxDelta *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertApproxEqAbsDecimal2(&_Vm.CallOpts, left, right, maxDelta, decimals, error)
}

// AssertApproxEqAbsDecimal2 is a free data retrieval call binding the contract method 0x6a5066d4.
//
// Solidity: function assertApproxEqAbsDecimal(int256 left, int256 right, uint256 maxDelta, uint256 decimals, string error) pure returns()
func (_Vm *VmCallerSession) AssertApproxEqAbsDecimal2(left *big.Int, right *big.Int, maxDelta *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertApproxEqAbsDecimal2(&_Vm.CallOpts, left, right, maxDelta, decimals, error)
}

// AssertApproxEqRel is a free data retrieval call binding the contract method 0x1ecb7d33.
//
// Solidity: function assertApproxEqRel(uint256 left, uint256 right, uint256 maxPercentDelta, string error) pure returns()
func (_Vm *VmCaller) AssertApproxEqRel(opts *bind.CallOpts, left *big.Int, right *big.Int, maxPercentDelta *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertApproxEqRel", left, right, maxPercentDelta, error)

	if err != nil {
		return err
	}

	return err

}

// AssertApproxEqRel is a free data retrieval call binding the contract method 0x1ecb7d33.
//
// Solidity: function assertApproxEqRel(uint256 left, uint256 right, uint256 maxPercentDelta, string error) pure returns()
func (_Vm *VmSession) AssertApproxEqRel(left *big.Int, right *big.Int, maxPercentDelta *big.Int, error string) error {
	return _Vm.Contract.AssertApproxEqRel(&_Vm.CallOpts, left, right, maxPercentDelta, error)
}

// AssertApproxEqRel is a free data retrieval call binding the contract method 0x1ecb7d33.
//
// Solidity: function assertApproxEqRel(uint256 left, uint256 right, uint256 maxPercentDelta, string error) pure returns()
func (_Vm *VmCallerSession) AssertApproxEqRel(left *big.Int, right *big.Int, maxPercentDelta *big.Int, error string) error {
	return _Vm.Contract.AssertApproxEqRel(&_Vm.CallOpts, left, right, maxPercentDelta, error)
}

// AssertApproxEqRel0 is a free data retrieval call binding the contract method 0x8cf25ef4.
//
// Solidity: function assertApproxEqRel(uint256 left, uint256 right, uint256 maxPercentDelta) pure returns()
func (_Vm *VmCaller) AssertApproxEqRel0(opts *bind.CallOpts, left *big.Int, right *big.Int, maxPercentDelta *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertApproxEqRel0", left, right, maxPercentDelta)

	if err != nil {
		return err
	}

	return err

}

// AssertApproxEqRel0 is a free data retrieval call binding the contract method 0x8cf25ef4.
//
// Solidity: function assertApproxEqRel(uint256 left, uint256 right, uint256 maxPercentDelta) pure returns()
func (_Vm *VmSession) AssertApproxEqRel0(left *big.Int, right *big.Int, maxPercentDelta *big.Int) error {
	return _Vm.Contract.AssertApproxEqRel0(&_Vm.CallOpts, left, right, maxPercentDelta)
}

// AssertApproxEqRel0 is a free data retrieval call binding the contract method 0x8cf25ef4.
//
// Solidity: function assertApproxEqRel(uint256 left, uint256 right, uint256 maxPercentDelta) pure returns()
func (_Vm *VmCallerSession) AssertApproxEqRel0(left *big.Int, right *big.Int, maxPercentDelta *big.Int) error {
	return _Vm.Contract.AssertApproxEqRel0(&_Vm.CallOpts, left, right, maxPercentDelta)
}

// AssertApproxEqRel1 is a free data retrieval call binding the contract method 0xef277d72.
//
// Solidity: function assertApproxEqRel(int256 left, int256 right, uint256 maxPercentDelta, string error) pure returns()
func (_Vm *VmCaller) AssertApproxEqRel1(opts *bind.CallOpts, left *big.Int, right *big.Int, maxPercentDelta *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertApproxEqRel1", left, right, maxPercentDelta, error)

	if err != nil {
		return err
	}

	return err

}

// AssertApproxEqRel1 is a free data retrieval call binding the contract method 0xef277d72.
//
// Solidity: function assertApproxEqRel(int256 left, int256 right, uint256 maxPercentDelta, string error) pure returns()
func (_Vm *VmSession) AssertApproxEqRel1(left *big.Int, right *big.Int, maxPercentDelta *big.Int, error string) error {
	return _Vm.Contract.AssertApproxEqRel1(&_Vm.CallOpts, left, right, maxPercentDelta, error)
}

// AssertApproxEqRel1 is a free data retrieval call binding the contract method 0xef277d72.
//
// Solidity: function assertApproxEqRel(int256 left, int256 right, uint256 maxPercentDelta, string error) pure returns()
func (_Vm *VmCallerSession) AssertApproxEqRel1(left *big.Int, right *big.Int, maxPercentDelta *big.Int, error string) error {
	return _Vm.Contract.AssertApproxEqRel1(&_Vm.CallOpts, left, right, maxPercentDelta, error)
}

// AssertApproxEqRel2 is a free data retrieval call binding the contract method 0xfea2d14f.
//
// Solidity: function assertApproxEqRel(int256 left, int256 right, uint256 maxPercentDelta) pure returns()
func (_Vm *VmCaller) AssertApproxEqRel2(opts *bind.CallOpts, left *big.Int, right *big.Int, maxPercentDelta *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertApproxEqRel2", left, right, maxPercentDelta)

	if err != nil {
		return err
	}

	return err

}

// AssertApproxEqRel2 is a free data retrieval call binding the contract method 0xfea2d14f.
//
// Solidity: function assertApproxEqRel(int256 left, int256 right, uint256 maxPercentDelta) pure returns()
func (_Vm *VmSession) AssertApproxEqRel2(left *big.Int, right *big.Int, maxPercentDelta *big.Int) error {
	return _Vm.Contract.AssertApproxEqRel2(&_Vm.CallOpts, left, right, maxPercentDelta)
}

// AssertApproxEqRel2 is a free data retrieval call binding the contract method 0xfea2d14f.
//
// Solidity: function assertApproxEqRel(int256 left, int256 right, uint256 maxPercentDelta) pure returns()
func (_Vm *VmCallerSession) AssertApproxEqRel2(left *big.Int, right *big.Int, maxPercentDelta *big.Int) error {
	return _Vm.Contract.AssertApproxEqRel2(&_Vm.CallOpts, left, right, maxPercentDelta)
}

// AssertApproxEqRelDecimal is a free data retrieval call binding the contract method 0x21ed2977.
//
// Solidity: function assertApproxEqRelDecimal(uint256 left, uint256 right, uint256 maxPercentDelta, uint256 decimals) pure returns()
func (_Vm *VmCaller) AssertApproxEqRelDecimal(opts *bind.CallOpts, left *big.Int, right *big.Int, maxPercentDelta *big.Int, decimals *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertApproxEqRelDecimal", left, right, maxPercentDelta, decimals)

	if err != nil {
		return err
	}

	return err

}

// AssertApproxEqRelDecimal is a free data retrieval call binding the contract method 0x21ed2977.
//
// Solidity: function assertApproxEqRelDecimal(uint256 left, uint256 right, uint256 maxPercentDelta, uint256 decimals) pure returns()
func (_Vm *VmSession) AssertApproxEqRelDecimal(left *big.Int, right *big.Int, maxPercentDelta *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertApproxEqRelDecimal(&_Vm.CallOpts, left, right, maxPercentDelta, decimals)
}

// AssertApproxEqRelDecimal is a free data retrieval call binding the contract method 0x21ed2977.
//
// Solidity: function assertApproxEqRelDecimal(uint256 left, uint256 right, uint256 maxPercentDelta, uint256 decimals) pure returns()
func (_Vm *VmCallerSession) AssertApproxEqRelDecimal(left *big.Int, right *big.Int, maxPercentDelta *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertApproxEqRelDecimal(&_Vm.CallOpts, left, right, maxPercentDelta, decimals)
}

// AssertApproxEqRelDecimal0 is a free data retrieval call binding the contract method 0x82d6c8fd.
//
// Solidity: function assertApproxEqRelDecimal(uint256 left, uint256 right, uint256 maxPercentDelta, uint256 decimals, string error) pure returns()
func (_Vm *VmCaller) AssertApproxEqRelDecimal0(opts *bind.CallOpts, left *big.Int, right *big.Int, maxPercentDelta *big.Int, decimals *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertApproxEqRelDecimal0", left, right, maxPercentDelta, decimals, error)

	if err != nil {
		return err
	}

	return err

}

// AssertApproxEqRelDecimal0 is a free data retrieval call binding the contract method 0x82d6c8fd.
//
// Solidity: function assertApproxEqRelDecimal(uint256 left, uint256 right, uint256 maxPercentDelta, uint256 decimals, string error) pure returns()
func (_Vm *VmSession) AssertApproxEqRelDecimal0(left *big.Int, right *big.Int, maxPercentDelta *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertApproxEqRelDecimal0(&_Vm.CallOpts, left, right, maxPercentDelta, decimals, error)
}

// AssertApproxEqRelDecimal0 is a free data retrieval call binding the contract method 0x82d6c8fd.
//
// Solidity: function assertApproxEqRelDecimal(uint256 left, uint256 right, uint256 maxPercentDelta, uint256 decimals, string error) pure returns()
func (_Vm *VmCallerSession) AssertApproxEqRelDecimal0(left *big.Int, right *big.Int, maxPercentDelta *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertApproxEqRelDecimal0(&_Vm.CallOpts, left, right, maxPercentDelta, decimals, error)
}

// AssertApproxEqRelDecimal1 is a free data retrieval call binding the contract method 0xabbf21cc.
//
// Solidity: function assertApproxEqRelDecimal(int256 left, int256 right, uint256 maxPercentDelta, uint256 decimals) pure returns()
func (_Vm *VmCaller) AssertApproxEqRelDecimal1(opts *bind.CallOpts, left *big.Int, right *big.Int, maxPercentDelta *big.Int, decimals *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertApproxEqRelDecimal1", left, right, maxPercentDelta, decimals)

	if err != nil {
		return err
	}

	return err

}

// AssertApproxEqRelDecimal1 is a free data retrieval call binding the contract method 0xabbf21cc.
//
// Solidity: function assertApproxEqRelDecimal(int256 left, int256 right, uint256 maxPercentDelta, uint256 decimals) pure returns()
func (_Vm *VmSession) AssertApproxEqRelDecimal1(left *big.Int, right *big.Int, maxPercentDelta *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertApproxEqRelDecimal1(&_Vm.CallOpts, left, right, maxPercentDelta, decimals)
}

// AssertApproxEqRelDecimal1 is a free data retrieval call binding the contract method 0xabbf21cc.
//
// Solidity: function assertApproxEqRelDecimal(int256 left, int256 right, uint256 maxPercentDelta, uint256 decimals) pure returns()
func (_Vm *VmCallerSession) AssertApproxEqRelDecimal1(left *big.Int, right *big.Int, maxPercentDelta *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertApproxEqRelDecimal1(&_Vm.CallOpts, left, right, maxPercentDelta, decimals)
}

// AssertApproxEqRelDecimal2 is a free data retrieval call binding the contract method 0xfccc11c4.
//
// Solidity: function assertApproxEqRelDecimal(int256 left, int256 right, uint256 maxPercentDelta, uint256 decimals, string error) pure returns()
func (_Vm *VmCaller) AssertApproxEqRelDecimal2(opts *bind.CallOpts, left *big.Int, right *big.Int, maxPercentDelta *big.Int, decimals *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertApproxEqRelDecimal2", left, right, maxPercentDelta, decimals, error)

	if err != nil {
		return err
	}

	return err

}

// AssertApproxEqRelDecimal2 is a free data retrieval call binding the contract method 0xfccc11c4.
//
// Solidity: function assertApproxEqRelDecimal(int256 left, int256 right, uint256 maxPercentDelta, uint256 decimals, string error) pure returns()
func (_Vm *VmSession) AssertApproxEqRelDecimal2(left *big.Int, right *big.Int, maxPercentDelta *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertApproxEqRelDecimal2(&_Vm.CallOpts, left, right, maxPercentDelta, decimals, error)
}

// AssertApproxEqRelDecimal2 is a free data retrieval call binding the contract method 0xfccc11c4.
//
// Solidity: function assertApproxEqRelDecimal(int256 left, int256 right, uint256 maxPercentDelta, uint256 decimals, string error) pure returns()
func (_Vm *VmCallerSession) AssertApproxEqRelDecimal2(left *big.Int, right *big.Int, maxPercentDelta *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertApproxEqRelDecimal2(&_Vm.CallOpts, left, right, maxPercentDelta, decimals, error)
}

// AssertEq is a free data retrieval call binding the contract method 0x0cc9ee84.
//
// Solidity: function assertEq(bytes32[] left, bytes32[] right) pure returns()
func (_Vm *VmCaller) AssertEq(opts *bind.CallOpts, left [][32]byte, right [][32]byte) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertEq is a free data retrieval call binding the contract method 0x0cc9ee84.
//
// Solidity: function assertEq(bytes32[] left, bytes32[] right) pure returns()
func (_Vm *VmSession) AssertEq(left [][32]byte, right [][32]byte) error {
	return _Vm.Contract.AssertEq(&_Vm.CallOpts, left, right)
}

// AssertEq is a free data retrieval call binding the contract method 0x0cc9ee84.
//
// Solidity: function assertEq(bytes32[] left, bytes32[] right) pure returns()
func (_Vm *VmCallerSession) AssertEq(left [][32]byte, right [][32]byte) error {
	return _Vm.Contract.AssertEq(&_Vm.CallOpts, left, right)
}

// AssertEq0 is a free data retrieval call binding the contract method 0x191f1b30.
//
// Solidity: function assertEq(int256[] left, int256[] right, string error) pure returns()
func (_Vm *VmCaller) AssertEq0(opts *bind.CallOpts, left []*big.Int, right []*big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq0", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertEq0 is a free data retrieval call binding the contract method 0x191f1b30.
//
// Solidity: function assertEq(int256[] left, int256[] right, string error) pure returns()
func (_Vm *VmSession) AssertEq0(left []*big.Int, right []*big.Int, error string) error {
	return _Vm.Contract.AssertEq0(&_Vm.CallOpts, left, right, error)
}

// AssertEq0 is a free data retrieval call binding the contract method 0x191f1b30.
//
// Solidity: function assertEq(int256[] left, int256[] right, string error) pure returns()
func (_Vm *VmCallerSession) AssertEq0(left []*big.Int, right []*big.Int, error string) error {
	return _Vm.Contract.AssertEq0(&_Vm.CallOpts, left, right, error)
}

// AssertEq1 is a free data retrieval call binding the contract method 0x2f2769d1.
//
// Solidity: function assertEq(address left, address right, string error) pure returns()
func (_Vm *VmCaller) AssertEq1(opts *bind.CallOpts, left common.Address, right common.Address, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq1", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertEq1 is a free data retrieval call binding the contract method 0x2f2769d1.
//
// Solidity: function assertEq(address left, address right, string error) pure returns()
func (_Vm *VmSession) AssertEq1(left common.Address, right common.Address, error string) error {
	return _Vm.Contract.AssertEq1(&_Vm.CallOpts, left, right, error)
}

// AssertEq1 is a free data retrieval call binding the contract method 0x2f2769d1.
//
// Solidity: function assertEq(address left, address right, string error) pure returns()
func (_Vm *VmCallerSession) AssertEq1(left common.Address, right common.Address, error string) error {
	return _Vm.Contract.AssertEq1(&_Vm.CallOpts, left, right, error)
}

// AssertEq10 is a free data retrieval call binding the contract method 0x714a2f13.
//
// Solidity: function assertEq(int256 left, int256 right, string error) pure returns()
func (_Vm *VmCaller) AssertEq10(opts *bind.CallOpts, left *big.Int, right *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq10", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertEq10 is a free data retrieval call binding the contract method 0x714a2f13.
//
// Solidity: function assertEq(int256 left, int256 right, string error) pure returns()
func (_Vm *VmSession) AssertEq10(left *big.Int, right *big.Int, error string) error {
	return _Vm.Contract.AssertEq10(&_Vm.CallOpts, left, right, error)
}

// AssertEq10 is a free data retrieval call binding the contract method 0x714a2f13.
//
// Solidity: function assertEq(int256 left, int256 right, string error) pure returns()
func (_Vm *VmCallerSession) AssertEq10(left *big.Int, right *big.Int, error string) error {
	return _Vm.Contract.AssertEq10(&_Vm.CallOpts, left, right, error)
}

// AssertEq11 is a free data retrieval call binding the contract method 0x7c84c69b.
//
// Solidity: function assertEq(bytes32 left, bytes32 right) pure returns()
func (_Vm *VmCaller) AssertEq11(opts *bind.CallOpts, left [32]byte, right [32]byte) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq11", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertEq11 is a free data retrieval call binding the contract method 0x7c84c69b.
//
// Solidity: function assertEq(bytes32 left, bytes32 right) pure returns()
func (_Vm *VmSession) AssertEq11(left [32]byte, right [32]byte) error {
	return _Vm.Contract.AssertEq11(&_Vm.CallOpts, left, right)
}

// AssertEq11 is a free data retrieval call binding the contract method 0x7c84c69b.
//
// Solidity: function assertEq(bytes32 left, bytes32 right) pure returns()
func (_Vm *VmCallerSession) AssertEq11(left [32]byte, right [32]byte) error {
	return _Vm.Contract.AssertEq11(&_Vm.CallOpts, left, right)
}

// AssertEq12 is a free data retrieval call binding the contract method 0x88b44c85.
//
// Solidity: function assertEq(uint256 left, uint256 right, string error) pure returns()
func (_Vm *VmCaller) AssertEq12(opts *bind.CallOpts, left *big.Int, right *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq12", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertEq12 is a free data retrieval call binding the contract method 0x88b44c85.
//
// Solidity: function assertEq(uint256 left, uint256 right, string error) pure returns()
func (_Vm *VmSession) AssertEq12(left *big.Int, right *big.Int, error string) error {
	return _Vm.Contract.AssertEq12(&_Vm.CallOpts, left, right, error)
}

// AssertEq12 is a free data retrieval call binding the contract method 0x88b44c85.
//
// Solidity: function assertEq(uint256 left, uint256 right, string error) pure returns()
func (_Vm *VmCallerSession) AssertEq12(left *big.Int, right *big.Int, error string) error {
	return _Vm.Contract.AssertEq12(&_Vm.CallOpts, left, right, error)
}

// AssertEq13 is a free data retrieval call binding the contract method 0x975d5a12.
//
// Solidity: function assertEq(uint256[] left, uint256[] right) pure returns()
func (_Vm *VmCaller) AssertEq13(opts *bind.CallOpts, left []*big.Int, right []*big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq13", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertEq13 is a free data retrieval call binding the contract method 0x975d5a12.
//
// Solidity: function assertEq(uint256[] left, uint256[] right) pure returns()
func (_Vm *VmSession) AssertEq13(left []*big.Int, right []*big.Int) error {
	return _Vm.Contract.AssertEq13(&_Vm.CallOpts, left, right)
}

// AssertEq13 is a free data retrieval call binding the contract method 0x975d5a12.
//
// Solidity: function assertEq(uint256[] left, uint256[] right) pure returns()
func (_Vm *VmCallerSession) AssertEq13(left []*big.Int, right []*big.Int) error {
	return _Vm.Contract.AssertEq13(&_Vm.CallOpts, left, right)
}

// AssertEq14 is a free data retrieval call binding the contract method 0x97624631.
//
// Solidity: function assertEq(bytes left, bytes right) pure returns()
func (_Vm *VmCaller) AssertEq14(opts *bind.CallOpts, left []byte, right []byte) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq14", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertEq14 is a free data retrieval call binding the contract method 0x97624631.
//
// Solidity: function assertEq(bytes left, bytes right) pure returns()
func (_Vm *VmSession) AssertEq14(left []byte, right []byte) error {
	return _Vm.Contract.AssertEq14(&_Vm.CallOpts, left, right)
}

// AssertEq14 is a free data retrieval call binding the contract method 0x97624631.
//
// Solidity: function assertEq(bytes left, bytes right) pure returns()
func (_Vm *VmCallerSession) AssertEq14(left []byte, right []byte) error {
	return _Vm.Contract.AssertEq14(&_Vm.CallOpts, left, right)
}

// AssertEq15 is a free data retrieval call binding the contract method 0x98296c54.
//
// Solidity: function assertEq(uint256 left, uint256 right) pure returns()
func (_Vm *VmCaller) AssertEq15(opts *bind.CallOpts, left *big.Int, right *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq15", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertEq15 is a free data retrieval call binding the contract method 0x98296c54.
//
// Solidity: function assertEq(uint256 left, uint256 right) pure returns()
func (_Vm *VmSession) AssertEq15(left *big.Int, right *big.Int) error {
	return _Vm.Contract.AssertEq15(&_Vm.CallOpts, left, right)
}

// AssertEq15 is a free data retrieval call binding the contract method 0x98296c54.
//
// Solidity: function assertEq(uint256 left, uint256 right) pure returns()
func (_Vm *VmCallerSession) AssertEq15(left *big.Int, right *big.Int) error {
	return _Vm.Contract.AssertEq15(&_Vm.CallOpts, left, right)
}

// AssertEq16 is a free data retrieval call binding the contract method 0xc1fa1ed0.
//
// Solidity: function assertEq(bytes32 left, bytes32 right, string error) pure returns()
func (_Vm *VmCaller) AssertEq16(opts *bind.CallOpts, left [32]byte, right [32]byte, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq16", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertEq16 is a free data retrieval call binding the contract method 0xc1fa1ed0.
//
// Solidity: function assertEq(bytes32 left, bytes32 right, string error) pure returns()
func (_Vm *VmSession) AssertEq16(left [32]byte, right [32]byte, error string) error {
	return _Vm.Contract.AssertEq16(&_Vm.CallOpts, left, right, error)
}

// AssertEq16 is a free data retrieval call binding the contract method 0xc1fa1ed0.
//
// Solidity: function assertEq(bytes32 left, bytes32 right, string error) pure returns()
func (_Vm *VmCallerSession) AssertEq16(left [32]byte, right [32]byte, error string) error {
	return _Vm.Contract.AssertEq16(&_Vm.CallOpts, left, right, error)
}

// AssertEq17 is a free data retrieval call binding the contract method 0xcf1c049c.
//
// Solidity: function assertEq(string[] left, string[] right) pure returns()
func (_Vm *VmCaller) AssertEq17(opts *bind.CallOpts, left []string, right []string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq17", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertEq17 is a free data retrieval call binding the contract method 0xcf1c049c.
//
// Solidity: function assertEq(string[] left, string[] right) pure returns()
func (_Vm *VmSession) AssertEq17(left []string, right []string) error {
	return _Vm.Contract.AssertEq17(&_Vm.CallOpts, left, right)
}

// AssertEq17 is a free data retrieval call binding the contract method 0xcf1c049c.
//
// Solidity: function assertEq(string[] left, string[] right) pure returns()
func (_Vm *VmCallerSession) AssertEq17(left []string, right []string) error {
	return _Vm.Contract.AssertEq17(&_Vm.CallOpts, left, right)
}

// AssertEq18 is a free data retrieval call binding the contract method 0xe03e9177.
//
// Solidity: function assertEq(bytes32[] left, bytes32[] right, string error) pure returns()
func (_Vm *VmCaller) AssertEq18(opts *bind.CallOpts, left [][32]byte, right [][32]byte, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq18", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertEq18 is a free data retrieval call binding the contract method 0xe03e9177.
//
// Solidity: function assertEq(bytes32[] left, bytes32[] right, string error) pure returns()
func (_Vm *VmSession) AssertEq18(left [][32]byte, right [][32]byte, error string) error {
	return _Vm.Contract.AssertEq18(&_Vm.CallOpts, left, right, error)
}

// AssertEq18 is a free data retrieval call binding the contract method 0xe03e9177.
//
// Solidity: function assertEq(bytes32[] left, bytes32[] right, string error) pure returns()
func (_Vm *VmCallerSession) AssertEq18(left [][32]byte, right [][32]byte, error string) error {
	return _Vm.Contract.AssertEq18(&_Vm.CallOpts, left, right, error)
}

// AssertEq19 is a free data retrieval call binding the contract method 0xe24fed00.
//
// Solidity: function assertEq(bytes left, bytes right, string error) pure returns()
func (_Vm *VmCaller) AssertEq19(opts *bind.CallOpts, left []byte, right []byte, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq19", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertEq19 is a free data retrieval call binding the contract method 0xe24fed00.
//
// Solidity: function assertEq(bytes left, bytes right, string error) pure returns()
func (_Vm *VmSession) AssertEq19(left []byte, right []byte, error string) error {
	return _Vm.Contract.AssertEq19(&_Vm.CallOpts, left, right, error)
}

// AssertEq19 is a free data retrieval call binding the contract method 0xe24fed00.
//
// Solidity: function assertEq(bytes left, bytes right, string error) pure returns()
func (_Vm *VmCallerSession) AssertEq19(left []byte, right []byte, error string) error {
	return _Vm.Contract.AssertEq19(&_Vm.CallOpts, left, right, error)
}

// AssertEq2 is a free data retrieval call binding the contract method 0x36f656d8.
//
// Solidity: function assertEq(string left, string right, string error) pure returns()
func (_Vm *VmCaller) AssertEq2(opts *bind.CallOpts, left string, right string, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq2", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertEq2 is a free data retrieval call binding the contract method 0x36f656d8.
//
// Solidity: function assertEq(string left, string right, string error) pure returns()
func (_Vm *VmSession) AssertEq2(left string, right string, error string) error {
	return _Vm.Contract.AssertEq2(&_Vm.CallOpts, left, right, error)
}

// AssertEq2 is a free data retrieval call binding the contract method 0x36f656d8.
//
// Solidity: function assertEq(string left, string right, string error) pure returns()
func (_Vm *VmCallerSession) AssertEq2(left string, right string, error string) error {
	return _Vm.Contract.AssertEq2(&_Vm.CallOpts, left, right, error)
}

// AssertEq20 is a free data retrieval call binding the contract method 0xe48a8f8d.
//
// Solidity: function assertEq(bool[] left, bool[] right, string error) pure returns()
func (_Vm *VmCaller) AssertEq20(opts *bind.CallOpts, left []bool, right []bool, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq20", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertEq20 is a free data retrieval call binding the contract method 0xe48a8f8d.
//
// Solidity: function assertEq(bool[] left, bool[] right, string error) pure returns()
func (_Vm *VmSession) AssertEq20(left []bool, right []bool, error string) error {
	return _Vm.Contract.AssertEq20(&_Vm.CallOpts, left, right, error)
}

// AssertEq20 is a free data retrieval call binding the contract method 0xe48a8f8d.
//
// Solidity: function assertEq(bool[] left, bool[] right, string error) pure returns()
func (_Vm *VmCallerSession) AssertEq20(left []bool, right []bool, error string) error {
	return _Vm.Contract.AssertEq20(&_Vm.CallOpts, left, right, error)
}

// AssertEq21 is a free data retrieval call binding the contract method 0xe5fb9b4a.
//
// Solidity: function assertEq(bytes[] left, bytes[] right) pure returns()
func (_Vm *VmCaller) AssertEq21(opts *bind.CallOpts, left [][]byte, right [][]byte) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq21", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertEq21 is a free data retrieval call binding the contract method 0xe5fb9b4a.
//
// Solidity: function assertEq(bytes[] left, bytes[] right) pure returns()
func (_Vm *VmSession) AssertEq21(left [][]byte, right [][]byte) error {
	return _Vm.Contract.AssertEq21(&_Vm.CallOpts, left, right)
}

// AssertEq21 is a free data retrieval call binding the contract method 0xe5fb9b4a.
//
// Solidity: function assertEq(bytes[] left, bytes[] right) pure returns()
func (_Vm *VmCallerSession) AssertEq21(left [][]byte, right [][]byte) error {
	return _Vm.Contract.AssertEq21(&_Vm.CallOpts, left, right)
}

// AssertEq22 is a free data retrieval call binding the contract method 0xeff6b27d.
//
// Solidity: function assertEq(string[] left, string[] right, string error) pure returns()
func (_Vm *VmCaller) AssertEq22(opts *bind.CallOpts, left []string, right []string, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq22", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertEq22 is a free data retrieval call binding the contract method 0xeff6b27d.
//
// Solidity: function assertEq(string[] left, string[] right, string error) pure returns()
func (_Vm *VmSession) AssertEq22(left []string, right []string, error string) error {
	return _Vm.Contract.AssertEq22(&_Vm.CallOpts, left, right, error)
}

// AssertEq22 is a free data retrieval call binding the contract method 0xeff6b27d.
//
// Solidity: function assertEq(string[] left, string[] right, string error) pure returns()
func (_Vm *VmCallerSession) AssertEq22(left []string, right []string, error string) error {
	return _Vm.Contract.AssertEq22(&_Vm.CallOpts, left, right, error)
}

// AssertEq23 is a free data retrieval call binding the contract method 0xf320d963.
//
// Solidity: function assertEq(string left, string right) pure returns()
func (_Vm *VmCaller) AssertEq23(opts *bind.CallOpts, left string, right string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq23", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertEq23 is a free data retrieval call binding the contract method 0xf320d963.
//
// Solidity: function assertEq(string left, string right) pure returns()
func (_Vm *VmSession) AssertEq23(left string, right string) error {
	return _Vm.Contract.AssertEq23(&_Vm.CallOpts, left, right)
}

// AssertEq23 is a free data retrieval call binding the contract method 0xf320d963.
//
// Solidity: function assertEq(string left, string right) pure returns()
func (_Vm *VmCallerSession) AssertEq23(left string, right string) error {
	return _Vm.Contract.AssertEq23(&_Vm.CallOpts, left, right)
}

// AssertEq24 is a free data retrieval call binding the contract method 0xf413f0b6.
//
// Solidity: function assertEq(bytes[] left, bytes[] right, string error) pure returns()
func (_Vm *VmCaller) AssertEq24(opts *bind.CallOpts, left [][]byte, right [][]byte, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq24", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertEq24 is a free data retrieval call binding the contract method 0xf413f0b6.
//
// Solidity: function assertEq(bytes[] left, bytes[] right, string error) pure returns()
func (_Vm *VmSession) AssertEq24(left [][]byte, right [][]byte, error string) error {
	return _Vm.Contract.AssertEq24(&_Vm.CallOpts, left, right, error)
}

// AssertEq24 is a free data retrieval call binding the contract method 0xf413f0b6.
//
// Solidity: function assertEq(bytes[] left, bytes[] right, string error) pure returns()
func (_Vm *VmCallerSession) AssertEq24(left [][]byte, right [][]byte, error string) error {
	return _Vm.Contract.AssertEq24(&_Vm.CallOpts, left, right, error)
}

// AssertEq25 is a free data retrieval call binding the contract method 0xf7fe3477.
//
// Solidity: function assertEq(bool left, bool right) pure returns()
func (_Vm *VmCaller) AssertEq25(opts *bind.CallOpts, left bool, right bool) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq25", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertEq25 is a free data retrieval call binding the contract method 0xf7fe3477.
//
// Solidity: function assertEq(bool left, bool right) pure returns()
func (_Vm *VmSession) AssertEq25(left bool, right bool) error {
	return _Vm.Contract.AssertEq25(&_Vm.CallOpts, left, right)
}

// AssertEq25 is a free data retrieval call binding the contract method 0xf7fe3477.
//
// Solidity: function assertEq(bool left, bool right) pure returns()
func (_Vm *VmCallerSession) AssertEq25(left bool, right bool) error {
	return _Vm.Contract.AssertEq25(&_Vm.CallOpts, left, right)
}

// AssertEq26 is a free data retrieval call binding the contract method 0xfe74f05b.
//
// Solidity: function assertEq(int256 left, int256 right) pure returns()
func (_Vm *VmCaller) AssertEq26(opts *bind.CallOpts, left *big.Int, right *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq26", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertEq26 is a free data retrieval call binding the contract method 0xfe74f05b.
//
// Solidity: function assertEq(int256 left, int256 right) pure returns()
func (_Vm *VmSession) AssertEq26(left *big.Int, right *big.Int) error {
	return _Vm.Contract.AssertEq26(&_Vm.CallOpts, left, right)
}

// AssertEq26 is a free data retrieval call binding the contract method 0xfe74f05b.
//
// Solidity: function assertEq(int256 left, int256 right) pure returns()
func (_Vm *VmCallerSession) AssertEq26(left *big.Int, right *big.Int) error {
	return _Vm.Contract.AssertEq26(&_Vm.CallOpts, left, right)
}

// AssertEq3 is a free data retrieval call binding the contract method 0x3868ac34.
//
// Solidity: function assertEq(address[] left, address[] right) pure returns()
func (_Vm *VmCaller) AssertEq3(opts *bind.CallOpts, left []common.Address, right []common.Address) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq3", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertEq3 is a free data retrieval call binding the contract method 0x3868ac34.
//
// Solidity: function assertEq(address[] left, address[] right) pure returns()
func (_Vm *VmSession) AssertEq3(left []common.Address, right []common.Address) error {
	return _Vm.Contract.AssertEq3(&_Vm.CallOpts, left, right)
}

// AssertEq3 is a free data retrieval call binding the contract method 0x3868ac34.
//
// Solidity: function assertEq(address[] left, address[] right) pure returns()
func (_Vm *VmCallerSession) AssertEq3(left []common.Address, right []common.Address) error {
	return _Vm.Contract.AssertEq3(&_Vm.CallOpts, left, right)
}

// AssertEq4 is a free data retrieval call binding the contract method 0x3e9173c5.
//
// Solidity: function assertEq(address[] left, address[] right, string error) pure returns()
func (_Vm *VmCaller) AssertEq4(opts *bind.CallOpts, left []common.Address, right []common.Address, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq4", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertEq4 is a free data retrieval call binding the contract method 0x3e9173c5.
//
// Solidity: function assertEq(address[] left, address[] right, string error) pure returns()
func (_Vm *VmSession) AssertEq4(left []common.Address, right []common.Address, error string) error {
	return _Vm.Contract.AssertEq4(&_Vm.CallOpts, left, right, error)
}

// AssertEq4 is a free data retrieval call binding the contract method 0x3e9173c5.
//
// Solidity: function assertEq(address[] left, address[] right, string error) pure returns()
func (_Vm *VmCallerSession) AssertEq4(left []common.Address, right []common.Address, error string) error {
	return _Vm.Contract.AssertEq4(&_Vm.CallOpts, left, right, error)
}

// AssertEq5 is a free data retrieval call binding the contract method 0x4db19e7e.
//
// Solidity: function assertEq(bool left, bool right, string error) pure returns()
func (_Vm *VmCaller) AssertEq5(opts *bind.CallOpts, left bool, right bool, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq5", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertEq5 is a free data retrieval call binding the contract method 0x4db19e7e.
//
// Solidity: function assertEq(bool left, bool right, string error) pure returns()
func (_Vm *VmSession) AssertEq5(left bool, right bool, error string) error {
	return _Vm.Contract.AssertEq5(&_Vm.CallOpts, left, right, error)
}

// AssertEq5 is a free data retrieval call binding the contract method 0x4db19e7e.
//
// Solidity: function assertEq(bool left, bool right, string error) pure returns()
func (_Vm *VmCallerSession) AssertEq5(left bool, right bool, error string) error {
	return _Vm.Contract.AssertEq5(&_Vm.CallOpts, left, right, error)
}

// AssertEq6 is a free data retrieval call binding the contract method 0x515361f6.
//
// Solidity: function assertEq(address left, address right) pure returns()
func (_Vm *VmCaller) AssertEq6(opts *bind.CallOpts, left common.Address, right common.Address) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq6", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertEq6 is a free data retrieval call binding the contract method 0x515361f6.
//
// Solidity: function assertEq(address left, address right) pure returns()
func (_Vm *VmSession) AssertEq6(left common.Address, right common.Address) error {
	return _Vm.Contract.AssertEq6(&_Vm.CallOpts, left, right)
}

// AssertEq6 is a free data retrieval call binding the contract method 0x515361f6.
//
// Solidity: function assertEq(address left, address right) pure returns()
func (_Vm *VmCallerSession) AssertEq6(left common.Address, right common.Address) error {
	return _Vm.Contract.AssertEq6(&_Vm.CallOpts, left, right)
}

// AssertEq7 is a free data retrieval call binding the contract method 0x5d18c73a.
//
// Solidity: function assertEq(uint256[] left, uint256[] right, string error) pure returns()
func (_Vm *VmCaller) AssertEq7(opts *bind.CallOpts, left []*big.Int, right []*big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq7", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertEq7 is a free data retrieval call binding the contract method 0x5d18c73a.
//
// Solidity: function assertEq(uint256[] left, uint256[] right, string error) pure returns()
func (_Vm *VmSession) AssertEq7(left []*big.Int, right []*big.Int, error string) error {
	return _Vm.Contract.AssertEq7(&_Vm.CallOpts, left, right, error)
}

// AssertEq7 is a free data retrieval call binding the contract method 0x5d18c73a.
//
// Solidity: function assertEq(uint256[] left, uint256[] right, string error) pure returns()
func (_Vm *VmCallerSession) AssertEq7(left []*big.Int, right []*big.Int, error string) error {
	return _Vm.Contract.AssertEq7(&_Vm.CallOpts, left, right, error)
}

// AssertEq8 is a free data retrieval call binding the contract method 0x707df785.
//
// Solidity: function assertEq(bool[] left, bool[] right) pure returns()
func (_Vm *VmCaller) AssertEq8(opts *bind.CallOpts, left []bool, right []bool) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq8", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertEq8 is a free data retrieval call binding the contract method 0x707df785.
//
// Solidity: function assertEq(bool[] left, bool[] right) pure returns()
func (_Vm *VmSession) AssertEq8(left []bool, right []bool) error {
	return _Vm.Contract.AssertEq8(&_Vm.CallOpts, left, right)
}

// AssertEq8 is a free data retrieval call binding the contract method 0x707df785.
//
// Solidity: function assertEq(bool[] left, bool[] right) pure returns()
func (_Vm *VmCallerSession) AssertEq8(left []bool, right []bool) error {
	return _Vm.Contract.AssertEq8(&_Vm.CallOpts, left, right)
}

// AssertEq9 is a free data retrieval call binding the contract method 0x711043ac.
//
// Solidity: function assertEq(int256[] left, int256[] right) pure returns()
func (_Vm *VmCaller) AssertEq9(opts *bind.CallOpts, left []*big.Int, right []*big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq9", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertEq9 is a free data retrieval call binding the contract method 0x711043ac.
//
// Solidity: function assertEq(int256[] left, int256[] right) pure returns()
func (_Vm *VmSession) AssertEq9(left []*big.Int, right []*big.Int) error {
	return _Vm.Contract.AssertEq9(&_Vm.CallOpts, left, right)
}

// AssertEq9 is a free data retrieval call binding the contract method 0x711043ac.
//
// Solidity: function assertEq(int256[] left, int256[] right) pure returns()
func (_Vm *VmCallerSession) AssertEq9(left []*big.Int, right []*big.Int) error {
	return _Vm.Contract.AssertEq9(&_Vm.CallOpts, left, right)
}

// AssertEqDecimal is a free data retrieval call binding the contract method 0x27af7d9c.
//
// Solidity: function assertEqDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_Vm *VmCaller) AssertEqDecimal(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEqDecimal", left, right, decimals)

	if err != nil {
		return err
	}

	return err

}

// AssertEqDecimal is a free data retrieval call binding the contract method 0x27af7d9c.
//
// Solidity: function assertEqDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_Vm *VmSession) AssertEqDecimal(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertEqDecimal(&_Vm.CallOpts, left, right, decimals)
}

// AssertEqDecimal is a free data retrieval call binding the contract method 0x27af7d9c.
//
// Solidity: function assertEqDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_Vm *VmCallerSession) AssertEqDecimal(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertEqDecimal(&_Vm.CallOpts, left, right, decimals)
}

// AssertEqDecimal0 is a free data retrieval call binding the contract method 0x48016c04.
//
// Solidity: function assertEqDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_Vm *VmCaller) AssertEqDecimal0(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEqDecimal0", left, right, decimals)

	if err != nil {
		return err
	}

	return err

}

// AssertEqDecimal0 is a free data retrieval call binding the contract method 0x48016c04.
//
// Solidity: function assertEqDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_Vm *VmSession) AssertEqDecimal0(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertEqDecimal0(&_Vm.CallOpts, left, right, decimals)
}

// AssertEqDecimal0 is a free data retrieval call binding the contract method 0x48016c04.
//
// Solidity: function assertEqDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_Vm *VmCallerSession) AssertEqDecimal0(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertEqDecimal0(&_Vm.CallOpts, left, right, decimals)
}

// AssertEqDecimal1 is a free data retrieval call binding the contract method 0x7e77b0c5.
//
// Solidity: function assertEqDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmCaller) AssertEqDecimal1(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEqDecimal1", left, right, decimals, error)

	if err != nil {
		return err
	}

	return err

}

// AssertEqDecimal1 is a free data retrieval call binding the contract method 0x7e77b0c5.
//
// Solidity: function assertEqDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmSession) AssertEqDecimal1(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertEqDecimal1(&_Vm.CallOpts, left, right, decimals, error)
}

// AssertEqDecimal1 is a free data retrieval call binding the contract method 0x7e77b0c5.
//
// Solidity: function assertEqDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmCallerSession) AssertEqDecimal1(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertEqDecimal1(&_Vm.CallOpts, left, right, decimals, error)
}

// AssertEqDecimal2 is a free data retrieval call binding the contract method 0xd0cbbdef.
//
// Solidity: function assertEqDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmCaller) AssertEqDecimal2(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEqDecimal2", left, right, decimals, error)

	if err != nil {
		return err
	}

	return err

}

// AssertEqDecimal2 is a free data retrieval call binding the contract method 0xd0cbbdef.
//
// Solidity: function assertEqDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmSession) AssertEqDecimal2(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertEqDecimal2(&_Vm.CallOpts, left, right, decimals, error)
}

// AssertEqDecimal2 is a free data retrieval call binding the contract method 0xd0cbbdef.
//
// Solidity: function assertEqDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmCallerSession) AssertEqDecimal2(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertEqDecimal2(&_Vm.CallOpts, left, right, decimals, error)
}

// AssertFalse is a free data retrieval call binding the contract method 0x7ba04809.
//
// Solidity: function assertFalse(bool condition, string error) pure returns()
func (_Vm *VmCaller) AssertFalse(opts *bind.CallOpts, condition bool, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertFalse", condition, error)

	if err != nil {
		return err
	}

	return err

}

// AssertFalse is a free data retrieval call binding the contract method 0x7ba04809.
//
// Solidity: function assertFalse(bool condition, string error) pure returns()
func (_Vm *VmSession) AssertFalse(condition bool, error string) error {
	return _Vm.Contract.AssertFalse(&_Vm.CallOpts, condition, error)
}

// AssertFalse is a free data retrieval call binding the contract method 0x7ba04809.
//
// Solidity: function assertFalse(bool condition, string error) pure returns()
func (_Vm *VmCallerSession) AssertFalse(condition bool, error string) error {
	return _Vm.Contract.AssertFalse(&_Vm.CallOpts, condition, error)
}

// AssertFalse0 is a free data retrieval call binding the contract method 0xa5982885.
//
// Solidity: function assertFalse(bool condition) pure returns()
func (_Vm *VmCaller) AssertFalse0(opts *bind.CallOpts, condition bool) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertFalse0", condition)

	if err != nil {
		return err
	}

	return err

}

// AssertFalse0 is a free data retrieval call binding the contract method 0xa5982885.
//
// Solidity: function assertFalse(bool condition) pure returns()
func (_Vm *VmSession) AssertFalse0(condition bool) error {
	return _Vm.Contract.AssertFalse0(&_Vm.CallOpts, condition)
}

// AssertFalse0 is a free data retrieval call binding the contract method 0xa5982885.
//
// Solidity: function assertFalse(bool condition) pure returns()
func (_Vm *VmCallerSession) AssertFalse0(condition bool) error {
	return _Vm.Contract.AssertFalse0(&_Vm.CallOpts, condition)
}

// AssertGe is a free data retrieval call binding the contract method 0x0a30b771.
//
// Solidity: function assertGe(int256 left, int256 right) pure returns()
func (_Vm *VmCaller) AssertGe(opts *bind.CallOpts, left *big.Int, right *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertGe", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertGe is a free data retrieval call binding the contract method 0x0a30b771.
//
// Solidity: function assertGe(int256 left, int256 right) pure returns()
func (_Vm *VmSession) AssertGe(left *big.Int, right *big.Int) error {
	return _Vm.Contract.AssertGe(&_Vm.CallOpts, left, right)
}

// AssertGe is a free data retrieval call binding the contract method 0x0a30b771.
//
// Solidity: function assertGe(int256 left, int256 right) pure returns()
func (_Vm *VmCallerSession) AssertGe(left *big.Int, right *big.Int) error {
	return _Vm.Contract.AssertGe(&_Vm.CallOpts, left, right)
}

// AssertGe0 is a free data retrieval call binding the contract method 0xa84328dd.
//
// Solidity: function assertGe(int256 left, int256 right, string error) pure returns()
func (_Vm *VmCaller) AssertGe0(opts *bind.CallOpts, left *big.Int, right *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertGe0", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertGe0 is a free data retrieval call binding the contract method 0xa84328dd.
//
// Solidity: function assertGe(int256 left, int256 right, string error) pure returns()
func (_Vm *VmSession) AssertGe0(left *big.Int, right *big.Int, error string) error {
	return _Vm.Contract.AssertGe0(&_Vm.CallOpts, left, right, error)
}

// AssertGe0 is a free data retrieval call binding the contract method 0xa84328dd.
//
// Solidity: function assertGe(int256 left, int256 right, string error) pure returns()
func (_Vm *VmCallerSession) AssertGe0(left *big.Int, right *big.Int, error string) error {
	return _Vm.Contract.AssertGe0(&_Vm.CallOpts, left, right, error)
}

// AssertGe1 is a free data retrieval call binding the contract method 0xa8d4d1d9.
//
// Solidity: function assertGe(uint256 left, uint256 right) pure returns()
func (_Vm *VmCaller) AssertGe1(opts *bind.CallOpts, left *big.Int, right *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertGe1", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertGe1 is a free data retrieval call binding the contract method 0xa8d4d1d9.
//
// Solidity: function assertGe(uint256 left, uint256 right) pure returns()
func (_Vm *VmSession) AssertGe1(left *big.Int, right *big.Int) error {
	return _Vm.Contract.AssertGe1(&_Vm.CallOpts, left, right)
}

// AssertGe1 is a free data retrieval call binding the contract method 0xa8d4d1d9.
//
// Solidity: function assertGe(uint256 left, uint256 right) pure returns()
func (_Vm *VmCallerSession) AssertGe1(left *big.Int, right *big.Int) error {
	return _Vm.Contract.AssertGe1(&_Vm.CallOpts, left, right)
}

// AssertGe2 is a free data retrieval call binding the contract method 0xe25242c0.
//
// Solidity: function assertGe(uint256 left, uint256 right, string error) pure returns()
func (_Vm *VmCaller) AssertGe2(opts *bind.CallOpts, left *big.Int, right *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertGe2", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertGe2 is a free data retrieval call binding the contract method 0xe25242c0.
//
// Solidity: function assertGe(uint256 left, uint256 right, string error) pure returns()
func (_Vm *VmSession) AssertGe2(left *big.Int, right *big.Int, error string) error {
	return _Vm.Contract.AssertGe2(&_Vm.CallOpts, left, right, error)
}

// AssertGe2 is a free data retrieval call binding the contract method 0xe25242c0.
//
// Solidity: function assertGe(uint256 left, uint256 right, string error) pure returns()
func (_Vm *VmCallerSession) AssertGe2(left *big.Int, right *big.Int, error string) error {
	return _Vm.Contract.AssertGe2(&_Vm.CallOpts, left, right, error)
}

// AssertGeDecimal is a free data retrieval call binding the contract method 0x3d1fe08a.
//
// Solidity: function assertGeDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_Vm *VmCaller) AssertGeDecimal(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertGeDecimal", left, right, decimals)

	if err != nil {
		return err
	}

	return err

}

// AssertGeDecimal is a free data retrieval call binding the contract method 0x3d1fe08a.
//
// Solidity: function assertGeDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_Vm *VmSession) AssertGeDecimal(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertGeDecimal(&_Vm.CallOpts, left, right, decimals)
}

// AssertGeDecimal is a free data retrieval call binding the contract method 0x3d1fe08a.
//
// Solidity: function assertGeDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_Vm *VmCallerSession) AssertGeDecimal(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertGeDecimal(&_Vm.CallOpts, left, right, decimals)
}

// AssertGeDecimal0 is a free data retrieval call binding the contract method 0x5df93c9b.
//
// Solidity: function assertGeDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmCaller) AssertGeDecimal0(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertGeDecimal0", left, right, decimals, error)

	if err != nil {
		return err
	}

	return err

}

// AssertGeDecimal0 is a free data retrieval call binding the contract method 0x5df93c9b.
//
// Solidity: function assertGeDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmSession) AssertGeDecimal0(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertGeDecimal0(&_Vm.CallOpts, left, right, decimals, error)
}

// AssertGeDecimal0 is a free data retrieval call binding the contract method 0x5df93c9b.
//
// Solidity: function assertGeDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmCallerSession) AssertGeDecimal0(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertGeDecimal0(&_Vm.CallOpts, left, right, decimals, error)
}

// AssertGeDecimal1 is a free data retrieval call binding the contract method 0x8bff9133.
//
// Solidity: function assertGeDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmCaller) AssertGeDecimal1(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertGeDecimal1", left, right, decimals, error)

	if err != nil {
		return err
	}

	return err

}

// AssertGeDecimal1 is a free data retrieval call binding the contract method 0x8bff9133.
//
// Solidity: function assertGeDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmSession) AssertGeDecimal1(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertGeDecimal1(&_Vm.CallOpts, left, right, decimals, error)
}

// AssertGeDecimal1 is a free data retrieval call binding the contract method 0x8bff9133.
//
// Solidity: function assertGeDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmCallerSession) AssertGeDecimal1(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertGeDecimal1(&_Vm.CallOpts, left, right, decimals, error)
}

// AssertGeDecimal2 is a free data retrieval call binding the contract method 0xdc28c0f1.
//
// Solidity: function assertGeDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_Vm *VmCaller) AssertGeDecimal2(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertGeDecimal2", left, right, decimals)

	if err != nil {
		return err
	}

	return err

}

// AssertGeDecimal2 is a free data retrieval call binding the contract method 0xdc28c0f1.
//
// Solidity: function assertGeDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_Vm *VmSession) AssertGeDecimal2(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertGeDecimal2(&_Vm.CallOpts, left, right, decimals)
}

// AssertGeDecimal2 is a free data retrieval call binding the contract method 0xdc28c0f1.
//
// Solidity: function assertGeDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_Vm *VmCallerSession) AssertGeDecimal2(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertGeDecimal2(&_Vm.CallOpts, left, right, decimals)
}

// AssertGt is a free data retrieval call binding the contract method 0x5a362d45.
//
// Solidity: function assertGt(int256 left, int256 right) pure returns()
func (_Vm *VmCaller) AssertGt(opts *bind.CallOpts, left *big.Int, right *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertGt", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertGt is a free data retrieval call binding the contract method 0x5a362d45.
//
// Solidity: function assertGt(int256 left, int256 right) pure returns()
func (_Vm *VmSession) AssertGt(left *big.Int, right *big.Int) error {
	return _Vm.Contract.AssertGt(&_Vm.CallOpts, left, right)
}

// AssertGt is a free data retrieval call binding the contract method 0x5a362d45.
//
// Solidity: function assertGt(int256 left, int256 right) pure returns()
func (_Vm *VmCallerSession) AssertGt(left *big.Int, right *big.Int) error {
	return _Vm.Contract.AssertGt(&_Vm.CallOpts, left, right)
}

// AssertGt0 is a free data retrieval call binding the contract method 0xd9a3c4d2.
//
// Solidity: function assertGt(uint256 left, uint256 right, string error) pure returns()
func (_Vm *VmCaller) AssertGt0(opts *bind.CallOpts, left *big.Int, right *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertGt0", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertGt0 is a free data retrieval call binding the contract method 0xd9a3c4d2.
//
// Solidity: function assertGt(uint256 left, uint256 right, string error) pure returns()
func (_Vm *VmSession) AssertGt0(left *big.Int, right *big.Int, error string) error {
	return _Vm.Contract.AssertGt0(&_Vm.CallOpts, left, right, error)
}

// AssertGt0 is a free data retrieval call binding the contract method 0xd9a3c4d2.
//
// Solidity: function assertGt(uint256 left, uint256 right, string error) pure returns()
func (_Vm *VmCallerSession) AssertGt0(left *big.Int, right *big.Int, error string) error {
	return _Vm.Contract.AssertGt0(&_Vm.CallOpts, left, right, error)
}

// AssertGt1 is a free data retrieval call binding the contract method 0xdb07fcd2.
//
// Solidity: function assertGt(uint256 left, uint256 right) pure returns()
func (_Vm *VmCaller) AssertGt1(opts *bind.CallOpts, left *big.Int, right *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertGt1", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertGt1 is a free data retrieval call binding the contract method 0xdb07fcd2.
//
// Solidity: function assertGt(uint256 left, uint256 right) pure returns()
func (_Vm *VmSession) AssertGt1(left *big.Int, right *big.Int) error {
	return _Vm.Contract.AssertGt1(&_Vm.CallOpts, left, right)
}

// AssertGt1 is a free data retrieval call binding the contract method 0xdb07fcd2.
//
// Solidity: function assertGt(uint256 left, uint256 right) pure returns()
func (_Vm *VmCallerSession) AssertGt1(left *big.Int, right *big.Int) error {
	return _Vm.Contract.AssertGt1(&_Vm.CallOpts, left, right)
}

// AssertGt2 is a free data retrieval call binding the contract method 0xf8d33b9b.
//
// Solidity: function assertGt(int256 left, int256 right, string error) pure returns()
func (_Vm *VmCaller) AssertGt2(opts *bind.CallOpts, left *big.Int, right *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertGt2", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertGt2 is a free data retrieval call binding the contract method 0xf8d33b9b.
//
// Solidity: function assertGt(int256 left, int256 right, string error) pure returns()
func (_Vm *VmSession) AssertGt2(left *big.Int, right *big.Int, error string) error {
	return _Vm.Contract.AssertGt2(&_Vm.CallOpts, left, right, error)
}

// AssertGt2 is a free data retrieval call binding the contract method 0xf8d33b9b.
//
// Solidity: function assertGt(int256 left, int256 right, string error) pure returns()
func (_Vm *VmCallerSession) AssertGt2(left *big.Int, right *big.Int, error string) error {
	return _Vm.Contract.AssertGt2(&_Vm.CallOpts, left, right, error)
}

// AssertGtDecimal is a free data retrieval call binding the contract method 0x04a5c7ab.
//
// Solidity: function assertGtDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmCaller) AssertGtDecimal(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertGtDecimal", left, right, decimals, error)

	if err != nil {
		return err
	}

	return err

}

// AssertGtDecimal is a free data retrieval call binding the contract method 0x04a5c7ab.
//
// Solidity: function assertGtDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmSession) AssertGtDecimal(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertGtDecimal(&_Vm.CallOpts, left, right, decimals, error)
}

// AssertGtDecimal is a free data retrieval call binding the contract method 0x04a5c7ab.
//
// Solidity: function assertGtDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmCallerSession) AssertGtDecimal(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertGtDecimal(&_Vm.CallOpts, left, right, decimals, error)
}

// AssertGtDecimal0 is a free data retrieval call binding the contract method 0x64949a8d.
//
// Solidity: function assertGtDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmCaller) AssertGtDecimal0(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertGtDecimal0", left, right, decimals, error)

	if err != nil {
		return err
	}

	return err

}

// AssertGtDecimal0 is a free data retrieval call binding the contract method 0x64949a8d.
//
// Solidity: function assertGtDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmSession) AssertGtDecimal0(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertGtDecimal0(&_Vm.CallOpts, left, right, decimals, error)
}

// AssertGtDecimal0 is a free data retrieval call binding the contract method 0x64949a8d.
//
// Solidity: function assertGtDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmCallerSession) AssertGtDecimal0(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertGtDecimal0(&_Vm.CallOpts, left, right, decimals, error)
}

// AssertGtDecimal1 is a free data retrieval call binding the contract method 0x78611f0e.
//
// Solidity: function assertGtDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_Vm *VmCaller) AssertGtDecimal1(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertGtDecimal1", left, right, decimals)

	if err != nil {
		return err
	}

	return err

}

// AssertGtDecimal1 is a free data retrieval call binding the contract method 0x78611f0e.
//
// Solidity: function assertGtDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_Vm *VmSession) AssertGtDecimal1(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertGtDecimal1(&_Vm.CallOpts, left, right, decimals)
}

// AssertGtDecimal1 is a free data retrieval call binding the contract method 0x78611f0e.
//
// Solidity: function assertGtDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_Vm *VmCallerSession) AssertGtDecimal1(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertGtDecimal1(&_Vm.CallOpts, left, right, decimals)
}

// AssertGtDecimal2 is a free data retrieval call binding the contract method 0xeccd2437.
//
// Solidity: function assertGtDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_Vm *VmCaller) AssertGtDecimal2(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertGtDecimal2", left, right, decimals)

	if err != nil {
		return err
	}

	return err

}

// AssertGtDecimal2 is a free data retrieval call binding the contract method 0xeccd2437.
//
// Solidity: function assertGtDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_Vm *VmSession) AssertGtDecimal2(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertGtDecimal2(&_Vm.CallOpts, left, right, decimals)
}

// AssertGtDecimal2 is a free data retrieval call binding the contract method 0xeccd2437.
//
// Solidity: function assertGtDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_Vm *VmCallerSession) AssertGtDecimal2(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertGtDecimal2(&_Vm.CallOpts, left, right, decimals)
}

// AssertLe is a free data retrieval call binding the contract method 0x4dfe692c.
//
// Solidity: function assertLe(int256 left, int256 right, string error) pure returns()
func (_Vm *VmCaller) AssertLe(opts *bind.CallOpts, left *big.Int, right *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertLe", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertLe is a free data retrieval call binding the contract method 0x4dfe692c.
//
// Solidity: function assertLe(int256 left, int256 right, string error) pure returns()
func (_Vm *VmSession) AssertLe(left *big.Int, right *big.Int, error string) error {
	return _Vm.Contract.AssertLe(&_Vm.CallOpts, left, right, error)
}

// AssertLe is a free data retrieval call binding the contract method 0x4dfe692c.
//
// Solidity: function assertLe(int256 left, int256 right, string error) pure returns()
func (_Vm *VmCallerSession) AssertLe(left *big.Int, right *big.Int, error string) error {
	return _Vm.Contract.AssertLe(&_Vm.CallOpts, left, right, error)
}

// AssertLe0 is a free data retrieval call binding the contract method 0x8466f415.
//
// Solidity: function assertLe(uint256 left, uint256 right) pure returns()
func (_Vm *VmCaller) AssertLe0(opts *bind.CallOpts, left *big.Int, right *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertLe0", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertLe0 is a free data retrieval call binding the contract method 0x8466f415.
//
// Solidity: function assertLe(uint256 left, uint256 right) pure returns()
func (_Vm *VmSession) AssertLe0(left *big.Int, right *big.Int) error {
	return _Vm.Contract.AssertLe0(&_Vm.CallOpts, left, right)
}

// AssertLe0 is a free data retrieval call binding the contract method 0x8466f415.
//
// Solidity: function assertLe(uint256 left, uint256 right) pure returns()
func (_Vm *VmCallerSession) AssertLe0(left *big.Int, right *big.Int) error {
	return _Vm.Contract.AssertLe0(&_Vm.CallOpts, left, right)
}

// AssertLe1 is a free data retrieval call binding the contract method 0x95fd154e.
//
// Solidity: function assertLe(int256 left, int256 right) pure returns()
func (_Vm *VmCaller) AssertLe1(opts *bind.CallOpts, left *big.Int, right *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertLe1", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertLe1 is a free data retrieval call binding the contract method 0x95fd154e.
//
// Solidity: function assertLe(int256 left, int256 right) pure returns()
func (_Vm *VmSession) AssertLe1(left *big.Int, right *big.Int) error {
	return _Vm.Contract.AssertLe1(&_Vm.CallOpts, left, right)
}

// AssertLe1 is a free data retrieval call binding the contract method 0x95fd154e.
//
// Solidity: function assertLe(int256 left, int256 right) pure returns()
func (_Vm *VmCallerSession) AssertLe1(left *big.Int, right *big.Int) error {
	return _Vm.Contract.AssertLe1(&_Vm.CallOpts, left, right)
}

// AssertLe2 is a free data retrieval call binding the contract method 0xd17d4b0d.
//
// Solidity: function assertLe(uint256 left, uint256 right, string error) pure returns()
func (_Vm *VmCaller) AssertLe2(opts *bind.CallOpts, left *big.Int, right *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertLe2", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertLe2 is a free data retrieval call binding the contract method 0xd17d4b0d.
//
// Solidity: function assertLe(uint256 left, uint256 right, string error) pure returns()
func (_Vm *VmSession) AssertLe2(left *big.Int, right *big.Int, error string) error {
	return _Vm.Contract.AssertLe2(&_Vm.CallOpts, left, right, error)
}

// AssertLe2 is a free data retrieval call binding the contract method 0xd17d4b0d.
//
// Solidity: function assertLe(uint256 left, uint256 right, string error) pure returns()
func (_Vm *VmCallerSession) AssertLe2(left *big.Int, right *big.Int, error string) error {
	return _Vm.Contract.AssertLe2(&_Vm.CallOpts, left, right, error)
}

// AssertLeDecimal is a free data retrieval call binding the contract method 0x11d1364a.
//
// Solidity: function assertLeDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_Vm *VmCaller) AssertLeDecimal(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertLeDecimal", left, right, decimals)

	if err != nil {
		return err
	}

	return err

}

// AssertLeDecimal is a free data retrieval call binding the contract method 0x11d1364a.
//
// Solidity: function assertLeDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_Vm *VmSession) AssertLeDecimal(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertLeDecimal(&_Vm.CallOpts, left, right, decimals)
}

// AssertLeDecimal is a free data retrieval call binding the contract method 0x11d1364a.
//
// Solidity: function assertLeDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_Vm *VmCallerSession) AssertLeDecimal(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertLeDecimal(&_Vm.CallOpts, left, right, decimals)
}

// AssertLeDecimal0 is a free data retrieval call binding the contract method 0x7fefbbe0.
//
// Solidity: function assertLeDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmCaller) AssertLeDecimal0(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertLeDecimal0", left, right, decimals, error)

	if err != nil {
		return err
	}

	return err

}

// AssertLeDecimal0 is a free data retrieval call binding the contract method 0x7fefbbe0.
//
// Solidity: function assertLeDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmSession) AssertLeDecimal0(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertLeDecimal0(&_Vm.CallOpts, left, right, decimals, error)
}

// AssertLeDecimal0 is a free data retrieval call binding the contract method 0x7fefbbe0.
//
// Solidity: function assertLeDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmCallerSession) AssertLeDecimal0(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertLeDecimal0(&_Vm.CallOpts, left, right, decimals, error)
}

// AssertLeDecimal1 is a free data retrieval call binding the contract method 0xaa5cf788.
//
// Solidity: function assertLeDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmCaller) AssertLeDecimal1(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertLeDecimal1", left, right, decimals, error)

	if err != nil {
		return err
	}

	return err

}

// AssertLeDecimal1 is a free data retrieval call binding the contract method 0xaa5cf788.
//
// Solidity: function assertLeDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmSession) AssertLeDecimal1(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertLeDecimal1(&_Vm.CallOpts, left, right, decimals, error)
}

// AssertLeDecimal1 is a free data retrieval call binding the contract method 0xaa5cf788.
//
// Solidity: function assertLeDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmCallerSession) AssertLeDecimal1(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertLeDecimal1(&_Vm.CallOpts, left, right, decimals, error)
}

// AssertLeDecimal2 is a free data retrieval call binding the contract method 0xc304aab7.
//
// Solidity: function assertLeDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_Vm *VmCaller) AssertLeDecimal2(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertLeDecimal2", left, right, decimals)

	if err != nil {
		return err
	}

	return err

}

// AssertLeDecimal2 is a free data retrieval call binding the contract method 0xc304aab7.
//
// Solidity: function assertLeDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_Vm *VmSession) AssertLeDecimal2(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertLeDecimal2(&_Vm.CallOpts, left, right, decimals)
}

// AssertLeDecimal2 is a free data retrieval call binding the contract method 0xc304aab7.
//
// Solidity: function assertLeDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_Vm *VmCallerSession) AssertLeDecimal2(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertLeDecimal2(&_Vm.CallOpts, left, right, decimals)
}

// AssertLt is a free data retrieval call binding the contract method 0x3e914080.
//
// Solidity: function assertLt(int256 left, int256 right) pure returns()
func (_Vm *VmCaller) AssertLt(opts *bind.CallOpts, left *big.Int, right *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertLt", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertLt is a free data retrieval call binding the contract method 0x3e914080.
//
// Solidity: function assertLt(int256 left, int256 right) pure returns()
func (_Vm *VmSession) AssertLt(left *big.Int, right *big.Int) error {
	return _Vm.Contract.AssertLt(&_Vm.CallOpts, left, right)
}

// AssertLt is a free data retrieval call binding the contract method 0x3e914080.
//
// Solidity: function assertLt(int256 left, int256 right) pure returns()
func (_Vm *VmCallerSession) AssertLt(left *big.Int, right *big.Int) error {
	return _Vm.Contract.AssertLt(&_Vm.CallOpts, left, right)
}

// AssertLt0 is a free data retrieval call binding the contract method 0x65d5c135.
//
// Solidity: function assertLt(uint256 left, uint256 right, string error) pure returns()
func (_Vm *VmCaller) AssertLt0(opts *bind.CallOpts, left *big.Int, right *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertLt0", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertLt0 is a free data retrieval call binding the contract method 0x65d5c135.
//
// Solidity: function assertLt(uint256 left, uint256 right, string error) pure returns()
func (_Vm *VmSession) AssertLt0(left *big.Int, right *big.Int, error string) error {
	return _Vm.Contract.AssertLt0(&_Vm.CallOpts, left, right, error)
}

// AssertLt0 is a free data retrieval call binding the contract method 0x65d5c135.
//
// Solidity: function assertLt(uint256 left, uint256 right, string error) pure returns()
func (_Vm *VmCallerSession) AssertLt0(left *big.Int, right *big.Int, error string) error {
	return _Vm.Contract.AssertLt0(&_Vm.CallOpts, left, right, error)
}

// AssertLt1 is a free data retrieval call binding the contract method 0x9ff531e3.
//
// Solidity: function assertLt(int256 left, int256 right, string error) pure returns()
func (_Vm *VmCaller) AssertLt1(opts *bind.CallOpts, left *big.Int, right *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertLt1", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertLt1 is a free data retrieval call binding the contract method 0x9ff531e3.
//
// Solidity: function assertLt(int256 left, int256 right, string error) pure returns()
func (_Vm *VmSession) AssertLt1(left *big.Int, right *big.Int, error string) error {
	return _Vm.Contract.AssertLt1(&_Vm.CallOpts, left, right, error)
}

// AssertLt1 is a free data retrieval call binding the contract method 0x9ff531e3.
//
// Solidity: function assertLt(int256 left, int256 right, string error) pure returns()
func (_Vm *VmCallerSession) AssertLt1(left *big.Int, right *big.Int, error string) error {
	return _Vm.Contract.AssertLt1(&_Vm.CallOpts, left, right, error)
}

// AssertLt2 is a free data retrieval call binding the contract method 0xb12fc005.
//
// Solidity: function assertLt(uint256 left, uint256 right) pure returns()
func (_Vm *VmCaller) AssertLt2(opts *bind.CallOpts, left *big.Int, right *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertLt2", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertLt2 is a free data retrieval call binding the contract method 0xb12fc005.
//
// Solidity: function assertLt(uint256 left, uint256 right) pure returns()
func (_Vm *VmSession) AssertLt2(left *big.Int, right *big.Int) error {
	return _Vm.Contract.AssertLt2(&_Vm.CallOpts, left, right)
}

// AssertLt2 is a free data retrieval call binding the contract method 0xb12fc005.
//
// Solidity: function assertLt(uint256 left, uint256 right) pure returns()
func (_Vm *VmCallerSession) AssertLt2(left *big.Int, right *big.Int) error {
	return _Vm.Contract.AssertLt2(&_Vm.CallOpts, left, right)
}

// AssertLtDecimal is a free data retrieval call binding the contract method 0x2077337e.
//
// Solidity: function assertLtDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_Vm *VmCaller) AssertLtDecimal(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertLtDecimal", left, right, decimals)

	if err != nil {
		return err
	}

	return err

}

// AssertLtDecimal is a free data retrieval call binding the contract method 0x2077337e.
//
// Solidity: function assertLtDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_Vm *VmSession) AssertLtDecimal(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertLtDecimal(&_Vm.CallOpts, left, right, decimals)
}

// AssertLtDecimal is a free data retrieval call binding the contract method 0x2077337e.
//
// Solidity: function assertLtDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_Vm *VmCallerSession) AssertLtDecimal(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertLtDecimal(&_Vm.CallOpts, left, right, decimals)
}

// AssertLtDecimal0 is a free data retrieval call binding the contract method 0x40f0b4e0.
//
// Solidity: function assertLtDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmCaller) AssertLtDecimal0(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertLtDecimal0", left, right, decimals, error)

	if err != nil {
		return err
	}

	return err

}

// AssertLtDecimal0 is a free data retrieval call binding the contract method 0x40f0b4e0.
//
// Solidity: function assertLtDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmSession) AssertLtDecimal0(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertLtDecimal0(&_Vm.CallOpts, left, right, decimals, error)
}

// AssertLtDecimal0 is a free data retrieval call binding the contract method 0x40f0b4e0.
//
// Solidity: function assertLtDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmCallerSession) AssertLtDecimal0(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertLtDecimal0(&_Vm.CallOpts, left, right, decimals, error)
}

// AssertLtDecimal1 is a free data retrieval call binding the contract method 0xa972d037.
//
// Solidity: function assertLtDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmCaller) AssertLtDecimal1(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertLtDecimal1", left, right, decimals, error)

	if err != nil {
		return err
	}

	return err

}

// AssertLtDecimal1 is a free data retrieval call binding the contract method 0xa972d037.
//
// Solidity: function assertLtDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmSession) AssertLtDecimal1(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertLtDecimal1(&_Vm.CallOpts, left, right, decimals, error)
}

// AssertLtDecimal1 is a free data retrieval call binding the contract method 0xa972d037.
//
// Solidity: function assertLtDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmCallerSession) AssertLtDecimal1(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertLtDecimal1(&_Vm.CallOpts, left, right, decimals, error)
}

// AssertLtDecimal2 is a free data retrieval call binding the contract method 0xdbe8d88b.
//
// Solidity: function assertLtDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_Vm *VmCaller) AssertLtDecimal2(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertLtDecimal2", left, right, decimals)

	if err != nil {
		return err
	}

	return err

}

// AssertLtDecimal2 is a free data retrieval call binding the contract method 0xdbe8d88b.
//
// Solidity: function assertLtDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_Vm *VmSession) AssertLtDecimal2(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertLtDecimal2(&_Vm.CallOpts, left, right, decimals)
}

// AssertLtDecimal2 is a free data retrieval call binding the contract method 0xdbe8d88b.
//
// Solidity: function assertLtDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_Vm *VmCallerSession) AssertLtDecimal2(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertLtDecimal2(&_Vm.CallOpts, left, right, decimals)
}

// AssertNotEq is a free data retrieval call binding the contract method 0x0603ea68.
//
// Solidity: function assertNotEq(bytes32[] left, bytes32[] right) pure returns()
func (_Vm *VmCaller) AssertNotEq(opts *bind.CallOpts, left [][32]byte, right [][32]byte) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq is a free data retrieval call binding the contract method 0x0603ea68.
//
// Solidity: function assertNotEq(bytes32[] left, bytes32[] right) pure returns()
func (_Vm *VmSession) AssertNotEq(left [][32]byte, right [][32]byte) error {
	return _Vm.Contract.AssertNotEq(&_Vm.CallOpts, left, right)
}

// AssertNotEq is a free data retrieval call binding the contract method 0x0603ea68.
//
// Solidity: function assertNotEq(bytes32[] left, bytes32[] right) pure returns()
func (_Vm *VmCallerSession) AssertNotEq(left [][32]byte, right [][32]byte) error {
	return _Vm.Contract.AssertNotEq(&_Vm.CallOpts, left, right)
}

// AssertNotEq0 is a free data retrieval call binding the contract method 0x0b72f4ef.
//
// Solidity: function assertNotEq(int256[] left, int256[] right) pure returns()
func (_Vm *VmCaller) AssertNotEq0(opts *bind.CallOpts, left []*big.Int, right []*big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq0", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq0 is a free data retrieval call binding the contract method 0x0b72f4ef.
//
// Solidity: function assertNotEq(int256[] left, int256[] right) pure returns()
func (_Vm *VmSession) AssertNotEq0(left []*big.Int, right []*big.Int) error {
	return _Vm.Contract.AssertNotEq0(&_Vm.CallOpts, left, right)
}

// AssertNotEq0 is a free data retrieval call binding the contract method 0x0b72f4ef.
//
// Solidity: function assertNotEq(int256[] left, int256[] right) pure returns()
func (_Vm *VmCallerSession) AssertNotEq0(left []*big.Int, right []*big.Int) error {
	return _Vm.Contract.AssertNotEq0(&_Vm.CallOpts, left, right)
}

// AssertNotEq1 is a free data retrieval call binding the contract method 0x1091a261.
//
// Solidity: function assertNotEq(bool left, bool right, string error) pure returns()
func (_Vm *VmCaller) AssertNotEq1(opts *bind.CallOpts, left bool, right bool, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq1", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq1 is a free data retrieval call binding the contract method 0x1091a261.
//
// Solidity: function assertNotEq(bool left, bool right, string error) pure returns()
func (_Vm *VmSession) AssertNotEq1(left bool, right bool, error string) error {
	return _Vm.Contract.AssertNotEq1(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq1 is a free data retrieval call binding the contract method 0x1091a261.
//
// Solidity: function assertNotEq(bool left, bool right, string error) pure returns()
func (_Vm *VmCallerSession) AssertNotEq1(left bool, right bool, error string) error {
	return _Vm.Contract.AssertNotEq1(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq10 is a free data retrieval call binding the contract method 0x6a8237b3.
//
// Solidity: function assertNotEq(string left, string right) pure returns()
func (_Vm *VmCaller) AssertNotEq10(opts *bind.CallOpts, left string, right string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq10", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq10 is a free data retrieval call binding the contract method 0x6a8237b3.
//
// Solidity: function assertNotEq(string left, string right) pure returns()
func (_Vm *VmSession) AssertNotEq10(left string, right string) error {
	return _Vm.Contract.AssertNotEq10(&_Vm.CallOpts, left, right)
}

// AssertNotEq10 is a free data retrieval call binding the contract method 0x6a8237b3.
//
// Solidity: function assertNotEq(string left, string right) pure returns()
func (_Vm *VmCallerSession) AssertNotEq10(left string, right string) error {
	return _Vm.Contract.AssertNotEq10(&_Vm.CallOpts, left, right)
}

// AssertNotEq11 is a free data retrieval call binding the contract method 0x72c7e0b5.
//
// Solidity: function assertNotEq(address[] left, address[] right, string error) pure returns()
func (_Vm *VmCaller) AssertNotEq11(opts *bind.CallOpts, left []common.Address, right []common.Address, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq11", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq11 is a free data retrieval call binding the contract method 0x72c7e0b5.
//
// Solidity: function assertNotEq(address[] left, address[] right, string error) pure returns()
func (_Vm *VmSession) AssertNotEq11(left []common.Address, right []common.Address, error string) error {
	return _Vm.Contract.AssertNotEq11(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq11 is a free data retrieval call binding the contract method 0x72c7e0b5.
//
// Solidity: function assertNotEq(address[] left, address[] right, string error) pure returns()
func (_Vm *VmCallerSession) AssertNotEq11(left []common.Address, right []common.Address, error string) error {
	return _Vm.Contract.AssertNotEq11(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq12 is a free data retrieval call binding the contract method 0x78bdcea7.
//
// Solidity: function assertNotEq(string left, string right, string error) pure returns()
func (_Vm *VmCaller) AssertNotEq12(opts *bind.CallOpts, left string, right string, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq12", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq12 is a free data retrieval call binding the contract method 0x78bdcea7.
//
// Solidity: function assertNotEq(string left, string right, string error) pure returns()
func (_Vm *VmSession) AssertNotEq12(left string, right string, error string) error {
	return _Vm.Contract.AssertNotEq12(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq12 is a free data retrieval call binding the contract method 0x78bdcea7.
//
// Solidity: function assertNotEq(string left, string right, string error) pure returns()
func (_Vm *VmCallerSession) AssertNotEq12(left string, right string, error string) error {
	return _Vm.Contract.AssertNotEq12(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq13 is a free data retrieval call binding the contract method 0x8775a591.
//
// Solidity: function assertNotEq(address left, address right, string error) pure returns()
func (_Vm *VmCaller) AssertNotEq13(opts *bind.CallOpts, left common.Address, right common.Address, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq13", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq13 is a free data retrieval call binding the contract method 0x8775a591.
//
// Solidity: function assertNotEq(address left, address right, string error) pure returns()
func (_Vm *VmSession) AssertNotEq13(left common.Address, right common.Address, error string) error {
	return _Vm.Contract.AssertNotEq13(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq13 is a free data retrieval call binding the contract method 0x8775a591.
//
// Solidity: function assertNotEq(address left, address right, string error) pure returns()
func (_Vm *VmCallerSession) AssertNotEq13(left common.Address, right common.Address, error string) error {
	return _Vm.Contract.AssertNotEq13(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq14 is a free data retrieval call binding the contract method 0x898e83fc.
//
// Solidity: function assertNotEq(bytes32 left, bytes32 right) pure returns()
func (_Vm *VmCaller) AssertNotEq14(opts *bind.CallOpts, left [32]byte, right [32]byte) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq14", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq14 is a free data retrieval call binding the contract method 0x898e83fc.
//
// Solidity: function assertNotEq(bytes32 left, bytes32 right) pure returns()
func (_Vm *VmSession) AssertNotEq14(left [32]byte, right [32]byte) error {
	return _Vm.Contract.AssertNotEq14(&_Vm.CallOpts, left, right)
}

// AssertNotEq14 is a free data retrieval call binding the contract method 0x898e83fc.
//
// Solidity: function assertNotEq(bytes32 left, bytes32 right) pure returns()
func (_Vm *VmCallerSession) AssertNotEq14(left [32]byte, right [32]byte) error {
	return _Vm.Contract.AssertNotEq14(&_Vm.CallOpts, left, right)
}

// AssertNotEq15 is a free data retrieval call binding the contract method 0x9507540e.
//
// Solidity: function assertNotEq(bytes left, bytes right, string error) pure returns()
func (_Vm *VmCaller) AssertNotEq15(opts *bind.CallOpts, left []byte, right []byte, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq15", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq15 is a free data retrieval call binding the contract method 0x9507540e.
//
// Solidity: function assertNotEq(bytes left, bytes right, string error) pure returns()
func (_Vm *VmSession) AssertNotEq15(left []byte, right []byte, error string) error {
	return _Vm.Contract.AssertNotEq15(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq15 is a free data retrieval call binding the contract method 0x9507540e.
//
// Solidity: function assertNotEq(bytes left, bytes right, string error) pure returns()
func (_Vm *VmCallerSession) AssertNotEq15(left []byte, right []byte, error string) error {
	return _Vm.Contract.AssertNotEq15(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq16 is a free data retrieval call binding the contract method 0x98f9bdbd.
//
// Solidity: function assertNotEq(uint256 left, uint256 right, string error) pure returns()
func (_Vm *VmCaller) AssertNotEq16(opts *bind.CallOpts, left *big.Int, right *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq16", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq16 is a free data retrieval call binding the contract method 0x98f9bdbd.
//
// Solidity: function assertNotEq(uint256 left, uint256 right, string error) pure returns()
func (_Vm *VmSession) AssertNotEq16(left *big.Int, right *big.Int, error string) error {
	return _Vm.Contract.AssertNotEq16(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq16 is a free data retrieval call binding the contract method 0x98f9bdbd.
//
// Solidity: function assertNotEq(uint256 left, uint256 right, string error) pure returns()
func (_Vm *VmCallerSession) AssertNotEq16(left *big.Int, right *big.Int, error string) error {
	return _Vm.Contract.AssertNotEq16(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq17 is a free data retrieval call binding the contract method 0x9a7fbd8f.
//
// Solidity: function assertNotEq(uint256[] left, uint256[] right, string error) pure returns()
func (_Vm *VmCaller) AssertNotEq17(opts *bind.CallOpts, left []*big.Int, right []*big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq17", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq17 is a free data retrieval call binding the contract method 0x9a7fbd8f.
//
// Solidity: function assertNotEq(uint256[] left, uint256[] right, string error) pure returns()
func (_Vm *VmSession) AssertNotEq17(left []*big.Int, right []*big.Int, error string) error {
	return _Vm.Contract.AssertNotEq17(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq17 is a free data retrieval call binding the contract method 0x9a7fbd8f.
//
// Solidity: function assertNotEq(uint256[] left, uint256[] right, string error) pure returns()
func (_Vm *VmCallerSession) AssertNotEq17(left []*big.Int, right []*big.Int, error string) error {
	return _Vm.Contract.AssertNotEq17(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq18 is a free data retrieval call binding the contract method 0xb12e1694.
//
// Solidity: function assertNotEq(address left, address right) pure returns()
func (_Vm *VmCaller) AssertNotEq18(opts *bind.CallOpts, left common.Address, right common.Address) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq18", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq18 is a free data retrieval call binding the contract method 0xb12e1694.
//
// Solidity: function assertNotEq(address left, address right) pure returns()
func (_Vm *VmSession) AssertNotEq18(left common.Address, right common.Address) error {
	return _Vm.Contract.AssertNotEq18(&_Vm.CallOpts, left, right)
}

// AssertNotEq18 is a free data retrieval call binding the contract method 0xb12e1694.
//
// Solidity: function assertNotEq(address left, address right) pure returns()
func (_Vm *VmCallerSession) AssertNotEq18(left common.Address, right common.Address) error {
	return _Vm.Contract.AssertNotEq18(&_Vm.CallOpts, left, right)
}

// AssertNotEq19 is a free data retrieval call binding the contract method 0xb2332f51.
//
// Solidity: function assertNotEq(bytes32 left, bytes32 right, string error) pure returns()
func (_Vm *VmCaller) AssertNotEq19(opts *bind.CallOpts, left [32]byte, right [32]byte, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq19", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq19 is a free data retrieval call binding the contract method 0xb2332f51.
//
// Solidity: function assertNotEq(bytes32 left, bytes32 right, string error) pure returns()
func (_Vm *VmSession) AssertNotEq19(left [32]byte, right [32]byte, error string) error {
	return _Vm.Contract.AssertNotEq19(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq19 is a free data retrieval call binding the contract method 0xb2332f51.
//
// Solidity: function assertNotEq(bytes32 left, bytes32 right, string error) pure returns()
func (_Vm *VmCallerSession) AssertNotEq19(left [32]byte, right [32]byte, error string) error {
	return _Vm.Contract.AssertNotEq19(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq2 is a free data retrieval call binding the contract method 0x1dcd1f68.
//
// Solidity: function assertNotEq(bytes[] left, bytes[] right, string error) pure returns()
func (_Vm *VmCaller) AssertNotEq2(opts *bind.CallOpts, left [][]byte, right [][]byte, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq2", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq2 is a free data retrieval call binding the contract method 0x1dcd1f68.
//
// Solidity: function assertNotEq(bytes[] left, bytes[] right, string error) pure returns()
func (_Vm *VmSession) AssertNotEq2(left [][]byte, right [][]byte, error string) error {
	return _Vm.Contract.AssertNotEq2(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq2 is a free data retrieval call binding the contract method 0x1dcd1f68.
//
// Solidity: function assertNotEq(bytes[] left, bytes[] right, string error) pure returns()
func (_Vm *VmCallerSession) AssertNotEq2(left [][]byte, right [][]byte, error string) error {
	return _Vm.Contract.AssertNotEq2(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq20 is a free data retrieval call binding the contract method 0xb67187f3.
//
// Solidity: function assertNotEq(string[] left, string[] right, string error) pure returns()
func (_Vm *VmCaller) AssertNotEq20(opts *bind.CallOpts, left []string, right []string, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq20", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq20 is a free data retrieval call binding the contract method 0xb67187f3.
//
// Solidity: function assertNotEq(string[] left, string[] right, string error) pure returns()
func (_Vm *VmSession) AssertNotEq20(left []string, right []string, error string) error {
	return _Vm.Contract.AssertNotEq20(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq20 is a free data retrieval call binding the contract method 0xb67187f3.
//
// Solidity: function assertNotEq(string[] left, string[] right, string error) pure returns()
func (_Vm *VmCallerSession) AssertNotEq20(left []string, right []string, error string) error {
	return _Vm.Contract.AssertNotEq20(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq21 is a free data retrieval call binding the contract method 0xb7909320.
//
// Solidity: function assertNotEq(uint256 left, uint256 right) pure returns()
func (_Vm *VmCaller) AssertNotEq21(opts *bind.CallOpts, left *big.Int, right *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq21", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq21 is a free data retrieval call binding the contract method 0xb7909320.
//
// Solidity: function assertNotEq(uint256 left, uint256 right) pure returns()
func (_Vm *VmSession) AssertNotEq21(left *big.Int, right *big.Int) error {
	return _Vm.Contract.AssertNotEq21(&_Vm.CallOpts, left, right)
}

// AssertNotEq21 is a free data retrieval call binding the contract method 0xb7909320.
//
// Solidity: function assertNotEq(uint256 left, uint256 right) pure returns()
func (_Vm *VmCallerSession) AssertNotEq21(left *big.Int, right *big.Int) error {
	return _Vm.Contract.AssertNotEq21(&_Vm.CallOpts, left, right)
}

// AssertNotEq22 is a free data retrieval call binding the contract method 0xb873634c.
//
// Solidity: function assertNotEq(bytes32[] left, bytes32[] right, string error) pure returns()
func (_Vm *VmCaller) AssertNotEq22(opts *bind.CallOpts, left [][32]byte, right [][32]byte, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq22", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq22 is a free data retrieval call binding the contract method 0xb873634c.
//
// Solidity: function assertNotEq(bytes32[] left, bytes32[] right, string error) pure returns()
func (_Vm *VmSession) AssertNotEq22(left [][32]byte, right [][32]byte, error string) error {
	return _Vm.Contract.AssertNotEq22(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq22 is a free data retrieval call binding the contract method 0xb873634c.
//
// Solidity: function assertNotEq(bytes32[] left, bytes32[] right, string error) pure returns()
func (_Vm *VmCallerSession) AssertNotEq22(left [][32]byte, right [][32]byte, error string) error {
	return _Vm.Contract.AssertNotEq22(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq23 is a free data retrieval call binding the contract method 0xbdfacbe8.
//
// Solidity: function assertNotEq(string[] left, string[] right) pure returns()
func (_Vm *VmCaller) AssertNotEq23(opts *bind.CallOpts, left []string, right []string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq23", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq23 is a free data retrieval call binding the contract method 0xbdfacbe8.
//
// Solidity: function assertNotEq(string[] left, string[] right) pure returns()
func (_Vm *VmSession) AssertNotEq23(left []string, right []string) error {
	return _Vm.Contract.AssertNotEq23(&_Vm.CallOpts, left, right)
}

// AssertNotEq23 is a free data retrieval call binding the contract method 0xbdfacbe8.
//
// Solidity: function assertNotEq(string[] left, string[] right) pure returns()
func (_Vm *VmCallerSession) AssertNotEq23(left []string, right []string) error {
	return _Vm.Contract.AssertNotEq23(&_Vm.CallOpts, left, right)
}

// AssertNotEq24 is a free data retrieval call binding the contract method 0xd3977322.
//
// Solidity: function assertNotEq(int256[] left, int256[] right, string error) pure returns()
func (_Vm *VmCaller) AssertNotEq24(opts *bind.CallOpts, left []*big.Int, right []*big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq24", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq24 is a free data retrieval call binding the contract method 0xd3977322.
//
// Solidity: function assertNotEq(int256[] left, int256[] right, string error) pure returns()
func (_Vm *VmSession) AssertNotEq24(left []*big.Int, right []*big.Int, error string) error {
	return _Vm.Contract.AssertNotEq24(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq24 is a free data retrieval call binding the contract method 0xd3977322.
//
// Solidity: function assertNotEq(int256[] left, int256[] right, string error) pure returns()
func (_Vm *VmCallerSession) AssertNotEq24(left []*big.Int, right []*big.Int, error string) error {
	return _Vm.Contract.AssertNotEq24(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq25 is a free data retrieval call binding the contract method 0xedecd035.
//
// Solidity: function assertNotEq(bytes[] left, bytes[] right) pure returns()
func (_Vm *VmCaller) AssertNotEq25(opts *bind.CallOpts, left [][]byte, right [][]byte) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq25", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq25 is a free data retrieval call binding the contract method 0xedecd035.
//
// Solidity: function assertNotEq(bytes[] left, bytes[] right) pure returns()
func (_Vm *VmSession) AssertNotEq25(left [][]byte, right [][]byte) error {
	return _Vm.Contract.AssertNotEq25(&_Vm.CallOpts, left, right)
}

// AssertNotEq25 is a free data retrieval call binding the contract method 0xedecd035.
//
// Solidity: function assertNotEq(bytes[] left, bytes[] right) pure returns()
func (_Vm *VmCallerSession) AssertNotEq25(left [][]byte, right [][]byte) error {
	return _Vm.Contract.AssertNotEq25(&_Vm.CallOpts, left, right)
}

// AssertNotEq26 is a free data retrieval call binding the contract method 0xf4c004e3.
//
// Solidity: function assertNotEq(int256 left, int256 right) pure returns()
func (_Vm *VmCaller) AssertNotEq26(opts *bind.CallOpts, left *big.Int, right *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq26", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq26 is a free data retrieval call binding the contract method 0xf4c004e3.
//
// Solidity: function assertNotEq(int256 left, int256 right) pure returns()
func (_Vm *VmSession) AssertNotEq26(left *big.Int, right *big.Int) error {
	return _Vm.Contract.AssertNotEq26(&_Vm.CallOpts, left, right)
}

// AssertNotEq26 is a free data retrieval call binding the contract method 0xf4c004e3.
//
// Solidity: function assertNotEq(int256 left, int256 right) pure returns()
func (_Vm *VmCallerSession) AssertNotEq26(left *big.Int, right *big.Int) error {
	return _Vm.Contract.AssertNotEq26(&_Vm.CallOpts, left, right)
}

// AssertNotEq3 is a free data retrieval call binding the contract method 0x236e4d66.
//
// Solidity: function assertNotEq(bool left, bool right) pure returns()
func (_Vm *VmCaller) AssertNotEq3(opts *bind.CallOpts, left bool, right bool) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq3", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq3 is a free data retrieval call binding the contract method 0x236e4d66.
//
// Solidity: function assertNotEq(bool left, bool right) pure returns()
func (_Vm *VmSession) AssertNotEq3(left bool, right bool) error {
	return _Vm.Contract.AssertNotEq3(&_Vm.CallOpts, left, right)
}

// AssertNotEq3 is a free data retrieval call binding the contract method 0x236e4d66.
//
// Solidity: function assertNotEq(bool left, bool right) pure returns()
func (_Vm *VmCallerSession) AssertNotEq3(left bool, right bool) error {
	return _Vm.Contract.AssertNotEq3(&_Vm.CallOpts, left, right)
}

// AssertNotEq4 is a free data retrieval call binding the contract method 0x286fafea.
//
// Solidity: function assertNotEq(bool[] left, bool[] right) pure returns()
func (_Vm *VmCaller) AssertNotEq4(opts *bind.CallOpts, left []bool, right []bool) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq4", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq4 is a free data retrieval call binding the contract method 0x286fafea.
//
// Solidity: function assertNotEq(bool[] left, bool[] right) pure returns()
func (_Vm *VmSession) AssertNotEq4(left []bool, right []bool) error {
	return _Vm.Contract.AssertNotEq4(&_Vm.CallOpts, left, right)
}

// AssertNotEq4 is a free data retrieval call binding the contract method 0x286fafea.
//
// Solidity: function assertNotEq(bool[] left, bool[] right) pure returns()
func (_Vm *VmCallerSession) AssertNotEq4(left []bool, right []bool) error {
	return _Vm.Contract.AssertNotEq4(&_Vm.CallOpts, left, right)
}

// AssertNotEq5 is a free data retrieval call binding the contract method 0x3cf78e28.
//
// Solidity: function assertNotEq(bytes left, bytes right) pure returns()
func (_Vm *VmCaller) AssertNotEq5(opts *bind.CallOpts, left []byte, right []byte) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq5", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq5 is a free data retrieval call binding the contract method 0x3cf78e28.
//
// Solidity: function assertNotEq(bytes left, bytes right) pure returns()
func (_Vm *VmSession) AssertNotEq5(left []byte, right []byte) error {
	return _Vm.Contract.AssertNotEq5(&_Vm.CallOpts, left, right)
}

// AssertNotEq5 is a free data retrieval call binding the contract method 0x3cf78e28.
//
// Solidity: function assertNotEq(bytes left, bytes right) pure returns()
func (_Vm *VmCallerSession) AssertNotEq5(left []byte, right []byte) error {
	return _Vm.Contract.AssertNotEq5(&_Vm.CallOpts, left, right)
}

// AssertNotEq6 is a free data retrieval call binding the contract method 0x46d0b252.
//
// Solidity: function assertNotEq(address[] left, address[] right) pure returns()
func (_Vm *VmCaller) AssertNotEq6(opts *bind.CallOpts, left []common.Address, right []common.Address) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq6", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq6 is a free data retrieval call binding the contract method 0x46d0b252.
//
// Solidity: function assertNotEq(address[] left, address[] right) pure returns()
func (_Vm *VmSession) AssertNotEq6(left []common.Address, right []common.Address) error {
	return _Vm.Contract.AssertNotEq6(&_Vm.CallOpts, left, right)
}

// AssertNotEq6 is a free data retrieval call binding the contract method 0x46d0b252.
//
// Solidity: function assertNotEq(address[] left, address[] right) pure returns()
func (_Vm *VmCallerSession) AssertNotEq6(left []common.Address, right []common.Address) error {
	return _Vm.Contract.AssertNotEq6(&_Vm.CallOpts, left, right)
}

// AssertNotEq7 is a free data retrieval call binding the contract method 0x4724c5b9.
//
// Solidity: function assertNotEq(int256 left, int256 right, string error) pure returns()
func (_Vm *VmCaller) AssertNotEq7(opts *bind.CallOpts, left *big.Int, right *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq7", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq7 is a free data retrieval call binding the contract method 0x4724c5b9.
//
// Solidity: function assertNotEq(int256 left, int256 right, string error) pure returns()
func (_Vm *VmSession) AssertNotEq7(left *big.Int, right *big.Int, error string) error {
	return _Vm.Contract.AssertNotEq7(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq7 is a free data retrieval call binding the contract method 0x4724c5b9.
//
// Solidity: function assertNotEq(int256 left, int256 right, string error) pure returns()
func (_Vm *VmCallerSession) AssertNotEq7(left *big.Int, right *big.Int, error string) error {
	return _Vm.Contract.AssertNotEq7(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq8 is a free data retrieval call binding the contract method 0x56f29cba.
//
// Solidity: function assertNotEq(uint256[] left, uint256[] right) pure returns()
func (_Vm *VmCaller) AssertNotEq8(opts *bind.CallOpts, left []*big.Int, right []*big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq8", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq8 is a free data retrieval call binding the contract method 0x56f29cba.
//
// Solidity: function assertNotEq(uint256[] left, uint256[] right) pure returns()
func (_Vm *VmSession) AssertNotEq8(left []*big.Int, right []*big.Int) error {
	return _Vm.Contract.AssertNotEq8(&_Vm.CallOpts, left, right)
}

// AssertNotEq8 is a free data retrieval call binding the contract method 0x56f29cba.
//
// Solidity: function assertNotEq(uint256[] left, uint256[] right) pure returns()
func (_Vm *VmCallerSession) AssertNotEq8(left []*big.Int, right []*big.Int) error {
	return _Vm.Contract.AssertNotEq8(&_Vm.CallOpts, left, right)
}

// AssertNotEq9 is a free data retrieval call binding the contract method 0x62c6f9fb.
//
// Solidity: function assertNotEq(bool[] left, bool[] right, string error) pure returns()
func (_Vm *VmCaller) AssertNotEq9(opts *bind.CallOpts, left []bool, right []bool, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq9", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq9 is a free data retrieval call binding the contract method 0x62c6f9fb.
//
// Solidity: function assertNotEq(bool[] left, bool[] right, string error) pure returns()
func (_Vm *VmSession) AssertNotEq9(left []bool, right []bool, error string) error {
	return _Vm.Contract.AssertNotEq9(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq9 is a free data retrieval call binding the contract method 0x62c6f9fb.
//
// Solidity: function assertNotEq(bool[] left, bool[] right, string error) pure returns()
func (_Vm *VmCallerSession) AssertNotEq9(left []bool, right []bool, error string) error {
	return _Vm.Contract.AssertNotEq9(&_Vm.CallOpts, left, right, error)
}

// AssertNotEqDecimal is a free data retrieval call binding the contract method 0x14e75680.
//
// Solidity: function assertNotEqDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_Vm *VmCaller) AssertNotEqDecimal(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEqDecimal", left, right, decimals)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEqDecimal is a free data retrieval call binding the contract method 0x14e75680.
//
// Solidity: function assertNotEqDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_Vm *VmSession) AssertNotEqDecimal(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertNotEqDecimal(&_Vm.CallOpts, left, right, decimals)
}

// AssertNotEqDecimal is a free data retrieval call binding the contract method 0x14e75680.
//
// Solidity: function assertNotEqDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_Vm *VmCallerSession) AssertNotEqDecimal(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertNotEqDecimal(&_Vm.CallOpts, left, right, decimals)
}

// AssertNotEqDecimal0 is a free data retrieval call binding the contract method 0x33949f0b.
//
// Solidity: function assertNotEqDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmCaller) AssertNotEqDecimal0(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEqDecimal0", left, right, decimals, error)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEqDecimal0 is a free data retrieval call binding the contract method 0x33949f0b.
//
// Solidity: function assertNotEqDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmSession) AssertNotEqDecimal0(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertNotEqDecimal0(&_Vm.CallOpts, left, right, decimals, error)
}

// AssertNotEqDecimal0 is a free data retrieval call binding the contract method 0x33949f0b.
//
// Solidity: function assertNotEqDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmCallerSession) AssertNotEqDecimal0(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertNotEqDecimal0(&_Vm.CallOpts, left, right, decimals, error)
}

// AssertNotEqDecimal1 is a free data retrieval call binding the contract method 0x669efca7.
//
// Solidity: function assertNotEqDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_Vm *VmCaller) AssertNotEqDecimal1(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEqDecimal1", left, right, decimals)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEqDecimal1 is a free data retrieval call binding the contract method 0x669efca7.
//
// Solidity: function assertNotEqDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_Vm *VmSession) AssertNotEqDecimal1(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertNotEqDecimal1(&_Vm.CallOpts, left, right, decimals)
}

// AssertNotEqDecimal1 is a free data retrieval call binding the contract method 0x669efca7.
//
// Solidity: function assertNotEqDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_Vm *VmCallerSession) AssertNotEqDecimal1(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertNotEqDecimal1(&_Vm.CallOpts, left, right, decimals)
}

// AssertNotEqDecimal2 is a free data retrieval call binding the contract method 0xf5a55558.
//
// Solidity: function assertNotEqDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmCaller) AssertNotEqDecimal2(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEqDecimal2", left, right, decimals, error)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEqDecimal2 is a free data retrieval call binding the contract method 0xf5a55558.
//
// Solidity: function assertNotEqDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmSession) AssertNotEqDecimal2(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertNotEqDecimal2(&_Vm.CallOpts, left, right, decimals, error)
}

// AssertNotEqDecimal2 is a free data retrieval call binding the contract method 0xf5a55558.
//
// Solidity: function assertNotEqDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmCallerSession) AssertNotEqDecimal2(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertNotEqDecimal2(&_Vm.CallOpts, left, right, decimals, error)
}

// AssertTrue is a free data retrieval call binding the contract method 0x0c9fd581.
//
// Solidity: function assertTrue(bool condition) pure returns()
func (_Vm *VmCaller) AssertTrue(opts *bind.CallOpts, condition bool) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertTrue", condition)

	if err != nil {
		return err
	}

	return err

}

// AssertTrue is a free data retrieval call binding the contract method 0x0c9fd581.
//
// Solidity: function assertTrue(bool condition) pure returns()
func (_Vm *VmSession) AssertTrue(condition bool) error {
	return _Vm.Contract.AssertTrue(&_Vm.CallOpts, condition)
}

// AssertTrue is a free data retrieval call binding the contract method 0x0c9fd581.
//
// Solidity: function assertTrue(bool condition) pure returns()
func (_Vm *VmCallerSession) AssertTrue(condition bool) error {
	return _Vm.Contract.AssertTrue(&_Vm.CallOpts, condition)
}

// AssertTrue0 is a free data retrieval call binding the contract method 0xa34edc03.
//
// Solidity: function assertTrue(bool condition, string error) pure returns()
func (_Vm *VmCaller) AssertTrue0(opts *bind.CallOpts, condition bool, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertTrue0", condition, error)

	if err != nil {
		return err
	}

	return err

}

// AssertTrue0 is a free data retrieval call binding the contract method 0xa34edc03.
//
// Solidity: function assertTrue(bool condition, string error) pure returns()
func (_Vm *VmSession) AssertTrue0(condition bool, error string) error {
	return _Vm.Contract.AssertTrue0(&_Vm.CallOpts, condition, error)
}

// AssertTrue0 is a free data retrieval call binding the contract method 0xa34edc03.
//
// Solidity: function assertTrue(bool condition, string error) pure returns()
func (_Vm *VmCallerSession) AssertTrue0(condition bool, error string) error {
	return _Vm.Contract.AssertTrue0(&_Vm.CallOpts, condition, error)
}

// Assume is a free data retrieval call binding the contract method 0x4c63e562.
//
// Solidity: function assume(bool condition) pure returns()
func (_Vm *VmCaller) Assume(opts *bind.CallOpts, condition bool) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assume", condition)

	if err != nil {
		return err
	}

	return err

}

// Assume is a free data retrieval call binding the contract method 0x4c63e562.
//
// Solidity: function assume(bool condition) pure returns()
func (_Vm *VmSession) Assume(condition bool) error {
	return _Vm.Contract.Assume(&_Vm.CallOpts, condition)
}

// Assume is a free data retrieval call binding the contract method 0x4c63e562.
//
// Solidity: function assume(bool condition) pure returns()
func (_Vm *VmCallerSession) Assume(condition bool) error {
	return _Vm.Contract.Assume(&_Vm.CallOpts, condition)
}

// ComputeCreate2Address is a free data retrieval call binding the contract method 0x890c283b.
//
// Solidity: function computeCreate2Address(bytes32 salt, bytes32 initCodeHash) pure returns(address)
func (_Vm *VmCaller) ComputeCreate2Address(opts *bind.CallOpts, salt [32]byte, initCodeHash [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "computeCreate2Address", salt, initCodeHash)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComputeCreate2Address is a free data retrieval call binding the contract method 0x890c283b.
//
// Solidity: function computeCreate2Address(bytes32 salt, bytes32 initCodeHash) pure returns(address)
func (_Vm *VmSession) ComputeCreate2Address(salt [32]byte, initCodeHash [32]byte) (common.Address, error) {
	return _Vm.Contract.ComputeCreate2Address(&_Vm.CallOpts, salt, initCodeHash)
}

// ComputeCreate2Address is a free data retrieval call binding the contract method 0x890c283b.
//
// Solidity: function computeCreate2Address(bytes32 salt, bytes32 initCodeHash) pure returns(address)
func (_Vm *VmCallerSession) ComputeCreate2Address(salt [32]byte, initCodeHash [32]byte) (common.Address, error) {
	return _Vm.Contract.ComputeCreate2Address(&_Vm.CallOpts, salt, initCodeHash)
}

// ComputeCreate2Address0 is a free data retrieval call binding the contract method 0xd323826a.
//
// Solidity: function computeCreate2Address(bytes32 salt, bytes32 initCodeHash, address deployer) pure returns(address)
func (_Vm *VmCaller) ComputeCreate2Address0(opts *bind.CallOpts, salt [32]byte, initCodeHash [32]byte, deployer common.Address) (common.Address, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "computeCreate2Address0", salt, initCodeHash, deployer)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComputeCreate2Address0 is a free data retrieval call binding the contract method 0xd323826a.
//
// Solidity: function computeCreate2Address(bytes32 salt, bytes32 initCodeHash, address deployer) pure returns(address)
func (_Vm *VmSession) ComputeCreate2Address0(salt [32]byte, initCodeHash [32]byte, deployer common.Address) (common.Address, error) {
	return _Vm.Contract.ComputeCreate2Address0(&_Vm.CallOpts, salt, initCodeHash, deployer)
}

// ComputeCreate2Address0 is a free data retrieval call binding the contract method 0xd323826a.
//
// Solidity: function computeCreate2Address(bytes32 salt, bytes32 initCodeHash, address deployer) pure returns(address)
func (_Vm *VmCallerSession) ComputeCreate2Address0(salt [32]byte, initCodeHash [32]byte, deployer common.Address) (common.Address, error) {
	return _Vm.Contract.ComputeCreate2Address0(&_Vm.CallOpts, salt, initCodeHash, deployer)
}

// ComputeCreateAddress is a free data retrieval call binding the contract method 0x74637a7a.
//
// Solidity: function computeCreateAddress(address deployer, uint256 nonce) pure returns(address)
func (_Vm *VmCaller) ComputeCreateAddress(opts *bind.CallOpts, deployer common.Address, nonce *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "computeCreateAddress", deployer, nonce)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComputeCreateAddress is a free data retrieval call binding the contract method 0x74637a7a.
//
// Solidity: function computeCreateAddress(address deployer, uint256 nonce) pure returns(address)
func (_Vm *VmSession) ComputeCreateAddress(deployer common.Address, nonce *big.Int) (common.Address, error) {
	return _Vm.Contract.ComputeCreateAddress(&_Vm.CallOpts, deployer, nonce)
}

// ComputeCreateAddress is a free data retrieval call binding the contract method 0x74637a7a.
//
// Solidity: function computeCreateAddress(address deployer, uint256 nonce) pure returns(address)
func (_Vm *VmCallerSession) ComputeCreateAddress(deployer common.Address, nonce *big.Int) (common.Address, error) {
	return _Vm.Contract.ComputeCreateAddress(&_Vm.CallOpts, deployer, nonce)
}

// DeriveKey is a free data retrieval call binding the contract method 0x29233b1f.
//
// Solidity: function deriveKey(string mnemonic, string derivationPath, uint32 index, string language) pure returns(uint256 privateKey)
func (_Vm *VmCaller) DeriveKey(opts *bind.CallOpts, mnemonic string, derivationPath string, index uint32, language string) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "deriveKey", mnemonic, derivationPath, index, language)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeriveKey is a free data retrieval call binding the contract method 0x29233b1f.
//
// Solidity: function deriveKey(string mnemonic, string derivationPath, uint32 index, string language) pure returns(uint256 privateKey)
func (_Vm *VmSession) DeriveKey(mnemonic string, derivationPath string, index uint32, language string) (*big.Int, error) {
	return _Vm.Contract.DeriveKey(&_Vm.CallOpts, mnemonic, derivationPath, index, language)
}

// DeriveKey is a free data retrieval call binding the contract method 0x29233b1f.
//
// Solidity: function deriveKey(string mnemonic, string derivationPath, uint32 index, string language) pure returns(uint256 privateKey)
func (_Vm *VmCallerSession) DeriveKey(mnemonic string, derivationPath string, index uint32, language string) (*big.Int, error) {
	return _Vm.Contract.DeriveKey(&_Vm.CallOpts, mnemonic, derivationPath, index, language)
}

// DeriveKey0 is a free data retrieval call binding the contract method 0x32c8176d.
//
// Solidity: function deriveKey(string mnemonic, uint32 index, string language) pure returns(uint256 privateKey)
func (_Vm *VmCaller) DeriveKey0(opts *bind.CallOpts, mnemonic string, index uint32, language string) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "deriveKey0", mnemonic, index, language)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeriveKey0 is a free data retrieval call binding the contract method 0x32c8176d.
//
// Solidity: function deriveKey(string mnemonic, uint32 index, string language) pure returns(uint256 privateKey)
func (_Vm *VmSession) DeriveKey0(mnemonic string, index uint32, language string) (*big.Int, error) {
	return _Vm.Contract.DeriveKey0(&_Vm.CallOpts, mnemonic, index, language)
}

// DeriveKey0 is a free data retrieval call binding the contract method 0x32c8176d.
//
// Solidity: function deriveKey(string mnemonic, uint32 index, string language) pure returns(uint256 privateKey)
func (_Vm *VmCallerSession) DeriveKey0(mnemonic string, index uint32, language string) (*big.Int, error) {
	return _Vm.Contract.DeriveKey0(&_Vm.CallOpts, mnemonic, index, language)
}

// DeriveKey1 is a free data retrieval call binding the contract method 0x6229498b.
//
// Solidity: function deriveKey(string mnemonic, uint32 index) pure returns(uint256 privateKey)
func (_Vm *VmCaller) DeriveKey1(opts *bind.CallOpts, mnemonic string, index uint32) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "deriveKey1", mnemonic, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeriveKey1 is a free data retrieval call binding the contract method 0x6229498b.
//
// Solidity: function deriveKey(string mnemonic, uint32 index) pure returns(uint256 privateKey)
func (_Vm *VmSession) DeriveKey1(mnemonic string, index uint32) (*big.Int, error) {
	return _Vm.Contract.DeriveKey1(&_Vm.CallOpts, mnemonic, index)
}

// DeriveKey1 is a free data retrieval call binding the contract method 0x6229498b.
//
// Solidity: function deriveKey(string mnemonic, uint32 index) pure returns(uint256 privateKey)
func (_Vm *VmCallerSession) DeriveKey1(mnemonic string, index uint32) (*big.Int, error) {
	return _Vm.Contract.DeriveKey1(&_Vm.CallOpts, mnemonic, index)
}

// DeriveKey2 is a free data retrieval call binding the contract method 0x6bcb2c1b.
//
// Solidity: function deriveKey(string mnemonic, string derivationPath, uint32 index) pure returns(uint256 privateKey)
func (_Vm *VmCaller) DeriveKey2(opts *bind.CallOpts, mnemonic string, derivationPath string, index uint32) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "deriveKey2", mnemonic, derivationPath, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeriveKey2 is a free data retrieval call binding the contract method 0x6bcb2c1b.
//
// Solidity: function deriveKey(string mnemonic, string derivationPath, uint32 index) pure returns(uint256 privateKey)
func (_Vm *VmSession) DeriveKey2(mnemonic string, derivationPath string, index uint32) (*big.Int, error) {
	return _Vm.Contract.DeriveKey2(&_Vm.CallOpts, mnemonic, derivationPath, index)
}

// DeriveKey2 is a free data retrieval call binding the contract method 0x6bcb2c1b.
//
// Solidity: function deriveKey(string mnemonic, string derivationPath, uint32 index) pure returns(uint256 privateKey)
func (_Vm *VmCallerSession) DeriveKey2(mnemonic string, derivationPath string, index uint32) (*big.Int, error) {
	return _Vm.Contract.DeriveKey2(&_Vm.CallOpts, mnemonic, derivationPath, index)
}

// EnsNamehash is a free data retrieval call binding the contract method 0x8c374c65.
//
// Solidity: function ensNamehash(string name) pure returns(bytes32)
func (_Vm *VmCaller) EnsNamehash(opts *bind.CallOpts, name string) ([32]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "ensNamehash", name)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EnsNamehash is a free data retrieval call binding the contract method 0x8c374c65.
//
// Solidity: function ensNamehash(string name) pure returns(bytes32)
func (_Vm *VmSession) EnsNamehash(name string) ([32]byte, error) {
	return _Vm.Contract.EnsNamehash(&_Vm.CallOpts, name)
}

// EnsNamehash is a free data retrieval call binding the contract method 0x8c374c65.
//
// Solidity: function ensNamehash(string name) pure returns(bytes32)
func (_Vm *VmCallerSession) EnsNamehash(name string) ([32]byte, error) {
	return _Vm.Contract.EnsNamehash(&_Vm.CallOpts, name)
}

// EnvAddress is a free data retrieval call binding the contract method 0x350d56bf.
//
// Solidity: function envAddress(string name) view returns(address value)
func (_Vm *VmCaller) EnvAddress(opts *bind.CallOpts, name string) (common.Address, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envAddress", name)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EnvAddress is a free data retrieval call binding the contract method 0x350d56bf.
//
// Solidity: function envAddress(string name) view returns(address value)
func (_Vm *VmSession) EnvAddress(name string) (common.Address, error) {
	return _Vm.Contract.EnvAddress(&_Vm.CallOpts, name)
}

// EnvAddress is a free data retrieval call binding the contract method 0x350d56bf.
//
// Solidity: function envAddress(string name) view returns(address value)
func (_Vm *VmCallerSession) EnvAddress(name string) (common.Address, error) {
	return _Vm.Contract.EnvAddress(&_Vm.CallOpts, name)
}

// EnvAddress0 is a free data retrieval call binding the contract method 0xad31b9fa.
//
// Solidity: function envAddress(string name, string delim) view returns(address[] value)
func (_Vm *VmCaller) EnvAddress0(opts *bind.CallOpts, name string, delim string) ([]common.Address, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envAddress0", name, delim)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// EnvAddress0 is a free data retrieval call binding the contract method 0xad31b9fa.
//
// Solidity: function envAddress(string name, string delim) view returns(address[] value)
func (_Vm *VmSession) EnvAddress0(name string, delim string) ([]common.Address, error) {
	return _Vm.Contract.EnvAddress0(&_Vm.CallOpts, name, delim)
}

// EnvAddress0 is a free data retrieval call binding the contract method 0xad31b9fa.
//
// Solidity: function envAddress(string name, string delim) view returns(address[] value)
func (_Vm *VmCallerSession) EnvAddress0(name string, delim string) ([]common.Address, error) {
	return _Vm.Contract.EnvAddress0(&_Vm.CallOpts, name, delim)
}

// EnvBool is a free data retrieval call binding the contract method 0x7ed1ec7d.
//
// Solidity: function envBool(string name) view returns(bool value)
func (_Vm *VmCaller) EnvBool(opts *bind.CallOpts, name string) (bool, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envBool", name)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EnvBool is a free data retrieval call binding the contract method 0x7ed1ec7d.
//
// Solidity: function envBool(string name) view returns(bool value)
func (_Vm *VmSession) EnvBool(name string) (bool, error) {
	return _Vm.Contract.EnvBool(&_Vm.CallOpts, name)
}

// EnvBool is a free data retrieval call binding the contract method 0x7ed1ec7d.
//
// Solidity: function envBool(string name) view returns(bool value)
func (_Vm *VmCallerSession) EnvBool(name string) (bool, error) {
	return _Vm.Contract.EnvBool(&_Vm.CallOpts, name)
}

// EnvBool0 is a free data retrieval call binding the contract method 0xaaaddeaf.
//
// Solidity: function envBool(string name, string delim) view returns(bool[] value)
func (_Vm *VmCaller) EnvBool0(opts *bind.CallOpts, name string, delim string) ([]bool, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envBool0", name, delim)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// EnvBool0 is a free data retrieval call binding the contract method 0xaaaddeaf.
//
// Solidity: function envBool(string name, string delim) view returns(bool[] value)
func (_Vm *VmSession) EnvBool0(name string, delim string) ([]bool, error) {
	return _Vm.Contract.EnvBool0(&_Vm.CallOpts, name, delim)
}

// EnvBool0 is a free data retrieval call binding the contract method 0xaaaddeaf.
//
// Solidity: function envBool(string name, string delim) view returns(bool[] value)
func (_Vm *VmCallerSession) EnvBool0(name string, delim string) ([]bool, error) {
	return _Vm.Contract.EnvBool0(&_Vm.CallOpts, name, delim)
}

// EnvBytes is a free data retrieval call binding the contract method 0x4d7baf06.
//
// Solidity: function envBytes(string name) view returns(bytes value)
func (_Vm *VmCaller) EnvBytes(opts *bind.CallOpts, name string) ([]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envBytes", name)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EnvBytes is a free data retrieval call binding the contract method 0x4d7baf06.
//
// Solidity: function envBytes(string name) view returns(bytes value)
func (_Vm *VmSession) EnvBytes(name string) ([]byte, error) {
	return _Vm.Contract.EnvBytes(&_Vm.CallOpts, name)
}

// EnvBytes is a free data retrieval call binding the contract method 0x4d7baf06.
//
// Solidity: function envBytes(string name) view returns(bytes value)
func (_Vm *VmCallerSession) EnvBytes(name string) ([]byte, error) {
	return _Vm.Contract.EnvBytes(&_Vm.CallOpts, name)
}

// EnvBytes0 is a free data retrieval call binding the contract method 0xddc2651b.
//
// Solidity: function envBytes(string name, string delim) view returns(bytes[] value)
func (_Vm *VmCaller) EnvBytes0(opts *bind.CallOpts, name string, delim string) ([][]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envBytes0", name, delim)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// EnvBytes0 is a free data retrieval call binding the contract method 0xddc2651b.
//
// Solidity: function envBytes(string name, string delim) view returns(bytes[] value)
func (_Vm *VmSession) EnvBytes0(name string, delim string) ([][]byte, error) {
	return _Vm.Contract.EnvBytes0(&_Vm.CallOpts, name, delim)
}

// EnvBytes0 is a free data retrieval call binding the contract method 0xddc2651b.
//
// Solidity: function envBytes(string name, string delim) view returns(bytes[] value)
func (_Vm *VmCallerSession) EnvBytes0(name string, delim string) ([][]byte, error) {
	return _Vm.Contract.EnvBytes0(&_Vm.CallOpts, name, delim)
}

// EnvBytes32 is a free data retrieval call binding the contract method 0x5af231c1.
//
// Solidity: function envBytes32(string name, string delim) view returns(bytes32[] value)
func (_Vm *VmCaller) EnvBytes32(opts *bind.CallOpts, name string, delim string) ([][32]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envBytes32", name, delim)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// EnvBytes32 is a free data retrieval call binding the contract method 0x5af231c1.
//
// Solidity: function envBytes32(string name, string delim) view returns(bytes32[] value)
func (_Vm *VmSession) EnvBytes32(name string, delim string) ([][32]byte, error) {
	return _Vm.Contract.EnvBytes32(&_Vm.CallOpts, name, delim)
}

// EnvBytes32 is a free data retrieval call binding the contract method 0x5af231c1.
//
// Solidity: function envBytes32(string name, string delim) view returns(bytes32[] value)
func (_Vm *VmCallerSession) EnvBytes32(name string, delim string) ([][32]byte, error) {
	return _Vm.Contract.EnvBytes32(&_Vm.CallOpts, name, delim)
}

// EnvBytes320 is a free data retrieval call binding the contract method 0x97949042.
//
// Solidity: function envBytes32(string name) view returns(bytes32 value)
func (_Vm *VmCaller) EnvBytes320(opts *bind.CallOpts, name string) ([32]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envBytes320", name)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EnvBytes320 is a free data retrieval call binding the contract method 0x97949042.
//
// Solidity: function envBytes32(string name) view returns(bytes32 value)
func (_Vm *VmSession) EnvBytes320(name string) ([32]byte, error) {
	return _Vm.Contract.EnvBytes320(&_Vm.CallOpts, name)
}

// EnvBytes320 is a free data retrieval call binding the contract method 0x97949042.
//
// Solidity: function envBytes32(string name) view returns(bytes32 value)
func (_Vm *VmCallerSession) EnvBytes320(name string) ([32]byte, error) {
	return _Vm.Contract.EnvBytes320(&_Vm.CallOpts, name)
}

// EnvExists is a free data retrieval call binding the contract method 0xce8365f9.
//
// Solidity: function envExists(string name) view returns(bool result)
func (_Vm *VmCaller) EnvExists(opts *bind.CallOpts, name string) (bool, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envExists", name)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EnvExists is a free data retrieval call binding the contract method 0xce8365f9.
//
// Solidity: function envExists(string name) view returns(bool result)
func (_Vm *VmSession) EnvExists(name string) (bool, error) {
	return _Vm.Contract.EnvExists(&_Vm.CallOpts, name)
}

// EnvExists is a free data retrieval call binding the contract method 0xce8365f9.
//
// Solidity: function envExists(string name) view returns(bool result)
func (_Vm *VmCallerSession) EnvExists(name string) (bool, error) {
	return _Vm.Contract.EnvExists(&_Vm.CallOpts, name)
}

// EnvInt is a free data retrieval call binding the contract method 0x42181150.
//
// Solidity: function envInt(string name, string delim) view returns(int256[] value)
func (_Vm *VmCaller) EnvInt(opts *bind.CallOpts, name string, delim string) ([]*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envInt", name, delim)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// EnvInt is a free data retrieval call binding the contract method 0x42181150.
//
// Solidity: function envInt(string name, string delim) view returns(int256[] value)
func (_Vm *VmSession) EnvInt(name string, delim string) ([]*big.Int, error) {
	return _Vm.Contract.EnvInt(&_Vm.CallOpts, name, delim)
}

// EnvInt is a free data retrieval call binding the contract method 0x42181150.
//
// Solidity: function envInt(string name, string delim) view returns(int256[] value)
func (_Vm *VmCallerSession) EnvInt(name string, delim string) ([]*big.Int, error) {
	return _Vm.Contract.EnvInt(&_Vm.CallOpts, name, delim)
}

// EnvInt0 is a free data retrieval call binding the contract method 0x892a0c61.
//
// Solidity: function envInt(string name) view returns(int256 value)
func (_Vm *VmCaller) EnvInt0(opts *bind.CallOpts, name string) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envInt0", name)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EnvInt0 is a free data retrieval call binding the contract method 0x892a0c61.
//
// Solidity: function envInt(string name) view returns(int256 value)
func (_Vm *VmSession) EnvInt0(name string) (*big.Int, error) {
	return _Vm.Contract.EnvInt0(&_Vm.CallOpts, name)
}

// EnvInt0 is a free data retrieval call binding the contract method 0x892a0c61.
//
// Solidity: function envInt(string name) view returns(int256 value)
func (_Vm *VmCallerSession) EnvInt0(name string) (*big.Int, error) {
	return _Vm.Contract.EnvInt0(&_Vm.CallOpts, name)
}

// EnvOr is a free data retrieval call binding the contract method 0x2281f367.
//
// Solidity: function envOr(string name, string delim, bytes32[] defaultValue) view returns(bytes32[] value)
func (_Vm *VmCaller) EnvOr(opts *bind.CallOpts, name string, delim string, defaultValue [][32]byte) ([][32]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envOr", name, delim, defaultValue)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// EnvOr is a free data retrieval call binding the contract method 0x2281f367.
//
// Solidity: function envOr(string name, string delim, bytes32[] defaultValue) view returns(bytes32[] value)
func (_Vm *VmSession) EnvOr(name string, delim string, defaultValue [][32]byte) ([][32]byte, error) {
	return _Vm.Contract.EnvOr(&_Vm.CallOpts, name, delim, defaultValue)
}

// EnvOr is a free data retrieval call binding the contract method 0x2281f367.
//
// Solidity: function envOr(string name, string delim, bytes32[] defaultValue) view returns(bytes32[] value)
func (_Vm *VmCallerSession) EnvOr(name string, delim string, defaultValue [][32]byte) ([][32]byte, error) {
	return _Vm.Contract.EnvOr(&_Vm.CallOpts, name, delim, defaultValue)
}

// EnvOr0 is a free data retrieval call binding the contract method 0x4700d74b.
//
// Solidity: function envOr(string name, string delim, int256[] defaultValue) view returns(int256[] value)
func (_Vm *VmCaller) EnvOr0(opts *bind.CallOpts, name string, delim string, defaultValue []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envOr0", name, delim, defaultValue)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// EnvOr0 is a free data retrieval call binding the contract method 0x4700d74b.
//
// Solidity: function envOr(string name, string delim, int256[] defaultValue) view returns(int256[] value)
func (_Vm *VmSession) EnvOr0(name string, delim string, defaultValue []*big.Int) ([]*big.Int, error) {
	return _Vm.Contract.EnvOr0(&_Vm.CallOpts, name, delim, defaultValue)
}

// EnvOr0 is a free data retrieval call binding the contract method 0x4700d74b.
//
// Solidity: function envOr(string name, string delim, int256[] defaultValue) view returns(int256[] value)
func (_Vm *VmCallerSession) EnvOr0(name string, delim string, defaultValue []*big.Int) ([]*big.Int, error) {
	return _Vm.Contract.EnvOr0(&_Vm.CallOpts, name, delim, defaultValue)
}

// EnvOr1 is a free data retrieval call binding the contract method 0x4777f3cf.
//
// Solidity: function envOr(string name, bool defaultValue) view returns(bool value)
func (_Vm *VmCaller) EnvOr1(opts *bind.CallOpts, name string, defaultValue bool) (bool, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envOr1", name, defaultValue)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EnvOr1 is a free data retrieval call binding the contract method 0x4777f3cf.
//
// Solidity: function envOr(string name, bool defaultValue) view returns(bool value)
func (_Vm *VmSession) EnvOr1(name string, defaultValue bool) (bool, error) {
	return _Vm.Contract.EnvOr1(&_Vm.CallOpts, name, defaultValue)
}

// EnvOr1 is a free data retrieval call binding the contract method 0x4777f3cf.
//
// Solidity: function envOr(string name, bool defaultValue) view returns(bool value)
func (_Vm *VmCallerSession) EnvOr1(name string, defaultValue bool) (bool, error) {
	return _Vm.Contract.EnvOr1(&_Vm.CallOpts, name, defaultValue)
}

// EnvOr10 is a free data retrieval call binding the contract method 0xc74e9deb.
//
// Solidity: function envOr(string name, string delim, address[] defaultValue) view returns(address[] value)
func (_Vm *VmCaller) EnvOr10(opts *bind.CallOpts, name string, delim string, defaultValue []common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envOr10", name, delim, defaultValue)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// EnvOr10 is a free data retrieval call binding the contract method 0xc74e9deb.
//
// Solidity: function envOr(string name, string delim, address[] defaultValue) view returns(address[] value)
func (_Vm *VmSession) EnvOr10(name string, delim string, defaultValue []common.Address) ([]common.Address, error) {
	return _Vm.Contract.EnvOr10(&_Vm.CallOpts, name, delim, defaultValue)
}

// EnvOr10 is a free data retrieval call binding the contract method 0xc74e9deb.
//
// Solidity: function envOr(string name, string delim, address[] defaultValue) view returns(address[] value)
func (_Vm *VmCallerSession) EnvOr10(name string, delim string, defaultValue []common.Address) ([]common.Address, error) {
	return _Vm.Contract.EnvOr10(&_Vm.CallOpts, name, delim, defaultValue)
}

// EnvOr11 is a free data retrieval call binding the contract method 0xd145736c.
//
// Solidity: function envOr(string name, string defaultValue) view returns(string value)
func (_Vm *VmCaller) EnvOr11(opts *bind.CallOpts, name string, defaultValue string) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envOr11", name, defaultValue)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// EnvOr11 is a free data retrieval call binding the contract method 0xd145736c.
//
// Solidity: function envOr(string name, string defaultValue) view returns(string value)
func (_Vm *VmSession) EnvOr11(name string, defaultValue string) (string, error) {
	return _Vm.Contract.EnvOr11(&_Vm.CallOpts, name, defaultValue)
}

// EnvOr11 is a free data retrieval call binding the contract method 0xd145736c.
//
// Solidity: function envOr(string name, string defaultValue) view returns(string value)
func (_Vm *VmCallerSession) EnvOr11(name string, defaultValue string) (string, error) {
	return _Vm.Contract.EnvOr11(&_Vm.CallOpts, name, defaultValue)
}

// EnvOr12 is a free data retrieval call binding the contract method 0xeb85e83b.
//
// Solidity: function envOr(string name, string delim, bool[] defaultValue) view returns(bool[] value)
func (_Vm *VmCaller) EnvOr12(opts *bind.CallOpts, name string, delim string, defaultValue []bool) ([]bool, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envOr12", name, delim, defaultValue)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// EnvOr12 is a free data retrieval call binding the contract method 0xeb85e83b.
//
// Solidity: function envOr(string name, string delim, bool[] defaultValue) view returns(bool[] value)
func (_Vm *VmSession) EnvOr12(name string, delim string, defaultValue []bool) ([]bool, error) {
	return _Vm.Contract.EnvOr12(&_Vm.CallOpts, name, delim, defaultValue)
}

// EnvOr12 is a free data retrieval call binding the contract method 0xeb85e83b.
//
// Solidity: function envOr(string name, string delim, bool[] defaultValue) view returns(bool[] value)
func (_Vm *VmCallerSession) EnvOr12(name string, delim string, defaultValue []bool) ([]bool, error) {
	return _Vm.Contract.EnvOr12(&_Vm.CallOpts, name, delim, defaultValue)
}

// EnvOr2 is a free data retrieval call binding the contract method 0x561fe540.
//
// Solidity: function envOr(string name, address defaultValue) view returns(address value)
func (_Vm *VmCaller) EnvOr2(opts *bind.CallOpts, name string, defaultValue common.Address) (common.Address, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envOr2", name, defaultValue)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EnvOr2 is a free data retrieval call binding the contract method 0x561fe540.
//
// Solidity: function envOr(string name, address defaultValue) view returns(address value)
func (_Vm *VmSession) EnvOr2(name string, defaultValue common.Address) (common.Address, error) {
	return _Vm.Contract.EnvOr2(&_Vm.CallOpts, name, defaultValue)
}

// EnvOr2 is a free data retrieval call binding the contract method 0x561fe540.
//
// Solidity: function envOr(string name, address defaultValue) view returns(address value)
func (_Vm *VmCallerSession) EnvOr2(name string, defaultValue common.Address) (common.Address, error) {
	return _Vm.Contract.EnvOr2(&_Vm.CallOpts, name, defaultValue)
}

// EnvOr3 is a free data retrieval call binding the contract method 0x5e97348f.
//
// Solidity: function envOr(string name, uint256 defaultValue) view returns(uint256 value)
func (_Vm *VmCaller) EnvOr3(opts *bind.CallOpts, name string, defaultValue *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envOr3", name, defaultValue)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EnvOr3 is a free data retrieval call binding the contract method 0x5e97348f.
//
// Solidity: function envOr(string name, uint256 defaultValue) view returns(uint256 value)
func (_Vm *VmSession) EnvOr3(name string, defaultValue *big.Int) (*big.Int, error) {
	return _Vm.Contract.EnvOr3(&_Vm.CallOpts, name, defaultValue)
}

// EnvOr3 is a free data retrieval call binding the contract method 0x5e97348f.
//
// Solidity: function envOr(string name, uint256 defaultValue) view returns(uint256 value)
func (_Vm *VmCallerSession) EnvOr3(name string, defaultValue *big.Int) (*big.Int, error) {
	return _Vm.Contract.EnvOr3(&_Vm.CallOpts, name, defaultValue)
}

// EnvOr4 is a free data retrieval call binding the contract method 0x64bc3e64.
//
// Solidity: function envOr(string name, string delim, bytes[] defaultValue) view returns(bytes[] value)
func (_Vm *VmCaller) EnvOr4(opts *bind.CallOpts, name string, delim string, defaultValue [][]byte) ([][]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envOr4", name, delim, defaultValue)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// EnvOr4 is a free data retrieval call binding the contract method 0x64bc3e64.
//
// Solidity: function envOr(string name, string delim, bytes[] defaultValue) view returns(bytes[] value)
func (_Vm *VmSession) EnvOr4(name string, delim string, defaultValue [][]byte) ([][]byte, error) {
	return _Vm.Contract.EnvOr4(&_Vm.CallOpts, name, delim, defaultValue)
}

// EnvOr4 is a free data retrieval call binding the contract method 0x64bc3e64.
//
// Solidity: function envOr(string name, string delim, bytes[] defaultValue) view returns(bytes[] value)
func (_Vm *VmCallerSession) EnvOr4(name string, delim string, defaultValue [][]byte) ([][]byte, error) {
	return _Vm.Contract.EnvOr4(&_Vm.CallOpts, name, delim, defaultValue)
}

// EnvOr5 is a free data retrieval call binding the contract method 0x74318528.
//
// Solidity: function envOr(string name, string delim, uint256[] defaultValue) view returns(uint256[] value)
func (_Vm *VmCaller) EnvOr5(opts *bind.CallOpts, name string, delim string, defaultValue []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envOr5", name, delim, defaultValue)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// EnvOr5 is a free data retrieval call binding the contract method 0x74318528.
//
// Solidity: function envOr(string name, string delim, uint256[] defaultValue) view returns(uint256[] value)
func (_Vm *VmSession) EnvOr5(name string, delim string, defaultValue []*big.Int) ([]*big.Int, error) {
	return _Vm.Contract.EnvOr5(&_Vm.CallOpts, name, delim, defaultValue)
}

// EnvOr5 is a free data retrieval call binding the contract method 0x74318528.
//
// Solidity: function envOr(string name, string delim, uint256[] defaultValue) view returns(uint256[] value)
func (_Vm *VmCallerSession) EnvOr5(name string, delim string, defaultValue []*big.Int) ([]*big.Int, error) {
	return _Vm.Contract.EnvOr5(&_Vm.CallOpts, name, delim, defaultValue)
}

// EnvOr6 is a free data retrieval call binding the contract method 0x859216bc.
//
// Solidity: function envOr(string name, string delim, string[] defaultValue) view returns(string[] value)
func (_Vm *VmCaller) EnvOr6(opts *bind.CallOpts, name string, delim string, defaultValue []string) ([]string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envOr6", name, delim, defaultValue)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// EnvOr6 is a free data retrieval call binding the contract method 0x859216bc.
//
// Solidity: function envOr(string name, string delim, string[] defaultValue) view returns(string[] value)
func (_Vm *VmSession) EnvOr6(name string, delim string, defaultValue []string) ([]string, error) {
	return _Vm.Contract.EnvOr6(&_Vm.CallOpts, name, delim, defaultValue)
}

// EnvOr6 is a free data retrieval call binding the contract method 0x859216bc.
//
// Solidity: function envOr(string name, string delim, string[] defaultValue) view returns(string[] value)
func (_Vm *VmCallerSession) EnvOr6(name string, delim string, defaultValue []string) ([]string, error) {
	return _Vm.Contract.EnvOr6(&_Vm.CallOpts, name, delim, defaultValue)
}

// EnvOr7 is a free data retrieval call binding the contract method 0xb3e47705.
//
// Solidity: function envOr(string name, bytes defaultValue) view returns(bytes value)
func (_Vm *VmCaller) EnvOr7(opts *bind.CallOpts, name string, defaultValue []byte) ([]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envOr7", name, defaultValue)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EnvOr7 is a free data retrieval call binding the contract method 0xb3e47705.
//
// Solidity: function envOr(string name, bytes defaultValue) view returns(bytes value)
func (_Vm *VmSession) EnvOr7(name string, defaultValue []byte) ([]byte, error) {
	return _Vm.Contract.EnvOr7(&_Vm.CallOpts, name, defaultValue)
}

// EnvOr7 is a free data retrieval call binding the contract method 0xb3e47705.
//
// Solidity: function envOr(string name, bytes defaultValue) view returns(bytes value)
func (_Vm *VmCallerSession) EnvOr7(name string, defaultValue []byte) ([]byte, error) {
	return _Vm.Contract.EnvOr7(&_Vm.CallOpts, name, defaultValue)
}

// EnvOr8 is a free data retrieval call binding the contract method 0xb4a85892.
//
// Solidity: function envOr(string name, bytes32 defaultValue) view returns(bytes32 value)
func (_Vm *VmCaller) EnvOr8(opts *bind.CallOpts, name string, defaultValue [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envOr8", name, defaultValue)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EnvOr8 is a free data retrieval call binding the contract method 0xb4a85892.
//
// Solidity: function envOr(string name, bytes32 defaultValue) view returns(bytes32 value)
func (_Vm *VmSession) EnvOr8(name string, defaultValue [32]byte) ([32]byte, error) {
	return _Vm.Contract.EnvOr8(&_Vm.CallOpts, name, defaultValue)
}

// EnvOr8 is a free data retrieval call binding the contract method 0xb4a85892.
//
// Solidity: function envOr(string name, bytes32 defaultValue) view returns(bytes32 value)
func (_Vm *VmCallerSession) EnvOr8(name string, defaultValue [32]byte) ([32]byte, error) {
	return _Vm.Contract.EnvOr8(&_Vm.CallOpts, name, defaultValue)
}

// EnvOr9 is a free data retrieval call binding the contract method 0xbbcb713e.
//
// Solidity: function envOr(string name, int256 defaultValue) view returns(int256 value)
func (_Vm *VmCaller) EnvOr9(opts *bind.CallOpts, name string, defaultValue *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envOr9", name, defaultValue)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EnvOr9 is a free data retrieval call binding the contract method 0xbbcb713e.
//
// Solidity: function envOr(string name, int256 defaultValue) view returns(int256 value)
func (_Vm *VmSession) EnvOr9(name string, defaultValue *big.Int) (*big.Int, error) {
	return _Vm.Contract.EnvOr9(&_Vm.CallOpts, name, defaultValue)
}

// EnvOr9 is a free data retrieval call binding the contract method 0xbbcb713e.
//
// Solidity: function envOr(string name, int256 defaultValue) view returns(int256 value)
func (_Vm *VmCallerSession) EnvOr9(name string, defaultValue *big.Int) (*big.Int, error) {
	return _Vm.Contract.EnvOr9(&_Vm.CallOpts, name, defaultValue)
}

// EnvString is a free data retrieval call binding the contract method 0x14b02bc9.
//
// Solidity: function envString(string name, string delim) view returns(string[] value)
func (_Vm *VmCaller) EnvString(opts *bind.CallOpts, name string, delim string) ([]string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envString", name, delim)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// EnvString is a free data retrieval call binding the contract method 0x14b02bc9.
//
// Solidity: function envString(string name, string delim) view returns(string[] value)
func (_Vm *VmSession) EnvString(name string, delim string) ([]string, error) {
	return _Vm.Contract.EnvString(&_Vm.CallOpts, name, delim)
}

// EnvString is a free data retrieval call binding the contract method 0x14b02bc9.
//
// Solidity: function envString(string name, string delim) view returns(string[] value)
func (_Vm *VmCallerSession) EnvString(name string, delim string) ([]string, error) {
	return _Vm.Contract.EnvString(&_Vm.CallOpts, name, delim)
}

// EnvString0 is a free data retrieval call binding the contract method 0xf877cb19.
//
// Solidity: function envString(string name) view returns(string value)
func (_Vm *VmCaller) EnvString0(opts *bind.CallOpts, name string) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envString0", name)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// EnvString0 is a free data retrieval call binding the contract method 0xf877cb19.
//
// Solidity: function envString(string name) view returns(string value)
func (_Vm *VmSession) EnvString0(name string) (string, error) {
	return _Vm.Contract.EnvString0(&_Vm.CallOpts, name)
}

// EnvString0 is a free data retrieval call binding the contract method 0xf877cb19.
//
// Solidity: function envString(string name) view returns(string value)
func (_Vm *VmCallerSession) EnvString0(name string) (string, error) {
	return _Vm.Contract.EnvString0(&_Vm.CallOpts, name)
}

// EnvUint is a free data retrieval call binding the contract method 0xc1978d1f.
//
// Solidity: function envUint(string name) view returns(uint256 value)
func (_Vm *VmCaller) EnvUint(opts *bind.CallOpts, name string) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envUint", name)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EnvUint is a free data retrieval call binding the contract method 0xc1978d1f.
//
// Solidity: function envUint(string name) view returns(uint256 value)
func (_Vm *VmSession) EnvUint(name string) (*big.Int, error) {
	return _Vm.Contract.EnvUint(&_Vm.CallOpts, name)
}

// EnvUint is a free data retrieval call binding the contract method 0xc1978d1f.
//
// Solidity: function envUint(string name) view returns(uint256 value)
func (_Vm *VmCallerSession) EnvUint(name string) (*big.Int, error) {
	return _Vm.Contract.EnvUint(&_Vm.CallOpts, name)
}

// EnvUint0 is a free data retrieval call binding the contract method 0xf3dec099.
//
// Solidity: function envUint(string name, string delim) view returns(uint256[] value)
func (_Vm *VmCaller) EnvUint0(opts *bind.CallOpts, name string, delim string) ([]*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envUint0", name, delim)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// EnvUint0 is a free data retrieval call binding the contract method 0xf3dec099.
//
// Solidity: function envUint(string name, string delim) view returns(uint256[] value)
func (_Vm *VmSession) EnvUint0(name string, delim string) ([]*big.Int, error) {
	return _Vm.Contract.EnvUint0(&_Vm.CallOpts, name, delim)
}

// EnvUint0 is a free data retrieval call binding the contract method 0xf3dec099.
//
// Solidity: function envUint(string name, string delim) view returns(uint256[] value)
func (_Vm *VmCallerSession) EnvUint0(name string, delim string) ([]*big.Int, error) {
	return _Vm.Contract.EnvUint0(&_Vm.CallOpts, name, delim)
}

// FsMetadata is a free data retrieval call binding the contract method 0xaf368a08.
//
// Solidity: function fsMetadata(string path) view returns((bool,bool,uint256,bool,uint256,uint256,uint256) metadata)
func (_Vm *VmCaller) FsMetadata(opts *bind.CallOpts, path string) (VmSafeFsMetadata, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "fsMetadata", path)

	if err != nil {
		return *new(VmSafeFsMetadata), err
	}

	out0 := *abi.ConvertType(out[0], new(VmSafeFsMetadata)).(*VmSafeFsMetadata)

	return out0, err

}

// FsMetadata is a free data retrieval call binding the contract method 0xaf368a08.
//
// Solidity: function fsMetadata(string path) view returns((bool,bool,uint256,bool,uint256,uint256,uint256) metadata)
func (_Vm *VmSession) FsMetadata(path string) (VmSafeFsMetadata, error) {
	return _Vm.Contract.FsMetadata(&_Vm.CallOpts, path)
}

// FsMetadata is a free data retrieval call binding the contract method 0xaf368a08.
//
// Solidity: function fsMetadata(string path) view returns((bool,bool,uint256,bool,uint256,uint256,uint256) metadata)
func (_Vm *VmCallerSession) FsMetadata(path string) (VmSafeFsMetadata, error) {
	return _Vm.Contract.FsMetadata(&_Vm.CallOpts, path)
}

// GetBlobBaseFee is a free data retrieval call binding the contract method 0x1f6d6ef7.
//
// Solidity: function getBlobBaseFee() view returns(uint256 blobBaseFee)
func (_Vm *VmCaller) GetBlobBaseFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "getBlobBaseFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlobBaseFee is a free data retrieval call binding the contract method 0x1f6d6ef7.
//
// Solidity: function getBlobBaseFee() view returns(uint256 blobBaseFee)
func (_Vm *VmSession) GetBlobBaseFee() (*big.Int, error) {
	return _Vm.Contract.GetBlobBaseFee(&_Vm.CallOpts)
}

// GetBlobBaseFee is a free data retrieval call binding the contract method 0x1f6d6ef7.
//
// Solidity: function getBlobBaseFee() view returns(uint256 blobBaseFee)
func (_Vm *VmCallerSession) GetBlobBaseFee() (*big.Int, error) {
	return _Vm.Contract.GetBlobBaseFee(&_Vm.CallOpts)
}

// GetBlobhashes is a free data retrieval call binding the contract method 0xf56ff18b.
//
// Solidity: function getBlobhashes() view returns(bytes32[] hashes)
func (_Vm *VmCaller) GetBlobhashes(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "getBlobhashes")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetBlobhashes is a free data retrieval call binding the contract method 0xf56ff18b.
//
// Solidity: function getBlobhashes() view returns(bytes32[] hashes)
func (_Vm *VmSession) GetBlobhashes() ([][32]byte, error) {
	return _Vm.Contract.GetBlobhashes(&_Vm.CallOpts)
}

// GetBlobhashes is a free data retrieval call binding the contract method 0xf56ff18b.
//
// Solidity: function getBlobhashes() view returns(bytes32[] hashes)
func (_Vm *VmCallerSession) GetBlobhashes() ([][32]byte, error) {
	return _Vm.Contract.GetBlobhashes(&_Vm.CallOpts)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 height)
func (_Vm *VmCaller) GetBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "getBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 height)
func (_Vm *VmSession) GetBlockNumber() (*big.Int, error) {
	return _Vm.Contract.GetBlockNumber(&_Vm.CallOpts)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 height)
func (_Vm *VmCallerSession) GetBlockNumber() (*big.Int, error) {
	return _Vm.Contract.GetBlockNumber(&_Vm.CallOpts)
}

// GetBlockTimestamp is a free data retrieval call binding the contract method 0x796b89b9.
//
// Solidity: function getBlockTimestamp() view returns(uint256 timestamp)
func (_Vm *VmCaller) GetBlockTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "getBlockTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockTimestamp is a free data retrieval call binding the contract method 0x796b89b9.
//
// Solidity: function getBlockTimestamp() view returns(uint256 timestamp)
func (_Vm *VmSession) GetBlockTimestamp() (*big.Int, error) {
	return _Vm.Contract.GetBlockTimestamp(&_Vm.CallOpts)
}

// GetBlockTimestamp is a free data retrieval call binding the contract method 0x796b89b9.
//
// Solidity: function getBlockTimestamp() view returns(uint256 timestamp)
func (_Vm *VmCallerSession) GetBlockTimestamp() (*big.Int, error) {
	return _Vm.Contract.GetBlockTimestamp(&_Vm.CallOpts)
}

// GetCode is a free data retrieval call binding the contract method 0x8d1cc925.
//
// Solidity: function getCode(string artifactPath) view returns(bytes creationBytecode)
func (_Vm *VmCaller) GetCode(opts *bind.CallOpts, artifactPath string) ([]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "getCode", artifactPath)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetCode is a free data retrieval call binding the contract method 0x8d1cc925.
//
// Solidity: function getCode(string artifactPath) view returns(bytes creationBytecode)
func (_Vm *VmSession) GetCode(artifactPath string) ([]byte, error) {
	return _Vm.Contract.GetCode(&_Vm.CallOpts, artifactPath)
}

// GetCode is a free data retrieval call binding the contract method 0x8d1cc925.
//
// Solidity: function getCode(string artifactPath) view returns(bytes creationBytecode)
func (_Vm *VmCallerSession) GetCode(artifactPath string) ([]byte, error) {
	return _Vm.Contract.GetCode(&_Vm.CallOpts, artifactPath)
}

// GetDeployedCode is a free data retrieval call binding the contract method 0x3ebf73b4.
//
// Solidity: function getDeployedCode(string artifactPath) view returns(bytes runtimeBytecode)
func (_Vm *VmCaller) GetDeployedCode(opts *bind.CallOpts, artifactPath string) ([]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "getDeployedCode", artifactPath)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetDeployedCode is a free data retrieval call binding the contract method 0x3ebf73b4.
//
// Solidity: function getDeployedCode(string artifactPath) view returns(bytes runtimeBytecode)
func (_Vm *VmSession) GetDeployedCode(artifactPath string) ([]byte, error) {
	return _Vm.Contract.GetDeployedCode(&_Vm.CallOpts, artifactPath)
}

// GetDeployedCode is a free data retrieval call binding the contract method 0x3ebf73b4.
//
// Solidity: function getDeployedCode(string artifactPath) view returns(bytes runtimeBytecode)
func (_Vm *VmCallerSession) GetDeployedCode(artifactPath string) ([]byte, error) {
	return _Vm.Contract.GetDeployedCode(&_Vm.CallOpts, artifactPath)
}

// GetLabel is a free data retrieval call binding the contract method 0x28a249b0.
//
// Solidity: function getLabel(address account) view returns(string currentLabel)
func (_Vm *VmCaller) GetLabel(opts *bind.CallOpts, account common.Address) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "getLabel", account)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetLabel is a free data retrieval call binding the contract method 0x28a249b0.
//
// Solidity: function getLabel(address account) view returns(string currentLabel)
func (_Vm *VmSession) GetLabel(account common.Address) (string, error) {
	return _Vm.Contract.GetLabel(&_Vm.CallOpts, account)
}

// GetLabel is a free data retrieval call binding the contract method 0x28a249b0.
//
// Solidity: function getLabel(address account) view returns(string currentLabel)
func (_Vm *VmCallerSession) GetLabel(account common.Address) (string, error) {
	return _Vm.Contract.GetLabel(&_Vm.CallOpts, account)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address account) view returns(uint64 nonce)
func (_Vm *VmCaller) GetNonce(opts *bind.CallOpts, account common.Address) (uint64, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "getNonce", account)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address account) view returns(uint64 nonce)
func (_Vm *VmSession) GetNonce(account common.Address) (uint64, error) {
	return _Vm.Contract.GetNonce(&_Vm.CallOpts, account)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address account) view returns(uint64 nonce)
func (_Vm *VmCallerSession) GetNonce(account common.Address) (uint64, error) {
	return _Vm.Contract.GetNonce(&_Vm.CallOpts, account)
}

// IndexOf is a free data retrieval call binding the contract method 0x8a0807b7.
//
// Solidity: function indexOf(string input, string key) pure returns(uint256)
func (_Vm *VmCaller) IndexOf(opts *bind.CallOpts, input string, key string) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "indexOf", input, key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IndexOf is a free data retrieval call binding the contract method 0x8a0807b7.
//
// Solidity: function indexOf(string input, string key) pure returns(uint256)
func (_Vm *VmSession) IndexOf(input string, key string) (*big.Int, error) {
	return _Vm.Contract.IndexOf(&_Vm.CallOpts, input, key)
}

// IndexOf is a free data retrieval call binding the contract method 0x8a0807b7.
//
// Solidity: function indexOf(string input, string key) pure returns(uint256)
func (_Vm *VmCallerSession) IndexOf(input string, key string) (*big.Int, error) {
	return _Vm.Contract.IndexOf(&_Vm.CallOpts, input, key)
}

// IsContext is a free data retrieval call binding the contract method 0x64af255d.
//
// Solidity: function isContext(uint8 context) view returns(bool result)
func (_Vm *VmCaller) IsContext(opts *bind.CallOpts, context uint8) (bool, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "isContext", context)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsContext is a free data retrieval call binding the contract method 0x64af255d.
//
// Solidity: function isContext(uint8 context) view returns(bool result)
func (_Vm *VmSession) IsContext(context uint8) (bool, error) {
	return _Vm.Contract.IsContext(&_Vm.CallOpts, context)
}

// IsContext is a free data retrieval call binding the contract method 0x64af255d.
//
// Solidity: function isContext(uint8 context) view returns(bool result)
func (_Vm *VmCallerSession) IsContext(context uint8) (bool, error) {
	return _Vm.Contract.IsContext(&_Vm.CallOpts, context)
}

// IsPersistent is a free data retrieval call binding the contract method 0xd92d8efd.
//
// Solidity: function isPersistent(address account) view returns(bool persistent)
func (_Vm *VmCaller) IsPersistent(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "isPersistent", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPersistent is a free data retrieval call binding the contract method 0xd92d8efd.
//
// Solidity: function isPersistent(address account) view returns(bool persistent)
func (_Vm *VmSession) IsPersistent(account common.Address) (bool, error) {
	return _Vm.Contract.IsPersistent(&_Vm.CallOpts, account)
}

// IsPersistent is a free data retrieval call binding the contract method 0xd92d8efd.
//
// Solidity: function isPersistent(address account) view returns(bool persistent)
func (_Vm *VmCallerSession) IsPersistent(account common.Address) (bool, error) {
	return _Vm.Contract.IsPersistent(&_Vm.CallOpts, account)
}

// KeyExists is a free data retrieval call binding the contract method 0x528a683c.
//
// Solidity: function keyExists(string json, string key) view returns(bool)
func (_Vm *VmCaller) KeyExists(opts *bind.CallOpts, json string, key string) (bool, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "keyExists", json, key)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// KeyExists is a free data retrieval call binding the contract method 0x528a683c.
//
// Solidity: function keyExists(string json, string key) view returns(bool)
func (_Vm *VmSession) KeyExists(json string, key string) (bool, error) {
	return _Vm.Contract.KeyExists(&_Vm.CallOpts, json, key)
}

// KeyExists is a free data retrieval call binding the contract method 0x528a683c.
//
// Solidity: function keyExists(string json, string key) view returns(bool)
func (_Vm *VmCallerSession) KeyExists(json string, key string) (bool, error) {
	return _Vm.Contract.KeyExists(&_Vm.CallOpts, json, key)
}

// KeyExistsJson is a free data retrieval call binding the contract method 0xdb4235f6.
//
// Solidity: function keyExistsJson(string json, string key) view returns(bool)
func (_Vm *VmCaller) KeyExistsJson(opts *bind.CallOpts, json string, key string) (bool, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "keyExistsJson", json, key)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// KeyExistsJson is a free data retrieval call binding the contract method 0xdb4235f6.
//
// Solidity: function keyExistsJson(string json, string key) view returns(bool)
func (_Vm *VmSession) KeyExistsJson(json string, key string) (bool, error) {
	return _Vm.Contract.KeyExistsJson(&_Vm.CallOpts, json, key)
}

// KeyExistsJson is a free data retrieval call binding the contract method 0xdb4235f6.
//
// Solidity: function keyExistsJson(string json, string key) view returns(bool)
func (_Vm *VmCallerSession) KeyExistsJson(json string, key string) (bool, error) {
	return _Vm.Contract.KeyExistsJson(&_Vm.CallOpts, json, key)
}

// KeyExistsToml is a free data retrieval call binding the contract method 0x600903ad.
//
// Solidity: function keyExistsToml(string toml, string key) view returns(bool)
func (_Vm *VmCaller) KeyExistsToml(opts *bind.CallOpts, toml string, key string) (bool, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "keyExistsToml", toml, key)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// KeyExistsToml is a free data retrieval call binding the contract method 0x600903ad.
//
// Solidity: function keyExistsToml(string toml, string key) view returns(bool)
func (_Vm *VmSession) KeyExistsToml(toml string, key string) (bool, error) {
	return _Vm.Contract.KeyExistsToml(&_Vm.CallOpts, toml, key)
}

// KeyExistsToml is a free data retrieval call binding the contract method 0x600903ad.
//
// Solidity: function keyExistsToml(string toml, string key) view returns(bool)
func (_Vm *VmCallerSession) KeyExistsToml(toml string, key string) (bool, error) {
	return _Vm.Contract.KeyExistsToml(&_Vm.CallOpts, toml, key)
}

// LastCallGas is a free data retrieval call binding the contract method 0x2b589b28.
//
// Solidity: function lastCallGas() view returns((uint64,uint64,uint64,int64,uint64) gas)
func (_Vm *VmCaller) LastCallGas(opts *bind.CallOpts) (VmSafeGas, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "lastCallGas")

	if err != nil {
		return *new(VmSafeGas), err
	}

	out0 := *abi.ConvertType(out[0], new(VmSafeGas)).(*VmSafeGas)

	return out0, err

}

// LastCallGas is a free data retrieval call binding the contract method 0x2b589b28.
//
// Solidity: function lastCallGas() view returns((uint64,uint64,uint64,int64,uint64) gas)
func (_Vm *VmSession) LastCallGas() (VmSafeGas, error) {
	return _Vm.Contract.LastCallGas(&_Vm.CallOpts)
}

// LastCallGas is a free data retrieval call binding the contract method 0x2b589b28.
//
// Solidity: function lastCallGas() view returns((uint64,uint64,uint64,int64,uint64) gas)
func (_Vm *VmCallerSession) LastCallGas() (VmSafeGas, error) {
	return _Vm.Contract.LastCallGas(&_Vm.CallOpts)
}

// Load is a free data retrieval call binding the contract method 0x667f9d70.
//
// Solidity: function load(address target, bytes32 slot) view returns(bytes32 data)
func (_Vm *VmCaller) Load(opts *bind.CallOpts, target common.Address, slot [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "load", target, slot)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Load is a free data retrieval call binding the contract method 0x667f9d70.
//
// Solidity: function load(address target, bytes32 slot) view returns(bytes32 data)
func (_Vm *VmSession) Load(target common.Address, slot [32]byte) ([32]byte, error) {
	return _Vm.Contract.Load(&_Vm.CallOpts, target, slot)
}

// Load is a free data retrieval call binding the contract method 0x667f9d70.
//
// Solidity: function load(address target, bytes32 slot) view returns(bytes32 data)
func (_Vm *VmCallerSession) Load(target common.Address, slot [32]byte) ([32]byte, error) {
	return _Vm.Contract.Load(&_Vm.CallOpts, target, slot)
}

// ParseAddress is a free data retrieval call binding the contract method 0xc6ce059d.
//
// Solidity: function parseAddress(string stringifiedValue) pure returns(address parsedValue)
func (_Vm *VmCaller) ParseAddress(opts *bind.CallOpts, stringifiedValue string) (common.Address, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseAddress", stringifiedValue)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ParseAddress is a free data retrieval call binding the contract method 0xc6ce059d.
//
// Solidity: function parseAddress(string stringifiedValue) pure returns(address parsedValue)
func (_Vm *VmSession) ParseAddress(stringifiedValue string) (common.Address, error) {
	return _Vm.Contract.ParseAddress(&_Vm.CallOpts, stringifiedValue)
}

// ParseAddress is a free data retrieval call binding the contract method 0xc6ce059d.
//
// Solidity: function parseAddress(string stringifiedValue) pure returns(address parsedValue)
func (_Vm *VmCallerSession) ParseAddress(stringifiedValue string) (common.Address, error) {
	return _Vm.Contract.ParseAddress(&_Vm.CallOpts, stringifiedValue)
}

// ParseBool is a free data retrieval call binding the contract method 0x974ef924.
//
// Solidity: function parseBool(string stringifiedValue) pure returns(bool parsedValue)
func (_Vm *VmCaller) ParseBool(opts *bind.CallOpts, stringifiedValue string) (bool, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseBool", stringifiedValue)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ParseBool is a free data retrieval call binding the contract method 0x974ef924.
//
// Solidity: function parseBool(string stringifiedValue) pure returns(bool parsedValue)
func (_Vm *VmSession) ParseBool(stringifiedValue string) (bool, error) {
	return _Vm.Contract.ParseBool(&_Vm.CallOpts, stringifiedValue)
}

// ParseBool is a free data retrieval call binding the contract method 0x974ef924.
//
// Solidity: function parseBool(string stringifiedValue) pure returns(bool parsedValue)
func (_Vm *VmCallerSession) ParseBool(stringifiedValue string) (bool, error) {
	return _Vm.Contract.ParseBool(&_Vm.CallOpts, stringifiedValue)
}

// ParseBytes is a free data retrieval call binding the contract method 0x8f5d232d.
//
// Solidity: function parseBytes(string stringifiedValue) pure returns(bytes parsedValue)
func (_Vm *VmCaller) ParseBytes(opts *bind.CallOpts, stringifiedValue string) ([]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseBytes", stringifiedValue)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ParseBytes is a free data retrieval call binding the contract method 0x8f5d232d.
//
// Solidity: function parseBytes(string stringifiedValue) pure returns(bytes parsedValue)
func (_Vm *VmSession) ParseBytes(stringifiedValue string) ([]byte, error) {
	return _Vm.Contract.ParseBytes(&_Vm.CallOpts, stringifiedValue)
}

// ParseBytes is a free data retrieval call binding the contract method 0x8f5d232d.
//
// Solidity: function parseBytes(string stringifiedValue) pure returns(bytes parsedValue)
func (_Vm *VmCallerSession) ParseBytes(stringifiedValue string) ([]byte, error) {
	return _Vm.Contract.ParseBytes(&_Vm.CallOpts, stringifiedValue)
}

// ParseBytes32 is a free data retrieval call binding the contract method 0x087e6e81.
//
// Solidity: function parseBytes32(string stringifiedValue) pure returns(bytes32 parsedValue)
func (_Vm *VmCaller) ParseBytes32(opts *bind.CallOpts, stringifiedValue string) ([32]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseBytes32", stringifiedValue)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ParseBytes32 is a free data retrieval call binding the contract method 0x087e6e81.
//
// Solidity: function parseBytes32(string stringifiedValue) pure returns(bytes32 parsedValue)
func (_Vm *VmSession) ParseBytes32(stringifiedValue string) ([32]byte, error) {
	return _Vm.Contract.ParseBytes32(&_Vm.CallOpts, stringifiedValue)
}

// ParseBytes32 is a free data retrieval call binding the contract method 0x087e6e81.
//
// Solidity: function parseBytes32(string stringifiedValue) pure returns(bytes32 parsedValue)
func (_Vm *VmCallerSession) ParseBytes32(stringifiedValue string) ([32]byte, error) {
	return _Vm.Contract.ParseBytes32(&_Vm.CallOpts, stringifiedValue)
}

// ParseInt is a free data retrieval call binding the contract method 0x42346c5e.
//
// Solidity: function parseInt(string stringifiedValue) pure returns(int256 parsedValue)
func (_Vm *VmCaller) ParseInt(opts *bind.CallOpts, stringifiedValue string) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseInt", stringifiedValue)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ParseInt is a free data retrieval call binding the contract method 0x42346c5e.
//
// Solidity: function parseInt(string stringifiedValue) pure returns(int256 parsedValue)
func (_Vm *VmSession) ParseInt(stringifiedValue string) (*big.Int, error) {
	return _Vm.Contract.ParseInt(&_Vm.CallOpts, stringifiedValue)
}

// ParseInt is a free data retrieval call binding the contract method 0x42346c5e.
//
// Solidity: function parseInt(string stringifiedValue) pure returns(int256 parsedValue)
func (_Vm *VmCallerSession) ParseInt(stringifiedValue string) (*big.Int, error) {
	return _Vm.Contract.ParseInt(&_Vm.CallOpts, stringifiedValue)
}

// ParseJson is a free data retrieval call binding the contract method 0x6a82600a.
//
// Solidity: function parseJson(string json) pure returns(bytes abiEncodedData)
func (_Vm *VmCaller) ParseJson(opts *bind.CallOpts, json string) ([]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJson", json)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ParseJson is a free data retrieval call binding the contract method 0x6a82600a.
//
// Solidity: function parseJson(string json) pure returns(bytes abiEncodedData)
func (_Vm *VmSession) ParseJson(json string) ([]byte, error) {
	return _Vm.Contract.ParseJson(&_Vm.CallOpts, json)
}

// ParseJson is a free data retrieval call binding the contract method 0x6a82600a.
//
// Solidity: function parseJson(string json) pure returns(bytes abiEncodedData)
func (_Vm *VmCallerSession) ParseJson(json string) ([]byte, error) {
	return _Vm.Contract.ParseJson(&_Vm.CallOpts, json)
}

// ParseJson0 is a free data retrieval call binding the contract method 0x85940ef1.
//
// Solidity: function parseJson(string json, string key) pure returns(bytes abiEncodedData)
func (_Vm *VmCaller) ParseJson0(opts *bind.CallOpts, json string, key string) ([]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJson0", json, key)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ParseJson0 is a free data retrieval call binding the contract method 0x85940ef1.
//
// Solidity: function parseJson(string json, string key) pure returns(bytes abiEncodedData)
func (_Vm *VmSession) ParseJson0(json string, key string) ([]byte, error) {
	return _Vm.Contract.ParseJson0(&_Vm.CallOpts, json, key)
}

// ParseJson0 is a free data retrieval call binding the contract method 0x85940ef1.
//
// Solidity: function parseJson(string json, string key) pure returns(bytes abiEncodedData)
func (_Vm *VmCallerSession) ParseJson0(json string, key string) ([]byte, error) {
	return _Vm.Contract.ParseJson0(&_Vm.CallOpts, json, key)
}

// ParseJsonAddress is a free data retrieval call binding the contract method 0x1e19e657.
//
// Solidity: function parseJsonAddress(string json, string key) pure returns(address)
func (_Vm *VmCaller) ParseJsonAddress(opts *bind.CallOpts, json string, key string) (common.Address, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJsonAddress", json, key)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ParseJsonAddress is a free data retrieval call binding the contract method 0x1e19e657.
//
// Solidity: function parseJsonAddress(string json, string key) pure returns(address)
func (_Vm *VmSession) ParseJsonAddress(json string, key string) (common.Address, error) {
	return _Vm.Contract.ParseJsonAddress(&_Vm.CallOpts, json, key)
}

// ParseJsonAddress is a free data retrieval call binding the contract method 0x1e19e657.
//
// Solidity: function parseJsonAddress(string json, string key) pure returns(address)
func (_Vm *VmCallerSession) ParseJsonAddress(json string, key string) (common.Address, error) {
	return _Vm.Contract.ParseJsonAddress(&_Vm.CallOpts, json, key)
}

// ParseJsonAddressArray is a free data retrieval call binding the contract method 0x2fce7883.
//
// Solidity: function parseJsonAddressArray(string json, string key) pure returns(address[])
func (_Vm *VmCaller) ParseJsonAddressArray(opts *bind.CallOpts, json string, key string) ([]common.Address, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJsonAddressArray", json, key)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ParseJsonAddressArray is a free data retrieval call binding the contract method 0x2fce7883.
//
// Solidity: function parseJsonAddressArray(string json, string key) pure returns(address[])
func (_Vm *VmSession) ParseJsonAddressArray(json string, key string) ([]common.Address, error) {
	return _Vm.Contract.ParseJsonAddressArray(&_Vm.CallOpts, json, key)
}

// ParseJsonAddressArray is a free data retrieval call binding the contract method 0x2fce7883.
//
// Solidity: function parseJsonAddressArray(string json, string key) pure returns(address[])
func (_Vm *VmCallerSession) ParseJsonAddressArray(json string, key string) ([]common.Address, error) {
	return _Vm.Contract.ParseJsonAddressArray(&_Vm.CallOpts, json, key)
}

// ParseJsonBool is a free data retrieval call binding the contract method 0x9f86dc91.
//
// Solidity: function parseJsonBool(string json, string key) pure returns(bool)
func (_Vm *VmCaller) ParseJsonBool(opts *bind.CallOpts, json string, key string) (bool, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJsonBool", json, key)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ParseJsonBool is a free data retrieval call binding the contract method 0x9f86dc91.
//
// Solidity: function parseJsonBool(string json, string key) pure returns(bool)
func (_Vm *VmSession) ParseJsonBool(json string, key string) (bool, error) {
	return _Vm.Contract.ParseJsonBool(&_Vm.CallOpts, json, key)
}

// ParseJsonBool is a free data retrieval call binding the contract method 0x9f86dc91.
//
// Solidity: function parseJsonBool(string json, string key) pure returns(bool)
func (_Vm *VmCallerSession) ParseJsonBool(json string, key string) (bool, error) {
	return _Vm.Contract.ParseJsonBool(&_Vm.CallOpts, json, key)
}

// ParseJsonBoolArray is a free data retrieval call binding the contract method 0x91f3b94f.
//
// Solidity: function parseJsonBoolArray(string json, string key) pure returns(bool[])
func (_Vm *VmCaller) ParseJsonBoolArray(opts *bind.CallOpts, json string, key string) ([]bool, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJsonBoolArray", json, key)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// ParseJsonBoolArray is a free data retrieval call binding the contract method 0x91f3b94f.
//
// Solidity: function parseJsonBoolArray(string json, string key) pure returns(bool[])
func (_Vm *VmSession) ParseJsonBoolArray(json string, key string) ([]bool, error) {
	return _Vm.Contract.ParseJsonBoolArray(&_Vm.CallOpts, json, key)
}

// ParseJsonBoolArray is a free data retrieval call binding the contract method 0x91f3b94f.
//
// Solidity: function parseJsonBoolArray(string json, string key) pure returns(bool[])
func (_Vm *VmCallerSession) ParseJsonBoolArray(json string, key string) ([]bool, error) {
	return _Vm.Contract.ParseJsonBoolArray(&_Vm.CallOpts, json, key)
}

// ParseJsonBytes is a free data retrieval call binding the contract method 0xfd921be8.
//
// Solidity: function parseJsonBytes(string json, string key) pure returns(bytes)
func (_Vm *VmCaller) ParseJsonBytes(opts *bind.CallOpts, json string, key string) ([]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJsonBytes", json, key)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ParseJsonBytes is a free data retrieval call binding the contract method 0xfd921be8.
//
// Solidity: function parseJsonBytes(string json, string key) pure returns(bytes)
func (_Vm *VmSession) ParseJsonBytes(json string, key string) ([]byte, error) {
	return _Vm.Contract.ParseJsonBytes(&_Vm.CallOpts, json, key)
}

// ParseJsonBytes is a free data retrieval call binding the contract method 0xfd921be8.
//
// Solidity: function parseJsonBytes(string json, string key) pure returns(bytes)
func (_Vm *VmCallerSession) ParseJsonBytes(json string, key string) ([]byte, error) {
	return _Vm.Contract.ParseJsonBytes(&_Vm.CallOpts, json, key)
}

// ParseJsonBytes32 is a free data retrieval call binding the contract method 0x1777e59d.
//
// Solidity: function parseJsonBytes32(string json, string key) pure returns(bytes32)
func (_Vm *VmCaller) ParseJsonBytes32(opts *bind.CallOpts, json string, key string) ([32]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJsonBytes32", json, key)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ParseJsonBytes32 is a free data retrieval call binding the contract method 0x1777e59d.
//
// Solidity: function parseJsonBytes32(string json, string key) pure returns(bytes32)
func (_Vm *VmSession) ParseJsonBytes32(json string, key string) ([32]byte, error) {
	return _Vm.Contract.ParseJsonBytes32(&_Vm.CallOpts, json, key)
}

// ParseJsonBytes32 is a free data retrieval call binding the contract method 0x1777e59d.
//
// Solidity: function parseJsonBytes32(string json, string key) pure returns(bytes32)
func (_Vm *VmCallerSession) ParseJsonBytes32(json string, key string) ([32]byte, error) {
	return _Vm.Contract.ParseJsonBytes32(&_Vm.CallOpts, json, key)
}

// ParseJsonBytes32Array is a free data retrieval call binding the contract method 0x91c75bc3.
//
// Solidity: function parseJsonBytes32Array(string json, string key) pure returns(bytes32[])
func (_Vm *VmCaller) ParseJsonBytes32Array(opts *bind.CallOpts, json string, key string) ([][32]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJsonBytes32Array", json, key)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// ParseJsonBytes32Array is a free data retrieval call binding the contract method 0x91c75bc3.
//
// Solidity: function parseJsonBytes32Array(string json, string key) pure returns(bytes32[])
func (_Vm *VmSession) ParseJsonBytes32Array(json string, key string) ([][32]byte, error) {
	return _Vm.Contract.ParseJsonBytes32Array(&_Vm.CallOpts, json, key)
}

// ParseJsonBytes32Array is a free data retrieval call binding the contract method 0x91c75bc3.
//
// Solidity: function parseJsonBytes32Array(string json, string key) pure returns(bytes32[])
func (_Vm *VmCallerSession) ParseJsonBytes32Array(json string, key string) ([][32]byte, error) {
	return _Vm.Contract.ParseJsonBytes32Array(&_Vm.CallOpts, json, key)
}

// ParseJsonBytesArray is a free data retrieval call binding the contract method 0x6631aa99.
//
// Solidity: function parseJsonBytesArray(string json, string key) pure returns(bytes[])
func (_Vm *VmCaller) ParseJsonBytesArray(opts *bind.CallOpts, json string, key string) ([][]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJsonBytesArray", json, key)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// ParseJsonBytesArray is a free data retrieval call binding the contract method 0x6631aa99.
//
// Solidity: function parseJsonBytesArray(string json, string key) pure returns(bytes[])
func (_Vm *VmSession) ParseJsonBytesArray(json string, key string) ([][]byte, error) {
	return _Vm.Contract.ParseJsonBytesArray(&_Vm.CallOpts, json, key)
}

// ParseJsonBytesArray is a free data retrieval call binding the contract method 0x6631aa99.
//
// Solidity: function parseJsonBytesArray(string json, string key) pure returns(bytes[])
func (_Vm *VmCallerSession) ParseJsonBytesArray(json string, key string) ([][]byte, error) {
	return _Vm.Contract.ParseJsonBytesArray(&_Vm.CallOpts, json, key)
}

// ParseJsonInt is a free data retrieval call binding the contract method 0x7b048ccd.
//
// Solidity: function parseJsonInt(string json, string key) pure returns(int256)
func (_Vm *VmCaller) ParseJsonInt(opts *bind.CallOpts, json string, key string) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJsonInt", json, key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ParseJsonInt is a free data retrieval call binding the contract method 0x7b048ccd.
//
// Solidity: function parseJsonInt(string json, string key) pure returns(int256)
func (_Vm *VmSession) ParseJsonInt(json string, key string) (*big.Int, error) {
	return _Vm.Contract.ParseJsonInt(&_Vm.CallOpts, json, key)
}

// ParseJsonInt is a free data retrieval call binding the contract method 0x7b048ccd.
//
// Solidity: function parseJsonInt(string json, string key) pure returns(int256)
func (_Vm *VmCallerSession) ParseJsonInt(json string, key string) (*big.Int, error) {
	return _Vm.Contract.ParseJsonInt(&_Vm.CallOpts, json, key)
}

// ParseJsonIntArray is a free data retrieval call binding the contract method 0x9983c28a.
//
// Solidity: function parseJsonIntArray(string json, string key) pure returns(int256[])
func (_Vm *VmCaller) ParseJsonIntArray(opts *bind.CallOpts, json string, key string) ([]*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJsonIntArray", json, key)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// ParseJsonIntArray is a free data retrieval call binding the contract method 0x9983c28a.
//
// Solidity: function parseJsonIntArray(string json, string key) pure returns(int256[])
func (_Vm *VmSession) ParseJsonIntArray(json string, key string) ([]*big.Int, error) {
	return _Vm.Contract.ParseJsonIntArray(&_Vm.CallOpts, json, key)
}

// ParseJsonIntArray is a free data retrieval call binding the contract method 0x9983c28a.
//
// Solidity: function parseJsonIntArray(string json, string key) pure returns(int256[])
func (_Vm *VmCallerSession) ParseJsonIntArray(json string, key string) ([]*big.Int, error) {
	return _Vm.Contract.ParseJsonIntArray(&_Vm.CallOpts, json, key)
}

// ParseJsonKeys is a free data retrieval call binding the contract method 0x213e4198.
//
// Solidity: function parseJsonKeys(string json, string key) pure returns(string[] keys)
func (_Vm *VmCaller) ParseJsonKeys(opts *bind.CallOpts, json string, key string) ([]string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJsonKeys", json, key)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// ParseJsonKeys is a free data retrieval call binding the contract method 0x213e4198.
//
// Solidity: function parseJsonKeys(string json, string key) pure returns(string[] keys)
func (_Vm *VmSession) ParseJsonKeys(json string, key string) ([]string, error) {
	return _Vm.Contract.ParseJsonKeys(&_Vm.CallOpts, json, key)
}

// ParseJsonKeys is a free data retrieval call binding the contract method 0x213e4198.
//
// Solidity: function parseJsonKeys(string json, string key) pure returns(string[] keys)
func (_Vm *VmCallerSession) ParseJsonKeys(json string, key string) ([]string, error) {
	return _Vm.Contract.ParseJsonKeys(&_Vm.CallOpts, json, key)
}

// ParseJsonString is a free data retrieval call binding the contract method 0x49c4fac8.
//
// Solidity: function parseJsonString(string json, string key) pure returns(string)
func (_Vm *VmCaller) ParseJsonString(opts *bind.CallOpts, json string, key string) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJsonString", json, key)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ParseJsonString is a free data retrieval call binding the contract method 0x49c4fac8.
//
// Solidity: function parseJsonString(string json, string key) pure returns(string)
func (_Vm *VmSession) ParseJsonString(json string, key string) (string, error) {
	return _Vm.Contract.ParseJsonString(&_Vm.CallOpts, json, key)
}

// ParseJsonString is a free data retrieval call binding the contract method 0x49c4fac8.
//
// Solidity: function parseJsonString(string json, string key) pure returns(string)
func (_Vm *VmCallerSession) ParseJsonString(json string, key string) (string, error) {
	return _Vm.Contract.ParseJsonString(&_Vm.CallOpts, json, key)
}

// ParseJsonStringArray is a free data retrieval call binding the contract method 0x498fdcf4.
//
// Solidity: function parseJsonStringArray(string json, string key) pure returns(string[])
func (_Vm *VmCaller) ParseJsonStringArray(opts *bind.CallOpts, json string, key string) ([]string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJsonStringArray", json, key)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// ParseJsonStringArray is a free data retrieval call binding the contract method 0x498fdcf4.
//
// Solidity: function parseJsonStringArray(string json, string key) pure returns(string[])
func (_Vm *VmSession) ParseJsonStringArray(json string, key string) ([]string, error) {
	return _Vm.Contract.ParseJsonStringArray(&_Vm.CallOpts, json, key)
}

// ParseJsonStringArray is a free data retrieval call binding the contract method 0x498fdcf4.
//
// Solidity: function parseJsonStringArray(string json, string key) pure returns(string[])
func (_Vm *VmCallerSession) ParseJsonStringArray(json string, key string) ([]string, error) {
	return _Vm.Contract.ParseJsonStringArray(&_Vm.CallOpts, json, key)
}

// ParseJsonUint is a free data retrieval call binding the contract method 0xaddde2b6.
//
// Solidity: function parseJsonUint(string json, string key) pure returns(uint256)
func (_Vm *VmCaller) ParseJsonUint(opts *bind.CallOpts, json string, key string) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJsonUint", json, key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ParseJsonUint is a free data retrieval call binding the contract method 0xaddde2b6.
//
// Solidity: function parseJsonUint(string json, string key) pure returns(uint256)
func (_Vm *VmSession) ParseJsonUint(json string, key string) (*big.Int, error) {
	return _Vm.Contract.ParseJsonUint(&_Vm.CallOpts, json, key)
}

// ParseJsonUint is a free data retrieval call binding the contract method 0xaddde2b6.
//
// Solidity: function parseJsonUint(string json, string key) pure returns(uint256)
func (_Vm *VmCallerSession) ParseJsonUint(json string, key string) (*big.Int, error) {
	return _Vm.Contract.ParseJsonUint(&_Vm.CallOpts, json, key)
}

// ParseJsonUintArray is a free data retrieval call binding the contract method 0x522074ab.
//
// Solidity: function parseJsonUintArray(string json, string key) pure returns(uint256[])
func (_Vm *VmCaller) ParseJsonUintArray(opts *bind.CallOpts, json string, key string) ([]*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJsonUintArray", json, key)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// ParseJsonUintArray is a free data retrieval call binding the contract method 0x522074ab.
//
// Solidity: function parseJsonUintArray(string json, string key) pure returns(uint256[])
func (_Vm *VmSession) ParseJsonUintArray(json string, key string) ([]*big.Int, error) {
	return _Vm.Contract.ParseJsonUintArray(&_Vm.CallOpts, json, key)
}

// ParseJsonUintArray is a free data retrieval call binding the contract method 0x522074ab.
//
// Solidity: function parseJsonUintArray(string json, string key) pure returns(uint256[])
func (_Vm *VmCallerSession) ParseJsonUintArray(json string, key string) ([]*big.Int, error) {
	return _Vm.Contract.ParseJsonUintArray(&_Vm.CallOpts, json, key)
}

// ParseToml is a free data retrieval call binding the contract method 0x37736e08.
//
// Solidity: function parseToml(string toml, string key) pure returns(bytes abiEncodedData)
func (_Vm *VmCaller) ParseToml(opts *bind.CallOpts, toml string, key string) ([]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseToml", toml, key)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ParseToml is a free data retrieval call binding the contract method 0x37736e08.
//
// Solidity: function parseToml(string toml, string key) pure returns(bytes abiEncodedData)
func (_Vm *VmSession) ParseToml(toml string, key string) ([]byte, error) {
	return _Vm.Contract.ParseToml(&_Vm.CallOpts, toml, key)
}

// ParseToml is a free data retrieval call binding the contract method 0x37736e08.
//
// Solidity: function parseToml(string toml, string key) pure returns(bytes abiEncodedData)
func (_Vm *VmCallerSession) ParseToml(toml string, key string) ([]byte, error) {
	return _Vm.Contract.ParseToml(&_Vm.CallOpts, toml, key)
}

// ParseToml0 is a free data retrieval call binding the contract method 0x592151f0.
//
// Solidity: function parseToml(string toml) pure returns(bytes abiEncodedData)
func (_Vm *VmCaller) ParseToml0(opts *bind.CallOpts, toml string) ([]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseToml0", toml)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ParseToml0 is a free data retrieval call binding the contract method 0x592151f0.
//
// Solidity: function parseToml(string toml) pure returns(bytes abiEncodedData)
func (_Vm *VmSession) ParseToml0(toml string) ([]byte, error) {
	return _Vm.Contract.ParseToml0(&_Vm.CallOpts, toml)
}

// ParseToml0 is a free data retrieval call binding the contract method 0x592151f0.
//
// Solidity: function parseToml(string toml) pure returns(bytes abiEncodedData)
func (_Vm *VmCallerSession) ParseToml0(toml string) ([]byte, error) {
	return _Vm.Contract.ParseToml0(&_Vm.CallOpts, toml)
}

// ParseTomlAddress is a free data retrieval call binding the contract method 0x65e7c844.
//
// Solidity: function parseTomlAddress(string toml, string key) pure returns(address)
func (_Vm *VmCaller) ParseTomlAddress(opts *bind.CallOpts, toml string, key string) (common.Address, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseTomlAddress", toml, key)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ParseTomlAddress is a free data retrieval call binding the contract method 0x65e7c844.
//
// Solidity: function parseTomlAddress(string toml, string key) pure returns(address)
func (_Vm *VmSession) ParseTomlAddress(toml string, key string) (common.Address, error) {
	return _Vm.Contract.ParseTomlAddress(&_Vm.CallOpts, toml, key)
}

// ParseTomlAddress is a free data retrieval call binding the contract method 0x65e7c844.
//
// Solidity: function parseTomlAddress(string toml, string key) pure returns(address)
func (_Vm *VmCallerSession) ParseTomlAddress(toml string, key string) (common.Address, error) {
	return _Vm.Contract.ParseTomlAddress(&_Vm.CallOpts, toml, key)
}

// ParseTomlAddressArray is a free data retrieval call binding the contract method 0x65c428e7.
//
// Solidity: function parseTomlAddressArray(string toml, string key) pure returns(address[])
func (_Vm *VmCaller) ParseTomlAddressArray(opts *bind.CallOpts, toml string, key string) ([]common.Address, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseTomlAddressArray", toml, key)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ParseTomlAddressArray is a free data retrieval call binding the contract method 0x65c428e7.
//
// Solidity: function parseTomlAddressArray(string toml, string key) pure returns(address[])
func (_Vm *VmSession) ParseTomlAddressArray(toml string, key string) ([]common.Address, error) {
	return _Vm.Contract.ParseTomlAddressArray(&_Vm.CallOpts, toml, key)
}

// ParseTomlAddressArray is a free data retrieval call binding the contract method 0x65c428e7.
//
// Solidity: function parseTomlAddressArray(string toml, string key) pure returns(address[])
func (_Vm *VmCallerSession) ParseTomlAddressArray(toml string, key string) ([]common.Address, error) {
	return _Vm.Contract.ParseTomlAddressArray(&_Vm.CallOpts, toml, key)
}

// ParseTomlBool is a free data retrieval call binding the contract method 0xd30dced6.
//
// Solidity: function parseTomlBool(string toml, string key) pure returns(bool)
func (_Vm *VmCaller) ParseTomlBool(opts *bind.CallOpts, toml string, key string) (bool, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseTomlBool", toml, key)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ParseTomlBool is a free data retrieval call binding the contract method 0xd30dced6.
//
// Solidity: function parseTomlBool(string toml, string key) pure returns(bool)
func (_Vm *VmSession) ParseTomlBool(toml string, key string) (bool, error) {
	return _Vm.Contract.ParseTomlBool(&_Vm.CallOpts, toml, key)
}

// ParseTomlBool is a free data retrieval call binding the contract method 0xd30dced6.
//
// Solidity: function parseTomlBool(string toml, string key) pure returns(bool)
func (_Vm *VmCallerSession) ParseTomlBool(toml string, key string) (bool, error) {
	return _Vm.Contract.ParseTomlBool(&_Vm.CallOpts, toml, key)
}

// ParseTomlBoolArray is a free data retrieval call binding the contract method 0x127cfe9a.
//
// Solidity: function parseTomlBoolArray(string toml, string key) pure returns(bool[])
func (_Vm *VmCaller) ParseTomlBoolArray(opts *bind.CallOpts, toml string, key string) ([]bool, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseTomlBoolArray", toml, key)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// ParseTomlBoolArray is a free data retrieval call binding the contract method 0x127cfe9a.
//
// Solidity: function parseTomlBoolArray(string toml, string key) pure returns(bool[])
func (_Vm *VmSession) ParseTomlBoolArray(toml string, key string) ([]bool, error) {
	return _Vm.Contract.ParseTomlBoolArray(&_Vm.CallOpts, toml, key)
}

// ParseTomlBoolArray is a free data retrieval call binding the contract method 0x127cfe9a.
//
// Solidity: function parseTomlBoolArray(string toml, string key) pure returns(bool[])
func (_Vm *VmCallerSession) ParseTomlBoolArray(toml string, key string) ([]bool, error) {
	return _Vm.Contract.ParseTomlBoolArray(&_Vm.CallOpts, toml, key)
}

// ParseTomlBytes is a free data retrieval call binding the contract method 0xd77bfdb9.
//
// Solidity: function parseTomlBytes(string toml, string key) pure returns(bytes)
func (_Vm *VmCaller) ParseTomlBytes(opts *bind.CallOpts, toml string, key string) ([]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseTomlBytes", toml, key)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ParseTomlBytes is a free data retrieval call binding the contract method 0xd77bfdb9.
//
// Solidity: function parseTomlBytes(string toml, string key) pure returns(bytes)
func (_Vm *VmSession) ParseTomlBytes(toml string, key string) ([]byte, error) {
	return _Vm.Contract.ParseTomlBytes(&_Vm.CallOpts, toml, key)
}

// ParseTomlBytes is a free data retrieval call binding the contract method 0xd77bfdb9.
//
// Solidity: function parseTomlBytes(string toml, string key) pure returns(bytes)
func (_Vm *VmCallerSession) ParseTomlBytes(toml string, key string) ([]byte, error) {
	return _Vm.Contract.ParseTomlBytes(&_Vm.CallOpts, toml, key)
}

// ParseTomlBytes32 is a free data retrieval call binding the contract method 0x8e214810.
//
// Solidity: function parseTomlBytes32(string toml, string key) pure returns(bytes32)
func (_Vm *VmCaller) ParseTomlBytes32(opts *bind.CallOpts, toml string, key string) ([32]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseTomlBytes32", toml, key)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ParseTomlBytes32 is a free data retrieval call binding the contract method 0x8e214810.
//
// Solidity: function parseTomlBytes32(string toml, string key) pure returns(bytes32)
func (_Vm *VmSession) ParseTomlBytes32(toml string, key string) ([32]byte, error) {
	return _Vm.Contract.ParseTomlBytes32(&_Vm.CallOpts, toml, key)
}

// ParseTomlBytes32 is a free data retrieval call binding the contract method 0x8e214810.
//
// Solidity: function parseTomlBytes32(string toml, string key) pure returns(bytes32)
func (_Vm *VmCallerSession) ParseTomlBytes32(toml string, key string) ([32]byte, error) {
	return _Vm.Contract.ParseTomlBytes32(&_Vm.CallOpts, toml, key)
}

// ParseTomlBytes32Array is a free data retrieval call binding the contract method 0x3e716f81.
//
// Solidity: function parseTomlBytes32Array(string toml, string key) pure returns(bytes32[])
func (_Vm *VmCaller) ParseTomlBytes32Array(opts *bind.CallOpts, toml string, key string) ([][32]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseTomlBytes32Array", toml, key)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// ParseTomlBytes32Array is a free data retrieval call binding the contract method 0x3e716f81.
//
// Solidity: function parseTomlBytes32Array(string toml, string key) pure returns(bytes32[])
func (_Vm *VmSession) ParseTomlBytes32Array(toml string, key string) ([][32]byte, error) {
	return _Vm.Contract.ParseTomlBytes32Array(&_Vm.CallOpts, toml, key)
}

// ParseTomlBytes32Array is a free data retrieval call binding the contract method 0x3e716f81.
//
// Solidity: function parseTomlBytes32Array(string toml, string key) pure returns(bytes32[])
func (_Vm *VmCallerSession) ParseTomlBytes32Array(toml string, key string) ([][32]byte, error) {
	return _Vm.Contract.ParseTomlBytes32Array(&_Vm.CallOpts, toml, key)
}

// ParseTomlBytesArray is a free data retrieval call binding the contract method 0xb197c247.
//
// Solidity: function parseTomlBytesArray(string toml, string key) pure returns(bytes[])
func (_Vm *VmCaller) ParseTomlBytesArray(opts *bind.CallOpts, toml string, key string) ([][]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseTomlBytesArray", toml, key)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// ParseTomlBytesArray is a free data retrieval call binding the contract method 0xb197c247.
//
// Solidity: function parseTomlBytesArray(string toml, string key) pure returns(bytes[])
func (_Vm *VmSession) ParseTomlBytesArray(toml string, key string) ([][]byte, error) {
	return _Vm.Contract.ParseTomlBytesArray(&_Vm.CallOpts, toml, key)
}

// ParseTomlBytesArray is a free data retrieval call binding the contract method 0xb197c247.
//
// Solidity: function parseTomlBytesArray(string toml, string key) pure returns(bytes[])
func (_Vm *VmCallerSession) ParseTomlBytesArray(toml string, key string) ([][]byte, error) {
	return _Vm.Contract.ParseTomlBytesArray(&_Vm.CallOpts, toml, key)
}

// ParseTomlInt is a free data retrieval call binding the contract method 0xc1350739.
//
// Solidity: function parseTomlInt(string toml, string key) pure returns(int256)
func (_Vm *VmCaller) ParseTomlInt(opts *bind.CallOpts, toml string, key string) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseTomlInt", toml, key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ParseTomlInt is a free data retrieval call binding the contract method 0xc1350739.
//
// Solidity: function parseTomlInt(string toml, string key) pure returns(int256)
func (_Vm *VmSession) ParseTomlInt(toml string, key string) (*big.Int, error) {
	return _Vm.Contract.ParseTomlInt(&_Vm.CallOpts, toml, key)
}

// ParseTomlInt is a free data retrieval call binding the contract method 0xc1350739.
//
// Solidity: function parseTomlInt(string toml, string key) pure returns(int256)
func (_Vm *VmCallerSession) ParseTomlInt(toml string, key string) (*big.Int, error) {
	return _Vm.Contract.ParseTomlInt(&_Vm.CallOpts, toml, key)
}

// ParseTomlIntArray is a free data retrieval call binding the contract method 0xd3522ae6.
//
// Solidity: function parseTomlIntArray(string toml, string key) pure returns(int256[])
func (_Vm *VmCaller) ParseTomlIntArray(opts *bind.CallOpts, toml string, key string) ([]*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseTomlIntArray", toml, key)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// ParseTomlIntArray is a free data retrieval call binding the contract method 0xd3522ae6.
//
// Solidity: function parseTomlIntArray(string toml, string key) pure returns(int256[])
func (_Vm *VmSession) ParseTomlIntArray(toml string, key string) ([]*big.Int, error) {
	return _Vm.Contract.ParseTomlIntArray(&_Vm.CallOpts, toml, key)
}

// ParseTomlIntArray is a free data retrieval call binding the contract method 0xd3522ae6.
//
// Solidity: function parseTomlIntArray(string toml, string key) pure returns(int256[])
func (_Vm *VmCallerSession) ParseTomlIntArray(toml string, key string) ([]*big.Int, error) {
	return _Vm.Contract.ParseTomlIntArray(&_Vm.CallOpts, toml, key)
}

// ParseTomlKeys is a free data retrieval call binding the contract method 0x812a44b2.
//
// Solidity: function parseTomlKeys(string toml, string key) pure returns(string[] keys)
func (_Vm *VmCaller) ParseTomlKeys(opts *bind.CallOpts, toml string, key string) ([]string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseTomlKeys", toml, key)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// ParseTomlKeys is a free data retrieval call binding the contract method 0x812a44b2.
//
// Solidity: function parseTomlKeys(string toml, string key) pure returns(string[] keys)
func (_Vm *VmSession) ParseTomlKeys(toml string, key string) ([]string, error) {
	return _Vm.Contract.ParseTomlKeys(&_Vm.CallOpts, toml, key)
}

// ParseTomlKeys is a free data retrieval call binding the contract method 0x812a44b2.
//
// Solidity: function parseTomlKeys(string toml, string key) pure returns(string[] keys)
func (_Vm *VmCallerSession) ParseTomlKeys(toml string, key string) ([]string, error) {
	return _Vm.Contract.ParseTomlKeys(&_Vm.CallOpts, toml, key)
}

// ParseTomlString is a free data retrieval call binding the contract method 0x8bb8dd43.
//
// Solidity: function parseTomlString(string toml, string key) pure returns(string)
func (_Vm *VmCaller) ParseTomlString(opts *bind.CallOpts, toml string, key string) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseTomlString", toml, key)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ParseTomlString is a free data retrieval call binding the contract method 0x8bb8dd43.
//
// Solidity: function parseTomlString(string toml, string key) pure returns(string)
func (_Vm *VmSession) ParseTomlString(toml string, key string) (string, error) {
	return _Vm.Contract.ParseTomlString(&_Vm.CallOpts, toml, key)
}

// ParseTomlString is a free data retrieval call binding the contract method 0x8bb8dd43.
//
// Solidity: function parseTomlString(string toml, string key) pure returns(string)
func (_Vm *VmCallerSession) ParseTomlString(toml string, key string) (string, error) {
	return _Vm.Contract.ParseTomlString(&_Vm.CallOpts, toml, key)
}

// ParseTomlStringArray is a free data retrieval call binding the contract method 0x9f629281.
//
// Solidity: function parseTomlStringArray(string toml, string key) pure returns(string[])
func (_Vm *VmCaller) ParseTomlStringArray(opts *bind.CallOpts, toml string, key string) ([]string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseTomlStringArray", toml, key)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// ParseTomlStringArray is a free data retrieval call binding the contract method 0x9f629281.
//
// Solidity: function parseTomlStringArray(string toml, string key) pure returns(string[])
func (_Vm *VmSession) ParseTomlStringArray(toml string, key string) ([]string, error) {
	return _Vm.Contract.ParseTomlStringArray(&_Vm.CallOpts, toml, key)
}

// ParseTomlStringArray is a free data retrieval call binding the contract method 0x9f629281.
//
// Solidity: function parseTomlStringArray(string toml, string key) pure returns(string[])
func (_Vm *VmCallerSession) ParseTomlStringArray(toml string, key string) ([]string, error) {
	return _Vm.Contract.ParseTomlStringArray(&_Vm.CallOpts, toml, key)
}

// ParseTomlUint is a free data retrieval call binding the contract method 0xcc7b0487.
//
// Solidity: function parseTomlUint(string toml, string key) pure returns(uint256)
func (_Vm *VmCaller) ParseTomlUint(opts *bind.CallOpts, toml string, key string) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseTomlUint", toml, key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ParseTomlUint is a free data retrieval call binding the contract method 0xcc7b0487.
//
// Solidity: function parseTomlUint(string toml, string key) pure returns(uint256)
func (_Vm *VmSession) ParseTomlUint(toml string, key string) (*big.Int, error) {
	return _Vm.Contract.ParseTomlUint(&_Vm.CallOpts, toml, key)
}

// ParseTomlUint is a free data retrieval call binding the contract method 0xcc7b0487.
//
// Solidity: function parseTomlUint(string toml, string key) pure returns(uint256)
func (_Vm *VmCallerSession) ParseTomlUint(toml string, key string) (*big.Int, error) {
	return _Vm.Contract.ParseTomlUint(&_Vm.CallOpts, toml, key)
}

// ParseTomlUintArray is a free data retrieval call binding the contract method 0xb5df27c8.
//
// Solidity: function parseTomlUintArray(string toml, string key) pure returns(uint256[])
func (_Vm *VmCaller) ParseTomlUintArray(opts *bind.CallOpts, toml string, key string) ([]*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseTomlUintArray", toml, key)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// ParseTomlUintArray is a free data retrieval call binding the contract method 0xb5df27c8.
//
// Solidity: function parseTomlUintArray(string toml, string key) pure returns(uint256[])
func (_Vm *VmSession) ParseTomlUintArray(toml string, key string) ([]*big.Int, error) {
	return _Vm.Contract.ParseTomlUintArray(&_Vm.CallOpts, toml, key)
}

// ParseTomlUintArray is a free data retrieval call binding the contract method 0xb5df27c8.
//
// Solidity: function parseTomlUintArray(string toml, string key) pure returns(uint256[])
func (_Vm *VmCallerSession) ParseTomlUintArray(toml string, key string) ([]*big.Int, error) {
	return _Vm.Contract.ParseTomlUintArray(&_Vm.CallOpts, toml, key)
}

// ParseUint is a free data retrieval call binding the contract method 0xfa91454d.
//
// Solidity: function parseUint(string stringifiedValue) pure returns(uint256 parsedValue)
func (_Vm *VmCaller) ParseUint(opts *bind.CallOpts, stringifiedValue string) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseUint", stringifiedValue)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ParseUint is a free data retrieval call binding the contract method 0xfa91454d.
//
// Solidity: function parseUint(string stringifiedValue) pure returns(uint256 parsedValue)
func (_Vm *VmSession) ParseUint(stringifiedValue string) (*big.Int, error) {
	return _Vm.Contract.ParseUint(&_Vm.CallOpts, stringifiedValue)
}

// ParseUint is a free data retrieval call binding the contract method 0xfa91454d.
//
// Solidity: function parseUint(string stringifiedValue) pure returns(uint256 parsedValue)
func (_Vm *VmCallerSession) ParseUint(stringifiedValue string) (*big.Int, error) {
	return _Vm.Contract.ParseUint(&_Vm.CallOpts, stringifiedValue)
}

// ProjectRoot is a free data retrieval call binding the contract method 0xd930a0e6.
//
// Solidity: function projectRoot() view returns(string path)
func (_Vm *VmCaller) ProjectRoot(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "projectRoot")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ProjectRoot is a free data retrieval call binding the contract method 0xd930a0e6.
//
// Solidity: function projectRoot() view returns(string path)
func (_Vm *VmSession) ProjectRoot() (string, error) {
	return _Vm.Contract.ProjectRoot(&_Vm.CallOpts)
}

// ProjectRoot is a free data retrieval call binding the contract method 0xd930a0e6.
//
// Solidity: function projectRoot() view returns(string path)
func (_Vm *VmCallerSession) ProjectRoot() (string, error) {
	return _Vm.Contract.ProjectRoot(&_Vm.CallOpts)
}

// ReadDir is a free data retrieval call binding the contract method 0x1497876c.
//
// Solidity: function readDir(string path, uint64 maxDepth) view returns((string,string,uint64,bool,bool)[] entries)
func (_Vm *VmCaller) ReadDir(opts *bind.CallOpts, path string, maxDepth uint64) ([]VmSafeDirEntry, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "readDir", path, maxDepth)

	if err != nil {
		return *new([]VmSafeDirEntry), err
	}

	out0 := *abi.ConvertType(out[0], new([]VmSafeDirEntry)).(*[]VmSafeDirEntry)

	return out0, err

}

// ReadDir is a free data retrieval call binding the contract method 0x1497876c.
//
// Solidity: function readDir(string path, uint64 maxDepth) view returns((string,string,uint64,bool,bool)[] entries)
func (_Vm *VmSession) ReadDir(path string, maxDepth uint64) ([]VmSafeDirEntry, error) {
	return _Vm.Contract.ReadDir(&_Vm.CallOpts, path, maxDepth)
}

// ReadDir is a free data retrieval call binding the contract method 0x1497876c.
//
// Solidity: function readDir(string path, uint64 maxDepth) view returns((string,string,uint64,bool,bool)[] entries)
func (_Vm *VmCallerSession) ReadDir(path string, maxDepth uint64) ([]VmSafeDirEntry, error) {
	return _Vm.Contract.ReadDir(&_Vm.CallOpts, path, maxDepth)
}

// ReadDir0 is a free data retrieval call binding the contract method 0x8102d70d.
//
// Solidity: function readDir(string path, uint64 maxDepth, bool followLinks) view returns((string,string,uint64,bool,bool)[] entries)
func (_Vm *VmCaller) ReadDir0(opts *bind.CallOpts, path string, maxDepth uint64, followLinks bool) ([]VmSafeDirEntry, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "readDir0", path, maxDepth, followLinks)

	if err != nil {
		return *new([]VmSafeDirEntry), err
	}

	out0 := *abi.ConvertType(out[0], new([]VmSafeDirEntry)).(*[]VmSafeDirEntry)

	return out0, err

}

// ReadDir0 is a free data retrieval call binding the contract method 0x8102d70d.
//
// Solidity: function readDir(string path, uint64 maxDepth, bool followLinks) view returns((string,string,uint64,bool,bool)[] entries)
func (_Vm *VmSession) ReadDir0(path string, maxDepth uint64, followLinks bool) ([]VmSafeDirEntry, error) {
	return _Vm.Contract.ReadDir0(&_Vm.CallOpts, path, maxDepth, followLinks)
}

// ReadDir0 is a free data retrieval call binding the contract method 0x8102d70d.
//
// Solidity: function readDir(string path, uint64 maxDepth, bool followLinks) view returns((string,string,uint64,bool,bool)[] entries)
func (_Vm *VmCallerSession) ReadDir0(path string, maxDepth uint64, followLinks bool) ([]VmSafeDirEntry, error) {
	return _Vm.Contract.ReadDir0(&_Vm.CallOpts, path, maxDepth, followLinks)
}

// ReadDir1 is a free data retrieval call binding the contract method 0xc4bc59e0.
//
// Solidity: function readDir(string path) view returns((string,string,uint64,bool,bool)[] entries)
func (_Vm *VmCaller) ReadDir1(opts *bind.CallOpts, path string) ([]VmSafeDirEntry, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "readDir1", path)

	if err != nil {
		return *new([]VmSafeDirEntry), err
	}

	out0 := *abi.ConvertType(out[0], new([]VmSafeDirEntry)).(*[]VmSafeDirEntry)

	return out0, err

}

// ReadDir1 is a free data retrieval call binding the contract method 0xc4bc59e0.
//
// Solidity: function readDir(string path) view returns((string,string,uint64,bool,bool)[] entries)
func (_Vm *VmSession) ReadDir1(path string) ([]VmSafeDirEntry, error) {
	return _Vm.Contract.ReadDir1(&_Vm.CallOpts, path)
}

// ReadDir1 is a free data retrieval call binding the contract method 0xc4bc59e0.
//
// Solidity: function readDir(string path) view returns((string,string,uint64,bool,bool)[] entries)
func (_Vm *VmCallerSession) ReadDir1(path string) ([]VmSafeDirEntry, error) {
	return _Vm.Contract.ReadDir1(&_Vm.CallOpts, path)
}

// ReadFile is a free data retrieval call binding the contract method 0x60f9bb11.
//
// Solidity: function readFile(string path) view returns(string data)
func (_Vm *VmCaller) ReadFile(opts *bind.CallOpts, path string) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "readFile", path)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ReadFile is a free data retrieval call binding the contract method 0x60f9bb11.
//
// Solidity: function readFile(string path) view returns(string data)
func (_Vm *VmSession) ReadFile(path string) (string, error) {
	return _Vm.Contract.ReadFile(&_Vm.CallOpts, path)
}

// ReadFile is a free data retrieval call binding the contract method 0x60f9bb11.
//
// Solidity: function readFile(string path) view returns(string data)
func (_Vm *VmCallerSession) ReadFile(path string) (string, error) {
	return _Vm.Contract.ReadFile(&_Vm.CallOpts, path)
}

// ReadFileBinary is a free data retrieval call binding the contract method 0x16ed7bc4.
//
// Solidity: function readFileBinary(string path) view returns(bytes data)
func (_Vm *VmCaller) ReadFileBinary(opts *bind.CallOpts, path string) ([]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "readFileBinary", path)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ReadFileBinary is a free data retrieval call binding the contract method 0x16ed7bc4.
//
// Solidity: function readFileBinary(string path) view returns(bytes data)
func (_Vm *VmSession) ReadFileBinary(path string) ([]byte, error) {
	return _Vm.Contract.ReadFileBinary(&_Vm.CallOpts, path)
}

// ReadFileBinary is a free data retrieval call binding the contract method 0x16ed7bc4.
//
// Solidity: function readFileBinary(string path) view returns(bytes data)
func (_Vm *VmCallerSession) ReadFileBinary(path string) ([]byte, error) {
	return _Vm.Contract.ReadFileBinary(&_Vm.CallOpts, path)
}

// ReadLine is a free data retrieval call binding the contract method 0x70f55728.
//
// Solidity: function readLine(string path) view returns(string line)
func (_Vm *VmCaller) ReadLine(opts *bind.CallOpts, path string) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "readLine", path)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ReadLine is a free data retrieval call binding the contract method 0x70f55728.
//
// Solidity: function readLine(string path) view returns(string line)
func (_Vm *VmSession) ReadLine(path string) (string, error) {
	return _Vm.Contract.ReadLine(&_Vm.CallOpts, path)
}

// ReadLine is a free data retrieval call binding the contract method 0x70f55728.
//
// Solidity: function readLine(string path) view returns(string line)
func (_Vm *VmCallerSession) ReadLine(path string) (string, error) {
	return _Vm.Contract.ReadLine(&_Vm.CallOpts, path)
}

// ReadLink is a free data retrieval call binding the contract method 0x9f5684a2.
//
// Solidity: function readLink(string linkPath) view returns(string targetPath)
func (_Vm *VmCaller) ReadLink(opts *bind.CallOpts, linkPath string) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "readLink", linkPath)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ReadLink is a free data retrieval call binding the contract method 0x9f5684a2.
//
// Solidity: function readLink(string linkPath) view returns(string targetPath)
func (_Vm *VmSession) ReadLink(linkPath string) (string, error) {
	return _Vm.Contract.ReadLink(&_Vm.CallOpts, linkPath)
}

// ReadLink is a free data retrieval call binding the contract method 0x9f5684a2.
//
// Solidity: function readLink(string linkPath) view returns(string targetPath)
func (_Vm *VmCallerSession) ReadLink(linkPath string) (string, error) {
	return _Vm.Contract.ReadLink(&_Vm.CallOpts, linkPath)
}

// Replace is a free data retrieval call binding the contract method 0xe00ad03e.
//
// Solidity: function replace(string input, string from, string to) pure returns(string output)
func (_Vm *VmCaller) Replace(opts *bind.CallOpts, input string, from string, to string) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "replace", input, from, to)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Replace is a free data retrieval call binding the contract method 0xe00ad03e.
//
// Solidity: function replace(string input, string from, string to) pure returns(string output)
func (_Vm *VmSession) Replace(input string, from string, to string) (string, error) {
	return _Vm.Contract.Replace(&_Vm.CallOpts, input, from, to)
}

// Replace is a free data retrieval call binding the contract method 0xe00ad03e.
//
// Solidity: function replace(string input, string from, string to) pure returns(string output)
func (_Vm *VmCallerSession) Replace(input string, from string, to string) (string, error) {
	return _Vm.Contract.Replace(&_Vm.CallOpts, input, from, to)
}

// RpcUrl is a free data retrieval call binding the contract method 0x975a6ce9.
//
// Solidity: function rpcUrl(string rpcAlias) view returns(string json)
func (_Vm *VmCaller) RpcUrl(opts *bind.CallOpts, rpcAlias string) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "rpcUrl", rpcAlias)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// RpcUrl is a free data retrieval call binding the contract method 0x975a6ce9.
//
// Solidity: function rpcUrl(string rpcAlias) view returns(string json)
func (_Vm *VmSession) RpcUrl(rpcAlias string) (string, error) {
	return _Vm.Contract.RpcUrl(&_Vm.CallOpts, rpcAlias)
}

// RpcUrl is a free data retrieval call binding the contract method 0x975a6ce9.
//
// Solidity: function rpcUrl(string rpcAlias) view returns(string json)
func (_Vm *VmCallerSession) RpcUrl(rpcAlias string) (string, error) {
	return _Vm.Contract.RpcUrl(&_Vm.CallOpts, rpcAlias)
}

// RpcUrlStructs is a free data retrieval call binding the contract method 0x9d2ad72a.
//
// Solidity: function rpcUrlStructs() view returns((string,string)[] urls)
func (_Vm *VmCaller) RpcUrlStructs(opts *bind.CallOpts) ([]VmSafeRpc, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "rpcUrlStructs")

	if err != nil {
		return *new([]VmSafeRpc), err
	}

	out0 := *abi.ConvertType(out[0], new([]VmSafeRpc)).(*[]VmSafeRpc)

	return out0, err

}

// RpcUrlStructs is a free data retrieval call binding the contract method 0x9d2ad72a.
//
// Solidity: function rpcUrlStructs() view returns((string,string)[] urls)
func (_Vm *VmSession) RpcUrlStructs() ([]VmSafeRpc, error) {
	return _Vm.Contract.RpcUrlStructs(&_Vm.CallOpts)
}

// RpcUrlStructs is a free data retrieval call binding the contract method 0x9d2ad72a.
//
// Solidity: function rpcUrlStructs() view returns((string,string)[] urls)
func (_Vm *VmCallerSession) RpcUrlStructs() ([]VmSafeRpc, error) {
	return _Vm.Contract.RpcUrlStructs(&_Vm.CallOpts)
}

// RpcUrls is a free data retrieval call binding the contract method 0xa85a8418.
//
// Solidity: function rpcUrls() view returns(string[2][] urls)
func (_Vm *VmCaller) RpcUrls(opts *bind.CallOpts) ([][2]string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "rpcUrls")

	if err != nil {
		return *new([][2]string), err
	}

	out0 := *abi.ConvertType(out[0], new([][2]string)).(*[][2]string)

	return out0, err

}

// RpcUrls is a free data retrieval call binding the contract method 0xa85a8418.
//
// Solidity: function rpcUrls() view returns(string[2][] urls)
func (_Vm *VmSession) RpcUrls() ([][2]string, error) {
	return _Vm.Contract.RpcUrls(&_Vm.CallOpts)
}

// RpcUrls is a free data retrieval call binding the contract method 0xa85a8418.
//
// Solidity: function rpcUrls() view returns(string[2][] urls)
func (_Vm *VmCallerSession) RpcUrls() ([][2]string, error) {
	return _Vm.Contract.RpcUrls(&_Vm.CallOpts)
}

// Sign is a free data retrieval call binding the contract method 0x799cd333.
//
// Solidity: function sign(bytes32 digest) pure returns(uint8 v, bytes32 r, bytes32 s)
func (_Vm *VmCaller) Sign(opts *bind.CallOpts, digest [32]byte) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "sign", digest)

	outstruct := new(struct {
		V uint8
		R [32]byte
		S [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.V = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.R = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.S = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Sign is a free data retrieval call binding the contract method 0x799cd333.
//
// Solidity: function sign(bytes32 digest) pure returns(uint8 v, bytes32 r, bytes32 s)
func (_Vm *VmSession) Sign(digest [32]byte) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	return _Vm.Contract.Sign(&_Vm.CallOpts, digest)
}

// Sign is a free data retrieval call binding the contract method 0x799cd333.
//
// Solidity: function sign(bytes32 digest) pure returns(uint8 v, bytes32 r, bytes32 s)
func (_Vm *VmCallerSession) Sign(digest [32]byte) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	return _Vm.Contract.Sign(&_Vm.CallOpts, digest)
}

// Sign0 is a free data retrieval call binding the contract method 0x8c1aa205.
//
// Solidity: function sign(address signer, bytes32 digest) pure returns(uint8 v, bytes32 r, bytes32 s)
func (_Vm *VmCaller) Sign0(opts *bind.CallOpts, signer common.Address, digest [32]byte) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "sign0", signer, digest)

	outstruct := new(struct {
		V uint8
		R [32]byte
		S [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.V = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.R = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.S = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Sign0 is a free data retrieval call binding the contract method 0x8c1aa205.
//
// Solidity: function sign(address signer, bytes32 digest) pure returns(uint8 v, bytes32 r, bytes32 s)
func (_Vm *VmSession) Sign0(signer common.Address, digest [32]byte) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	return _Vm.Contract.Sign0(&_Vm.CallOpts, signer, digest)
}

// Sign0 is a free data retrieval call binding the contract method 0x8c1aa205.
//
// Solidity: function sign(address signer, bytes32 digest) pure returns(uint8 v, bytes32 r, bytes32 s)
func (_Vm *VmCallerSession) Sign0(signer common.Address, digest [32]byte) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	return _Vm.Contract.Sign0(&_Vm.CallOpts, signer, digest)
}

// Sign2 is a free data retrieval call binding the contract method 0xe341eaa4.
//
// Solidity: function sign(uint256 privateKey, bytes32 digest) pure returns(uint8 v, bytes32 r, bytes32 s)
func (_Vm *VmCaller) Sign2(opts *bind.CallOpts, privateKey *big.Int, digest [32]byte) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "sign2", privateKey, digest)

	outstruct := new(struct {
		V uint8
		R [32]byte
		S [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.V = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.R = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.S = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Sign2 is a free data retrieval call binding the contract method 0xe341eaa4.
//
// Solidity: function sign(uint256 privateKey, bytes32 digest) pure returns(uint8 v, bytes32 r, bytes32 s)
func (_Vm *VmSession) Sign2(privateKey *big.Int, digest [32]byte) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	return _Vm.Contract.Sign2(&_Vm.CallOpts, privateKey, digest)
}

// Sign2 is a free data retrieval call binding the contract method 0xe341eaa4.
//
// Solidity: function sign(uint256 privateKey, bytes32 digest) pure returns(uint8 v, bytes32 r, bytes32 s)
func (_Vm *VmCallerSession) Sign2(privateKey *big.Int, digest [32]byte) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	return _Vm.Contract.Sign2(&_Vm.CallOpts, privateKey, digest)
}

// SignP256 is a free data retrieval call binding the contract method 0x83211b40.
//
// Solidity: function signP256(uint256 privateKey, bytes32 digest) pure returns(bytes32 r, bytes32 s)
func (_Vm *VmCaller) SignP256(opts *bind.CallOpts, privateKey *big.Int, digest [32]byte) (struct {
	R [32]byte
	S [32]byte
}, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "signP256", privateKey, digest)

	outstruct := new(struct {
		R [32]byte
		S [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.R = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.S = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// SignP256 is a free data retrieval call binding the contract method 0x83211b40.
//
// Solidity: function signP256(uint256 privateKey, bytes32 digest) pure returns(bytes32 r, bytes32 s)
func (_Vm *VmSession) SignP256(privateKey *big.Int, digest [32]byte) (struct {
	R [32]byte
	S [32]byte
}, error) {
	return _Vm.Contract.SignP256(&_Vm.CallOpts, privateKey, digest)
}

// SignP256 is a free data retrieval call binding the contract method 0x83211b40.
//
// Solidity: function signP256(uint256 privateKey, bytes32 digest) pure returns(bytes32 r, bytes32 s)
func (_Vm *VmCallerSession) SignP256(privateKey *big.Int, digest [32]byte) (struct {
	R [32]byte
	S [32]byte
}, error) {
	return _Vm.Contract.SignP256(&_Vm.CallOpts, privateKey, digest)
}

// Split is a free data retrieval call binding the contract method 0x8bb75533.
//
// Solidity: function split(string input, string delimiter) pure returns(string[] outputs)
func (_Vm *VmCaller) Split(opts *bind.CallOpts, input string, delimiter string) ([]string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "split", input, delimiter)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// Split is a free data retrieval call binding the contract method 0x8bb75533.
//
// Solidity: function split(string input, string delimiter) pure returns(string[] outputs)
func (_Vm *VmSession) Split(input string, delimiter string) ([]string, error) {
	return _Vm.Contract.Split(&_Vm.CallOpts, input, delimiter)
}

// Split is a free data retrieval call binding the contract method 0x8bb75533.
//
// Solidity: function split(string input, string delimiter) pure returns(string[] outputs)
func (_Vm *VmCallerSession) Split(input string, delimiter string) ([]string, error) {
	return _Vm.Contract.Split(&_Vm.CallOpts, input, delimiter)
}

// ToBase64 is a free data retrieval call binding the contract method 0x3f8be2c8.
//
// Solidity: function toBase64(string data) pure returns(string)
func (_Vm *VmCaller) ToBase64(opts *bind.CallOpts, data string) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "toBase64", data)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToBase64 is a free data retrieval call binding the contract method 0x3f8be2c8.
//
// Solidity: function toBase64(string data) pure returns(string)
func (_Vm *VmSession) ToBase64(data string) (string, error) {
	return _Vm.Contract.ToBase64(&_Vm.CallOpts, data)
}

// ToBase64 is a free data retrieval call binding the contract method 0x3f8be2c8.
//
// Solidity: function toBase64(string data) pure returns(string)
func (_Vm *VmCallerSession) ToBase64(data string) (string, error) {
	return _Vm.Contract.ToBase64(&_Vm.CallOpts, data)
}

// ToBase640 is a free data retrieval call binding the contract method 0xa5cbfe65.
//
// Solidity: function toBase64(bytes data) pure returns(string)
func (_Vm *VmCaller) ToBase640(opts *bind.CallOpts, data []byte) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "toBase640", data)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToBase640 is a free data retrieval call binding the contract method 0xa5cbfe65.
//
// Solidity: function toBase64(bytes data) pure returns(string)
func (_Vm *VmSession) ToBase640(data []byte) (string, error) {
	return _Vm.Contract.ToBase640(&_Vm.CallOpts, data)
}

// ToBase640 is a free data retrieval call binding the contract method 0xa5cbfe65.
//
// Solidity: function toBase64(bytes data) pure returns(string)
func (_Vm *VmCallerSession) ToBase640(data []byte) (string, error) {
	return _Vm.Contract.ToBase640(&_Vm.CallOpts, data)
}

// ToBase64URL is a free data retrieval call binding the contract method 0xae3165b3.
//
// Solidity: function toBase64URL(string data) pure returns(string)
func (_Vm *VmCaller) ToBase64URL(opts *bind.CallOpts, data string) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "toBase64URL", data)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToBase64URL is a free data retrieval call binding the contract method 0xae3165b3.
//
// Solidity: function toBase64URL(string data) pure returns(string)
func (_Vm *VmSession) ToBase64URL(data string) (string, error) {
	return _Vm.Contract.ToBase64URL(&_Vm.CallOpts, data)
}

// ToBase64URL is a free data retrieval call binding the contract method 0xae3165b3.
//
// Solidity: function toBase64URL(string data) pure returns(string)
func (_Vm *VmCallerSession) ToBase64URL(data string) (string, error) {
	return _Vm.Contract.ToBase64URL(&_Vm.CallOpts, data)
}

// ToBase64URL0 is a free data retrieval call binding the contract method 0xc8bd0e4a.
//
// Solidity: function toBase64URL(bytes data) pure returns(string)
func (_Vm *VmCaller) ToBase64URL0(opts *bind.CallOpts, data []byte) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "toBase64URL0", data)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToBase64URL0 is a free data retrieval call binding the contract method 0xc8bd0e4a.
//
// Solidity: function toBase64URL(bytes data) pure returns(string)
func (_Vm *VmSession) ToBase64URL0(data []byte) (string, error) {
	return _Vm.Contract.ToBase64URL0(&_Vm.CallOpts, data)
}

// ToBase64URL0 is a free data retrieval call binding the contract method 0xc8bd0e4a.
//
// Solidity: function toBase64URL(bytes data) pure returns(string)
func (_Vm *VmCallerSession) ToBase64URL0(data []byte) (string, error) {
	return _Vm.Contract.ToBase64URL0(&_Vm.CallOpts, data)
}

// ToLowercase is a free data retrieval call binding the contract method 0x50bb0884.
//
// Solidity: function toLowercase(string input) pure returns(string output)
func (_Vm *VmCaller) ToLowercase(opts *bind.CallOpts, input string) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "toLowercase", input)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToLowercase is a free data retrieval call binding the contract method 0x50bb0884.
//
// Solidity: function toLowercase(string input) pure returns(string output)
func (_Vm *VmSession) ToLowercase(input string) (string, error) {
	return _Vm.Contract.ToLowercase(&_Vm.CallOpts, input)
}

// ToLowercase is a free data retrieval call binding the contract method 0x50bb0884.
//
// Solidity: function toLowercase(string input) pure returns(string output)
func (_Vm *VmCallerSession) ToLowercase(input string) (string, error) {
	return _Vm.Contract.ToLowercase(&_Vm.CallOpts, input)
}

// ToString is a free data retrieval call binding the contract method 0x56ca623e.
//
// Solidity: function toString(address value) pure returns(string stringifiedValue)
func (_Vm *VmCaller) ToString(opts *bind.CallOpts, value common.Address) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "toString", value)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToString is a free data retrieval call binding the contract method 0x56ca623e.
//
// Solidity: function toString(address value) pure returns(string stringifiedValue)
func (_Vm *VmSession) ToString(value common.Address) (string, error) {
	return _Vm.Contract.ToString(&_Vm.CallOpts, value)
}

// ToString is a free data retrieval call binding the contract method 0x56ca623e.
//
// Solidity: function toString(address value) pure returns(string stringifiedValue)
func (_Vm *VmCallerSession) ToString(value common.Address) (string, error) {
	return _Vm.Contract.ToString(&_Vm.CallOpts, value)
}

// ToString0 is a free data retrieval call binding the contract method 0x6900a3ae.
//
// Solidity: function toString(uint256 value) pure returns(string stringifiedValue)
func (_Vm *VmCaller) ToString0(opts *bind.CallOpts, value *big.Int) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "toString0", value)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToString0 is a free data retrieval call binding the contract method 0x6900a3ae.
//
// Solidity: function toString(uint256 value) pure returns(string stringifiedValue)
func (_Vm *VmSession) ToString0(value *big.Int) (string, error) {
	return _Vm.Contract.ToString0(&_Vm.CallOpts, value)
}

// ToString0 is a free data retrieval call binding the contract method 0x6900a3ae.
//
// Solidity: function toString(uint256 value) pure returns(string stringifiedValue)
func (_Vm *VmCallerSession) ToString0(value *big.Int) (string, error) {
	return _Vm.Contract.ToString0(&_Vm.CallOpts, value)
}

// ToString1 is a free data retrieval call binding the contract method 0x71aad10d.
//
// Solidity: function toString(bytes value) pure returns(string stringifiedValue)
func (_Vm *VmCaller) ToString1(opts *bind.CallOpts, value []byte) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "toString1", value)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToString1 is a free data retrieval call binding the contract method 0x71aad10d.
//
// Solidity: function toString(bytes value) pure returns(string stringifiedValue)
func (_Vm *VmSession) ToString1(value []byte) (string, error) {
	return _Vm.Contract.ToString1(&_Vm.CallOpts, value)
}

// ToString1 is a free data retrieval call binding the contract method 0x71aad10d.
//
// Solidity: function toString(bytes value) pure returns(string stringifiedValue)
func (_Vm *VmCallerSession) ToString1(value []byte) (string, error) {
	return _Vm.Contract.ToString1(&_Vm.CallOpts, value)
}

// ToString2 is a free data retrieval call binding the contract method 0x71dce7da.
//
// Solidity: function toString(bool value) pure returns(string stringifiedValue)
func (_Vm *VmCaller) ToString2(opts *bind.CallOpts, value bool) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "toString2", value)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToString2 is a free data retrieval call binding the contract method 0x71dce7da.
//
// Solidity: function toString(bool value) pure returns(string stringifiedValue)
func (_Vm *VmSession) ToString2(value bool) (string, error) {
	return _Vm.Contract.ToString2(&_Vm.CallOpts, value)
}

// ToString2 is a free data retrieval call binding the contract method 0x71dce7da.
//
// Solidity: function toString(bool value) pure returns(string stringifiedValue)
func (_Vm *VmCallerSession) ToString2(value bool) (string, error) {
	return _Vm.Contract.ToString2(&_Vm.CallOpts, value)
}

// ToString3 is a free data retrieval call binding the contract method 0xa322c40e.
//
// Solidity: function toString(int256 value) pure returns(string stringifiedValue)
func (_Vm *VmCaller) ToString3(opts *bind.CallOpts, value *big.Int) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "toString3", value)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToString3 is a free data retrieval call binding the contract method 0xa322c40e.
//
// Solidity: function toString(int256 value) pure returns(string stringifiedValue)
func (_Vm *VmSession) ToString3(value *big.Int) (string, error) {
	return _Vm.Contract.ToString3(&_Vm.CallOpts, value)
}

// ToString3 is a free data retrieval call binding the contract method 0xa322c40e.
//
// Solidity: function toString(int256 value) pure returns(string stringifiedValue)
func (_Vm *VmCallerSession) ToString3(value *big.Int) (string, error) {
	return _Vm.Contract.ToString3(&_Vm.CallOpts, value)
}

// ToString4 is a free data retrieval call binding the contract method 0xb11a19e8.
//
// Solidity: function toString(bytes32 value) pure returns(string stringifiedValue)
func (_Vm *VmCaller) ToString4(opts *bind.CallOpts, value [32]byte) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "toString4", value)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToString4 is a free data retrieval call binding the contract method 0xb11a19e8.
//
// Solidity: function toString(bytes32 value) pure returns(string stringifiedValue)
func (_Vm *VmSession) ToString4(value [32]byte) (string, error) {
	return _Vm.Contract.ToString4(&_Vm.CallOpts, value)
}

// ToString4 is a free data retrieval call binding the contract method 0xb11a19e8.
//
// Solidity: function toString(bytes32 value) pure returns(string stringifiedValue)
func (_Vm *VmCallerSession) ToString4(value [32]byte) (string, error) {
	return _Vm.Contract.ToString4(&_Vm.CallOpts, value)
}

// ToUppercase is a free data retrieval call binding the contract method 0x074ae3d7.
//
// Solidity: function toUppercase(string input) pure returns(string output)
func (_Vm *VmCaller) ToUppercase(opts *bind.CallOpts, input string) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "toUppercase", input)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToUppercase is a free data retrieval call binding the contract method 0x074ae3d7.
//
// Solidity: function toUppercase(string input) pure returns(string output)
func (_Vm *VmSession) ToUppercase(input string) (string, error) {
	return _Vm.Contract.ToUppercase(&_Vm.CallOpts, input)
}

// ToUppercase is a free data retrieval call binding the contract method 0x074ae3d7.
//
// Solidity: function toUppercase(string input) pure returns(string output)
func (_Vm *VmCallerSession) ToUppercase(input string) (string, error) {
	return _Vm.Contract.ToUppercase(&_Vm.CallOpts, input)
}

// Trim is a free data retrieval call binding the contract method 0xb2dad155.
//
// Solidity: function trim(string input) pure returns(string output)
func (_Vm *VmCaller) Trim(opts *bind.CallOpts, input string) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "trim", input)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Trim is a free data retrieval call binding the contract method 0xb2dad155.
//
// Solidity: function trim(string input) pure returns(string output)
func (_Vm *VmSession) Trim(input string) (string, error) {
	return _Vm.Contract.Trim(&_Vm.CallOpts, input)
}

// Trim is a free data retrieval call binding the contract method 0xb2dad155.
//
// Solidity: function trim(string input) pure returns(string output)
func (_Vm *VmCallerSession) Trim(input string) (string, error) {
	return _Vm.Contract.Trim(&_Vm.CallOpts, input)
}

// Accesses is a paid mutator transaction binding the contract method 0x65bc9481.
//
// Solidity: function accesses(address target) returns(bytes32[] readSlots, bytes32[] writeSlots)
func (_Vm *VmTransactor) Accesses(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "accesses", target)
}

// Accesses is a paid mutator transaction binding the contract method 0x65bc9481.
//
// Solidity: function accesses(address target) returns(bytes32[] readSlots, bytes32[] writeSlots)
func (_Vm *VmSession) Accesses(target common.Address) (*types.Transaction, error) {
	return _Vm.Contract.Accesses(&_Vm.TransactOpts, target)
}

// Accesses is a paid mutator transaction binding the contract method 0x65bc9481.
//
// Solidity: function accesses(address target) returns(bytes32[] readSlots, bytes32[] writeSlots)
func (_Vm *VmTransactorSession) Accesses(target common.Address) (*types.Transaction, error) {
	return _Vm.Contract.Accesses(&_Vm.TransactOpts, target)
}

// AllowCheatcodes is a paid mutator transaction binding the contract method 0xea060291.
//
// Solidity: function allowCheatcodes(address account) returns()
func (_Vm *VmTransactor) AllowCheatcodes(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "allowCheatcodes", account)
}

// AllowCheatcodes is a paid mutator transaction binding the contract method 0xea060291.
//
// Solidity: function allowCheatcodes(address account) returns()
func (_Vm *VmSession) AllowCheatcodes(account common.Address) (*types.Transaction, error) {
	return _Vm.Contract.AllowCheatcodes(&_Vm.TransactOpts, account)
}

// AllowCheatcodes is a paid mutator transaction binding the contract method 0xea060291.
//
// Solidity: function allowCheatcodes(address account) returns()
func (_Vm *VmTransactorSession) AllowCheatcodes(account common.Address) (*types.Transaction, error) {
	return _Vm.Contract.AllowCheatcodes(&_Vm.TransactOpts, account)
}

// BlobBaseFee is a paid mutator transaction binding the contract method 0x6d315d7e.
//
// Solidity: function blobBaseFee(uint256 newBlobBaseFee) returns()
func (_Vm *VmTransactor) BlobBaseFee(opts *bind.TransactOpts, newBlobBaseFee *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "blobBaseFee", newBlobBaseFee)
}

// BlobBaseFee is a paid mutator transaction binding the contract method 0x6d315d7e.
//
// Solidity: function blobBaseFee(uint256 newBlobBaseFee) returns()
func (_Vm *VmSession) BlobBaseFee(newBlobBaseFee *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.BlobBaseFee(&_Vm.TransactOpts, newBlobBaseFee)
}

// BlobBaseFee is a paid mutator transaction binding the contract method 0x6d315d7e.
//
// Solidity: function blobBaseFee(uint256 newBlobBaseFee) returns()
func (_Vm *VmTransactorSession) BlobBaseFee(newBlobBaseFee *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.BlobBaseFee(&_Vm.TransactOpts, newBlobBaseFee)
}

// Blobhashes is a paid mutator transaction binding the contract method 0x129de7eb.
//
// Solidity: function blobhashes(bytes32[] hashes) returns()
func (_Vm *VmTransactor) Blobhashes(opts *bind.TransactOpts, hashes [][32]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "blobhashes", hashes)
}

// Blobhashes is a paid mutator transaction binding the contract method 0x129de7eb.
//
// Solidity: function blobhashes(bytes32[] hashes) returns()
func (_Vm *VmSession) Blobhashes(hashes [][32]byte) (*types.Transaction, error) {
	return _Vm.Contract.Blobhashes(&_Vm.TransactOpts, hashes)
}

// Blobhashes is a paid mutator transaction binding the contract method 0x129de7eb.
//
// Solidity: function blobhashes(bytes32[] hashes) returns()
func (_Vm *VmTransactorSession) Blobhashes(hashes [][32]byte) (*types.Transaction, error) {
	return _Vm.Contract.Blobhashes(&_Vm.TransactOpts, hashes)
}

// Breakpoint is a paid mutator transaction binding the contract method 0xf0259e92.
//
// Solidity: function breakpoint(string char) returns()
func (_Vm *VmTransactor) Breakpoint(opts *bind.TransactOpts, char string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "breakpoint", char)
}

// Breakpoint is a paid mutator transaction binding the contract method 0xf0259e92.
//
// Solidity: function breakpoint(string char) returns()
func (_Vm *VmSession) Breakpoint(char string) (*types.Transaction, error) {
	return _Vm.Contract.Breakpoint(&_Vm.TransactOpts, char)
}

// Breakpoint is a paid mutator transaction binding the contract method 0xf0259e92.
//
// Solidity: function breakpoint(string char) returns()
func (_Vm *VmTransactorSession) Breakpoint(char string) (*types.Transaction, error) {
	return _Vm.Contract.Breakpoint(&_Vm.TransactOpts, char)
}

// Breakpoint0 is a paid mutator transaction binding the contract method 0xf7d39a8d.
//
// Solidity: function breakpoint(string char, bool value) returns()
func (_Vm *VmTransactor) Breakpoint0(opts *bind.TransactOpts, char string, value bool) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "breakpoint0", char, value)
}

// Breakpoint0 is a paid mutator transaction binding the contract method 0xf7d39a8d.
//
// Solidity: function breakpoint(string char, bool value) returns()
func (_Vm *VmSession) Breakpoint0(char string, value bool) (*types.Transaction, error) {
	return _Vm.Contract.Breakpoint0(&_Vm.TransactOpts, char, value)
}

// Breakpoint0 is a paid mutator transaction binding the contract method 0xf7d39a8d.
//
// Solidity: function breakpoint(string char, bool value) returns()
func (_Vm *VmTransactorSession) Breakpoint0(char string, value bool) (*types.Transaction, error) {
	return _Vm.Contract.Breakpoint0(&_Vm.TransactOpts, char, value)
}

// Broadcast is a paid mutator transaction binding the contract method 0xafc98040.
//
// Solidity: function broadcast() returns()
func (_Vm *VmTransactor) Broadcast(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "broadcast")
}

// Broadcast is a paid mutator transaction binding the contract method 0xafc98040.
//
// Solidity: function broadcast() returns()
func (_Vm *VmSession) Broadcast() (*types.Transaction, error) {
	return _Vm.Contract.Broadcast(&_Vm.TransactOpts)
}

// Broadcast is a paid mutator transaction binding the contract method 0xafc98040.
//
// Solidity: function broadcast() returns()
func (_Vm *VmTransactorSession) Broadcast() (*types.Transaction, error) {
	return _Vm.Contract.Broadcast(&_Vm.TransactOpts)
}

// Broadcast0 is a paid mutator transaction binding the contract method 0xe6962cdb.
//
// Solidity: function broadcast(address signer) returns()
func (_Vm *VmTransactor) Broadcast0(opts *bind.TransactOpts, signer common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "broadcast0", signer)
}

// Broadcast0 is a paid mutator transaction binding the contract method 0xe6962cdb.
//
// Solidity: function broadcast(address signer) returns()
func (_Vm *VmSession) Broadcast0(signer common.Address) (*types.Transaction, error) {
	return _Vm.Contract.Broadcast0(&_Vm.TransactOpts, signer)
}

// Broadcast0 is a paid mutator transaction binding the contract method 0xe6962cdb.
//
// Solidity: function broadcast(address signer) returns()
func (_Vm *VmTransactorSession) Broadcast0(signer common.Address) (*types.Transaction, error) {
	return _Vm.Contract.Broadcast0(&_Vm.TransactOpts, signer)
}

// Broadcast1 is a paid mutator transaction binding the contract method 0xf67a965b.
//
// Solidity: function broadcast(uint256 privateKey) returns()
func (_Vm *VmTransactor) Broadcast1(opts *bind.TransactOpts, privateKey *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "broadcast1", privateKey)
}

// Broadcast1 is a paid mutator transaction binding the contract method 0xf67a965b.
//
// Solidity: function broadcast(uint256 privateKey) returns()
func (_Vm *VmSession) Broadcast1(privateKey *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.Broadcast1(&_Vm.TransactOpts, privateKey)
}

// Broadcast1 is a paid mutator transaction binding the contract method 0xf67a965b.
//
// Solidity: function broadcast(uint256 privateKey) returns()
func (_Vm *VmTransactorSession) Broadcast1(privateKey *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.Broadcast1(&_Vm.TransactOpts, privateKey)
}

// ChainId is a paid mutator transaction binding the contract method 0x4049ddd2.
//
// Solidity: function chainId(uint256 newChainId) returns()
func (_Vm *VmTransactor) ChainId(opts *bind.TransactOpts, newChainId *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "chainId", newChainId)
}

// ChainId is a paid mutator transaction binding the contract method 0x4049ddd2.
//
// Solidity: function chainId(uint256 newChainId) returns()
func (_Vm *VmSession) ChainId(newChainId *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.ChainId(&_Vm.TransactOpts, newChainId)
}

// ChainId is a paid mutator transaction binding the contract method 0x4049ddd2.
//
// Solidity: function chainId(uint256 newChainId) returns()
func (_Vm *VmTransactorSession) ChainId(newChainId *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.ChainId(&_Vm.TransactOpts, newChainId)
}

// ClearMockedCalls is a paid mutator transaction binding the contract method 0x3fdf4e15.
//
// Solidity: function clearMockedCalls() returns()
func (_Vm *VmTransactor) ClearMockedCalls(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "clearMockedCalls")
}

// ClearMockedCalls is a paid mutator transaction binding the contract method 0x3fdf4e15.
//
// Solidity: function clearMockedCalls() returns()
func (_Vm *VmSession) ClearMockedCalls() (*types.Transaction, error) {
	return _Vm.Contract.ClearMockedCalls(&_Vm.TransactOpts)
}

// ClearMockedCalls is a paid mutator transaction binding the contract method 0x3fdf4e15.
//
// Solidity: function clearMockedCalls() returns()
func (_Vm *VmTransactorSession) ClearMockedCalls() (*types.Transaction, error) {
	return _Vm.Contract.ClearMockedCalls(&_Vm.TransactOpts)
}

// CloseFile is a paid mutator transaction binding the contract method 0x48c3241f.
//
// Solidity: function closeFile(string path) returns()
func (_Vm *VmTransactor) CloseFile(opts *bind.TransactOpts, path string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "closeFile", path)
}

// CloseFile is a paid mutator transaction binding the contract method 0x48c3241f.
//
// Solidity: function closeFile(string path) returns()
func (_Vm *VmSession) CloseFile(path string) (*types.Transaction, error) {
	return _Vm.Contract.CloseFile(&_Vm.TransactOpts, path)
}

// CloseFile is a paid mutator transaction binding the contract method 0x48c3241f.
//
// Solidity: function closeFile(string path) returns()
func (_Vm *VmTransactorSession) CloseFile(path string) (*types.Transaction, error) {
	return _Vm.Contract.CloseFile(&_Vm.TransactOpts, path)
}

// Coinbase is a paid mutator transaction binding the contract method 0xff483c54.
//
// Solidity: function coinbase(address newCoinbase) returns()
func (_Vm *VmTransactor) Coinbase(opts *bind.TransactOpts, newCoinbase common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "coinbase", newCoinbase)
}

// Coinbase is a paid mutator transaction binding the contract method 0xff483c54.
//
// Solidity: function coinbase(address newCoinbase) returns()
func (_Vm *VmSession) Coinbase(newCoinbase common.Address) (*types.Transaction, error) {
	return _Vm.Contract.Coinbase(&_Vm.TransactOpts, newCoinbase)
}

// Coinbase is a paid mutator transaction binding the contract method 0xff483c54.
//
// Solidity: function coinbase(address newCoinbase) returns()
func (_Vm *VmTransactorSession) Coinbase(newCoinbase common.Address) (*types.Transaction, error) {
	return _Vm.Contract.Coinbase(&_Vm.TransactOpts, newCoinbase)
}

// CopyFile is a paid mutator transaction binding the contract method 0xa54a87d8.
//
// Solidity: function copyFile(string from, string to) returns(uint64 copied)
func (_Vm *VmTransactor) CopyFile(opts *bind.TransactOpts, from string, to string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "copyFile", from, to)
}

// CopyFile is a paid mutator transaction binding the contract method 0xa54a87d8.
//
// Solidity: function copyFile(string from, string to) returns(uint64 copied)
func (_Vm *VmSession) CopyFile(from string, to string) (*types.Transaction, error) {
	return _Vm.Contract.CopyFile(&_Vm.TransactOpts, from, to)
}

// CopyFile is a paid mutator transaction binding the contract method 0xa54a87d8.
//
// Solidity: function copyFile(string from, string to) returns(uint64 copied)
func (_Vm *VmTransactorSession) CopyFile(from string, to string) (*types.Transaction, error) {
	return _Vm.Contract.CopyFile(&_Vm.TransactOpts, from, to)
}

// CreateDir is a paid mutator transaction binding the contract method 0x168b64d3.
//
// Solidity: function createDir(string path, bool recursive) returns()
func (_Vm *VmTransactor) CreateDir(opts *bind.TransactOpts, path string, recursive bool) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "createDir", path, recursive)
}

// CreateDir is a paid mutator transaction binding the contract method 0x168b64d3.
//
// Solidity: function createDir(string path, bool recursive) returns()
func (_Vm *VmSession) CreateDir(path string, recursive bool) (*types.Transaction, error) {
	return _Vm.Contract.CreateDir(&_Vm.TransactOpts, path, recursive)
}

// CreateDir is a paid mutator transaction binding the contract method 0x168b64d3.
//
// Solidity: function createDir(string path, bool recursive) returns()
func (_Vm *VmTransactorSession) CreateDir(path string, recursive bool) (*types.Transaction, error) {
	return _Vm.Contract.CreateDir(&_Vm.TransactOpts, path, recursive)
}

// CreateFork is a paid mutator transaction binding the contract method 0x31ba3498.
//
// Solidity: function createFork(string urlOrAlias) returns(uint256 forkId)
func (_Vm *VmTransactor) CreateFork(opts *bind.TransactOpts, urlOrAlias string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "createFork", urlOrAlias)
}

// CreateFork is a paid mutator transaction binding the contract method 0x31ba3498.
//
// Solidity: function createFork(string urlOrAlias) returns(uint256 forkId)
func (_Vm *VmSession) CreateFork(urlOrAlias string) (*types.Transaction, error) {
	return _Vm.Contract.CreateFork(&_Vm.TransactOpts, urlOrAlias)
}

// CreateFork is a paid mutator transaction binding the contract method 0x31ba3498.
//
// Solidity: function createFork(string urlOrAlias) returns(uint256 forkId)
func (_Vm *VmTransactorSession) CreateFork(urlOrAlias string) (*types.Transaction, error) {
	return _Vm.Contract.CreateFork(&_Vm.TransactOpts, urlOrAlias)
}

// CreateFork0 is a paid mutator transaction binding the contract method 0x6ba3ba2b.
//
// Solidity: function createFork(string urlOrAlias, uint256 blockNumber) returns(uint256 forkId)
func (_Vm *VmTransactor) CreateFork0(opts *bind.TransactOpts, urlOrAlias string, blockNumber *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "createFork0", urlOrAlias, blockNumber)
}

// CreateFork0 is a paid mutator transaction binding the contract method 0x6ba3ba2b.
//
// Solidity: function createFork(string urlOrAlias, uint256 blockNumber) returns(uint256 forkId)
func (_Vm *VmSession) CreateFork0(urlOrAlias string, blockNumber *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.CreateFork0(&_Vm.TransactOpts, urlOrAlias, blockNumber)
}

// CreateFork0 is a paid mutator transaction binding the contract method 0x6ba3ba2b.
//
// Solidity: function createFork(string urlOrAlias, uint256 blockNumber) returns(uint256 forkId)
func (_Vm *VmTransactorSession) CreateFork0(urlOrAlias string, blockNumber *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.CreateFork0(&_Vm.TransactOpts, urlOrAlias, blockNumber)
}

// CreateFork1 is a paid mutator transaction binding the contract method 0x7ca29682.
//
// Solidity: function createFork(string urlOrAlias, bytes32 txHash) returns(uint256 forkId)
func (_Vm *VmTransactor) CreateFork1(opts *bind.TransactOpts, urlOrAlias string, txHash [32]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "createFork1", urlOrAlias, txHash)
}

// CreateFork1 is a paid mutator transaction binding the contract method 0x7ca29682.
//
// Solidity: function createFork(string urlOrAlias, bytes32 txHash) returns(uint256 forkId)
func (_Vm *VmSession) CreateFork1(urlOrAlias string, txHash [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.CreateFork1(&_Vm.TransactOpts, urlOrAlias, txHash)
}

// CreateFork1 is a paid mutator transaction binding the contract method 0x7ca29682.
//
// Solidity: function createFork(string urlOrAlias, bytes32 txHash) returns(uint256 forkId)
func (_Vm *VmTransactorSession) CreateFork1(urlOrAlias string, txHash [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.CreateFork1(&_Vm.TransactOpts, urlOrAlias, txHash)
}

// CreateSelectFork is a paid mutator transaction binding the contract method 0x71ee464d.
//
// Solidity: function createSelectFork(string urlOrAlias, uint256 blockNumber) returns(uint256 forkId)
func (_Vm *VmTransactor) CreateSelectFork(opts *bind.TransactOpts, urlOrAlias string, blockNumber *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "createSelectFork", urlOrAlias, blockNumber)
}

// CreateSelectFork is a paid mutator transaction binding the contract method 0x71ee464d.
//
// Solidity: function createSelectFork(string urlOrAlias, uint256 blockNumber) returns(uint256 forkId)
func (_Vm *VmSession) CreateSelectFork(urlOrAlias string, blockNumber *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.CreateSelectFork(&_Vm.TransactOpts, urlOrAlias, blockNumber)
}

// CreateSelectFork is a paid mutator transaction binding the contract method 0x71ee464d.
//
// Solidity: function createSelectFork(string urlOrAlias, uint256 blockNumber) returns(uint256 forkId)
func (_Vm *VmTransactorSession) CreateSelectFork(urlOrAlias string, blockNumber *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.CreateSelectFork(&_Vm.TransactOpts, urlOrAlias, blockNumber)
}

// CreateSelectFork0 is a paid mutator transaction binding the contract method 0x84d52b7a.
//
// Solidity: function createSelectFork(string urlOrAlias, bytes32 txHash) returns(uint256 forkId)
func (_Vm *VmTransactor) CreateSelectFork0(opts *bind.TransactOpts, urlOrAlias string, txHash [32]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "createSelectFork0", urlOrAlias, txHash)
}

// CreateSelectFork0 is a paid mutator transaction binding the contract method 0x84d52b7a.
//
// Solidity: function createSelectFork(string urlOrAlias, bytes32 txHash) returns(uint256 forkId)
func (_Vm *VmSession) CreateSelectFork0(urlOrAlias string, txHash [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.CreateSelectFork0(&_Vm.TransactOpts, urlOrAlias, txHash)
}

// CreateSelectFork0 is a paid mutator transaction binding the contract method 0x84d52b7a.
//
// Solidity: function createSelectFork(string urlOrAlias, bytes32 txHash) returns(uint256 forkId)
func (_Vm *VmTransactorSession) CreateSelectFork0(urlOrAlias string, txHash [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.CreateSelectFork0(&_Vm.TransactOpts, urlOrAlias, txHash)
}

// CreateSelectFork1 is a paid mutator transaction binding the contract method 0x98680034.
//
// Solidity: function createSelectFork(string urlOrAlias) returns(uint256 forkId)
func (_Vm *VmTransactor) CreateSelectFork1(opts *bind.TransactOpts, urlOrAlias string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "createSelectFork1", urlOrAlias)
}

// CreateSelectFork1 is a paid mutator transaction binding the contract method 0x98680034.
//
// Solidity: function createSelectFork(string urlOrAlias) returns(uint256 forkId)
func (_Vm *VmSession) CreateSelectFork1(urlOrAlias string) (*types.Transaction, error) {
	return _Vm.Contract.CreateSelectFork1(&_Vm.TransactOpts, urlOrAlias)
}

// CreateSelectFork1 is a paid mutator transaction binding the contract method 0x98680034.
//
// Solidity: function createSelectFork(string urlOrAlias) returns(uint256 forkId)
func (_Vm *VmTransactorSession) CreateSelectFork1(urlOrAlias string) (*types.Transaction, error) {
	return _Vm.Contract.CreateSelectFork1(&_Vm.TransactOpts, urlOrAlias)
}

// CreateWallet is a paid mutator transaction binding the contract method 0x7404f1d2.
//
// Solidity: function createWallet(string walletLabel) returns((address,uint256,uint256,uint256) wallet)
func (_Vm *VmTransactor) CreateWallet(opts *bind.TransactOpts, walletLabel string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "createWallet", walletLabel)
}

// CreateWallet is a paid mutator transaction binding the contract method 0x7404f1d2.
//
// Solidity: function createWallet(string walletLabel) returns((address,uint256,uint256,uint256) wallet)
func (_Vm *VmSession) CreateWallet(walletLabel string) (*types.Transaction, error) {
	return _Vm.Contract.CreateWallet(&_Vm.TransactOpts, walletLabel)
}

// CreateWallet is a paid mutator transaction binding the contract method 0x7404f1d2.
//
// Solidity: function createWallet(string walletLabel) returns((address,uint256,uint256,uint256) wallet)
func (_Vm *VmTransactorSession) CreateWallet(walletLabel string) (*types.Transaction, error) {
	return _Vm.Contract.CreateWallet(&_Vm.TransactOpts, walletLabel)
}

// CreateWallet0 is a paid mutator transaction binding the contract method 0x7a675bb6.
//
// Solidity: function createWallet(uint256 privateKey) returns((address,uint256,uint256,uint256) wallet)
func (_Vm *VmTransactor) CreateWallet0(opts *bind.TransactOpts, privateKey *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "createWallet0", privateKey)
}

// CreateWallet0 is a paid mutator transaction binding the contract method 0x7a675bb6.
//
// Solidity: function createWallet(uint256 privateKey) returns((address,uint256,uint256,uint256) wallet)
func (_Vm *VmSession) CreateWallet0(privateKey *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.CreateWallet0(&_Vm.TransactOpts, privateKey)
}

// CreateWallet0 is a paid mutator transaction binding the contract method 0x7a675bb6.
//
// Solidity: function createWallet(uint256 privateKey) returns((address,uint256,uint256,uint256) wallet)
func (_Vm *VmTransactorSession) CreateWallet0(privateKey *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.CreateWallet0(&_Vm.TransactOpts, privateKey)
}

// CreateWallet1 is a paid mutator transaction binding the contract method 0xed7c5462.
//
// Solidity: function createWallet(uint256 privateKey, string walletLabel) returns((address,uint256,uint256,uint256) wallet)
func (_Vm *VmTransactor) CreateWallet1(opts *bind.TransactOpts, privateKey *big.Int, walletLabel string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "createWallet1", privateKey, walletLabel)
}

// CreateWallet1 is a paid mutator transaction binding the contract method 0xed7c5462.
//
// Solidity: function createWallet(uint256 privateKey, string walletLabel) returns((address,uint256,uint256,uint256) wallet)
func (_Vm *VmSession) CreateWallet1(privateKey *big.Int, walletLabel string) (*types.Transaction, error) {
	return _Vm.Contract.CreateWallet1(&_Vm.TransactOpts, privateKey, walletLabel)
}

// CreateWallet1 is a paid mutator transaction binding the contract method 0xed7c5462.
//
// Solidity: function createWallet(uint256 privateKey, string walletLabel) returns((address,uint256,uint256,uint256) wallet)
func (_Vm *VmTransactorSession) CreateWallet1(privateKey *big.Int, walletLabel string) (*types.Transaction, error) {
	return _Vm.Contract.CreateWallet1(&_Vm.TransactOpts, privateKey, walletLabel)
}

// Deal is a paid mutator transaction binding the contract method 0xc88a5e6d.
//
// Solidity: function deal(address account, uint256 newBalance) returns()
func (_Vm *VmTransactor) Deal(opts *bind.TransactOpts, account common.Address, newBalance *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "deal", account, newBalance)
}

// Deal is a paid mutator transaction binding the contract method 0xc88a5e6d.
//
// Solidity: function deal(address account, uint256 newBalance) returns()
func (_Vm *VmSession) Deal(account common.Address, newBalance *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.Deal(&_Vm.TransactOpts, account, newBalance)
}

// Deal is a paid mutator transaction binding the contract method 0xc88a5e6d.
//
// Solidity: function deal(address account, uint256 newBalance) returns()
func (_Vm *VmTransactorSession) Deal(account common.Address, newBalance *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.Deal(&_Vm.TransactOpts, account, newBalance)
}

// DeleteSnapshot is a paid mutator transaction binding the contract method 0xa6368557.
//
// Solidity: function deleteSnapshot(uint256 snapshotId) returns(bool success)
func (_Vm *VmTransactor) DeleteSnapshot(opts *bind.TransactOpts, snapshotId *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "deleteSnapshot", snapshotId)
}

// DeleteSnapshot is a paid mutator transaction binding the contract method 0xa6368557.
//
// Solidity: function deleteSnapshot(uint256 snapshotId) returns(bool success)
func (_Vm *VmSession) DeleteSnapshot(snapshotId *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.DeleteSnapshot(&_Vm.TransactOpts, snapshotId)
}

// DeleteSnapshot is a paid mutator transaction binding the contract method 0xa6368557.
//
// Solidity: function deleteSnapshot(uint256 snapshotId) returns(bool success)
func (_Vm *VmTransactorSession) DeleteSnapshot(snapshotId *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.DeleteSnapshot(&_Vm.TransactOpts, snapshotId)
}

// DeleteSnapshots is a paid mutator transaction binding the contract method 0x421ae469.
//
// Solidity: function deleteSnapshots() returns()
func (_Vm *VmTransactor) DeleteSnapshots(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "deleteSnapshots")
}

// DeleteSnapshots is a paid mutator transaction binding the contract method 0x421ae469.
//
// Solidity: function deleteSnapshots() returns()
func (_Vm *VmSession) DeleteSnapshots() (*types.Transaction, error) {
	return _Vm.Contract.DeleteSnapshots(&_Vm.TransactOpts)
}

// DeleteSnapshots is a paid mutator transaction binding the contract method 0x421ae469.
//
// Solidity: function deleteSnapshots() returns()
func (_Vm *VmTransactorSession) DeleteSnapshots() (*types.Transaction, error) {
	return _Vm.Contract.DeleteSnapshots(&_Vm.TransactOpts)
}

// Difficulty is a paid mutator transaction binding the contract method 0x46cc92d9.
//
// Solidity: function difficulty(uint256 newDifficulty) returns()
func (_Vm *VmTransactor) Difficulty(opts *bind.TransactOpts, newDifficulty *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "difficulty", newDifficulty)
}

// Difficulty is a paid mutator transaction binding the contract method 0x46cc92d9.
//
// Solidity: function difficulty(uint256 newDifficulty) returns()
func (_Vm *VmSession) Difficulty(newDifficulty *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.Difficulty(&_Vm.TransactOpts, newDifficulty)
}

// Difficulty is a paid mutator transaction binding the contract method 0x46cc92d9.
//
// Solidity: function difficulty(uint256 newDifficulty) returns()
func (_Vm *VmTransactorSession) Difficulty(newDifficulty *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.Difficulty(&_Vm.TransactOpts, newDifficulty)
}

// DumpState is a paid mutator transaction binding the contract method 0x709ecd3f.
//
// Solidity: function dumpState(string pathToStateJson) returns()
func (_Vm *VmTransactor) DumpState(opts *bind.TransactOpts, pathToStateJson string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "dumpState", pathToStateJson)
}

// DumpState is a paid mutator transaction binding the contract method 0x709ecd3f.
//
// Solidity: function dumpState(string pathToStateJson) returns()
func (_Vm *VmSession) DumpState(pathToStateJson string) (*types.Transaction, error) {
	return _Vm.Contract.DumpState(&_Vm.TransactOpts, pathToStateJson)
}

// DumpState is a paid mutator transaction binding the contract method 0x709ecd3f.
//
// Solidity: function dumpState(string pathToStateJson) returns()
func (_Vm *VmTransactorSession) DumpState(pathToStateJson string) (*types.Transaction, error) {
	return _Vm.Contract.DumpState(&_Vm.TransactOpts, pathToStateJson)
}

// Etch is a paid mutator transaction binding the contract method 0xb4d6c782.
//
// Solidity: function etch(address target, bytes newRuntimeBytecode) returns()
func (_Vm *VmTransactor) Etch(opts *bind.TransactOpts, target common.Address, newRuntimeBytecode []byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "etch", target, newRuntimeBytecode)
}

// Etch is a paid mutator transaction binding the contract method 0xb4d6c782.
//
// Solidity: function etch(address target, bytes newRuntimeBytecode) returns()
func (_Vm *VmSession) Etch(target common.Address, newRuntimeBytecode []byte) (*types.Transaction, error) {
	return _Vm.Contract.Etch(&_Vm.TransactOpts, target, newRuntimeBytecode)
}

// Etch is a paid mutator transaction binding the contract method 0xb4d6c782.
//
// Solidity: function etch(address target, bytes newRuntimeBytecode) returns()
func (_Vm *VmTransactorSession) Etch(target common.Address, newRuntimeBytecode []byte) (*types.Transaction, error) {
	return _Vm.Contract.Etch(&_Vm.TransactOpts, target, newRuntimeBytecode)
}

// EthGetLogs is a paid mutator transaction binding the contract method 0x35e1349b.
//
// Solidity: function eth_getLogs(uint256 fromBlock, uint256 toBlock, address target, bytes32[] topics) returns((address,bytes32[],bytes,bytes32,uint64,bytes32,uint64,uint256,bool)[] logs)
func (_Vm *VmTransactor) EthGetLogs(opts *bind.TransactOpts, fromBlock *big.Int, toBlock *big.Int, target common.Address, topics [][32]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "eth_getLogs", fromBlock, toBlock, target, topics)
}

// EthGetLogs is a paid mutator transaction binding the contract method 0x35e1349b.
//
// Solidity: function eth_getLogs(uint256 fromBlock, uint256 toBlock, address target, bytes32[] topics) returns((address,bytes32[],bytes,bytes32,uint64,bytes32,uint64,uint256,bool)[] logs)
func (_Vm *VmSession) EthGetLogs(fromBlock *big.Int, toBlock *big.Int, target common.Address, topics [][32]byte) (*types.Transaction, error) {
	return _Vm.Contract.EthGetLogs(&_Vm.TransactOpts, fromBlock, toBlock, target, topics)
}

// EthGetLogs is a paid mutator transaction binding the contract method 0x35e1349b.
//
// Solidity: function eth_getLogs(uint256 fromBlock, uint256 toBlock, address target, bytes32[] topics) returns((address,bytes32[],bytes,bytes32,uint64,bytes32,uint64,uint256,bool)[] logs)
func (_Vm *VmTransactorSession) EthGetLogs(fromBlock *big.Int, toBlock *big.Int, target common.Address, topics [][32]byte) (*types.Transaction, error) {
	return _Vm.Contract.EthGetLogs(&_Vm.TransactOpts, fromBlock, toBlock, target, topics)
}

// Exists is a paid mutator transaction binding the contract method 0x261a323e.
//
// Solidity: function exists(string path) returns(bool result)
func (_Vm *VmTransactor) Exists(opts *bind.TransactOpts, path string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "exists", path)
}

// Exists is a paid mutator transaction binding the contract method 0x261a323e.
//
// Solidity: function exists(string path) returns(bool result)
func (_Vm *VmSession) Exists(path string) (*types.Transaction, error) {
	return _Vm.Contract.Exists(&_Vm.TransactOpts, path)
}

// Exists is a paid mutator transaction binding the contract method 0x261a323e.
//
// Solidity: function exists(string path) returns(bool result)
func (_Vm *VmTransactorSession) Exists(path string) (*types.Transaction, error) {
	return _Vm.Contract.Exists(&_Vm.TransactOpts, path)
}

// ExpectCall is a paid mutator transaction binding the contract method 0x23361207.
//
// Solidity: function expectCall(address callee, uint256 msgValue, uint64 gas, bytes data) returns()
func (_Vm *VmTransactor) ExpectCall(opts *bind.TransactOpts, callee common.Address, msgValue *big.Int, gas uint64, data []byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectCall", callee, msgValue, gas, data)
}

// ExpectCall is a paid mutator transaction binding the contract method 0x23361207.
//
// Solidity: function expectCall(address callee, uint256 msgValue, uint64 gas, bytes data) returns()
func (_Vm *VmSession) ExpectCall(callee common.Address, msgValue *big.Int, gas uint64, data []byte) (*types.Transaction, error) {
	return _Vm.Contract.ExpectCall(&_Vm.TransactOpts, callee, msgValue, gas, data)
}

// ExpectCall is a paid mutator transaction binding the contract method 0x23361207.
//
// Solidity: function expectCall(address callee, uint256 msgValue, uint64 gas, bytes data) returns()
func (_Vm *VmTransactorSession) ExpectCall(callee common.Address, msgValue *big.Int, gas uint64, data []byte) (*types.Transaction, error) {
	return _Vm.Contract.ExpectCall(&_Vm.TransactOpts, callee, msgValue, gas, data)
}

// ExpectCall0 is a paid mutator transaction binding the contract method 0x65b7b7cc.
//
// Solidity: function expectCall(address callee, uint256 msgValue, uint64 gas, bytes data, uint64 count) returns()
func (_Vm *VmTransactor) ExpectCall0(opts *bind.TransactOpts, callee common.Address, msgValue *big.Int, gas uint64, data []byte, count uint64) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectCall0", callee, msgValue, gas, data, count)
}

// ExpectCall0 is a paid mutator transaction binding the contract method 0x65b7b7cc.
//
// Solidity: function expectCall(address callee, uint256 msgValue, uint64 gas, bytes data, uint64 count) returns()
func (_Vm *VmSession) ExpectCall0(callee common.Address, msgValue *big.Int, gas uint64, data []byte, count uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectCall0(&_Vm.TransactOpts, callee, msgValue, gas, data, count)
}

// ExpectCall0 is a paid mutator transaction binding the contract method 0x65b7b7cc.
//
// Solidity: function expectCall(address callee, uint256 msgValue, uint64 gas, bytes data, uint64 count) returns()
func (_Vm *VmTransactorSession) ExpectCall0(callee common.Address, msgValue *big.Int, gas uint64, data []byte, count uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectCall0(&_Vm.TransactOpts, callee, msgValue, gas, data, count)
}

// ExpectCall1 is a paid mutator transaction binding the contract method 0xa2b1a1ae.
//
// Solidity: function expectCall(address callee, uint256 msgValue, bytes data, uint64 count) returns()
func (_Vm *VmTransactor) ExpectCall1(opts *bind.TransactOpts, callee common.Address, msgValue *big.Int, data []byte, count uint64) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectCall1", callee, msgValue, data, count)
}

// ExpectCall1 is a paid mutator transaction binding the contract method 0xa2b1a1ae.
//
// Solidity: function expectCall(address callee, uint256 msgValue, bytes data, uint64 count) returns()
func (_Vm *VmSession) ExpectCall1(callee common.Address, msgValue *big.Int, data []byte, count uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectCall1(&_Vm.TransactOpts, callee, msgValue, data, count)
}

// ExpectCall1 is a paid mutator transaction binding the contract method 0xa2b1a1ae.
//
// Solidity: function expectCall(address callee, uint256 msgValue, bytes data, uint64 count) returns()
func (_Vm *VmTransactorSession) ExpectCall1(callee common.Address, msgValue *big.Int, data []byte, count uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectCall1(&_Vm.TransactOpts, callee, msgValue, data, count)
}

// ExpectCall2 is a paid mutator transaction binding the contract method 0xbd6af434.
//
// Solidity: function expectCall(address callee, bytes data) returns()
func (_Vm *VmTransactor) ExpectCall2(opts *bind.TransactOpts, callee common.Address, data []byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectCall2", callee, data)
}

// ExpectCall2 is a paid mutator transaction binding the contract method 0xbd6af434.
//
// Solidity: function expectCall(address callee, bytes data) returns()
func (_Vm *VmSession) ExpectCall2(callee common.Address, data []byte) (*types.Transaction, error) {
	return _Vm.Contract.ExpectCall2(&_Vm.TransactOpts, callee, data)
}

// ExpectCall2 is a paid mutator transaction binding the contract method 0xbd6af434.
//
// Solidity: function expectCall(address callee, bytes data) returns()
func (_Vm *VmTransactorSession) ExpectCall2(callee common.Address, data []byte) (*types.Transaction, error) {
	return _Vm.Contract.ExpectCall2(&_Vm.TransactOpts, callee, data)
}

// ExpectCall3 is a paid mutator transaction binding the contract method 0xc1adbbff.
//
// Solidity: function expectCall(address callee, bytes data, uint64 count) returns()
func (_Vm *VmTransactor) ExpectCall3(opts *bind.TransactOpts, callee common.Address, data []byte, count uint64) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectCall3", callee, data, count)
}

// ExpectCall3 is a paid mutator transaction binding the contract method 0xc1adbbff.
//
// Solidity: function expectCall(address callee, bytes data, uint64 count) returns()
func (_Vm *VmSession) ExpectCall3(callee common.Address, data []byte, count uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectCall3(&_Vm.TransactOpts, callee, data, count)
}

// ExpectCall3 is a paid mutator transaction binding the contract method 0xc1adbbff.
//
// Solidity: function expectCall(address callee, bytes data, uint64 count) returns()
func (_Vm *VmTransactorSession) ExpectCall3(callee common.Address, data []byte, count uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectCall3(&_Vm.TransactOpts, callee, data, count)
}

// ExpectCall4 is a paid mutator transaction binding the contract method 0xf30c7ba3.
//
// Solidity: function expectCall(address callee, uint256 msgValue, bytes data) returns()
func (_Vm *VmTransactor) ExpectCall4(opts *bind.TransactOpts, callee common.Address, msgValue *big.Int, data []byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectCall4", callee, msgValue, data)
}

// ExpectCall4 is a paid mutator transaction binding the contract method 0xf30c7ba3.
//
// Solidity: function expectCall(address callee, uint256 msgValue, bytes data) returns()
func (_Vm *VmSession) ExpectCall4(callee common.Address, msgValue *big.Int, data []byte) (*types.Transaction, error) {
	return _Vm.Contract.ExpectCall4(&_Vm.TransactOpts, callee, msgValue, data)
}

// ExpectCall4 is a paid mutator transaction binding the contract method 0xf30c7ba3.
//
// Solidity: function expectCall(address callee, uint256 msgValue, bytes data) returns()
func (_Vm *VmTransactorSession) ExpectCall4(callee common.Address, msgValue *big.Int, data []byte) (*types.Transaction, error) {
	return _Vm.Contract.ExpectCall4(&_Vm.TransactOpts, callee, msgValue, data)
}

// ExpectCallMinGas is a paid mutator transaction binding the contract method 0x08e4e116.
//
// Solidity: function expectCallMinGas(address callee, uint256 msgValue, uint64 minGas, bytes data) returns()
func (_Vm *VmTransactor) ExpectCallMinGas(opts *bind.TransactOpts, callee common.Address, msgValue *big.Int, minGas uint64, data []byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectCallMinGas", callee, msgValue, minGas, data)
}

// ExpectCallMinGas is a paid mutator transaction binding the contract method 0x08e4e116.
//
// Solidity: function expectCallMinGas(address callee, uint256 msgValue, uint64 minGas, bytes data) returns()
func (_Vm *VmSession) ExpectCallMinGas(callee common.Address, msgValue *big.Int, minGas uint64, data []byte) (*types.Transaction, error) {
	return _Vm.Contract.ExpectCallMinGas(&_Vm.TransactOpts, callee, msgValue, minGas, data)
}

// ExpectCallMinGas is a paid mutator transaction binding the contract method 0x08e4e116.
//
// Solidity: function expectCallMinGas(address callee, uint256 msgValue, uint64 minGas, bytes data) returns()
func (_Vm *VmTransactorSession) ExpectCallMinGas(callee common.Address, msgValue *big.Int, minGas uint64, data []byte) (*types.Transaction, error) {
	return _Vm.Contract.ExpectCallMinGas(&_Vm.TransactOpts, callee, msgValue, minGas, data)
}

// ExpectCallMinGas0 is a paid mutator transaction binding the contract method 0xe13a1834.
//
// Solidity: function expectCallMinGas(address callee, uint256 msgValue, uint64 minGas, bytes data, uint64 count) returns()
func (_Vm *VmTransactor) ExpectCallMinGas0(opts *bind.TransactOpts, callee common.Address, msgValue *big.Int, minGas uint64, data []byte, count uint64) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectCallMinGas0", callee, msgValue, minGas, data, count)
}

// ExpectCallMinGas0 is a paid mutator transaction binding the contract method 0xe13a1834.
//
// Solidity: function expectCallMinGas(address callee, uint256 msgValue, uint64 minGas, bytes data, uint64 count) returns()
func (_Vm *VmSession) ExpectCallMinGas0(callee common.Address, msgValue *big.Int, minGas uint64, data []byte, count uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectCallMinGas0(&_Vm.TransactOpts, callee, msgValue, minGas, data, count)
}

// ExpectCallMinGas0 is a paid mutator transaction binding the contract method 0xe13a1834.
//
// Solidity: function expectCallMinGas(address callee, uint256 msgValue, uint64 minGas, bytes data, uint64 count) returns()
func (_Vm *VmTransactorSession) ExpectCallMinGas0(callee common.Address, msgValue *big.Int, minGas uint64, data []byte, count uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectCallMinGas0(&_Vm.TransactOpts, callee, msgValue, minGas, data, count)
}

// ExpectEmit is a paid mutator transaction binding the contract method 0x440ed10d.
//
// Solidity: function expectEmit() returns()
func (_Vm *VmTransactor) ExpectEmit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectEmit")
}

// ExpectEmit is a paid mutator transaction binding the contract method 0x440ed10d.
//
// Solidity: function expectEmit() returns()
func (_Vm *VmSession) ExpectEmit() (*types.Transaction, error) {
	return _Vm.Contract.ExpectEmit(&_Vm.TransactOpts)
}

// ExpectEmit is a paid mutator transaction binding the contract method 0x440ed10d.
//
// Solidity: function expectEmit() returns()
func (_Vm *VmTransactorSession) ExpectEmit() (*types.Transaction, error) {
	return _Vm.Contract.ExpectEmit(&_Vm.TransactOpts)
}

// ExpectEmit0 is a paid mutator transaction binding the contract method 0x491cc7c2.
//
// Solidity: function expectEmit(bool checkTopic1, bool checkTopic2, bool checkTopic3, bool checkData) returns()
func (_Vm *VmTransactor) ExpectEmit0(opts *bind.TransactOpts, checkTopic1 bool, checkTopic2 bool, checkTopic3 bool, checkData bool) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectEmit0", checkTopic1, checkTopic2, checkTopic3, checkData)
}

// ExpectEmit0 is a paid mutator transaction binding the contract method 0x491cc7c2.
//
// Solidity: function expectEmit(bool checkTopic1, bool checkTopic2, bool checkTopic3, bool checkData) returns()
func (_Vm *VmSession) ExpectEmit0(checkTopic1 bool, checkTopic2 bool, checkTopic3 bool, checkData bool) (*types.Transaction, error) {
	return _Vm.Contract.ExpectEmit0(&_Vm.TransactOpts, checkTopic1, checkTopic2, checkTopic3, checkData)
}

// ExpectEmit0 is a paid mutator transaction binding the contract method 0x491cc7c2.
//
// Solidity: function expectEmit(bool checkTopic1, bool checkTopic2, bool checkTopic3, bool checkData) returns()
func (_Vm *VmTransactorSession) ExpectEmit0(checkTopic1 bool, checkTopic2 bool, checkTopic3 bool, checkData bool) (*types.Transaction, error) {
	return _Vm.Contract.ExpectEmit0(&_Vm.TransactOpts, checkTopic1, checkTopic2, checkTopic3, checkData)
}

// ExpectEmit1 is a paid mutator transaction binding the contract method 0x81bad6f3.
//
// Solidity: function expectEmit(bool checkTopic1, bool checkTopic2, bool checkTopic3, bool checkData, address emitter) returns()
func (_Vm *VmTransactor) ExpectEmit1(opts *bind.TransactOpts, checkTopic1 bool, checkTopic2 bool, checkTopic3 bool, checkData bool, emitter common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectEmit1", checkTopic1, checkTopic2, checkTopic3, checkData, emitter)
}

// ExpectEmit1 is a paid mutator transaction binding the contract method 0x81bad6f3.
//
// Solidity: function expectEmit(bool checkTopic1, bool checkTopic2, bool checkTopic3, bool checkData, address emitter) returns()
func (_Vm *VmSession) ExpectEmit1(checkTopic1 bool, checkTopic2 bool, checkTopic3 bool, checkData bool, emitter common.Address) (*types.Transaction, error) {
	return _Vm.Contract.ExpectEmit1(&_Vm.TransactOpts, checkTopic1, checkTopic2, checkTopic3, checkData, emitter)
}

// ExpectEmit1 is a paid mutator transaction binding the contract method 0x81bad6f3.
//
// Solidity: function expectEmit(bool checkTopic1, bool checkTopic2, bool checkTopic3, bool checkData, address emitter) returns()
func (_Vm *VmTransactorSession) ExpectEmit1(checkTopic1 bool, checkTopic2 bool, checkTopic3 bool, checkData bool, emitter common.Address) (*types.Transaction, error) {
	return _Vm.Contract.ExpectEmit1(&_Vm.TransactOpts, checkTopic1, checkTopic2, checkTopic3, checkData, emitter)
}

// ExpectEmit2 is a paid mutator transaction binding the contract method 0x86b9620d.
//
// Solidity: function expectEmit(address emitter) returns()
func (_Vm *VmTransactor) ExpectEmit2(opts *bind.TransactOpts, emitter common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectEmit2", emitter)
}

// ExpectEmit2 is a paid mutator transaction binding the contract method 0x86b9620d.
//
// Solidity: function expectEmit(address emitter) returns()
func (_Vm *VmSession) ExpectEmit2(emitter common.Address) (*types.Transaction, error) {
	return _Vm.Contract.ExpectEmit2(&_Vm.TransactOpts, emitter)
}

// ExpectEmit2 is a paid mutator transaction binding the contract method 0x86b9620d.
//
// Solidity: function expectEmit(address emitter) returns()
func (_Vm *VmTransactorSession) ExpectEmit2(emitter common.Address) (*types.Transaction, error) {
	return _Vm.Contract.ExpectEmit2(&_Vm.TransactOpts, emitter)
}

// ExpectRevert is a paid mutator transaction binding the contract method 0xc31eb0e0.
//
// Solidity: function expectRevert(bytes4 revertData) returns()
func (_Vm *VmTransactor) ExpectRevert(opts *bind.TransactOpts, revertData [4]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectRevert", revertData)
}

// ExpectRevert is a paid mutator transaction binding the contract method 0xc31eb0e0.
//
// Solidity: function expectRevert(bytes4 revertData) returns()
func (_Vm *VmSession) ExpectRevert(revertData [4]byte) (*types.Transaction, error) {
	return _Vm.Contract.ExpectRevert(&_Vm.TransactOpts, revertData)
}

// ExpectRevert is a paid mutator transaction binding the contract method 0xc31eb0e0.
//
// Solidity: function expectRevert(bytes4 revertData) returns()
func (_Vm *VmTransactorSession) ExpectRevert(revertData [4]byte) (*types.Transaction, error) {
	return _Vm.Contract.ExpectRevert(&_Vm.TransactOpts, revertData)
}

// ExpectRevert0 is a paid mutator transaction binding the contract method 0xf28dceb3.
//
// Solidity: function expectRevert(bytes revertData) returns()
func (_Vm *VmTransactor) ExpectRevert0(opts *bind.TransactOpts, revertData []byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectRevert0", revertData)
}

// ExpectRevert0 is a paid mutator transaction binding the contract method 0xf28dceb3.
//
// Solidity: function expectRevert(bytes revertData) returns()
func (_Vm *VmSession) ExpectRevert0(revertData []byte) (*types.Transaction, error) {
	return _Vm.Contract.ExpectRevert0(&_Vm.TransactOpts, revertData)
}

// ExpectRevert0 is a paid mutator transaction binding the contract method 0xf28dceb3.
//
// Solidity: function expectRevert(bytes revertData) returns()
func (_Vm *VmTransactorSession) ExpectRevert0(revertData []byte) (*types.Transaction, error) {
	return _Vm.Contract.ExpectRevert0(&_Vm.TransactOpts, revertData)
}

// ExpectRevert1 is a paid mutator transaction binding the contract method 0xf4844814.
//
// Solidity: function expectRevert() returns()
func (_Vm *VmTransactor) ExpectRevert1(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectRevert1")
}

// ExpectRevert1 is a paid mutator transaction binding the contract method 0xf4844814.
//
// Solidity: function expectRevert() returns()
func (_Vm *VmSession) ExpectRevert1() (*types.Transaction, error) {
	return _Vm.Contract.ExpectRevert1(&_Vm.TransactOpts)
}

// ExpectRevert1 is a paid mutator transaction binding the contract method 0xf4844814.
//
// Solidity: function expectRevert() returns()
func (_Vm *VmTransactorSession) ExpectRevert1() (*types.Transaction, error) {
	return _Vm.Contract.ExpectRevert1(&_Vm.TransactOpts)
}

// ExpectSafeMemory is a paid mutator transaction binding the contract method 0x6d016688.
//
// Solidity: function expectSafeMemory(uint64 min, uint64 max) returns()
func (_Vm *VmTransactor) ExpectSafeMemory(opts *bind.TransactOpts, min uint64, max uint64) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectSafeMemory", min, max)
}

// ExpectSafeMemory is a paid mutator transaction binding the contract method 0x6d016688.
//
// Solidity: function expectSafeMemory(uint64 min, uint64 max) returns()
func (_Vm *VmSession) ExpectSafeMemory(min uint64, max uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectSafeMemory(&_Vm.TransactOpts, min, max)
}

// ExpectSafeMemory is a paid mutator transaction binding the contract method 0x6d016688.
//
// Solidity: function expectSafeMemory(uint64 min, uint64 max) returns()
func (_Vm *VmTransactorSession) ExpectSafeMemory(min uint64, max uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectSafeMemory(&_Vm.TransactOpts, min, max)
}

// ExpectSafeMemoryCall is a paid mutator transaction binding the contract method 0x05838bf4.
//
// Solidity: function expectSafeMemoryCall(uint64 min, uint64 max) returns()
func (_Vm *VmTransactor) ExpectSafeMemoryCall(opts *bind.TransactOpts, min uint64, max uint64) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectSafeMemoryCall", min, max)
}

// ExpectSafeMemoryCall is a paid mutator transaction binding the contract method 0x05838bf4.
//
// Solidity: function expectSafeMemoryCall(uint64 min, uint64 max) returns()
func (_Vm *VmSession) ExpectSafeMemoryCall(min uint64, max uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectSafeMemoryCall(&_Vm.TransactOpts, min, max)
}

// ExpectSafeMemoryCall is a paid mutator transaction binding the contract method 0x05838bf4.
//
// Solidity: function expectSafeMemoryCall(uint64 min, uint64 max) returns()
func (_Vm *VmTransactorSession) ExpectSafeMemoryCall(min uint64, max uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectSafeMemoryCall(&_Vm.TransactOpts, min, max)
}

// Fee is a paid mutator transaction binding the contract method 0x39b37ab0.
//
// Solidity: function fee(uint256 newBasefee) returns()
func (_Vm *VmTransactor) Fee(opts *bind.TransactOpts, newBasefee *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "fee", newBasefee)
}

// Fee is a paid mutator transaction binding the contract method 0x39b37ab0.
//
// Solidity: function fee(uint256 newBasefee) returns()
func (_Vm *VmSession) Fee(newBasefee *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.Fee(&_Vm.TransactOpts, newBasefee)
}

// Fee is a paid mutator transaction binding the contract method 0x39b37ab0.
//
// Solidity: function fee(uint256 newBasefee) returns()
func (_Vm *VmTransactorSession) Fee(newBasefee *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.Fee(&_Vm.TransactOpts, newBasefee)
}

// Ffi is a paid mutator transaction binding the contract method 0x89160467.
//
// Solidity: function ffi(string[] commandInput) returns(bytes result)
func (_Vm *VmTransactor) Ffi(opts *bind.TransactOpts, commandInput []string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "ffi", commandInput)
}

// Ffi is a paid mutator transaction binding the contract method 0x89160467.
//
// Solidity: function ffi(string[] commandInput) returns(bytes result)
func (_Vm *VmSession) Ffi(commandInput []string) (*types.Transaction, error) {
	return _Vm.Contract.Ffi(&_Vm.TransactOpts, commandInput)
}

// Ffi is a paid mutator transaction binding the contract method 0x89160467.
//
// Solidity: function ffi(string[] commandInput) returns(bytes result)
func (_Vm *VmTransactorSession) Ffi(commandInput []string) (*types.Transaction, error) {
	return _Vm.Contract.Ffi(&_Vm.TransactOpts, commandInput)
}

// GetMappingKeyAndParentOf is a paid mutator transaction binding the contract method 0x876e24e6.
//
// Solidity: function getMappingKeyAndParentOf(address target, bytes32 elementSlot) returns(bool found, bytes32 key, bytes32 parent)
func (_Vm *VmTransactor) GetMappingKeyAndParentOf(opts *bind.TransactOpts, target common.Address, elementSlot [32]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "getMappingKeyAndParentOf", target, elementSlot)
}

// GetMappingKeyAndParentOf is a paid mutator transaction binding the contract method 0x876e24e6.
//
// Solidity: function getMappingKeyAndParentOf(address target, bytes32 elementSlot) returns(bool found, bytes32 key, bytes32 parent)
func (_Vm *VmSession) GetMappingKeyAndParentOf(target common.Address, elementSlot [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.GetMappingKeyAndParentOf(&_Vm.TransactOpts, target, elementSlot)
}

// GetMappingKeyAndParentOf is a paid mutator transaction binding the contract method 0x876e24e6.
//
// Solidity: function getMappingKeyAndParentOf(address target, bytes32 elementSlot) returns(bool found, bytes32 key, bytes32 parent)
func (_Vm *VmTransactorSession) GetMappingKeyAndParentOf(target common.Address, elementSlot [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.GetMappingKeyAndParentOf(&_Vm.TransactOpts, target, elementSlot)
}

// GetMappingLength is a paid mutator transaction binding the contract method 0x2f2fd63f.
//
// Solidity: function getMappingLength(address target, bytes32 mappingSlot) returns(uint256 length)
func (_Vm *VmTransactor) GetMappingLength(opts *bind.TransactOpts, target common.Address, mappingSlot [32]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "getMappingLength", target, mappingSlot)
}

// GetMappingLength is a paid mutator transaction binding the contract method 0x2f2fd63f.
//
// Solidity: function getMappingLength(address target, bytes32 mappingSlot) returns(uint256 length)
func (_Vm *VmSession) GetMappingLength(target common.Address, mappingSlot [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.GetMappingLength(&_Vm.TransactOpts, target, mappingSlot)
}

// GetMappingLength is a paid mutator transaction binding the contract method 0x2f2fd63f.
//
// Solidity: function getMappingLength(address target, bytes32 mappingSlot) returns(uint256 length)
func (_Vm *VmTransactorSession) GetMappingLength(target common.Address, mappingSlot [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.GetMappingLength(&_Vm.TransactOpts, target, mappingSlot)
}

// GetMappingSlotAt is a paid mutator transaction binding the contract method 0xebc73ab4.
//
// Solidity: function getMappingSlotAt(address target, bytes32 mappingSlot, uint256 idx) returns(bytes32 value)
func (_Vm *VmTransactor) GetMappingSlotAt(opts *bind.TransactOpts, target common.Address, mappingSlot [32]byte, idx *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "getMappingSlotAt", target, mappingSlot, idx)
}

// GetMappingSlotAt is a paid mutator transaction binding the contract method 0xebc73ab4.
//
// Solidity: function getMappingSlotAt(address target, bytes32 mappingSlot, uint256 idx) returns(bytes32 value)
func (_Vm *VmSession) GetMappingSlotAt(target common.Address, mappingSlot [32]byte, idx *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.GetMappingSlotAt(&_Vm.TransactOpts, target, mappingSlot, idx)
}

// GetMappingSlotAt is a paid mutator transaction binding the contract method 0xebc73ab4.
//
// Solidity: function getMappingSlotAt(address target, bytes32 mappingSlot, uint256 idx) returns(bytes32 value)
func (_Vm *VmTransactorSession) GetMappingSlotAt(target common.Address, mappingSlot [32]byte, idx *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.GetMappingSlotAt(&_Vm.TransactOpts, target, mappingSlot, idx)
}

// GetNonce0 is a paid mutator transaction binding the contract method 0xa5748aad.
//
// Solidity: function getNonce((address,uint256,uint256,uint256) wallet) returns(uint64 nonce)
func (_Vm *VmTransactor) GetNonce0(opts *bind.TransactOpts, wallet VmSafeWallet) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "getNonce0", wallet)
}

// GetNonce0 is a paid mutator transaction binding the contract method 0xa5748aad.
//
// Solidity: function getNonce((address,uint256,uint256,uint256) wallet) returns(uint64 nonce)
func (_Vm *VmSession) GetNonce0(wallet VmSafeWallet) (*types.Transaction, error) {
	return _Vm.Contract.GetNonce0(&_Vm.TransactOpts, wallet)
}

// GetNonce0 is a paid mutator transaction binding the contract method 0xa5748aad.
//
// Solidity: function getNonce((address,uint256,uint256,uint256) wallet) returns(uint64 nonce)
func (_Vm *VmTransactorSession) GetNonce0(wallet VmSafeWallet) (*types.Transaction, error) {
	return _Vm.Contract.GetNonce0(&_Vm.TransactOpts, wallet)
}

// GetRecordedLogs is a paid mutator transaction binding the contract method 0x191553a4.
//
// Solidity: function getRecordedLogs() returns((bytes32[],bytes,address)[] logs)
func (_Vm *VmTransactor) GetRecordedLogs(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "getRecordedLogs")
}

// GetRecordedLogs is a paid mutator transaction binding the contract method 0x191553a4.
//
// Solidity: function getRecordedLogs() returns((bytes32[],bytes,address)[] logs)
func (_Vm *VmSession) GetRecordedLogs() (*types.Transaction, error) {
	return _Vm.Contract.GetRecordedLogs(&_Vm.TransactOpts)
}

// GetRecordedLogs is a paid mutator transaction binding the contract method 0x191553a4.
//
// Solidity: function getRecordedLogs() returns((bytes32[],bytes,address)[] logs)
func (_Vm *VmTransactorSession) GetRecordedLogs() (*types.Transaction, error) {
	return _Vm.Contract.GetRecordedLogs(&_Vm.TransactOpts)
}

// IsDir is a paid mutator transaction binding the contract method 0x7d15d019.
//
// Solidity: function isDir(string path) returns(bool result)
func (_Vm *VmTransactor) IsDir(opts *bind.TransactOpts, path string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "isDir", path)
}

// IsDir is a paid mutator transaction binding the contract method 0x7d15d019.
//
// Solidity: function isDir(string path) returns(bool result)
func (_Vm *VmSession) IsDir(path string) (*types.Transaction, error) {
	return _Vm.Contract.IsDir(&_Vm.TransactOpts, path)
}

// IsDir is a paid mutator transaction binding the contract method 0x7d15d019.
//
// Solidity: function isDir(string path) returns(bool result)
func (_Vm *VmTransactorSession) IsDir(path string) (*types.Transaction, error) {
	return _Vm.Contract.IsDir(&_Vm.TransactOpts, path)
}

// IsFile is a paid mutator transaction binding the contract method 0xe0eb04d4.
//
// Solidity: function isFile(string path) returns(bool result)
func (_Vm *VmTransactor) IsFile(opts *bind.TransactOpts, path string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "isFile", path)
}

// IsFile is a paid mutator transaction binding the contract method 0xe0eb04d4.
//
// Solidity: function isFile(string path) returns(bool result)
func (_Vm *VmSession) IsFile(path string) (*types.Transaction, error) {
	return _Vm.Contract.IsFile(&_Vm.TransactOpts, path)
}

// IsFile is a paid mutator transaction binding the contract method 0xe0eb04d4.
//
// Solidity: function isFile(string path) returns(bool result)
func (_Vm *VmTransactorSession) IsFile(path string) (*types.Transaction, error) {
	return _Vm.Contract.IsFile(&_Vm.TransactOpts, path)
}

// Label is a paid mutator transaction binding the contract method 0xc657c718.
//
// Solidity: function label(address account, string newLabel) returns()
func (_Vm *VmTransactor) Label(opts *bind.TransactOpts, account common.Address, newLabel string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "label", account, newLabel)
}

// Label is a paid mutator transaction binding the contract method 0xc657c718.
//
// Solidity: function label(address account, string newLabel) returns()
func (_Vm *VmSession) Label(account common.Address, newLabel string) (*types.Transaction, error) {
	return _Vm.Contract.Label(&_Vm.TransactOpts, account, newLabel)
}

// Label is a paid mutator transaction binding the contract method 0xc657c718.
//
// Solidity: function label(address account, string newLabel) returns()
func (_Vm *VmTransactorSession) Label(account common.Address, newLabel string) (*types.Transaction, error) {
	return _Vm.Contract.Label(&_Vm.TransactOpts, account, newLabel)
}

// LoadAllocs is a paid mutator transaction binding the contract method 0xb3a056d7.
//
// Solidity: function loadAllocs(string pathToAllocsJson) returns()
func (_Vm *VmTransactor) LoadAllocs(opts *bind.TransactOpts, pathToAllocsJson string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "loadAllocs", pathToAllocsJson)
}

// LoadAllocs is a paid mutator transaction binding the contract method 0xb3a056d7.
//
// Solidity: function loadAllocs(string pathToAllocsJson) returns()
func (_Vm *VmSession) LoadAllocs(pathToAllocsJson string) (*types.Transaction, error) {
	return _Vm.Contract.LoadAllocs(&_Vm.TransactOpts, pathToAllocsJson)
}

// LoadAllocs is a paid mutator transaction binding the contract method 0xb3a056d7.
//
// Solidity: function loadAllocs(string pathToAllocsJson) returns()
func (_Vm *VmTransactorSession) LoadAllocs(pathToAllocsJson string) (*types.Transaction, error) {
	return _Vm.Contract.LoadAllocs(&_Vm.TransactOpts, pathToAllocsJson)
}

// MakePersistent is a paid mutator transaction binding the contract method 0x1d9e269e.
//
// Solidity: function makePersistent(address[] accounts) returns()
func (_Vm *VmTransactor) MakePersistent(opts *bind.TransactOpts, accounts []common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "makePersistent", accounts)
}

// MakePersistent is a paid mutator transaction binding the contract method 0x1d9e269e.
//
// Solidity: function makePersistent(address[] accounts) returns()
func (_Vm *VmSession) MakePersistent(accounts []common.Address) (*types.Transaction, error) {
	return _Vm.Contract.MakePersistent(&_Vm.TransactOpts, accounts)
}

// MakePersistent is a paid mutator transaction binding the contract method 0x1d9e269e.
//
// Solidity: function makePersistent(address[] accounts) returns()
func (_Vm *VmTransactorSession) MakePersistent(accounts []common.Address) (*types.Transaction, error) {
	return _Vm.Contract.MakePersistent(&_Vm.TransactOpts, accounts)
}

// MakePersistent0 is a paid mutator transaction binding the contract method 0x4074e0a8.
//
// Solidity: function makePersistent(address account0, address account1) returns()
func (_Vm *VmTransactor) MakePersistent0(opts *bind.TransactOpts, account0 common.Address, account1 common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "makePersistent0", account0, account1)
}

// MakePersistent0 is a paid mutator transaction binding the contract method 0x4074e0a8.
//
// Solidity: function makePersistent(address account0, address account1) returns()
func (_Vm *VmSession) MakePersistent0(account0 common.Address, account1 common.Address) (*types.Transaction, error) {
	return _Vm.Contract.MakePersistent0(&_Vm.TransactOpts, account0, account1)
}

// MakePersistent0 is a paid mutator transaction binding the contract method 0x4074e0a8.
//
// Solidity: function makePersistent(address account0, address account1) returns()
func (_Vm *VmTransactorSession) MakePersistent0(account0 common.Address, account1 common.Address) (*types.Transaction, error) {
	return _Vm.Contract.MakePersistent0(&_Vm.TransactOpts, account0, account1)
}

// MakePersistent1 is a paid mutator transaction binding the contract method 0x57e22dde.
//
// Solidity: function makePersistent(address account) returns()
func (_Vm *VmTransactor) MakePersistent1(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "makePersistent1", account)
}

// MakePersistent1 is a paid mutator transaction binding the contract method 0x57e22dde.
//
// Solidity: function makePersistent(address account) returns()
func (_Vm *VmSession) MakePersistent1(account common.Address) (*types.Transaction, error) {
	return _Vm.Contract.MakePersistent1(&_Vm.TransactOpts, account)
}

// MakePersistent1 is a paid mutator transaction binding the contract method 0x57e22dde.
//
// Solidity: function makePersistent(address account) returns()
func (_Vm *VmTransactorSession) MakePersistent1(account common.Address) (*types.Transaction, error) {
	return _Vm.Contract.MakePersistent1(&_Vm.TransactOpts, account)
}

// MakePersistent2 is a paid mutator transaction binding the contract method 0xefb77a75.
//
// Solidity: function makePersistent(address account0, address account1, address account2) returns()
func (_Vm *VmTransactor) MakePersistent2(opts *bind.TransactOpts, account0 common.Address, account1 common.Address, account2 common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "makePersistent2", account0, account1, account2)
}

// MakePersistent2 is a paid mutator transaction binding the contract method 0xefb77a75.
//
// Solidity: function makePersistent(address account0, address account1, address account2) returns()
func (_Vm *VmSession) MakePersistent2(account0 common.Address, account1 common.Address, account2 common.Address) (*types.Transaction, error) {
	return _Vm.Contract.MakePersistent2(&_Vm.TransactOpts, account0, account1, account2)
}

// MakePersistent2 is a paid mutator transaction binding the contract method 0xefb77a75.
//
// Solidity: function makePersistent(address account0, address account1, address account2) returns()
func (_Vm *VmTransactorSession) MakePersistent2(account0 common.Address, account1 common.Address, account2 common.Address) (*types.Transaction, error) {
	return _Vm.Contract.MakePersistent2(&_Vm.TransactOpts, account0, account1, account2)
}

// MockCall is a paid mutator transaction binding the contract method 0x81409b91.
//
// Solidity: function mockCall(address callee, uint256 msgValue, bytes data, bytes returnData) returns()
func (_Vm *VmTransactor) MockCall(opts *bind.TransactOpts, callee common.Address, msgValue *big.Int, data []byte, returnData []byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "mockCall", callee, msgValue, data, returnData)
}

// MockCall is a paid mutator transaction binding the contract method 0x81409b91.
//
// Solidity: function mockCall(address callee, uint256 msgValue, bytes data, bytes returnData) returns()
func (_Vm *VmSession) MockCall(callee common.Address, msgValue *big.Int, data []byte, returnData []byte) (*types.Transaction, error) {
	return _Vm.Contract.MockCall(&_Vm.TransactOpts, callee, msgValue, data, returnData)
}

// MockCall is a paid mutator transaction binding the contract method 0x81409b91.
//
// Solidity: function mockCall(address callee, uint256 msgValue, bytes data, bytes returnData) returns()
func (_Vm *VmTransactorSession) MockCall(callee common.Address, msgValue *big.Int, data []byte, returnData []byte) (*types.Transaction, error) {
	return _Vm.Contract.MockCall(&_Vm.TransactOpts, callee, msgValue, data, returnData)
}

// MockCall0 is a paid mutator transaction binding the contract method 0xb96213e4.
//
// Solidity: function mockCall(address callee, bytes data, bytes returnData) returns()
func (_Vm *VmTransactor) MockCall0(opts *bind.TransactOpts, callee common.Address, data []byte, returnData []byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "mockCall0", callee, data, returnData)
}

// MockCall0 is a paid mutator transaction binding the contract method 0xb96213e4.
//
// Solidity: function mockCall(address callee, bytes data, bytes returnData) returns()
func (_Vm *VmSession) MockCall0(callee common.Address, data []byte, returnData []byte) (*types.Transaction, error) {
	return _Vm.Contract.MockCall0(&_Vm.TransactOpts, callee, data, returnData)
}

// MockCall0 is a paid mutator transaction binding the contract method 0xb96213e4.
//
// Solidity: function mockCall(address callee, bytes data, bytes returnData) returns()
func (_Vm *VmTransactorSession) MockCall0(callee common.Address, data []byte, returnData []byte) (*types.Transaction, error) {
	return _Vm.Contract.MockCall0(&_Vm.TransactOpts, callee, data, returnData)
}

// MockCallRevert is a paid mutator transaction binding the contract method 0xd23cd037.
//
// Solidity: function mockCallRevert(address callee, uint256 msgValue, bytes data, bytes revertData) returns()
func (_Vm *VmTransactor) MockCallRevert(opts *bind.TransactOpts, callee common.Address, msgValue *big.Int, data []byte, revertData []byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "mockCallRevert", callee, msgValue, data, revertData)
}

// MockCallRevert is a paid mutator transaction binding the contract method 0xd23cd037.
//
// Solidity: function mockCallRevert(address callee, uint256 msgValue, bytes data, bytes revertData) returns()
func (_Vm *VmSession) MockCallRevert(callee common.Address, msgValue *big.Int, data []byte, revertData []byte) (*types.Transaction, error) {
	return _Vm.Contract.MockCallRevert(&_Vm.TransactOpts, callee, msgValue, data, revertData)
}

// MockCallRevert is a paid mutator transaction binding the contract method 0xd23cd037.
//
// Solidity: function mockCallRevert(address callee, uint256 msgValue, bytes data, bytes revertData) returns()
func (_Vm *VmTransactorSession) MockCallRevert(callee common.Address, msgValue *big.Int, data []byte, revertData []byte) (*types.Transaction, error) {
	return _Vm.Contract.MockCallRevert(&_Vm.TransactOpts, callee, msgValue, data, revertData)
}

// MockCallRevert0 is a paid mutator transaction binding the contract method 0xdbaad147.
//
// Solidity: function mockCallRevert(address callee, bytes data, bytes revertData) returns()
func (_Vm *VmTransactor) MockCallRevert0(opts *bind.TransactOpts, callee common.Address, data []byte, revertData []byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "mockCallRevert0", callee, data, revertData)
}

// MockCallRevert0 is a paid mutator transaction binding the contract method 0xdbaad147.
//
// Solidity: function mockCallRevert(address callee, bytes data, bytes revertData) returns()
func (_Vm *VmSession) MockCallRevert0(callee common.Address, data []byte, revertData []byte) (*types.Transaction, error) {
	return _Vm.Contract.MockCallRevert0(&_Vm.TransactOpts, callee, data, revertData)
}

// MockCallRevert0 is a paid mutator transaction binding the contract method 0xdbaad147.
//
// Solidity: function mockCallRevert(address callee, bytes data, bytes revertData) returns()
func (_Vm *VmTransactorSession) MockCallRevert0(callee common.Address, data []byte, revertData []byte) (*types.Transaction, error) {
	return _Vm.Contract.MockCallRevert0(&_Vm.TransactOpts, callee, data, revertData)
}

// PauseGasMetering is a paid mutator transaction binding the contract method 0xd1a5b36f.
//
// Solidity: function pauseGasMetering() returns()
func (_Vm *VmTransactor) PauseGasMetering(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "pauseGasMetering")
}

// PauseGasMetering is a paid mutator transaction binding the contract method 0xd1a5b36f.
//
// Solidity: function pauseGasMetering() returns()
func (_Vm *VmSession) PauseGasMetering() (*types.Transaction, error) {
	return _Vm.Contract.PauseGasMetering(&_Vm.TransactOpts)
}

// PauseGasMetering is a paid mutator transaction binding the contract method 0xd1a5b36f.
//
// Solidity: function pauseGasMetering() returns()
func (_Vm *VmTransactorSession) PauseGasMetering() (*types.Transaction, error) {
	return _Vm.Contract.PauseGasMetering(&_Vm.TransactOpts)
}

// Prank is a paid mutator transaction binding the contract method 0x47e50cce.
//
// Solidity: function prank(address msgSender, address txOrigin) returns()
func (_Vm *VmTransactor) Prank(opts *bind.TransactOpts, msgSender common.Address, txOrigin common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "prank", msgSender, txOrigin)
}

// Prank is a paid mutator transaction binding the contract method 0x47e50cce.
//
// Solidity: function prank(address msgSender, address txOrigin) returns()
func (_Vm *VmSession) Prank(msgSender common.Address, txOrigin common.Address) (*types.Transaction, error) {
	return _Vm.Contract.Prank(&_Vm.TransactOpts, msgSender, txOrigin)
}

// Prank is a paid mutator transaction binding the contract method 0x47e50cce.
//
// Solidity: function prank(address msgSender, address txOrigin) returns()
func (_Vm *VmTransactorSession) Prank(msgSender common.Address, txOrigin common.Address) (*types.Transaction, error) {
	return _Vm.Contract.Prank(&_Vm.TransactOpts, msgSender, txOrigin)
}

// Prank0 is a paid mutator transaction binding the contract method 0xca669fa7.
//
// Solidity: function prank(address msgSender) returns()
func (_Vm *VmTransactor) Prank0(opts *bind.TransactOpts, msgSender common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "prank0", msgSender)
}

// Prank0 is a paid mutator transaction binding the contract method 0xca669fa7.
//
// Solidity: function prank(address msgSender) returns()
func (_Vm *VmSession) Prank0(msgSender common.Address) (*types.Transaction, error) {
	return _Vm.Contract.Prank0(&_Vm.TransactOpts, msgSender)
}

// Prank0 is a paid mutator transaction binding the contract method 0xca669fa7.
//
// Solidity: function prank(address msgSender) returns()
func (_Vm *VmTransactorSession) Prank0(msgSender common.Address) (*types.Transaction, error) {
	return _Vm.Contract.Prank0(&_Vm.TransactOpts, msgSender)
}

// Prevrandao is a paid mutator transaction binding the contract method 0x3b925549.
//
// Solidity: function prevrandao(bytes32 newPrevrandao) returns()
func (_Vm *VmTransactor) Prevrandao(opts *bind.TransactOpts, newPrevrandao [32]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "prevrandao", newPrevrandao)
}

// Prevrandao is a paid mutator transaction binding the contract method 0x3b925549.
//
// Solidity: function prevrandao(bytes32 newPrevrandao) returns()
func (_Vm *VmSession) Prevrandao(newPrevrandao [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.Prevrandao(&_Vm.TransactOpts, newPrevrandao)
}

// Prevrandao is a paid mutator transaction binding the contract method 0x3b925549.
//
// Solidity: function prevrandao(bytes32 newPrevrandao) returns()
func (_Vm *VmTransactorSession) Prevrandao(newPrevrandao [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.Prevrandao(&_Vm.TransactOpts, newPrevrandao)
}

// Prevrandao0 is a paid mutator transaction binding the contract method 0x9cb1c0d4.
//
// Solidity: function prevrandao(uint256 newPrevrandao) returns()
func (_Vm *VmTransactor) Prevrandao0(opts *bind.TransactOpts, newPrevrandao *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "prevrandao0", newPrevrandao)
}

// Prevrandao0 is a paid mutator transaction binding the contract method 0x9cb1c0d4.
//
// Solidity: function prevrandao(uint256 newPrevrandao) returns()
func (_Vm *VmSession) Prevrandao0(newPrevrandao *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.Prevrandao0(&_Vm.TransactOpts, newPrevrandao)
}

// Prevrandao0 is a paid mutator transaction binding the contract method 0x9cb1c0d4.
//
// Solidity: function prevrandao(uint256 newPrevrandao) returns()
func (_Vm *VmTransactorSession) Prevrandao0(newPrevrandao *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.Prevrandao0(&_Vm.TransactOpts, newPrevrandao)
}

// Prompt is a paid mutator transaction binding the contract method 0x47eaf474.
//
// Solidity: function prompt(string promptText) returns(string input)
func (_Vm *VmTransactor) Prompt(opts *bind.TransactOpts, promptText string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "prompt", promptText)
}

// Prompt is a paid mutator transaction binding the contract method 0x47eaf474.
//
// Solidity: function prompt(string promptText) returns(string input)
func (_Vm *VmSession) Prompt(promptText string) (*types.Transaction, error) {
	return _Vm.Contract.Prompt(&_Vm.TransactOpts, promptText)
}

// Prompt is a paid mutator transaction binding the contract method 0x47eaf474.
//
// Solidity: function prompt(string promptText) returns(string input)
func (_Vm *VmTransactorSession) Prompt(promptText string) (*types.Transaction, error) {
	return _Vm.Contract.Prompt(&_Vm.TransactOpts, promptText)
}

// PromptAddress is a paid mutator transaction binding the contract method 0x62ee05f4.
//
// Solidity: function promptAddress(string promptText) returns(address)
func (_Vm *VmTransactor) PromptAddress(opts *bind.TransactOpts, promptText string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "promptAddress", promptText)
}

// PromptAddress is a paid mutator transaction binding the contract method 0x62ee05f4.
//
// Solidity: function promptAddress(string promptText) returns(address)
func (_Vm *VmSession) PromptAddress(promptText string) (*types.Transaction, error) {
	return _Vm.Contract.PromptAddress(&_Vm.TransactOpts, promptText)
}

// PromptAddress is a paid mutator transaction binding the contract method 0x62ee05f4.
//
// Solidity: function promptAddress(string promptText) returns(address)
func (_Vm *VmTransactorSession) PromptAddress(promptText string) (*types.Transaction, error) {
	return _Vm.Contract.PromptAddress(&_Vm.TransactOpts, promptText)
}

// PromptSecret is a paid mutator transaction binding the contract method 0x1e279d41.
//
// Solidity: function promptSecret(string promptText) returns(string input)
func (_Vm *VmTransactor) PromptSecret(opts *bind.TransactOpts, promptText string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "promptSecret", promptText)
}

// PromptSecret is a paid mutator transaction binding the contract method 0x1e279d41.
//
// Solidity: function promptSecret(string promptText) returns(string input)
func (_Vm *VmSession) PromptSecret(promptText string) (*types.Transaction, error) {
	return _Vm.Contract.PromptSecret(&_Vm.TransactOpts, promptText)
}

// PromptSecret is a paid mutator transaction binding the contract method 0x1e279d41.
//
// Solidity: function promptSecret(string promptText) returns(string input)
func (_Vm *VmTransactorSession) PromptSecret(promptText string) (*types.Transaction, error) {
	return _Vm.Contract.PromptSecret(&_Vm.TransactOpts, promptText)
}

// PromptSecretUint is a paid mutator transaction binding the contract method 0x69ca02b7.
//
// Solidity: function promptSecretUint(string promptText) returns(uint256)
func (_Vm *VmTransactor) PromptSecretUint(opts *bind.TransactOpts, promptText string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "promptSecretUint", promptText)
}

// PromptSecretUint is a paid mutator transaction binding the contract method 0x69ca02b7.
//
// Solidity: function promptSecretUint(string promptText) returns(uint256)
func (_Vm *VmSession) PromptSecretUint(promptText string) (*types.Transaction, error) {
	return _Vm.Contract.PromptSecretUint(&_Vm.TransactOpts, promptText)
}

// PromptSecretUint is a paid mutator transaction binding the contract method 0x69ca02b7.
//
// Solidity: function promptSecretUint(string promptText) returns(uint256)
func (_Vm *VmTransactorSession) PromptSecretUint(promptText string) (*types.Transaction, error) {
	return _Vm.Contract.PromptSecretUint(&_Vm.TransactOpts, promptText)
}

// PromptUint is a paid mutator transaction binding the contract method 0x652fd489.
//
// Solidity: function promptUint(string promptText) returns(uint256)
func (_Vm *VmTransactor) PromptUint(opts *bind.TransactOpts, promptText string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "promptUint", promptText)
}

// PromptUint is a paid mutator transaction binding the contract method 0x652fd489.
//
// Solidity: function promptUint(string promptText) returns(uint256)
func (_Vm *VmSession) PromptUint(promptText string) (*types.Transaction, error) {
	return _Vm.Contract.PromptUint(&_Vm.TransactOpts, promptText)
}

// PromptUint is a paid mutator transaction binding the contract method 0x652fd489.
//
// Solidity: function promptUint(string promptText) returns(uint256)
func (_Vm *VmTransactorSession) PromptUint(promptText string) (*types.Transaction, error) {
	return _Vm.Contract.PromptUint(&_Vm.TransactOpts, promptText)
}

// RandomAddress is a paid mutator transaction binding the contract method 0xd5bee9f5.
//
// Solidity: function randomAddress() returns(address)
func (_Vm *VmTransactor) RandomAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "randomAddress")
}

// RandomAddress is a paid mutator transaction binding the contract method 0xd5bee9f5.
//
// Solidity: function randomAddress() returns(address)
func (_Vm *VmSession) RandomAddress() (*types.Transaction, error) {
	return _Vm.Contract.RandomAddress(&_Vm.TransactOpts)
}

// RandomAddress is a paid mutator transaction binding the contract method 0xd5bee9f5.
//
// Solidity: function randomAddress() returns(address)
func (_Vm *VmTransactorSession) RandomAddress() (*types.Transaction, error) {
	return _Vm.Contract.RandomAddress(&_Vm.TransactOpts)
}

// RandomUint is a paid mutator transaction binding the contract method 0x25124730.
//
// Solidity: function randomUint() returns(uint256)
func (_Vm *VmTransactor) RandomUint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "randomUint")
}

// RandomUint is a paid mutator transaction binding the contract method 0x25124730.
//
// Solidity: function randomUint() returns(uint256)
func (_Vm *VmSession) RandomUint() (*types.Transaction, error) {
	return _Vm.Contract.RandomUint(&_Vm.TransactOpts)
}

// RandomUint is a paid mutator transaction binding the contract method 0x25124730.
//
// Solidity: function randomUint() returns(uint256)
func (_Vm *VmTransactorSession) RandomUint() (*types.Transaction, error) {
	return _Vm.Contract.RandomUint(&_Vm.TransactOpts)
}

// RandomUint0 is a paid mutator transaction binding the contract method 0xd61b051b.
//
// Solidity: function randomUint(uint256 min, uint256 max) returns(uint256)
func (_Vm *VmTransactor) RandomUint0(opts *bind.TransactOpts, min *big.Int, max *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "randomUint0", min, max)
}

// RandomUint0 is a paid mutator transaction binding the contract method 0xd61b051b.
//
// Solidity: function randomUint(uint256 min, uint256 max) returns(uint256)
func (_Vm *VmSession) RandomUint0(min *big.Int, max *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.RandomUint0(&_Vm.TransactOpts, min, max)
}

// RandomUint0 is a paid mutator transaction binding the contract method 0xd61b051b.
//
// Solidity: function randomUint(uint256 min, uint256 max) returns(uint256)
func (_Vm *VmTransactorSession) RandomUint0(min *big.Int, max *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.RandomUint0(&_Vm.TransactOpts, min, max)
}

// ReadCallers is a paid mutator transaction binding the contract method 0x4ad0bac9.
//
// Solidity: function readCallers() returns(uint8 callerMode, address msgSender, address txOrigin)
func (_Vm *VmTransactor) ReadCallers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "readCallers")
}

// ReadCallers is a paid mutator transaction binding the contract method 0x4ad0bac9.
//
// Solidity: function readCallers() returns(uint8 callerMode, address msgSender, address txOrigin)
func (_Vm *VmSession) ReadCallers() (*types.Transaction, error) {
	return _Vm.Contract.ReadCallers(&_Vm.TransactOpts)
}

// ReadCallers is a paid mutator transaction binding the contract method 0x4ad0bac9.
//
// Solidity: function readCallers() returns(uint8 callerMode, address msgSender, address txOrigin)
func (_Vm *VmTransactorSession) ReadCallers() (*types.Transaction, error) {
	return _Vm.Contract.ReadCallers(&_Vm.TransactOpts)
}

// Record is a paid mutator transaction binding the contract method 0x266cf109.
//
// Solidity: function record() returns()
func (_Vm *VmTransactor) Record(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "record")
}

// Record is a paid mutator transaction binding the contract method 0x266cf109.
//
// Solidity: function record() returns()
func (_Vm *VmSession) Record() (*types.Transaction, error) {
	return _Vm.Contract.Record(&_Vm.TransactOpts)
}

// Record is a paid mutator transaction binding the contract method 0x266cf109.
//
// Solidity: function record() returns()
func (_Vm *VmTransactorSession) Record() (*types.Transaction, error) {
	return _Vm.Contract.Record(&_Vm.TransactOpts)
}

// RecordLogs is a paid mutator transaction binding the contract method 0x41af2f52.
//
// Solidity: function recordLogs() returns()
func (_Vm *VmTransactor) RecordLogs(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "recordLogs")
}

// RecordLogs is a paid mutator transaction binding the contract method 0x41af2f52.
//
// Solidity: function recordLogs() returns()
func (_Vm *VmSession) RecordLogs() (*types.Transaction, error) {
	return _Vm.Contract.RecordLogs(&_Vm.TransactOpts)
}

// RecordLogs is a paid mutator transaction binding the contract method 0x41af2f52.
//
// Solidity: function recordLogs() returns()
func (_Vm *VmTransactorSession) RecordLogs() (*types.Transaction, error) {
	return _Vm.Contract.RecordLogs(&_Vm.TransactOpts)
}

// RememberKey is a paid mutator transaction binding the contract method 0x22100064.
//
// Solidity: function rememberKey(uint256 privateKey) returns(address keyAddr)
func (_Vm *VmTransactor) RememberKey(opts *bind.TransactOpts, privateKey *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "rememberKey", privateKey)
}

// RememberKey is a paid mutator transaction binding the contract method 0x22100064.
//
// Solidity: function rememberKey(uint256 privateKey) returns(address keyAddr)
func (_Vm *VmSession) RememberKey(privateKey *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.RememberKey(&_Vm.TransactOpts, privateKey)
}

// RememberKey is a paid mutator transaction binding the contract method 0x22100064.
//
// Solidity: function rememberKey(uint256 privateKey) returns(address keyAddr)
func (_Vm *VmTransactorSession) RememberKey(privateKey *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.RememberKey(&_Vm.TransactOpts, privateKey)
}

// RemoveDir is a paid mutator transaction binding the contract method 0x45c62011.
//
// Solidity: function removeDir(string path, bool recursive) returns()
func (_Vm *VmTransactor) RemoveDir(opts *bind.TransactOpts, path string, recursive bool) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "removeDir", path, recursive)
}

// RemoveDir is a paid mutator transaction binding the contract method 0x45c62011.
//
// Solidity: function removeDir(string path, bool recursive) returns()
func (_Vm *VmSession) RemoveDir(path string, recursive bool) (*types.Transaction, error) {
	return _Vm.Contract.RemoveDir(&_Vm.TransactOpts, path, recursive)
}

// RemoveDir is a paid mutator transaction binding the contract method 0x45c62011.
//
// Solidity: function removeDir(string path, bool recursive) returns()
func (_Vm *VmTransactorSession) RemoveDir(path string, recursive bool) (*types.Transaction, error) {
	return _Vm.Contract.RemoveDir(&_Vm.TransactOpts, path, recursive)
}

// RemoveFile is a paid mutator transaction binding the contract method 0xf1afe04d.
//
// Solidity: function removeFile(string path) returns()
func (_Vm *VmTransactor) RemoveFile(opts *bind.TransactOpts, path string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "removeFile", path)
}

// RemoveFile is a paid mutator transaction binding the contract method 0xf1afe04d.
//
// Solidity: function removeFile(string path) returns()
func (_Vm *VmSession) RemoveFile(path string) (*types.Transaction, error) {
	return _Vm.Contract.RemoveFile(&_Vm.TransactOpts, path)
}

// RemoveFile is a paid mutator transaction binding the contract method 0xf1afe04d.
//
// Solidity: function removeFile(string path) returns()
func (_Vm *VmTransactorSession) RemoveFile(path string) (*types.Transaction, error) {
	return _Vm.Contract.RemoveFile(&_Vm.TransactOpts, path)
}

// ResetNonce is a paid mutator transaction binding the contract method 0x1c72346d.
//
// Solidity: function resetNonce(address account) returns()
func (_Vm *VmTransactor) ResetNonce(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "resetNonce", account)
}

// ResetNonce is a paid mutator transaction binding the contract method 0x1c72346d.
//
// Solidity: function resetNonce(address account) returns()
func (_Vm *VmSession) ResetNonce(account common.Address) (*types.Transaction, error) {
	return _Vm.Contract.ResetNonce(&_Vm.TransactOpts, account)
}

// ResetNonce is a paid mutator transaction binding the contract method 0x1c72346d.
//
// Solidity: function resetNonce(address account) returns()
func (_Vm *VmTransactorSession) ResetNonce(account common.Address) (*types.Transaction, error) {
	return _Vm.Contract.ResetNonce(&_Vm.TransactOpts, account)
}

// ResumeGasMetering is a paid mutator transaction binding the contract method 0x2bcd50e0.
//
// Solidity: function resumeGasMetering() returns()
func (_Vm *VmTransactor) ResumeGasMetering(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "resumeGasMetering")
}

// ResumeGasMetering is a paid mutator transaction binding the contract method 0x2bcd50e0.
//
// Solidity: function resumeGasMetering() returns()
func (_Vm *VmSession) ResumeGasMetering() (*types.Transaction, error) {
	return _Vm.Contract.ResumeGasMetering(&_Vm.TransactOpts)
}

// ResumeGasMetering is a paid mutator transaction binding the contract method 0x2bcd50e0.
//
// Solidity: function resumeGasMetering() returns()
func (_Vm *VmTransactorSession) ResumeGasMetering() (*types.Transaction, error) {
	return _Vm.Contract.ResumeGasMetering(&_Vm.TransactOpts)
}

// RevertTo is a paid mutator transaction binding the contract method 0x44d7f0a4.
//
// Solidity: function revertTo(uint256 snapshotId) returns(bool success)
func (_Vm *VmTransactor) RevertTo(opts *bind.TransactOpts, snapshotId *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "revertTo", snapshotId)
}

// RevertTo is a paid mutator transaction binding the contract method 0x44d7f0a4.
//
// Solidity: function revertTo(uint256 snapshotId) returns(bool success)
func (_Vm *VmSession) RevertTo(snapshotId *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.RevertTo(&_Vm.TransactOpts, snapshotId)
}

// RevertTo is a paid mutator transaction binding the contract method 0x44d7f0a4.
//
// Solidity: function revertTo(uint256 snapshotId) returns(bool success)
func (_Vm *VmTransactorSession) RevertTo(snapshotId *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.RevertTo(&_Vm.TransactOpts, snapshotId)
}

// RevertToAndDelete is a paid mutator transaction binding the contract method 0x03e0aca9.
//
// Solidity: function revertToAndDelete(uint256 snapshotId) returns(bool success)
func (_Vm *VmTransactor) RevertToAndDelete(opts *bind.TransactOpts, snapshotId *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "revertToAndDelete", snapshotId)
}

// RevertToAndDelete is a paid mutator transaction binding the contract method 0x03e0aca9.
//
// Solidity: function revertToAndDelete(uint256 snapshotId) returns(bool success)
func (_Vm *VmSession) RevertToAndDelete(snapshotId *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.RevertToAndDelete(&_Vm.TransactOpts, snapshotId)
}

// RevertToAndDelete is a paid mutator transaction binding the contract method 0x03e0aca9.
//
// Solidity: function revertToAndDelete(uint256 snapshotId) returns(bool success)
func (_Vm *VmTransactorSession) RevertToAndDelete(snapshotId *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.RevertToAndDelete(&_Vm.TransactOpts, snapshotId)
}

// RevokePersistent is a paid mutator transaction binding the contract method 0x3ce969e6.
//
// Solidity: function revokePersistent(address[] accounts) returns()
func (_Vm *VmTransactor) RevokePersistent(opts *bind.TransactOpts, accounts []common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "revokePersistent", accounts)
}

// RevokePersistent is a paid mutator transaction binding the contract method 0x3ce969e6.
//
// Solidity: function revokePersistent(address[] accounts) returns()
func (_Vm *VmSession) RevokePersistent(accounts []common.Address) (*types.Transaction, error) {
	return _Vm.Contract.RevokePersistent(&_Vm.TransactOpts, accounts)
}

// RevokePersistent is a paid mutator transaction binding the contract method 0x3ce969e6.
//
// Solidity: function revokePersistent(address[] accounts) returns()
func (_Vm *VmTransactorSession) RevokePersistent(accounts []common.Address) (*types.Transaction, error) {
	return _Vm.Contract.RevokePersistent(&_Vm.TransactOpts, accounts)
}

// RevokePersistent0 is a paid mutator transaction binding the contract method 0x997a0222.
//
// Solidity: function revokePersistent(address account) returns()
func (_Vm *VmTransactor) RevokePersistent0(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "revokePersistent0", account)
}

// RevokePersistent0 is a paid mutator transaction binding the contract method 0x997a0222.
//
// Solidity: function revokePersistent(address account) returns()
func (_Vm *VmSession) RevokePersistent0(account common.Address) (*types.Transaction, error) {
	return _Vm.Contract.RevokePersistent0(&_Vm.TransactOpts, account)
}

// RevokePersistent0 is a paid mutator transaction binding the contract method 0x997a0222.
//
// Solidity: function revokePersistent(address account) returns()
func (_Vm *VmTransactorSession) RevokePersistent0(account common.Address) (*types.Transaction, error) {
	return _Vm.Contract.RevokePersistent0(&_Vm.TransactOpts, account)
}

// Roll is a paid mutator transaction binding the contract method 0x1f7b4f30.
//
// Solidity: function roll(uint256 newHeight) returns()
func (_Vm *VmTransactor) Roll(opts *bind.TransactOpts, newHeight *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "roll", newHeight)
}

// Roll is a paid mutator transaction binding the contract method 0x1f7b4f30.
//
// Solidity: function roll(uint256 newHeight) returns()
func (_Vm *VmSession) Roll(newHeight *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.Roll(&_Vm.TransactOpts, newHeight)
}

// Roll is a paid mutator transaction binding the contract method 0x1f7b4f30.
//
// Solidity: function roll(uint256 newHeight) returns()
func (_Vm *VmTransactorSession) Roll(newHeight *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.Roll(&_Vm.TransactOpts, newHeight)
}

// RollFork is a paid mutator transaction binding the contract method 0x0f29772b.
//
// Solidity: function rollFork(bytes32 txHash) returns()
func (_Vm *VmTransactor) RollFork(opts *bind.TransactOpts, txHash [32]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "rollFork", txHash)
}

// RollFork is a paid mutator transaction binding the contract method 0x0f29772b.
//
// Solidity: function rollFork(bytes32 txHash) returns()
func (_Vm *VmSession) RollFork(txHash [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.RollFork(&_Vm.TransactOpts, txHash)
}

// RollFork is a paid mutator transaction binding the contract method 0x0f29772b.
//
// Solidity: function rollFork(bytes32 txHash) returns()
func (_Vm *VmTransactorSession) RollFork(txHash [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.RollFork(&_Vm.TransactOpts, txHash)
}

// RollFork0 is a paid mutator transaction binding the contract method 0xd74c83a4.
//
// Solidity: function rollFork(uint256 forkId, uint256 blockNumber) returns()
func (_Vm *VmTransactor) RollFork0(opts *bind.TransactOpts, forkId *big.Int, blockNumber *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "rollFork0", forkId, blockNumber)
}

// RollFork0 is a paid mutator transaction binding the contract method 0xd74c83a4.
//
// Solidity: function rollFork(uint256 forkId, uint256 blockNumber) returns()
func (_Vm *VmSession) RollFork0(forkId *big.Int, blockNumber *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.RollFork0(&_Vm.TransactOpts, forkId, blockNumber)
}

// RollFork0 is a paid mutator transaction binding the contract method 0xd74c83a4.
//
// Solidity: function rollFork(uint256 forkId, uint256 blockNumber) returns()
func (_Vm *VmTransactorSession) RollFork0(forkId *big.Int, blockNumber *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.RollFork0(&_Vm.TransactOpts, forkId, blockNumber)
}

// RollFork1 is a paid mutator transaction binding the contract method 0xd9bbf3a1.
//
// Solidity: function rollFork(uint256 blockNumber) returns()
func (_Vm *VmTransactor) RollFork1(opts *bind.TransactOpts, blockNumber *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "rollFork1", blockNumber)
}

// RollFork1 is a paid mutator transaction binding the contract method 0xd9bbf3a1.
//
// Solidity: function rollFork(uint256 blockNumber) returns()
func (_Vm *VmSession) RollFork1(blockNumber *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.RollFork1(&_Vm.TransactOpts, blockNumber)
}

// RollFork1 is a paid mutator transaction binding the contract method 0xd9bbf3a1.
//
// Solidity: function rollFork(uint256 blockNumber) returns()
func (_Vm *VmTransactorSession) RollFork1(blockNumber *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.RollFork1(&_Vm.TransactOpts, blockNumber)
}

// RollFork2 is a paid mutator transaction binding the contract method 0xf2830f7b.
//
// Solidity: function rollFork(uint256 forkId, bytes32 txHash) returns()
func (_Vm *VmTransactor) RollFork2(opts *bind.TransactOpts, forkId *big.Int, txHash [32]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "rollFork2", forkId, txHash)
}

// RollFork2 is a paid mutator transaction binding the contract method 0xf2830f7b.
//
// Solidity: function rollFork(uint256 forkId, bytes32 txHash) returns()
func (_Vm *VmSession) RollFork2(forkId *big.Int, txHash [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.RollFork2(&_Vm.TransactOpts, forkId, txHash)
}

// RollFork2 is a paid mutator transaction binding the contract method 0xf2830f7b.
//
// Solidity: function rollFork(uint256 forkId, bytes32 txHash) returns()
func (_Vm *VmTransactorSession) RollFork2(forkId *big.Int, txHash [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.RollFork2(&_Vm.TransactOpts, forkId, txHash)
}

// Rpc is a paid mutator transaction binding the contract method 0x1206c8a8.
//
// Solidity: function rpc(string method, string params) returns(bytes data)
func (_Vm *VmTransactor) Rpc(opts *bind.TransactOpts, method string, params string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "rpc", method, params)
}

// Rpc is a paid mutator transaction binding the contract method 0x1206c8a8.
//
// Solidity: function rpc(string method, string params) returns(bytes data)
func (_Vm *VmSession) Rpc(method string, params string) (*types.Transaction, error) {
	return _Vm.Contract.Rpc(&_Vm.TransactOpts, method, params)
}

// Rpc is a paid mutator transaction binding the contract method 0x1206c8a8.
//
// Solidity: function rpc(string method, string params) returns(bytes data)
func (_Vm *VmTransactorSession) Rpc(method string, params string) (*types.Transaction, error) {
	return _Vm.Contract.Rpc(&_Vm.TransactOpts, method, params)
}

// SelectFork is a paid mutator transaction binding the contract method 0x9ebf6827.
//
// Solidity: function selectFork(uint256 forkId) returns()
func (_Vm *VmTransactor) SelectFork(opts *bind.TransactOpts, forkId *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "selectFork", forkId)
}

// SelectFork is a paid mutator transaction binding the contract method 0x9ebf6827.
//
// Solidity: function selectFork(uint256 forkId) returns()
func (_Vm *VmSession) SelectFork(forkId *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.SelectFork(&_Vm.TransactOpts, forkId)
}

// SelectFork is a paid mutator transaction binding the contract method 0x9ebf6827.
//
// Solidity: function selectFork(uint256 forkId) returns()
func (_Vm *VmTransactorSession) SelectFork(forkId *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.SelectFork(&_Vm.TransactOpts, forkId)
}

// SerializeAddress is a paid mutator transaction binding the contract method 0x1e356e1a.
//
// Solidity: function serializeAddress(string objectKey, string valueKey, address[] values) returns(string json)
func (_Vm *VmTransactor) SerializeAddress(opts *bind.TransactOpts, objectKey string, valueKey string, values []common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "serializeAddress", objectKey, valueKey, values)
}

// SerializeAddress is a paid mutator transaction binding the contract method 0x1e356e1a.
//
// Solidity: function serializeAddress(string objectKey, string valueKey, address[] values) returns(string json)
func (_Vm *VmSession) SerializeAddress(objectKey string, valueKey string, values []common.Address) (*types.Transaction, error) {
	return _Vm.Contract.SerializeAddress(&_Vm.TransactOpts, objectKey, valueKey, values)
}

// SerializeAddress is a paid mutator transaction binding the contract method 0x1e356e1a.
//
// Solidity: function serializeAddress(string objectKey, string valueKey, address[] values) returns(string json)
func (_Vm *VmTransactorSession) SerializeAddress(objectKey string, valueKey string, values []common.Address) (*types.Transaction, error) {
	return _Vm.Contract.SerializeAddress(&_Vm.TransactOpts, objectKey, valueKey, values)
}

// SerializeAddress0 is a paid mutator transaction binding the contract method 0x972c6062.
//
// Solidity: function serializeAddress(string objectKey, string valueKey, address value) returns(string json)
func (_Vm *VmTransactor) SerializeAddress0(opts *bind.TransactOpts, objectKey string, valueKey string, value common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "serializeAddress0", objectKey, valueKey, value)
}

// SerializeAddress0 is a paid mutator transaction binding the contract method 0x972c6062.
//
// Solidity: function serializeAddress(string objectKey, string valueKey, address value) returns(string json)
func (_Vm *VmSession) SerializeAddress0(objectKey string, valueKey string, value common.Address) (*types.Transaction, error) {
	return _Vm.Contract.SerializeAddress0(&_Vm.TransactOpts, objectKey, valueKey, value)
}

// SerializeAddress0 is a paid mutator transaction binding the contract method 0x972c6062.
//
// Solidity: function serializeAddress(string objectKey, string valueKey, address value) returns(string json)
func (_Vm *VmTransactorSession) SerializeAddress0(objectKey string, valueKey string, value common.Address) (*types.Transaction, error) {
	return _Vm.Contract.SerializeAddress0(&_Vm.TransactOpts, objectKey, valueKey, value)
}

// SerializeBool is a paid mutator transaction binding the contract method 0x92925aa1.
//
// Solidity: function serializeBool(string objectKey, string valueKey, bool[] values) returns(string json)
func (_Vm *VmTransactor) SerializeBool(opts *bind.TransactOpts, objectKey string, valueKey string, values []bool) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "serializeBool", objectKey, valueKey, values)
}

// SerializeBool is a paid mutator transaction binding the contract method 0x92925aa1.
//
// Solidity: function serializeBool(string objectKey, string valueKey, bool[] values) returns(string json)
func (_Vm *VmSession) SerializeBool(objectKey string, valueKey string, values []bool) (*types.Transaction, error) {
	return _Vm.Contract.SerializeBool(&_Vm.TransactOpts, objectKey, valueKey, values)
}

// SerializeBool is a paid mutator transaction binding the contract method 0x92925aa1.
//
// Solidity: function serializeBool(string objectKey, string valueKey, bool[] values) returns(string json)
func (_Vm *VmTransactorSession) SerializeBool(objectKey string, valueKey string, values []bool) (*types.Transaction, error) {
	return _Vm.Contract.SerializeBool(&_Vm.TransactOpts, objectKey, valueKey, values)
}

// SerializeBool0 is a paid mutator transaction binding the contract method 0xac22e971.
//
// Solidity: function serializeBool(string objectKey, string valueKey, bool value) returns(string json)
func (_Vm *VmTransactor) SerializeBool0(opts *bind.TransactOpts, objectKey string, valueKey string, value bool) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "serializeBool0", objectKey, valueKey, value)
}

// SerializeBool0 is a paid mutator transaction binding the contract method 0xac22e971.
//
// Solidity: function serializeBool(string objectKey, string valueKey, bool value) returns(string json)
func (_Vm *VmSession) SerializeBool0(objectKey string, valueKey string, value bool) (*types.Transaction, error) {
	return _Vm.Contract.SerializeBool0(&_Vm.TransactOpts, objectKey, valueKey, value)
}

// SerializeBool0 is a paid mutator transaction binding the contract method 0xac22e971.
//
// Solidity: function serializeBool(string objectKey, string valueKey, bool value) returns(string json)
func (_Vm *VmTransactorSession) SerializeBool0(objectKey string, valueKey string, value bool) (*types.Transaction, error) {
	return _Vm.Contract.SerializeBool0(&_Vm.TransactOpts, objectKey, valueKey, value)
}

// SerializeBytes is a paid mutator transaction binding the contract method 0x9884b232.
//
// Solidity: function serializeBytes(string objectKey, string valueKey, bytes[] values) returns(string json)
func (_Vm *VmTransactor) SerializeBytes(opts *bind.TransactOpts, objectKey string, valueKey string, values [][]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "serializeBytes", objectKey, valueKey, values)
}

// SerializeBytes is a paid mutator transaction binding the contract method 0x9884b232.
//
// Solidity: function serializeBytes(string objectKey, string valueKey, bytes[] values) returns(string json)
func (_Vm *VmSession) SerializeBytes(objectKey string, valueKey string, values [][]byte) (*types.Transaction, error) {
	return _Vm.Contract.SerializeBytes(&_Vm.TransactOpts, objectKey, valueKey, values)
}

// SerializeBytes is a paid mutator transaction binding the contract method 0x9884b232.
//
// Solidity: function serializeBytes(string objectKey, string valueKey, bytes[] values) returns(string json)
func (_Vm *VmTransactorSession) SerializeBytes(objectKey string, valueKey string, values [][]byte) (*types.Transaction, error) {
	return _Vm.Contract.SerializeBytes(&_Vm.TransactOpts, objectKey, valueKey, values)
}

// SerializeBytes0 is a paid mutator transaction binding the contract method 0xf21d52c7.
//
// Solidity: function serializeBytes(string objectKey, string valueKey, bytes value) returns(string json)
func (_Vm *VmTransactor) SerializeBytes0(opts *bind.TransactOpts, objectKey string, valueKey string, value []byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "serializeBytes0", objectKey, valueKey, value)
}

// SerializeBytes0 is a paid mutator transaction binding the contract method 0xf21d52c7.
//
// Solidity: function serializeBytes(string objectKey, string valueKey, bytes value) returns(string json)
func (_Vm *VmSession) SerializeBytes0(objectKey string, valueKey string, value []byte) (*types.Transaction, error) {
	return _Vm.Contract.SerializeBytes0(&_Vm.TransactOpts, objectKey, valueKey, value)
}

// SerializeBytes0 is a paid mutator transaction binding the contract method 0xf21d52c7.
//
// Solidity: function serializeBytes(string objectKey, string valueKey, bytes value) returns(string json)
func (_Vm *VmTransactorSession) SerializeBytes0(objectKey string, valueKey string, value []byte) (*types.Transaction, error) {
	return _Vm.Contract.SerializeBytes0(&_Vm.TransactOpts, objectKey, valueKey, value)
}

// SerializeBytes32 is a paid mutator transaction binding the contract method 0x201e43e2.
//
// Solidity: function serializeBytes32(string objectKey, string valueKey, bytes32[] values) returns(string json)
func (_Vm *VmTransactor) SerializeBytes32(opts *bind.TransactOpts, objectKey string, valueKey string, values [][32]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "serializeBytes32", objectKey, valueKey, values)
}

// SerializeBytes32 is a paid mutator transaction binding the contract method 0x201e43e2.
//
// Solidity: function serializeBytes32(string objectKey, string valueKey, bytes32[] values) returns(string json)
func (_Vm *VmSession) SerializeBytes32(objectKey string, valueKey string, values [][32]byte) (*types.Transaction, error) {
	return _Vm.Contract.SerializeBytes32(&_Vm.TransactOpts, objectKey, valueKey, values)
}

// SerializeBytes32 is a paid mutator transaction binding the contract method 0x201e43e2.
//
// Solidity: function serializeBytes32(string objectKey, string valueKey, bytes32[] values) returns(string json)
func (_Vm *VmTransactorSession) SerializeBytes32(objectKey string, valueKey string, values [][32]byte) (*types.Transaction, error) {
	return _Vm.Contract.SerializeBytes32(&_Vm.TransactOpts, objectKey, valueKey, values)
}

// SerializeBytes320 is a paid mutator transaction binding the contract method 0x2d812b44.
//
// Solidity: function serializeBytes32(string objectKey, string valueKey, bytes32 value) returns(string json)
func (_Vm *VmTransactor) SerializeBytes320(opts *bind.TransactOpts, objectKey string, valueKey string, value [32]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "serializeBytes320", objectKey, valueKey, value)
}

// SerializeBytes320 is a paid mutator transaction binding the contract method 0x2d812b44.
//
// Solidity: function serializeBytes32(string objectKey, string valueKey, bytes32 value) returns(string json)
func (_Vm *VmSession) SerializeBytes320(objectKey string, valueKey string, value [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.SerializeBytes320(&_Vm.TransactOpts, objectKey, valueKey, value)
}

// SerializeBytes320 is a paid mutator transaction binding the contract method 0x2d812b44.
//
// Solidity: function serializeBytes32(string objectKey, string valueKey, bytes32 value) returns(string json)
func (_Vm *VmTransactorSession) SerializeBytes320(objectKey string, valueKey string, value [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.SerializeBytes320(&_Vm.TransactOpts, objectKey, valueKey, value)
}

// SerializeInt is a paid mutator transaction binding the contract method 0x3f33db60.
//
// Solidity: function serializeInt(string objectKey, string valueKey, int256 value) returns(string json)
func (_Vm *VmTransactor) SerializeInt(opts *bind.TransactOpts, objectKey string, valueKey string, value *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "serializeInt", objectKey, valueKey, value)
}

// SerializeInt is a paid mutator transaction binding the contract method 0x3f33db60.
//
// Solidity: function serializeInt(string objectKey, string valueKey, int256 value) returns(string json)
func (_Vm *VmSession) SerializeInt(objectKey string, valueKey string, value *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.SerializeInt(&_Vm.TransactOpts, objectKey, valueKey, value)
}

// SerializeInt is a paid mutator transaction binding the contract method 0x3f33db60.
//
// Solidity: function serializeInt(string objectKey, string valueKey, int256 value) returns(string json)
func (_Vm *VmTransactorSession) SerializeInt(objectKey string, valueKey string, value *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.SerializeInt(&_Vm.TransactOpts, objectKey, valueKey, value)
}

// SerializeInt0 is a paid mutator transaction binding the contract method 0x7676e127.
//
// Solidity: function serializeInt(string objectKey, string valueKey, int256[] values) returns(string json)
func (_Vm *VmTransactor) SerializeInt0(opts *bind.TransactOpts, objectKey string, valueKey string, values []*big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "serializeInt0", objectKey, valueKey, values)
}

// SerializeInt0 is a paid mutator transaction binding the contract method 0x7676e127.
//
// Solidity: function serializeInt(string objectKey, string valueKey, int256[] values) returns(string json)
func (_Vm *VmSession) SerializeInt0(objectKey string, valueKey string, values []*big.Int) (*types.Transaction, error) {
	return _Vm.Contract.SerializeInt0(&_Vm.TransactOpts, objectKey, valueKey, values)
}

// SerializeInt0 is a paid mutator transaction binding the contract method 0x7676e127.
//
// Solidity: function serializeInt(string objectKey, string valueKey, int256[] values) returns(string json)
func (_Vm *VmTransactorSession) SerializeInt0(objectKey string, valueKey string, values []*big.Int) (*types.Transaction, error) {
	return _Vm.Contract.SerializeInt0(&_Vm.TransactOpts, objectKey, valueKey, values)
}

// SerializeJson is a paid mutator transaction binding the contract method 0x9b3358b0.
//
// Solidity: function serializeJson(string objectKey, string value) returns(string json)
func (_Vm *VmTransactor) SerializeJson(opts *bind.TransactOpts, objectKey string, value string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "serializeJson", objectKey, value)
}

// SerializeJson is a paid mutator transaction binding the contract method 0x9b3358b0.
//
// Solidity: function serializeJson(string objectKey, string value) returns(string json)
func (_Vm *VmSession) SerializeJson(objectKey string, value string) (*types.Transaction, error) {
	return _Vm.Contract.SerializeJson(&_Vm.TransactOpts, objectKey, value)
}

// SerializeJson is a paid mutator transaction binding the contract method 0x9b3358b0.
//
// Solidity: function serializeJson(string objectKey, string value) returns(string json)
func (_Vm *VmTransactorSession) SerializeJson(objectKey string, value string) (*types.Transaction, error) {
	return _Vm.Contract.SerializeJson(&_Vm.TransactOpts, objectKey, value)
}

// SerializeString is a paid mutator transaction binding the contract method 0x561cd6f3.
//
// Solidity: function serializeString(string objectKey, string valueKey, string[] values) returns(string json)
func (_Vm *VmTransactor) SerializeString(opts *bind.TransactOpts, objectKey string, valueKey string, values []string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "serializeString", objectKey, valueKey, values)
}

// SerializeString is a paid mutator transaction binding the contract method 0x561cd6f3.
//
// Solidity: function serializeString(string objectKey, string valueKey, string[] values) returns(string json)
func (_Vm *VmSession) SerializeString(objectKey string, valueKey string, values []string) (*types.Transaction, error) {
	return _Vm.Contract.SerializeString(&_Vm.TransactOpts, objectKey, valueKey, values)
}

// SerializeString is a paid mutator transaction binding the contract method 0x561cd6f3.
//
// Solidity: function serializeString(string objectKey, string valueKey, string[] values) returns(string json)
func (_Vm *VmTransactorSession) SerializeString(objectKey string, valueKey string, values []string) (*types.Transaction, error) {
	return _Vm.Contract.SerializeString(&_Vm.TransactOpts, objectKey, valueKey, values)
}

// SerializeString0 is a paid mutator transaction binding the contract method 0x88da6d35.
//
// Solidity: function serializeString(string objectKey, string valueKey, string value) returns(string json)
func (_Vm *VmTransactor) SerializeString0(opts *bind.TransactOpts, objectKey string, valueKey string, value string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "serializeString0", objectKey, valueKey, value)
}

// SerializeString0 is a paid mutator transaction binding the contract method 0x88da6d35.
//
// Solidity: function serializeString(string objectKey, string valueKey, string value) returns(string json)
func (_Vm *VmSession) SerializeString0(objectKey string, valueKey string, value string) (*types.Transaction, error) {
	return _Vm.Contract.SerializeString0(&_Vm.TransactOpts, objectKey, valueKey, value)
}

// SerializeString0 is a paid mutator transaction binding the contract method 0x88da6d35.
//
// Solidity: function serializeString(string objectKey, string valueKey, string value) returns(string json)
func (_Vm *VmTransactorSession) SerializeString0(objectKey string, valueKey string, value string) (*types.Transaction, error) {
	return _Vm.Contract.SerializeString0(&_Vm.TransactOpts, objectKey, valueKey, value)
}

// SerializeUint is a paid mutator transaction binding the contract method 0x129e9002.
//
// Solidity: function serializeUint(string objectKey, string valueKey, uint256 value) returns(string json)
func (_Vm *VmTransactor) SerializeUint(opts *bind.TransactOpts, objectKey string, valueKey string, value *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "serializeUint", objectKey, valueKey, value)
}

// SerializeUint is a paid mutator transaction binding the contract method 0x129e9002.
//
// Solidity: function serializeUint(string objectKey, string valueKey, uint256 value) returns(string json)
func (_Vm *VmSession) SerializeUint(objectKey string, valueKey string, value *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.SerializeUint(&_Vm.TransactOpts, objectKey, valueKey, value)
}

// SerializeUint is a paid mutator transaction binding the contract method 0x129e9002.
//
// Solidity: function serializeUint(string objectKey, string valueKey, uint256 value) returns(string json)
func (_Vm *VmTransactorSession) SerializeUint(objectKey string, valueKey string, value *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.SerializeUint(&_Vm.TransactOpts, objectKey, valueKey, value)
}

// SerializeUint0 is a paid mutator transaction binding the contract method 0xfee9a469.
//
// Solidity: function serializeUint(string objectKey, string valueKey, uint256[] values) returns(string json)
func (_Vm *VmTransactor) SerializeUint0(opts *bind.TransactOpts, objectKey string, valueKey string, values []*big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "serializeUint0", objectKey, valueKey, values)
}

// SerializeUint0 is a paid mutator transaction binding the contract method 0xfee9a469.
//
// Solidity: function serializeUint(string objectKey, string valueKey, uint256[] values) returns(string json)
func (_Vm *VmSession) SerializeUint0(objectKey string, valueKey string, values []*big.Int) (*types.Transaction, error) {
	return _Vm.Contract.SerializeUint0(&_Vm.TransactOpts, objectKey, valueKey, values)
}

// SerializeUint0 is a paid mutator transaction binding the contract method 0xfee9a469.
//
// Solidity: function serializeUint(string objectKey, string valueKey, uint256[] values) returns(string json)
func (_Vm *VmTransactorSession) SerializeUint0(objectKey string, valueKey string, values []*big.Int) (*types.Transaction, error) {
	return _Vm.Contract.SerializeUint0(&_Vm.TransactOpts, objectKey, valueKey, values)
}

// SerializeUintToHex is a paid mutator transaction binding the contract method 0xae5a2ae8.
//
// Solidity: function serializeUintToHex(string objectKey, string valueKey, uint256 value) returns(string json)
func (_Vm *VmTransactor) SerializeUintToHex(opts *bind.TransactOpts, objectKey string, valueKey string, value *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "serializeUintToHex", objectKey, valueKey, value)
}

// SerializeUintToHex is a paid mutator transaction binding the contract method 0xae5a2ae8.
//
// Solidity: function serializeUintToHex(string objectKey, string valueKey, uint256 value) returns(string json)
func (_Vm *VmSession) SerializeUintToHex(objectKey string, valueKey string, value *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.SerializeUintToHex(&_Vm.TransactOpts, objectKey, valueKey, value)
}

// SerializeUintToHex is a paid mutator transaction binding the contract method 0xae5a2ae8.
//
// Solidity: function serializeUintToHex(string objectKey, string valueKey, uint256 value) returns(string json)
func (_Vm *VmTransactorSession) SerializeUintToHex(objectKey string, valueKey string, value *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.SerializeUintToHex(&_Vm.TransactOpts, objectKey, valueKey, value)
}

// SetEnv is a paid mutator transaction binding the contract method 0x3d5923ee.
//
// Solidity: function setEnv(string name, string value) returns()
func (_Vm *VmTransactor) SetEnv(opts *bind.TransactOpts, name string, value string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "setEnv", name, value)
}

// SetEnv is a paid mutator transaction binding the contract method 0x3d5923ee.
//
// Solidity: function setEnv(string name, string value) returns()
func (_Vm *VmSession) SetEnv(name string, value string) (*types.Transaction, error) {
	return _Vm.Contract.SetEnv(&_Vm.TransactOpts, name, value)
}

// SetEnv is a paid mutator transaction binding the contract method 0x3d5923ee.
//
// Solidity: function setEnv(string name, string value) returns()
func (_Vm *VmTransactorSession) SetEnv(name string, value string) (*types.Transaction, error) {
	return _Vm.Contract.SetEnv(&_Vm.TransactOpts, name, value)
}

// SetNonce is a paid mutator transaction binding the contract method 0xf8e18b57.
//
// Solidity: function setNonce(address account, uint64 newNonce) returns()
func (_Vm *VmTransactor) SetNonce(opts *bind.TransactOpts, account common.Address, newNonce uint64) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "setNonce", account, newNonce)
}

// SetNonce is a paid mutator transaction binding the contract method 0xf8e18b57.
//
// Solidity: function setNonce(address account, uint64 newNonce) returns()
func (_Vm *VmSession) SetNonce(account common.Address, newNonce uint64) (*types.Transaction, error) {
	return _Vm.Contract.SetNonce(&_Vm.TransactOpts, account, newNonce)
}

// SetNonce is a paid mutator transaction binding the contract method 0xf8e18b57.
//
// Solidity: function setNonce(address account, uint64 newNonce) returns()
func (_Vm *VmTransactorSession) SetNonce(account common.Address, newNonce uint64) (*types.Transaction, error) {
	return _Vm.Contract.SetNonce(&_Vm.TransactOpts, account, newNonce)
}

// SetNonceUnsafe is a paid mutator transaction binding the contract method 0x9b67b21c.
//
// Solidity: function setNonceUnsafe(address account, uint64 newNonce) returns()
func (_Vm *VmTransactor) SetNonceUnsafe(opts *bind.TransactOpts, account common.Address, newNonce uint64) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "setNonceUnsafe", account, newNonce)
}

// SetNonceUnsafe is a paid mutator transaction binding the contract method 0x9b67b21c.
//
// Solidity: function setNonceUnsafe(address account, uint64 newNonce) returns()
func (_Vm *VmSession) SetNonceUnsafe(account common.Address, newNonce uint64) (*types.Transaction, error) {
	return _Vm.Contract.SetNonceUnsafe(&_Vm.TransactOpts, account, newNonce)
}

// SetNonceUnsafe is a paid mutator transaction binding the contract method 0x9b67b21c.
//
// Solidity: function setNonceUnsafe(address account, uint64 newNonce) returns()
func (_Vm *VmTransactorSession) SetNonceUnsafe(account common.Address, newNonce uint64) (*types.Transaction, error) {
	return _Vm.Contract.SetNonceUnsafe(&_Vm.TransactOpts, account, newNonce)
}

// Sign1 is a paid mutator transaction binding the contract method 0xb25c5a25.
//
// Solidity: function sign((address,uint256,uint256,uint256) wallet, bytes32 digest) returns(uint8 v, bytes32 r, bytes32 s)
func (_Vm *VmTransactor) Sign1(opts *bind.TransactOpts, wallet VmSafeWallet, digest [32]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "sign1", wallet, digest)
}

// Sign1 is a paid mutator transaction binding the contract method 0xb25c5a25.
//
// Solidity: function sign((address,uint256,uint256,uint256) wallet, bytes32 digest) returns(uint8 v, bytes32 r, bytes32 s)
func (_Vm *VmSession) Sign1(wallet VmSafeWallet, digest [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.Sign1(&_Vm.TransactOpts, wallet, digest)
}

// Sign1 is a paid mutator transaction binding the contract method 0xb25c5a25.
//
// Solidity: function sign((address,uint256,uint256,uint256) wallet, bytes32 digest) returns(uint8 v, bytes32 r, bytes32 s)
func (_Vm *VmTransactorSession) Sign1(wallet VmSafeWallet, digest [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.Sign1(&_Vm.TransactOpts, wallet, digest)
}

// Skip is a paid mutator transaction binding the contract method 0xdd82d13e.
//
// Solidity: function skip(bool skipTest) returns()
func (_Vm *VmTransactor) Skip(opts *bind.TransactOpts, skipTest bool) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "skip", skipTest)
}

// Skip is a paid mutator transaction binding the contract method 0xdd82d13e.
//
// Solidity: function skip(bool skipTest) returns()
func (_Vm *VmSession) Skip(skipTest bool) (*types.Transaction, error) {
	return _Vm.Contract.Skip(&_Vm.TransactOpts, skipTest)
}

// Skip is a paid mutator transaction binding the contract method 0xdd82d13e.
//
// Solidity: function skip(bool skipTest) returns()
func (_Vm *VmTransactorSession) Skip(skipTest bool) (*types.Transaction, error) {
	return _Vm.Contract.Skip(&_Vm.TransactOpts, skipTest)
}

// Sleep is a paid mutator transaction binding the contract method 0xfa9d8713.
//
// Solidity: function sleep(uint256 duration) returns()
func (_Vm *VmTransactor) Sleep(opts *bind.TransactOpts, duration *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "sleep", duration)
}

// Sleep is a paid mutator transaction binding the contract method 0xfa9d8713.
//
// Solidity: function sleep(uint256 duration) returns()
func (_Vm *VmSession) Sleep(duration *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.Sleep(&_Vm.TransactOpts, duration)
}

// Sleep is a paid mutator transaction binding the contract method 0xfa9d8713.
//
// Solidity: function sleep(uint256 duration) returns()
func (_Vm *VmTransactorSession) Sleep(duration *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.Sleep(&_Vm.TransactOpts, duration)
}

// Snapshot is a paid mutator transaction binding the contract method 0x9711715a.
//
// Solidity: function snapshot() returns(uint256 snapshotId)
func (_Vm *VmTransactor) Snapshot(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "snapshot")
}

// Snapshot is a paid mutator transaction binding the contract method 0x9711715a.
//
// Solidity: function snapshot() returns(uint256 snapshotId)
func (_Vm *VmSession) Snapshot() (*types.Transaction, error) {
	return _Vm.Contract.Snapshot(&_Vm.TransactOpts)
}

// Snapshot is a paid mutator transaction binding the contract method 0x9711715a.
//
// Solidity: function snapshot() returns(uint256 snapshotId)
func (_Vm *VmTransactorSession) Snapshot() (*types.Transaction, error) {
	return _Vm.Contract.Snapshot(&_Vm.TransactOpts)
}

// StartBroadcast is a paid mutator transaction binding the contract method 0x7fb5297f.
//
// Solidity: function startBroadcast() returns()
func (_Vm *VmTransactor) StartBroadcast(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "startBroadcast")
}

// StartBroadcast is a paid mutator transaction binding the contract method 0x7fb5297f.
//
// Solidity: function startBroadcast() returns()
func (_Vm *VmSession) StartBroadcast() (*types.Transaction, error) {
	return _Vm.Contract.StartBroadcast(&_Vm.TransactOpts)
}

// StartBroadcast is a paid mutator transaction binding the contract method 0x7fb5297f.
//
// Solidity: function startBroadcast() returns()
func (_Vm *VmTransactorSession) StartBroadcast() (*types.Transaction, error) {
	return _Vm.Contract.StartBroadcast(&_Vm.TransactOpts)
}

// StartBroadcast0 is a paid mutator transaction binding the contract method 0x7fec2a8d.
//
// Solidity: function startBroadcast(address signer) returns()
func (_Vm *VmTransactor) StartBroadcast0(opts *bind.TransactOpts, signer common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "startBroadcast0", signer)
}

// StartBroadcast0 is a paid mutator transaction binding the contract method 0x7fec2a8d.
//
// Solidity: function startBroadcast(address signer) returns()
func (_Vm *VmSession) StartBroadcast0(signer common.Address) (*types.Transaction, error) {
	return _Vm.Contract.StartBroadcast0(&_Vm.TransactOpts, signer)
}

// StartBroadcast0 is a paid mutator transaction binding the contract method 0x7fec2a8d.
//
// Solidity: function startBroadcast(address signer) returns()
func (_Vm *VmTransactorSession) StartBroadcast0(signer common.Address) (*types.Transaction, error) {
	return _Vm.Contract.StartBroadcast0(&_Vm.TransactOpts, signer)
}

// StartBroadcast1 is a paid mutator transaction binding the contract method 0xce817d47.
//
// Solidity: function startBroadcast(uint256 privateKey) returns()
func (_Vm *VmTransactor) StartBroadcast1(opts *bind.TransactOpts, privateKey *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "startBroadcast1", privateKey)
}

// StartBroadcast1 is a paid mutator transaction binding the contract method 0xce817d47.
//
// Solidity: function startBroadcast(uint256 privateKey) returns()
func (_Vm *VmSession) StartBroadcast1(privateKey *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.StartBroadcast1(&_Vm.TransactOpts, privateKey)
}

// StartBroadcast1 is a paid mutator transaction binding the contract method 0xce817d47.
//
// Solidity: function startBroadcast(uint256 privateKey) returns()
func (_Vm *VmTransactorSession) StartBroadcast1(privateKey *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.StartBroadcast1(&_Vm.TransactOpts, privateKey)
}

// StartMappingRecording is a paid mutator transaction binding the contract method 0x3e9705c0.
//
// Solidity: function startMappingRecording() returns()
func (_Vm *VmTransactor) StartMappingRecording(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "startMappingRecording")
}

// StartMappingRecording is a paid mutator transaction binding the contract method 0x3e9705c0.
//
// Solidity: function startMappingRecording() returns()
func (_Vm *VmSession) StartMappingRecording() (*types.Transaction, error) {
	return _Vm.Contract.StartMappingRecording(&_Vm.TransactOpts)
}

// StartMappingRecording is a paid mutator transaction binding the contract method 0x3e9705c0.
//
// Solidity: function startMappingRecording() returns()
func (_Vm *VmTransactorSession) StartMappingRecording() (*types.Transaction, error) {
	return _Vm.Contract.StartMappingRecording(&_Vm.TransactOpts)
}

// StartPrank is a paid mutator transaction binding the contract method 0x06447d56.
//
// Solidity: function startPrank(address msgSender) returns()
func (_Vm *VmTransactor) StartPrank(opts *bind.TransactOpts, msgSender common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "startPrank", msgSender)
}

// StartPrank is a paid mutator transaction binding the contract method 0x06447d56.
//
// Solidity: function startPrank(address msgSender) returns()
func (_Vm *VmSession) StartPrank(msgSender common.Address) (*types.Transaction, error) {
	return _Vm.Contract.StartPrank(&_Vm.TransactOpts, msgSender)
}

// StartPrank is a paid mutator transaction binding the contract method 0x06447d56.
//
// Solidity: function startPrank(address msgSender) returns()
func (_Vm *VmTransactorSession) StartPrank(msgSender common.Address) (*types.Transaction, error) {
	return _Vm.Contract.StartPrank(&_Vm.TransactOpts, msgSender)
}

// StartPrank0 is a paid mutator transaction binding the contract method 0x45b56078.
//
// Solidity: function startPrank(address msgSender, address txOrigin) returns()
func (_Vm *VmTransactor) StartPrank0(opts *bind.TransactOpts, msgSender common.Address, txOrigin common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "startPrank0", msgSender, txOrigin)
}

// StartPrank0 is a paid mutator transaction binding the contract method 0x45b56078.
//
// Solidity: function startPrank(address msgSender, address txOrigin) returns()
func (_Vm *VmSession) StartPrank0(msgSender common.Address, txOrigin common.Address) (*types.Transaction, error) {
	return _Vm.Contract.StartPrank0(&_Vm.TransactOpts, msgSender, txOrigin)
}

// StartPrank0 is a paid mutator transaction binding the contract method 0x45b56078.
//
// Solidity: function startPrank(address msgSender, address txOrigin) returns()
func (_Vm *VmTransactorSession) StartPrank0(msgSender common.Address, txOrigin common.Address) (*types.Transaction, error) {
	return _Vm.Contract.StartPrank0(&_Vm.TransactOpts, msgSender, txOrigin)
}

// StartStateDiffRecording is a paid mutator transaction binding the contract method 0xcf22e3c9.
//
// Solidity: function startStateDiffRecording() returns()
func (_Vm *VmTransactor) StartStateDiffRecording(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "startStateDiffRecording")
}

// StartStateDiffRecording is a paid mutator transaction binding the contract method 0xcf22e3c9.
//
// Solidity: function startStateDiffRecording() returns()
func (_Vm *VmSession) StartStateDiffRecording() (*types.Transaction, error) {
	return _Vm.Contract.StartStateDiffRecording(&_Vm.TransactOpts)
}

// StartStateDiffRecording is a paid mutator transaction binding the contract method 0xcf22e3c9.
//
// Solidity: function startStateDiffRecording() returns()
func (_Vm *VmTransactorSession) StartStateDiffRecording() (*types.Transaction, error) {
	return _Vm.Contract.StartStateDiffRecording(&_Vm.TransactOpts)
}

// StopAndReturnStateDiff is a paid mutator transaction binding the contract method 0xaa5cf90e.
//
// Solidity: function stopAndReturnStateDiff() returns(((uint256,uint256),uint8,address,address,bool,uint256,uint256,bytes,uint256,bytes,bool,(address,bytes32,bool,bytes32,bytes32,bool)[],uint64)[] accountAccesses)
func (_Vm *VmTransactor) StopAndReturnStateDiff(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "stopAndReturnStateDiff")
}

// StopAndReturnStateDiff is a paid mutator transaction binding the contract method 0xaa5cf90e.
//
// Solidity: function stopAndReturnStateDiff() returns(((uint256,uint256),uint8,address,address,bool,uint256,uint256,bytes,uint256,bytes,bool,(address,bytes32,bool,bytes32,bytes32,bool)[],uint64)[] accountAccesses)
func (_Vm *VmSession) StopAndReturnStateDiff() (*types.Transaction, error) {
	return _Vm.Contract.StopAndReturnStateDiff(&_Vm.TransactOpts)
}

// StopAndReturnStateDiff is a paid mutator transaction binding the contract method 0xaa5cf90e.
//
// Solidity: function stopAndReturnStateDiff() returns(((uint256,uint256),uint8,address,address,bool,uint256,uint256,bytes,uint256,bytes,bool,(address,bytes32,bool,bytes32,bytes32,bool)[],uint64)[] accountAccesses)
func (_Vm *VmTransactorSession) StopAndReturnStateDiff() (*types.Transaction, error) {
	return _Vm.Contract.StopAndReturnStateDiff(&_Vm.TransactOpts)
}

// StopBroadcast is a paid mutator transaction binding the contract method 0x76eadd36.
//
// Solidity: function stopBroadcast() returns()
func (_Vm *VmTransactor) StopBroadcast(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "stopBroadcast")
}

// StopBroadcast is a paid mutator transaction binding the contract method 0x76eadd36.
//
// Solidity: function stopBroadcast() returns()
func (_Vm *VmSession) StopBroadcast() (*types.Transaction, error) {
	return _Vm.Contract.StopBroadcast(&_Vm.TransactOpts)
}

// StopBroadcast is a paid mutator transaction binding the contract method 0x76eadd36.
//
// Solidity: function stopBroadcast() returns()
func (_Vm *VmTransactorSession) StopBroadcast() (*types.Transaction, error) {
	return _Vm.Contract.StopBroadcast(&_Vm.TransactOpts)
}

// StopExpectSafeMemory is a paid mutator transaction binding the contract method 0x0956441b.
//
// Solidity: function stopExpectSafeMemory() returns()
func (_Vm *VmTransactor) StopExpectSafeMemory(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "stopExpectSafeMemory")
}

// StopExpectSafeMemory is a paid mutator transaction binding the contract method 0x0956441b.
//
// Solidity: function stopExpectSafeMemory() returns()
func (_Vm *VmSession) StopExpectSafeMemory() (*types.Transaction, error) {
	return _Vm.Contract.StopExpectSafeMemory(&_Vm.TransactOpts)
}

// StopExpectSafeMemory is a paid mutator transaction binding the contract method 0x0956441b.
//
// Solidity: function stopExpectSafeMemory() returns()
func (_Vm *VmTransactorSession) StopExpectSafeMemory() (*types.Transaction, error) {
	return _Vm.Contract.StopExpectSafeMemory(&_Vm.TransactOpts)
}

// StopMappingRecording is a paid mutator transaction binding the contract method 0x0d4aae9b.
//
// Solidity: function stopMappingRecording() returns()
func (_Vm *VmTransactor) StopMappingRecording(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "stopMappingRecording")
}

// StopMappingRecording is a paid mutator transaction binding the contract method 0x0d4aae9b.
//
// Solidity: function stopMappingRecording() returns()
func (_Vm *VmSession) StopMappingRecording() (*types.Transaction, error) {
	return _Vm.Contract.StopMappingRecording(&_Vm.TransactOpts)
}

// StopMappingRecording is a paid mutator transaction binding the contract method 0x0d4aae9b.
//
// Solidity: function stopMappingRecording() returns()
func (_Vm *VmTransactorSession) StopMappingRecording() (*types.Transaction, error) {
	return _Vm.Contract.StopMappingRecording(&_Vm.TransactOpts)
}

// StopPrank is a paid mutator transaction binding the contract method 0x90c5013b.
//
// Solidity: function stopPrank() returns()
func (_Vm *VmTransactor) StopPrank(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "stopPrank")
}

// StopPrank is a paid mutator transaction binding the contract method 0x90c5013b.
//
// Solidity: function stopPrank() returns()
func (_Vm *VmSession) StopPrank() (*types.Transaction, error) {
	return _Vm.Contract.StopPrank(&_Vm.TransactOpts)
}

// StopPrank is a paid mutator transaction binding the contract method 0x90c5013b.
//
// Solidity: function stopPrank() returns()
func (_Vm *VmTransactorSession) StopPrank() (*types.Transaction, error) {
	return _Vm.Contract.StopPrank(&_Vm.TransactOpts)
}

// Store is a paid mutator transaction binding the contract method 0x70ca10bb.
//
// Solidity: function store(address target, bytes32 slot, bytes32 value) returns()
func (_Vm *VmTransactor) Store(opts *bind.TransactOpts, target common.Address, slot [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "store", target, slot, value)
}

// Store is a paid mutator transaction binding the contract method 0x70ca10bb.
//
// Solidity: function store(address target, bytes32 slot, bytes32 value) returns()
func (_Vm *VmSession) Store(target common.Address, slot [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.Store(&_Vm.TransactOpts, target, slot, value)
}

// Store is a paid mutator transaction binding the contract method 0x70ca10bb.
//
// Solidity: function store(address target, bytes32 slot, bytes32 value) returns()
func (_Vm *VmTransactorSession) Store(target common.Address, slot [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.Store(&_Vm.TransactOpts, target, slot, value)
}

// Transact is a paid mutator transaction binding the contract method 0x4d8abc4b.
//
// Solidity: function transact(uint256 forkId, bytes32 txHash) returns()
func (_Vm *VmTransactor) Transact(opts *bind.TransactOpts, forkId *big.Int, txHash [32]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "transact", forkId, txHash)
}

// Transact is a paid mutator transaction binding the contract method 0x4d8abc4b.
//
// Solidity: function transact(uint256 forkId, bytes32 txHash) returns()
func (_Vm *VmSession) Transact(forkId *big.Int, txHash [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.Transact(&_Vm.TransactOpts, forkId, txHash)
}

// Transact is a paid mutator transaction binding the contract method 0x4d8abc4b.
//
// Solidity: function transact(uint256 forkId, bytes32 txHash) returns()
func (_Vm *VmTransactorSession) Transact(forkId *big.Int, txHash [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.Transact(&_Vm.TransactOpts, forkId, txHash)
}

// Transact0 is a paid mutator transaction binding the contract method 0xbe646da1.
//
// Solidity: function transact(bytes32 txHash) returns()
func (_Vm *VmTransactor) Transact0(opts *bind.TransactOpts, txHash [32]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "transact0", txHash)
}

// Transact0 is a paid mutator transaction binding the contract method 0xbe646da1.
//
// Solidity: function transact(bytes32 txHash) returns()
func (_Vm *VmSession) Transact0(txHash [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.Transact0(&_Vm.TransactOpts, txHash)
}

// Transact0 is a paid mutator transaction binding the contract method 0xbe646da1.
//
// Solidity: function transact(bytes32 txHash) returns()
func (_Vm *VmTransactorSession) Transact0(txHash [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.Transact0(&_Vm.TransactOpts, txHash)
}

// TryFfi is a paid mutator transaction binding the contract method 0xf45c1ce7.
//
// Solidity: function tryFfi(string[] commandInput) returns((int32,bytes,bytes) result)
func (_Vm *VmTransactor) TryFfi(opts *bind.TransactOpts, commandInput []string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "tryFfi", commandInput)
}

// TryFfi is a paid mutator transaction binding the contract method 0xf45c1ce7.
//
// Solidity: function tryFfi(string[] commandInput) returns((int32,bytes,bytes) result)
func (_Vm *VmSession) TryFfi(commandInput []string) (*types.Transaction, error) {
	return _Vm.Contract.TryFfi(&_Vm.TransactOpts, commandInput)
}

// TryFfi is a paid mutator transaction binding the contract method 0xf45c1ce7.
//
// Solidity: function tryFfi(string[] commandInput) returns((int32,bytes,bytes) result)
func (_Vm *VmTransactorSession) TryFfi(commandInput []string) (*types.Transaction, error) {
	return _Vm.Contract.TryFfi(&_Vm.TransactOpts, commandInput)
}

// TxGasPrice is a paid mutator transaction binding the contract method 0x48f50c0f.
//
// Solidity: function txGasPrice(uint256 newGasPrice) returns()
func (_Vm *VmTransactor) TxGasPrice(opts *bind.TransactOpts, newGasPrice *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "txGasPrice", newGasPrice)
}

// TxGasPrice is a paid mutator transaction binding the contract method 0x48f50c0f.
//
// Solidity: function txGasPrice(uint256 newGasPrice) returns()
func (_Vm *VmSession) TxGasPrice(newGasPrice *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.TxGasPrice(&_Vm.TransactOpts, newGasPrice)
}

// TxGasPrice is a paid mutator transaction binding the contract method 0x48f50c0f.
//
// Solidity: function txGasPrice(uint256 newGasPrice) returns()
func (_Vm *VmTransactorSession) TxGasPrice(newGasPrice *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.TxGasPrice(&_Vm.TransactOpts, newGasPrice)
}

// UnixTime is a paid mutator transaction binding the contract method 0x625387dc.
//
// Solidity: function unixTime() returns(uint256 milliseconds)
func (_Vm *VmTransactor) UnixTime(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "unixTime")
}

// UnixTime is a paid mutator transaction binding the contract method 0x625387dc.
//
// Solidity: function unixTime() returns(uint256 milliseconds)
func (_Vm *VmSession) UnixTime() (*types.Transaction, error) {
	return _Vm.Contract.UnixTime(&_Vm.TransactOpts)
}

// UnixTime is a paid mutator transaction binding the contract method 0x625387dc.
//
// Solidity: function unixTime() returns(uint256 milliseconds)
func (_Vm *VmTransactorSession) UnixTime() (*types.Transaction, error) {
	return _Vm.Contract.UnixTime(&_Vm.TransactOpts)
}

// Warp is a paid mutator transaction binding the contract method 0xe5d6bf02.
//
// Solidity: function warp(uint256 newTimestamp) returns()
func (_Vm *VmTransactor) Warp(opts *bind.TransactOpts, newTimestamp *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "warp", newTimestamp)
}

// Warp is a paid mutator transaction binding the contract method 0xe5d6bf02.
//
// Solidity: function warp(uint256 newTimestamp) returns()
func (_Vm *VmSession) Warp(newTimestamp *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.Warp(&_Vm.TransactOpts, newTimestamp)
}

// Warp is a paid mutator transaction binding the contract method 0xe5d6bf02.
//
// Solidity: function warp(uint256 newTimestamp) returns()
func (_Vm *VmTransactorSession) Warp(newTimestamp *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.Warp(&_Vm.TransactOpts, newTimestamp)
}

// WriteFile is a paid mutator transaction binding the contract method 0x897e0a97.
//
// Solidity: function writeFile(string path, string data) returns()
func (_Vm *VmTransactor) WriteFile(opts *bind.TransactOpts, path string, data string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "writeFile", path, data)
}

// WriteFile is a paid mutator transaction binding the contract method 0x897e0a97.
//
// Solidity: function writeFile(string path, string data) returns()
func (_Vm *VmSession) WriteFile(path string, data string) (*types.Transaction, error) {
	return _Vm.Contract.WriteFile(&_Vm.TransactOpts, path, data)
}

// WriteFile is a paid mutator transaction binding the contract method 0x897e0a97.
//
// Solidity: function writeFile(string path, string data) returns()
func (_Vm *VmTransactorSession) WriteFile(path string, data string) (*types.Transaction, error) {
	return _Vm.Contract.WriteFile(&_Vm.TransactOpts, path, data)
}

// WriteFileBinary is a paid mutator transaction binding the contract method 0x1f21fc80.
//
// Solidity: function writeFileBinary(string path, bytes data) returns()
func (_Vm *VmTransactor) WriteFileBinary(opts *bind.TransactOpts, path string, data []byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "writeFileBinary", path, data)
}

// WriteFileBinary is a paid mutator transaction binding the contract method 0x1f21fc80.
//
// Solidity: function writeFileBinary(string path, bytes data) returns()
func (_Vm *VmSession) WriteFileBinary(path string, data []byte) (*types.Transaction, error) {
	return _Vm.Contract.WriteFileBinary(&_Vm.TransactOpts, path, data)
}

// WriteFileBinary is a paid mutator transaction binding the contract method 0x1f21fc80.
//
// Solidity: function writeFileBinary(string path, bytes data) returns()
func (_Vm *VmTransactorSession) WriteFileBinary(path string, data []byte) (*types.Transaction, error) {
	return _Vm.Contract.WriteFileBinary(&_Vm.TransactOpts, path, data)
}

// WriteJson is a paid mutator transaction binding the contract method 0x35d6ad46.
//
// Solidity: function writeJson(string json, string path, string valueKey) returns()
func (_Vm *VmTransactor) WriteJson(opts *bind.TransactOpts, json string, path string, valueKey string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "writeJson", json, path, valueKey)
}

// WriteJson is a paid mutator transaction binding the contract method 0x35d6ad46.
//
// Solidity: function writeJson(string json, string path, string valueKey) returns()
func (_Vm *VmSession) WriteJson(json string, path string, valueKey string) (*types.Transaction, error) {
	return _Vm.Contract.WriteJson(&_Vm.TransactOpts, json, path, valueKey)
}

// WriteJson is a paid mutator transaction binding the contract method 0x35d6ad46.
//
// Solidity: function writeJson(string json, string path, string valueKey) returns()
func (_Vm *VmTransactorSession) WriteJson(json string, path string, valueKey string) (*types.Transaction, error) {
	return _Vm.Contract.WriteJson(&_Vm.TransactOpts, json, path, valueKey)
}

// WriteJson0 is a paid mutator transaction binding the contract method 0xe23cd19f.
//
// Solidity: function writeJson(string json, string path) returns()
func (_Vm *VmTransactor) WriteJson0(opts *bind.TransactOpts, json string, path string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "writeJson0", json, path)
}

// WriteJson0 is a paid mutator transaction binding the contract method 0xe23cd19f.
//
// Solidity: function writeJson(string json, string path) returns()
func (_Vm *VmSession) WriteJson0(json string, path string) (*types.Transaction, error) {
	return _Vm.Contract.WriteJson0(&_Vm.TransactOpts, json, path)
}

// WriteJson0 is a paid mutator transaction binding the contract method 0xe23cd19f.
//
// Solidity: function writeJson(string json, string path) returns()
func (_Vm *VmTransactorSession) WriteJson0(json string, path string) (*types.Transaction, error) {
	return _Vm.Contract.WriteJson0(&_Vm.TransactOpts, json, path)
}

// WriteLine is a paid mutator transaction binding the contract method 0x619d897f.
//
// Solidity: function writeLine(string path, string data) returns()
func (_Vm *VmTransactor) WriteLine(opts *bind.TransactOpts, path string, data string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "writeLine", path, data)
}

// WriteLine is a paid mutator transaction binding the contract method 0x619d897f.
//
// Solidity: function writeLine(string path, string data) returns()
func (_Vm *VmSession) WriteLine(path string, data string) (*types.Transaction, error) {
	return _Vm.Contract.WriteLine(&_Vm.TransactOpts, path, data)
}

// WriteLine is a paid mutator transaction binding the contract method 0x619d897f.
//
// Solidity: function writeLine(string path, string data) returns()
func (_Vm *VmTransactorSession) WriteLine(path string, data string) (*types.Transaction, error) {
	return _Vm.Contract.WriteLine(&_Vm.TransactOpts, path, data)
}

// WriteToml is a paid mutator transaction binding the contract method 0x51ac6a33.
//
// Solidity: function writeToml(string json, string path, string valueKey) returns()
func (_Vm *VmTransactor) WriteToml(opts *bind.TransactOpts, json string, path string, valueKey string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "writeToml", json, path, valueKey)
}

// WriteToml is a paid mutator transaction binding the contract method 0x51ac6a33.
//
// Solidity: function writeToml(string json, string path, string valueKey) returns()
func (_Vm *VmSession) WriteToml(json string, path string, valueKey string) (*types.Transaction, error) {
	return _Vm.Contract.WriteToml(&_Vm.TransactOpts, json, path, valueKey)
}

// WriteToml is a paid mutator transaction binding the contract method 0x51ac6a33.
//
// Solidity: function writeToml(string json, string path, string valueKey) returns()
func (_Vm *VmTransactorSession) WriteToml(json string, path string, valueKey string) (*types.Transaction, error) {
	return _Vm.Contract.WriteToml(&_Vm.TransactOpts, json, path, valueKey)
}

// WriteToml0 is a paid mutator transaction binding the contract method 0xc0865ba7.
//
// Solidity: function writeToml(string json, string path) returns()
func (_Vm *VmTransactor) WriteToml0(opts *bind.TransactOpts, json string, path string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "writeToml0", json, path)
}

// WriteToml0 is a paid mutator transaction binding the contract method 0xc0865ba7.
//
// Solidity: function writeToml(string json, string path) returns()
func (_Vm *VmSession) WriteToml0(json string, path string) (*types.Transaction, error) {
	return _Vm.Contract.WriteToml0(&_Vm.TransactOpts, json, path)
}

// WriteToml0 is a paid mutator transaction binding the contract method 0xc0865ba7.
//
// Solidity: function writeToml(string json, string path) returns()
func (_Vm *VmTransactorSession) WriteToml0(json string, path string) (*types.Transaction, error) {
	return _Vm.Contract.WriteToml0(&_Vm.TransactOpts, json, path)
}
