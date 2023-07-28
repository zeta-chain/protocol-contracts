## INonfungiblePositionManager

```solidity
import "@zetachain/protocol-contracts/contracts/evm/testing/TestUniswapV3Contracts.sol";
```

Source: https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/evm/testing/TestUniswapV3Contracts.sol

### Function List

* [positions(tokenId)](#INonfungiblePositionManager-positions-uint256-)
* [mint(params)](#INonfungiblePositionManager-mint-struct-INonfungiblePositionManager-MintParams-)
* [increaseLiquidity(params)](#INonfungiblePositionManager-increaseLiquidity-struct-INonfungiblePositionManager-IncreaseLiquidityParams-)
* [decreaseLiquidity(params)](#INonfungiblePositionManager-decreaseLiquidity-struct-INonfungiblePositionManager-DecreaseLiquidityParams-)
* [collect(params)](#INonfungiblePositionManager-collect-struct-INonfungiblePositionManager-CollectParams-)
* [burn(tokenId)](#INonfungiblePositionManager-burn-uint256-)

### Event List

* [IncreaseLiquidity(tokenId, liquidity, amount0, amount1)](#INonfungiblePositionManager-IncreaseLiquidity-uint256-uint128-uint256-uint256-)
* [DecreaseLiquidity(tokenId, liquidity, amount0, amount1)](#INonfungiblePositionManager-DecreaseLiquidity-uint256-uint128-uint256-uint256-)
* [Collect(tokenId, recipient, amount0, amount1)](#INonfungiblePositionManager-Collect-uint256-address-uint256-uint256-)

### Modifiers

### Functions

```
positions(uint256 tokenId) → uint96 nonce, address operator, address token0, address token1, uint24 fee, int24 tickLower, int24 tickUpper, uint128 liquidity, uint256 feeGrowthInside0LastX128, uint256 feeGrowthInside1LastX128, uint128 tokensOwed0, uint128 tokensOwed1 (external function)
```

<a name="INonfungiblePositionManager-positions-uint256-"></a>

Throws if the token ID is not valid.

```
mint(struct INonfungiblePositionManager.MintParams params) → uint256 tokenId, uint128 liquidity, uint256 amount0, uint256 amount1 (external function)
```

<a name="INonfungiblePositionManager-mint-struct-INonfungiblePositionManager-MintParams-"></a>

Call this when the pool does exist and is initialized. Note that if the pool is created but not initialized
a method does not exist, i.e. the pool is assumed to be initialized.

```
increaseLiquidity(struct INonfungiblePositionManager.IncreaseLiquidityParams params) → uint128 liquidity, uint256 amount0, uint256 amount1 (external function)
```

<a name="INonfungiblePositionManager-increaseLiquidity-struct-INonfungiblePositionManager-IncreaseLiquidityParams-"></a>

```
decreaseLiquidity(struct INonfungiblePositionManager.DecreaseLiquidityParams params) → uint256 amount0, uint256 amount1 (external function)
```

<a name="INonfungiblePositionManager-decreaseLiquidity-struct-INonfungiblePositionManager-DecreaseLiquidityParams-"></a>

```
collect(struct INonfungiblePositionManager.CollectParams params) → uint256 amount0, uint256 amount1 (external function)
```

<a name="INonfungiblePositionManager-collect-struct-INonfungiblePositionManager-CollectParams-"></a>

```
burn(uint256 tokenId) (external function)
```

<a name="INonfungiblePositionManager-burn-uint256-"></a>

### Events

```
IncreaseLiquidity(uint256 indexed tokenId, uint128 liquidity, uint256 amount0, uint256 amount1) (event)
```

<a name="INonfungiblePositionManager-IncreaseLiquidity-uint256-uint128-uint256-uint256-"></a>

Also emitted when a token is minted

```
DecreaseLiquidity(uint256 indexed tokenId, uint128 liquidity, uint256 amount0, uint256 amount1) (event)
```

<a name="INonfungiblePositionManager-DecreaseLiquidity-uint256-uint128-uint256-uint256-"></a>

```
Collect(uint256 indexed tokenId, address recipient, uint256 amount0, uint256 amount1) (event)
```

<a name="INonfungiblePositionManager-Collect-uint256-address-uint256-uint256-"></a>

The amounts reported may not be exactly equivalent to the amounts transferred, due to rounding behavior

## IPoolInitializer

```solidity
import "@zetachain/protocol-contracts/contracts/evm/testing/TestUniswapV3Contracts.sol";
```

Source: https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/evm/testing/TestUniswapV3Contracts.sol

### Function List

* [createAndInitializePoolIfNecessary(token0, token1, fee, sqrtPriceX96)](#IPoolInitializer-createAndInitializePoolIfNecessary-address-address-uint24-uint160-)

### Modifiers

### Functions

```
createAndInitializePoolIfNecessary(address token0, address token1, uint24 fee, uint160 sqrtPriceX96) → address pool (external function)
```

<a name="IPoolInitializer-createAndInitializePoolIfNecessary-address-address-uint24-uint160-"></a>

This method can be bundled with others via IMulticall for the first action (e.g. mint) performed against a pool

