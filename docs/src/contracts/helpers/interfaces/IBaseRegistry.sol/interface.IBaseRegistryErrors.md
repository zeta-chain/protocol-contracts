# IBaseRegistryErrors
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/main/v2/contracts/helpers/interfaces/IBaseRegistry.sol)

Interface for the errors used by the BaseRegistry contract.


## Errors
### ZeroAddress
Error thrown when a zero address is provided where a non-zero address is required.


```solidity
error ZeroAddress();
```

### InvalidSender
Error thrown when the sender is invalid


```solidity
error InvalidSender();
```

### TransferFailed
Error thrown when a ZRC20 token transfer failed.


```solidity
error TransferFailed();
```

### ChainActive
Error thrown when a chain is already active.


```solidity
error ChainActive(uint256 chainId);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain that is already active.|

### ChainNonActive
Error thrown when a chain is not active.


```solidity
error ChainNonActive(uint256 chainId);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain that is not active.|

### InvalidContractType
Error thrown when a contract type is invalid.


```solidity
error InvalidContractType(string message);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`message`|`string`|Describes why error happened|

### ContractAlreadyRegistered
Error thrown when a contract is already registered.


```solidity
error ContractAlreadyRegistered(uint256 chainId, string contractType, bytes addressBytes);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain.|
|`contractType`|`string`|The type of the contract.|
|`addressBytes`|`bytes`|The address of the contract.|

### ContractNotFound
Error thrown when a contract is not found in the registry.


```solidity
error ContractNotFound(uint256 chainId, string contractType);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain,|
|`contractType`|`string`|The type of the contract.|

### ZRC20AlreadyRegistered
Error thrown when a ZRC20 token is already registered.


```solidity
error ZRC20AlreadyRegistered(address address_);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`address_`|`address`|The address of the ZRC20 token.|

### ZRC20SymbolAlreadyInUse
Error thrown when a ZRC20 token symbol is already in use.


```solidity
error ZRC20SymbolAlreadyInUse(string symbol);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`symbol`|`string`|The symbol that is already in use.|

