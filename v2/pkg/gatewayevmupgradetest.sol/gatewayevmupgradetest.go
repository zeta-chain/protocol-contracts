// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gatewayevmupgradetest

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

// GatewayEVMUpgradeTestMetaData contains all meta data concerning the GatewayEVMUpgradeTest contract.
var GatewayEVMUpgradeTestMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"call\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"custody\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositAndCall\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"depositAndCall\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"destination\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"executeRevert\",\"inputs\":[{\"name\":\"destination\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"executeWithERC20\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_tssAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_zetaToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revertWithERC20\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setConnector\",\"inputs\":[{\"name\":\"_zetaConnector\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCustody\",\"inputs\":[{\"name\":\"_custody\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"tssAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"zetaConnector\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"zetaToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Call\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Executed\",\"inputs\":[{\"name\":\"destination\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecutedV2\",\"inputs\":[{\"name\":\"destination\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecutedWithERC20\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Reverted\",\"inputs\":[{\"name\":\"destination\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RevertedWithERC20\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ApprovalFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CustodyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DepositFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExecutionFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientERC20Amount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientETHAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x60a060405230608052348015601357600080fd5b5060805161254461003d600039600081816114730152818161149c01526118e201526125446000f3fe6080604052600436106101755760003560e01c80635b112591116100cb578063ae7a3a6f1161007f578063f2fde38b11610059578063f2fde38b14610414578063f340fa0114610434578063f45346dc1461044757600080fd5b8063ae7a3a6f146103b4578063b8969bd4146103d4578063dda79b75146103f457600080fd5b80638c6f037f116100b05780638c6f037f1461030e5780638da5cb5b1461032e578063ad3cb1cc1461036b57600080fd5b80635b112591146102d9578063715018a6146102f957600080fd5b806335c018db1161012d5780635131ab59116101075780635131ab591461027657806352d1902d1461029657806357bec62f146102b957600080fd5b806335c018db14610230578063485cc955146102435780634f1ef2861461026357600080fd5b80631cff79cd1161015e5780631cff79cd146101bc57806321e093b1146101e557806329c59b5d1461021d57600080fd5b806310188aef1461017a5780631b8b921d1461019c575b600080fd5b34801561018657600080fd5b5061019a610195366004612061565b610467565b005b3480156101a857600080fd5b5061019a6101b73660046120c5565b610524565b6101cf6101ca3660046120c5565b610576565b6040516101dc9190612186565b60405180910390f35b3480156101f157600080fd5b50600354610205906001600160a01b031681565b6040516001600160a01b0390911681526020016101dc565b61019a61022b3660046120c5565b61061b565b61019a61023e3660046120c5565b610740565b34801561024f57600080fd5b5061019a61025e366004612199565b6108d8565b61019a6102713660046121fb565b610b05565b34801561028257600080fd5b5061019a610291366004612302565b610b24565b3480156102a257600080fd5b506102ab610e2b565b6040519081526020016101dc565b3480156102c557600080fd5b50600254610205906001600160a01b031681565b3480156102e557600080fd5b50600154610205906001600160a01b031681565b34801561030557600080fd5b5061019a610e5a565b34801561031a57600080fd5b5061019a610329366004612371565b610e6e565b34801561033a57600080fd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b0316610205565b34801561037757600080fd5b506101cf6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b3480156103c057600080fd5b5061019a6103cf366004612061565b610f0b565b3480156103e057600080fd5b5061019a6103ef366004612302565b610fc8565b34801561040057600080fd5b50600054610205906001600160a01b031681565b34801561042057600080fd5b5061019a61042f366004612061565b61116c565b61019a610442366004612061565b6111c8565b34801561045357600080fd5b5061019a6104623660046123c3565b6112ed565b6002546001600160a01b0316156104aa576040517fb337f37800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0381166104ea576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b826001600160a01b0316336001600160a01b03167f2a21062ee9199c2e205622999eeb7c3da73153674f36a0acd3f74fa6af67bde38484604051610569929190612448565b60405180910390a3505050565b6001546060906001600160a01b031633146105bd576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006105ca858585611398565b9050846001600160a01b03167f373df382b9c587826f3de13f16d67f8d99f28ee947fc0924c6ef2d6d2c7e854634868660405161060993929190612464565b60405180910390a290505b9392505050565b34600003610655576040517f7671265e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001546040516000916001600160a01b03169034908381818185875af1925050503d80600081146106a2576040519150601f19603f3d011682016040523d82523d6000602084013e6106a7565b606091505b50909150508015156000036106e8576040517f79cacff100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b836001600160a01b0316336001600160a01b03167f2103daedac6c1eee9e5bfbd02064d751c9ec3c03fb9bc3e4f94ca41afa38c1a43460008787604051610732949392919061247e565b60405180910390a350505050565b6001546001600160a01b03163314610784576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080846001600160a01b03163460405160006040518083038185875af1925050503d80600081146107d2576040519150601f19603f3d011682016040523d82523d6000602084013e6107d7565b606091505b509150915081610813576040517facfdb44400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f8fcaa0b50000000000000000000000000000000000000000000000000000000081526001600160a01b03861690638fcaa0b59061085a9087908790600401612448565b600060405180830381600087803b15801561087457600080fd5b505af1158015610888573d6000803e3d6000fd5b50505050846001600160a01b03167fd5d7616b1678354a0dea9d7e57e6a090bff5babe9f8d6381fdbad16e89ba311c3486866040516108c993929190612464565b60405180910390a25050505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff166000811580156109235750825b905060008267ffffffffffffffff1660011480156109405750303b155b90508115801561094e575080155b15610985576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000016600117855583156109e65784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b6001600160a01b0387161580610a0357506001600160a01b038616155b15610a3a576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610a433361143f565b610a4b611450565b610a53611458565b600180546001600160a01b03808a167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161790925560038054928916929091169190911790558315610afc5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050565b610b0d611468565b610b1682611538565b610b208282611540565b5050565b610b2c611664565b6000546001600160a01b03163314801590610b5257506002546001600160a01b03163314155b15610b89576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b82600003610bc3576040517f951e19ed00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610bcd85856116e5565b610c03576040517f8164f84200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b0385811660048301526024820185905286169063095ea7b3906044016020604051808303816000875af1158015610c6b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c8f91906124a7565b610cc5576040517f8164f84200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000610cd2858484611398565b9050610cde86866116e5565b610d14576040517f8164f84200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526000906001600160a01b038816906370a0823190602401602060405180830381865afa158015610d74573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d9891906124c9565b90508015610daa57610daa8782611775565b856001600160a01b0316876001600160a01b03167f29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382878787604051610df193929190612464565b60405180910390a35050610e2460017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b5050505050565b6000610e356118d7565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b610e62611939565b610e6c60006119ad565b565b83600003610ea8576040517f951e19ed00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610eb3338486611a36565b846001600160a01b0316336001600160a01b03167f2103daedac6c1eee9e5bfbd02064d751c9ec3c03fb9bc3e4f94ca41afa38c1a486868686604051610efc949392919061247e565b60405180910390a35050505050565b6000546001600160a01b031615610f4e576040517fb337f37800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038116610f8e576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b610fd0611664565b6000546001600160a01b03163314801590610ff657506002546001600160a01b03163314155b1561102d576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b82600003611067576040517f951e19ed00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61107b6001600160a01b0386168585611b81565b6040517f8fcaa0b50000000000000000000000000000000000000000000000000000000081526001600160a01b03851690638fcaa0b5906110c29085908590600401612448565b600060405180830381600087803b1580156110dc57600080fd5b505af11580156110f0573d6000803e3d6000fd5b50505050836001600160a01b0316856001600160a01b03167f723fc7be2448075379e4fdf1e6bf5fead954d2668d2da05dcb44ccfec4beeda785858560405161113b93929190612464565b60405180910390a3610e2460017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b611174611939565b6001600160a01b0381166111bc576040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600060048201526024015b60405180910390fd5b6111c5816119ad565b50565b34600003611202576040517f7671265e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001546040516000916001600160a01b03169034908381818185875af1925050503d806000811461124f576040519150601f19603f3d011682016040523d82523d6000602084013e611254565b606091505b5090915050801515600003611295576040517f79cacff100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60408051348152600060208201819052606082840181905282015290516001600160a01b0384169133917f2103daedac6c1eee9e5bfbd02064d751c9ec3c03fb9bc3e4f94ca41afa38c1a49181900360800190a35050565b81600003611327576040517f951e19ed00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611332338284611a36565b826001600160a01b0316336001600160a01b03167f2103daedac6c1eee9e5bfbd02064d751c9ec3c03fb9bc3e4f94ca41afa38c1a484846040516105699291909182526001600160a01b0316602082015260606040820181905260009082015260800190565b6060600080856001600160a01b03163486866040516113b89291906124e2565b60006040518083038185875af1925050503d80600081146113f5576040519150601f19603f3d011682016040523d82523d6000602084013e6113fa565b606091505b509150915081611436576040517facfdb44400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b95945050505050565b611447611bf5565b6111c581611c5c565b610e6c611bf5565b611460611bf5565b610e6c611c64565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148061150157507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166114f57f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b031614155b15610e6c576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6111c5611939565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156115b8575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682019092526115b5918101906124c9565b60015b6115f9576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b03831660048201526024016111b3565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8114611655576040517faa1d49a4000000000000000000000000000000000000000000000000000000008152600481018290526024016111b3565b61165f8383611c6c565b505050565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0080547ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016116df576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60029055565b6040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b038281166004830152600060248301819052919084169063095ea7b3906044016020604051808303816000875af1158015611751573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061061491906124a7565b6003546001600160a01b0390811690831603611897576002546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152602481018390529083169063095ea7b3906044016020604051808303816000875af11580156117f7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061181b91906124a7565b506002546040517f743e0c9b000000000000000000000000000000000000000000000000000000008152600481018390526001600160a01b039091169063743e0c9b90602401600060405180830381600087803b15801561187b57600080fd5b505af115801561188f573d6000803e3d6000fd5b505050505050565b600054610b20906001600160a01b03848116911683611b81565b60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610e6c576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3361196b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b031614610e6c576040517f118cdaa70000000000000000000000000000000000000000000000000000000081523360048201526024016111b3565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080547fffffffffffffffffffffffff000000000000000000000000000000000000000081166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3505050565b6003546001600160a01b0390811690831603611b6557611a616001600160a01b038316843084611cc2565b6002546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152602481018390529083169063095ea7b3906044016020604051808303816000875af1158015611acd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611af191906124a7565b506002546040517f743e0c9b000000000000000000000000000000000000000000000000000000008152600481018390526001600160a01b039091169063743e0c9b90602401600060405180830381600087803b158015611b5157600080fd5b505af1158015610afc573d6000803e3d6000fd5b60005461165f906001600160a01b038481169186911684611cc2565b6040516001600160a01b0383811660248301526044820183905261165f91859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611d01565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff16610e6c576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611174611bf5565b6118b1611bf5565b611c7582611d7d565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a2805115611cba5761165f8282611e25565b610b20611e92565b6040516001600160a01b038481166024830152838116604483015260648201839052611cfb9186918216906323b872dd90608401611bae565b50505050565b6000611d166001600160a01b03841683611eca565b90508051600014158015611d3b575080806020019051810190611d3991906124a7565b155b1561165f576040517f5274afe70000000000000000000000000000000000000000000000000000000081526001600160a01b03841660048201526024016111b3565b806001600160a01b03163b600003611dcc576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b03821660048201526024016111b3565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b6060600080846001600160a01b031684604051611e4291906124f2565b600060405180830381855af49150503d8060008114611e7d576040519150601f19603f3d011682016040523d82523d6000602084013e611e82565b606091505b5091509150611436858383611ed8565b3415610e6c576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606061061483836000611f4d565b606082611eed57611ee882612003565b610614565b8151158015611f0457506001600160a01b0384163b155b15611f46576040517f9996b3150000000000000000000000000000000000000000000000000000000081526001600160a01b03851660048201526024016111b3565b5080610614565b606081471015611f8b576040517fcd7860590000000000000000000000000000000000000000000000000000000081523060048201526024016111b3565b600080856001600160a01b03168486604051611fa791906124f2565b60006040518083038185875af1925050503d8060008114611fe4576040519150601f19603f3d011682016040523d82523d6000602084013e611fe9565b606091505b5091509150611ff9868383611ed8565b9695505050505050565b8051156120135780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80356001600160a01b038116811461205c57600080fd5b919050565b60006020828403121561207357600080fd5b61061482612045565b60008083601f84011261208e57600080fd5b50813567ffffffffffffffff8111156120a657600080fd5b6020830191508360208285010111156120be57600080fd5b9250929050565b6000806000604084860312156120da57600080fd5b6120e384612045565b9250602084013567ffffffffffffffff8111156120ff57600080fd5b61210b8682870161207c565b9497909650939450505050565b60005b8381101561213357818101518382015260200161211b565b50506000910152565b60008151808452612154816020860160208601612118565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000610614602083018461213c565b600080604083850312156121ac57600080fd5b6121b583612045565b91506121c360208401612045565b90509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000806040838503121561220e57600080fd5b61221783612045565b9150602083013567ffffffffffffffff81111561223357600080fd5b8301601f8101851361224457600080fd5b803567ffffffffffffffff81111561225e5761225e6121cc565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff821117156122ca576122ca6121cc565b6040528181528282016020018710156122e257600080fd5b816020840160208301376000602083830101528093505050509250929050565b60008060008060006080868803121561231a57600080fd5b61232386612045565b945061233160208701612045565b935060408601359250606086013567ffffffffffffffff81111561235457600080fd5b6123608882890161207c565b969995985093965092949392505050565b60008060008060006080868803121561238957600080fd5b61239286612045565b9450602086013593506123a760408701612045565b9250606086013567ffffffffffffffff81111561235457600080fd5b6000806000606084860312156123d857600080fd5b6123e184612045565b9250602084013591506123f660408501612045565b90509250925092565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b60208152600061245c6020830184866123ff565b949350505050565b8381526040602082015260006114366040830184866123ff565b8481526001600160a01b0384166020820152606060408201526000611ff96060830184866123ff565b6000602082840312156124b957600080fd5b8151801515811461061457600080fd5b6000602082840312156124db57600080fd5b5051919050565b8183823760009101908152919050565b60008251612504818460208701612118565b919091019291505056fea2646970667358221220de349acfc8741efa6db137df12b279146aa84392e947dfc852fc83fbe89954ff64736f6c634300081a0033",
}

