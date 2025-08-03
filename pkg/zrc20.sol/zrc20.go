// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zrc20

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

// ZRC20MetaData contains all meta data concerning the ZRC20 contract.
var ZRC20MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"name_\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol_\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals_\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"chainid_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"coinType_\",\"type\":\"uint8\",\"internalType\":\"enumCoinType\"},{\"name\":\"gasLimit_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"systemContractAddress_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gatewayAddress_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"CHAIN_ID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"COIN_TYPE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumCoinType\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"FUNGIBLE_MODULE_ADDRESS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GAS_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_BATCH_SIZE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PROTOCOL_FLAT_FEE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SYSTEM_CONTRACT_ADDRESS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"batchWithdraw\",\"inputs\":[{\"name\":\"recipients\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"amounts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"gatewayAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setName\",\"inputs\":[{\"name\":\"newName\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSymbol\",\"inputs\":[{\"name\":\"newSymbol\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateGasLimit\",\"inputs\":[{\"name\":\"gasLimit_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateGatewayAddress\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateProtocolFlatFee\",\"inputs\":[{\"name\":\"protocolFlatFee_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateSystemContractAddress\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"to\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawGasFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawGasFeeWithGasLimit\",\"inputs\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BatchWithdrawal\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipients\",\"type\":\"bytes[]\",\"indexed\":false,\"internalType\":\"bytes[]\"},{\"name\":\"values\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"gasFeePerTx\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"from\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedGasLimit\",\"inputs\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedGateway\",\"inputs\":[{\"name\":\"gateway\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedProtocolFlatFee\",\"inputs\":[{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedSystemContract\",\"inputs\":[{\"name\":\"systemContract\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawal\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BatchWithdrawSizeExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotFungibleModule\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyArray\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GasFeeTransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LowAllowance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LowBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroGasCoin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroGasPrice\",\"inputs\":[]}]",
	Bin: "0x60c060405234801561001057600080fd5b5060405161257238038061257283398101604081905261002f916101f0565b6001600160a01b038216158061004c57506001600160a01b038116155b1561006a5760405163d92e233d60e01b815260040160405180910390fd5b60066100768982610342565b5060076100838882610342565b506008805460ff191660ff881617905560808590528360028111156100aa576100aa610400565b60a08160028111156100be576100be610400565b905250600192909255600080546001600160a01b039283166001600160a01b0319909116179055600880549190921661010002610100600160a81b0319909116179055506104169350505050565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261013357600080fd5b81516001600160401b0381111561014c5761014c61010c565b604051601f8201601f19908116603f011681016001600160401b038111828210171561017a5761017a61010c565b60405281815283820160200185101561019257600080fd5b60005b828110156101b157602081860181015183830182015201610195565b506000918101602001919091529392505050565b8051600381106101d457600080fd5b919050565b80516001600160a01b03811681146101d457600080fd5b600080600080600080600080610100898b03121561020d57600080fd5b88516001600160401b0381111561022357600080fd5b61022f8b828c01610122565b60208b015190995090506001600160401b0381111561024d57600080fd5b6102598b828c01610122565b975050604089015160ff8116811461027057600080fd5b60608a0151909650945061028660808a016101c5565b60a08a0151909450925061029c60c08a016101d9565b91506102aa60e08a016101d9565b90509295985092959890939650565b600181811c908216806102cd57607f821691505b6020821081036102ed57634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561033d57806000526020600020601f840160051c8101602085101561031a5750805b601f840160051c820191505b8181101561033a5760008155600101610326565b50505b505050565b81516001600160401b0381111561035b5761035b61010c565b61036f8161036984546102b9565b846102f3565b6020601f8211600181146103a3576000831561038b5750848201515b600019600385901b1c1916600184901b17845561033a565b600084815260208120601f198516915b828110156103d357878501518255602094850194600190920191016103b3565b50848210156103f15786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b634e487b7160e01b600052602160045260246000fd5b60805160a05161211b610457600039600061036d01526000818161031901528181610e5d01528181610f630152818161117f0152611285015261211b6000f3fe608060405234801561001057600080fd5b50600436106101cf5760003560e01c806395d89b4111610104578063ccc77599116100a2578063eddeb12311610071578063eddeb12314610492578063f2441b32146104a5578063f687d12a146104c5578063fc5fecd5146104d857600080fd5b8063ccc77599146103fd578063cfdbf25414610410578063d9eeebed14610418578063dd62ed3e1461044c57600080fd5b8063b84c8246116100de578063b84c8246146103af578063c47f0027146103c4578063c7012626146103d7578063c835d7cc146103ea57600080fd5b806395d89b4114610360578063a3413d0314610368578063a9059cbb1461039c57600080fd5b806342966c681161017157806370a082311161014b57806370a08231146102cb57806373486a781461030157806385e1f4d0146103145780638b851b951461033b57600080fd5b806342966c681461029c57806347e7ef24146102af5780634d8943bb146102c257600080fd5b806318160ddd116101ad57806318160ddd1461022c57806323b872dd14610234578063313ce567146102475780633ce4a5bc1461025c57600080fd5b806306fdde03146101d4578063091d2788146101f2578063095ea7b314610209575b600080fd5b6101dc6104eb565b6040516101e991906118c9565b60405180910390f35b6101fb60015481565b6040519081526020016101e9565b61021c610217366004611908565b61057d565b60405190151581526020016101e9565b6005546101fb565b61021c610242366004611934565b610594565b60085460405160ff90911681526020016101e9565b61027773735b14bb79463307aacbed86daf3322b1e6226ab81565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101e9565b61021c6102aa366004611975565b61062b565b61021c6102bd366004611908565b61063f565b6101fb60025481565b6101fb6102d936600461198e565b73ffffffffffffffffffffffffffffffffffffffff1660009081526003602052604090205490565b61021c61030f3660046119f7565b610798565b6101fb7f000000000000000000000000000000000000000000000000000000000000000081565b60085461027790610100900473ffffffffffffffffffffffffffffffffffffffff1681565b6101dc6109e8565b61038f7f000000000000000000000000000000000000000000000000000000000000000081565b6040516101e99190611a68565b61021c6103aa366004611908565b6109f7565b6103c26103bd366004611b70565b610a04565b005b6103c26103d2366004611b70565b610a61565b61021c6103e5366004611bc1565b610aba565b6103c26103f836600461198e565b610c09565b6103c261040b36600461198e565b610d1d565b6101fb608081565b610420610e31565b6040805173ffffffffffffffffffffffffffffffffffffffff90931683526020830191909152016101e9565b6101fb61045a366004611c1a565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260046020908152604080832093909416825291909152205490565b6103c26104a0366004611975565b61104f565b6000546102779073ffffffffffffffffffffffffffffffffffffffff1681565b6103c26104d3366004611975565b6110d1565b6104206104e6366004611975565b611153565b6060600680546104fa90611c53565b80601f016020809104026020016040519081016040528092919081815260200182805461052690611c53565b80156105735780601f1061054857610100808354040283529160200191610573565b820191906000526020600020905b81548152906001019060200180831161055657829003601f168201915b5050505050905090565b600061058a33848461136f565b5060015b92915050565b60006105a1848484611478565b73ffffffffffffffffffffffffffffffffffffffff841660009081526004602090815260408083203384529091529020548281101561060c576040517f10bad14700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610620853361061b8685611cd5565b61136f565b506001949350505050565b60006106373383611633565b506001919050565b60003373735b14bb79463307aacbed86daf3322b1e6226ab1480159061067d575060005473ffffffffffffffffffffffffffffffffffffffff163314155b80156106a65750600854610100900473ffffffffffffffffffffffffffffffffffffffff163314155b156106dd576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6106e78383611775565b6040517f735b14bb79463307aacbed86daf3322b1e6226ab000000000000000000000000602082015273ffffffffffffffffffffffffffffffffffffffff8416907f67fc7bdaed5b0ec550d8706b87d60568ab70c6b781263c70101d54cd1564aab390603401604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815290829052610787918690611ce8565b60405180910390a250600192915050565b6000838082036107d4576040517f521299a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80831461080d576040517fa24a13a600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6080811115610848576040517fa9cf382700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080610853610e31565b909250905060006108648483611d0a565b6040517f23b872dd00000000000000000000000000000000000000000000000000000000815233600482015273735b14bb79463307aacbed86daf3322b1e6226ab60248201526044810182905290915073ffffffffffffffffffffffffffffffffffffffff8416906323b872dd906064016020604051808303816000875af11580156108f4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109189190611d21565b61094e576040517f0a7cd6d600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000805b858110156109885788888281811061096c5761096c611d43565b905060200201358261097e9190611d72565b9150600101610952565b506109933382611633565b60025460405133917f2512245b924c3fcf7f4aed2a20c0e35ee10a0f15114842e5e46732128d8cab3a916109d0918e918e918e918e918b91611e19565b60405180910390a25060019998505050505050505050565b6060600780546104fa90611c53565b600061058a338484611478565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610a51576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6007610a5d8282611f67565b5050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610aae576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6006610a5d8282611f67565b6000806000610ac7610e31565b6040517f23b872dd00000000000000000000000000000000000000000000000000000000815233600482015273735b14bb79463307aacbed86daf3322b1e6226ab602482015260448101829052919350915073ffffffffffffffffffffffffffffffffffffffff8316906323b872dd906064016020604051808303816000875af1158015610b59573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b7d9190611d21565b610bb3576040517f0a7cd6d600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610bbd3385611633565b60025460405133917f9ffbffc04a397460ee1dbe8c9503e098090567d6b7f4b3c02a8617d800b6d95591610bf691899189918791612080565b60405180910390a2506001949350505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610c56576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116610ca3576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fd55614e962c5fd6ece71614f6348d702468a997a394dd5e5c1677950226d97ae906020015b60405180910390a150565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610d6a576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116610db7576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600880547fffffffffffffffffffffff0000000000000000000000000000000000000000ff1661010073ffffffffffffffffffffffffffffffffffffffff8416908102919091179091556040519081527f88815d964e380677e86d817e7d65dea59cb7b4c3b5b7a0c8ec7ea4a74f90a38790602001610d12565b600080546040517f0be155470000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060048201528291829173ffffffffffffffffffffffffffffffffffffffff90911690630be1554790602401602060405180830381865afa158015610ec4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ee891906120af565b905073ffffffffffffffffffffffffffffffffffffffff8116610f37576040517f78fff39600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080546040517fd7fd7afb0000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000600482015273ffffffffffffffffffffffffffffffffffffffff9091169063d7fd7afb90602401602060405180830381865afa158015610fc6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fea91906120cc565b905080600003611026576040517fe661aed000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000600254600154836110399190611d0a565b6110439190611d72565b92959294509192505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461109c576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60028190556040518181527fef13af88e424b5d15f49c77758542c1938b08b8b95b91ed0751f98ba99000d8f90602001610d12565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461111e576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018190556040518181527fff5788270f43bfc1ca41c503606d2594aa3023a1a7547de403a3e2f146a4a80a90602001610d12565b600080546040517f0be155470000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060048201528291829173ffffffffffffffffffffffffffffffffffffffff90911690630be1554790602401602060405180830381865afa1580156111e6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061120a91906120af565b905073ffffffffffffffffffffffffffffffffffffffff8116611259576040517f78fff39600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080546040517fd7fd7afb0000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000600482015273ffffffffffffffffffffffffffffffffffffffff9091169063d7fd7afb90602401602060405180830381865afa1580156112e8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061130c91906120cc565b905080600003611348576040517fe661aed000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002546000906113588784611d0a565b6113629190611d72565b9296929550919350505050565b73ffffffffffffffffffffffffffffffffffffffff83166113bc576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216611409576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff83811660008181526004602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff83166114c5576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216611512576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff831660009081526003602052604090205481811015611572576040517ffe382aa700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61157c8282611cd5565b73ffffffffffffffffffffffffffffffffffffffff80861660009081526003602052604080822093909355908516815290812080548492906115bf908490611d72565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161162591815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff8216611680576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216600090815260036020526040902054818110156116e0576040517ffe382aa700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6116ea8282611cd5565b73ffffffffffffffffffffffffffffffffffffffff841660009081526003602052604081209190915560058054849290611725908490611cd5565b909155505060405182815260009073ffffffffffffffffffffffffffffffffffffffff8516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200161146b565b73ffffffffffffffffffffffffffffffffffffffff82166117c2576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600560008282546117d49190611d72565b909155505073ffffffffffffffffffffffffffffffffffffffff82166000908152600360205260408120805483929061180e908490611d72565b909155505060405181815273ffffffffffffffffffffffffffffffffffffffff8316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b6000815180845260005b8181101561188b5760208185018101518683018201520161186f565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b6020815260006118dc6020830184611865565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff8116811461190557600080fd5b50565b6000806040838503121561191b57600080fd5b8235611926816118e3565b946020939093013593505050565b60008060006060848603121561194957600080fd5b8335611954816118e3565b92506020840135611964816118e3565b929592945050506040919091013590565b60006020828403121561198757600080fd5b5035919050565b6000602082840312156119a057600080fd5b81356118dc816118e3565b60008083601f8401126119bd57600080fd5b50813567ffffffffffffffff8111156119d557600080fd5b6020830191508360208260051b85010111156119f057600080fd5b9250929050565b60008060008060408587031215611a0d57600080fd5b843567ffffffffffffffff811115611a2457600080fd5b611a30878288016119ab565b909550935050602085013567ffffffffffffffff811115611a5057600080fd5b611a5c878288016119ab565b95989497509550505050565b6020810160038310611aa3577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008067ffffffffffffffff841115611af357611af3611aa9565b506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85018116603f0116810181811067ffffffffffffffff82111715611b4057611b40611aa9565b604052838152905080828401851015611b5857600080fd5b83836020830137600060208583010152509392505050565b600060208284031215611b8257600080fd5b813567ffffffffffffffff811115611b9957600080fd5b8201601f81018413611baa57600080fd5b611bb984823560208401611ad8565b949350505050565b60008060408385031215611bd457600080fd5b823567ffffffffffffffff811115611beb57600080fd5b8301601f81018513611bfc57600080fd5b611c0b85823560208401611ad8565b95602094909401359450505050565b60008060408385031215611c2d57600080fd5b8235611c38816118e3565b91506020830135611c48816118e3565b809150509250929050565b600181811c90821680611c6757607f821691505b602082108103611ca0577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8181038181111561058e5761058e611ca6565b604081526000611cfb6040830185611865565b90508260208301529392505050565b808202811582820484141761058e5761058e611ca6565b600060208284031215611d3357600080fd5b815180151581146118dc57600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b8082018082111561058e5761058e611ca6565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b81835260007f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff831115611e0057600080fd5b8260051b80836020870137939093016020019392505050565b60808082528101869052600060a0600588901b8301810190830189837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe136839003015b8b821015611eeb577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff608786030184528235818112611e9957600080fd5b8d0160208101903567ffffffffffffffff811115611eb657600080fd5b803603821315611ec557600080fd5b611ed0878284611d85565b96505050602083019250602084019350600182019150611e5c565b505050508281036020840152611f02818789611dce565b6040840195909552505060600152949350505050565b601f821115611f6257806000526020600020601f840160051c81016020851015611f3f5750805b601f840160051c820191505b81811015611f5f5760008155600101611f4b565b50505b505050565b815167ffffffffffffffff811115611f8157611f81611aa9565b611f9581611f8f8454611c53565b84611f18565b6020601f821160018114611fe75760008315611fb15750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b178455611f5f565b6000848152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b828110156120355787850151825560209485019460019092019101612015565b508482101561207157868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b6080815260006120936080830187611865565b6020830195909552506040810192909252606090910152919050565b6000602082840312156120c157600080fd5b81516118dc816118e3565b6000602082840312156120de57600080fd5b505191905056fea26469706673582212202e6f7cf07b30388f6da980f3d117c7dec350444468bcd09758f2cc08bcbecce964736f6c634300081a0033",
}

