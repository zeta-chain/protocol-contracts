## IUniswapV2Router02

```solidity
import "@zetachain/protocol-contracts/contracts/zevm/interfaces/IUniswapV2Router02.sol";
```

Source: https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/zevm/interfaces/IUniswapV2Router02.sol

### Function List

* [removeLiquidityETHSupportingFeeOnTransferTokens(token, liquidity, amountTokenMin, amountETHMin, to, deadline)](#IUniswapV2Router02-removeLiquidityETHSupportingFeeOnTransferTokens-address-uint256-uint256-uint256-address-uint256-)
* [removeLiquidityETHWithPermitSupportingFeeOnTransferTokens(token, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)](#IUniswapV2Router02-removeLiquidityETHWithPermitSupportingFeeOnTransferTokens-address-uint256-uint256-uint256-address-uint256-bool-uint8-bytes32-bytes32-)
* [swapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn, amountOutMin, path, to, deadline)](#IUniswapV2Router02-swapExactTokensForTokensSupportingFeeOnTransferTokens-uint256-uint256-address---address-uint256-)
* [swapExactETHForTokensSupportingFeeOnTransferTokens(amountOutMin, path, to, deadline)](#IUniswapV2Router02-swapExactETHForTokensSupportingFeeOnTransferTokens-uint256-address---address-uint256-)
* [swapExactTokensForETHSupportingFeeOnTransferTokens(amountIn, amountOutMin, path, to, deadline)](#IUniswapV2Router02-swapExactTokensForETHSupportingFeeOnTransferTokens-uint256-uint256-address---address-uint256-)

### Modifiers

### Functions

```
removeLiquidityETHSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) → uint256 amountETH (external function)
```

<a name="IUniswapV2Router02-removeLiquidityETHSupportingFeeOnTransferTokens-address-uint256-uint256-uint256-address-uint256-"></a>

```
removeLiquidityETHWithPermitSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) → uint256 amountETH (external function)
```

<a name="IUniswapV2Router02-removeLiquidityETHWithPermitSupportingFeeOnTransferTokens-address-uint256-uint256-uint256-address-uint256-bool-uint8-bytes32-bytes32-"></a>

```
swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) (external function)
```

<a name="IUniswapV2Router02-swapExactTokensForTokensSupportingFeeOnTransferTokens-uint256-uint256-address---address-uint256-"></a>

```
swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) (external function)
```

<a name="IUniswapV2Router02-swapExactETHForTokensSupportingFeeOnTransferTokens-uint256-address---address-uint256-"></a>

```
swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) (external function)
```

<a name="IUniswapV2Router02-swapExactTokensForETHSupportingFeeOnTransferTokens-uint256-uint256-address---address-uint256-"></a>

