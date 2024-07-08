// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mockerc721

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

// MockERC721MetaData contains all meta data concerning the MockERC721 contract.
var MockERC721MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611e69806100206000396000f3fe6080604052600436106100dd5760003560e01c80636352211e1161007f578063a22cb46511610059578063a22cb465146102a9578063b88d4fde146102d2578063c87b56dd146102ee578063e985e9c51461032b576100dd565b80636352211e1461020457806370a082311461024157806395d89b411461027e576100dd565b8063095ea7b3116100bb578063095ea7b31461018757806323b872dd146101a357806342842e0e146101bf5780634cd88b76146101db576100dd565b806301ffc9a7146100e257806306fdde031461011f578063081812fc1461014a575b600080fd5b3480156100ee57600080fd5b5061010960048036038101906101049190611519565b610368565b6040516101169190611880565b60405180910390f35b34801561012b57600080fd5b506101346103fa565b604051610141919061189b565b60405180910390f35b34801561015657600080fd5b50610171600480360381019061016c91906115eb565b61048c565b60405161017e91906117cf565b60405180910390f35b6101a1600480360381019061019c91906114d9565b6104c9565b005b6101bd60048036038101906101b891906113c3565b6106b2565b005b6101d960048036038101906101d491906113c3565b610abd565b005b3480156101e757600080fd5b5061020260048036038101906101fd9190611573565b610bf3565b005b34801561021057600080fd5b5061022b600480360381019061022691906115eb565b610c90565b60405161023891906117cf565b60405180910390f35b34801561024d57600080fd5b5061026860048036038101906102639190611356565b610d3c565b604051610275919061199d565b60405180910390f35b34801561028a57600080fd5b50610293610df4565b6040516102a0919061189b565b60405180910390f35b3480156102b557600080fd5b506102d060048036038101906102cb9190611499565b610e86565b005b6102ec60048036038101906102e79190611416565b610f83565b005b3480156102fa57600080fd5b50610315600480360381019061031091906115eb565b6110bc565b604051610322919061189b565b60405180910390f35b34801561033757600080fd5b50610352600480360381019061034d9190611383565b6110c3565b60405161035f9190611880565b60405180910390f35b60006301ffc9a760e01b827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614806103c357506380ac58cd60e01b827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b806103f35750635b5e139f60e01b827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b9050919050565b60606000805461040990611b57565b80601f016020809104026020016040519081016040528092919081815260200182805461043590611b57565b80156104825780601f1061045757610100808354040283529160200191610482565b820191906000526020600020905b81548152906001019060200180831161046557829003601f168201915b5050505050905090565b60006004600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b60006002600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614806105c15750600560008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff165b610600576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105f79061193d565b60405180910390fd5b826004600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550818373ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a4505050565b6002600082815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614610753576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161074a9061197d565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156107c3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107ba906118dd565b60405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614806108835750600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff165b806108ec57506004600082815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b61092b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109229061193d565b60405180910390fd5b600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600081548092919061097b90611b2d565b9190505550600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008154809291906109d090611bba565b9190505550816002600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506004600082815260200190815260200160002060006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055808273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4505050565b610ac88383836106b2565b610ad182611157565b1580610baf575063150b7a0260e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168273ffffffffffffffffffffffffffffffffffffffff1663150b7a023386856040518463ffffffff1660e01b8152600401610b3c93929190611836565b602060405180830381600087803b158015610b5657600080fd5b505af1158015610b6a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b8e9190611546565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b610bee576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610be59061191d565b60405180910390fd5b505050565b600660009054906101000a900460ff1615610c43576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c3a906118bd565b60405180910390fd5b8160009080519060200190610c5992919061116a565b508060019080519060200190610c7092919061116a565b506001600660006101000a81548160ff0219169083151502179055505050565b60008073ffffffffffffffffffffffffffffffffffffffff166002600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1691508173ffffffffffffffffffffffffffffffffffffffff161415610d37576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d2e9061195d565b60405180910390fd5b919050565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610dad576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610da4906118fd565b60405180910390fd5b600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b606060018054610e0390611b57565b80601f0160208091040260200160405190810160405280929190818152602001828054610e2f90611b57565b8015610e7c5780601f10610e5157610100808354040283529160200191610e7c565b820191906000526020600020905b815481529060010190602001808311610e5f57829003601f168201915b5050505050905090565b80600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c3183604051610f779190611880565b60405180910390a35050565b610f8e8484846106b2565b610f9783611157565b1580611077575063150b7a0260e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168373ffffffffffffffffffffffffffffffffffffffff1663150b7a02338786866040518563ffffffff1660e01b815260040161100494939291906117ea565b602060405180830381600087803b15801561101e57600080fd5b505af1158015611032573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110569190611546565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b6110b6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110ad9061191d565b60405180910390fd5b50505050565b6060919050565b6000600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b600080823b905060008111915050919050565b82805461117690611b57565b90600052602060002090601f01602090048101928261119857600085556111df565b82601f106111b157805160ff19168380011785556111df565b828001600101855582156111df579182015b828111156111de5782518255916020019190600101906111c3565b5b5090506111ec91906111f0565b5090565b5b808211156112095760008160009055506001016111f1565b5090565b600061122061121b846119dd565b6119b8565b90508281526020810184848401111561123c5761123b611c95565b5b611247848285611aeb565b509392505050565b600061126261125d84611a0e565b6119b8565b90508281526020810184848401111561127e5761127d611c95565b5b611289848285611aeb565b509392505050565b6000813590506112a081611dd7565b92915050565b6000813590506112b581611dee565b92915050565b6000813590506112ca81611e05565b92915050565b6000815190506112df81611e05565b92915050565b600082601f8301126112fa576112f9611c90565b5b813561130a84826020860161120d565b91505092915050565b600082601f83011261132857611327611c90565b5b813561133884826020860161124f565b91505092915050565b60008135905061135081611e1c565b92915050565b60006020828403121561136c5761136b611c9f565b5b600061137a84828501611291565b91505092915050565b6000806040838503121561139a57611399611c9f565b5b60006113a885828601611291565b92505060206113b985828601611291565b9150509250929050565b6000806000606084860312156113dc576113db611c9f565b5b60006113ea86828701611291565b93505060206113fb86828701611291565b925050604061140c86828701611341565b9150509250925092565b600080600080608085870312156114305761142f611c9f565b5b600061143e87828801611291565b945050602061144f87828801611291565b935050604061146087828801611341565b925050606085013567ffffffffffffffff81111561148157611480611c9a565b5b61148d878288016112e5565b91505092959194509250565b600080604083850312156114b0576114af611c9f565b5b60006114be85828601611291565b92505060206114cf858286016112a6565b9150509250929050565b600080604083850312156114f0576114ef611c9f565b5b60006114fe85828601611291565b925050602061150f85828601611341565b9150509250929050565b60006020828403121561152f5761152e611c9f565b5b600061153d848285016112bb565b91505092915050565b60006020828403121561155c5761155b611c9f565b5b600061156a848285016112d0565b91505092915050565b6000806040838503121561158a57611589611c9f565b5b600083013567ffffffffffffffff8111156115a8576115a7611c9a565b5b6115b485828601611313565b925050602083013567ffffffffffffffff8111156115d5576115d4611c9a565b5b6115e185828601611313565b9150509250929050565b60006020828403121561160157611600611c9f565b5b600061160f84828501611341565b91505092915050565b61162181611a77565b82525050565b61163081611a89565b82525050565b600061164182611a3f565b61164b8185611a55565b935061165b818560208601611afa565b61166481611ca4565b840191505092915050565b600061167a82611a4a565b6116848185611a66565b9350611694818560208601611afa565b61169d81611ca4565b840191505092915050565b60006116b5601383611a66565b91506116c082611cb5565b602082019050919050565b60006116d8601183611a66565b91506116e382611cde565b602082019050919050565b60006116fb600c83611a66565b915061170682611d07565b602082019050919050565b600061171e601083611a66565b915061172982611d30565b602082019050919050565b6000611741600083611a55565b915061174c82611d59565b600082019050919050565b6000611764600e83611a66565b915061176f82611d5c565b602082019050919050565b6000611787600a83611a66565b915061179282611d85565b602082019050919050565b60006117aa600a83611a66565b91506117b582611dae565b602082019050919050565b6117c981611ae1565b82525050565b60006020820190506117e46000830184611618565b92915050565b60006080820190506117ff6000830187611618565b61180c6020830186611618565b61181960408301856117c0565b818103606083015261182b8184611636565b905095945050505050565b600060808201905061184b6000830186611618565b6118586020830185611618565b61186560408301846117c0565b818103606083015261187681611734565b9050949350505050565b60006020820190506118956000830184611627565b92915050565b600060208201905081810360008301526118b5818461166f565b905092915050565b600060208201905081810360008301526118d6816116a8565b9050919050565b600060208201905081810360008301526118f6816116cb565b9050919050565b60006020820190508181036000830152611916816116ee565b9050919050565b6000602082019050818103600083015261193681611711565b9050919050565b6000602082019050818103600083015261195681611757565b9050919050565b600060208201905081810360008301526119768161177a565b9050919050565b600060208201905081810360008301526119968161179d565b9050919050565b60006020820190506119b260008301846117c0565b92915050565b60006119c26119d3565b90506119ce8282611b89565b919050565b6000604051905090565b600067ffffffffffffffff8211156119f8576119f7611c61565b5b611a0182611ca4565b9050602081019050919050565b600067ffffffffffffffff821115611a2957611a28611c61565b5b611a3282611ca4565b9050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b6000611a8282611ac1565b9050919050565b60008115159050919050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b83811015611b18578082015181840152602081019050611afd565b83811115611b27576000848401525b50505050565b6000611b3882611ae1565b91506000821415611b4c57611b4b611c03565b5b600182039050919050565b60006002820490506001821680611b6f57607f821691505b60208210811415611b8357611b82611c32565b5b50919050565b611b9282611ca4565b810181811067ffffffffffffffff82111715611bb157611bb0611c61565b5b80604052505050565b6000611bc582611ae1565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415611bf857611bf7611c03565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f414c52454144595f494e495449414c495a454400000000000000000000000000600082015250565b7f494e56414c49445f524543495049454e54000000000000000000000000000000600082015250565b7f5a45524f5f414444524553530000000000000000000000000000000000000000600082015250565b7f554e534146455f524543495049454e5400000000000000000000000000000000600082015250565b50565b7f4e4f545f415554484f52495a4544000000000000000000000000000000000000600082015250565b7f4e4f545f4d494e54454400000000000000000000000000000000000000000000600082015250565b7f57524f4e475f46524f4d00000000000000000000000000000000000000000000600082015250565b611de081611a77565b8114611deb57600080fd5b50565b611df781611a89565b8114611e0257600080fd5b50565b611e0e81611a95565b8114611e1957600080fd5b50565b611e2581611ae1565b8114611e3057600080fd5b5056fea264697066735822122096b381bcc1d3bc0c5ed578328f1d7697016e32ad0e05efb995b62d3457df066664736f6c63430008070033",
}

