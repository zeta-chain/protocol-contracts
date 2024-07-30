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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"name_\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol_\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals_\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"chainid_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"coinType_\",\"type\":\"uint8\",\"internalType\":\"enumCoinType\"},{\"name\":\"gasLimit_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"systemContractAddress_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gatewayContractAddress_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"CHAIN_ID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"COIN_TYPE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumCoinType\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"FUNGIBLE_MODULE_ADDRESS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GAS_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GATEWAY_CONTRACT_ADDRESS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PROTOCOL_FLAT_FEE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SYSTEM_CONTRACT_ADDRESS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateGasLimit\",\"inputs\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateProtocolFlatFee\",\"inputs\":[{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateSystemContractAddress\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"to\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawGasFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"from\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedGasLimit\",\"inputs\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedProtocolFlatFee\",\"inputs\":[{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedSystemContract\",\"inputs\":[{\"name\":\"systemContract\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawal\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CallerIsNotFungibleModule\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GasFeeTransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LowAllowance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LowBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroGasCoin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroGasPrice\",\"inputs\":[]}]",
	Bin: "0x60c060405234801561001057600080fd5b506040516119e83803806119e883398101604081905261002f916101db565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461006357604051632b2add3d60e01b815260040160405180910390fd5b600761006f898261032d565b50600861007c888261032d565b506009805460ff191660ff881617905560808590528360028111156100a3576100a36103eb565b60a08160028111156100b7576100b76103eb565b905250600292909255600080546001600160a01b039283166001600160a01b03199182161790915560018054929093169116179055506104019350505050565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261011e57600080fd5b81516001600160401b03811115610137576101376100f7565b604051601f8201601f19908116603f011681016001600160401b0381118282101715610165576101656100f7565b60405281815283820160200185101561017d57600080fd5b60005b8281101561019c57602081860181015183830182015201610180565b506000918101602001919091529392505050565b8051600381106101bf57600080fd5b919050565b80516001600160a01b03811681146101bf57600080fd5b600080600080600080600080610100898b0312156101f857600080fd5b88516001600160401b0381111561020e57600080fd5b61021a8b828c0161010d565b60208b015190995090506001600160401b0381111561023857600080fd5b6102448b828c0161010d565b975050604089015160ff8116811461025b57600080fd5b60608a0151909650945061027160808a016101b0565b60a08a0151909450925061028760c08a016101c4565b915061029560e08a016101c4565b90509295985092959890939650565b600181811c908216806102b857607f821691505b6020821081036102d857634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561032857806000526020600020601f840160051c810160208510156103055750805b601f840160051c820191505b818110156103255760008155600101610311565b50505b505050565b81516001600160401b03811115610346576103466100f7565b61035a8161035484546102a4565b846102de565b6020601f82116001811461038e57600083156103765750848201515b600019600385901b1c1916600184901b178455610325565b600084815260208120601f198516915b828110156103be578785015182556020948501946001909201910161039e565b50848210156103dc5786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b634e487b7160e01b600052602160045260246000fd5b60805160a0516115b461043460003960006102f30152600081816102c4015281816109430152610a4901526115b46000f3fe608060405234801561001057600080fd5b506004361061018d5760003560e01c806385e1f4d0116100e3578063d9eeebed1161008c578063eddeb12311610066578063eddeb123146103f7578063f2441b321461040a578063f687d12a1461042a57600080fd5b8063d9eeebed1461035d578063dd62ed3e14610391578063e2f535b8146103d757600080fd5b8063a9059cbb116100bd578063a9059cbb14610322578063c701262614610335578063c835d7cc1461034857600080fd5b806385e1f4d0146102bf57806395d89b41146102e6578063a3413d03146102ee57600080fd5b8063313ce5671161014557806347e7ef241161011f57806347e7ef241461026d5780634d8943bb1461028057806370a082311461028957600080fd5b8063313ce567146102055780633ce4a5bc1461021a57806342966c681461025a57600080fd5b8063095ea7b311610176578063095ea7b3146101c757806318160ddd146101ea57806323b872dd146101f257600080fd5b806306fdde0314610192578063091d2788146101b0575b600080fd5b61019a61043d565b6040516101a79190611193565b60405180910390f35b6101b960025481565b6040519081526020016101a7565b6101da6101d53660046111d2565b6104cf565b60405190151581526020016101a7565b6006546101b9565b6101da6102003660046111fe565b6104e6565b60095460405160ff90911681526020016101a7565b61023573735b14bb79463307aacbed86daf3322b1e6226ab81565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101a7565b6101da61026836600461123f565b61057d565b6101da61027b3660046111d2565b610591565b6101b960035481565b6101b9610297366004611258565b73ffffffffffffffffffffffffffffffffffffffff1660009081526004602052604090205490565b6101b97f000000000000000000000000000000000000000000000000000000000000000081565b61019a6106e5565b6103157f000000000000000000000000000000000000000000000000000000000000000081565b6040516101a79190611275565b6101da6103303660046111d2565b6106f4565b6101da6103433660046112e5565b610701565b61035b610356366004611258565b610850565b005b610365610917565b6040805173ffffffffffffffffffffffffffffffffffffffff90931683526020830191909152016101a7565b6101b961039f3660046113dd565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260056020908152604080832093909416825291909152205490565b6001546102359073ffffffffffffffffffffffffffffffffffffffff1681565b61035b61040536600461123f565b610b35565b6000546102359073ffffffffffffffffffffffffffffffffffffffff1681565b61035b61043836600461123f565b610bb7565b60606007805461044c90611416565b80601f016020809104026020016040519081016040528092919081815260200182805461047890611416565b80156104c55780601f1061049a576101008083540402835291602001916104c5565b820191906000526020600020905b8154815290600101906020018083116104a857829003601f168201915b5050505050905090565b60006104dc338484610c39565b5060015b92915050565b60006104f3848484610d42565b73ffffffffffffffffffffffffffffffffffffffff841660009081526005602090815260408083203384529091529020548281101561055e576040517f10bad14700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610572853361056d8685611498565b610c39565b506001949350505050565b60006105893383610efd565b506001919050565b60003373735b14bb79463307aacbed86daf3322b1e6226ab148015906105cf575060005473ffffffffffffffffffffffffffffffffffffffff163314155b80156105f3575060015473ffffffffffffffffffffffffffffffffffffffff163314155b1561062a576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610634838361103f565b6040517f735b14bb79463307aacbed86daf3322b1e6226ab000000000000000000000000602082015273ffffffffffffffffffffffffffffffffffffffff8416907f67fc7bdaed5b0ec550d8706b87d60568ab70c6b781263c70101d54cd1564aab390603401604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290526106d49186906114ab565b60405180910390a250600192915050565b60606008805461044c90611416565b60006104dc338484610d42565b600080600061070e610917565b6040517f23b872dd00000000000000000000000000000000000000000000000000000000815233600482015273735b14bb79463307aacbed86daf3322b1e6226ab602482015260448101829052919350915073ffffffffffffffffffffffffffffffffffffffff8316906323b872dd906064016020604051808303816000875af11580156107a0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107c491906114cd565b6107fa576040517f0a7cd6d600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6108043385610efd565b60035460405133917f9ffbffc04a397460ee1dbe8c9503e098090567d6b7f4b3c02a8617d800b6d9559161083d918991899187916114ef565b60405180910390a2506001949350505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461089d576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fd55614e962c5fd6ece71614f6348d702468a997a394dd5e5c1677950226d97ae906020015b60405180910390a150565b600080546040517f0be155470000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060048201528291829173ffffffffffffffffffffffffffffffffffffffff90911690630be1554790602401602060405180830381865afa1580156109aa573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109ce919061151e565b905073ffffffffffffffffffffffffffffffffffffffff8116610a1d576040517f78fff39600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080546040517fd7fd7afb0000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000600482015273ffffffffffffffffffffffffffffffffffffffff9091169063d7fd7afb90602401602060405180830381865afa158015610aac573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ad0919061153b565b905080600003610b0c576040517fe661aed000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060035460025483610b1f9190611554565b610b29919061156b565b92959294509192505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610b82576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60038190556040518181527fef13af88e424b5d15f49c77758542c1938b08b8b95b91ed0751f98ba99000d8f9060200161090c565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610c04576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60028190556040518181527fff5788270f43bfc1ca41c503606d2594aa3023a1a7547de403a3e2f146a4a80a9060200161090c565b73ffffffffffffffffffffffffffffffffffffffff8316610c86576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216610cd3576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff83811660008181526005602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8316610d8f576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216610ddc576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff831660009081526004602052604090205481811015610e3c576040517ffe382aa700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610e468282611498565b73ffffffffffffffffffffffffffffffffffffffff8086166000908152600460205260408082209390935590851681529081208054849290610e8990849061156b565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610eef91815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff8216610f4a576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff821660009081526004602052604090205481811015610faa576040517ffe382aa700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610fb48282611498565b73ffffffffffffffffffffffffffffffffffffffff841660009081526004602052604081209190915560068054849290610fef908490611498565b909155505060405182815260009073ffffffffffffffffffffffffffffffffffffffff8516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef90602001610d35565b73ffffffffffffffffffffffffffffffffffffffff821661108c576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806006600082825461109e919061156b565b909155505073ffffffffffffffffffffffffffffffffffffffff8216600090815260046020526040812080548392906110d890849061156b565b909155505060405181815273ffffffffffffffffffffffffffffffffffffffff8316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b6000815180845260005b8181101561115557602081850181015186830182015201611139565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b6020815260006111a6602083018461112f565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff811681146111cf57600080fd5b50565b600080604083850312156111e557600080fd5b82356111f0816111ad565b946020939093013593505050565b60008060006060848603121561121357600080fd5b833561121e816111ad565b9250602084013561122e816111ad565b929592945050506040919091013590565b60006020828403121561125157600080fd5b5035919050565b60006020828403121561126a57600080fd5b81356111a6816111ad565b60208101600383106112b0577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080604083850312156112f857600080fd5b823567ffffffffffffffff81111561130f57600080fd5b8301601f8101851361132057600080fd5b803567ffffffffffffffff81111561133a5761133a6112b6565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff821117156113a6576113a66112b6565b6040528181528282016020018710156113be57600080fd5b8160208401602083013760006020928201830152969401359450505050565b600080604083850312156113f057600080fd5b82356113fb816111ad565b9150602083013561140b816111ad565b809150509250929050565b600181811c9082168061142a57607f821691505b602082108103611463577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b818103818111156104e0576104e0611469565b6040815260006114be604083018561112f565b90508260208301529392505050565b6000602082840312156114df57600080fd5b815180151581146111a657600080fd5b608081526000611502608083018761112f565b6020830195909552506040810192909252606090910152919050565b60006020828403121561153057600080fd5b81516111a6816111ad565b60006020828403121561154d57600080fd5b5051919050565b80820281158282048414176104e0576104e0611469565b808201808211156104e0576104e061146956fea264697066735822122062e8c81bc642fb8e9a1eaabde5abdece7e58510051b4aa2bee706dc57435041764736f6c634300081a0033",
}

