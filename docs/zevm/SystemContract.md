## SystemContractErrors

```solidity
import "@zetachain/protocol-contracts/contracts/zevm/SystemContract.sol";
```

Source: https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/zevm/SystemContract.sol

Custom errors for SystemContract

### Error List

* [CallerIsNotFungibleModule()](#SystemContractErrors-CallerIsNotFungibleModule--)
* [InvalidTarget()](#SystemContractErrors-InvalidTarget--)
* [CantBeIdenticalAddresses()](#SystemContractErrors-CantBeIdenticalAddresses--)
* [CantBeZeroAddress()](#SystemContractErrors-CantBeZeroAddress--)
* [ZeroAddress()](#SystemContractErrors-ZeroAddress--)

### Modifiers

### Errors

```
CallerIsNotFungibleModule() (error)
```

<a name="SystemContractErrors-CallerIsNotFungibleModule--"></a>

```
InvalidTarget() (error)
```

<a name="SystemContractErrors-InvalidTarget--"></a>

```
CantBeIdenticalAddresses() (error)
```

<a name="SystemContractErrors-CantBeIdenticalAddresses--"></a>

```
CantBeZeroAddress() (error)
```

<a name="SystemContractErrors-CantBeZeroAddress--"></a>

```
ZeroAddress() (error)
```

<a name="SystemContractErrors-ZeroAddress--"></a>

## SystemContract

```solidity
import "@zetachain/protocol-contracts/contracts/zevm/SystemContract.sol";
```

Source: https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/zevm/SystemContract.sol

The system contract it's called by the protocol to interact with the blockchain.
Also includes a lot of tools to make easier to interact with ZetaChain.

### Function List

* [constructor(wzeta_, uniswapv2Factory_, uniswapv2Router02_)](#SystemContract-constructor-address-address-address-)
* [depositAndCall(zrc20, amount, target, message)](#SystemContract-depositAndCall-address-uint256-address-bytes-)
* [sortTokens(tokenA, tokenB)](#SystemContract-sortTokens-address-address-)
* [uniswapv2PairFor(factory, tokenA, tokenB)](#SystemContract-uniswapv2PairFor-address-address-address-)
* [setGasPrice(chainID, price)](#SystemContract-setGasPrice-uint256-uint256-)
* [setGasCoinZRC20(chainID, zrc20)](#SystemContract-setGasCoinZRC20-uint256-address-)
* [setGasZetaPool(chainID, erc20)](#SystemContract-setGasZetaPool-uint256-address-)
* [setWZETAContractAddress(addr)](#SystemContract-setWZETAContractAddress-address-)
* [setConnectorZEVMAddress(addr)](#SystemContract-setConnectorZEVMAddress-address-)

### Event List

* [SystemContractDeployed()](#SystemContract-SystemContractDeployed--)
* [SetGasPrice(, )](#SystemContract-SetGasPrice-uint256-uint256-)
* [SetGasCoin(, )](#SystemContract-SetGasCoin-uint256-address-)
* [SetGasZetaPool(, )](#SystemContract-SetGasZetaPool-uint256-address-)
* [SetWZeta()](#SystemContract-SetWZeta-address-)
* [SetConnectorZEVM()](#SystemContract-SetConnectorZEVM-address-)

### Error List

* [CallerIsNotFungibleModule()](#SystemContractErrors-CallerIsNotFungibleModule--)
* [InvalidTarget()](#SystemContractErrors-InvalidTarget--)
* [CantBeIdenticalAddresses()](#SystemContractErrors-CantBeIdenticalAddresses--)
* [CantBeZeroAddress()](#SystemContractErrors-CantBeZeroAddress--)
* [ZeroAddress()](#SystemContractErrors-ZeroAddress--)

### Modifiers

### Functions

```
constructor(address wzeta_, address uniswapv2Factory_, address uniswapv2Router02_) (public function)
```

<a name="SystemContract-constructor-address-address-address-"></a>

Only fungible module can deploy a system contract.

```
depositAndCall(address zrc20, uint256 amount, address target, bytes message) (external function)
```

<a name="SystemContract-depositAndCall-address-uint256-address-bytes-"></a>

Deposit foreign coins into ZRC20 and call user specified contract on zEVM.

```
sortTokens(address tokenA, address tokenB) → address token0, address token1 (internal function)
```

<a name="SystemContract-sortTokens-address-address-"></a>

Sort token addresses lexicographically. Used to handle return values from pairs sorted in the order.

```
uniswapv2PairFor(address factory, address tokenA, address tokenB) → address pair (public function)
```

<a name="SystemContract-uniswapv2PairFor-address-address-address-"></a>

Calculates the CREATE2 address for a pair without making any external calls.

```
setGasPrice(uint256 chainID, uint256 price) (external function)
```

<a name="SystemContract-setGasPrice-uint256-uint256-"></a>

Fungible module updates the gas price oracle periodically.

```
setGasCoinZRC20(uint256 chainID, address zrc20) (external function)
```

<a name="SystemContract-setGasCoinZRC20-uint256-address-"></a>

Setter for gasCoinZRC20ByChainId map.

```
setGasZetaPool(uint256 chainID, address erc20) (external function)
```

<a name="SystemContract-setGasZetaPool-uint256-address-"></a>

Set the pool wzeta/erc20 address.

```
setWZETAContractAddress(address addr) (external function)
```

<a name="SystemContract-setWZETAContractAddress-address-"></a>

Setter for wrapped ZETA address.

```
setConnectorZEVMAddress(address addr) (external function)
```

<a name="SystemContract-setConnectorZEVMAddress-address-"></a>

Setter for zetaConnector ZEVM Address

### Events

```
SystemContractDeployed() (event)
```

<a name="SystemContract-SystemContractDeployed--"></a>

```
SetGasPrice(uint256, uint256) (event)
```

<a name="SystemContract-SetGasPrice-uint256-uint256-"></a>

```
SetGasCoin(uint256, address) (event)
```

<a name="SystemContract-SetGasCoin-uint256-address-"></a>

```
SetGasZetaPool(uint256, address) (event)
```

<a name="SystemContract-SetGasZetaPool-uint256-address-"></a>

```
SetWZeta(address) (event)
```

<a name="SystemContract-SetWZeta-address-"></a>

```
SetConnectorZEVM(address) (event)
```

<a name="SystemContract-SetConnectorZEVM-address-"></a>

### Errors

```
CallerIsNotFungibleModule() (error)
```

<a name="SystemContractErrors-CallerIsNotFungibleModule--"></a>

```
InvalidTarget() (error)
```

<a name="SystemContractErrors-InvalidTarget--"></a>

```
CantBeIdenticalAddresses() (error)
```

<a name="SystemContractErrors-CantBeIdenticalAddresses--"></a>

```
CantBeZeroAddress() (error)
```

<a name="SystemContractErrors-CantBeZeroAddress--"></a>

```
ZeroAddress() (error)
```

<a name="SystemContractErrors-ZeroAddress--"></a>

