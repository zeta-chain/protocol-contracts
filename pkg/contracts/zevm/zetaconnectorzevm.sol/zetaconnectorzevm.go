// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zetaconnectorzevm

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

// ZetaInterfacesSendInput is an auto generated low-level Go binding around an user-defined struct.
type ZetaInterfacesSendInput struct {
	DestinationChainId  *big.Int
	DestinationAddress  []byte
	DestinationGasLimit *big.Int
	Message             []byte
	ZetaValueAndGas     *big.Int
	ZetaParams          []byte
}

// ZetaConnectorZEVMMetaData contains all meta data concerning the ZetaConnectorZEVM contract.
var ZetaConnectorZEVMMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"zetaTokenAddress_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tssAddress_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tssAddressUpdater_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pauserAddress_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerIsNotPauser\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerIsNotTss\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerIsNotTssOrUpdater\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerIsNotTssUpdater\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"}],\"name\":\"ExceedsMaxSupply\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedZetaSent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyFungibleModule\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyWZETA\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WZETATransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZetaTransferError\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"callerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTssAddress\",\"type\":\"address\"}],\"name\":\"PauserAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"callerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTssAddress\",\"type\":\"address\"}],\"name\":\"TSSAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"callerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTssUpdaterAddress\",\"type\":\"address\"}],\"name\":\"TSSAddressUpdaterUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"zetaTxSenderAddress\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"zetaValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"internalSendHash\",\"type\":\"bytes32\"}],\"name\":\"ZetaReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"zetaTxSenderAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"destinationAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remainingZetaValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"internalSendHash\",\"type\":\"bytes32\"}],\"name\":\"ZetaReverted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sourceTxOriginAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"zetaTxSenderAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"destinationAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"zetaValueAndGas\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destinationGasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"zetaParams\",\"type\":\"bytes\"}],\"name\":\"ZetaSent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FUNGIBLE_MODULE_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"zetaTxSenderAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"zetaValue\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"internalSendHash\",\"type\":\"bytes32\"}],\"name\":\"onReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"zetaTxSenderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"destinationAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainingZetaValue\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"internalSendHash\",\"type\":\"bytes32\"}],\"name\":\"onRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauserAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceTssAddressUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"destinationAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"destinationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"zetaValueAndGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"zetaParams\",\"type\":\"bytes\"}],\"internalType\":\"structZetaInterfaces.SendInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tssAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tssAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pauserAddress_\",\"type\":\"address\"}],\"name\":\"updatePauserAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tssAddress_\",\"type\":\"address\"}],\"name\":\"updateTssAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zetaToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b506040516200238338038062002383833981810160405281019062000037919062000284565b8383838360008060006101000a81548160ff021916908315150217905550600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161480620000bd5750600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16145b80620000f55750600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16145b806200012d5750600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b1562000165576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8373ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1660601b8152505082600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600060016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050505050505062000349565b6000815190506200027e816200032f565b92915050565b60008060008060808587031215620002a157620002a06200032a565b5b6000620002b1878288016200026d565b9450506020620002c4878288016200026d565b9350506040620002d7878288016200026d565b9250506060620002ea878288016200026d565b91505092959194509250565b600062000303826200030a565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b6200033a81620002f6565b81146200034657600080fd5b50565b60805160601c611feb620003986000396000818160e80152818161038801528181610425015281816104a601528181610e7f01528181610f000152818161117f01526112680152611feb6000f3fe6080604052600436106100e15760003560e01c80636128480f1161007f5780639122c344116100595780639122c344146102e0578063942a5e1614610309578063ec02690114610332578063f7fb869b1461035b5761016d565b80636128480f14610289578063779e3b63146102b25780638456cb59146102c95761016d565b80633ce4a5bc116100bb5780633ce4a5bc146101f15780633f4ba83a1461021c5780635b112591146102335780635c975abb1461025e5761016d565b806321e093b11461017257806329dd214d1461019d578063328a01d0146101c65761016d565b3661016d577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461016b576040517f6e6b6de700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b005b600080fd5b34801561017e57600080fd5b50610187610386565b6040516101949190611b27565b60405180910390f35b3480156101a957600080fd5b506101c460048036038101906101bf91906117d3565b6103aa565b005b3480156101d257600080fd5b506101db610727565b6040516101e89190611b27565b60405180910390f35b3480156101fd57600080fd5b5061020661074d565b6040516102139190611b27565b60405180910390f35b34801561022857600080fd5b50610231610765565b005b34801561023f57600080fd5b50610248610801565b6040516102559190611b27565b60405180910390f35b34801561026a57600080fd5b50610273610827565b6040516102809190611c83565b60405180910390f35b34801561029557600080fd5b506102b060048036038101906102ab9190611697565b61083d565b005b3480156102be57600080fd5b506102c76109b3565b005b3480156102d557600080fd5b506102de610b8e565b005b3480156102ec57600080fd5b5061030760048036038101906103029190611697565b610c2a565b005b34801561031557600080fd5b50610330600480360381019061032b91906116c4565b610dfc565b005b34801561033e57600080fd5b50610359600480360381019061035491906118a2565b611175565b005b34801561036757600080fd5b50610370611449565b60405161037d9190611b27565b60405180910390f35b7f000000000000000000000000000000000000000000000000000000000000000081565b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610423576040517fea02b3f300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663d0e30db0856040518263ffffffff1660e01b81526004016000604051808303818588803b15801561048b57600080fd5b505af115801561049f573d6000803e3d6000fd5b50505050507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166323b872dd3087876040518463ffffffff1660e01b815260040161050193929190611b6b565b602060405180830381600087803b15801561051b57600080fd5b505af115801561052f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061055391906117a6565b610589576040517fa8c6fd4a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008383905011156106c5578473ffffffffffffffffffffffffffffffffffffffff16633749c51a6040518060a001604052808b8b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505081526020018981526020018873ffffffffffffffffffffffffffffffffffffffff16815260200187815260200186868080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508152506040518263ffffffff1660e01b81526004016106929190611d27565b600060405180830381600087803b1580156106ac57600080fd5b505af11580156106c0573d6000803e3d6000fd5b505050505b808573ffffffffffffffffffffffffffffffffffffffff16877ff1302855733b40d8acb467ee990b6d56c05c80e28ebcabfa6e6f3f57cb50d6988b8b898989604051610715959493929190611c9e565b60405180910390a45050505050505050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b73735b14bb79463307aacbed86daf3322b1e6226ab81565b600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146107f757336040517f4677a0d30000000000000000000000000000000000000000000000000000000081526004016107ee9190611b27565b60405180910390fd5b6107ff61146f565b565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060009054906101000a900460ff16905090565b600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146108cf57336040517f4677a0d30000000000000000000000000000000000000000000000000000000081526004016108c69190611b27565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610936576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600060016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fd41d83655d484bdf299598751c371b2d92088667266fe3774b25a97bdd5d039733826040516109a8929190611b42565b60405180910390a150565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a4557336040517fe700765e000000000000000000000000000000000000000000000000000000008152600401610a3c9190611b27565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415610ace576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f5104c9abdc7d111c2aeb4ce890ac70274a4be2ee83f46a62551be5e6ebc82dd033600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051610b84929190611b42565b60405180910390a1565b600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610c2057336040517f4677a0d3000000000000000000000000000000000000000000000000000000008152600401610c179190611b27565b60405180910390fd5b610c286114d1565b565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614158015610cd65750600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614155b15610d1857336040517fcdfcef97000000000000000000000000000000000000000000000000000000008152600401610d0f9190611b27565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610d7f576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fe79965b5c67dcfb2cf5fe152715e4a7256cee62a3d5dd8484fd8a8539eb8beff3382604051610df1929190611b42565b60405180910390a150565b610e04611533565b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610e7d576040517fea02b3f300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663d0e30db0856040518263ffffffff1660e01b81526004016000604051808303818588803b158015610ee557600080fd5b505af1158015610ef9573d6000803e3d6000fd5b50505050507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166323b872dd308b876040518463ffffffff1660e01b8152600401610f5b93929190611b6b565b602060405180830381600087803b158015610f7557600080fd5b505af1158015610f89573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fad91906117a6565b610fe3576040517fa8c6fd4a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000838390501115611125578873ffffffffffffffffffffffffffffffffffffffff16633ff0693c6040518060c001604052808c73ffffffffffffffffffffffffffffffffffffffff1681526020018b81526020018a8a8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050815260200188815260200187815260200186868080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508152506040518263ffffffff1660e01b81526004016110f29190611d49565b600060405180830381600087803b15801561110c57600080fd5b505af1158015611120573d6000803e3d6000fd5b505050505b80857f521fb0b407c2eb9b1375530e9b9a569889992140a688bc076aa72c1712012c888b8b8b8b8a8a8a6040516111629796959493929190611c1e565b60405180910390a3505050505050505050565b61117d611533565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166323b872dd333084608001356040518463ffffffff1660e01b81526004016111de93929190611b6b565b602060405180830381600087803b1580156111f857600080fd5b505af115801561120c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061123091906117a6565b611266576040517fa8c6fd4a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16632e1a7d4d82608001356040518263ffffffff1660e01b81526004016112c39190611d6b565b600060405180830381600087803b1580156112dd57600080fd5b505af11580156112f1573d6000803e3d6000fd5b50505050600073735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff16826080013560405161133390611b12565b60006040518083038185875af1925050503d8060008114611370576040519150601f19603f3d011682016040523d82523d6000602084013e611375565b606091505b50509050806113b0576040517fc7ffc47b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81600001353373ffffffffffffffffffffffffffffffffffffffff167f7ec1c94701e09b1652f3e1d307e60c4b9ebf99aff8c2079fd1d8c585e031c4e4328580602001906113fe9190611d86565b876080013588604001358980606001906114189190611d86565b8b8060a001906114289190611d86565b60405161143d99989796959493929190611ba2565b60405180910390a35050565b600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b61147761157d565b60008060006101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6114ba6115c6565b6040516114c79190611b27565b60405180910390a1565b6114d9611533565b60016000806101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25861151c6115c6565b6040516115299190611b27565b60405180910390a1565b61153b610827565b1561157b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161157290611d07565b60405180910390fd5b565b611585610827565b6115c4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115bb90611ce7565b60405180910390fd5b565b600033905090565b6000813590506115dd81611f59565b92915050565b6000815190506115f281611f70565b92915050565b60008135905061160781611f87565b92915050565b60008083601f84011261162357611622611ecb565b5b8235905067ffffffffffffffff8111156116405761163f611ec6565b5b60208301915083600182028301111561165c5761165b611edf565b5b9250929050565b600060c0828403121561167957611678611ed5565b5b81905092915050565b60008135905061169181611f9e565b92915050565b6000602082840312156116ad576116ac611eee565b5b60006116bb848285016115ce565b91505092915050565b600080600080600080600080600060e08a8c0312156116e6576116e5611eee565b5b60006116f48c828d016115ce565b99505060206117058c828d01611682565b98505060408a013567ffffffffffffffff81111561172657611725611ee9565b5b6117328c828d0161160d565b975097505060606117458c828d01611682565b95505060806117568c828d01611682565b94505060a08a013567ffffffffffffffff81111561177757611776611ee9565b5b6117838c828d0161160d565b935093505060c06117968c828d016115f8565b9150509295985092959850929598565b6000602082840312156117bc576117bb611eee565b5b60006117ca848285016115e3565b91505092915050565b60008060008060008060008060c0898b0312156117f3576117f2611eee565b5b600089013567ffffffffffffffff81111561181157611810611ee9565b5b61181d8b828c0161160d565b985098505060206118308b828c01611682565b96505060406118418b828c016115ce565b95505060606118528b828c01611682565b945050608089013567ffffffffffffffff81111561187357611872611ee9565b5b61187f8b828c0161160d565b935093505060a06118928b828c016115f8565b9150509295985092959890939650565b6000602082840312156118b8576118b7611eee565b5b600082013567ffffffffffffffff8111156118d6576118d5611ee9565b5b6118e284828501611663565b91505092915050565b6118f481611e32565b82525050565b61190381611e32565b82525050565b61191281611e44565b82525050565b60006119248385611e05565b9350611931838584611e84565b61193a83611ef3565b840190509392505050565b600061195082611de9565b61195a8185611df4565b935061196a818560208601611e93565b61197381611ef3565b840191505092915050565b600061198b601483611e21565b915061199682611f04565b602082019050919050565b60006119ae601083611e21565b91506119b982611f2d565b602082019050919050565b60006119d1600083611e16565b91506119dc82611f56565b600082019050919050565b600060a0830160008301518482036000860152611a048282611945565b9150506020830151611a196020860182611af4565b506040830151611a2c60408601826118eb565b506060830151611a3f6060860182611af4565b5060808301518482036080860152611a578282611945565b9150508091505092915050565b600060c083016000830151611a7c60008601826118eb565b506020830151611a8f6020860182611af4565b5060408301518482036040860152611aa78282611945565b9150506060830151611abc6060860182611af4565b506080830151611acf6080860182611af4565b5060a083015184820360a0860152611ae78282611945565b9150508091505092915050565b611afd81611e7a565b82525050565b611b0c81611e7a565b82525050565b6000611b1d826119c4565b9150819050919050565b6000602082019050611b3c60008301846118fa565b92915050565b6000604082019050611b5760008301856118fa565b611b6460208301846118fa565b9392505050565b6000606082019050611b8060008301866118fa565b611b8d60208301856118fa565b611b9a6040830184611b03565b949350505050565b600060c082019050611bb7600083018c6118fa565b8181036020830152611bca818a8c611918565b9050611bd96040830189611b03565b611be66060830188611b03565b8181036080830152611bf9818688611918565b905081810360a0830152611c0e818486611918565b90509a9950505050505050505050565b600060a082019050611c33600083018a6118fa565b611c406020830189611b03565b8181036040830152611c53818789611918565b9050611c626060830186611b03565b8181036080830152611c75818486611918565b905098975050505050505050565b6000602082019050611c986000830184611909565b92915050565b60006060820190508181036000830152611cb9818789611918565b9050611cc86020830186611b03565b8181036040830152611cdb818486611918565b90509695505050505050565b60006020820190508181036000830152611d008161197e565b9050919050565b60006020820190508181036000830152611d20816119a1565b9050919050565b60006020820190508181036000830152611d4181846119e7565b905092915050565b60006020820190508181036000830152611d638184611a64565b905092915050565b6000602082019050611d806000830184611b03565b92915050565b60008083356001602003843603038112611da357611da2611eda565b5b80840192508235915067ffffffffffffffff821115611dc557611dc4611ed0565b5b602083019250600182023603831315611de157611de0611ee4565b5b509250929050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b6000611e3d82611e5a565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b83811015611eb1578082015181840152602081019050611e96565b83811115611ec0576000848401525b50505050565b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f5061757361626c653a206e6f7420706175736564000000000000000000000000600082015250565b7f5061757361626c653a2070617573656400000000000000000000000000000000600082015250565b50565b611f6281611e32565b8114611f6d57600080fd5b50565b611f7981611e44565b8114611f8457600080fd5b50565b611f9081611e50565b8114611f9b57600080fd5b50565b611fa781611e7a565b8114611fb257600080fd5b5056fea2646970667358221220e26d28d0e1a524eee4ee8ed3c5fafc87b2b8e13ef2506abc1564a0edce9d0be764736f6c63430008070033",
}