// GatewayEVMUpgradeTestABI is the input ABI used to generate the binding from.
// Deprecated: Use GatewayEVMUpgradeTestMetaData.ABI instead.
var GatewayEVMUpgradeTestABI = GatewayEVMUpgradeTestMetaData.ABI

// GatewayEVMUpgradeTestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GatewayEVMUpgradeTestMetaData.Bin instead.
var GatewayEVMUpgradeTestBin = GatewayEVMUpgradeTestMetaData.Bin

// DeployGatewayEVMUpgradeTest deploys a new Ethereum contract, binding an instance of GatewayEVMUpgradeTest to it.
func DeployGatewayEVMUpgradeTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GatewayEVMUpgradeTest, error) {
	parsed, err := GatewayEVMUpgradeTestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GatewayEVMUpgradeTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GatewayEVMUpgradeTest{GatewayEVMUpgradeTestCaller: GatewayEVMUpgradeTestCaller{contract: contract}, GatewayEVMUpgradeTestTransactor: GatewayEVMUpgradeTestTransactor{contract: contract}, GatewayEVMUpgradeTestFilterer: GatewayEVMUpgradeTestFilterer{contract: contract}}, nil
}

// GatewayEVMUpgradeTest is an auto generated Go binding around an Ethereum contract.
type GatewayEVMUpgradeTest struct {
	GatewayEVMUpgradeTestCaller     // Read-only binding to the contract
	GatewayEVMUpgradeTestTransactor // Write-only binding to the contract
	GatewayEVMUpgradeTestFilterer   // Log filterer for contract events
}

// GatewayEVMUpgradeTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type GatewayEVMUpgradeTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayEVMUpgradeTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GatewayEVMUpgradeTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayEVMUpgradeTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GatewayEVMUpgradeTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayEVMUpgradeTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GatewayEVMUpgradeTestSession struct {
	Contract     *GatewayEVMUpgradeTest // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// GatewayEVMUpgradeTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GatewayEVMUpgradeTestCallerSession struct {
	Contract *GatewayEVMUpgradeTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// GatewayEVMUpgradeTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GatewayEVMUpgradeTestTransactorSession struct {
	Contract     *GatewayEVMUpgradeTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// GatewayEVMUpgradeTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type GatewayEVMUpgradeTestRaw struct {
	Contract *GatewayEVMUpgradeTest // Generic contract binding to access the raw methods on
}

// GatewayEVMUpgradeTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GatewayEVMUpgradeTestCallerRaw struct {
	Contract *GatewayEVMUpgradeTestCaller // Generic read-only contract binding to access the raw methods on
}

// GatewayEVMUpgradeTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GatewayEVMUpgradeTestTransactorRaw struct {
	Contract *GatewayEVMUpgradeTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGatewayEVMUpgradeTest creates a new instance of GatewayEVMUpgradeTest, bound to a specific deployed contract.
func NewGatewayEVMUpgradeTest(address common.Address, backend bind.ContractBackend) (*GatewayEVMUpgradeTest, error) {
	contract, err := bindGatewayEVMUpgradeTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMUpgradeTest{GatewayEVMUpgradeTestCaller: GatewayEVMUpgradeTestCaller{contract: contract}, GatewayEVMUpgradeTestTransactor: GatewayEVMUpgradeTestTransactor{contract: contract}, GatewayEVMUpgradeTestFilterer: GatewayEVMUpgradeTestFilterer{contract: contract}}, nil
}

// NewGatewayEVMUpgradeTestCaller creates a new read-only instance of GatewayEVMUpgradeTest, bound to a specific deployed contract.
func NewGatewayEVMUpgradeTestCaller(address common.Address, caller bind.ContractCaller) (*GatewayEVMUpgradeTestCaller, error) {
	contract, err := bindGatewayEVMUpgradeTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMUpgradeTestCaller{contract: contract}, nil
}

// NewGatewayEVMUpgradeTestTransactor creates a new write-only instance of GatewayEVMUpgradeTest, bound to a specific deployed contract.
func NewGatewayEVMUpgradeTestTransactor(address common.Address, transactor bind.ContractTransactor) (*GatewayEVMUpgradeTestTransactor, error) {
	contract, err := bindGatewayEVMUpgradeTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMUpgradeTestTransactor{contract: contract}, nil
}

