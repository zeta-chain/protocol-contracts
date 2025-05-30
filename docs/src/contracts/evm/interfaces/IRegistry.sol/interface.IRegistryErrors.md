# IRegistryErrors
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/evm/interfaces/IRegistry.sol)


## Errors
### NotGateway
Error thrown when an operation requires the caller to be the gateway


```solidity
error NotGateway(address caller);
```

### ChainActive
Error thrown when the chain ID is already active


```solidity
error ChainActive(uint256 chainId);
```

### ChainNonActive
Error thrown when the chain ID is not active


```solidity
error ChainNonActive(uint256 chainId);
```

### ZeroAddress
Error thrown when a zero address is provided


```solidity
error ZeroAddress();
```

### InvalidContractType
Error thrown when an invalid contract type is provided


```solidity
error InvalidContractType(string contractType);
```

### ContractAlreadyRegistered
Error thrown when a contract is already registered


```solidity
error ContractAlreadyRegistered(uint256 chainId, string contractType, bytes addressBytes);
```

### ContractNotFound
Error thrown when a contract is not found


```solidity
error ContractNotFound(uint256 chainId, string contractType);
```

### ZRC20AlreadyRegistered
Error thrown when a ZRC20 token is already registered


```solidity
error ZRC20AlreadyRegistered(address zrc20);
```

### ZRC20SymbolAlreadyInUse
Error thrown when a ZRC20 symbol is already in use


```solidity
error ZRC20SymbolAlreadyInUse(string symbol);
```

### InvalidSender
Error thrown when the sender is invalid


```solidity
error InvalidSender();
```

### NotAuthorized
Error thrown when the caller is not authorized


```solidity
error NotAuthorized(address caller);
```