// ZetaConnectorZEVMABI is the input ABI used to generate the binding from.
// Deprecated: Use ZetaConnectorZEVMMetaData.ABI instead.
var ZetaConnectorZEVMABI = ZetaConnectorZEVMMetaData.ABI

// ZetaConnectorZEVMBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZetaConnectorZEVMMetaData.Bin instead.
var ZetaConnectorZEVMBin = ZetaConnectorZEVMMetaData.Bin

// DeployZetaConnectorZEVM deploys a new Ethereum contract, binding an instance of ZetaConnectorZEVM to it.
func DeployZetaConnectorZEVM(auth *bind.TransactOpts, backend bind.ContractBackend, zetaTokenAddress_ common.Address, tssAddress_ common.Address, tssAddressUpdater_ common.Address, pauserAddress_ common.Address) (common.Address, *types.Transaction, *ZetaConnectorZEVM, error) {
	parsed, err := ZetaConnectorZEVMMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZetaConnectorZEVMBin), backend, zetaTokenAddress_, tssAddress_, tssAddressUpdater_, pauserAddress_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZetaConnectorZEVM{ZetaConnectorZEVMCaller: ZetaConnectorZEVMCaller{contract: contract}, ZetaConnectorZEVMTransactor: ZetaConnectorZEVMTransactor{contract: contract}, ZetaConnectorZEVMFilterer: ZetaConnectorZEVMFilterer{contract: contract}}, nil
}