// NewGatewayEVMUpgradeTestFilterer creates a new log filterer instance of GatewayEVMUpgradeTest, bound to a specific deployed contract.
func NewGatewayEVMUpgradeTestFilterer(address common.Address, filterer bind.ContractFilterer) (*GatewayEVMUpgradeTestFilterer, error) {
	contract, err := bindGatewayEVMUpgradeTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMUpgradeTestFilterer{contract: contract}, nil
}

// bindGatewayEVMUpgradeTest binds a generic wrapper to an already deployed contract.
func bindGatewayEVMUpgradeTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GatewayEVMUpgradeTestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayEVMUpgradeTest.Contract.GatewayEVMUpgradeTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.GatewayEVMUpgradeTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.GatewayEVMUpgradeTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayEVMUpgradeTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GatewayEVMUpgradeTest.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _GatewayEVMUpgradeTest.Contract.UPGRADEINTERFACEVERSION(&_GatewayEVMUpgradeTest.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _GatewayEVMUpgradeTest.Contract.UPGRADEINTERFACEVERSION(&_GatewayEVMUpgradeTest.CallOpts)
}

// Custody is a free data retrieval call binding the contract method 0xdda79b75.
//
// Solidity: function custody() view returns(address)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCaller) Custody(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayEVMUpgradeTest.contract.Call(opts, &out, "custody")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Custody is a free data retrieval call binding the contract method 0xdda79b75.
//
// Solidity: function custody() view returns(address)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) Custody() (common.Address, error) {
	return _GatewayEVMUpgradeTest.Contract.Custody(&_GatewayEVMUpgradeTest.CallOpts)
}

// Custody is a free data retrieval call binding the contract method 0xdda79b75.
//
// Solidity: function custody() view returns(address)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCallerSession) Custody() (common.Address, error) {
	return _GatewayEVMUpgradeTest.Contract.Custody(&_GatewayEVMUpgradeTest.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayEVMUpgradeTest.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) Owner() (common.Address, error) {
	return _GatewayEVMUpgradeTest.Contract.Owner(&_GatewayEVMUpgradeTest.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCallerSession) Owner() (common.Address, error) {
	return _GatewayEVMUpgradeTest.Contract.Owner(&_GatewayEVMUpgradeTest.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayEVMUpgradeTest.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) ProxiableUUID() ([32]byte, error) {
	return _GatewayEVMUpgradeTest.Contract.ProxiableUUID(&_GatewayEVMUpgradeTest.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCallerSession) ProxiableUUID() ([32]byte, error) {
	return _GatewayEVMUpgradeTest.Contract.ProxiableUUID(&_GatewayEVMUpgradeTest.CallOpts)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCaller) TssAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayEVMUpgradeTest.contract.Call(opts, &out, "tssAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) TssAddress() (common.Address, error) {
	return _GatewayEVMUpgradeTest.Contract.TssAddress(&_GatewayEVMUpgradeTest.CallOpts)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCallerSession) TssAddress() (common.Address, error) {
	return _GatewayEVMUpgradeTest.Contract.TssAddress(&_GatewayEVMUpgradeTest.CallOpts)
}

// ZetaConnector is a free data retrieval call binding the contract method 0x57bec62f.
//
// Solidity: function zetaConnector() view returns(address)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCaller) ZetaConnector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayEVMUpgradeTest.contract.Call(opts, &out, "zetaConnector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZetaConnector is a free data retrieval call binding the contract method 0x57bec62f.
//
// Solidity: function zetaConnector() view returns(address)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) ZetaConnector() (common.Address, error) {
	return _GatewayEVMUpgradeTest.Contract.ZetaConnector(&_GatewayEVMUpgradeTest.CallOpts)
}

// ZetaConnector is a free data retrieval call binding the contract method 0x57bec62f.
//
// Solidity: function zetaConnector() view returns(address)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCallerSession) ZetaConnector() (common.Address, error) {
	return _GatewayEVMUpgradeTest.Contract.ZetaConnector(&_GatewayEVMUpgradeTest.CallOpts)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCaller) ZetaToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayEVMUpgradeTest.contract.Call(opts, &out, "zetaToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) ZetaToken() (common.Address, error) {
	return _GatewayEVMUpgradeTest.Contract.ZetaToken(&_GatewayEVMUpgradeTest.CallOpts)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCallerSession) ZetaToken() (common.Address, error) {
	return _GatewayEVMUpgradeTest.Contract.ZetaToken(&_GatewayEVMUpgradeTest.CallOpts)
}

// Call is a paid mutator transaction binding the contract method 0x1b8b921d.
//
// Solidity: function call(address receiver, bytes payload) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactor) Call(opts *bind.TransactOpts, receiver common.Address, payload []byte) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.contract.Transact(opts, "call", receiver, payload)
}

// Call is a paid mutator transaction binding the contract method 0x1b8b921d.
//
// Solidity: function call(address receiver, bytes payload) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) Call(receiver common.Address, payload []byte) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.Call(&_GatewayEVMUpgradeTest.TransactOpts, receiver, payload)
}

// Call is a paid mutator transaction binding the contract method 0x1b8b921d.
//
// Solidity: function call(address receiver, bytes payload) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorSession) Call(receiver common.Address, payload []byte) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.Call(&_GatewayEVMUpgradeTest.TransactOpts, receiver, payload)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address receiver) payable returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactor) Deposit(opts *bind.TransactOpts, receiver common.Address) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.contract.Transact(opts, "deposit", receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address receiver) payable returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) Deposit(receiver common.Address) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.Deposit(&_GatewayEVMUpgradeTest.TransactOpts, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address receiver) payable returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorSession) Deposit(receiver common.Address) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.Deposit(&_GatewayEVMUpgradeTest.TransactOpts, receiver)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xf45346dc.
//
// Solidity: function deposit(address receiver, uint256 amount, address asset) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactor) Deposit0(opts *bind.TransactOpts, receiver common.Address, amount *big.Int, asset common.Address) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.contract.Transact(opts, "deposit0", receiver, amount, asset)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xf45346dc.
//
// Solidity: function deposit(address receiver, uint256 amount, address asset) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) Deposit0(receiver common.Address, amount *big.Int, asset common.Address) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.Deposit0(&_GatewayEVMUpgradeTest.TransactOpts, receiver, amount, asset)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xf45346dc.
//
// Solidity: function deposit(address receiver, uint256 amount, address asset) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorSession) Deposit0(receiver common.Address, amount *big.Int, asset common.Address) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.Deposit0(&_GatewayEVMUpgradeTest.TransactOpts, receiver, amount, asset)
}

