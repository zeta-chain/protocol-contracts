## ZetaTokenConsumerUniV3Errors

```solidity
import "@zetachain/protocol-contracts/contracts/evm/tools/ZetaTokenConsumerUniV3.strategy.sol";
```

Source: https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/evm/tools/ZetaTokenConsumerUniV3.strategy.sol

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
import "@zetachain/protocol-contracts/contracts/evm/tools/ZetaTokenConsumerUniV3.strategy.sol";
```

Source: https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/evm/tools/ZetaTokenConsumerUniV3.strategy.sol

### Function List

* [withdraw(wad)](#WETH9-withdraw-uint256-)

### Modifiers

### Functions

```
withdraw(uint256 wad) (external function)
```

<a name="WETH9-withdraw-uint256-"></a>

## ZetaTokenConsumerUniV3

```solidity
import "@zetachain/protocol-contracts/contracts/evm/tools/ZetaTokenConsumerUniV3.strategy.sol";
```

Source: https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/evm/tools/ZetaTokenConsumerUniV3.strategy.sol

Uniswap V3 strategy for ZetaTokenConsumer

### Modifier List

* [nonReentrant()](#ZetaTokenConsumerUniV3-nonReentrant--)

### Function List

* [constructor(zetaToken_, uniswapV3Router_, quoter_, WETH9Address_, zetaPoolFee_, tokenPoolFee_)](#ZetaTokenConsumerUniV3-constructor-address-address-address-address-uint24-uint24-)
* [receive()](#ZetaTokenConsumerUniV3-receive--)
* [getZetaFromEth(destinationAddress, minAmountOut)](#ZetaTokenConsumerUniV3-getZetaFromEth-address-uint256-)
* [getZetaFromToken(destinationAddress, minAmountOut, inputToken, inputTokenAmount)](#ZetaTokenConsumerUniV3-getZetaFromToken-address-uint256-address-uint256-)
* [getEthFromZeta(destinationAddress, minAmountOut, zetaTokenAmount)](#ZetaTokenConsumerUniV3-getEthFromZeta-address-uint256-uint256-)
* [getTokenFromZeta(destinationAddress, minAmountOut, outputToken, zetaTokenAmount)](#ZetaTokenConsumerUniV3-getTokenFromZeta-address-uint256-address-uint256-)

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

<a name="ZetaTokenConsumerUniV3-nonReentrant--"></a>

### Functions

```
constructor(address zetaToken_, address uniswapV3Router_, address quoter_, address WETH9Address_, uint24 zetaPoolFee_, uint24 tokenPoolFee_) (public function)
```

<a name="ZetaTokenConsumerUniV3-constructor-address-address-address-address-uint24-uint24-"></a>

```
receive() (external function)
```

<a name="ZetaTokenConsumerUniV3-receive--"></a>

```
getZetaFromEth(address destinationAddress, uint256 minAmountOut) → uint256 (external function)
```

<a name="ZetaTokenConsumerUniV3-getZetaFromEth-address-uint256-"></a>

```
getZetaFromToken(address destinationAddress, uint256 minAmountOut, address inputToken, uint256 inputTokenAmount) → uint256 (external function)
```

<a name="ZetaTokenConsumerUniV3-getZetaFromToken-address-uint256-address-uint256-"></a>

```
getEthFromZeta(address destinationAddress, uint256 minAmountOut, uint256 zetaTokenAmount) → uint256 (external function)
```

<a name="ZetaTokenConsumerUniV3-getEthFromZeta-address-uint256-uint256-"></a>

```
getTokenFromZeta(address destinationAddress, uint256 minAmountOut, address outputToken, uint256 zetaTokenAmount) → uint256 (external function)
```

<a name="ZetaTokenConsumerUniV3-getTokenFromZeta-address-uint256-address-uint256-"></a>

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