// ZetaConnectorZEVM is an auto generated Go binding around an Ethereum contract.
type ZetaConnectorZEVM struct {
	ZetaConnectorZEVMCaller     // Read-only binding to the contract
	ZetaConnectorZEVMTransactor // Write-only binding to the contract
	ZetaConnectorZEVMFilterer   // Log filterer for contract events
}

// ZetaConnectorZEVMCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZetaConnectorZEVMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorZEVMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZetaConnectorZEVMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorZEVMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZetaConnectorZEVMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorZEVMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZetaConnectorZEVMSession struct {
	Contract     *ZetaConnectorZEVM // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ZetaConnectorZEVMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZetaConnectorZEVMCallerSession struct {
	Contract *ZetaConnectorZEVMCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ZetaConnectorZEVMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZetaConnectorZEVMTransactorSession struct {
	Contract     *ZetaConnectorZEVMTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ZetaConnectorZEVMRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZetaConnectorZEVMRaw struct {
	Contract *ZetaConnectorZEVM // Generic contract binding to access the raw methods on
}

// ZetaConnectorZEVMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZetaConnectorZEVMCallerRaw struct {
	Contract *ZetaConnectorZEVMCaller // Generic read-only contract binding to access the raw methods on
}

// ZetaConnectorZEVMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZetaConnectorZEVMTransactorRaw struct {
	Contract *ZetaConnectorZEVMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZetaConnectorZEVM creates a new instance of ZetaConnectorZEVM, bound to a specific deployed contract.
func NewZetaConnectorZEVM(address common.Address, backend bind.ContractBackend) (*ZetaConnectorZEVM, error) {
	contract, err := bindZetaConnectorZEVM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorZEVM{ZetaConnectorZEVMCaller: ZetaConnectorZEVMCaller{contract: contract}, ZetaConnectorZEVMTransactor: ZetaConnectorZEVMTransactor{contract: contract}, ZetaConnectorZEVMFilterer: ZetaConnectorZEVMFilterer{contract: contract}}, nil
}

// NewZetaConnectorZEVMCaller creates a new read-only instance of ZetaConnectorZEVM, bound to a specific deployed contract.
func NewZetaConnectorZEVMCaller(address common.Address, caller bind.ContractCaller) (*ZetaConnectorZEVMCaller, error) {
	contract, err := bindZetaConnectorZEVM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorZEVMCaller{contract: contract}, nil
}

// NewZetaConnectorZEVMTransactor creates a new write-only instance of ZetaConnectorZEVM, bound to a specific deployed contract.
func NewZetaConnectorZEVMTransactor(address common.Address, transactor bind.ContractTransactor) (*ZetaConnectorZEVMTransactor, error) {
	contract, err := bindZetaConnectorZEVM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorZEVMTransactor{contract: contract}, nil
}

// NewZetaConnectorZEVMFilterer creates a new log filterer instance of ZetaConnectorZEVM, bound to a specific deployed contract.
func NewZetaConnectorZEVMFilterer(address common.Address, filterer bind.ContractFilterer) (*ZetaConnectorZEVMFilterer, error) {
	contract, err := bindZetaConnectorZEVM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorZEVMFilterer{contract: contract}, nil
}

// bindZetaConnectorZEVM binds a generic wrapper to an already deployed contract.
func bindZetaConnectorZEVM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZetaConnectorZEVMMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaConnectorZEVM *ZetaConnectorZEVMRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaConnectorZEVM.Contract.ZetaConnectorZEVMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaConnectorZEVM *ZetaConnectorZEVMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.Contract.ZetaConnectorZEVMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaConnectorZEVM *ZetaConnectorZEVMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.Contract.ZetaConnectorZEVMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaConnectorZEVM *ZetaConnectorZEVMCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaConnectorZEVM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaConnectorZEVM *ZetaConnectorZEVMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaConnectorZEVM *ZetaConnectorZEVMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.Contract.contract.Transact(opts, method, params...)
}

// FUNGIBLEMODULEADDRESS is a free data retrieval call binding the contract method 0x3ce4a5bc.
//
// Solidity: function FUNGIBLE_MODULE_ADDRESS() view returns(address)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMCaller) FUNGIBLEMODULEADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorZEVM.contract.Call(opts, &out, "FUNGIBLE_MODULE_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FUNGIBLEMODULEADDRESS is a free data retrieval call binding the contract method 0x3ce4a5bc.
//
// Solidity: function FUNGIBLE_MODULE_ADDRESS() view returns(address)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMSession) FUNGIBLEMODULEADDRESS() (common.Address, error) {
	return _ZetaConnectorZEVM.Contract.FUNGIBLEMODULEADDRESS(&_ZetaConnectorZEVM.CallOpts)
}

