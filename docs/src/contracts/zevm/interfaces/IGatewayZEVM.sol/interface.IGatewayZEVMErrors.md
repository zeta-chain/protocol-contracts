# IGatewayZEVMErrors
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/main/v2/contracts/zevm/interfaces/IGatewayZEVM.sol)

Interface for the errors used in the GatewayZEVM contract.


## Errors
### WithdrawalFailed
Error indicating a withdrawal failure.


```solidity
error WithdrawalFailed(address token, address recipient, uint256 amount);
```

### InsufficientZRC20Amount
Error indicating an insufficient ZRC20 token amount.


```solidity
error InsufficientZRC20Amount();
```

### InsufficientZetaAmount
Error indicating an insufficient zeta amount.


```solidity
error InsufficientZetaAmount();
```

### ZRC20BurnFailed
Error indicating a failure to burn ZRC20 tokens.


```solidity
error ZRC20BurnFailed(address zrc20, uint256 amount);
```

### ZRC20TransferFailed
Error indicating a failure to transfer ZRC20 tokens.


```solidity
error ZRC20TransferFailed(address zrc20, address from, address to, uint256 amount);
```

### ZRC20DepositFailed
Error indicating a failure to deposit ZRC20 tokens.


```solidity
error ZRC20DepositFailed(address zrc20, address to, uint256 amount);
```

### GasFeeTransferFailed
Error indicating a failure to transfer gas fee.


```solidity
error GasFeeTransferFailed(address token, address to, uint256 amount);
```

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