// ZRC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use ZRC20MetaData.ABI instead.
var ZRC20ABI = ZRC20MetaData.ABI

// ZRC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZRC20MetaData.Bin instead.
var ZRC20Bin = ZRC20MetaData.Bin

// DeployZRC20 deploys a new Ethereum contract, binding an instance of ZRC20 to it.
func DeployZRC20(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string, decimals_ uint8, chainid_ *big.Int, coinType_ uint8, gasLimit_ *big.Int, systemContractAddress_ common.Address, gatewayAddress_ common.Address) (common.Address, *types.Transaction, *ZRC20, error) {
	parsed, err := ZRC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZRC20Bin), backend, name_, symbol_, decimals_, chainid_, coinType_, gasLimit_, systemContractAddress_, gatewayAddress_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZRC20{ZRC20Caller: ZRC20Caller{contract: contract}, ZRC20Transactor: ZRC20Transactor{contract: contract}, ZRC20Filterer: ZRC20Filterer{contract: contract}}, nil
}

// ZRC20 is an auto generated Go binding around an Ethereum contract.
type ZRC20 struct {
	ZRC20Caller     // Read-only binding to the contract
	ZRC20Transactor // Write-only binding to the contract
	ZRC20Filterer   // Log filterer for contract events
}

// ZRC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ZRC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZRC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ZRC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZRC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZRC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZRC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZRC20Session struct {
	Contract     *ZRC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZRC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZRC20CallerSession struct {
	Contract *ZRC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ZRC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZRC20TransactorSession struct {
	Contract     *ZRC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZRC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ZRC20Raw struct {
	Contract *ZRC20 // Generic contract binding to access the raw methods on
}

// ZRC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZRC20CallerRaw struct {
	Contract *ZRC20Caller // Generic read-only contract binding to access the raw methods on
}

// ZRC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZRC20TransactorRaw struct {
	Contract *ZRC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewZRC20 creates a new instance of ZRC20, bound to a specific deployed contract.
func NewZRC20(address common.Address, backend bind.ContractBackend) (*ZRC20, error) {
	contract, err := bindZRC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZRC20{ZRC20Caller: ZRC20Caller{contract: contract}, ZRC20Transactor: ZRC20Transactor{contract: contract}, ZRC20Filterer: ZRC20Filterer{contract: contract}}, nil
}

// NewZRC20Caller creates a new read-only instance of ZRC20, bound to a specific deployed contract.
func NewZRC20Caller(address common.Address, caller bind.ContractCaller) (*ZRC20Caller, error) {
	contract, err := bindZRC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZRC20Caller{contract: contract}, nil
}

// NewZRC20Transactor creates a new write-only instance of ZRC20, bound to a specific deployed contract.
func NewZRC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ZRC20Transactor, error) {
	contract, err := bindZRC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZRC20Transactor{contract: contract}, nil
}