// FUNGIBLEMODULEADDRESS is a free data retrieval call binding the contract method 0x3ce4a5bc.
//
// Solidity: function FUNGIBLE_MODULE_ADDRESS() view returns(address)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMCallerSession) FUNGIBLEMODULEADDRESS() (common.Address, error) {
	return _ZetaConnectorZEVM.Contract.FUNGIBLEMODULEADDRESS(&_ZetaConnectorZEVM.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ZetaConnectorZEVM.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMSession) Paused() (bool, error) {
	return _ZetaConnectorZEVM.Contract.Paused(&_ZetaConnectorZEVM.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMCallerSession) Paused() (bool, error) {
	return _ZetaConnectorZEVM.Contract.Paused(&_ZetaConnectorZEVM.CallOpts)
}

// PauserAddress is a free data retrieval call binding the contract method 0xf7fb869b.
//
// Solidity: function pauserAddress() view returns(address)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMCaller) PauserAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorZEVM.contract.Call(opts, &out, "pauserAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserAddress is a free data retrieval call binding the contract method 0xf7fb869b.
//
// Solidity: function pauserAddress() view returns(address)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMSession) PauserAddress() (common.Address, error) {
	return _ZetaConnectorZEVM.Contract.PauserAddress(&_ZetaConnectorZEVM.CallOpts)
}

// PauserAddress is a free data retrieval call binding the contract method 0xf7fb869b.
//
// Solidity: function pauserAddress() view returns(address)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMCallerSession) PauserAddress() (common.Address, error) {
	return _ZetaConnectorZEVM.Contract.PauserAddress(&_ZetaConnectorZEVM.CallOpts)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMCaller) TssAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorZEVM.contract.Call(opts, &out, "tssAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMSession) TssAddress() (common.Address, error) {
	return _ZetaConnectorZEVM.Contract.TssAddress(&_ZetaConnectorZEVM.CallOpts)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMCallerSession) TssAddress() (common.Address, error) {
	return _ZetaConnectorZEVM.Contract.TssAddress(&_ZetaConnectorZEVM.CallOpts)
}

// TssAddressUpdater is a free data retrieval call binding the contract method 0x328a01d0.
//
// Solidity: function tssAddressUpdater() view returns(address)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMCaller) TssAddressUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorZEVM.contract.Call(opts, &out, "tssAddressUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TssAddressUpdater is a free data retrieval call binding the contract method 0x328a01d0.
//
// Solidity: function tssAddressUpdater() view returns(address)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMSession) TssAddressUpdater() (common.Address, error) {
	return _ZetaConnectorZEVM.Contract.TssAddressUpdater(&_ZetaConnectorZEVM.CallOpts)
}

// TssAddressUpdater is a free data retrieval call binding the contract method 0x328a01d0.
//
// Solidity: function tssAddressUpdater() view returns(address)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMCallerSession) TssAddressUpdater() (common.Address, error) {
	return _ZetaConnectorZEVM.Contract.TssAddressUpdater(&_ZetaConnectorZEVM.CallOpts)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMCaller) ZetaToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorZEVM.contract.Call(opts, &out, "zetaToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMSession) ZetaToken() (common.Address, error) {
	return _ZetaConnectorZEVM.Contract.ZetaToken(&_ZetaConnectorZEVM.CallOpts)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMCallerSession) ZetaToken() (common.Address, error) {
	return _ZetaConnectorZEVM.Contract.ZetaToken(&_ZetaConnectorZEVM.CallOpts)
}

// OnReceive is a paid mutator transaction binding the contract method 0x29dd214d.
//
// Solidity: function onReceive(bytes zetaTxSenderAddress, uint256 sourceChainId, address destinationAddress, uint256 zetaValue, bytes message, bytes32 internalSendHash) returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMTransactor) OnReceive(opts *bind.TransactOpts, zetaTxSenderAddress []byte, sourceChainId *big.Int, destinationAddress common.Address, zetaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.contract.Transact(opts, "onReceive", zetaTxSenderAddress, sourceChainId, destinationAddress, zetaValue, message, internalSendHash)
}

// OnReceive is a paid mutator transaction binding the contract method 0x29dd214d.
//
// Solidity: function onReceive(bytes zetaTxSenderAddress, uint256 sourceChainId, address destinationAddress, uint256 zetaValue, bytes message, bytes32 internalSendHash) returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMSession) OnReceive(zetaTxSenderAddress []byte, sourceChainId *big.Int, destinationAddress common.Address, zetaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.Contract.OnReceive(&_ZetaConnectorZEVM.TransactOpts, zetaTxSenderAddress, sourceChainId, destinationAddress, zetaValue, message, internalSendHash)
}

// OnReceive is a paid mutator transaction binding the contract method 0x29dd214d.
//
// Solidity: function onReceive(bytes zetaTxSenderAddress, uint256 sourceChainId, address destinationAddress, uint256 zetaValue, bytes message, bytes32 internalSendHash) returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMTransactorSession) OnReceive(zetaTxSenderAddress []byte, sourceChainId *big.Int, destinationAddress common.Address, zetaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.Contract.OnReceive(&_ZetaConnectorZEVM.TransactOpts, zetaTxSenderAddress, sourceChainId, destinationAddress, zetaValue, message, internalSendHash)
}

// OnRevert is a paid mutator transaction binding the contract method 0x942a5e16.
//
// Solidity: function onRevert(address zetaTxSenderAddress, uint256 sourceChainId, bytes destinationAddress, uint256 destinationChainId, uint256 remainingZetaValue, bytes message, bytes32 internalSendHash) returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMTransactor) OnRevert(opts *bind.TransactOpts, zetaTxSenderAddress common.Address, sourceChainId *big.Int, destinationAddress []byte, destinationChainId *big.Int, remainingZetaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.contract.Transact(opts, "onRevert", zetaTxSenderAddress, sourceChainId, destinationAddress, destinationChainId, remainingZetaValue, message, internalSendHash)
}

// OnRevert is a paid mutator transaction binding the contract method 0x942a5e16.
//
// Solidity: function onRevert(address zetaTxSenderAddress, uint256 sourceChainId, bytes destinationAddress, uint256 destinationChainId, uint256 remainingZetaValue, bytes message, bytes32 internalSendHash) returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMSession) OnRevert(zetaTxSenderAddress common.Address, sourceChainId *big.Int, destinationAddress []byte, destinationChainId *big.Int, remainingZetaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.Contract.OnRevert(&_ZetaConnectorZEVM.TransactOpts, zetaTxSenderAddress, sourceChainId, destinationAddress, destinationChainId, remainingZetaValue, message, internalSendHash)
}

