// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testzcontract

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

// RevertContext is an auto generated low-level Go binding around an user-defined struct.
type RevertContext struct {
	Origin  []byte
	Sender  common.Address
	ChainID *big.Int
}

// ZContext is an auto generated low-level Go binding around an user-defined struct.
type ZContext struct {
	Origin  []byte
	Sender  common.Address
	ChainID *big.Int
}

// TestZContractMetaData contains all meta data concerning the TestZContract contract.
var TestZContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"origin\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"ContextData\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"origin\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"internalType\":\"structzContext\",\"name\":\"context\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"zrc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"onCrossChainCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"origin\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"internalType\":\"structrevertContext\",\"name\":\"context\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"zrc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"onRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506107b1806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806369582bee1461003b578063de43156e14610057575b600080fd5b61005560048036038101906100509190610317565b610073565b005b610071600480360381019061006c91906103bb565b6100f3565b005b6000828281019061008491906102ce565b90507fcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e8680600001906100b79190610546565b8860200160208101906100ca91906102a1565b896040013533866040516100e3969594939291906104e3565b60405180910390a1505050505050565b6000828281019061010491906102ce565b90507fcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e8680600001906101379190610546565b88602001602081019061014a91906102a1565b89604001353386604051610163969594939291906104e3565b60405180910390a1505050505050565b6000610186610181846105ce565b6105a9565b9050828152602081018484840111156101a2576101a161072d565b5b6101ad848285610668565b509392505050565b6000813590506101c48161074d565b92915050565b60008083601f8401126101e0576101df61070f565b5b8235905067ffffffffffffffff8111156101fd576101fc61070a565b5b60208301915083600182028301111561021957610218610723565b5b9250929050565b600082601f8301126102355761023461070f565b5b8135610245848260208601610173565b91505092915050565b60006060828403121561026457610263610719565b5b81905092915050565b60006060828403121561028357610282610719565b5b81905092915050565b60008135905061029b81610764565b92915050565b6000602082840312156102b7576102b6610737565b5b60006102c5848285016101b5565b91505092915050565b6000602082840312156102e4576102e3610737565b5b600082013567ffffffffffffffff81111561030257610301610732565b5b61030e84828501610220565b91505092915050565b60008060008060006080868803121561033357610332610737565b5b600086013567ffffffffffffffff81111561035157610350610732565b5b61035d8882890161024e565b955050602061036e888289016101b5565b945050604061037f8882890161028c565b935050606086013567ffffffffffffffff8111156103a05761039f610732565b5b6103ac888289016101ca565b92509250509295509295909350565b6000806000806000608086880312156103d7576103d6610737565b5b600086013567ffffffffffffffff8111156103f5576103f4610732565b5b6104018882890161026d565b9550506020610412888289016101b5565b94505060406104238882890161028c565b935050606086013567ffffffffffffffff81111561044457610443610732565b5b610450888289016101ca565b92509250509295509295909350565b6104688161062c565b82525050565b600061047a838561060a565b9350610487838584610668565b6104908361073c565b840190509392505050565b60006104a6826105ff565b6104b0818561061b565b93506104c0818560208601610677565b6104c98161073c565b840191505092915050565b6104dd8161065e565b82525050565b600060a08201905081810360008301526104fe81888a61046e565b905061050d602083018761045f565b61051a60408301866104d4565b610527606083018561045f565b8181036080830152610539818461049b565b9050979650505050505050565b600080833560016020038436030381126105635761056261071e565b5b80840192508235915067ffffffffffffffff82111561058557610584610714565b5b6020830192506001820236038313156105a1576105a0610728565b5b509250929050565b60006105b36105c4565b90506105bf82826106aa565b919050565b6000604051905090565b600067ffffffffffffffff8211156105e9576105e86106db565b5b6105f28261073c565b9050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b60006106378261063e565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b8381101561069557808201518184015260208101905061067a565b838111156106a4576000848401525b50505050565b6106b38261073c565b810181811067ffffffffffffffff821117156106d2576106d16106db565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b6107568161062c565b811461076157600080fd5b50565b61076d8161065e565b811461077857600080fd5b5056fea264697066735822122066cf44b92a869772f8263906f54b9c640301bcfb59e1f9b384defb872c4de02564736f6c63430008070033",
}

