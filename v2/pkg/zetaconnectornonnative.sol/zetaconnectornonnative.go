// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zetaconnectornonnative

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

// ZetaConnectorNonNativeMetaData contains all meta data concerning the ZetaConnectorNonNative contract.
var ZetaConnectorNonNativeMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"gateway_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"zetaToken_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tssAddress_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"admin_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TSS_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WITHDRAWER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gateway\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIGatewayEVM\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"receiveTokens\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxSupply\",\"inputs\":[{\"name\":\"maxSupply_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"internalSendHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAndCall\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"internalSendHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAndRevert\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"internalSendHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"zetaToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"MaxSupplyUpdated\",\"inputs\":[{\"name\":\"maxSupply\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdraw\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawAndCall\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawAndRevert\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExceedsMaxSupply\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x60c060405260001960035534801561001657600080fd5b5060405161155a38038061155a83398101604081905261003591610220565b60016000819055805460ff19169055838383836001600160a01b038416158061006557506001600160a01b038316155b8061007757506001600160a01b038216155b8061008957506001600160a01b038116155b156100a75760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b03808516608052831660a0526100c5600082610154565b506100f07f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e483610154565b5061011b7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb83610154565b506101467f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a82610154565b505050505050505050610274565b60008281526002602090815260408083206001600160a01b038516845290915281205460ff166101fa5760008381526002602090815260408083206001600160a01b03861684529091529020805460ff191660011790556101b23390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016101fe565b5060005b92915050565b80516001600160a01b038116811461021b57600080fd5b919050565b6000806000806080858703121561023657600080fd5b61023f85610204565b935061024d60208601610204565b925061025b60408601610204565b915061026960608601610204565b905092959194509250565b60805160a0516112826102d86000396000818161021d01528181610531015281816108240152818161097b01528181610ad00152610bf20152600081816101d1015281816104a1015281816105040152818161079401526107f701526112826000f3fe608060405234801561001057600080fd5b50600436106101775760003560e01c80635e3e9fef116100d857806391d148541161008c578063d547741f11610066578063d547741f14610386578063d5abeb0114610399578063e63ab1e9146103a257600080fd5b806391d1485414610311578063a217fddf14610357578063a783c7891461035f57600080fd5b8063743e0c9b116100bd578063743e0c9b146102cf5780638456cb59146102e257806385f438c1146102ea57600080fd5b80635e3e9fef146102a95780636f8b44b0146102bc57600080fd5b8063248a9ca31161012f57806336568abe1161011457806336568abe146102835780633f4ba83a146102965780635c975abb1461029e57600080fd5b8063248a9ca31461023f5780632f2ff15d1461027057600080fd5b8063106e629011610160578063106e6290146101b9578063116191b6146101cc57806321e093b11461021857600080fd5b806301ffc9a71461017c57806302d5c899146101a4575b600080fd5b61018f61018a366004610fb4565b6103c9565b60405190151581526020015b60405180910390f35b6101b76101b2366004611026565b610462565b005b6101b76101c73660046110b8565b6105f7565b6101f37f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161019b565b6101f37f000000000000000000000000000000000000000000000000000000000000000081565b61026261024d3660046110eb565b60009081526002602052604090206001015490565b60405190815260200161019b565b6101b761027e366004611104565b61069c565b6101b7610291366004611104565b6106c7565b6101b7610720565b60015460ff1661018f565b6101b76102b7366004611026565b610755565b6101b76102ca3660046110eb565b6108d0565b6101b76102dd3660046110eb565b61093e565b6101b76109e8565b6102627f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e481565b61018f61031f366004611104565b600091825260026020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b610262600081565b6102627f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb81565b6101b7610394366004611104565b610a1a565b61026260035481565b6102627f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061045c57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b61046a610a3f565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461049481610a82565b61049c610a8c565b6104c77f00000000000000000000000000000000000000000000000000000000000000008684610acb565b6040517fb8969bd400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063b8969bd490610561907f0000000000000000000000000000000000000000000000000000000000000000908a908a908a908a90600401611179565b600060405180830381600087803b15801561057b57600080fd5b505af115801561058f573d6000803e3d6000fd5b505050508573ffffffffffffffffffffffffffffffffffffffff167fba96f26bdda53eb8c8ba39045dfb4ff39753fbc7a6edcf250a88e75e78d102fe8686866040516105dd939291906111d6565b60405180910390a2506105f06001600055565b5050505050565b6105ff610a3f565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461062981610a82565b610631610a8c565b61063c848484610acb565b8373ffffffffffffffffffffffffffffffffffffffff167f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a94243648460405161068491815260200190565b60405180910390a2506106976001600055565b505050565b6000828152600260205260409020600101546106b781610a82565b6106c18383610c53565b50505050565b73ffffffffffffffffffffffffffffffffffffffff81163314610716576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6106978282610d53565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a61074a81610a82565b610752610e12565b50565b61075d610a3f565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461078781610a82565b61078f610a8c565b6107ba7f00000000000000000000000000000000000000000000000000000000000000008684610acb565b6040517f5131ab5900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690635131ab5990610854907f0000000000000000000000000000000000000000000000000000000000000000908a908a908a908a90600401611179565b600060405180830381600087803b15801561086e57600080fd5b505af1158015610882573d6000803e3d6000fd5b505050508573ffffffffffffffffffffffffffffffffffffffff167f7772f56296d3a5202974a45c61c9188d844ab4d6eeb18c851e4b8d5384ca6ced8686866040516105dd939291906111d6565b7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb6108fa81610a82565b610902610a8c565b60038290556040518281527f7810bd47de260c3e9ee10061cf438099dd12256c79485f12f94dbccc981e806c9060200160405180910390a15050565b610946610a8c565b6040517f79cc6790000000000000000000000000000000000000000000000000000000008152336004820152602481018290527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906379cc679090604401600060405180830381600087803b1580156109d457600080fd5b505af11580156105f0573d6000803e3d6000fd5b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a610a1281610a82565b610752610e8f565b600082815260026020526040902060010154610a3581610a82565b6106c18383610d53565b600260005403610a7b576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600055565b6107528133610ee8565b60015460ff1615610ac9576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6003547f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b39573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b5d91906111f9565b610b679084611212565b1115610b9f576040517fc30436e900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f1e458bee00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff848116600483015260248201849052604482018390527f00000000000000000000000000000000000000000000000000000000000000001690631e458bee90606401600060405180830381600087803b158015610c3657600080fd5b505af1158015610c4a573d6000803e3d6000fd5b50505050505050565b600082815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff16610d4b57600083815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff86168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055610ce93390565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a450600161045c565b50600061045c565b600082815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff1615610d4b57600083815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8616808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a450600161045c565b610e1a610f78565b600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b610e97610a8c565b600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016811790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833610e65565b600082815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16610f74576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024810183905260440160405180910390fd5b5050565b60015460ff16610ac9576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060208284031215610fc657600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114610ff657600080fd5b9392505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461102157600080fd5b919050565b60008060008060006080868803121561103e57600080fd5b61104786610ffd565b945060208601359350604086013567ffffffffffffffff81111561106a57600080fd5b8601601f8101881361107b57600080fd5b803567ffffffffffffffff81111561109257600080fd5b8860208284010111156110a457600080fd5b959894975060200195606001359392505050565b6000806000606084860312156110cd57600080fd5b6110d684610ffd565b95602085013595506040909401359392505050565b6000602082840312156110fd57600080fd5b5035919050565b6000806040838503121561111757600080fd5b8235915061112760208401610ffd565b90509250929050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff8616815273ffffffffffffffffffffffffffffffffffffffff851660208201528360408201526080606082015260006111cb608083018486611130565b979650505050505050565b8381526040602082015260006111f0604083018486611130565b95945050505050565b60006020828403121561120b57600080fd5b5051919050565b8082018082111561045c577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea2646970667358221220de5fa55e0ca79ea2f08903481d01ea00bbbf346b9cfba44e12a2dd6f3a6f2e8a64736f6c634300081a0033",
}

