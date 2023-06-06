# evm/tools/ZetaTokenConsumerUniV2.strategy.md

## ZetaTokenConsumerUniV2Errors

### InputCantBeZero

```solidity
error InputCantBeZero()
```

## ZetaTokenConsumerUniV2

_Uniswap V2 strategy for ZetaTokenConsumer_

### MAX_DEADLINE

```solidity
uint256 MAX_DEADLINE
```

### wETH

```solidity
address wETH
```

### zetaToken

```solidity
address zetaToken
```

### uniswapV2Router

```solidity
contract IUniswapV2Router02 uniswapV2Router
```

### constructor

```solidity
constructor(address zetaToken_, address uniswapV2Router_) public
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