// TestZContractABI is the input ABI used to generate the binding from.
// Deprecated: Use TestZContractMetaData.ABI instead.
var TestZContractABI = TestZContractMetaData.ABI

// TestZContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestZContractMetaData.Bin instead.
var TestZContractBin = TestZContractMetaData.Bin

// DeployTestZContract deploys a new Ethereum contract, binding an instance of TestZContract to it.
func DeployTestZContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TestZContract, error) {
	parsed, err := TestZContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestZContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestZContract{TestZContractCaller: TestZContractCaller{contract: contract}, TestZContractTransactor: TestZContractTransactor{contract: contract}, TestZContractFilterer: TestZContractFilterer{contract: contract}}, nil
}

// TestZContract is an auto generated Go binding around an Ethereum contract.
type TestZContract struct {
	TestZContractCaller     // Read-only binding to the contract
	TestZContractTransactor // Write-only binding to the contract
	TestZContractFilterer   // Log filterer for contract events
}

// TestZContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestZContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestZContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestZContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestZContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestZContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestZContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestZContractSession struct {
	Contract     *TestZContract    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestZContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestZContractCallerSession struct {
	Contract *TestZContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TestZContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestZContractTransactorSession struct {
	Contract     *TestZContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TestZContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestZContractRaw struct {
	Contract *TestZContract // Generic contract binding to access the raw methods on
}

// TestZContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestZContractCallerRaw struct {
	Contract *TestZContractCaller // Generic read-only contract binding to access the raw methods on
}

// TestZContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestZContractTransactorRaw struct {
	Contract *TestZContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestZContract creates a new instance of TestZContract, bound to a specific deployed contract.
func NewTestZContract(address common.Address, backend bind.ContractBackend) (*TestZContract, error) {
	contract, err := bindTestZContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestZContract{TestZContractCaller: TestZContractCaller{contract: contract}, TestZContractTransactor: TestZContractTransactor{contract: contract}, TestZContractFilterer: TestZContractFilterer{contract: contract}}, nil
}

// NewTestZContractCaller creates a new read-only instance of TestZContract, bound to a specific deployed contract.
func NewTestZContractCaller(address common.Address, caller bind.ContractCaller) (*TestZContractCaller, error) {
	contract, err := bindTestZContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestZContractCaller{contract: contract}, nil
}

// NewTestZContractTransactor creates a new write-only instance of TestZContract, bound to a specific deployed contract.
func NewTestZContractTransactor(address common.Address, transactor bind.ContractTransactor) (*TestZContractTransactor, error) {
	contract, err := bindTestZContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestZContractTransactor{contract: contract}, nil
}

// NewTestZContractFilterer creates a new log filterer instance of TestZContract, bound to a specific deployed contract.
func NewTestZContractFilterer(address common.Address, filterer bind.ContractFilterer) (*TestZContractFilterer, error) {
	contract, err := bindTestZContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestZContractFilterer{contract: contract}, nil
}

// bindTestZContract binds a generic wrapper to an already deployed contract.
func bindTestZContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestZContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestZContract *TestZContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestZContract.Contract.TestZContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestZContract *TestZContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestZContract.Contract.TestZContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestZContract *TestZContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestZContract.Contract.TestZContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestZContract *TestZContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestZContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestZContract *TestZContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestZContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestZContract *TestZContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestZContract.Contract.contract.Transact(opts, method, params...)
}

// OnCrossChainCall is a paid mutator transaction binding the contract method 0xde43156e.
//
// Solidity: function onCrossChainCall((bytes,address,uint256) context, address zrc20, uint256 amount, bytes message) returns()
func (_TestZContract *TestZContractTransactor) OnCrossChainCall(opts *bind.TransactOpts, context ZContext, zrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _TestZContract.contract.Transact(opts, "onCrossChainCall", context, zrc20, amount, message)
}