// DepositAndCall is a paid mutator transaction binding the contract method 0x29c59b5d.
//
// Solidity: function depositAndCall(address receiver, bytes payload) payable returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactor) DepositAndCall(opts *bind.TransactOpts, receiver common.Address, payload []byte) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.contract.Transact(opts, "depositAndCall", receiver, payload)
}

// DepositAndCall is a paid mutator transaction binding the contract method 0x29c59b5d.
//
// Solidity: function depositAndCall(address receiver, bytes payload) payable returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) DepositAndCall(receiver common.Address, payload []byte) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.DepositAndCall(&_GatewayEVMUpgradeTest.TransactOpts, receiver, payload)
}

// DepositAndCall is a paid mutator transaction binding the contract method 0x29c59b5d.
//
// Solidity: function depositAndCall(address receiver, bytes payload) payable returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorSession) DepositAndCall(receiver common.Address, payload []byte) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.DepositAndCall(&_GatewayEVMUpgradeTest.TransactOpts, receiver, payload)
}

// DepositAndCall0 is a paid mutator transaction binding the contract method 0x8c6f037f.
//
// Solidity: function depositAndCall(address receiver, uint256 amount, address asset, bytes payload) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactor) DepositAndCall0(opts *bind.TransactOpts, receiver common.Address, amount *big.Int, asset common.Address, payload []byte) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.contract.Transact(opts, "depositAndCall0", receiver, amount, asset, payload)
}

// DepositAndCall0 is a paid mutator transaction binding the contract method 0x8c6f037f.
//
// Solidity: function depositAndCall(address receiver, uint256 amount, address asset, bytes payload) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) DepositAndCall0(receiver common.Address, amount *big.Int, asset common.Address, payload []byte) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.DepositAndCall0(&_GatewayEVMUpgradeTest.TransactOpts, receiver, amount, asset, payload)
}

// DepositAndCall0 is a paid mutator transaction binding the contract method 0x8c6f037f.
//
// Solidity: function depositAndCall(address receiver, uint256 amount, address asset, bytes payload) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorSession) DepositAndCall0(receiver common.Address, amount *big.Int, asset common.Address, payload []byte) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.DepositAndCall0(&_GatewayEVMUpgradeTest.TransactOpts, receiver, amount, asset, payload)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address destination, bytes data) payable returns(bytes)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactor) Execute(opts *bind.TransactOpts, destination common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.contract.Transact(opts, "execute", destination, data)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address destination, bytes data) payable returns(bytes)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) Execute(destination common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.Execute(&_GatewayEVMUpgradeTest.TransactOpts, destination, data)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address destination, bytes data) payable returns(bytes)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorSession) Execute(destination common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.Execute(&_GatewayEVMUpgradeTest.TransactOpts, destination, data)
}

// ExecuteRevert is a paid mutator transaction binding the contract method 0x35c018db.
//
// Solidity: function executeRevert(address destination, bytes data) payable returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactor) ExecuteRevert(opts *bind.TransactOpts, destination common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.contract.Transact(opts, "executeRevert", destination, data)
}

// ExecuteRevert is a paid mutator transaction binding the contract method 0x35c018db.
//
// Solidity: function executeRevert(address destination, bytes data) payable returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) ExecuteRevert(destination common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.ExecuteRevert(&_GatewayEVMUpgradeTest.TransactOpts, destination, data)
}

// ExecuteRevert is a paid mutator transaction binding the contract method 0x35c018db.
//
// Solidity: function executeRevert(address destination, bytes data) payable returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorSession) ExecuteRevert(destination common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.ExecuteRevert(&_GatewayEVMUpgradeTest.TransactOpts, destination, data)
}

// ExecuteWithERC20 is a paid mutator transaction binding the contract method 0x5131ab59.
//
// Solidity: function executeWithERC20(address token, address to, uint256 amount, bytes data) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactor) ExecuteWithERC20(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.contract.Transact(opts, "executeWithERC20", token, to, amount, data)
}

// ExecuteWithERC20 is a paid mutator transaction binding the contract method 0x5131ab59.
//
// Solidity: function executeWithERC20(address token, address to, uint256 amount, bytes data) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) ExecuteWithERC20(token common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.ExecuteWithERC20(&_GatewayEVMUpgradeTest.TransactOpts, token, to, amount, data)
}

// ExecuteWithERC20 is a paid mutator transaction binding the contract method 0x5131ab59.
//
// Solidity: function executeWithERC20(address token, address to, uint256 amount, bytes data) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorSession) ExecuteWithERC20(token common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.ExecuteWithERC20(&_GatewayEVMUpgradeTest.TransactOpts, token, to, amount, data)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _tssAddress, address _zetaToken) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactor) Initialize(opts *bind.TransactOpts, _tssAddress common.Address, _zetaToken common.Address) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.contract.Transact(opts, "initialize", _tssAddress, _zetaToken)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _tssAddress, address _zetaToken) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) Initialize(_tssAddress common.Address, _zetaToken common.Address) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.Initialize(&_GatewayEVMUpgradeTest.TransactOpts, _tssAddress, _zetaToken)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _tssAddress, address _zetaToken) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorSession) Initialize(_tssAddress common.Address, _zetaToken common.Address) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.Initialize(&_GatewayEVMUpgradeTest.TransactOpts, _tssAddress, _zetaToken)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) RenounceOwnership() (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.RenounceOwnership(&_GatewayEVMUpgradeTest.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.RenounceOwnership(&_GatewayEVMUpgradeTest.TransactOpts)
}

// RevertWithERC20 is a paid mutator transaction binding the contract method 0xb8969bd4.
//
// Solidity: function revertWithERC20(address token, address to, uint256 amount, bytes data) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactor) RevertWithERC20(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.contract.Transact(opts, "revertWithERC20", token, to, amount, data)
}