// MockERC721ABI is the input ABI used to generate the binding from.
// Deprecated: Use MockERC721MetaData.ABI instead.
var MockERC721ABI = MockERC721MetaData.ABI

// MockERC721Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MockERC721MetaData.Bin instead.
var MockERC721Bin = MockERC721MetaData.Bin

// DeployMockERC721 deploys a new Ethereum contract, binding an instance of MockERC721 to it.
func DeployMockERC721(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MockERC721, error) {
	parsed, err := MockERC721MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockERC721Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockERC721{MockERC721Caller: MockERC721Caller{contract: contract}, MockERC721Transactor: MockERC721Transactor{contract: contract}, MockERC721Filterer: MockERC721Filterer{contract: contract}}, nil
}

// MockERC721 is an auto generated Go binding around an Ethereum contract.
type MockERC721 struct {
	MockERC721Caller     // Read-only binding to the contract
	MockERC721Transactor // Write-only binding to the contract
	MockERC721Filterer   // Log filterer for contract events
}

// MockERC721Caller is an auto generated read-only Go binding around an Ethereum contract.
type MockERC721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockERC721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type MockERC721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockERC721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockERC721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockERC721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockERC721Session struct {
	Contract     *MockERC721       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MockERC721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockERC721CallerSession struct {
	Contract *MockERC721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MockERC721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockERC721TransactorSession struct {
	Contract     *MockERC721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MockERC721Raw is an auto generated low-level Go binding around an Ethereum contract.
type MockERC721Raw struct {
	Contract *MockERC721 // Generic contract binding to access the raw methods on
}

// MockERC721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockERC721CallerRaw struct {
	Contract *MockERC721Caller // Generic read-only contract binding to access the raw methods on
}

// MockERC721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockERC721TransactorRaw struct {
	Contract *MockERC721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMockERC721 creates a new instance of MockERC721, bound to a specific deployed contract.
func NewMockERC721(address common.Address, backend bind.ContractBackend) (*MockERC721, error) {
	contract, err := bindMockERC721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockERC721{MockERC721Caller: MockERC721Caller{contract: contract}, MockERC721Transactor: MockERC721Transactor{contract: contract}, MockERC721Filterer: MockERC721Filterer{contract: contract}}, nil
}

// NewMockERC721Caller creates a new read-only instance of MockERC721, bound to a specific deployed contract.
func NewMockERC721Caller(address common.Address, caller bind.ContractCaller) (*MockERC721Caller, error) {
	contract, err := bindMockERC721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockERC721Caller{contract: contract}, nil
}

// NewMockERC721Transactor creates a new write-only instance of MockERC721, bound to a specific deployed contract.
func NewMockERC721Transactor(address common.Address, transactor bind.ContractTransactor) (*MockERC721Transactor, error) {
	contract, err := bindMockERC721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockERC721Transactor{contract: contract}, nil
}

// NewMockERC721Filterer creates a new log filterer instance of MockERC721, bound to a specific deployed contract.
func NewMockERC721Filterer(address common.Address, filterer bind.ContractFilterer) (*MockERC721Filterer, error) {
	contract, err := bindMockERC721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockERC721Filterer{contract: contract}, nil
}

// bindMockERC721 binds a generic wrapper to an already deployed contract.
func bindMockERC721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockERC721MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockERC721 *MockERC721Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockERC721.Contract.MockERC721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockERC721 *MockERC721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockERC721.Contract.MockERC721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockERC721 *MockERC721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockERC721.Contract.MockERC721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockERC721 *MockERC721CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockERC721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockERC721 *MockERC721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockERC721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockERC721 *MockERC721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockERC721.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_MockERC721 *MockERC721Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MockERC721.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_MockERC721 *MockERC721Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _MockERC721.Contract.BalanceOf(&_MockERC721.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_MockERC721 *MockERC721CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _MockERC721.Contract.BalanceOf(&_MockERC721.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 id) view returns(address)
func (_MockERC721 *MockERC721Caller) GetApproved(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MockERC721.contract.Call(opts, &out, "getApproved", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 id) view returns(address)
func (_MockERC721 *MockERC721Session) GetApproved(id *big.Int) (common.Address, error) {
	return _MockERC721.Contract.GetApproved(&_MockERC721.CallOpts, id)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 id) view returns(address)
func (_MockERC721 *MockERC721CallerSession) GetApproved(id *big.Int) (common.Address, error) {
	return _MockERC721.Contract.GetApproved(&_MockERC721.CallOpts, id)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_MockERC721 *MockERC721Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _MockERC721.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_MockERC721 *MockERC721Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _MockERC721.Contract.IsApprovedForAll(&_MockERC721.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_MockERC721 *MockERC721CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _MockERC721.Contract.IsApprovedForAll(&_MockERC721.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MockERC721 *MockERC721Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MockERC721.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MockERC721 *MockERC721Session) Name() (string, error) {
	return _MockERC721.Contract.Name(&_MockERC721.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MockERC721 *MockERC721CallerSession) Name() (string, error) {
	return _MockERC721.Contract.Name(&_MockERC721.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 id) view returns(address owner)
func (_MockERC721 *MockERC721Caller) OwnerOf(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MockERC721.contract.Call(opts, &out, "ownerOf", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 id) view returns(address owner)
func (_MockERC721 *MockERC721Session) OwnerOf(id *big.Int) (common.Address, error) {
	return _MockERC721.Contract.OwnerOf(&_MockERC721.CallOpts, id)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 id) view returns(address owner)
func (_MockERC721 *MockERC721CallerSession) OwnerOf(id *big.Int) (common.Address, error) {
	return _MockERC721.Contract.OwnerOf(&_MockERC721.CallOpts, id)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MockERC721 *MockERC721Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _MockERC721.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MockERC721 *MockERC721Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MockERC721.Contract.SupportsInterface(&_MockERC721.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MockERC721 *MockERC721CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MockERC721.Contract.SupportsInterface(&_MockERC721.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MockERC721 *MockERC721Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MockERC721.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MockERC721 *MockERC721Session) Symbol() (string, error) {
	return _MockERC721.Contract.Symbol(&_MockERC721.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MockERC721 *MockERC721CallerSession) Symbol() (string, error) {
	return _MockERC721.Contract.Symbol(&_MockERC721.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 id) view returns(string)
func (_MockERC721 *MockERC721Caller) TokenURI(opts *bind.CallOpts, id *big.Int) (string, error) {
	var out []interface{}
	err := _MockERC721.contract.Call(opts, &out, "tokenURI", id)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 id) view returns(string)
func (_MockERC721 *MockERC721Session) TokenURI(id *big.Int) (string, error) {
	return _MockERC721.Contract.TokenURI(&_MockERC721.CallOpts, id)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 id) view returns(string)
func (_MockERC721 *MockERC721CallerSession) TokenURI(id *big.Int) (string, error) {
	return _MockERC721.Contract.TokenURI(&_MockERC721.CallOpts, id)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 id) payable returns()
func (_MockERC721 *MockERC721Transactor) Approve(opts *bind.TransactOpts, spender common.Address, id *big.Int) (*types.Transaction, error) {
	return _MockERC721.contract.Transact(opts, "approve", spender, id)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 id) payable returns()
func (_MockERC721 *MockERC721Session) Approve(spender common.Address, id *big.Int) (*types.Transaction, error) {
	return _MockERC721.Contract.Approve(&_MockERC721.TransactOpts, spender, id)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 id) payable returns()
func (_MockERC721 *MockERC721TransactorSession) Approve(spender common.Address, id *big.Int) (*types.Transaction, error) {
	return _MockERC721.Contract.Approve(&_MockERC721.TransactOpts, spender, id)
}

// Initialize is a paid mutator transaction binding the contract method 0x4cd88b76.
//
// Solidity: function initialize(string name_, string symbol_) returns()
func (_MockERC721 *MockERC721Transactor) Initialize(opts *bind.TransactOpts, name_ string, symbol_ string) (*types.Transaction, error) {
	return _MockERC721.contract.Transact(opts, "initialize", name_, symbol_)
}

// Initialize is a paid mutator transaction binding the contract method 0x4cd88b76.
//
// Solidity: function initialize(string name_, string symbol_) returns()
func (_MockERC721 *MockERC721Session) Initialize(name_ string, symbol_ string) (*types.Transaction, error) {
	return _MockERC721.Contract.Initialize(&_MockERC721.TransactOpts, name_, symbol_)
}

// Initialize is a paid mutator transaction binding the contract method 0x4cd88b76.
//
// Solidity: function initialize(string name_, string symbol_) returns()
func (_MockERC721 *MockERC721TransactorSession) Initialize(name_ string, symbol_ string) (*types.Transaction, error) {
	return _MockERC721.Contract.Initialize(&_MockERC721.TransactOpts, name_, symbol_)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id) payable returns()
func (_MockERC721 *MockERC721Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _MockERC721.contract.Transact(opts, "safeTransferFrom", from, to, id)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id) payable returns()
func (_MockERC721 *MockERC721Session) SafeTransferFrom(from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _MockERC721.Contract.SafeTransferFrom(&_MockERC721.TransactOpts, from, to, id)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id) payable returns()
func (_MockERC721 *MockERC721TransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _MockERC721.Contract.SafeTransferFrom(&_MockERC721.TransactOpts, from, to, id)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, bytes data) payable returns()
func (_MockERC721 *MockERC721Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, data []byte) (*types.Transaction, error) {
	return _MockERC721.contract.Transact(opts, "safeTransferFrom0", from, to, id, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, bytes data) payable returns()
func (_MockERC721 *MockERC721Session) SafeTransferFrom0(from common.Address, to common.Address, id *big.Int, data []byte) (*types.Transaction, error) {
	return _MockERC721.Contract.SafeTransferFrom0(&_MockERC721.TransactOpts, from, to, id, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, bytes data) payable returns()
func (_MockERC721 *MockERC721TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, id *big.Int, data []byte) (*types.Transaction, error) {
	return _MockERC721.Contract.SafeTransferFrom0(&_MockERC721.TransactOpts, from, to, id, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MockERC721 *MockERC721Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _MockERC721.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MockERC721 *MockERC721Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _MockERC721.Contract.SetApprovalForAll(&_MockERC721.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MockERC721 *MockERC721TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _MockERC721.Contract.SetApprovalForAll(&_MockERC721.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 id) payable returns()
func (_MockERC721 *MockERC721Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _MockERC721.contract.Transact(opts, "transferFrom", from, to, id)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 id) payable returns()
func (_MockERC721 *MockERC721Session) TransferFrom(from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _MockERC721.Contract.TransferFrom(&_MockERC721.TransactOpts, from, to, id)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 id) payable returns()
func (_MockERC721 *MockERC721TransactorSession) TransferFrom(from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _MockERC721.Contract.TransferFrom(&_MockERC721.TransactOpts, from, to, id)
}

// MockERC721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MockERC721 contract.
type MockERC721ApprovalIterator struct {
	Event *MockERC721Approval // Event containing the contract specifics and raw log

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
func (it *MockERC721ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockERC721Approval)
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
		it.Event = new(MockERC721Approval)
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
func (it *MockERC721ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockERC721ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockERC721Approval represents a Approval event raised by the MockERC721 contract.
type MockERC721Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 indexed _tokenId)
func (_MockERC721 *MockERC721Filterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (*MockERC721ApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _MockERC721.contract.FilterLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MockERC721ApprovalIterator{contract: _MockERC721.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 indexed _tokenId)
func (_MockERC721 *MockERC721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MockERC721Approval, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _MockERC721.contract.WatchLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockERC721Approval)
				if err := _MockERC721.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 indexed _tokenId)
func (_MockERC721 *MockERC721Filterer) ParseApproval(log types.Log) (*MockERC721Approval, error) {
	event := new(MockERC721Approval)
	if err := _MockERC721.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockERC721ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the MockERC721 contract.
type MockERC721ApprovalForAllIterator struct {
	Event *MockERC721ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *MockERC721ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockERC721ApprovalForAll)
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
		it.Event = new(MockERC721ApprovalForAll)
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
func (it *MockERC721ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockERC721ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockERC721ApprovalForAll represents a ApprovalForAll event raised by the MockERC721 contract.
type MockERC721ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_MockERC721 *MockERC721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, _owner []common.Address, _operator []common.Address) (*MockERC721ApprovalForAllIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _MockERC721.contract.FilterLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return &MockERC721ApprovalForAllIterator{contract: _MockERC721.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_MockERC721 *MockERC721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *MockERC721ApprovalForAll, _owner []common.Address, _operator []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _MockERC721.contract.WatchLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockERC721ApprovalForAll)
				if err := _MockERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_MockERC721 *MockERC721Filterer) ParseApprovalForAll(log types.Log) (*MockERC721ApprovalForAll, error) {
	event := new(MockERC721ApprovalForAll)
	if err := _MockERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockERC721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MockERC721 contract.
type MockERC721TransferIterator struct {
	Event *MockERC721Transfer // Event containing the contract specifics and raw log

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
func (it *MockERC721TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockERC721Transfer)
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
		it.Event = new(MockERC721Transfer)
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
func (it *MockERC721TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockERC721TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockERC721Transfer represents a Transfer event raised by the MockERC721 contract.
type MockERC721Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
func (_MockERC721 *MockERC721Filterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (*MockERC721TransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _MockERC721.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MockERC721TransferIterator{contract: _MockERC721.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
func (_MockERC721 *MockERC721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MockERC721Transfer, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _MockERC721.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockERC721Transfer)
				if err := _MockERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
func (_MockERC721 *MockERC721Filterer) ParseTransfer(log types.Log) (*MockERC721Transfer, error) {
	event := new(MockERC721Transfer)
	if err := _MockERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
