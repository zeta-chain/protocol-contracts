# Solidity API

## ZetaInterfaces

### SendInput

```solidity
struct SendInput {
  uint256 destinationChainId;
  bytes destinationAddress;
  uint256 destinationGasLimit;
  bytes message;
  uint256 zetaValueAndGas;
  bytes zetaParams;
}
```

### ZetaMessage

```solidity
struct ZetaMessage {
  bytes zetaTxSenderAddress;
  uint256 sourceChainId;
  address destinationAddress;
  uint256 zetaValue;
  bytes message;
}
```

### ZetaRevert

```solidity
struct ZetaRevert {
  address zetaTxSenderAddress;
  uint256 sourceChainId;
  bytes destinationAddress;
  uint256 destinationChainId;
  uint256 remainingZetaValue;
  bytes message;
}
```

## ZetaConnector

### send

```solidity
function send(struct ZetaInterfaces.SendInput input) external
```

_Sending value and data cross-chain is as easy as calling connector.send(SendInput)_

## ZetaReceiver

### onZetaMessage

```solidity
function onZetaMessage(struct ZetaInterfaces.ZetaMessage zetaMessage) external
```

_onZetaMessage is called when a cross-chain message reaches a contract_

### onZetaRevert

```solidity
function onZetaRevert(struct ZetaInterfaces.ZetaRevert zetaRevert) external
```

_onZetaRevert is called when a cross-chain message reverts.
It's useful to rollback to the original state_

## ZetaTokenConsumer

_ZetaTokenConsumer makes it easier to handle the following situations:
  - Getting Zeta using native coin (to pay for destination gas while using `connector.send`)
  - Getting Zeta using a token (to pay for destination gas while using `connector.send`)
  - Getting native coin using Zeta (to return unused destination gas when `onZetaRevert` is executed)
  - Getting a token using Zeta (to return unused destination gas when `onZetaRevert` is executed)
The interface can be implemented using different strategies, like UniswapV2, UniswapV3, etc_

### EthExchangedForZeta

```solidity
event EthExchangedForZeta(uint256 amountIn, uint256 amountOut)
```

### TokenExchangedForZeta

```solidity
event TokenExchangedForZeta(address token, uint256 amountIn, uint256 amountOut)
```

### ZetaExchangedForEth

```solidity
event ZetaExchangedForEth(uint256 amountIn, uint256 amountOut)
```

### ZetaExchangedForToken

```solidity
event ZetaExchangedForToken(address token, uint256 amountIn, uint256 amountOut)
```

### getZetaFromEth

```solidity
function getZetaFromEth(address destinationAddress, uint256 minAmountOut) external payable returns (uint256)
```

### getZetaFromToken

```solidity
function getZetaFromToken(address destinationAddress, uint256 minAmountOut, address inputToken, uint256 inputTokenAmount) external returns (uint256)
```

### getEthFromZeta

```solidity
function getEthFromZeta(address destinationAddress, uint256 minAmountOut, uint256 zetaTokenAmount) external returns (uint256)
```

### getTokenFromZeta

```solidity
function getTokenFromZeta(address destinationAddress, uint256 minAmountOut, address outputToken, uint256 zetaTokenAmount) external returns (uint256)
```

## ZetaCommonErrors

### InvalidAddress

```solidity
error InvalidAddress()
```

