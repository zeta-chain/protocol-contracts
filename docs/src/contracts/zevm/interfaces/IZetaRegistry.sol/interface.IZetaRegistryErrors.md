# IZetaRegistryErrors
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/main/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/contracts/zevm/interfaces/IZetaRegistry.sol)

Interface for the errors used in the ZetaRegistry contract.


## Errors
### ChainAlreadyExists
Error indicating the chain already exists.


```solidity
error ChainAlreadyExists(bytes32 chainIdentifier);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainIdentifier`|`bytes32`|Unique identifier for the blockchain.|

### ChainDoesNotExist
Error indicating the chain does not exist.


```solidity
error ChainDoesNotExist(bytes32 chainIdentifier);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainIdentifier`|`bytes32`|Unique identifier for the blockchain.|

### ContractAlreadyExists
Error indicating the contract already exists.


```solidity
error ContractAlreadyExists(bytes32 chainIdentifier, bytes32 contractIdentifier);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainIdentifier`|`bytes32`|Unique identifier for the blockchain.|
|`contractIdentifier`|`bytes32`|Unique identifier for the contract.|

### ContractDoesNotExist
Error indicating the contract does not exist.


```solidity
error ContractDoesNotExist(bytes32 chainIdentifier, bytes32 contractIdentifier);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainIdentifier`|`bytes32`|Unique identifier for the blockchain.|
|`contractIdentifier`|`bytes32`|Unique identifier for the contract.|

### InvalidContractAddress
Error indicating an invalid contract address.


```solidity
error InvalidContractAddress(address contractAddress);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`contractAddress`|`address`|The invalid contract address.|

### ZRC20TokenAlreadyExists
Error indicating the ZRC20 token already exists.


```solidity
error ZRC20TokenAlreadyExists(bytes32 chainIdentifier, bytes32 tokenIdentifier);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainIdentifier`|`bytes32`|Unique identifier for the blockchain.|
|`tokenIdentifier`|`bytes32`|Unique identifier for the token.|

### ZRC20TokenDoesNotExist
Error indicating the ZRC20 token does not exist.


```solidity
error ZRC20TokenDoesNotExist(bytes32 chainIdentifier, bytes32 tokenIdentifier);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainIdentifier`|`bytes32`|Unique identifier for the blockchain.|
|`tokenIdentifier`|`bytes32`|Unique identifier for the token.|

### UnauthorizedCaller
Error indicating an unauthorized caller.


```solidity
error UnauthorizedCaller(address caller);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`caller`|`address`|The address that attempted to call the function.|

### ArrayLengthMismatch
Error indicating array length mismatch.


```solidity
error ArrayLengthMismatch(uint256 expected, uint256 actual);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`expected`|`uint256`|Expected length.|
|`actual`|`uint256`|Actual length.|