// ZetaConnectorNonNativeABI is the input ABI used to generate the binding from.
// Deprecated: Use ZetaConnectorNonNativeMetaData.ABI instead.
var ZetaConnectorNonNativeABI = ZetaConnectorNonNativeMetaData.ABI

// ZetaConnectorNonNativeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZetaConnectorNonNativeMetaData.Bin instead.
var ZetaConnectorNonNativeBin = ZetaConnectorNonNativeMetaData.Bin

// DeployZetaConnectorNonNative deploys a new Ethereum contract, binding an instance of ZetaConnectorNonNative to it.
func DeployZetaConnectorNonNative(auth *bind.TransactOpts, backend bind.ContractBackend, gateway_ common.Address, zetaToken_ common.Address, tssAddress_ common.Address, admin_ common.Address) (common.Address, *types.Transaction, *ZetaConnectorNonNative, error) {
	parsed, err := ZetaConnectorNonNativeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZetaConnectorNonNativeBin), backend, gateway_, zetaToken_, tssAddress_, admin_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZetaConnectorNonNative{ZetaConnectorNonNativeCaller: ZetaConnectorNonNativeCaller{contract: contract}, ZetaConnectorNonNativeTransactor: ZetaConnectorNonNativeTransactor{contract: contract}, ZetaConnectorNonNativeFilterer: ZetaConnectorNonNativeFilterer{contract: contract}}, nil
}

