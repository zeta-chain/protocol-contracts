## IUniswapV2Router01

```solidity
import "@zetachain/protocol-contracts/contracts/zevm/interfaces/IUniswapV2Router01.sol";
```

Source: https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/zevm/interfaces/IUniswapV2Router01.sol

### Function List

* [factory()](#IUniswapV2Router01-factory--)
* [WETH()](#IUniswapV2Router01-WETH--)
* [addLiquidity(tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)](#IUniswapV2Router01-addLiquidity-address-address-uint256-uint256-uint256-uint256-address-uint256-)
* [addLiquidityETH(token, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)](#IUniswapV2Router01-addLiquidityETH-address-uint256-uint256-uint256-address-uint256-)
* [removeLiquidity(tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)](#IUniswapV2Router01-removeLiquidity-address-address-uint256-uint256-uint256-address-uint256-)
* [removeLiquidityETH(token, liquidity, amountTokenMin, amountETHMin, to, deadline)](#IUniswapV2Router01-removeLiquidityETH-address-uint256-uint256-uint256-address-uint256-)
* [removeLiquidityWithPermit(tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)](#IUniswapV2Router01-removeLiquidityWithPermit-address-address-uint256-uint256-uint256-address-uint256-bool-uint8-bytes32-bytes32-)
* [removeLiquidityETHWithPermit(token, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)](#IUniswapV2Router01-removeLiquidityETHWithPermit-address-uint256-uint256-uint256-address-uint256-bool-uint8-bytes32-bytes32-)
* [swapExactTokensForTokens(amountIn, amountOutMin, path, to, deadline)](#IUniswapV2Router01-swapExactTokensForTokens-uint256-uint256-address---address-uint256-)
* [swapTokensForExactTokens(amountOut, amountInMax, path, to, deadline)](#IUniswapV2Router01-swapTokensForExactTokens-uint256-uint256-address---address-uint256-)
* [swapExactETHForTokens(amountOutMin, path, to, deadline)](#IUniswapV2Router01-swapExactETHForTokens-uint256-address---address-uint256-)
* [swapTokensForExactETH(amountOut, amountInMax, path, to, deadline)](#IUniswapV2Router01-swapTokensForExactETH-uint256-uint256-address---address-uint256-)
* [swapExactTokensForETH(amountIn, amountOutMin, path, to, deadline)](#IUniswapV2Router01-swapExactTokensForETH-uint256-uint256-address---address-uint256-)
* [swapETHForExactTokens(amountOut, path, to, deadline)](#IUniswapV2Router01-swapETHForExactTokens-uint256-address---address-uint256-)
* [quote(amountA, reserveA, reserveB)](#IUniswapV2Router01-quote-uint256-uint256-uint256-)
* [getAmountOut(amountIn, reserveIn, reserveOut)](#IUniswapV2Router01-getAmountOut-uint256-uint256-uint256-)
* [getAmountIn(amountOut, reserveIn, reserveOut)](#IUniswapV2Router01-getAmountIn-uint256-uint256-uint256-)
* [getAmountsOut(amountIn, path)](#IUniswapV2Router01-getAmountsOut-uint256-address---)
* [getAmountsIn(amountOut, path)](#IUniswapV2Router01-getAmountsIn-uint256-address---)

### Modifiers

### Functions

```
factory() → address (external function)
```

<a name="IUniswapV2Router01-factory--"></a>

```
WETH() → address (external function)
```

<a name="IUniswapV2Router01-WETH--"></a>

```
addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) → uint256 amountA, uint256 amountB, uint256 liquidity (external function)
```

<a name="IUniswapV2Router01-addLiquidity-address-address-uint256-uint256-uint256-uint256-address-uint256-"></a>

```
addLiquidityETH(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) → uint256 amountToken, uint256 amountETH, uint256 liquidity (external function)
```

<a name="IUniswapV2Router01-addLiquidityETH-address-uint256-uint256-uint256-address-uint256-"></a>

```
removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) → uint256 amountA, uint256 amountB (external function)
```

<a name="IUniswapV2Router01-removeLiquidity-address-address-uint256-uint256-uint256-address-uint256-"></a>

```
removeLiquidityETH(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) → uint256 amountToken, uint256 amountETH (external function)
```

<a name="IUniswapV2Router01-removeLiquidityETH-address-uint256-uint256-uint256-address-uint256-"></a>

```
removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) → uint256 amountA, uint256 amountB (external function)
```

<a name="IUniswapV2Router01-removeLiquidityWithPermit-address-address-uint256-uint256-uint256-address-uint256-bool-uint8-bytes32-bytes32-"></a>

```
removeLiquidityETHWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) → uint256 amountToken, uint256 amountETH (external function)
```

<a name="IUniswapV2Router01-removeLiquidityETHWithPermit-address-uint256-uint256-uint256-address-uint256-bool-uint8-bytes32-bytes32-"></a>

```
swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) → uint256[] amounts (external function)
```

<a name="IUniswapV2Router01-swapExactTokensForTokens-uint256-uint256-address---address-uint256-"></a>

```
swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) → uint256[] amounts (external function)
```

<a name="IUniswapV2Router01-swapTokensForExactTokens-uint256-uint256-address---address-uint256-"></a>

```
swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) → uint256[] amounts (external function)
```

<a name="IUniswapV2Router01-swapExactETHForTokens-uint256-address---address-uint256-"></a>

```
swapTokensForExactETH(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) → uint256[] amounts (external function)
```

<a name="IUniswapV2Router01-swapTokensForExactETH-uint256-uint256-address---address-uint256-"></a>

```
swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) → uint256[] amounts (external function)
```

<a name="IUniswapV2Router01-swapExactTokensForETH-uint256-uint256-address---address-uint256-"></a>

```
swapETHForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) → uint256[] amounts (external function)
```

<a name="IUniswapV2Router01-swapETHForExactTokens-uint256-address---address-uint256-"></a>

```
quote(uint256 amountA, uint256 reserveA, uint256 reserveB) → uint256 amountB (external function)
```

<a name="IUniswapV2Router01-quote-uint256-uint256-uint256-"></a>

```
getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) → uint256 amountOut (external function)
```

<a name="IUniswapV2Router01-getAmountOut-uint256-uint256-uint256-"></a>

```
getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) → uint256 amountIn (external function)
```

<a name="IUniswapV2Router01-getAmountIn-uint256-uint256-uint256-"></a>

```
getAmountsOut(uint256 amountIn, address[] path) → uint256[] amounts (external function)
```

<a name="IUniswapV2Router01-getAmountsOut-uint256-address---"></a>

```
getAmountsIn(uint256 amountOut, address[] path) → uint256[] amounts (external function)
```

<a name="IUniswapV2Router01-getAmountsIn-uint256-address---"></a>

