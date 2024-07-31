// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zetaconnectornative

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

// ZetaConnectorNativeMetaData contains all meta data concerning the ZetaConnectorNative contract.
var ZetaConnectorNativeMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_gateway\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_zetaToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tssAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_admin\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TSS_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gateway\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIGatewayEVM\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"receiveTokens\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tssAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"internalSendHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAndCall\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"internalSendHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAndRevert\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"internalSendHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"zetaToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdraw\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawAndCall\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawAndRevert\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x60c060405234801561001057600080fd5b5060405161163738038061163783398101604081905261002f916101f5565b60016000819055805460ff19169055838383836001600160a01b038416158061005f57506001600160a01b038316155b8061007157506001600160a01b038216155b1561008f5760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b0384811660805283811660a052600380546001600160a01b0319169184169190911790556100c5600082610129565b506100f07f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb83610129565b5061011b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a82610129565b505050505050505050610249565b60008281526002602090815260408083206001600160a01b038516845290915281205460ff166101cf5760008381526002602090815260408083206001600160a01b03861684529091529020805460ff191660011790556101873390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016101d3565b5060005b92915050565b80516001600160a01b03811681146101f057600080fd5b919050565b6000806000806080858703121561020b57600080fd5b610214856101d9565b9350610222602086016101d9565b9250610230604086016101d9565b915061023e606086016101d9565b905092959194509250565b60805160a0516113836102b4600039600081816101f70152818161046d0152818161051e01528181610638015281816107cc0152818161087d015261094b0152600081816101ab0152818161048f015281816104f1015281816107ee015261085001526113836000f3fe608060405234801561001057600080fd5b50600436106101515760003560e01c80635b112591116100cd57806391d1485411610081578063a783c78911610066578063a783c7891461031f578063d547741f14610346578063e63ab1e91461035957600080fd5b806391d14854146102d1578063a217fddf1461031757600080fd5b80635e3e9fef116100b25780635e3e9fef146102a3578063743e0c9b146102b65780638456cb59146102c957600080fd5b80635b112591146102785780635c975abb1461029857600080fd5b806321e093b1116101245780632f2ff15d116101095780632f2ff15d1461024a57806336568abe1461025d5780633f4ba83a1461027057600080fd5b806321e093b1146101f2578063248a9ca31461021957600080fd5b806301ffc9a71461015657806302d5c8991461017e578063106e629014610193578063116191b6146101a6575b600080fd5b6101696101643660046110be565b610380565b60405190151581526020015b60405180910390f35b61019161018c366004611129565b610419565b005b6101916101a13660046111bb565b6105e4565b6101cd7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610175565b6101cd7f000000000000000000000000000000000000000000000000000000000000000081565b61023c6102273660046111ee565b60009081526002602052604090206001015490565b604051908152602001610175565b610191610258366004611207565b6106bf565b61019161026b366004611207565b6106ea565b610191610743565b6003546101cd9073ffffffffffffffffffffffffffffffffffffffff1681565b60015460ff16610169565b6101916102b1366004611129565b610778565b6101916102c43660046111ee565b610929565b610191610973565b6101696102df366004611207565b600091825260026020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b61023c600081565b61023c7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb81565b610191610354366004611207565b6109a5565b61023c7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061041357507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b6104216109ca565b7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb61044b81610a0d565b610453610a17565b6104b473ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000167f000000000000000000000000000000000000000000000000000000000000000087610a56565b6040517fb8969bd400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063b8969bd49061054e907f0000000000000000000000000000000000000000000000000000000000000000908a908a908a908a9060040161127c565b600060405180830381600087803b15801561056857600080fd5b505af115801561057c573d6000803e3d6000fd5b505050508573ffffffffffffffffffffffffffffffffffffffff167fba96f26bdda53eb8c8ba39045dfb4ff39753fbc7a6edcf250a88e75e78d102fe8686866040516105ca939291906112d9565b60405180910390a2506105dd6001600055565b5050505050565b6105ec6109ca565b7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb61061681610a0d565b61061e610a17565b61065f73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168585610a56565b8373ffffffffffffffffffffffffffffffffffffffff167f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364846040516106a791815260200190565b60405180910390a2506106ba6001600055565b505050565b6000828152600260205260409020600101546106da81610a0d565b6106e48383610ad7565b50505050565b73ffffffffffffffffffffffffffffffffffffffff81163314610739576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6106ba8282610bd7565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a61076d81610a0d565b610775610c96565b50565b6107806109ca565b7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb6107aa81610a0d565b6107b2610a17565b61081373ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000167f000000000000000000000000000000000000000000000000000000000000000087610a56565b6040517f5131ab5900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690635131ab59906108ad907f0000000000000000000000000000000000000000000000000000000000000000908a908a908a908a9060040161127c565b600060405180830381600087803b1580156108c757600080fd5b505af11580156108db573d6000803e3d6000fd5b505050508573ffffffffffffffffffffffffffffffffffffffff167f7772f56296d3a5202974a45c61c9188d844ab4d6eeb18c851e4b8d5384ca6ced8686866040516105ca939291906112d9565b610931610a17565b61077573ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016333084610d13565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a61099d81610a0d565b610775610d59565b6000828152600260205260409020600101546109c081610a0d565b6106e48383610bd7565b600260005403610a06576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600055565b6107758133610db2565b60015460ff1615610a54576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b60405173ffffffffffffffffffffffffffffffffffffffff8381166024830152604482018390526106ba91859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610e43565b600082815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff16610bcf57600083815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff86168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055610b6d3390565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001610413565b506000610413565b600082815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff1615610bcf57600083815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8616808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4506001610413565b610c9e610ed9565b600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b60405173ffffffffffffffffffffffffffffffffffffffff84811660248301528381166044830152606482018390526106e49186918216906323b872dd90608401610a90565b610d61610a17565b600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016811790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833610ce9565b600082815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16610e3f576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602481018390526044015b60405180910390fd5b5050565b6000610e6573ffffffffffffffffffffffffffffffffffffffff841683610f15565b90508051600014158015610e8a575080806020019051810190610e8891906112fc565b155b156106ba576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84166004820152602401610e36565b60015460ff16610a54576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6060610f2383836000610f2a565b9392505050565b606081471015610f68576040517fcd786059000000000000000000000000000000000000000000000000000000008152306004820152602401610e36565b6000808573ffffffffffffffffffffffffffffffffffffffff168486604051610f91919061131e565b60006040518083038185875af1925050503d8060008114610fce576040519150601f19603f3d011682016040523d82523d6000602084013e610fd3565b606091505b5091509150610fe3868383610fed565b9695505050505050565b60608261100257610ffd8261107c565b610f23565b8151158015611026575073ffffffffffffffffffffffffffffffffffffffff84163b155b15611075576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610e36565b5080610f23565b80511561108c5780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000602082840312156110d057600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114610f2357600080fd5b803573ffffffffffffffffffffffffffffffffffffffff8116811461112457600080fd5b919050565b60008060008060006080868803121561114157600080fd5b61114a86611100565b945060208601359350604086013567ffffffffffffffff81111561116d57600080fd5b8601601f8101881361117e57600080fd5b803567ffffffffffffffff81111561119557600080fd5b8860208284010111156111a757600080fd5b959894975060200195606001359392505050565b6000806000606084860312156111d057600080fd5b6111d984611100565b95602085013595506040909401359392505050565b60006020828403121561120057600080fd5b5035919050565b6000806040838503121561121a57600080fd5b8235915061122a60208401611100565b90509250929050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff8616815273ffffffffffffffffffffffffffffffffffffffff851660208201528360408201526080606082015260006112ce608083018486611233565b979650505050505050565b8381526040602082015260006112f3604083018486611233565b95945050505050565b60006020828403121561130e57600080fd5b81518015158114610f2357600080fd5b6000825160005b8181101561133f5760208186018101518583015201611325565b50600092019182525091905056fea264697066735822122049fc4f7005945c94be87b64761390ef0a7a9420599ed8603d7ad523eef4e31d764736f6c634300081a0033",
}

