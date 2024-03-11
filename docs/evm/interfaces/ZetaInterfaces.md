## ZetaInterfaces

```solidity
import "@zetachain/protocol-contracts/contracts/evm/interfaces/ZetaInterfaces.sol";
```

Source: https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/evm/interfaces/ZetaInterfaces.sol

### Modifiers

## ZetaConnector

```solidity
import "@zetachain/protocol-contracts/contracts/evm/interfaces/ZetaInterfaces.sol";
```

Source: https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/evm/interfaces/ZetaInterfaces.sol

### Function List

* [send(input)](#ZetaConnector-send-struct-ZetaInterfaces-SendInput-)

### Modifiers

### Functions

```
send(struct ZetaInterfaces.SendInput input) (external function)
```

<a name="ZetaConnector-send-struct-ZetaInterfaces-SendInput-"></a>

Sending value and data cross-chain is as easy as calling connector.send(SendInput)

## ZetaReceiver

```solidity
import "@zetachain/protocol-contracts/contracts/evm/interfaces/ZetaInterfaces.sol";
```

Source: https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/evm/interfaces/ZetaInterfaces.sol

### Function List

* [onZetaMessage(zetaMessage)](#ZetaReceiver-onZetaMessage-struct-ZetaInterfaces-ZetaMessage-)
* [onZetaRevert(zetaRevert)](#ZetaReceiver-onZetaRevert-struct-ZetaInterfaces-ZetaRevert-)

### Modifiers

### Functions

```
onZetaMessage(struct ZetaInterfaces.ZetaMessage zetaMessage) (external function)
```

<a name="ZetaReceiver-onZetaMessage-struct-ZetaInterfaces-ZetaMessage-"></a>

onZetaMessage is called when a cross-chain message reaches a contract

```
onZetaRevert(struct ZetaInterfaces.ZetaRevert zetaRevert) (external function)
```

<a name="ZetaReceiver-onZetaRevert-struct-ZetaInterfaces-ZetaRevert-"></a>

onZetaRevert is called when a cross-chain message reverts.
It's useful to rollback to the original state

## ZetaTokenConsumer

```solidity
import "@zetachain/protocol-contracts/contracts/evm/interfaces/ZetaInterfaces.sol";
```

Source: https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/evm/interfaces/ZetaInterfaces.sol

ZetaTokenConsumer makes it easier to handle the following situations:
  - Getting Zeta using native coin (to pay for destination gas while using `connector.send`)
  - Getting Zeta using a token (to pay for destination gas while using `connector.send`)
  - Getting native coin using Zeta (to return unused destination gas when `onZetaRevert` is executed)
  - Getting a token using Zeta (to return unused destination gas when `onZetaRevert` is executed)
The interface can be implemented using different strategies, like UniswapV2, UniswapV3, etc

### Function List

* [getZetaFromEth(destinationAddress, minAmountOut)](#ZetaTokenConsumer-getZetaFromEth-address-uint256-)
* [getZetaFromToken(destinationAddress, minAmountOut, inputToken, inputTokenAmount)](#ZetaTokenConsumer-getZetaFromToken-address-uint256-address-uint256-)
* [getEthFromZeta(destinationAddress, minAmountOut, zetaTokenAmount)](#ZetaTokenConsumer-getEthFromZeta-address-uint256-uint256-)
* [getTokenFromZeta(destinationAddress, minAmountOut, outputToken, zetaTokenAmount)](#ZetaTokenConsumer-getTokenFromZeta-address-uint256-address-uint256-)
* [hasZetaLiquidity()](#ZetaTokenConsumer-hasZetaLiquidity--)

### Event List

* [EthExchangedForZeta(amountIn, amountOut)](#ZetaTokenConsumer-EthExchangedForZeta-uint256-uint256-)
* [TokenExchangedForZeta(token, amountIn, amountOut)](#ZetaTokenConsumer-TokenExchangedForZeta-address-uint256-uint256-)
* [ZetaExchangedForEth(amountIn, amountOut)](#ZetaTokenConsumer-ZetaExchangedForEth-uint256-uint256-)
* [ZetaExchangedForToken(token, amountIn, amountOut)](#ZetaTokenConsumer-ZetaExchangedForToken-address-uint256-uint256-)

### Modifiers

### Functions

```
getZetaFromEth(address destinationAddress, uint256 minAmountOut) → uint256 (external function)
```

<a name="ZetaTokenConsumer-getZetaFromEth-address-uint256-"></a>

```
getZetaFromToken(address destinationAddress, uint256 minAmountOut, address inputToken, uint256 inputTokenAmount) → uint256 (external function)
```

<a name="ZetaTokenConsumer-getZetaFromToken-address-uint256-address-uint256-"></a>

```
getEthFromZeta(address destinationAddress, uint256 minAmountOut, uint256 zetaTokenAmount) → uint256 (external function)
```

<a name="ZetaTokenConsumer-getEthFromZeta-address-uint256-uint256-"></a>

```
getTokenFromZeta(address destinationAddress, uint256 minAmountOut, address outputToken, uint256 zetaTokenAmount) → uint256 (external function)
```

<a name="ZetaTokenConsumer-getTokenFromZeta-address-uint256-address-uint256-"></a>

```
hasZetaLiquidity() → bool (external function)
```

<a name="ZetaTokenConsumer-hasZetaLiquidity--"></a>

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

## ZetaCommonErrors

```solidity
import "@zetachain/protocol-contracts/contracts/evm/interfaces/ZetaInterfaces.sol";
```

Source: https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/evm/interfaces/ZetaInterfaces.sol

### Error List

* [InvalidAddress()](#ZetaCommonErrors-InvalidAddress--)

### Modifiers

### Errors

```
InvalidAddress() (error)
```

<a name="ZetaCommonErrors-InvalidAddress--"></a>

