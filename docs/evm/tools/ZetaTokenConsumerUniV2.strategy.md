## ZetaTokenConsumerUniV2Errors

```solidity
import "@zetachain/protocol-contracts/contracts/evm/tools/ZetaTokenConsumerUniV2.strategy.sol";
```

Source: https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/evm/tools/ZetaTokenConsumerUniV2.strategy.sol

### Error List

* [InputCantBeZero()](#ZetaTokenConsumerUniV2Errors-InputCantBeZero--)

### Modifiers

### Errors

```
InputCantBeZero() (error)
```

<a name="ZetaTokenConsumerUniV2Errors-InputCantBeZero--"></a>

## ZetaTokenConsumerUniV2

```solidity
import "@zetachain/protocol-contracts/contracts/evm/tools/ZetaTokenConsumerUniV2.strategy.sol";
```

Source: https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/evm/tools/ZetaTokenConsumerUniV2.strategy.sol

Uniswap V2 strategy for ZetaTokenConsumer

### Function List

* [constructor(zetaToken_, uniswapV2Router_)](#ZetaTokenConsumerUniV2-constructor-address-address-)
* [getZetaFromEth(destinationAddress, minAmountOut)](#ZetaTokenConsumerUniV2-getZetaFromEth-address-uint256-)
* [getZetaFromToken(destinationAddress, minAmountOut, inputToken, inputTokenAmount)](#ZetaTokenConsumerUniV2-getZetaFromToken-address-uint256-address-uint256-)
* [getEthFromZeta(destinationAddress, minAmountOut, zetaTokenAmount)](#ZetaTokenConsumerUniV2-getEthFromZeta-address-uint256-uint256-)
* [getTokenFromZeta(destinationAddress, minAmountOut, outputToken, zetaTokenAmount)](#ZetaTokenConsumerUniV2-getTokenFromZeta-address-uint256-address-uint256-)

### Event List

* [EthExchangedForZeta(amountIn, amountOut)](#ZetaTokenConsumer-EthExchangedForZeta-uint256-uint256-)
* [TokenExchangedForZeta(token, amountIn, amountOut)](#ZetaTokenConsumer-TokenExchangedForZeta-address-uint256-uint256-)
* [ZetaExchangedForEth(amountIn, amountOut)](#ZetaTokenConsumer-ZetaExchangedForEth-uint256-uint256-)
* [ZetaExchangedForToken(token, amountIn, amountOut)](#ZetaTokenConsumer-ZetaExchangedForToken-address-uint256-uint256-)

### Error List

* [InputCantBeZero()](#ZetaTokenConsumerUniV2Errors-InputCantBeZero--)

### Modifiers

### Functions

```
constructor(address zetaToken_, address uniswapV2Router_) (public function)
```

<a name="ZetaTokenConsumerUniV2-constructor-address-address-"></a>

```
getZetaFromEth(address destinationAddress, uint256 minAmountOut) → uint256 (external function)
```

<a name="ZetaTokenConsumerUniV2-getZetaFromEth-address-uint256-"></a>

```
getZetaFromToken(address destinationAddress, uint256 minAmountOut, address inputToken, uint256 inputTokenAmount) → uint256 (external function)
```

<a name="ZetaTokenConsumerUniV2-getZetaFromToken-address-uint256-address-uint256-"></a>

```
getEthFromZeta(address destinationAddress, uint256 minAmountOut, uint256 zetaTokenAmount) → uint256 (external function)
```

<a name="ZetaTokenConsumerUniV2-getEthFromZeta-address-uint256-uint256-"></a>

```
getTokenFromZeta(address destinationAddress, uint256 minAmountOut, address outputToken, uint256 zetaTokenAmount) → uint256 (external function)
```

<a name="ZetaTokenConsumerUniV2-getTokenFromZeta-address-uint256-address-uint256-"></a>

### Events

```
EthExchangedForZeta(uint256 amountIn, uint256 amountOut) (event)
```

<a name="ZetaTokenConsumer-EthExchangedForZeta-uint256-uint256-"></a>

```
TokenExchangedForZeta(address token, uint256 amountIn, uint256 amountOut) (event)
```

<a name="ZetaTokenConsumer-TokenExchangedForZeta-address-uint256-uint256-"></a>

```
ZetaExchangedForEth(uint256 amountIn, uint256 amountOut) (event)
```

<a name="ZetaTokenConsumer-ZetaExchangedForEth-uint256-uint256-"></a>

```
ZetaExchangedForToken(address token, uint256 amountIn, uint256 amountOut) (event)
```

<a name="ZetaTokenConsumer-ZetaExchangedForToken-address-uint256-uint256-"></a>

### Errors

```
InputCantBeZero() (error)
```

<a name="ZetaTokenConsumerUniV2Errors-InputCantBeZero--"></a>

