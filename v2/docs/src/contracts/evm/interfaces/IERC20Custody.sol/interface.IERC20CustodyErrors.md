# IERC20CustodyErrors
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/1ebf02353d4ffe1e3d159fe4887220a0672a2035/contracts/evm/interfaces/IERC20Custody.sol)

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

