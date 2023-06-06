# Solidity API

## ZetaTokenConsumerUniV3Errors

### InputCantBeZero

```solidity
error InputCantBeZero()
```

### ErrorSendingETH

```solidity
error ErrorSendingETH()
```

### ReentrancyError

```solidity
error ReentrancyError()
```

## WETH9

### withdraw

```solidity
function withdraw(uint256 wad) external
```

## ZetaTokenConsumerUniV3

_Uniswap V3 strategy for ZetaTokenConsumer_

### MAX_DEADLINE

```solidity
uint256 MAX_DEADLINE
```

### zetaPoolFee

```solidity
uint24 zetaPoolFee
```

### tokenPoolFee

```solidity
uint24 tokenPoolFee
```

### WETH9Address

```solidity
address WETH9Address
```

### zetaToken

```solidity
address zetaToken
```

### uniswapV3Router

```solidity
contract ISwapRouter uniswapV3Router
```

### quoter

```solidity
contract IQuoter quoter
```

### _locked

```solidity
bool _locked
```

### constructor

```solidity
constructor(address zetaToken_, address uniswapV3Router_, address quoter_, address WETH9Address_, uint24 zetaPoolFee_, uint24 tokenPoolFee_) public
```

### nonReentrant

```solidity
modifier nonReentrant()
```

### receive

```solidity
receive() external payable
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

