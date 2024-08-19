# IERC20CustodyErrors
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/639b8fd295b6e67d3a12f04f00a74c5a31cff469/contracts/evm/interfaces/IERC20Custody.sol)

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

