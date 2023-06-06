# evm/tools/interfaces/TridentIPoolRouter.md

## IPoolRouter

### ExactInputSingleParams

```solidity
struct ExactInputSingleParams {
  address tokenIn;
  uint256 amountIn;
  uint256 amountOutMinimum;
  address pool;
  address to;
  bool unwrap;
}
```

### ExactInputParams

```solidity
struct ExactInputParams {
  address tokenIn;
  uint256 amountIn;
  uint256 amountOutMinimum;
  address[] path;
  address to;
  bool unwrap;
}
```

### ExactOutputSingleParams

```solidity
struct ExactOutputSingleParams {
  address tokenIn;
  uint256 amountOut;
  uint256 amountInMaximum;
  address pool;
  address to;
  bool unwrap;
}
```

### ExactOutputParams

```solidity
struct ExactOutputParams {
  address tokenIn;
  uint256 amountOut;
  uint256 amountInMaximum;
  address[] path;
  address to;
  bool unwrap;
}
```

### exactInputSingle

```solidity
function exactInputSingle(struct IPoolRouter.ExactInputSingleParams params) external payable returns (uint256 amountOut)
```

Swap amountIn of one token for as much as possible of another token

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| params | struct IPoolRouter.ExactInputSingleParams | The parameters necessary for the swap, encoded as ExactInputSingleParams in calldata |

### exactInput

```solidity
function exactInput(struct IPoolRouter.ExactInputParams params) external payable returns (uint256 amountOut)
```

Swap amountIn of one token for as much as possible of another along the specified path

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| params | struct IPoolRouter.ExactInputParams | The parameters necessary for the multi-hop swap, encoded as ExactInputParams in calldata |

### exactOutputSingle

```solidity
function exactOutputSingle(struct IPoolRouter.ExactOutputSingleParams params) external payable returns (uint256 amountIn)
```

Swaps as little as possible of one token for `amountOut` of another token

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| params | struct IPoolRouter.ExactOutputSingleParams | The parameters necessary for the swap, encoded as ExactOutputSingleParams in calldata |

### exactOutput

```solidity
function exactOutput(struct IPoolRouter.ExactOutputParams params) external payable returns (uint256 amountIn)
```

Swaps as little as possible of one token for `amountOut` of another along the specified path (reversed)

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| params | struct IPoolRouter.ExactOutputParams | The parameters necessary for the multi-hop swap, encoded as ExactOutputParams in calldata |

### sweep

```solidity
function sweep(address token, uint256 amount, address recipient) external payable
```

Recover mistakenly sent tokens