// ZetaConnectorNonNative is an auto generated Go binding around an Ethereum contract.
type ZetaConnectorNonNative struct {
	ZetaConnectorNonNativeCaller     // Read-only binding to the contract
	ZetaConnectorNonNativeTransactor // Write-only binding to the contract
	ZetaConnectorNonNativeFilterer   // Log filterer for contract events
}

// ZetaConnectorNonNativeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZetaConnectorNonNativeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorNonNativeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZetaConnectorNonNativeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorNonNativeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZetaConnectorNonNativeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorNonNativeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZetaConnectorNonNativeSession struct {
	Contract     *ZetaConnectorNonNative // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ZetaConnectorNonNativeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZetaConnectorNonNativeCallerSession struct {
	Contract *ZetaConnectorNonNativeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// ZetaConnectorNonNativeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZetaConnectorNonNativeTransactorSession struct {
	Contract     *ZetaConnectorNonNativeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// ZetaConnectorNonNativeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZetaConnectorNonNativeRaw struct {
	Contract *ZetaConnectorNonNative // Generic contract binding to access the raw methods on
}

// ZetaConnectorNonNativeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZetaConnectorNonNativeCallerRaw struct {
	Contract *ZetaConnectorNonNativeCaller // Generic read-only contract binding to access the raw methods on
}

// ZetaConnectorNonNativeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZetaConnectorNonNativeTransactorRaw struct {
	Contract *ZetaConnectorNonNativeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZetaConnectorNonNative creates a new instance of ZetaConnectorNonNative, bound to a specific deployed contract.
