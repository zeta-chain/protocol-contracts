## ZetaTokenConsumerTridentErrors

```solidity
import "@zetachain/protocol-contracts/contracts/evm/tools/ZetaTokenConsumerTrident.strategy.sol";
```

Source: https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/evm/tools/ZetaTokenConsumerTrident.strategy.sol

### Error List

* [InputCantBeZero()](#ZetaTokenConsumerTridentErrors-InputCantBeZero--)
* [ErrorSendingETH()](#ZetaTokenConsumerTridentErrors-ErrorSendingETH--)
* [ReentrancyError()](#ZetaTokenConsumerTridentErrors-ReentrancyError--)

### Modifiers

### Errors

```
InputCantBeZero() (error)
```

<a name="ZetaTokenConsumerTridentErrors-InputCantBeZero--"></a>

```
ErrorSendingETH() (error)
```

<a name="ZetaTokenConsumerTridentErrors-ErrorSendingETH--"></a>

```
ReentrancyError() (error)
```

<a name="ZetaTokenConsumerTridentErrors-ReentrancyError--"></a>

## WETH9

```solidity
import "@zetachain/protocol-contracts/contracts/evm/tools/ZetaTokenConsumerTrident.strategy.sol";
```

Source: https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/evm/tools/ZetaTokenConsumerTrident.strategy.sol

### Function List

* [deposit()](#WETH9-deposit--)
* [withdraw(wad)](#WETH9-withdraw-uint256-)
* [depositTo(to)](#WETH9-depositTo-address-)
* [withdrawTo(to, value)](#WETH9-withdrawTo-address-payable-uint256-)

### Modifiers

### Functions

```
deposit() (external function)
```

<a name="WETH9-deposit--"></a>

```
withdraw(uint256 wad) (external function)
```

<a name="WETH9-withdraw-uint256-"></a>

```
depositTo(address to) (external function)
```

<a name="WETH9-depositTo-address-"></a>

```
withdrawTo(address payable to, uint256 value) (external function)
```

<a name="WETH9-withdrawTo-address-payable-uint256-"></a>

## ZetaTokenConsumerTrident

```solidity
import "@zetachain/protocol-contracts/contracts/evm/tools/ZetaTokenConsumerTrident.strategy.sol";
```

Source: https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/evm/tools/ZetaTokenConsumerTrident.strategy.sol

Trident strategy for ZetaTokenConsumer

### Modifier List

* [nonReentrant()](#ZetaTokenConsumerTrident-nonReentrant--)

### Function List

* [constructor(zetaToken_, uniswapV3Router_, WETH9Address_, poolFactory_)](#ZetaTokenConsumerTrident-constructor-address-address-address-address-)
* [receive()](#ZetaTokenConsumerTrident-receive--)
* [getPair(token0, token1)](#ZetaTokenConsumerTrident-getPair-address-address-)
* [getZetaFromEth(destinationAddress, minAmountOut)](#ZetaTokenConsumerTrident-getZetaFromEth-address-uint256-)
* [getZetaFromToken(destinationAddress, minAmountOut, inputToken, inputTokenAmount)](#ZetaTokenConsumerTrident-getZetaFromToken-address-uint256-address-uint256-)
* [getEthFromZeta(destinationAddress, minAmountOut, zetaTokenAmount)](#ZetaTokenConsumerTrident-getEthFromZeta-address-uint256-uint256-)
* [getTokenFromZeta(destinationAddress, minAmountOut, outputToken, zetaTokenAmount)](#ZetaTokenConsumerTrident-getTokenFromZeta-address-uint256-address-uint256-)

### Event List

* [EthExchangedForZeta(amountIn, amountOut)](#ZetaTokenConsumer-EthExchangedForZeta-uint256-uint256-)
* [TokenExchangedForZeta(token, amountIn, amountOut)](#ZetaTokenConsumer-TokenExchangedForZeta-address-uint256-uint256-)
* [ZetaExchangedForEth(amountIn, amountOut)](#ZetaTokenConsumer-ZetaExchangedForEth-uint256-uint256-)
* [ZetaExchangedForToken(token, amountIn, amountOut)](#ZetaTokenConsumer-ZetaExchangedForToken-address-uint256-uint256-)

### Error List

* [InputCantBeZero()](#ZetaTokenConsumerTridentErrors-InputCantBeZero--)
* [ErrorSendingETH()](#ZetaTokenConsumerTridentErrors-ErrorSendingETH--)
* [ReentrancyError()](#ZetaTokenConsumerTridentErrors-ReentrancyError--)

### Modifiers

```
nonReentrant() (modifier)
```

<a name="ZetaTokenConsumerTrident-nonReentrant--"></a>

### Functions

```
constructor(address zetaToken_, address uniswapV3Router_, address WETH9Address_, address poolFactory_) (public function)
```

<a name="ZetaTokenConsumerTrident-constructor-address-address-address-address-"></a>

```
receive() (external function)
```

<a name="ZetaTokenConsumerTrident-receive--"></a>

```
getPair(address token0, address token1) → address, address (internal function)
```

<a name="ZetaTokenConsumerTrident-getPair-address-address-"></a>

```
getZetaFromEth(address destinationAddress, uint256 minAmountOut) → uint256 (external function)
```

<a name="ZetaTokenConsumerTrident-getZetaFromEth-address-uint256-"></a>

```
getZetaFromToken(address destinationAddress, uint256 minAmountOut, address inputToken, uint256 inputTokenAmount) → uint256 (external function)
```

<a name="ZetaTokenConsumerTrident-getZetaFromToken-address-uint256-address-uint256-"></a>

```
getEthFromZeta(address destinationAddress, uint256 minAmountOut, uint256 zetaTokenAmount) → uint256 (external function)
```

<a name="ZetaTokenConsumerTrident-getEthFromZeta-address-uint256-uint256-"></a>

```
getTokenFromZeta(address destinationAddress, uint256 minAmountOut, address outputToken, uint256 zetaTokenAmount) → uint256 (external function)
```

<a name="ZetaTokenConsumerTrident-getTokenFromZeta-address-uint256-address-uint256-"></a>

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

<a name="ZetaTokenConsumerTridentErrors-InputCantBeZero--"></a>

```
ErrorSendingETH() (error)
```

<a name="ZetaTokenConsumerTridentErrors-ErrorSendingETH--"></a>

```
ReentrancyError() (error)
```

<a name="ZetaTokenConsumerTridentErrors-ReentrancyError--"></a>