// NewZRC20Filterer creates a new log filterer instance of ZRC20, bound to a specific deployed contract.
func NewZRC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ZRC20Filterer, error) {
	contract, err := bindZRC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZRC20Filterer{contract: contract}, nil
}

// bindZRC20 binds a generic wrapper to an already deployed contract.
func bindZRC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZRC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZRC20 *ZRC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZRC20.Contract.ZRC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZRC20 *ZRC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZRC20.Contract.ZRC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZRC20 *ZRC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZRC20.Contract.ZRC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZRC20 *ZRC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZRC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZRC20 *ZRC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZRC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZRC20 *ZRC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZRC20.Contract.contract.Transact(opts, method, params...)
}

// CHAINID is a free data retrieval call binding the contract method 0x85e1f4d0.
//
// Solidity: function CHAIN_ID() view returns(uint256)
func (_ZRC20 *ZRC20Caller) CHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZRC20.contract.Call(opts, &out, "CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CHAINID is a free data retrieval call binding the contract method 0x85e1f4d0.
//
// Solidity: function CHAIN_ID() view returns(uint256)
func (_ZRC20 *ZRC20Session) CHAINID() (*big.Int, error) {
	return _ZRC20.Contract.CHAINID(&_ZRC20.CallOpts)
}

// CHAINID is a free data retrieval call binding the contract method 0x85e1f4d0.
//
// Solidity: function CHAIN_ID() view returns(uint256)
func (_ZRC20 *ZRC20CallerSession) CHAINID() (*big.Int, error) {
	return _ZRC20.Contract.CHAINID(&_ZRC20.CallOpts)
}

// COINTYPE is a free data retrieval call binding the contract method 0xa3413d03.
//
// Solidity: function COIN_TYPE() view returns(uint8)
func (_ZRC20 *ZRC20Caller) COINTYPE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ZRC20.contract.Call(opts, &out, "COIN_TYPE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// COINTYPE is a free data retrieval call binding the contract method 0xa3413d03.
//
// Solidity: function COIN_TYPE() view returns(uint8)
func (_ZRC20 *ZRC20Session) COINTYPE() (uint8, error) {
	return _ZRC20.Contract.COINTYPE(&_ZRC20.CallOpts)
}

// COINTYPE is a free data retrieval call binding the contract method 0xa3413d03.
//
// Solidity: function COIN_TYPE() view returns(uint8)
func (_ZRC20 *ZRC20CallerSession) COINTYPE() (uint8, error) {
	return _ZRC20.Contract.COINTYPE(&_ZRC20.CallOpts)
}

// FUNGIBLEMODULEADDRESS is a free data retrieval call binding the contract method 0x3ce4a5bc.
//
// Solidity: function FUNGIBLE_MODULE_ADDRESS() view returns(address)
func (_ZRC20 *ZRC20Caller) FUNGIBLEMODULEADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZRC20.contract.Call(opts, &out, "FUNGIBLE_MODULE_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FUNGIBLEMODULEADDRESS is a free data retrieval call binding the contract method 0x3ce4a5bc.
//
// Solidity: function FUNGIBLE_MODULE_ADDRESS() view returns(address)
func (_ZRC20 *ZRC20Session) FUNGIBLEMODULEADDRESS() (common.Address, error) {
	return _ZRC20.Contract.FUNGIBLEMODULEADDRESS(&_ZRC20.CallOpts)
}

// FUNGIBLEMODULEADDRESS is a free data retrieval call binding the contract method 0x3ce4a5bc.
//
// Solidity: function FUNGIBLE_MODULE_ADDRESS() view returns(address)
func (_ZRC20 *ZRC20CallerSession) FUNGIBLEMODULEADDRESS() (common.Address, error) {
	return _ZRC20.Contract.FUNGIBLEMODULEADDRESS(&_ZRC20.CallOpts)
}

// GASLIMIT is a free data retrieval call binding the contract method 0x091d2788.
//
// Solidity: function GAS_LIMIT() view returns(uint256)
func (_ZRC20 *ZRC20Caller) GASLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZRC20.contract.Call(opts, &out, "GAS_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GASLIMIT is a free data retrieval call binding the contract method 0x091d2788.
//
// Solidity: function GAS_LIMIT() view returns(uint256)
func (_ZRC20 *ZRC20Session) GASLIMIT() (*big.Int, error) {
	return _ZRC20.Contract.GASLIMIT(&_ZRC20.CallOpts)
}

// GASLIMIT is a free data retrieval call binding the contract method 0x091d2788.
//
// Solidity: function GAS_LIMIT() view returns(uint256)
func (_ZRC20 *ZRC20CallerSession) GASLIMIT() (*big.Int, error) {
	return _ZRC20.Contract.GASLIMIT(&_ZRC20.CallOpts)
}

// MAXBATCHSIZE is a free data retrieval call binding the contract method 0xcfdbf254.
//
// Solidity: function MAX_BATCH_SIZE() view returns(uint256)
func (_ZRC20 *ZRC20Caller) MAXBATCHSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZRC20.contract.Call(opts, &out, "MAX_BATCH_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXBATCHSIZE is a free data retrieval call binding the contract method 0xcfdbf254.
//
// Solidity: function MAX_BATCH_SIZE() view returns(uint256)
func (_ZRC20 *ZRC20Session) MAXBATCHSIZE() (*big.Int, error) {
	return _ZRC20.Contract.MAXBATCHSIZE(&_ZRC20.CallOpts)
}

// MAXBATCHSIZE is a free data retrieval call binding the contract method 0xcfdbf254.
//
// Solidity: function MAX_BATCH_SIZE() view returns(uint256)
func (_ZRC20 *ZRC20CallerSession) MAXBATCHSIZE() (*big.Int, error) {
	return _ZRC20.Contract.MAXBATCHSIZE(&_ZRC20.CallOpts)
}

// PROTOCOLFLATFEE is a free data retrieval call binding the contract method 0x4d8943bb.
//
// Solidity: function PROTOCOL_FLAT_FEE() view returns(uint256)
func (_ZRC20 *ZRC20Caller) PROTOCOLFLATFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZRC20.contract.Call(opts, &out, "PROTOCOL_FLAT_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PROTOCOLFLATFEE is a free data retrieval call binding the contract method 0x4d8943bb.
//
// Solidity: function PROTOCOL_FLAT_FEE() view returns(uint256)
func (_ZRC20 *ZRC20Session) PROTOCOLFLATFEE() (*big.Int, error) {
	return _ZRC20.Contract.PROTOCOLFLATFEE(&_ZRC20.CallOpts)
}

// PROTOCOLFLATFEE is a free data retrieval call binding the contract method 0x4d8943bb.
//
// Solidity: function PROTOCOL_FLAT_FEE() view returns(uint256)
func (_ZRC20 *ZRC20CallerSession) PROTOCOLFLATFEE() (*big.Int, error) {
	return _ZRC20.Contract.PROTOCOLFLATFEE(&_ZRC20.CallOpts)
}

// SYSTEMCONTRACTADDRESS is a free data retrieval call binding the contract method 0xf2441b32.
//
// Solidity: function SYSTEM_CONTRACT_ADDRESS() view returns(address)
func (_ZRC20 *ZRC20Caller) SYSTEMCONTRACTADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZRC20.contract.Call(opts, &out, "SYSTEM_CONTRACT_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SYSTEMCONTRACTADDRESS is a free data retrieval call binding the contract method 0xf2441b32.
//
// Solidity: function SYSTEM_CONTRACT_ADDRESS() view returns(address)
func (_ZRC20 *ZRC20Session) SYSTEMCONTRACTADDRESS() (common.Address, error) {
	return _ZRC20.Contract.SYSTEMCONTRACTADDRESS(&_ZRC20.CallOpts)
}

// SYSTEMCONTRACTADDRESS is a free data retrieval call binding the contract method 0xf2441b32.
//
// Solidity: function SYSTEM_CONTRACT_ADDRESS() view returns(address)
func (_ZRC20 *ZRC20CallerSession) SYSTEMCONTRACTADDRESS() (common.Address, error) {
	return _ZRC20.Contract.SYSTEMCONTRACTADDRESS(&_ZRC20.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ZRC20 *ZRC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZRC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ZRC20 *ZRC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ZRC20.Contract.Allowance(&_ZRC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ZRC20 *ZRC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ZRC20.Contract.Allowance(&_ZRC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ZRC20 *ZRC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZRC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ZRC20 *ZRC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _ZRC20.Contract.BalanceOf(&_ZRC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ZRC20 *ZRC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ZRC20.Contract.BalanceOf(&_ZRC20.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ZRC20 *ZRC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ZRC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ZRC20 *ZRC20Session) Decimals() (uint8, error) {
	return _ZRC20.Contract.Decimals(&_ZRC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ZRC20 *ZRC20CallerSession) Decimals() (uint8, error) {
	return _ZRC20.Contract.Decimals(&_ZRC20.CallOpts)
}

// GatewayAddress is a free data retrieval call binding the contract method 0x8b851b95.
//
// Solidity: function gatewayAddress() view returns(address)
func (_ZRC20 *ZRC20Caller) GatewayAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZRC20.contract.Call(opts, &out, "gatewayAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GatewayAddress is a free data retrieval call binding the contract method 0x8b851b95.
//
// Solidity: function gatewayAddress() view returns(address)
func (_ZRC20 *ZRC20Session) GatewayAddress() (common.Address, error) {
	return _ZRC20.Contract.GatewayAddress(&_ZRC20.CallOpts)
}

// GatewayAddress is a free data retrieval call binding the contract method 0x8b851b95.
//
// Solidity: function gatewayAddress() view returns(address)
func (_ZRC20 *ZRC20CallerSession) GatewayAddress() (common.Address, error) {
	return _ZRC20.Contract.GatewayAddress(&_ZRC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ZRC20 *ZRC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ZRC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ZRC20 *ZRC20Session) Name() (string, error) {
	return _ZRC20.Contract.Name(&_ZRC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ZRC20 *ZRC20CallerSession) Name() (string, error) {
	return _ZRC20.Contract.Name(&_ZRC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ZRC20 *ZRC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ZRC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ZRC20 *ZRC20Session) Symbol() (string, error) {
	return _ZRC20.Contract.Symbol(&_ZRC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ZRC20 *ZRC20CallerSession) Symbol() (string, error) {
	return _ZRC20.Contract.Symbol(&_ZRC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ZRC20 *ZRC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZRC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ZRC20 *ZRC20Session) TotalSupply() (*big.Int, error) {
	return _ZRC20.Contract.TotalSupply(&_ZRC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ZRC20 *ZRC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ZRC20.Contract.TotalSupply(&_ZRC20.CallOpts)
}

// WithdrawGasFee is a free data retrieval call binding the contract method 0xd9eeebed.
//
// Solidity: function withdrawGasFee() view returns(address, uint256)
func (_ZRC20 *ZRC20Caller) WithdrawGasFee(opts *bind.CallOpts) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _ZRC20.contract.Call(opts, &out, "withdrawGasFee")

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// WithdrawGasFee is a free data retrieval call binding the contract method 0xd9eeebed.
//
// Solidity: function withdrawGasFee() view returns(address, uint256)
func (_ZRC20 *ZRC20Session) WithdrawGasFee() (common.Address, *big.Int, error) {
	return _ZRC20.Contract.WithdrawGasFee(&_ZRC20.CallOpts)
}

// WithdrawGasFee is a free data retrieval call binding the contract method 0xd9eeebed.
//
// Solidity: function withdrawGasFee() view returns(address, uint256)
func (_ZRC20 *ZRC20CallerSession) WithdrawGasFee() (common.Address, *big.Int, error) {
	return _ZRC20.Contract.WithdrawGasFee(&_ZRC20.CallOpts)
}

// WithdrawGasFeeWithGasLimit is a free data retrieval call binding the contract method 0xfc5fecd5.
//
// Solidity: function withdrawGasFeeWithGasLimit(uint256 gasLimit) view returns(address, uint256)
func (_ZRC20 *ZRC20Caller) WithdrawGasFeeWithGasLimit(opts *bind.CallOpts, gasLimit *big.Int) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _ZRC20.contract.Call(opts, &out, "withdrawGasFeeWithGasLimit", gasLimit)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// WithdrawGasFeeWithGasLimit is a free data retrieval call binding the contract method 0xfc5fecd5.
//
// Solidity: function withdrawGasFeeWithGasLimit(uint256 gasLimit) view returns(address, uint256)
func (_ZRC20 *ZRC20Session) WithdrawGasFeeWithGasLimit(gasLimit *big.Int) (common.Address, *big.Int, error) {
	return _ZRC20.Contract.WithdrawGasFeeWithGasLimit(&_ZRC20.CallOpts, gasLimit)
}

// WithdrawGasFeeWithGasLimit is a free data retrieval call binding the contract method 0xfc5fecd5.
//
// Solidity: function withdrawGasFeeWithGasLimit(uint256 gasLimit) view returns(address, uint256)
func (_ZRC20 *ZRC20CallerSession) WithdrawGasFeeWithGasLimit(gasLimit *big.Int) (common.Address, *big.Int, error) {
	return _ZRC20.Contract.WithdrawGasFeeWithGasLimit(&_ZRC20.CallOpts, gasLimit)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ZRC20 *ZRC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZRC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ZRC20 *ZRC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZRC20.Contract.Approve(&_ZRC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ZRC20 *ZRC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZRC20.Contract.Approve(&_ZRC20.TransactOpts, spender, amount)
}

// BatchWithdraw is a paid mutator transaction binding the contract method 0x73486a78.
//
// Solidity: function batchWithdraw(bytes[] recipients, uint256[] amounts) returns(bool success)
func (_ZRC20 *ZRC20Transactor) BatchWithdraw(opts *bind.TransactOpts, recipients [][]byte, amounts []*big.Int) (*types.Transaction, error) {
	return _ZRC20.contract.Transact(opts, "batchWithdraw", recipients, amounts)
}

// BatchWithdraw is a paid mutator transaction binding the contract method 0x73486a78.
//
// Solidity: function batchWithdraw(bytes[] recipients, uint256[] amounts) returns(bool success)
func (_ZRC20 *ZRC20Session) BatchWithdraw(recipients [][]byte, amounts []*big.Int) (*types.Transaction, error) {
	return _ZRC20.Contract.BatchWithdraw(&_ZRC20.TransactOpts, recipients, amounts)
}

// BatchWithdraw is a paid mutator transaction binding the contract method 0x73486a78.
//
// Solidity: function batchWithdraw(bytes[] recipients, uint256[] amounts) returns(bool success)
func (_ZRC20 *ZRC20TransactorSession) BatchWithdraw(recipients [][]byte, amounts []*big.Int) (*types.Transaction, error) {
	return _ZRC20.Contract.BatchWithdraw(&_ZRC20.TransactOpts, recipients, amounts)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns(bool)
func (_ZRC20 *ZRC20Transactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ZRC20.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns(bool)
func (_ZRC20 *ZRC20Session) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ZRC20.Contract.Burn(&_ZRC20.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns(bool)
func (_ZRC20 *ZRC20TransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ZRC20.Contract.Burn(&_ZRC20.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address to, uint256 amount) returns(bool)
func (_ZRC20 *ZRC20Transactor) Deposit(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZRC20.contract.Transact(opts, "deposit", to, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address to, uint256 amount) returns(bool)
func (_ZRC20 *ZRC20Session) Deposit(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZRC20.Contract.Deposit(&_ZRC20.TransactOpts, to, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address to, uint256 amount) returns(bool)
func (_ZRC20 *ZRC20TransactorSession) Deposit(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZRC20.Contract.Deposit(&_ZRC20.TransactOpts, to, amount)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string newName) returns()
func (_ZRC20 *ZRC20Transactor) SetName(opts *bind.TransactOpts, newName string) (*types.Transaction, error) {
	return _ZRC20.contract.Transact(opts, "setName", newName)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string newName) returns()
func (_ZRC20 *ZRC20Session) SetName(newName string) (*types.Transaction, error) {
	return _ZRC20.Contract.SetName(&_ZRC20.TransactOpts, newName)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string newName) returns()
func (_ZRC20 *ZRC20TransactorSession) SetName(newName string) (*types.Transaction, error) {
	return _ZRC20.Contract.SetName(&_ZRC20.TransactOpts, newName)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string newSymbol) returns()
func (_ZRC20 *ZRC20Transactor) SetSymbol(opts *bind.TransactOpts, newSymbol string) (*types.Transaction, error) {
	return _ZRC20.contract.Transact(opts, "setSymbol", newSymbol)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string newSymbol) returns()
func (_ZRC20 *ZRC20Session) SetSymbol(newSymbol string) (*types.Transaction, error) {
	return _ZRC20.Contract.SetSymbol(&_ZRC20.TransactOpts, newSymbol)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string newSymbol) returns()
func (_ZRC20 *ZRC20TransactorSession) SetSymbol(newSymbol string) (*types.Transaction, error) {
	return _ZRC20.Contract.SetSymbol(&_ZRC20.TransactOpts, newSymbol)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ZRC20 *ZRC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZRC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ZRC20 *ZRC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZRC20.Contract.Transfer(&_ZRC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ZRC20 *ZRC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZRC20.Contract.Transfer(&_ZRC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ZRC20 *ZRC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZRC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ZRC20 *ZRC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZRC20.Contract.TransferFrom(&_ZRC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ZRC20 *ZRC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZRC20.Contract.TransferFrom(&_ZRC20.TransactOpts, sender, recipient, amount)
}

// UpdateGasLimit is a paid mutator transaction binding the contract method 0xf687d12a.
//
// Solidity: function updateGasLimit(uint256 gasLimit_) returns()
func (_ZRC20 *ZRC20Transactor) UpdateGasLimit(opts *bind.TransactOpts, gasLimit_ *big.Int) (*types.Transaction, error) {
	return _ZRC20.contract.Transact(opts, "updateGasLimit", gasLimit_)
}

// UpdateGasLimit is a paid mutator transaction binding the contract method 0xf687d12a.
//
// Solidity: function updateGasLimit(uint256 gasLimit_) returns()
func (_ZRC20 *ZRC20Session) UpdateGasLimit(gasLimit_ *big.Int) (*types.Transaction, error) {
	return _ZRC20.Contract.UpdateGasLimit(&_ZRC20.TransactOpts, gasLimit_)
}

// UpdateGasLimit is a paid mutator transaction binding the contract method 0xf687d12a.
//
// Solidity: function updateGasLimit(uint256 gasLimit_) returns()
func (_ZRC20 *ZRC20TransactorSession) UpdateGasLimit(gasLimit_ *big.Int) (*types.Transaction, error) {
	return _ZRC20.Contract.UpdateGasLimit(&_ZRC20.TransactOpts, gasLimit_)
}

// UpdateGatewayAddress is a paid mutator transaction binding the contract method 0xccc77599.
//
// Solidity: function updateGatewayAddress(address addr) returns()
func (_ZRC20 *ZRC20Transactor) UpdateGatewayAddress(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _ZRC20.contract.Transact(opts, "updateGatewayAddress", addr)
}

// UpdateGatewayAddress is a paid mutator transaction binding the contract method 0xccc77599.
//
// Solidity: function updateGatewayAddress(address addr) returns()
func (_ZRC20 *ZRC20Session) UpdateGatewayAddress(addr common.Address) (*types.Transaction, error) {
	return _ZRC20.Contract.UpdateGatewayAddress(&_ZRC20.TransactOpts, addr)
}

// UpdateGatewayAddress is a paid mutator transaction binding the contract method 0xccc77599.
//
// Solidity: function updateGatewayAddress(address addr) returns()
func (_ZRC20 *ZRC20TransactorSession) UpdateGatewayAddress(addr common.Address) (*types.Transaction, error) {
	return _ZRC20.Contract.UpdateGatewayAddress(&_ZRC20.TransactOpts, addr)
}

// UpdateProtocolFlatFee is a paid mutator transaction binding the contract method 0xeddeb123.
//
// Solidity: function updateProtocolFlatFee(uint256 protocolFlatFee_) returns()
func (_ZRC20 *ZRC20Transactor) UpdateProtocolFlatFee(opts *bind.TransactOpts, protocolFlatFee_ *big.Int) (*types.Transaction, error) {
	return _ZRC20.contract.Transact(opts, "updateProtocolFlatFee", protocolFlatFee_)
}

// UpdateProtocolFlatFee is a paid mutator transaction binding the contract method 0xeddeb123.
//
// Solidity: function updateProtocolFlatFee(uint256 protocolFlatFee_) returns()
func (_ZRC20 *ZRC20Session) UpdateProtocolFlatFee(protocolFlatFee_ *big.Int) (*types.Transaction, error) {
	return _ZRC20.Contract.UpdateProtocolFlatFee(&_ZRC20.TransactOpts, protocolFlatFee_)
}

// UpdateProtocolFlatFee is a paid mutator transaction binding the contract method 0xeddeb123.
//
// Solidity: function updateProtocolFlatFee(uint256 protocolFlatFee_) returns()
func (_ZRC20 *ZRC20TransactorSession) UpdateProtocolFlatFee(protocolFlatFee_ *big.Int) (*types.Transaction, error) {
	return _ZRC20.Contract.UpdateProtocolFlatFee(&_ZRC20.TransactOpts, protocolFlatFee_)
}

// UpdateSystemContractAddress is a paid mutator transaction binding the contract method 0xc835d7cc.
//
// Solidity: function updateSystemContractAddress(address addr) returns()
func (_ZRC20 *ZRC20Transactor) UpdateSystemContractAddress(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _ZRC20.contract.Transact(opts, "updateSystemContractAddress", addr)
}

// UpdateSystemContractAddress is a paid mutator transaction binding the contract method 0xc835d7cc.
//
// Solidity: function updateSystemContractAddress(address addr) returns()
func (_ZRC20 *ZRC20Session) UpdateSystemContractAddress(addr common.Address) (*types.Transaction, error) {
	return _ZRC20.Contract.UpdateSystemContractAddress(&_ZRC20.TransactOpts, addr)
}

// UpdateSystemContractAddress is a paid mutator transaction binding the contract method 0xc835d7cc.
//
// Solidity: function updateSystemContractAddress(address addr) returns()
func (_ZRC20 *ZRC20TransactorSession) UpdateSystemContractAddress(addr common.Address) (*types.Transaction, error) {
	return _ZRC20.Contract.UpdateSystemContractAddress(&_ZRC20.TransactOpts, addr)
}

// Withdraw is a paid mutator transaction binding the contract method 0xc7012626.
//
// Solidity: function withdraw(bytes to, uint256 amount) returns(bool)
func (_ZRC20 *ZRC20Transactor) Withdraw(opts *bind.TransactOpts, to []byte, amount *big.Int) (*types.Transaction, error) {
	return _ZRC20.contract.Transact(opts, "withdraw", to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xc7012626.
//
// Solidity: function withdraw(bytes to, uint256 amount) returns(bool)
func (_ZRC20 *ZRC20Session) Withdraw(to []byte, amount *big.Int) (*types.Transaction, error) {
	return _ZRC20.Contract.Withdraw(&_ZRC20.TransactOpts, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xc7012626.
//
// Solidity: function withdraw(bytes to, uint256 amount) returns(bool)
func (_ZRC20 *ZRC20TransactorSession) Withdraw(to []byte, amount *big.Int) (*types.Transaction, error) {
	return _ZRC20.Contract.Withdraw(&_ZRC20.TransactOpts, to, amount)
}

// ZRC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ZRC20 contract.
type ZRC20ApprovalIterator struct {
	Event *ZRC20Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZRC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZRC20Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZRC20Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZRC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZRC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZRC20Approval represents a Approval event raised by the ZRC20 contract.
type ZRC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ZRC20 *ZRC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ZRC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ZRC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ZRC20ApprovalIterator{contract: _ZRC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ZRC20 *ZRC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ZRC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ZRC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZRC20Approval)
				if err := _ZRC20.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ZRC20 *ZRC20Filterer) ParseApproval(log types.Log) (*ZRC20Approval, error) {
	event := new(ZRC20Approval)
	if err := _ZRC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZRC20BatchWithdrawalIterator is returned from FilterBatchWithdrawal and is used to iterate over the raw logs and unpacked data for BatchWithdrawal events raised by the ZRC20 contract.
type ZRC20BatchWithdrawalIterator struct {
	Event *ZRC20BatchWithdrawal // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZRC20BatchWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZRC20BatchWithdrawal)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZRC20BatchWithdrawal)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZRC20BatchWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZRC20BatchWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZRC20BatchWithdrawal represents a BatchWithdrawal event raised by the ZRC20 contract.
type ZRC20BatchWithdrawal struct {
	From            common.Address
	Recipients      [][]byte
	Values          []*big.Int
	GasFeePerTx     *big.Int
	ProtocolFlatFee *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBatchWithdrawal is a free log retrieval operation binding the contract event 0x2512245b924c3fcf7f4aed2a20c0e35ee10a0f15114842e5e46732128d8cab3a.
//
// Solidity: event BatchWithdrawal(address indexed from, bytes[] recipients, uint256[] values, uint256 gasFeePerTx, uint256 protocolFlatFee)
func (_ZRC20 *ZRC20Filterer) FilterBatchWithdrawal(opts *bind.FilterOpts, from []common.Address) (*ZRC20BatchWithdrawalIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _ZRC20.contract.FilterLogs(opts, "BatchWithdrawal", fromRule)
	if err != nil {
		return nil, err
	}
	return &ZRC20BatchWithdrawalIterator{contract: _ZRC20.contract, event: "BatchWithdrawal", logs: logs, sub: sub}, nil
}

// WatchBatchWithdrawal is a free log subscription operation binding the contract event 0x2512245b924c3fcf7f4aed2a20c0e35ee10a0f15114842e5e46732128d8cab3a.
//
// Solidity: event BatchWithdrawal(address indexed from, bytes[] recipients, uint256[] values, uint256 gasFeePerTx, uint256 protocolFlatFee)
func (_ZRC20 *ZRC20Filterer) WatchBatchWithdrawal(opts *bind.WatchOpts, sink chan<- *ZRC20BatchWithdrawal, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _ZRC20.contract.WatchLogs(opts, "BatchWithdrawal", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZRC20BatchWithdrawal)
				if err := _ZRC20.contract.UnpackLog(event, "BatchWithdrawal", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBatchWithdrawal is a log parse operation binding the contract event 0x2512245b924c3fcf7f4aed2a20c0e35ee10a0f15114842e5e46732128d8cab3a.
//
// Solidity: event BatchWithdrawal(address indexed from, bytes[] recipients, uint256[] values, uint256 gasFeePerTx, uint256 protocolFlatFee)
func (_ZRC20 *ZRC20Filterer) ParseBatchWithdrawal(log types.Log) (*ZRC20BatchWithdrawal, error) {
	event := new(ZRC20BatchWithdrawal)
	if err := _ZRC20.contract.UnpackLog(event, "BatchWithdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZRC20DepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the ZRC20 contract.
type ZRC20DepositIterator struct {
	Event *ZRC20Deposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZRC20DepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZRC20Deposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZRC20Deposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZRC20DepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZRC20DepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZRC20Deposit represents a Deposit event raised by the ZRC20 contract.
type ZRC20Deposit struct {
	From  []byte
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x67fc7bdaed5b0ec550d8706b87d60568ab70c6b781263c70101d54cd1564aab3.
//
// Solidity: event Deposit(bytes from, address indexed to, uint256 value)
func (_ZRC20 *ZRC20Filterer) FilterDeposit(opts *bind.FilterOpts, to []common.Address) (*ZRC20DepositIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZRC20.contract.FilterLogs(opts, "Deposit", toRule)
	if err != nil {
		return nil, err
	}
	return &ZRC20DepositIterator{contract: _ZRC20.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x67fc7bdaed5b0ec550d8706b87d60568ab70c6b781263c70101d54cd1564aab3.
//
// Solidity: event Deposit(bytes from, address indexed to, uint256 value)
func (_ZRC20 *ZRC20Filterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ZRC20Deposit, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZRC20.contract.WatchLogs(opts, "Deposit", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZRC20Deposit)
				if err := _ZRC20.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposit is a log parse operation binding the contract event 0x67fc7bdaed5b0ec550d8706b87d60568ab70c6b781263c70101d54cd1564aab3.
//
// Solidity: event Deposit(bytes from, address indexed to, uint256 value)
func (_ZRC20 *ZRC20Filterer) ParseDeposit(log types.Log) (*ZRC20Deposit, error) {
	event := new(ZRC20Deposit)
	if err := _ZRC20.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZRC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ZRC20 contract.
type ZRC20TransferIterator struct {
	Event *ZRC20Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZRC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZRC20Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZRC20Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZRC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZRC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZRC20Transfer represents a Transfer event raised by the ZRC20 contract.
type ZRC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ZRC20 *ZRC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ZRC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZRC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ZRC20TransferIterator{contract: _ZRC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ZRC20 *ZRC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ZRC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZRC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZRC20Transfer)
				if err := _ZRC20.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ZRC20 *ZRC20Filterer) ParseTransfer(log types.Log) (*ZRC20Transfer, error) {
	event := new(ZRC20Transfer)
	if err := _ZRC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZRC20UpdatedGasLimitIterator is returned from FilterUpdatedGasLimit and is used to iterate over the raw logs and unpacked data for UpdatedGasLimit events raised by the ZRC20 contract.
type ZRC20UpdatedGasLimitIterator struct {
	Event *ZRC20UpdatedGasLimit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZRC20UpdatedGasLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZRC20UpdatedGasLimit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZRC20UpdatedGasLimit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZRC20UpdatedGasLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZRC20UpdatedGasLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZRC20UpdatedGasLimit represents a UpdatedGasLimit event raised by the ZRC20 contract.
type ZRC20UpdatedGasLimit struct {
	GasLimit *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdatedGasLimit is a free log retrieval operation binding the contract event 0xff5788270f43bfc1ca41c503606d2594aa3023a1a7547de403a3e2f146a4a80a.
//
// Solidity: event UpdatedGasLimit(uint256 gasLimit)
func (_ZRC20 *ZRC20Filterer) FilterUpdatedGasLimit(opts *bind.FilterOpts) (*ZRC20UpdatedGasLimitIterator, error) {

	logs, sub, err := _ZRC20.contract.FilterLogs(opts, "UpdatedGasLimit")
	if err != nil {
		return nil, err
	}
	return &ZRC20UpdatedGasLimitIterator{contract: _ZRC20.contract, event: "UpdatedGasLimit", logs: logs, sub: sub}, nil
}

// WatchUpdatedGasLimit is a free log subscription operation binding the contract event 0xff5788270f43bfc1ca41c503606d2594aa3023a1a7547de403a3e2f146a4a80a.
//
// Solidity: event UpdatedGasLimit(uint256 gasLimit)
func (_ZRC20 *ZRC20Filterer) WatchUpdatedGasLimit(opts *bind.WatchOpts, sink chan<- *ZRC20UpdatedGasLimit) (event.Subscription, error) {

	logs, sub, err := _ZRC20.contract.WatchLogs(opts, "UpdatedGasLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZRC20UpdatedGasLimit)
				if err := _ZRC20.contract.UnpackLog(event, "UpdatedGasLimit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdatedGasLimit is a log parse operation binding the contract event 0xff5788270f43bfc1ca41c503606d2594aa3023a1a7547de403a3e2f146a4a80a.
//
// Solidity: event UpdatedGasLimit(uint256 gasLimit)
func (_ZRC20 *ZRC20Filterer) ParseUpdatedGasLimit(log types.Log) (*ZRC20UpdatedGasLimit, error) {
	event := new(ZRC20UpdatedGasLimit)
	if err := _ZRC20.contract.UnpackLog(event, "UpdatedGasLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZRC20UpdatedGatewayIterator is returned from FilterUpdatedGateway and is used to iterate over the raw logs and unpacked data for UpdatedGateway events raised by the ZRC20 contract.
type ZRC20UpdatedGatewayIterator struct {
	Event *ZRC20UpdatedGateway // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZRC20UpdatedGatewayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZRC20UpdatedGateway)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZRC20UpdatedGateway)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZRC20UpdatedGatewayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZRC20UpdatedGatewayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZRC20UpdatedGateway represents a UpdatedGateway event raised by the ZRC20 contract.
type ZRC20UpdatedGateway struct {
	Gateway common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdatedGateway is a free log retrieval operation binding the contract event 0x88815d964e380677e86d817e7d65dea59cb7b4c3b5b7a0c8ec7ea4a74f90a387.
//
// Solidity: event UpdatedGateway(address gateway)
func (_ZRC20 *ZRC20Filterer) FilterUpdatedGateway(opts *bind.FilterOpts) (*ZRC20UpdatedGatewayIterator, error) {

	logs, sub, err := _ZRC20.contract.FilterLogs(opts, "UpdatedGateway")
	if err != nil {
		return nil, err
	}
	return &ZRC20UpdatedGatewayIterator{contract: _ZRC20.contract, event: "UpdatedGateway", logs: logs, sub: sub}, nil
}

// WatchUpdatedGateway is a free log subscription operation binding the contract event 0x88815d964e380677e86d817e7d65dea59cb7b4c3b5b7a0c8ec7ea4a74f90a387.
//
// Solidity: event UpdatedGateway(address gateway)
func (_ZRC20 *ZRC20Filterer) WatchUpdatedGateway(opts *bind.WatchOpts, sink chan<- *ZRC20UpdatedGateway) (event.Subscription, error) {

	logs, sub, err := _ZRC20.contract.WatchLogs(opts, "UpdatedGateway")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZRC20UpdatedGateway)
				if err := _ZRC20.contract.UnpackLog(event, "UpdatedGateway", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdatedGateway is a log parse operation binding the contract event 0x88815d964e380677e86d817e7d65dea59cb7b4c3b5b7a0c8ec7ea4a74f90a387.
//
// Solidity: event UpdatedGateway(address gateway)
func (_ZRC20 *ZRC20Filterer) ParseUpdatedGateway(log types.Log) (*ZRC20UpdatedGateway, error) {
	event := new(ZRC20UpdatedGateway)
	if err := _ZRC20.contract.UnpackLog(event, "UpdatedGateway", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZRC20UpdatedProtocolFlatFeeIterator is returned from FilterUpdatedProtocolFlatFee and is used to iterate over the raw logs and unpacked data for UpdatedProtocolFlatFee events raised by the ZRC20 contract.
type ZRC20UpdatedProtocolFlatFeeIterator struct {
	Event *ZRC20UpdatedProtocolFlatFee // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZRC20UpdatedProtocolFlatFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZRC20UpdatedProtocolFlatFee)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZRC20UpdatedProtocolFlatFee)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZRC20UpdatedProtocolFlatFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZRC20UpdatedProtocolFlatFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZRC20UpdatedProtocolFlatFee represents a UpdatedProtocolFlatFee event raised by the ZRC20 contract.
type ZRC20UpdatedProtocolFlatFee struct {
	ProtocolFlatFee *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdatedProtocolFlatFee is a free log retrieval operation binding the contract event 0xef13af88e424b5d15f49c77758542c1938b08b8b95b91ed0751f98ba99000d8f.
//
// Solidity: event UpdatedProtocolFlatFee(uint256 protocolFlatFee)
func (_ZRC20 *ZRC20Filterer) FilterUpdatedProtocolFlatFee(opts *bind.FilterOpts) (*ZRC20UpdatedProtocolFlatFeeIterator, error) {

	logs, sub, err := _ZRC20.contract.FilterLogs(opts, "UpdatedProtocolFlatFee")
	if err != nil {
		return nil, err
	}
	return &ZRC20UpdatedProtocolFlatFeeIterator{contract: _ZRC20.contract, event: "UpdatedProtocolFlatFee", logs: logs, sub: sub}, nil
}

// WatchUpdatedProtocolFlatFee is a free log subscription operation binding the contract event 0xef13af88e424b5d15f49c77758542c1938b08b8b95b91ed0751f98ba99000d8f.
//
// Solidity: event UpdatedProtocolFlatFee(uint256 protocolFlatFee)
func (_ZRC20 *ZRC20Filterer) WatchUpdatedProtocolFlatFee(opts *bind.WatchOpts, sink chan<- *ZRC20UpdatedProtocolFlatFee) (event.Subscription, error) {

	logs, sub, err := _ZRC20.contract.WatchLogs(opts, "UpdatedProtocolFlatFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZRC20UpdatedProtocolFlatFee)
				if err := _ZRC20.contract.UnpackLog(event, "UpdatedProtocolFlatFee", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdatedProtocolFlatFee is a log parse operation binding the contract event 0xef13af88e424b5d15f49c77758542c1938b08b8b95b91ed0751f98ba99000d8f.
//
// Solidity: event UpdatedProtocolFlatFee(uint256 protocolFlatFee)
func (_ZRC20 *ZRC20Filterer) ParseUpdatedProtocolFlatFee(log types.Log) (*ZRC20UpdatedProtocolFlatFee, error) {
	event := new(ZRC20UpdatedProtocolFlatFee)
	if err := _ZRC20.contract.UnpackLog(event, "UpdatedProtocolFlatFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZRC20UpdatedSystemContractIterator is returned from FilterUpdatedSystemContract and is used to iterate over the raw logs and unpacked data for UpdatedSystemContract events raised by the ZRC20 contract.
type ZRC20UpdatedSystemContractIterator struct {
	Event *ZRC20UpdatedSystemContract // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZRC20UpdatedSystemContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZRC20UpdatedSystemContract)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZRC20UpdatedSystemContract)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZRC20UpdatedSystemContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZRC20UpdatedSystemContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZRC20UpdatedSystemContract represents a UpdatedSystemContract event raised by the ZRC20 contract.
type ZRC20UpdatedSystemContract struct {
	SystemContract common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpdatedSystemContract is a free log retrieval operation binding the contract event 0xd55614e962c5fd6ece71614f6348d702468a997a394dd5e5c1677950226d97ae.
//
// Solidity: event UpdatedSystemContract(address systemContract)
func (_ZRC20 *ZRC20Filterer) FilterUpdatedSystemContract(opts *bind.FilterOpts) (*ZRC20UpdatedSystemContractIterator, error) {

	logs, sub, err := _ZRC20.contract.FilterLogs(opts, "UpdatedSystemContract")
	if err != nil {
		return nil, err
	}
	return &ZRC20UpdatedSystemContractIterator{contract: _ZRC20.contract, event: "UpdatedSystemContract", logs: logs, sub: sub}, nil
}

// WatchUpdatedSystemContract is a free log subscription operation binding the contract event 0xd55614e962c5fd6ece71614f6348d702468a997a394dd5e5c1677950226d97ae.
//
// Solidity: event UpdatedSystemContract(address systemContract)
func (_ZRC20 *ZRC20Filterer) WatchUpdatedSystemContract(opts *bind.WatchOpts, sink chan<- *ZRC20UpdatedSystemContract) (event.Subscription, error) {

	logs, sub, err := _ZRC20.contract.WatchLogs(opts, "UpdatedSystemContract")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZRC20UpdatedSystemContract)
				if err := _ZRC20.contract.UnpackLog(event, "UpdatedSystemContract", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdatedSystemContract is a log parse operation binding the contract event 0xd55614e962c5fd6ece71614f6348d702468a997a394dd5e5c1677950226d97ae.
//
// Solidity: event UpdatedSystemContract(address systemContract)
func (_ZRC20 *ZRC20Filterer) ParseUpdatedSystemContract(log types.Log) (*ZRC20UpdatedSystemContract, error) {
	event := new(ZRC20UpdatedSystemContract)
	if err := _ZRC20.contract.UnpackLog(event, "UpdatedSystemContract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZRC20WithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the ZRC20 contract.
type ZRC20WithdrawalIterator struct {
	Event *ZRC20Withdrawal // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZRC20WithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZRC20Withdrawal)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZRC20Withdrawal)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZRC20WithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZRC20WithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZRC20Withdrawal represents a Withdrawal event raised by the ZRC20 contract.
type ZRC20Withdrawal struct {
	From            common.Address
	To              []byte
	Value           *big.Int
	GasFee          *big.Int
	ProtocolFlatFee *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x9ffbffc04a397460ee1dbe8c9503e098090567d6b7f4b3c02a8617d800b6d955.
//
// Solidity: event Withdrawal(address indexed from, bytes to, uint256 value, uint256 gasFee, uint256 protocolFlatFee)
func (_ZRC20 *ZRC20Filterer) FilterWithdrawal(opts *bind.FilterOpts, from []common.Address) (*ZRC20WithdrawalIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _ZRC20.contract.FilterLogs(opts, "Withdrawal", fromRule)
	if err != nil {
		return nil, err
	}
	return &ZRC20WithdrawalIterator{contract: _ZRC20.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x9ffbffc04a397460ee1dbe8c9503e098090567d6b7f4b3c02a8617d800b6d955.
//
// Solidity: event Withdrawal(address indexed from, bytes to, uint256 value, uint256 gasFee, uint256 protocolFlatFee)
func (_ZRC20 *ZRC20Filterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *ZRC20Withdrawal, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _ZRC20.contract.WatchLogs(opts, "Withdrawal", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZRC20Withdrawal)
				if err := _ZRC20.contract.UnpackLog(event, "Withdrawal", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawal is a log parse operation binding the contract event 0x9ffbffc04a397460ee1dbe8c9503e098090567d6b7f4b3c02a8617d800b6d955.
//
// Solidity: event Withdrawal(address indexed from, bytes to, uint256 value, uint256 gasFee, uint256 protocolFlatFee)
func (_ZRC20 *ZRC20Filterer) ParseWithdrawal(log types.Log) (*ZRC20Withdrawal, error) {
	event := new(ZRC20Withdrawal)
	if err := _ZRC20.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