// ZRC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use ZRC20MetaData.ABI instead.
var ZRC20ABI = ZRC20MetaData.ABI

// ZRC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZRC20MetaData.Bin instead.
var ZRC20Bin = ZRC20MetaData.Bin

// DeployZRC20 deploys a new Ethereum contract, binding an instance of ZRC20 to it.
func DeployZRC20(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string, decimals_ uint8, chainid_ *big.Int, coinType_ uint8, gasLimit_ *big.Int, systemContractAddress_ common.Address, gatewayContractAddress_ common.Address) (common.Address, *types.Transaction, *ZRC20, error) {
	parsed, err := ZRC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZRC20Bin), backend, name_, symbol_, decimals_, chainid_, coinType_, gasLimit_, systemContractAddress_, gatewayContractAddress_)
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

// GATEWAYCONTRACTADDRESS is a free data retrieval call binding the contract method 0xe2f535b8.
//
// Solidity: function GATEWAY_CONTRACT_ADDRESS() view returns(address)
func (_ZRC20 *ZRC20Caller) GATEWAYCONTRACTADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZRC20.contract.Call(opts, &out, "GATEWAY_CONTRACT_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GATEWAYCONTRACTADDRESS is a free data retrieval call binding the contract method 0xe2f535b8.
//
// Solidity: function GATEWAY_CONTRACT_ADDRESS() view returns(address)
func (_ZRC20 *ZRC20Session) GATEWAYCONTRACTADDRESS() (common.Address, error) {
	return _ZRC20.Contract.GATEWAYCONTRACTADDRESS(&_ZRC20.CallOpts)
}

// GATEWAYCONTRACTADDRESS is a free data retrieval call binding the contract method 0xe2f535b8.
//
// Solidity: function GATEWAY_CONTRACT_ADDRESS() view returns(address)
func (_ZRC20 *ZRC20CallerSession) GATEWAYCONTRACTADDRESS() (common.Address, error) {
	return _ZRC20.Contract.GATEWAYCONTRACTADDRESS(&_ZRC20.CallOpts)
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
// Solidity: function updateGasLimit(uint256 gasLimit) returns()
func (_ZRC20 *ZRC20Transactor) UpdateGasLimit(opts *bind.TransactOpts, gasLimit *big.Int) (*types.Transaction, error) {
	return _ZRC20.contract.Transact(opts, "updateGasLimit", gasLimit)
}

// UpdateGasLimit is a paid mutator transaction binding the contract method 0xf687d12a.
//
// Solidity: function updateGasLimit(uint256 gasLimit) returns()
func (_ZRC20 *ZRC20Session) UpdateGasLimit(gasLimit *big.Int) (*types.Transaction, error) {
	return _ZRC20.Contract.UpdateGasLimit(&_ZRC20.TransactOpts, gasLimit)
}

// UpdateGasLimit is a paid mutator transaction binding the contract method 0xf687d12a.
//
// Solidity: function updateGasLimit(uint256 gasLimit) returns()
func (_ZRC20 *ZRC20TransactorSession) UpdateGasLimit(gasLimit *big.Int) (*types.Transaction, error) {
	return _ZRC20.Contract.UpdateGasLimit(&_ZRC20.TransactOpts, gasLimit)
}

// UpdateProtocolFlatFee is a paid mutator transaction binding the contract method 0xeddeb123.
//
// Solidity: function updateProtocolFlatFee(uint256 protocolFlatFee) returns()
func (_ZRC20 *ZRC20Transactor) UpdateProtocolFlatFee(opts *bind.TransactOpts, protocolFlatFee *big.Int) (*types.Transaction, error) {
	return _ZRC20.contract.Transact(opts, "updateProtocolFlatFee", protocolFlatFee)
}

// UpdateProtocolFlatFee is a paid mutator transaction binding the contract method 0xeddeb123.
//
// Solidity: function updateProtocolFlatFee(uint256 protocolFlatFee) returns()
func (_ZRC20 *ZRC20Session) UpdateProtocolFlatFee(protocolFlatFee *big.Int) (*types.Transaction, error) {
	return _ZRC20.Contract.UpdateProtocolFlatFee(&_ZRC20.TransactOpts, protocolFlatFee)
}

// UpdateProtocolFlatFee is a paid mutator transaction binding the contract method 0xeddeb123.
//
// Solidity: function updateProtocolFlatFee(uint256 protocolFlatFee) returns()
func (_ZRC20 *ZRC20TransactorSession) UpdateProtocolFlatFee(protocolFlatFee *big.Int) (*types.Transaction, error) {
	return _ZRC20.Contract.UpdateProtocolFlatFee(&_ZRC20.TransactOpts, protocolFlatFee)
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
