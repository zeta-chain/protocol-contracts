## SystemContractErrors

```solidity
import "@zetachain/protocol-contracts/contracts/zevm/testing/SystemContractMock.sol";
```

Source: https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/zevm/testing/SystemContractMock.sol

### Error List

* [CallerIsNotFungibleModule()](#SystemContractErrors-CallerIsNotFungibleModule--)
* [InvalidTarget()](#SystemContractErrors-InvalidTarget--)
* [CantBeIdenticalAddresses()](#SystemContractErrors-CantBeIdenticalAddresses--)
* [CantBeZeroAddress()](#SystemContractErrors-CantBeZeroAddress--)

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

## SystemContractMock

```solidity
import "@zetachain/protocol-contracts/contracts/zevm/testing/SystemContractMock.sol";
```

Source: https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/zevm/testing/SystemContractMock.sol

### Function List

* [constructor(wzeta_, uniswapv2Factory_, uniswapv2Router02_)](#SystemContractMock-constructor-address-address-address-)
* [setGasPrice(chainID, price)](#SystemContractMock-setGasPrice-uint256-uint256-)
* [setGasCoinZRC20(chainID, zrc20)](#SystemContractMock-setGasCoinZRC20-uint256-address-)
* [setWZETAContractAddress(addr)](#SystemContractMock-setWZETAContractAddress-address-)
* [sortTokens(tokenA, tokenB)](#SystemContractMock-sortTokens-address-address-)
* [uniswapv2PairFor(factory, tokenA, tokenB)](#SystemContractMock-uniswapv2PairFor-address-address-address-)
* [onCrossChainCall(target, zrc20, amount, message)](#SystemContractMock-onCrossChainCall-address-address-uint256-bytes-)

### Event List

* [SystemContractDeployed()](#SystemContractMock-SystemContractDeployed--)
* [SetGasPrice(, )](#SystemContractMock-SetGasPrice-uint256-uint256-)
* [SetGasCoin(, )](#SystemContractMock-SetGasCoin-uint256-address-)
* [SetGasZetaPool(, )](#SystemContractMock-SetGasZetaPool-uint256-address-)
* [SetWZeta()](#SystemContractMock-SetWZeta-address-)

### Error List

* [CallerIsNotFungibleModule()](#SystemContractErrors-CallerIsNotFungibleModule--)
* [InvalidTarget()](#SystemContractErrors-InvalidTarget--)
* [CantBeIdenticalAddresses()](#SystemContractErrors-CantBeIdenticalAddresses--)
* [CantBeZeroAddress()](#SystemContractErrors-CantBeZeroAddress--)

### Modifiers

### Functions

```
constructor(address wzeta_, address uniswapv2Factory_, address uniswapv2Router02_) (public function)
```

<a name="SystemContractMock-constructor-address-address-address-"></a>

```
setGasPrice(uint256 chainID, uint256 price) (external function)
```

<a name="SystemContractMock-setGasPrice-uint256-uint256-"></a>

```
setGasCoinZRC20(uint256 chainID, address zrc20) (external function)
```

<a name="SystemContractMock-setGasCoinZRC20-uint256-address-"></a>

```
setWZETAContractAddress(address addr) (external function)
```

<a name="SystemContractMock-setWZETAContractAddress-address-"></a>

```
sortTokens(address tokenA, address tokenB) → address token0, address token1 (internal function)
```

<a name="SystemContractMock-sortTokens-address-address-"></a>

```
uniswapv2PairFor(address factory, address tokenA, address tokenB) → address pair (public function)
```

<a name="SystemContractMock-uniswapv2PairFor-address-address-address-"></a>

```
onCrossChainCall(address target, address zrc20, uint256 amount, bytes message) (external function)
```

<a name="SystemContractMock-onCrossChainCall-address-address-uint256-bytes-"></a>

### Events

```
SystemContractDeployed() (event)
```

<a name="SystemContractMock-SystemContractDeployed--"></a>

```
SetGasPrice(uint256, uint256) (event)
```

<a name="SystemContractMock-SetGasPrice-uint256-uint256-"></a>

```
SetGasCoin(uint256, address) (event)
```

<a name="SystemContractMock-SetGasCoin-uint256-address-"></a>

```
SetGasZetaPool(uint256, address) (event)
```

<a name="SystemContractMock-SetGasZetaPool-uint256-address-"></a>

```
SetWZeta(address) (event)
```

<a name="SystemContractMock-SetWZeta-address-"></a>

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