// RevertWithERC20 is a paid mutator transaction binding the contract method 0xb8969bd4.
//
// Solidity: function revertWithERC20(address token, address to, uint256 amount, bytes data) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) RevertWithERC20(token common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.RevertWithERC20(&_GatewayEVMUpgradeTest.TransactOpts, token, to, amount, data)
}

// RevertWithERC20 is a paid mutator transaction binding the contract method 0xb8969bd4.
//
// Solidity: function revertWithERC20(address token, address to, uint256 amount, bytes data) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorSession) RevertWithERC20(token common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.RevertWithERC20(&_GatewayEVMUpgradeTest.TransactOpts, token, to, amount, data)
}

// SetConnector is a paid mutator transaction binding the contract method 0x10188aef.
//
// Solidity: function setConnector(address _zetaConnector) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactor) SetConnector(opts *bind.TransactOpts, _zetaConnector common.Address) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.contract.Transact(opts, "setConnector", _zetaConnector)
}

// SetConnector is a paid mutator transaction binding the contract method 0x10188aef.
//
// Solidity: function setConnector(address _zetaConnector) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) SetConnector(_zetaConnector common.Address) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.SetConnector(&_GatewayEVMUpgradeTest.TransactOpts, _zetaConnector)
}

// SetConnector is a paid mutator transaction binding the contract method 0x10188aef.
//
// Solidity: function setConnector(address _zetaConnector) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorSession) SetConnector(_zetaConnector common.Address) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.SetConnector(&_GatewayEVMUpgradeTest.TransactOpts, _zetaConnector)
}

// SetCustody is a paid mutator transaction binding the contract method 0xae7a3a6f.
//
// Solidity: function setCustody(address _custody) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactor) SetCustody(opts *bind.TransactOpts, _custody common.Address) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.contract.Transact(opts, "setCustody", _custody)
}

// SetCustody is a paid mutator transaction binding the contract method 0xae7a3a6f.
//
// Solidity: function setCustody(address _custody) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) SetCustody(_custody common.Address) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.SetCustody(&_GatewayEVMUpgradeTest.TransactOpts, _custody)
}

// SetCustody is a paid mutator transaction binding the contract method 0xae7a3a6f.
//
// Solidity: function setCustody(address _custody) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorSession) SetCustody(_custody common.Address) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.SetCustody(&_GatewayEVMUpgradeTest.TransactOpts, _custody)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.TransferOwnership(&_GatewayEVMUpgradeTest.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.TransferOwnership(&_GatewayEVMUpgradeTest.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.UpgradeToAndCall(&_GatewayEVMUpgradeTest.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.UpgradeToAndCall(&_GatewayEVMUpgradeTest.TransactOpts, newImplementation, data)
}