// OnCrossChainCall is a paid mutator transaction binding the contract method 0xde43156e.
//
// Solidity: function onCrossChainCall((bytes,address,uint256) context, address zrc20, uint256 amount, bytes message) returns()
func (_TestZContract *TestZContractSession) OnCrossChainCall(context ZContext, zrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _TestZContract.Contract.OnCrossChainCall(&_TestZContract.TransactOpts, context, zrc20, amount, message)
}

// OnCrossChainCall is a paid mutator transaction binding the contract method 0xde43156e.
//
// Solidity: function onCrossChainCall((bytes,address,uint256) context, address zrc20, uint256 amount, bytes message) returns()
func (_TestZContract *TestZContractTransactorSession) OnCrossChainCall(context ZContext, zrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _TestZContract.Contract.OnCrossChainCall(&_TestZContract.TransactOpts, context, zrc20, amount, message)
}

// OnRevert is a paid mutator transaction binding the contract method 0x69582bee.
//
// Solidity: function onRevert((bytes,address,uint256) context, address zrc20, uint256 amount, bytes message) returns()
func (_TestZContract *TestZContractTransactor) OnRevert(opts *bind.TransactOpts, context RevertContext, zrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _TestZContract.contract.Transact(opts, "onRevert", context, zrc20, amount, message)
}

// OnRevert is a paid mutator transaction binding the contract method 0x69582bee.
//
// Solidity: function onRevert((bytes,address,uint256) context, address zrc20, uint256 amount, bytes message) returns()
func (_TestZContract *TestZContractSession) OnRevert(context RevertContext, zrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _TestZContract.Contract.OnRevert(&_TestZContract.TransactOpts, context, zrc20, amount, message)
}

// OnRevert is a paid mutator transaction binding the contract method 0x69582bee.
//
// Solidity: function onRevert((bytes,address,uint256) context, address zrc20, uint256 amount, bytes message) returns()
func (_TestZContract *TestZContractTransactorSession) OnRevert(context RevertContext, zrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _TestZContract.Contract.OnRevert(&_TestZContract.TransactOpts, context, zrc20, amount, message)
}

// TestZContractContextDataIterator is returned from FilterContextData and is used to iterate over the raw logs and unpacked data for ContextData events raised by the TestZContract contract.
type TestZContractContextDataIterator struct {
	Event *TestZContractContextData // Event containing the contract specifics and raw log

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
func (it *TestZContractContextDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestZContractContextData)
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
		it.Event = new(TestZContractContextData)
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
func (it *TestZContractContextDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestZContractContextDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestZContractContextData represents a ContextData event raised by the TestZContract contract.
type TestZContractContextData struct {
	Origin    []byte
	Sender    common.Address
	ChainID   *big.Int
	MsgSender common.Address
	Message   string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterContextData is a free log retrieval operation binding the contract event 0xcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e.
//
// Solidity: event ContextData(bytes origin, address sender, uint256 chainID, address msgSender, string message)
func (_TestZContract *TestZContractFilterer) FilterContextData(opts *bind.FilterOpts) (*TestZContractContextDataIterator, error) {

	logs, sub, err := _TestZContract.contract.FilterLogs(opts, "ContextData")
	if err != nil {
		return nil, err
	}
	return &TestZContractContextDataIterator{contract: _TestZContract.contract, event: "ContextData", logs: logs, sub: sub}, nil
}

// WatchContextData is a free log subscription operation binding the contract event 0xcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e.
//
// Solidity: event ContextData(bytes origin, address sender, uint256 chainID, address msgSender, string message)
func (_TestZContract *TestZContractFilterer) WatchContextData(opts *bind.WatchOpts, sink chan<- *TestZContractContextData) (event.Subscription, error) {

	logs, sub, err := _TestZContract.contract.WatchLogs(opts, "ContextData")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestZContractContextData)
				if err := _TestZContract.contract.UnpackLog(event, "ContextData", log); err != nil {
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

// ParseContextData is a log parse operation binding the contract event 0xcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e.
//
// Solidity: event ContextData(bytes origin, address sender, uint256 chainID, address msgSender, string message)
func (_TestZContract *TestZContractFilterer) ParseContextData(log types.Log) (*TestZContractContextData, error) {
	event := new(TestZContractContextData)
	if err := _TestZContract.contract.UnpackLog(event, "ContextData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
