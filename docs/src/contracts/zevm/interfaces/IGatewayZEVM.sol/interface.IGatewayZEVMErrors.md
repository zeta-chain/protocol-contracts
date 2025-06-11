# IGatewayZEVMErrors
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/zevm/interfaces/IGatewayZEVM.sol)

Interface for the errors used in the GatewayZEVM contract.


## Errors
### WithdrawalFailed
Error indicating a withdrawal failure.


```solidity
error WithdrawalFailed(address token, address recipient, uint256 amount);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`token`|`address`|The address of the token that failed to withdraw.|
|`recipient`|`address`|The address that was supposed to receive the tokens.|
|`amount`|`uint256`|The amount of tokens that failed to withdraw.|

### InsufficientAmount
Error indicating an insufficient token amount.


```solidity
error InsufficientAmount();
```

### ZRC20BurnFailed
Error indicating a failure to burn ZRC20 tokens.


```solidity
error ZRC20BurnFailed(address zrc20, uint256 amount);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`zrc20`|`address`|The address of the ZRC20 token that failed to burn.|
|`amount`|`uint256`|The amount of tokens that failed to burn.|

### ZRC20TransferFailed
Error indicating a failure to transfer ZRC20 tokens.


```solidity
error ZRC20TransferFailed(address zrc20, address from, address to, uint256 amount);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`zrc20`|`address`|The address of the ZRC20 token that failed to transfer.|
|`from`|`address`|The address sending the tokens.|
|`to`|`address`|The address receiving the tokens.|
|`amount`|`uint256`|The amount of tokens that failed to transfer.|

### ZRC20DepositFailed
Error indicating a failure to deposit ZRC20 tokens.


```solidity
error ZRC20DepositFailed(address zrc20, address to, uint256 amount);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`zrc20`|`address`|The address of the ZRC20 token that failed to deposit.|
|`to`|`address`|The address that was supposed to receive the deposit.|
|`amount`|`uint256`|The amount of tokens that failed to deposit.|

### GasFeeTransferFailed
Error indicating a failure to transfer gas fee.


```solidity
error GasFeeTransferFailed(address token, address to, uint256 amount);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`token`|`address`|The address of the token used for gas fee.|
|`to`|`address`|The address that was supposed to receive the gas fee.|
|`amount`|`uint256`|The amount of gas fee that failed to transfer.|

### CallerIsNotProtocol
Error indicating that the caller is not the protocol account.


```solidity
error CallerIsNotProtocol();
```

### InvalidTarget
Error indicating an invalid target address.


```solidity
error InvalidTarget();
```

### FailedZetaSent
Error indicating a failure to send ZETA tokens.


```solidity
error FailedZetaSent(address recipient, uint256 amount);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`recipient`|`address`|The address that was supposed to receive the ZETA tokens.|
|`amount`|`uint256`|The amount of ZETA tokens that failed to send.|

### OnlyWZETAOrProtocol
Error indicating that only WZETA or the protocol address can call the function.


```solidity
error OnlyWZETAOrProtocol();
```

### InsufficientGasLimit
Error indicating an insufficient gas limit.


```solidity
error InsufficientGasLimit();
```

### MessageSizeExceeded
Error indicating message size exceeded in external functions.


```solidity
error MessageSizeExceeded(uint256 provided, uint256 maximum);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`provided`|`uint256`|The size of the message that was provided.|
|`maximum`|`uint256`|The maximum allowed message size.|

### ZeroGasPrice
Error indicating an invalid gas price.


```solidity
error ZeroGasPrice();
```