// ZetaConnectorNativeABI is the input ABI used to generate the binding from.
// Deprecated: Use ZetaConnectorNativeMetaData.ABI instead.
var ZetaConnectorNativeABI = ZetaConnectorNativeMetaData.ABI

// ZetaConnectorNativeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZetaConnectorNativeMetaData.Bin instead.
var ZetaConnectorNativeBin = ZetaConnectorNativeMetaData.Bin

// DeployZetaConnectorNative deploys a new Ethereum contract, binding an instance of ZetaConnectorNative to it.
func DeployZetaConnectorNative(auth *bind.TransactOpts, backend bind.ContractBackend, _gateway common.Address, _zetaToken common.Address, _tssAddress common.Address, _admin common.Address) (common.Address, *types.Transaction, *ZetaConnectorNative, error) {
	parsed, err := ZetaConnectorNativeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZetaConnectorNativeBin), backend, _gateway, _zetaToken, _tssAddress, _admin)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZetaConnectorNative{ZetaConnectorNativeCaller: ZetaConnectorNativeCaller{contract: contract}, ZetaConnectorNativeTransactor: ZetaConnectorNativeTransactor{contract: contract}, ZetaConnectorNativeFilterer: ZetaConnectorNativeFilterer{contract: contract}}, nil
}