func NewZetaConnectorNonNative(address common.Address, backend bind.ContractBackend) (*ZetaConnectorNonNative, error) {
	contract, err := bindZetaConnectorNonNative(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNative{ZetaConnectorNonNativeCaller: ZetaConnectorNonNativeCaller{contract: contract}, ZetaConnectorNonNativeTransactor: ZetaConnectorNonNativeTransactor{contract: contract}, ZetaConnectorNonNativeFilterer: ZetaConnectorNonNativeFilterer{contract: contract}}, nil
}

// NewZetaConnectorNonNativeCaller creates a new read-only instance of ZetaConnectorNonNative, bound to a specific deployed contract.
func NewZetaConnectorNonNativeCaller(address common.Address, caller bind.ContractCaller) (*ZetaConnectorNonNativeCaller, error) {
	contract, err := bindZetaConnectorNonNative(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeCaller{contract: contract}, nil
}

// NewZetaConnectorNonNativeTransactor creates a new write-only instance of ZetaConnectorNonNative, bound to a specific deployed contract.
func NewZetaConnectorNonNativeTransactor(address common.Address, transactor bind.ContractTransactor) (*ZetaConnectorNonNativeTransactor, error) {
	contract, err := bindZetaConnectorNonNative(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeTransactor{contract: contract}, nil
}

// NewZetaConnectorNonNativeFilterer creates a new log filterer instance of ZetaConnectorNonNative, bound to a specific deployed contract.
func NewZetaConnectorNonNativeFilterer(address common.Address, filterer bind.ContractFilterer) (*ZetaConnectorNonNativeFilterer, error) {
	contract, err := bindZetaConnectorNonNative(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeFilterer{contract: contract}, nil
}

// bindZetaConnectorNonNative binds a generic wrapper to an already deployed contract.
func bindZetaConnectorNonNative(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZetaConnectorNonNativeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaConnectorNonNative.Contract.ZetaConnectorNonNativeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNonNative.Contract.ZetaConnectorNonNativeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaConnectorNonNative.Contract.ZetaConnectorNonNativeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaConnectorNonNative.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNonNative.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaConnectorNonNative.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZetaConnectorNonNative.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ZetaConnectorNonNative.Contract.DEFAULTADMINROLE(&_ZetaConnectorNonNative.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ZetaConnectorNonNative.Contract.DEFAULTADMINROLE(&_ZetaConnectorNonNative.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZetaConnectorNonNative.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeSession) PAUSERROLE() ([32]byte, error) {
	return _ZetaConnectorNonNative.Contract.PAUSERROLE(&_ZetaConnectorNonNative.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeCallerSession) PAUSERROLE() ([32]byte, error) {
	return _ZetaConnectorNonNative.Contract.PAUSERROLE(&_ZetaConnectorNonNative.CallOpts)
}

// TSSROLE is a free data retrieval call binding the contract method 0xa783c789.
//
// Solidity: function TSS_ROLE() view returns(bytes32)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeCaller) TSSROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZetaConnectorNonNative.contract.Call(opts, &out, "TSS_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TSSROLE is a free data retrieval call binding the contract method 0xa783c789.
//
// Solidity: function TSS_ROLE() view returns(bytes32)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeSession) TSSROLE() ([32]byte, error) {
	return _ZetaConnectorNonNative.Contract.TSSROLE(&_ZetaConnectorNonNative.CallOpts)
}

// TSSROLE is a free data retrieval call binding the contract method 0xa783c789.
//
// Solidity: function TSS_ROLE() view returns(bytes32)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeCallerSession) TSSROLE() ([32]byte, error) {
	return _ZetaConnectorNonNative.Contract.TSSROLE(&_ZetaConnectorNonNative.CallOpts)
}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeCaller) WITHDRAWERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZetaConnectorNonNative.contract.Call(opts, &out, "WITHDRAWER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeSession) WITHDRAWERROLE() ([32]byte, error) {
	return _ZetaConnectorNonNative.Contract.WITHDRAWERROLE(&_ZetaConnectorNonNative.CallOpts)
}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeCallerSession) WITHDRAWERROLE() ([32]byte, error) {
	return _ZetaConnectorNonNative.Contract.WITHDRAWERROLE(&_ZetaConnectorNonNative.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeCaller) Gateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorNonNative.contract.Call(opts, &out, "gateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeSession) Gateway() (common.Address, error) {
	return _ZetaConnectorNonNative.Contract.Gateway(&_ZetaConnectorNonNative.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeCallerSession) Gateway() (common.Address, error) {
	return _ZetaConnectorNonNative.Contract.Gateway(&_ZetaConnectorNonNative.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ZetaConnectorNonNative.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ZetaConnectorNonNative.Contract.GetRoleAdmin(&_ZetaConnectorNonNative.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ZetaConnectorNonNative.Contract.GetRoleAdmin(&_ZetaConnectorNonNative.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ZetaConnectorNonNative.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ZetaConnectorNonNative.Contract.HasRole(&_ZetaConnectorNonNative.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ZetaConnectorNonNative.Contract.HasRole(&_ZetaConnectorNonNative.CallOpts, role, account)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeCaller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZetaConnectorNonNative.contract.Call(opts, &out, "maxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeSession) MaxSupply() (*big.Int, error) {
	return _ZetaConnectorNonNative.Contract.MaxSupply(&_ZetaConnectorNonNative.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeCallerSession) MaxSupply() (*big.Int, error) {
	return _ZetaConnectorNonNative.Contract.MaxSupply(&_ZetaConnectorNonNative.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ZetaConnectorNonNative.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeSession) Paused() (bool, error) {
	return _ZetaConnectorNonNative.Contract.Paused(&_ZetaConnectorNonNative.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeCallerSession) Paused() (bool, error) {
	return _ZetaConnectorNonNative.Contract.Paused(&_ZetaConnectorNonNative.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ZetaConnectorNonNative.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ZetaConnectorNonNative.Contract.SupportsInterface(&_ZetaConnectorNonNative.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ZetaConnectorNonNative.Contract.SupportsInterface(&_ZetaConnectorNonNative.CallOpts, interfaceId)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeCaller) ZetaToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorNonNative.contract.Call(opts, &out, "zetaToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeSession) ZetaToken() (common.Address, error) {
	return _ZetaConnectorNonNative.Contract.ZetaToken(&_ZetaConnectorNonNative.CallOpts)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeCallerSession) ZetaToken() (common.Address, error) {
	return _ZetaConnectorNonNative.Contract.ZetaToken(&_ZetaConnectorNonNative.CallOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNonNative.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNonNative.Contract.GrantRole(&_ZetaConnectorNonNative.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNonNative.Contract.GrantRole(&_ZetaConnectorNonNative.TransactOpts, role, account)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNonNative.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeSession) Pause() (*types.Transaction, error) {
	return _ZetaConnectorNonNative.Contract.Pause(&_ZetaConnectorNonNative.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeTransactorSession) Pause() (*types.Transaction, error) {
	return _ZetaConnectorNonNative.Contract.Pause(&_ZetaConnectorNonNative.TransactOpts)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x743e0c9b.
//
// Solidity: function receiveTokens(uint256 amount) returns()
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeTransactor) ReceiveTokens(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ZetaConnectorNonNative.contract.Transact(opts, "receiveTokens", amount)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x743e0c9b.
//
// Solidity: function receiveTokens(uint256 amount) returns()
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeSession) ReceiveTokens(amount *big.Int) (*types.Transaction, error) {
	return _ZetaConnectorNonNative.Contract.ReceiveTokens(&_ZetaConnectorNonNative.TransactOpts, amount)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x743e0c9b.
//
// Solidity: function receiveTokens(uint256 amount) returns()
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeTransactorSession) ReceiveTokens(amount *big.Int) (*types.Transaction, error) {
	return _ZetaConnectorNonNative.Contract.ReceiveTokens(&_ZetaConnectorNonNative.TransactOpts, amount)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNonNative.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNonNative.Contract.RenounceRole(&_ZetaConnectorNonNative.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNonNative.Contract.RenounceRole(&_ZetaConnectorNonNative.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNonNative.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNonNative.Contract.RevokeRole(&_ZetaConnectorNonNative.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNonNative.Contract.RevokeRole(&_ZetaConnectorNonNative.TransactOpts, role, account)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 maxSupply_) returns()
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeTransactor) SetMaxSupply(opts *bind.TransactOpts, maxSupply_ *big.Int) (*types.Transaction, error) {
	return _ZetaConnectorNonNative.contract.Transact(opts, "setMaxSupply", maxSupply_)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 maxSupply_) returns()
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeSession) SetMaxSupply(maxSupply_ *big.Int) (*types.Transaction, error) {
	return _ZetaConnectorNonNative.Contract.SetMaxSupply(&_ZetaConnectorNonNative.TransactOpts, maxSupply_)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 maxSupply_) returns()
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeTransactorSession) SetMaxSupply(maxSupply_ *big.Int) (*types.Transaction, error) {
	return _ZetaConnectorNonNative.Contract.SetMaxSupply(&_ZetaConnectorNonNative.TransactOpts, maxSupply_)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNonNative.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeSession) Unpause() (*types.Transaction, error) {
	return _ZetaConnectorNonNative.Contract.Unpause(&_ZetaConnectorNonNative.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeTransactorSession) Unpause() (*types.Transaction, error) {
	return _ZetaConnectorNonNative.Contract.Unpause(&_ZetaConnectorNonNative.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x106e6290.
//
// Solidity: function withdraw(address to, uint256 amount, bytes32 internalSendHash) returns()
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeTransactor) Withdraw(opts *bind.TransactOpts, to common.Address, amount *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNonNative.contract.Transact(opts, "withdraw", to, amount, internalSendHash)
}

// Withdraw is a paid mutator transaction binding the contract method 0x106e6290.
//
// Solidity: function withdraw(address to, uint256 amount, bytes32 internalSendHash) returns()
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeSession) Withdraw(to common.Address, amount *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNonNative.Contract.Withdraw(&_ZetaConnectorNonNative.TransactOpts, to, amount, internalSendHash)
}

// Withdraw is a paid mutator transaction binding the contract method 0x106e6290.
//
// Solidity: function withdraw(address to, uint256 amount, bytes32 internalSendHash) returns()
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeTransactorSession) Withdraw(to common.Address, amount *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNonNative.Contract.Withdraw(&_ZetaConnectorNonNative.TransactOpts, to, amount, internalSendHash)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x5e3e9fef.
//
// Solidity: function withdrawAndCall(address to, uint256 amount, bytes data, bytes32 internalSendHash) returns()
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeTransactor) WithdrawAndCall(opts *bind.TransactOpts, to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNonNative.contract.Transact(opts, "withdrawAndCall", to, amount, data, internalSendHash)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x5e3e9fef.
//
// Solidity: function withdrawAndCall(address to, uint256 amount, bytes data, bytes32 internalSendHash) returns()
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeSession) WithdrawAndCall(to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNonNative.Contract.WithdrawAndCall(&_ZetaConnectorNonNative.TransactOpts, to, amount, data, internalSendHash)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x5e3e9fef.
//
// Solidity: function withdrawAndCall(address to, uint256 amount, bytes data, bytes32 internalSendHash) returns()
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeTransactorSession) WithdrawAndCall(to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNonNative.Contract.WithdrawAndCall(&_ZetaConnectorNonNative.TransactOpts, to, amount, data, internalSendHash)
}

// WithdrawAndRevert is a paid mutator transaction binding the contract method 0x02d5c899.
//
// Solidity: function withdrawAndRevert(address to, uint256 amount, bytes data, bytes32 internalSendHash) returns()
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeTransactor) WithdrawAndRevert(opts *bind.TransactOpts, to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNonNative.contract.Transact(opts, "withdrawAndRevert", to, amount, data, internalSendHash)
}

// WithdrawAndRevert is a paid mutator transaction binding the contract method 0x02d5c899.
//
// Solidity: function withdrawAndRevert(address to, uint256 amount, bytes data, bytes32 internalSendHash) returns()
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeSession) WithdrawAndRevert(to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNonNative.Contract.WithdrawAndRevert(&_ZetaConnectorNonNative.TransactOpts, to, amount, data, internalSendHash)
}

// WithdrawAndRevert is a paid mutator transaction binding the contract method 0x02d5c899.
//
// Solidity: function withdrawAndRevert(address to, uint256 amount, bytes data, bytes32 internalSendHash) returns()
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeTransactorSession) WithdrawAndRevert(to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNonNative.Contract.WithdrawAndRevert(&_ZetaConnectorNonNative.TransactOpts, to, amount, data, internalSendHash)
}

// ZetaConnectorNonNativeMaxSupplyUpdatedIterator is returned from FilterMaxSupplyUpdated and is used to iterate over the raw logs and unpacked data for MaxSupplyUpdated events raised by the ZetaConnectorNonNative contract.
type ZetaConnectorNonNativeMaxSupplyUpdatedIterator struct {
	Event *ZetaConnectorNonNativeMaxSupplyUpdated // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNonNativeMaxSupplyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeMaxSupplyUpdated)
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
		it.Event = new(ZetaConnectorNonNativeMaxSupplyUpdated)
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
func (it *ZetaConnectorNonNativeMaxSupplyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeMaxSupplyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeMaxSupplyUpdated represents a MaxSupplyUpdated event raised by the ZetaConnectorNonNative contract.
type ZetaConnectorNonNativeMaxSupplyUpdated struct {
	MaxSupply *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMaxSupplyUpdated is a free log retrieval operation binding the contract event 0x7810bd47de260c3e9ee10061cf438099dd12256c79485f12f94dbccc981e806c.
//
// Solidity: event MaxSupplyUpdated(uint256 maxSupply)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeFilterer) FilterMaxSupplyUpdated(opts *bind.FilterOpts) (*ZetaConnectorNonNativeMaxSupplyUpdatedIterator, error) {

	logs, sub, err := _ZetaConnectorNonNative.contract.FilterLogs(opts, "MaxSupplyUpdated")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeMaxSupplyUpdatedIterator{contract: _ZetaConnectorNonNative.contract, event: "MaxSupplyUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxSupplyUpdated is a free log subscription operation binding the contract event 0x7810bd47de260c3e9ee10061cf438099dd12256c79485f12f94dbccc981e806c.
//
// Solidity: event MaxSupplyUpdated(uint256 maxSupply)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeFilterer) WatchMaxSupplyUpdated(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeMaxSupplyUpdated) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNonNative.contract.WatchLogs(opts, "MaxSupplyUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeMaxSupplyUpdated)
				if err := _ZetaConnectorNonNative.contract.UnpackLog(event, "MaxSupplyUpdated", log); err != nil {
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

// ParseMaxSupplyUpdated is a log parse operation binding the contract event 0x7810bd47de260c3e9ee10061cf438099dd12256c79485f12f94dbccc981e806c.
//
// Solidity: event MaxSupplyUpdated(uint256 maxSupply)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeFilterer) ParseMaxSupplyUpdated(log types.Log) (*ZetaConnectorNonNativeMaxSupplyUpdated, error) {
	event := new(ZetaConnectorNonNativeMaxSupplyUpdated)
	if err := _ZetaConnectorNonNative.contract.UnpackLog(event, "MaxSupplyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ZetaConnectorNonNative contract.
type ZetaConnectorNonNativePausedIterator struct {
	Event *ZetaConnectorNonNativePaused // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNonNativePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativePaused)
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
		it.Event = new(ZetaConnectorNonNativePaused)
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
func (it *ZetaConnectorNonNativePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativePaused represents a Paused event raised by the ZetaConnectorNonNative contract.
type ZetaConnectorNonNativePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeFilterer) FilterPaused(opts *bind.FilterOpts) (*ZetaConnectorNonNativePausedIterator, error) {

	logs, sub, err := _ZetaConnectorNonNative.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativePausedIterator{contract: _ZetaConnectorNonNative.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativePaused) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNonNative.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativePaused)
				if err := _ZetaConnectorNonNative.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeFilterer) ParsePaused(log types.Log) (*ZetaConnectorNonNativePaused, error) {
	event := new(ZetaConnectorNonNativePaused)
	if err := _ZetaConnectorNonNative.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ZetaConnectorNonNative contract.
type ZetaConnectorNonNativeRoleAdminChangedIterator struct {
	Event *ZetaConnectorNonNativeRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNonNativeRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeRoleAdminChanged)
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
		it.Event = new(ZetaConnectorNonNativeRoleAdminChanged)
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
func (it *ZetaConnectorNonNativeRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeRoleAdminChanged represents a RoleAdminChanged event raised by the ZetaConnectorNonNative contract.
type ZetaConnectorNonNativeRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ZetaConnectorNonNativeRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ZetaConnectorNonNative.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeRoleAdminChangedIterator{contract: _ZetaConnectorNonNative.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ZetaConnectorNonNative.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeRoleAdminChanged)
				if err := _ZetaConnectorNonNative.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeFilterer) ParseRoleAdminChanged(log types.Log) (*ZetaConnectorNonNativeRoleAdminChanged, error) {
	event := new(ZetaConnectorNonNativeRoleAdminChanged)
	if err := _ZetaConnectorNonNative.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ZetaConnectorNonNative contract.
type ZetaConnectorNonNativeRoleGrantedIterator struct {
	Event *ZetaConnectorNonNativeRoleGranted // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNonNativeRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeRoleGranted)
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
		it.Event = new(ZetaConnectorNonNativeRoleGranted)
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
func (it *ZetaConnectorNonNativeRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeRoleGranted represents a RoleGranted event raised by the ZetaConnectorNonNative contract.
type ZetaConnectorNonNativeRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ZetaConnectorNonNativeRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ZetaConnectorNonNative.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeRoleGrantedIterator{contract: _ZetaConnectorNonNative.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ZetaConnectorNonNative.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeRoleGranted)
				if err := _ZetaConnectorNonNative.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeFilterer) ParseRoleGranted(log types.Log) (*ZetaConnectorNonNativeRoleGranted, error) {
	event := new(ZetaConnectorNonNativeRoleGranted)
	if err := _ZetaConnectorNonNative.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ZetaConnectorNonNative contract.
type ZetaConnectorNonNativeRoleRevokedIterator struct {
	Event *ZetaConnectorNonNativeRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNonNativeRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeRoleRevoked)
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
		it.Event = new(ZetaConnectorNonNativeRoleRevoked)
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
func (it *ZetaConnectorNonNativeRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeRoleRevoked represents a RoleRevoked event raised by the ZetaConnectorNonNative contract.
type ZetaConnectorNonNativeRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ZetaConnectorNonNativeRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ZetaConnectorNonNative.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeRoleRevokedIterator{contract: _ZetaConnectorNonNative.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ZetaConnectorNonNative.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeRoleRevoked)
				if err := _ZetaConnectorNonNative.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeFilterer) ParseRoleRevoked(log types.Log) (*ZetaConnectorNonNativeRoleRevoked, error) {
	event := new(ZetaConnectorNonNativeRoleRevoked)
	if err := _ZetaConnectorNonNative.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ZetaConnectorNonNative contract.
type ZetaConnectorNonNativeUnpausedIterator struct {
	Event *ZetaConnectorNonNativeUnpaused // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNonNativeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeUnpaused)
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
		it.Event = new(ZetaConnectorNonNativeUnpaused)
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
func (it *ZetaConnectorNonNativeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeUnpaused represents a Unpaused event raised by the ZetaConnectorNonNative contract.
type ZetaConnectorNonNativeUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ZetaConnectorNonNativeUnpausedIterator, error) {

	logs, sub, err := _ZetaConnectorNonNative.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeUnpausedIterator{contract: _ZetaConnectorNonNative.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeUnpaused) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNonNative.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeUnpaused)
				if err := _ZetaConnectorNonNative.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeFilterer) ParseUnpaused(log types.Log) (*ZetaConnectorNonNativeUnpaused, error) {
	event := new(ZetaConnectorNonNativeUnpaused)
	if err := _ZetaConnectorNonNative.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the ZetaConnectorNonNative contract.
type ZetaConnectorNonNativeWithdrawIterator struct {
	Event *ZetaConnectorNonNativeWithdraw // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNonNativeWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeWithdraw)
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
		it.Event = new(ZetaConnectorNonNativeWithdraw)
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
func (it *ZetaConnectorNonNativeWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeWithdraw represents a Withdraw event raised by the ZetaConnectorNonNative contract.
type ZetaConnectorNonNativeWithdraw struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed to, uint256 amount)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeFilterer) FilterWithdraw(opts *bind.FilterOpts, to []common.Address) (*ZetaConnectorNonNativeWithdrawIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNonNative.contract.FilterLogs(opts, "Withdraw", toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeWithdrawIterator{contract: _ZetaConnectorNonNative.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed to, uint256 amount)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeWithdraw, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNonNative.contract.WatchLogs(opts, "Withdraw", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeWithdraw)
				if err := _ZetaConnectorNonNative.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed to, uint256 amount)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeFilterer) ParseWithdraw(log types.Log) (*ZetaConnectorNonNativeWithdraw, error) {
	event := new(ZetaConnectorNonNativeWithdraw)
	if err := _ZetaConnectorNonNative.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeWithdrawAndCallIterator is returned from FilterWithdrawAndCall and is used to iterate over the raw logs and unpacked data for WithdrawAndCall events raised by the ZetaConnectorNonNative contract.
type ZetaConnectorNonNativeWithdrawAndCallIterator struct {
	Event *ZetaConnectorNonNativeWithdrawAndCall // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNonNativeWithdrawAndCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeWithdrawAndCall)
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
		it.Event = new(ZetaConnectorNonNativeWithdrawAndCall)
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
func (it *ZetaConnectorNonNativeWithdrawAndCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeWithdrawAndCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeWithdrawAndCall represents a WithdrawAndCall event raised by the ZetaConnectorNonNative contract.
type ZetaConnectorNonNativeWithdrawAndCall struct {
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawAndCall is a free log retrieval operation binding the contract event 0x7772f56296d3a5202974a45c61c9188d844ab4d6eeb18c851e4b8d5384ca6ced.
//
// Solidity: event WithdrawAndCall(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeFilterer) FilterWithdrawAndCall(opts *bind.FilterOpts, to []common.Address) (*ZetaConnectorNonNativeWithdrawAndCallIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNonNative.contract.FilterLogs(opts, "WithdrawAndCall", toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeWithdrawAndCallIterator{contract: _ZetaConnectorNonNative.contract, event: "WithdrawAndCall", logs: logs, sub: sub}, nil
}

// WatchWithdrawAndCall is a free log subscription operation binding the contract event 0x7772f56296d3a5202974a45c61c9188d844ab4d6eeb18c851e4b8d5384ca6ced.
//
// Solidity: event WithdrawAndCall(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeFilterer) WatchWithdrawAndCall(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeWithdrawAndCall, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNonNative.contract.WatchLogs(opts, "WithdrawAndCall", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeWithdrawAndCall)
				if err := _ZetaConnectorNonNative.contract.UnpackLog(event, "WithdrawAndCall", log); err != nil {
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

// ParseWithdrawAndCall is a log parse operation binding the contract event 0x7772f56296d3a5202974a45c61c9188d844ab4d6eeb18c851e4b8d5384ca6ced.
//
// Solidity: event WithdrawAndCall(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeFilterer) ParseWithdrawAndCall(log types.Log) (*ZetaConnectorNonNativeWithdrawAndCall, error) {
	event := new(ZetaConnectorNonNativeWithdrawAndCall)
	if err := _ZetaConnectorNonNative.contract.UnpackLog(event, "WithdrawAndCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeWithdrawAndRevertIterator is returned from FilterWithdrawAndRevert and is used to iterate over the raw logs and unpacked data for WithdrawAndRevert events raised by the ZetaConnectorNonNative contract.
type ZetaConnectorNonNativeWithdrawAndRevertIterator struct {
	Event *ZetaConnectorNonNativeWithdrawAndRevert // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNonNativeWithdrawAndRevertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeWithdrawAndRevert)
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
		it.Event = new(ZetaConnectorNonNativeWithdrawAndRevert)
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
func (it *ZetaConnectorNonNativeWithdrawAndRevertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeWithdrawAndRevertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeWithdrawAndRevert represents a WithdrawAndRevert event raised by the ZetaConnectorNonNative contract.
type ZetaConnectorNonNativeWithdrawAndRevert struct {
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawAndRevert is a free log retrieval operation binding the contract event 0xba96f26bdda53eb8c8ba39045dfb4ff39753fbc7a6edcf250a88e75e78d102fe.
//
// Solidity: event WithdrawAndRevert(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeFilterer) FilterWithdrawAndRevert(opts *bind.FilterOpts, to []common.Address) (*ZetaConnectorNonNativeWithdrawAndRevertIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNonNative.contract.FilterLogs(opts, "WithdrawAndRevert", toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeWithdrawAndRevertIterator{contract: _ZetaConnectorNonNative.contract, event: "WithdrawAndRevert", logs: logs, sub: sub}, nil
}

// WatchWithdrawAndRevert is a free log subscription operation binding the contract event 0xba96f26bdda53eb8c8ba39045dfb4ff39753fbc7a6edcf250a88e75e78d102fe.
//
// Solidity: event WithdrawAndRevert(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeFilterer) WatchWithdrawAndRevert(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeWithdrawAndRevert, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNonNative.contract.WatchLogs(opts, "WithdrawAndRevert", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeWithdrawAndRevert)
				if err := _ZetaConnectorNonNative.contract.UnpackLog(event, "WithdrawAndRevert", log); err != nil {
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

// ParseWithdrawAndRevert is a log parse operation binding the contract event 0xba96f26bdda53eb8c8ba39045dfb4ff39753fbc7a6edcf250a88e75e78d102fe.
//
// Solidity: event WithdrawAndRevert(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeFilterer) ParseWithdrawAndRevert(log types.Log) (*ZetaConnectorNonNativeWithdrawAndRevert, error) {
	event := new(ZetaConnectorNonNativeWithdrawAndRevert)
	if err := _ZetaConnectorNonNative.contract.UnpackLog(event, "WithdrawAndRevert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
