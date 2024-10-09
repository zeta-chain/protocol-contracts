# IERC20CustodyErrors
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/0f5a289f3179440ee2e4f23b1aa3613d2e644a78/contracts/evm/interfaces/IERC20Custody.sol)

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