// ZetaConnectorNative is an auto generated Go binding around an Ethereum contract.
type ZetaConnectorNative struct {
	ZetaConnectorNativeCaller     // Read-only binding to the contract
	ZetaConnectorNativeTransactor // Write-only binding to the contract
	ZetaConnectorNativeFilterer   // Log filterer for contract events
}

// ZetaConnectorNativeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZetaConnectorNativeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorNativeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZetaConnectorNativeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorNativeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZetaConnectorNativeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorNativeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZetaConnectorNativeSession struct {
	Contract     *ZetaConnectorNative // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ZetaConnectorNativeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZetaConnectorNativeCallerSession struct {
	Contract *ZetaConnectorNativeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// ZetaConnectorNativeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZetaConnectorNativeTransactorSession struct {
	Contract     *ZetaConnectorNativeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ZetaConnectorNativeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZetaConnectorNativeRaw struct {
	Contract *ZetaConnectorNative // Generic contract binding to access the raw methods on
}

// ZetaConnectorNativeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZetaConnectorNativeCallerRaw struct {
	Contract *ZetaConnectorNativeCaller // Generic read-only contract binding to access the raw methods on
}

// ZetaConnectorNativeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZetaConnectorNativeTransactorRaw struct {
	Contract *ZetaConnectorNativeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZetaConnectorNative creates a new instance of ZetaConnectorNative, bound to a specific deployed contract.