// GatewayEVMUpgradeTestCallIterator is returned from FilterCall and is used to iterate over the raw logs and unpacked data for Call events raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestCallIterator struct {
	Event *GatewayEVMUpgradeTestCall // Event containing the contract specifics and raw log

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
func (it *GatewayEVMUpgradeTestCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMUpgradeTestCall)
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
		it.Event = new(GatewayEVMUpgradeTestCall)
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
func (it *GatewayEVMUpgradeTestCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMUpgradeTestCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMUpgradeTestCall represents a Call event raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestCall struct {
	Sender   common.Address
	Receiver common.Address
	Payload  []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCall is a free log retrieval operation binding the contract event 0x2a21062ee9199c2e205622999eeb7c3da73153674f36a0acd3f74fa6af67bde3.
//
// Solidity: event Call(address indexed sender, address indexed receiver, bytes payload)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) FilterCall(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*GatewayEVMUpgradeTestCallIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMUpgradeTest.contract.FilterLogs(opts, "Call", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMUpgradeTestCallIterator{contract: _GatewayEVMUpgradeTest.contract, event: "Call", logs: logs, sub: sub}, nil
}

// WatchCall is a free log subscription operation binding the contract event 0x2a21062ee9199c2e205622999eeb7c3da73153674f36a0acd3f74fa6af67bde3.
//
// Solidity: event Call(address indexed sender, address indexed receiver, bytes payload)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) WatchCall(opts *bind.WatchOpts, sink chan<- *GatewayEVMUpgradeTestCall, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMUpgradeTest.contract.WatchLogs(opts, "Call", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMUpgradeTestCall)
				if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "Call", log); err != nil {
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

// ParseCall is a log parse operation binding the contract event 0x2a21062ee9199c2e205622999eeb7c3da73153674f36a0acd3f74fa6af67bde3.
//
// Solidity: event Call(address indexed sender, address indexed receiver, bytes payload)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) ParseCall(log types.Log) (*GatewayEVMUpgradeTestCall, error) {
	event := new(GatewayEVMUpgradeTestCall)
	if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "Call", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMUpgradeTestDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestDepositIterator struct {
	Event *GatewayEVMUpgradeTestDeposit // Event containing the contract specifics and raw log

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
func (it *GatewayEVMUpgradeTestDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMUpgradeTestDeposit)
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
		it.Event = new(GatewayEVMUpgradeTestDeposit)
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
func (it *GatewayEVMUpgradeTestDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMUpgradeTestDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMUpgradeTestDeposit represents a Deposit event raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestDeposit struct {
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
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*GatewayEVMUpgradeTestDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMUpgradeTest.contract.FilterLogs(opts, "Deposit", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMUpgradeTestDepositIterator{contract: _GatewayEVMUpgradeTest.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x2103daedac6c1eee9e5bfbd02064d751c9ec3c03fb9bc3e4f94ca41afa38c1a4.
//
// Solidity: event Deposit(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *GatewayEVMUpgradeTestDeposit, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMUpgradeTest.contract.WatchLogs(opts, "Deposit", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMUpgradeTestDeposit)
				if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) ParseDeposit(log types.Log) (*GatewayEVMUpgradeTestDeposit, error) {
	event := new(GatewayEVMUpgradeTestDeposit)
	if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMUpgradeTestExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestExecutedIterator struct {
	Event *GatewayEVMUpgradeTestExecuted // Event containing the contract specifics and raw log

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
func (it *GatewayEVMUpgradeTestExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMUpgradeTestExecuted)
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
		it.Event = new(GatewayEVMUpgradeTestExecuted)
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
func (it *GatewayEVMUpgradeTestExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMUpgradeTestExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMUpgradeTestExecuted represents a Executed event raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestExecuted struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0xcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f.
//
// Solidity: event Executed(address indexed destination, uint256 value, bytes data)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) FilterExecuted(opts *bind.FilterOpts, destination []common.Address) (*GatewayEVMUpgradeTestExecutedIterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _GatewayEVMUpgradeTest.contract.FilterLogs(opts, "Executed", destinationRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMUpgradeTestExecutedIterator{contract: _GatewayEVMUpgradeTest.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0xcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f.
//
// Solidity: event Executed(address indexed destination, uint256 value, bytes data)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *GatewayEVMUpgradeTestExecuted, destination []common.Address) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _GatewayEVMUpgradeTest.contract.WatchLogs(opts, "Executed", destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMUpgradeTestExecuted)
				if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "Executed", log); err != nil {
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
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) ParseExecuted(log types.Log) (*GatewayEVMUpgradeTestExecuted, error) {
	event := new(GatewayEVMUpgradeTestExecuted)
	if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMUpgradeTestExecutedV2Iterator is returned from FilterExecutedV2 and is used to iterate over the raw logs and unpacked data for ExecutedV2 events raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestExecutedV2Iterator struct {
	Event *GatewayEVMUpgradeTestExecutedV2 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMUpgradeTestExecutedV2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMUpgradeTestExecutedV2)
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
		it.Event = new(GatewayEVMUpgradeTestExecutedV2)
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
func (it *GatewayEVMUpgradeTestExecutedV2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMUpgradeTestExecutedV2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMUpgradeTestExecutedV2 represents a ExecutedV2 event raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestExecutedV2 struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExecutedV2 is a free log retrieval operation binding the contract event 0x373df382b9c587826f3de13f16d67f8d99f28ee947fc0924c6ef2d6d2c7e8546.
//
// Solidity: event ExecutedV2(address indexed destination, uint256 value, bytes data)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) FilterExecutedV2(opts *bind.FilterOpts, destination []common.Address) (*GatewayEVMUpgradeTestExecutedV2Iterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _GatewayEVMUpgradeTest.contract.FilterLogs(opts, "ExecutedV2", destinationRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMUpgradeTestExecutedV2Iterator{contract: _GatewayEVMUpgradeTest.contract, event: "ExecutedV2", logs: logs, sub: sub}, nil
}

// WatchExecutedV2 is a free log subscription operation binding the contract event 0x373df382b9c587826f3de13f16d67f8d99f28ee947fc0924c6ef2d6d2c7e8546.
//
// Solidity: event ExecutedV2(address indexed destination, uint256 value, bytes data)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) WatchExecutedV2(opts *bind.WatchOpts, sink chan<- *GatewayEVMUpgradeTestExecutedV2, destination []common.Address) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _GatewayEVMUpgradeTest.contract.WatchLogs(opts, "ExecutedV2", destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMUpgradeTestExecutedV2)
				if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "ExecutedV2", log); err != nil {
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

// ParseExecutedV2 is a log parse operation binding the contract event 0x373df382b9c587826f3de13f16d67f8d99f28ee947fc0924c6ef2d6d2c7e8546.
//
// Solidity: event ExecutedV2(address indexed destination, uint256 value, bytes data)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) ParseExecutedV2(log types.Log) (*GatewayEVMUpgradeTestExecutedV2, error) {
	event := new(GatewayEVMUpgradeTestExecutedV2)
	if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "ExecutedV2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMUpgradeTestExecutedWithERC20Iterator is returned from FilterExecutedWithERC20 and is used to iterate over the raw logs and unpacked data for ExecutedWithERC20 events raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestExecutedWithERC20Iterator struct {
	Event *GatewayEVMUpgradeTestExecutedWithERC20 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMUpgradeTestExecutedWithERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMUpgradeTestExecutedWithERC20)
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
		it.Event = new(GatewayEVMUpgradeTestExecutedWithERC20)
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
func (it *GatewayEVMUpgradeTestExecutedWithERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMUpgradeTestExecutedWithERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMUpgradeTestExecutedWithERC20 represents a ExecutedWithERC20 event raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestExecutedWithERC20 struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExecutedWithERC20 is a free log retrieval operation binding the contract event 0x29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382.
//
// Solidity: event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) FilterExecutedWithERC20(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*GatewayEVMUpgradeTestExecutedWithERC20Iterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GatewayEVMUpgradeTest.contract.FilterLogs(opts, "ExecutedWithERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMUpgradeTestExecutedWithERC20Iterator{contract: _GatewayEVMUpgradeTest.contract, event: "ExecutedWithERC20", logs: logs, sub: sub}, nil
}

// WatchExecutedWithERC20 is a free log subscription operation binding the contract event 0x29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382.
//
// Solidity: event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) WatchExecutedWithERC20(opts *bind.WatchOpts, sink chan<- *GatewayEVMUpgradeTestExecutedWithERC20, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GatewayEVMUpgradeTest.contract.WatchLogs(opts, "ExecutedWithERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMUpgradeTestExecutedWithERC20)
				if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "ExecutedWithERC20", log); err != nil {
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
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) ParseExecutedWithERC20(log types.Log) (*GatewayEVMUpgradeTestExecutedWithERC20, error) {
	event := new(GatewayEVMUpgradeTestExecutedWithERC20)
	if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "ExecutedWithERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMUpgradeTestInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestInitializedIterator struct {
	Event *GatewayEVMUpgradeTestInitialized // Event containing the contract specifics and raw log

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
func (it *GatewayEVMUpgradeTestInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMUpgradeTestInitialized)
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
		it.Event = new(GatewayEVMUpgradeTestInitialized)
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
func (it *GatewayEVMUpgradeTestInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMUpgradeTestInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMUpgradeTestInitialized represents a Initialized event raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) FilterInitialized(opts *bind.FilterOpts) (*GatewayEVMUpgradeTestInitializedIterator, error) {

	logs, sub, err := _GatewayEVMUpgradeTest.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMUpgradeTestInitializedIterator{contract: _GatewayEVMUpgradeTest.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *GatewayEVMUpgradeTestInitialized) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMUpgradeTest.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMUpgradeTestInitialized)
				if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) ParseInitialized(log types.Log) (*GatewayEVMUpgradeTestInitialized, error) {
	event := new(GatewayEVMUpgradeTestInitialized)
	if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMUpgradeTestOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestOwnershipTransferredIterator struct {
	Event *GatewayEVMUpgradeTestOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GatewayEVMUpgradeTestOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMUpgradeTestOwnershipTransferred)
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
		it.Event = new(GatewayEVMUpgradeTestOwnershipTransferred)
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
func (it *GatewayEVMUpgradeTestOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMUpgradeTestOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMUpgradeTestOwnershipTransferred represents a OwnershipTransferred event raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GatewayEVMUpgradeTestOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GatewayEVMUpgradeTest.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMUpgradeTestOwnershipTransferredIterator{contract: _GatewayEVMUpgradeTest.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GatewayEVMUpgradeTestOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GatewayEVMUpgradeTest.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMUpgradeTestOwnershipTransferred)
				if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) ParseOwnershipTransferred(log types.Log) (*GatewayEVMUpgradeTestOwnershipTransferred, error) {
	event := new(GatewayEVMUpgradeTestOwnershipTransferred)
	if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMUpgradeTestRevertedIterator is returned from FilterReverted and is used to iterate over the raw logs and unpacked data for Reverted events raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestRevertedIterator struct {
	Event *GatewayEVMUpgradeTestReverted // Event containing the contract specifics and raw log

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
func (it *GatewayEVMUpgradeTestRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMUpgradeTestReverted)
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
		it.Event = new(GatewayEVMUpgradeTestReverted)
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
func (it *GatewayEVMUpgradeTestRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMUpgradeTestRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMUpgradeTestReverted represents a Reverted event raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestReverted struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterReverted is a free log retrieval operation binding the contract event 0xd5d7616b1678354a0dea9d7e57e6a090bff5babe9f8d6381fdbad16e89ba311c.
//
// Solidity: event Reverted(address indexed destination, uint256 value, bytes data)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) FilterReverted(opts *bind.FilterOpts, destination []common.Address) (*GatewayEVMUpgradeTestRevertedIterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _GatewayEVMUpgradeTest.contract.FilterLogs(opts, "Reverted", destinationRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMUpgradeTestRevertedIterator{contract: _GatewayEVMUpgradeTest.contract, event: "Reverted", logs: logs, sub: sub}, nil
}

// WatchReverted is a free log subscription operation binding the contract event 0xd5d7616b1678354a0dea9d7e57e6a090bff5babe9f8d6381fdbad16e89ba311c.
//
// Solidity: event Reverted(address indexed destination, uint256 value, bytes data)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) WatchReverted(opts *bind.WatchOpts, sink chan<- *GatewayEVMUpgradeTestReverted, destination []common.Address) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _GatewayEVMUpgradeTest.contract.WatchLogs(opts, "Reverted", destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMUpgradeTestReverted)
				if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "Reverted", log); err != nil {
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
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) ParseReverted(log types.Log) (*GatewayEVMUpgradeTestReverted, error) {
	event := new(GatewayEVMUpgradeTestReverted)
	if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "Reverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMUpgradeTestRevertedWithERC20Iterator is returned from FilterRevertedWithERC20 and is used to iterate over the raw logs and unpacked data for RevertedWithERC20 events raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestRevertedWithERC20Iterator struct {
	Event *GatewayEVMUpgradeTestRevertedWithERC20 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMUpgradeTestRevertedWithERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMUpgradeTestRevertedWithERC20)
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
		it.Event = new(GatewayEVMUpgradeTestRevertedWithERC20)
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
func (it *GatewayEVMUpgradeTestRevertedWithERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMUpgradeTestRevertedWithERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMUpgradeTestRevertedWithERC20 represents a RevertedWithERC20 event raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestRevertedWithERC20 struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRevertedWithERC20 is a free log retrieval operation binding the contract event 0x723fc7be2448075379e4fdf1e6bf5fead954d2668d2da05dcb44ccfec4beeda7.
//
// Solidity: event RevertedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) FilterRevertedWithERC20(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*GatewayEVMUpgradeTestRevertedWithERC20Iterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GatewayEVMUpgradeTest.contract.FilterLogs(opts, "RevertedWithERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMUpgradeTestRevertedWithERC20Iterator{contract: _GatewayEVMUpgradeTest.contract, event: "RevertedWithERC20", logs: logs, sub: sub}, nil
}

// WatchRevertedWithERC20 is a free log subscription operation binding the contract event 0x723fc7be2448075379e4fdf1e6bf5fead954d2668d2da05dcb44ccfec4beeda7.
//
// Solidity: event RevertedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) WatchRevertedWithERC20(opts *bind.WatchOpts, sink chan<- *GatewayEVMUpgradeTestRevertedWithERC20, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GatewayEVMUpgradeTest.contract.WatchLogs(opts, "RevertedWithERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMUpgradeTestRevertedWithERC20)
				if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "RevertedWithERC20", log); err != nil {
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
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) ParseRevertedWithERC20(log types.Log) (*GatewayEVMUpgradeTestRevertedWithERC20, error) {
	event := new(GatewayEVMUpgradeTestRevertedWithERC20)
	if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "RevertedWithERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMUpgradeTestUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestUpgradedIterator struct {
	Event *GatewayEVMUpgradeTestUpgraded // Event containing the contract specifics and raw log

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
func (it *GatewayEVMUpgradeTestUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMUpgradeTestUpgraded)
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
		it.Event = new(GatewayEVMUpgradeTestUpgraded)
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
func (it *GatewayEVMUpgradeTestUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMUpgradeTestUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMUpgradeTestUpgraded represents a Upgraded event raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*GatewayEVMUpgradeTestUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _GatewayEVMUpgradeTest.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMUpgradeTestUpgradedIterator{contract: _GatewayEVMUpgradeTest.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *GatewayEVMUpgradeTestUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _GatewayEVMUpgradeTest.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMUpgradeTestUpgraded)
				if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) ParseUpgraded(log types.Log) (*GatewayEVMUpgradeTestUpgraded, error) {
	event := new(GatewayEVMUpgradeTestUpgraded)
	if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
