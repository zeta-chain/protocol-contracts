# evm/tools/ZetaTokenConsumerTrident.strategy.md

## ZetaTokenConsumerTridentErrors

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

### deposit

```solidity
function deposit() external payable
```

### withdraw

```solidity
function withdraw(uint256 wad) external
```

### depositTo

```solidity
function depositTo(address to) external payable
```

### withdrawTo

```solidity
function withdrawTo(address payable to, uint256 value) external
```

## ZetaTokenConsumerTrident

_Trident strategy for ZetaTokenConsumer_

### MAX_DEADLINE

```solidity
uint256 MAX_DEADLINE
```

### WETH9Address

```solidity
address WETH9Address
```

### zetaToken

```solidity
address zetaToken
```

### tridentRouter

```solidity
contract IPoolRouter tridentRouter
```

### poolFactory

```solidity
contract ConcentratedLiquidityPoolFactory poolFactory
```

### _locked

```solidity
bool _locked
```

### constructor

```solidity
constructor(address zetaToken_, address uniswapV3Router_, address WETH9Address_, address poolFactory_) public
```

### nonReentrant

```solidity
modifier nonReentrant()
```

### receive

```solidity
receive() external payable
```

### getPair

```solidity
function getPair(address token0, address token1) internal pure returns (address, address)
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