// OnRevert is a paid mutator transaction binding the contract method 0x942a5e16.
//
// Solidity: function onRevert(address zetaTxSenderAddress, uint256 sourceChainId, bytes destinationAddress, uint256 destinationChainId, uint256 remainingZetaValue, bytes message, bytes32 internalSendHash) returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMTransactorSession) OnRevert(zetaTxSenderAddress common.Address, sourceChainId *big.Int, destinationAddress []byte, destinationChainId *big.Int, remainingZetaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.Contract.OnRevert(&_ZetaConnectorZEVM.TransactOpts, zetaTxSenderAddress, sourceChainId, destinationAddress, destinationChainId, remainingZetaValue, message, internalSendHash)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMSession) Pause() (*types.Transaction, error) {
	return _ZetaConnectorZEVM.Contract.Pause(&_ZetaConnectorZEVM.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMTransactorSession) Pause() (*types.Transaction, error) {
	return _ZetaConnectorZEVM.Contract.Pause(&_ZetaConnectorZEVM.TransactOpts)
}

// RenounceTssAddressUpdater is a paid mutator transaction binding the contract method 0x779e3b63.
//
// Solidity: function renounceTssAddressUpdater() returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMTransactor) RenounceTssAddressUpdater(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.contract.Transact(opts, "renounceTssAddressUpdater")
}

// RenounceTssAddressUpdater is a paid mutator transaction binding the contract method 0x779e3b63.
//
// Solidity: function renounceTssAddressUpdater() returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMSession) RenounceTssAddressUpdater() (*types.Transaction, error) {
	return _ZetaConnectorZEVM.Contract.RenounceTssAddressUpdater(&_ZetaConnectorZEVM.TransactOpts)
}

// RenounceTssAddressUpdater is a paid mutator transaction binding the contract method 0x779e3b63.
//
// Solidity: function renounceTssAddressUpdater() returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMTransactorSession) RenounceTssAddressUpdater() (*types.Transaction, error) {
	return _ZetaConnectorZEVM.Contract.RenounceTssAddressUpdater(&_ZetaConnectorZEVM.TransactOpts)
}

// Send is a paid mutator transaction binding the contract method 0xec026901.
//
// Solidity: function send((uint256,bytes,uint256,bytes,uint256,bytes) input) returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMTransactor) Send(opts *bind.TransactOpts, input ZetaInterfacesSendInput) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.contract.Transact(opts, "send", input)
}

// Send is a paid mutator transaction binding the contract method 0xec026901.
//
// Solidity: function send((uint256,bytes,uint256,bytes,uint256,bytes) input) returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMSession) Send(input ZetaInterfacesSendInput) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.Contract.Send(&_ZetaConnectorZEVM.TransactOpts, input)
}

// Send is a paid mutator transaction binding the contract method 0xec026901.
//
// Solidity: function send((uint256,bytes,uint256,bytes,uint256,bytes) input) returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMTransactorSession) Send(input ZetaInterfacesSendInput) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.Contract.Send(&_ZetaConnectorZEVM.TransactOpts, input)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMSession) Unpause() (*types.Transaction, error) {
	return _ZetaConnectorZEVM.Contract.Unpause(&_ZetaConnectorZEVM.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMTransactorSession) Unpause() (*types.Transaction, error) {
	return _ZetaConnectorZEVM.Contract.Unpause(&_ZetaConnectorZEVM.TransactOpts)
}

// UpdatePauserAddress is a paid mutator transaction binding the contract method 0x6128480f.
//
// Solidity: function updatePauserAddress(address pauserAddress_) returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMTransactor) UpdatePauserAddress(opts *bind.TransactOpts, pauserAddress_ common.Address) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.contract.Transact(opts, "updatePauserAddress", pauserAddress_)
}

// UpdatePauserAddress is a paid mutator transaction binding the contract method 0x6128480f.
//
// Solidity: function updatePauserAddress(address pauserAddress_) returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMSession) UpdatePauserAddress(pauserAddress_ common.Address) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.Contract.UpdatePauserAddress(&_ZetaConnectorZEVM.TransactOpts, pauserAddress_)
}

// UpdatePauserAddress is a paid mutator transaction binding the contract method 0x6128480f.
//
// Solidity: function updatePauserAddress(address pauserAddress_) returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMTransactorSession) UpdatePauserAddress(pauserAddress_ common.Address) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.Contract.UpdatePauserAddress(&_ZetaConnectorZEVM.TransactOpts, pauserAddress_)
}

// UpdateTssAddress is a paid mutator transaction binding the contract method 0x9122c344.
//
// Solidity: function updateTssAddress(address tssAddress_) returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMTransactor) UpdateTssAddress(opts *bind.TransactOpts, tssAddress_ common.Address) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.contract.Transact(opts, "updateTssAddress", tssAddress_)
}

// UpdateTssAddress is a paid mutator transaction binding the contract method 0x9122c344.
//
// Solidity: function updateTssAddress(address tssAddress_) returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMSession) UpdateTssAddress(tssAddress_ common.Address) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.Contract.UpdateTssAddress(&_ZetaConnectorZEVM.TransactOpts, tssAddress_)
}

// UpdateTssAddress is a paid mutator transaction binding the contract method 0x9122c344.
//
// Solidity: function updateTssAddress(address tssAddress_) returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMTransactorSession) UpdateTssAddress(tssAddress_ common.Address) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.Contract.UpdateTssAddress(&_ZetaConnectorZEVM.TransactOpts, tssAddress_)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMSession) Receive() (*types.Transaction, error) {
	return _ZetaConnectorZEVM.Contract.Receive(&_ZetaConnectorZEVM.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMTransactorSession) Receive() (*types.Transaction, error) {
	return _ZetaConnectorZEVM.Contract.Receive(&_ZetaConnectorZEVM.TransactOpts)
}

// ZetaConnectorZEVMPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ZetaConnectorZEVM contract.
type ZetaConnectorZEVMPausedIterator struct {
	Event *ZetaConnectorZEVMPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorZEVMPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorZEVMPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorZEVMPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorZEVMPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorZEVMPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorZEVMPaused represents a Paused event raised by the ZetaConnectorZEVM contract.
type ZetaConnectorZEVMPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMFilterer) FilterPaused(opts *bind.FilterOpts) (*ZetaConnectorZEVMPausedIterator, error) {

	logs, sub, err := _ZetaConnectorZEVM.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorZEVMPausedIterator{contract: _ZetaConnectorZEVM.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ZetaConnectorZEVMPaused) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorZEVM.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorZEVMPaused)
				if err := _ZetaConnectorZEVM.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMFilterer) ParsePaused(log types.Log) (*ZetaConnectorZEVMPaused, error) {
	event := new(ZetaConnectorZEVMPaused)
	if err := _ZetaConnectorZEVM.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorZEVMPauserAddressUpdatedIterator is returned from FilterPauserAddressUpdated and is used to iterate over the raw logs and unpacked data for PauserAddressUpdated events raised by the ZetaConnectorZEVM contract.
type ZetaConnectorZEVMPauserAddressUpdatedIterator struct {
	Event *ZetaConnectorZEVMPauserAddressUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorZEVMPauserAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorZEVMPauserAddressUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorZEVMPauserAddressUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorZEVMPauserAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorZEVMPauserAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorZEVMPauserAddressUpdated represents a PauserAddressUpdated event raised by the ZetaConnectorZEVM contract.
type ZetaConnectorZEVMPauserAddressUpdated struct {
	CallerAddress common.Address
	NewTssAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPauserAddressUpdated is a free log retrieval operation binding the contract event 0xd41d83655d484bdf299598751c371b2d92088667266fe3774b25a97bdd5d0397.
//
// Solidity: event PauserAddressUpdated(address callerAddress, address newTssAddress)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMFilterer) FilterPauserAddressUpdated(opts *bind.FilterOpts) (*ZetaConnectorZEVMPauserAddressUpdatedIterator, error) {

	logs, sub, err := _ZetaConnectorZEVM.contract.FilterLogs(opts, "PauserAddressUpdated")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorZEVMPauserAddressUpdatedIterator{contract: _ZetaConnectorZEVM.contract, event: "PauserAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchPauserAddressUpdated is a free log subscription operation binding the contract event 0xd41d83655d484bdf299598751c371b2d92088667266fe3774b25a97bdd5d0397.
//
// Solidity: event PauserAddressUpdated(address callerAddress, address newTssAddress)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMFilterer) WatchPauserAddressUpdated(opts *bind.WatchOpts, sink chan<- *ZetaConnectorZEVMPauserAddressUpdated) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorZEVM.contract.WatchLogs(opts, "PauserAddressUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorZEVMPauserAddressUpdated)
				if err := _ZetaConnectorZEVM.contract.UnpackLog(event, "PauserAddressUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePauserAddressUpdated is a log parse operation binding the contract event 0xd41d83655d484bdf299598751c371b2d92088667266fe3774b25a97bdd5d0397.
//
// Solidity: event PauserAddressUpdated(address callerAddress, address newTssAddress)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMFilterer) ParsePauserAddressUpdated(log types.Log) (*ZetaConnectorZEVMPauserAddressUpdated, error) {
	event := new(ZetaConnectorZEVMPauserAddressUpdated)
	if err := _ZetaConnectorZEVM.contract.UnpackLog(event, "PauserAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorZEVMTSSAddressUpdatedIterator is returned from FilterTSSAddressUpdated and is used to iterate over the raw logs and unpacked data for TSSAddressUpdated events raised by the ZetaConnectorZEVM contract.
type ZetaConnectorZEVMTSSAddressUpdatedIterator struct {
	Event *ZetaConnectorZEVMTSSAddressUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorZEVMTSSAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorZEVMTSSAddressUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorZEVMTSSAddressUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorZEVMTSSAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorZEVMTSSAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorZEVMTSSAddressUpdated represents a TSSAddressUpdated event raised by the ZetaConnectorZEVM contract.
type ZetaConnectorZEVMTSSAddressUpdated struct {
	CallerAddress common.Address
	NewTssAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTSSAddressUpdated is a free log retrieval operation binding the contract event 0xe79965b5c67dcfb2cf5fe152715e4a7256cee62a3d5dd8484fd8a8539eb8beff.
//
// Solidity: event TSSAddressUpdated(address callerAddress, address newTssAddress)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMFilterer) FilterTSSAddressUpdated(opts *bind.FilterOpts) (*ZetaConnectorZEVMTSSAddressUpdatedIterator, error) {

	logs, sub, err := _ZetaConnectorZEVM.contract.FilterLogs(opts, "TSSAddressUpdated")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorZEVMTSSAddressUpdatedIterator{contract: _ZetaConnectorZEVM.contract, event: "TSSAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchTSSAddressUpdated is a free log subscription operation binding the contract event 0xe79965b5c67dcfb2cf5fe152715e4a7256cee62a3d5dd8484fd8a8539eb8beff.
//
// Solidity: event TSSAddressUpdated(address callerAddress, address newTssAddress)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMFilterer) WatchTSSAddressUpdated(opts *bind.WatchOpts, sink chan<- *ZetaConnectorZEVMTSSAddressUpdated) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorZEVM.contract.WatchLogs(opts, "TSSAddressUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorZEVMTSSAddressUpdated)
				if err := _ZetaConnectorZEVM.contract.UnpackLog(event, "TSSAddressUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTSSAddressUpdated is a log parse operation binding the contract event 0xe79965b5c67dcfb2cf5fe152715e4a7256cee62a3d5dd8484fd8a8539eb8beff.
//
// Solidity: event TSSAddressUpdated(address callerAddress, address newTssAddress)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMFilterer) ParseTSSAddressUpdated(log types.Log) (*ZetaConnectorZEVMTSSAddressUpdated, error) {
	event := new(ZetaConnectorZEVMTSSAddressUpdated)
	if err := _ZetaConnectorZEVM.contract.UnpackLog(event, "TSSAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorZEVMTSSAddressUpdaterUpdatedIterator is returned from FilterTSSAddressUpdaterUpdated and is used to iterate over the raw logs and unpacked data for TSSAddressUpdaterUpdated events raised by the ZetaConnectorZEVM contract.
type ZetaConnectorZEVMTSSAddressUpdaterUpdatedIterator struct {
	Event *ZetaConnectorZEVMTSSAddressUpdaterUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorZEVMTSSAddressUpdaterUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorZEVMTSSAddressUpdaterUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorZEVMTSSAddressUpdaterUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorZEVMTSSAddressUpdaterUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorZEVMTSSAddressUpdaterUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorZEVMTSSAddressUpdaterUpdated represents a TSSAddressUpdaterUpdated event raised by the ZetaConnectorZEVM contract.
type ZetaConnectorZEVMTSSAddressUpdaterUpdated struct {
	CallerAddress        common.Address
	NewTssUpdaterAddress common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterTSSAddressUpdaterUpdated is a free log retrieval operation binding the contract event 0x5104c9abdc7d111c2aeb4ce890ac70274a4be2ee83f46a62551be5e6ebc82dd0.
//
// Solidity: event TSSAddressUpdaterUpdated(address callerAddress, address newTssUpdaterAddress)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMFilterer) FilterTSSAddressUpdaterUpdated(opts *bind.FilterOpts) (*ZetaConnectorZEVMTSSAddressUpdaterUpdatedIterator, error) {

	logs, sub, err := _ZetaConnectorZEVM.contract.FilterLogs(opts, "TSSAddressUpdaterUpdated")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorZEVMTSSAddressUpdaterUpdatedIterator{contract: _ZetaConnectorZEVM.contract, event: "TSSAddressUpdaterUpdated", logs: logs, sub: sub}, nil
}

// WatchTSSAddressUpdaterUpdated is a free log subscription operation binding the contract event 0x5104c9abdc7d111c2aeb4ce890ac70274a4be2ee83f46a62551be5e6ebc82dd0.
//
// Solidity: event TSSAddressUpdaterUpdated(address callerAddress, address newTssUpdaterAddress)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMFilterer) WatchTSSAddressUpdaterUpdated(opts *bind.WatchOpts, sink chan<- *ZetaConnectorZEVMTSSAddressUpdaterUpdated) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorZEVM.contract.WatchLogs(opts, "TSSAddressUpdaterUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorZEVMTSSAddressUpdaterUpdated)
				if err := _ZetaConnectorZEVM.contract.UnpackLog(event, "TSSAddressUpdaterUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTSSAddressUpdaterUpdated is a log parse operation binding the contract event 0x5104c9abdc7d111c2aeb4ce890ac70274a4be2ee83f46a62551be5e6ebc82dd0.
//
// Solidity: event TSSAddressUpdaterUpdated(address callerAddress, address newTssUpdaterAddress)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMFilterer) ParseTSSAddressUpdaterUpdated(log types.Log) (*ZetaConnectorZEVMTSSAddressUpdaterUpdated, error) {
	event := new(ZetaConnectorZEVMTSSAddressUpdaterUpdated)
	if err := _ZetaConnectorZEVM.contract.UnpackLog(event, "TSSAddressUpdaterUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorZEVMUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ZetaConnectorZEVM contract.
type ZetaConnectorZEVMUnpausedIterator struct {
	Event *ZetaConnectorZEVMUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorZEVMUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorZEVMUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorZEVMUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorZEVMUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorZEVMUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorZEVMUnpaused represents a Unpaused event raised by the ZetaConnectorZEVM contract.
type ZetaConnectorZEVMUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ZetaConnectorZEVMUnpausedIterator, error) {

	logs, sub, err := _ZetaConnectorZEVM.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorZEVMUnpausedIterator{contract: _ZetaConnectorZEVM.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ZetaConnectorZEVMUnpaused) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorZEVM.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorZEVMUnpaused)
				if err := _ZetaConnectorZEVM.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMFilterer) ParseUnpaused(log types.Log) (*ZetaConnectorZEVMUnpaused, error) {
	event := new(ZetaConnectorZEVMUnpaused)
	if err := _ZetaConnectorZEVM.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorZEVMZetaReceivedIterator is returned from FilterZetaReceived and is used to iterate over the raw logs and unpacked data for ZetaReceived events raised by the ZetaConnectorZEVM contract.
type ZetaConnectorZEVMZetaReceivedIterator struct {
	Event *ZetaConnectorZEVMZetaReceived // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorZEVMZetaReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorZEVMZetaReceived)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorZEVMZetaReceived)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorZEVMZetaReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorZEVMZetaReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorZEVMZetaReceived represents a ZetaReceived event raised by the ZetaConnectorZEVM contract.
type ZetaConnectorZEVMZetaReceived struct {
	ZetaTxSenderAddress []byte
	SourceChainId       *big.Int
	DestinationAddress  common.Address
	ZetaValue           *big.Int
	Message             []byte
	InternalSendHash    [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterZetaReceived is a free log retrieval operation binding the contract event 0xf1302855733b40d8acb467ee990b6d56c05c80e28ebcabfa6e6f3f57cb50d698.
//
// Solidity: event ZetaReceived(bytes zetaTxSenderAddress, uint256 indexed sourceChainId, address indexed destinationAddress, uint256 zetaValue, bytes message, bytes32 indexed internalSendHash)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMFilterer) FilterZetaReceived(opts *bind.FilterOpts, sourceChainId []*big.Int, destinationAddress []common.Address, internalSendHash [][32]byte) (*ZetaConnectorZEVMZetaReceivedIterator, error) {

	var sourceChainIdRule []interface{}
	for _, sourceChainIdItem := range sourceChainId {
		sourceChainIdRule = append(sourceChainIdRule, sourceChainIdItem)
	}
	var destinationAddressRule []interface{}
	for _, destinationAddressItem := range destinationAddress {
		destinationAddressRule = append(destinationAddressRule, destinationAddressItem)
	}

	var internalSendHashRule []interface{}
	for _, internalSendHashItem := range internalSendHash {
		internalSendHashRule = append(internalSendHashRule, internalSendHashItem)
	}

	logs, sub, err := _ZetaConnectorZEVM.contract.FilterLogs(opts, "ZetaReceived", sourceChainIdRule, destinationAddressRule, internalSendHashRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorZEVMZetaReceivedIterator{contract: _ZetaConnectorZEVM.contract, event: "ZetaReceived", logs: logs, sub: sub}, nil
}

// WatchZetaReceived is a free log subscription operation binding the contract event 0xf1302855733b40d8acb467ee990b6d56c05c80e28ebcabfa6e6f3f57cb50d698.
//
// Solidity: event ZetaReceived(bytes zetaTxSenderAddress, uint256 indexed sourceChainId, address indexed destinationAddress, uint256 zetaValue, bytes message, bytes32 indexed internalSendHash)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMFilterer) WatchZetaReceived(opts *bind.WatchOpts, sink chan<- *ZetaConnectorZEVMZetaReceived, sourceChainId []*big.Int, destinationAddress []common.Address, internalSendHash [][32]byte) (event.Subscription, error) {

	var sourceChainIdRule []interface{}
	for _, sourceChainIdItem := range sourceChainId {
		sourceChainIdRule = append(sourceChainIdRule, sourceChainIdItem)
	}
	var destinationAddressRule []interface{}
	for _, destinationAddressItem := range destinationAddress {
		destinationAddressRule = append(destinationAddressRule, destinationAddressItem)
	}

	var internalSendHashRule []interface{}
	for _, internalSendHashItem := range internalSendHash {
		internalSendHashRule = append(internalSendHashRule, internalSendHashItem)
	}

	logs, sub, err := _ZetaConnectorZEVM.contract.WatchLogs(opts, "ZetaReceived", sourceChainIdRule, destinationAddressRule, internalSendHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorZEVMZetaReceived)
				if err := _ZetaConnectorZEVM.contract.UnpackLog(event, "ZetaReceived", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseZetaReceived is a log parse operation binding the contract event 0xf1302855733b40d8acb467ee990b6d56c05c80e28ebcabfa6e6f3f57cb50d698.
//
// Solidity: event ZetaReceived(bytes zetaTxSenderAddress, uint256 indexed sourceChainId, address indexed destinationAddress, uint256 zetaValue, bytes message, bytes32 indexed internalSendHash)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMFilterer) ParseZetaReceived(log types.Log) (*ZetaConnectorZEVMZetaReceived, error) {
	event := new(ZetaConnectorZEVMZetaReceived)
	if err := _ZetaConnectorZEVM.contract.UnpackLog(event, "ZetaReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorZEVMZetaRevertedIterator is returned from FilterZetaReverted and is used to iterate over the raw logs and unpacked data for ZetaReverted events raised by the ZetaConnectorZEVM contract.
type ZetaConnectorZEVMZetaRevertedIterator struct {
	Event *ZetaConnectorZEVMZetaReverted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorZEVMZetaRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorZEVMZetaReverted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorZEVMZetaReverted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorZEVMZetaRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorZEVMZetaRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorZEVMZetaReverted represents a ZetaReverted event raised by the ZetaConnectorZEVM contract.
type ZetaConnectorZEVMZetaReverted struct {
	ZetaTxSenderAddress common.Address
	SourceChainId       *big.Int
	DestinationChainId  *big.Int
	DestinationAddress  []byte
	RemainingZetaValue  *big.Int
	Message             []byte
	InternalSendHash    [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterZetaReverted is a free log retrieval operation binding the contract event 0x521fb0b407c2eb9b1375530e9b9a569889992140a688bc076aa72c1712012c88.
//
// Solidity: event ZetaReverted(address zetaTxSenderAddress, uint256 sourceChainId, uint256 indexed destinationChainId, bytes destinationAddress, uint256 remainingZetaValue, bytes message, bytes32 indexed internalSendHash)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMFilterer) FilterZetaReverted(opts *bind.FilterOpts, destinationChainId []*big.Int, internalSendHash [][32]byte) (*ZetaConnectorZEVMZetaRevertedIterator, error) {

	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	var internalSendHashRule []interface{}
	for _, internalSendHashItem := range internalSendHash {
		internalSendHashRule = append(internalSendHashRule, internalSendHashItem)
	}

	logs, sub, err := _ZetaConnectorZEVM.contract.FilterLogs(opts, "ZetaReverted", destinationChainIdRule, internalSendHashRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorZEVMZetaRevertedIterator{contract: _ZetaConnectorZEVM.contract, event: "ZetaReverted", logs: logs, sub: sub}, nil
}

// WatchZetaReverted is a free log subscription operation binding the contract event 0x521fb0b407c2eb9b1375530e9b9a569889992140a688bc076aa72c1712012c88.
//
// Solidity: event ZetaReverted(address zetaTxSenderAddress, uint256 sourceChainId, uint256 indexed destinationChainId, bytes destinationAddress, uint256 remainingZetaValue, bytes message, bytes32 indexed internalSendHash)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMFilterer) WatchZetaReverted(opts *bind.WatchOpts, sink chan<- *ZetaConnectorZEVMZetaReverted, destinationChainId []*big.Int, internalSendHash [][32]byte) (event.Subscription, error) {

	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	var internalSendHashRule []interface{}
	for _, internalSendHashItem := range internalSendHash {
		internalSendHashRule = append(internalSendHashRule, internalSendHashItem)
	}

	logs, sub, err := _ZetaConnectorZEVM.contract.WatchLogs(opts, "ZetaReverted", destinationChainIdRule, internalSendHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorZEVMZetaReverted)
				if err := _ZetaConnectorZEVM.contract.UnpackLog(event, "ZetaReverted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseZetaReverted is a log parse operation binding the contract event 0x521fb0b407c2eb9b1375530e9b9a569889992140a688bc076aa72c1712012c88.
//
// Solidity: event ZetaReverted(address zetaTxSenderAddress, uint256 sourceChainId, uint256 indexed destinationChainId, bytes destinationAddress, uint256 remainingZetaValue, bytes message, bytes32 indexed internalSendHash)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMFilterer) ParseZetaReverted(log types.Log) (*ZetaConnectorZEVMZetaReverted, error) {
	event := new(ZetaConnectorZEVMZetaReverted)
	if err := _ZetaConnectorZEVM.contract.UnpackLog(event, "ZetaReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorZEVMZetaSentIterator is returned from FilterZetaSent and is used to iterate over the raw logs and unpacked data for ZetaSent events raised by the ZetaConnectorZEVM contract.
type ZetaConnectorZEVMZetaSentIterator struct {
	Event *ZetaConnectorZEVMZetaSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorZEVMZetaSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorZEVMZetaSent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorZEVMZetaSent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorZEVMZetaSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorZEVMZetaSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorZEVMZetaSent represents a ZetaSent event raised by the ZetaConnectorZEVM contract.
type ZetaConnectorZEVMZetaSent struct {
	SourceTxOriginAddress common.Address
	ZetaTxSenderAddress   common.Address
	DestinationChainId    *big.Int
	DestinationAddress    []byte
	ZetaValueAndGas       *big.Int
	DestinationGasLimit   *big.Int
	Message               []byte
	ZetaParams            []byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterZetaSent is a free log retrieval operation binding the contract event 0x7ec1c94701e09b1652f3e1d307e60c4b9ebf99aff8c2079fd1d8c585e031c4e4.
//
// Solidity: event ZetaSent(address sourceTxOriginAddress, address indexed zetaTxSenderAddress, uint256 indexed destinationChainId, bytes destinationAddress, uint256 zetaValueAndGas, uint256 destinationGasLimit, bytes message, bytes zetaParams)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMFilterer) FilterZetaSent(opts *bind.FilterOpts, zetaTxSenderAddress []common.Address, destinationChainId []*big.Int) (*ZetaConnectorZEVMZetaSentIterator, error) {

	var zetaTxSenderAddressRule []interface{}
	for _, zetaTxSenderAddressItem := range zetaTxSenderAddress {
		zetaTxSenderAddressRule = append(zetaTxSenderAddressRule, zetaTxSenderAddressItem)
	}
	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	logs, sub, err := _ZetaConnectorZEVM.contract.FilterLogs(opts, "ZetaSent", zetaTxSenderAddressRule, destinationChainIdRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorZEVMZetaSentIterator{contract: _ZetaConnectorZEVM.contract, event: "ZetaSent", logs: logs, sub: sub}, nil
}

// WatchZetaSent is a free log subscription operation binding the contract event 0x7ec1c94701e09b1652f3e1d307e60c4b9ebf99aff8c2079fd1d8c585e031c4e4.
//
// Solidity: event ZetaSent(address sourceTxOriginAddress, address indexed zetaTxSenderAddress, uint256 indexed destinationChainId, bytes destinationAddress, uint256 zetaValueAndGas, uint256 destinationGasLimit, bytes message, bytes zetaParams)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMFilterer) WatchZetaSent(opts *bind.WatchOpts, sink chan<- *ZetaConnectorZEVMZetaSent, zetaTxSenderAddress []common.Address, destinationChainId []*big.Int) (event.Subscription, error) {

	var zetaTxSenderAddressRule []interface{}
	for _, zetaTxSenderAddressItem := range zetaTxSenderAddress {
		zetaTxSenderAddressRule = append(zetaTxSenderAddressRule, zetaTxSenderAddressItem)
	}
	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	logs, sub, err := _ZetaConnectorZEVM.contract.WatchLogs(opts, "ZetaSent", zetaTxSenderAddressRule, destinationChainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorZEVMZetaSent)
				if err := _ZetaConnectorZEVM.contract.UnpackLog(event, "ZetaSent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseZetaSent is a log parse operation binding the contract event 0x7ec1c94701e09b1652f3e1d307e60c4b9ebf99aff8c2079fd1d8c585e031c4e4.
//
// Solidity: event ZetaSent(address sourceTxOriginAddress, address indexed zetaTxSenderAddress, uint256 indexed destinationChainId, bytes destinationAddress, uint256 zetaValueAndGas, uint256 destinationGasLimit, bytes message, bytes zetaParams)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMFilterer) ParseZetaSent(log types.Log) (*ZetaConnectorZEVMZetaSent, error) {
	event := new(ZetaConnectorZEVMZetaSent)
	if err := _ZetaConnectorZEVM.contract.UnpackLog(event, "ZetaSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
