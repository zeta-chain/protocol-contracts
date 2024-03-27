## ZetaTokenConsumerUniV3Errors

```solidity
import "@zetachain/protocol-contracts/contracts/evm/tools/ZetaTokenConsumerPancakeV3.strategy.sol";
```

Source: https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/evm/tools/ZetaTokenConsumerPancakeV3.strategy.sol

### Error List

* [InputCantBeZero()](#ZetaTokenConsumerUniV3Errors-InputCantBeZero--)
* [ErrorSendingETH()](#ZetaTokenConsumerUniV3Errors-ErrorSendingETH--)
* [ReentrancyError()](#ZetaTokenConsumerUniV3Errors-ReentrancyError--)

### Modifiers

### Errors

```
InputCantBeZero() (error)
```

<a name="ZetaTokenConsumerUniV3Errors-InputCantBeZero--"></a>

```
ErrorSendingETH() (error)
```

<a name="ZetaTokenConsumerUniV3Errors-ErrorSendingETH--"></a>

```
ReentrancyError() (error)
```

<a name="ZetaTokenConsumerUniV3Errors-ReentrancyError--"></a>

## WETH9

```solidity
import "@zetachain/protocol-contracts/contracts/evm/tools/ZetaTokenConsumerPancakeV3.strategy.sol";
```

Source: https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/evm/tools/ZetaTokenConsumerPancakeV3.strategy.sol

### Function List

* [withdraw(wad)](#WETH9-withdraw-uint256-)

### Modifiers

### Functions

```
withdraw(uint256 wad) (external function)
```

<a name="WETH9-withdraw-uint256-"></a>

## ISwapRouterPancake

```solidity
import "@zetachain/protocol-contracts/contracts/evm/tools/ZetaTokenConsumerPancakeV3.strategy.sol";
```

Source: https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/evm/tools/ZetaTokenConsumerPancakeV3.strategy.sol

### Function List

* [exactInputSingle(params)](#ISwapRouterPancake-exactInputSingle-struct-ISwapRouterPancake-ExactInputSingleParams-)
* [exactInput(params)](#ISwapRouterPancake-exactInput-struct-ISwapRouterPancake-ExactInputParams-)

### Modifiers

### Functions

```
exactInputSingle(struct ISwapRouterPancake.ExactInputSingleParams params) → uint256 amountOut (external function)
```

<a name="ISwapRouterPancake-exactInputSingle-struct-ISwapRouterPancake-ExactInputSingleParams-"></a>

```
exactInput(struct ISwapRouterPancake.ExactInputParams params) → uint256 amountOut (external function)
```

<a name="ISwapRouterPancake-exactInput-struct-ISwapRouterPancake-ExactInputParams-"></a>

## ZetaTokenConsumerPancakeV3

```solidity
import "@zetachain/protocol-contracts/contracts/evm/tools/ZetaTokenConsumerPancakeV3.strategy.sol";
```

Source: https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/evm/tools/ZetaTokenConsumerPancakeV3.strategy.sol

Uniswap V3 strategy for ZetaTokenConsumer

### Modifier List

* [nonReentrant()](#ZetaTokenConsumerPancakeV3-nonReentrant--)

### Function List

* [constructor(zetaToken_, pancakeV3Router_, uniswapV3Factory_, WETH9Address_, zetaPoolFee_, tokenPoolFee_)](#ZetaTokenConsumerPancakeV3-constructor-address-address-address-address-uint24-uint24-)
* [receive()](#ZetaTokenConsumerPancakeV3-receive--)
* [getZetaFromEth(destinationAddress, minAmountOut)](#ZetaTokenConsumerPancakeV3-getZetaFromEth-address-uint256-)
* [getZetaFromToken(destinationAddress, minAmountOut, inputToken, inputTokenAmount)](#ZetaTokenConsumerPancakeV3-getZetaFromToken-address-uint256-address-uint256-)
* [getEthFromZeta(destinationAddress, minAmountOut, zetaTokenAmount)](#ZetaTokenConsumerPancakeV3-getEthFromZeta-address-uint256-uint256-)
* [getTokenFromZeta(destinationAddress, minAmountOut, outputToken, zetaTokenAmount)](#ZetaTokenConsumerPancakeV3-getTokenFromZeta-address-uint256-address-uint256-)
* [hasZetaLiquidity()](#ZetaTokenConsumerPancakeV3-hasZetaLiquidity--)

### Event List

* [EthExchangedForZeta(amountIn, amountOut)](#ZetaTokenConsumer-EthExchangedForZeta-uint256-uint256-)
* [TokenExchangedForZeta(token, amountIn, amountOut)](#ZetaTokenConsumer-TokenExchangedForZeta-address-uint256-uint256-)
* [ZetaExchangedForEth(amountIn, amountOut)](#ZetaTokenConsumer-ZetaExchangedForEth-uint256-uint256-)
* [ZetaExchangedForToken(token, amountIn, amountOut)](#ZetaTokenConsumer-ZetaExchangedForToken-address-uint256-uint256-)

### Error List

* [InputCantBeZero()](#ZetaTokenConsumerUniV3Errors-InputCantBeZero--)
* [ErrorSendingETH()](#ZetaTokenConsumerUniV3Errors-ErrorSendingETH--)
* [ReentrancyError()](#ZetaTokenConsumerUniV3Errors-ReentrancyError--)

### Modifiers

```
nonReentrant() (modifier)
```

<a name="ZetaTokenConsumerPancakeV3-nonReentrant--"></a>

### Functions

```
constructor(address zetaToken_, address pancakeV3Router_, address uniswapV3Factory_, address WETH9Address_, uint24 zetaPoolFee_, uint24 tokenPoolFee_) (public function)
```

<a name="ZetaTokenConsumerPancakeV3-constructor-address-address-address-address-uint24-uint24-"></a>

```
receive() (external function)
```

<a name="ZetaTokenConsumerPancakeV3-receive--"></a>

```
getZetaFromEth(address destinationAddress, uint256 minAmountOut) → uint256 (external function)
```

<a name="ZetaTokenConsumerPancakeV3-getZetaFromEth-address-uint256-"></a>

```
getZetaFromToken(address destinationAddress, uint256 minAmountOut, address inputToken, uint256 inputTokenAmount) → uint256 (external function)
```

<a name="ZetaTokenConsumerPancakeV3-getZetaFromToken-address-uint256-address-uint256-"></a>

```
getEthFromZeta(address destinationAddress, uint256 minAmountOut, uint256 zetaTokenAmount) → uint256 (external function)
```

<a name="ZetaTokenConsumerPancakeV3-getEthFromZeta-address-uint256-uint256-"></a>

```
getTokenFromZeta(address destinationAddress, uint256 minAmountOut, address outputToken, uint256 zetaTokenAmount) → uint256 (external function)
```

<a name="ZetaTokenConsumerPancakeV3-getTokenFromZeta-address-uint256-address-uint256-"></a>

```
hasZetaLiquidity() → bool (external function)
```

<a name="ZetaTokenConsumerPancakeV3-hasZetaLiquidity--"></a>

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

<a name="ZetaTokenConsumerUniV3Errors-InputCantBeZero--"></a>

```
ErrorSendingETH() (error)
```

<a name="ZetaTokenConsumerUniV3Errors-ErrorSendingETH--"></a>

```
ReentrancyError() (error)
```

<a name="ZetaTokenConsumerUniV3Errors-ReentrancyError--"></a>

