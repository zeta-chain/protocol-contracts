# IERC20CustodyErrors
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/57d1446373e344f7366de3008e0bf2b97aeeabf8/contracts/evm/interfaces/IERC20Custody.sol)

Interface for the errors used in the ERC20 custody contract.


## Errors
### ZeroAddress
Error for zero address input.


```solidity
error ZeroAddress();
```

### NotWhitelisted
Error for not whitelisted ERC20 token


```solidity
error NotWhitelisted();
```

### LegacyMethodsNotSupported
Error for calling not supported legacy methods.


```solidity
error LegacyMethodsNotSupported();
```