func NewZetaConnectorNative(address common.Address, backend bind.ContractBackend) (*ZetaConnectorNative, error) {
	contract, err := bindZetaConnectorNative(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNative{ZetaConnectorNativeCaller: ZetaConnectorNativeCaller{contract: contract}, ZetaConnectorNativeTransactor: ZetaConnectorNativeTransactor{contract: contract}, ZetaConnectorNativeFilterer: ZetaConnectorNativeFilterer{contract: contract}}, nil
}

// NewZetaConnectorNativeCaller creates a new read-only instance of ZetaConnectorNative, bound to a specific deployed contract.
func NewZetaConnectorNativeCaller(address common.Address, caller bind.ContractCaller) (*ZetaConnectorNativeCaller, error) {
	contract, err := bindZetaConnectorNative(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeCaller{contract: contract}, nil
}

// NewZetaConnectorNativeTransactor creates a new write-only instance of ZetaConnectorNative, bound to a specific deployed contract.
func NewZetaConnectorNativeTransactor(address common.Address, transactor bind.ContractTransactor) (*ZetaConnectorNativeTransactor, error) {
	contract, err := bindZetaConnectorNative(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTransactor{contract: contract}, nil
}

// NewZetaConnectorNativeFilterer creates a new log filterer instance of ZetaConnectorNative, bound to a specific deployed contract.
func NewZetaConnectorNativeFilterer(address common.Address, filterer bind.ContractFilterer) (*ZetaConnectorNativeFilterer, error) {
	contract, err := bindZetaConnectorNative(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeFilterer{contract: contract}, nil
}

// bindZetaConnectorNative binds a generic wrapper to an already deployed contract.
func bindZetaConnectorNative(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZetaConnectorNativeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaConnectorNative *ZetaConnectorNativeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaConnectorNative.Contract.ZetaConnectorNativeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaConnectorNative *ZetaConnectorNativeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNative.Contract.ZetaConnectorNativeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaConnectorNative *ZetaConnectorNativeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaConnectorNative.Contract.ZetaConnectorNativeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaConnectorNative *ZetaConnectorNativeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaConnectorNative.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaConnectorNative *ZetaConnectorNativeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNative.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaConnectorNative *ZetaConnectorNativeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaConnectorNative.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ZetaConnectorNative *ZetaConnectorNativeCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZetaConnectorNative.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ZetaConnectorNative *ZetaConnectorNativeSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ZetaConnectorNative.Contract.DEFAULTADMINROLE(&_ZetaConnectorNative.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ZetaConnectorNative *ZetaConnectorNativeCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ZetaConnectorNative.Contract.DEFAULTADMINROLE(&_ZetaConnectorNative.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ZetaConnectorNative *ZetaConnectorNativeCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZetaConnectorNative.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ZetaConnectorNative *ZetaConnectorNativeSession) PAUSERROLE() ([32]byte, error) {
	return _ZetaConnectorNative.Contract.PAUSERROLE(&_ZetaConnectorNative.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ZetaConnectorNative *ZetaConnectorNativeCallerSession) PAUSERROLE() ([32]byte, error) {
	return _ZetaConnectorNative.Contract.PAUSERROLE(&_ZetaConnectorNative.CallOpts)
}

// TSSROLE is a free data retrieval call binding the contract method 0xa783c789.
//
// Solidity: function TSS_ROLE() view returns(bytes32)
func (_ZetaConnectorNative *ZetaConnectorNativeCaller) TSSROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZetaConnectorNative.contract.Call(opts, &out, "TSS_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TSSROLE is a free data retrieval call binding the contract method 0xa783c789.
//
// Solidity: function TSS_ROLE() view returns(bytes32)
func (_ZetaConnectorNative *ZetaConnectorNativeSession) TSSROLE() ([32]byte, error) {
	return _ZetaConnectorNative.Contract.TSSROLE(&_ZetaConnectorNative.CallOpts)
}

// TSSROLE is a free data retrieval call binding the contract method 0xa783c789.
//
// Solidity: function TSS_ROLE() view returns(bytes32)
func (_ZetaConnectorNative *ZetaConnectorNativeCallerSession) TSSROLE() ([32]byte, error) {
	return _ZetaConnectorNative.Contract.TSSROLE(&_ZetaConnectorNative.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_ZetaConnectorNative *ZetaConnectorNativeCaller) Gateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorNative.contract.Call(opts, &out, "gateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_ZetaConnectorNative *ZetaConnectorNativeSession) Gateway() (common.Address, error) {
	return _ZetaConnectorNative.Contract.Gateway(&_ZetaConnectorNative.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_ZetaConnectorNative *ZetaConnectorNativeCallerSession) Gateway() (common.Address, error) {
	return _ZetaConnectorNative.Contract.Gateway(&_ZetaConnectorNative.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ZetaConnectorNative *ZetaConnectorNativeCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ZetaConnectorNative.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ZetaConnectorNative *ZetaConnectorNativeSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ZetaConnectorNative.Contract.GetRoleAdmin(&_ZetaConnectorNative.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ZetaConnectorNative *ZetaConnectorNativeCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ZetaConnectorNative.Contract.GetRoleAdmin(&_ZetaConnectorNative.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ZetaConnectorNative *ZetaConnectorNativeCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ZetaConnectorNative.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ZetaConnectorNative *ZetaConnectorNativeSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ZetaConnectorNative.Contract.HasRole(&_ZetaConnectorNative.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ZetaConnectorNative *ZetaConnectorNativeCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ZetaConnectorNative.Contract.HasRole(&_ZetaConnectorNative.CallOpts, role, account)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ZetaConnectorNative *ZetaConnectorNativeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ZetaConnectorNative.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ZetaConnectorNative *ZetaConnectorNativeSession) Paused() (bool, error) {
	return _ZetaConnectorNative.Contract.Paused(&_ZetaConnectorNative.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ZetaConnectorNative *ZetaConnectorNativeCallerSession) Paused() (bool, error) {
	return _ZetaConnectorNative.Contract.Paused(&_ZetaConnectorNative.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ZetaConnectorNative *ZetaConnectorNativeCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ZetaConnectorNative.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ZetaConnectorNative *ZetaConnectorNativeSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ZetaConnectorNative.Contract.SupportsInterface(&_ZetaConnectorNative.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ZetaConnectorNative *ZetaConnectorNativeCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ZetaConnectorNative.Contract.SupportsInterface(&_ZetaConnectorNative.CallOpts, interfaceId)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_ZetaConnectorNative *ZetaConnectorNativeCaller) TssAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorNative.contract.Call(opts, &out, "tssAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_ZetaConnectorNative *ZetaConnectorNativeSession) TssAddress() (common.Address, error) {
	return _ZetaConnectorNative.Contract.TssAddress(&_ZetaConnectorNative.CallOpts)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_ZetaConnectorNative *ZetaConnectorNativeCallerSession) TssAddress() (common.Address, error) {
	return _ZetaConnectorNative.Contract.TssAddress(&_ZetaConnectorNative.CallOpts)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaConnectorNative *ZetaConnectorNativeCaller) ZetaToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorNative.contract.Call(opts, &out, "zetaToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaConnectorNative *ZetaConnectorNativeSession) ZetaToken() (common.Address, error) {
	return _ZetaConnectorNative.Contract.ZetaToken(&_ZetaConnectorNative.CallOpts)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaConnectorNative *ZetaConnectorNativeCallerSession) ZetaToken() (common.Address, error) {
	return _ZetaConnectorNative.Contract.ZetaToken(&_ZetaConnectorNative.CallOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ZetaConnectorNative *ZetaConnectorNativeTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNative.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ZetaConnectorNative *ZetaConnectorNativeSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNative.Contract.GrantRole(&_ZetaConnectorNative.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ZetaConnectorNative *ZetaConnectorNativeTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNative.Contract.GrantRole(&_ZetaConnectorNative.TransactOpts, role, account)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ZetaConnectorNative *ZetaConnectorNativeTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNative.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ZetaConnectorNative *ZetaConnectorNativeSession) Pause() (*types.Transaction, error) {
	return _ZetaConnectorNative.Contract.Pause(&_ZetaConnectorNative.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ZetaConnectorNative *ZetaConnectorNativeTransactorSession) Pause() (*types.Transaction, error) {
	return _ZetaConnectorNative.Contract.Pause(&_ZetaConnectorNative.TransactOpts)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x743e0c9b.
//
// Solidity: function receiveTokens(uint256 amount) returns()
func (_ZetaConnectorNative *ZetaConnectorNativeTransactor) ReceiveTokens(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ZetaConnectorNative.contract.Transact(opts, "receiveTokens", amount)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x743e0c9b.
//
// Solidity: function receiveTokens(uint256 amount) returns()
func (_ZetaConnectorNative *ZetaConnectorNativeSession) ReceiveTokens(amount *big.Int) (*types.Transaction, error) {
	return _ZetaConnectorNative.Contract.ReceiveTokens(&_ZetaConnectorNative.TransactOpts, amount)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x743e0c9b.
//
// Solidity: function receiveTokens(uint256 amount) returns()
func (_ZetaConnectorNative *ZetaConnectorNativeTransactorSession) ReceiveTokens(amount *big.Int) (*types.Transaction, error) {
	return _ZetaConnectorNative.Contract.ReceiveTokens(&_ZetaConnectorNative.TransactOpts, amount)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ZetaConnectorNative *ZetaConnectorNativeTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNative.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ZetaConnectorNative *ZetaConnectorNativeSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNative.Contract.RenounceRole(&_ZetaConnectorNative.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ZetaConnectorNative *ZetaConnectorNativeTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNative.Contract.RenounceRole(&_ZetaConnectorNative.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ZetaConnectorNative *ZetaConnectorNativeTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNative.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ZetaConnectorNative *ZetaConnectorNativeSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNative.Contract.RevokeRole(&_ZetaConnectorNative.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ZetaConnectorNative *ZetaConnectorNativeTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNative.Contract.RevokeRole(&_ZetaConnectorNative.TransactOpts, role, account)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ZetaConnectorNative *ZetaConnectorNativeTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNative.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ZetaConnectorNative *ZetaConnectorNativeSession) Unpause() (*types.Transaction, error) {
	return _ZetaConnectorNative.Contract.Unpause(&_ZetaConnectorNative.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ZetaConnectorNative *ZetaConnectorNativeTransactorSession) Unpause() (*types.Transaction, error) {
	return _ZetaConnectorNative.Contract.Unpause(&_ZetaConnectorNative.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x106e6290.
//
// Solidity: function withdraw(address to, uint256 amount, bytes32 internalSendHash) returns()
func (_ZetaConnectorNative *ZetaConnectorNativeTransactor) Withdraw(opts *bind.TransactOpts, to common.Address, amount *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNative.contract.Transact(opts, "withdraw", to, amount, internalSendHash)
}

// Withdraw is a paid mutator transaction binding the contract method 0x106e6290.
//
// Solidity: function withdraw(address to, uint256 amount, bytes32 internalSendHash) returns()
func (_ZetaConnectorNative *ZetaConnectorNativeSession) Withdraw(to common.Address, amount *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNative.Contract.Withdraw(&_ZetaConnectorNative.TransactOpts, to, amount, internalSendHash)
}

// Withdraw is a paid mutator transaction binding the contract method 0x106e6290.
//
// Solidity: function withdraw(address to, uint256 amount, bytes32 internalSendHash) returns()
func (_ZetaConnectorNative *ZetaConnectorNativeTransactorSession) Withdraw(to common.Address, amount *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNative.Contract.Withdraw(&_ZetaConnectorNative.TransactOpts, to, amount, internalSendHash)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x5e3e9fef.
//
// Solidity: function withdrawAndCall(address to, uint256 amount, bytes data, bytes32 internalSendHash) returns()
func (_ZetaConnectorNative *ZetaConnectorNativeTransactor) WithdrawAndCall(opts *bind.TransactOpts, to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNative.contract.Transact(opts, "withdrawAndCall", to, amount, data, internalSendHash)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x5e3e9fef.
//
// Solidity: function withdrawAndCall(address to, uint256 amount, bytes data, bytes32 internalSendHash) returns()
func (_ZetaConnectorNative *ZetaConnectorNativeSession) WithdrawAndCall(to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNative.Contract.WithdrawAndCall(&_ZetaConnectorNative.TransactOpts, to, amount, data, internalSendHash)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x5e3e9fef.
//
// Solidity: function withdrawAndCall(address to, uint256 amount, bytes data, bytes32 internalSendHash) returns()
func (_ZetaConnectorNative *ZetaConnectorNativeTransactorSession) WithdrawAndCall(to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNative.Contract.WithdrawAndCall(&_ZetaConnectorNative.TransactOpts, to, amount, data, internalSendHash)
}

// WithdrawAndRevert is a paid mutator transaction binding the contract method 0x02d5c899.
//
// Solidity: function withdrawAndRevert(address to, uint256 amount, bytes data, bytes32 internalSendHash) returns()
func (_ZetaConnectorNative *ZetaConnectorNativeTransactor) WithdrawAndRevert(opts *bind.TransactOpts, to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNative.contract.Transact(opts, "withdrawAndRevert", to, amount, data, internalSendHash)
}

// WithdrawAndRevert is a paid mutator transaction binding the contract method 0x02d5c899.
//
// Solidity: function withdrawAndRevert(address to, uint256 amount, bytes data, bytes32 internalSendHash) returns()
func (_ZetaConnectorNative *ZetaConnectorNativeSession) WithdrawAndRevert(to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNative.Contract.WithdrawAndRevert(&_ZetaConnectorNative.TransactOpts, to, amount, data, internalSendHash)
}

// WithdrawAndRevert is a paid mutator transaction binding the contract method 0x02d5c899.
//
// Solidity: function withdrawAndRevert(address to, uint256 amount, bytes data, bytes32 internalSendHash) returns()
func (_ZetaConnectorNative *ZetaConnectorNativeTransactorSession) WithdrawAndRevert(to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNative.Contract.WithdrawAndRevert(&_ZetaConnectorNative.TransactOpts, to, amount, data, internalSendHash)
}

// ZetaConnectorNativePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ZetaConnectorNative contract.
type ZetaConnectorNativePausedIterator struct {
	Event *ZetaConnectorNativePaused // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativePaused)
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
		it.Event = new(ZetaConnectorNativePaused)
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
func (it *ZetaConnectorNativePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativePaused represents a Paused event raised by the ZetaConnectorNative contract.
type ZetaConnectorNativePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ZetaConnectorNative *ZetaConnectorNativeFilterer) FilterPaused(opts *bind.FilterOpts) (*ZetaConnectorNativePausedIterator, error) {

	logs, sub, err := _ZetaConnectorNative.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativePausedIterator{contract: _ZetaConnectorNative.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ZetaConnectorNative *ZetaConnectorNativeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativePaused) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNative.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativePaused)
				if err := _ZetaConnectorNative.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_ZetaConnectorNative *ZetaConnectorNativeFilterer) ParsePaused(log types.Log) (*ZetaConnectorNativePaused, error) {
	event := new(ZetaConnectorNativePaused)
	if err := _ZetaConnectorNative.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ZetaConnectorNative contract.
type ZetaConnectorNativeRoleAdminChangedIterator struct {
	Event *ZetaConnectorNativeRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeRoleAdminChanged)
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
		it.Event = new(ZetaConnectorNativeRoleAdminChanged)
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
func (it *ZetaConnectorNativeRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeRoleAdminChanged represents a RoleAdminChanged event raised by the ZetaConnectorNative contract.
type ZetaConnectorNativeRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ZetaConnectorNative *ZetaConnectorNativeFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ZetaConnectorNativeRoleAdminChangedIterator, error) {

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

	logs, sub, err := _ZetaConnectorNative.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeRoleAdminChangedIterator{contract: _ZetaConnectorNative.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ZetaConnectorNative *ZetaConnectorNativeFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _ZetaConnectorNative.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeRoleAdminChanged)
				if err := _ZetaConnectorNative.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_ZetaConnectorNative *ZetaConnectorNativeFilterer) ParseRoleAdminChanged(log types.Log) (*ZetaConnectorNativeRoleAdminChanged, error) {
	event := new(ZetaConnectorNativeRoleAdminChanged)
	if err := _ZetaConnectorNative.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ZetaConnectorNative contract.
type ZetaConnectorNativeRoleGrantedIterator struct {
	Event *ZetaConnectorNativeRoleGranted // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeRoleGranted)
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
		it.Event = new(ZetaConnectorNativeRoleGranted)
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
func (it *ZetaConnectorNativeRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeRoleGranted represents a RoleGranted event raised by the ZetaConnectorNative contract.
type ZetaConnectorNativeRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZetaConnectorNative *ZetaConnectorNativeFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ZetaConnectorNativeRoleGrantedIterator, error) {

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

	logs, sub, err := _ZetaConnectorNative.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeRoleGrantedIterator{contract: _ZetaConnectorNative.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZetaConnectorNative *ZetaConnectorNativeFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ZetaConnectorNative.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeRoleGranted)
				if err := _ZetaConnectorNative.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_ZetaConnectorNative *ZetaConnectorNativeFilterer) ParseRoleGranted(log types.Log) (*ZetaConnectorNativeRoleGranted, error) {
	event := new(ZetaConnectorNativeRoleGranted)
	if err := _ZetaConnectorNative.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ZetaConnectorNative contract.
type ZetaConnectorNativeRoleRevokedIterator struct {
	Event *ZetaConnectorNativeRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeRoleRevoked)
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
		it.Event = new(ZetaConnectorNativeRoleRevoked)
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
func (it *ZetaConnectorNativeRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeRoleRevoked represents a RoleRevoked event raised by the ZetaConnectorNative contract.
type ZetaConnectorNativeRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZetaConnectorNative *ZetaConnectorNativeFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ZetaConnectorNativeRoleRevokedIterator, error) {

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

	logs, sub, err := _ZetaConnectorNative.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeRoleRevokedIterator{contract: _ZetaConnectorNative.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZetaConnectorNative *ZetaConnectorNativeFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ZetaConnectorNative.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeRoleRevoked)
				if err := _ZetaConnectorNative.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_ZetaConnectorNative *ZetaConnectorNativeFilterer) ParseRoleRevoked(log types.Log) (*ZetaConnectorNativeRoleRevoked, error) {
	event := new(ZetaConnectorNativeRoleRevoked)
	if err := _ZetaConnectorNative.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ZetaConnectorNative contract.
type ZetaConnectorNativeUnpausedIterator struct {
	Event *ZetaConnectorNativeUnpaused // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeUnpaused)
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
		it.Event = new(ZetaConnectorNativeUnpaused)
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
func (it *ZetaConnectorNativeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeUnpaused represents a Unpaused event raised by the ZetaConnectorNative contract.
type ZetaConnectorNativeUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ZetaConnectorNative *ZetaConnectorNativeFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ZetaConnectorNativeUnpausedIterator, error) {

	logs, sub, err := _ZetaConnectorNative.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeUnpausedIterator{contract: _ZetaConnectorNative.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ZetaConnectorNative *ZetaConnectorNativeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeUnpaused) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNative.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeUnpaused)
				if err := _ZetaConnectorNative.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_ZetaConnectorNative *ZetaConnectorNativeFilterer) ParseUnpaused(log types.Log) (*ZetaConnectorNativeUnpaused, error) {
	event := new(ZetaConnectorNativeUnpaused)
	if err := _ZetaConnectorNative.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the ZetaConnectorNative contract.
type ZetaConnectorNativeWithdrawIterator struct {
	Event *ZetaConnectorNativeWithdraw // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeWithdraw)
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
		it.Event = new(ZetaConnectorNativeWithdraw)
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
func (it *ZetaConnectorNativeWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeWithdraw represents a Withdraw event raised by the ZetaConnectorNative contract.
type ZetaConnectorNativeWithdraw struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed to, uint256 amount)
func (_ZetaConnectorNative *ZetaConnectorNativeFilterer) FilterWithdraw(opts *bind.FilterOpts, to []common.Address) (*ZetaConnectorNativeWithdrawIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNative.contract.FilterLogs(opts, "Withdraw", toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeWithdrawIterator{contract: _ZetaConnectorNative.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed to, uint256 amount)
func (_ZetaConnectorNative *ZetaConnectorNativeFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeWithdraw, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNative.contract.WatchLogs(opts, "Withdraw", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeWithdraw)
				if err := _ZetaConnectorNative.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_ZetaConnectorNative *ZetaConnectorNativeFilterer) ParseWithdraw(log types.Log) (*ZetaConnectorNativeWithdraw, error) {
	event := new(ZetaConnectorNativeWithdraw)
	if err := _ZetaConnectorNative.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeWithdrawAndCallIterator is returned from FilterWithdrawAndCall and is used to iterate over the raw logs and unpacked data for WithdrawAndCall events raised by the ZetaConnectorNative contract.
type ZetaConnectorNativeWithdrawAndCallIterator struct {
	Event *ZetaConnectorNativeWithdrawAndCall // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeWithdrawAndCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeWithdrawAndCall)
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
		it.Event = new(ZetaConnectorNativeWithdrawAndCall)
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
func (it *ZetaConnectorNativeWithdrawAndCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeWithdrawAndCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeWithdrawAndCall represents a WithdrawAndCall event raised by the ZetaConnectorNative contract.
type ZetaConnectorNativeWithdrawAndCall struct {
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawAndCall is a free log retrieval operation binding the contract event 0x7772f56296d3a5202974a45c61c9188d844ab4d6eeb18c851e4b8d5384ca6ced.
//
// Solidity: event WithdrawAndCall(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNative *ZetaConnectorNativeFilterer) FilterWithdrawAndCall(opts *bind.FilterOpts, to []common.Address) (*ZetaConnectorNativeWithdrawAndCallIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNative.contract.FilterLogs(opts, "WithdrawAndCall", toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeWithdrawAndCallIterator{contract: _ZetaConnectorNative.contract, event: "WithdrawAndCall", logs: logs, sub: sub}, nil
}

// WatchWithdrawAndCall is a free log subscription operation binding the contract event 0x7772f56296d3a5202974a45c61c9188d844ab4d6eeb18c851e4b8d5384ca6ced.
//
// Solidity: event WithdrawAndCall(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNative *ZetaConnectorNativeFilterer) WatchWithdrawAndCall(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeWithdrawAndCall, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNative.contract.WatchLogs(opts, "WithdrawAndCall", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeWithdrawAndCall)
				if err := _ZetaConnectorNative.contract.UnpackLog(event, "WithdrawAndCall", log); err != nil {
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
func (_ZetaConnectorNative *ZetaConnectorNativeFilterer) ParseWithdrawAndCall(log types.Log) (*ZetaConnectorNativeWithdrawAndCall, error) {
	event := new(ZetaConnectorNativeWithdrawAndCall)
	if err := _ZetaConnectorNative.contract.UnpackLog(event, "WithdrawAndCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeWithdrawAndRevertIterator is returned from FilterWithdrawAndRevert and is used to iterate over the raw logs and unpacked data for WithdrawAndRevert events raised by the ZetaConnectorNative contract.
type ZetaConnectorNativeWithdrawAndRevertIterator struct {
	Event *ZetaConnectorNativeWithdrawAndRevert // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeWithdrawAndRevertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeWithdrawAndRevert)
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
		it.Event = new(ZetaConnectorNativeWithdrawAndRevert)
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
func (it *ZetaConnectorNativeWithdrawAndRevertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeWithdrawAndRevertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeWithdrawAndRevert represents a WithdrawAndRevert event raised by the ZetaConnectorNative contract.
type ZetaConnectorNativeWithdrawAndRevert struct {
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawAndRevert is a free log retrieval operation binding the contract event 0xba96f26bdda53eb8c8ba39045dfb4ff39753fbc7a6edcf250a88e75e78d102fe.
//
// Solidity: event WithdrawAndRevert(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNative *ZetaConnectorNativeFilterer) FilterWithdrawAndRevert(opts *bind.FilterOpts, to []common.Address) (*ZetaConnectorNativeWithdrawAndRevertIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNative.contract.FilterLogs(opts, "WithdrawAndRevert", toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeWithdrawAndRevertIterator{contract: _ZetaConnectorNative.contract, event: "WithdrawAndRevert", logs: logs, sub: sub}, nil
}

// WatchWithdrawAndRevert is a free log subscription operation binding the contract event 0xba96f26bdda53eb8c8ba39045dfb4ff39753fbc7a6edcf250a88e75e78d102fe.
//
// Solidity: event WithdrawAndRevert(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNative *ZetaConnectorNativeFilterer) WatchWithdrawAndRevert(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeWithdrawAndRevert, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNative.contract.WatchLogs(opts, "WithdrawAndRevert", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeWithdrawAndRevert)
				if err := _ZetaConnectorNative.contract.UnpackLog(event, "WithdrawAndRevert", log); err != nil {
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
func (_ZetaConnectorNative *ZetaConnectorNativeFilterer) ParseWithdrawAndRevert(log types.Log) (*ZetaConnectorNativeWithdrawAndRevert, error) {
	event := new(ZetaConnectorNativeWithdrawAndRevert)
	if err := _ZetaConnectorNative.contract.UnpackLog(event, "WithdrawAndRevert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
