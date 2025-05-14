# IZetaRegistryEvents
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/main/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/contracts/zevm/interfaces/IZetaRegistry.sol)

Interface for the events emitted by the ZetaRegistry contract.


## Events
### ChainAdded
Emitted when a new blockchain is added to the registry.


```solidity
event ChainAdded(bytes32 indexed chainIdentifier, uint256 chainId, string name);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainIdentifier`|`bytes32`|Unique identifier for the blockchain.|
|`chainId`|`uint256`|Blockchain network ID.|
|`name`|`string`|Human-readable name of the blockchain.|

### ChainStatusUpdated
Emitted when a blockchain's active status is updated.


```solidity
event ChainStatusUpdated(bytes32 indexed chainIdentifier, bool active);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainIdentifier`|`bytes32`|Unique identifier for the blockchain.|
|`active`|`bool`|New active status.|

### ContractAddressUpdated
Emitted when a contract address is updated.


```solidity
event ContractAddressUpdated(
    bytes32 indexed chainIdentifier, bytes32 indexed contractIdentifier, uint8 indexed category, address contractAddress
);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainIdentifier`|`bytes32`|Unique identifier for the blockchain.|
|`contractIdentifier`|`bytes32`|Unique identifier for the contract.|
|`category`|`uint8`|Contract category.|
|`contractAddress`|`address`|New contract address.|

### ContractConfigurationUpdated
Emitted when a contract's configuration is updated.


```solidity
event ContractConfigurationUpdated(
    bytes32 indexed chainIdentifier, bytes32 indexed contractIdentifier, bytes configurationData
);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainIdentifier`|`bytes32`|Unique identifier for the blockchain.|
|`contractIdentifier`|`bytes32`|Unique identifier for the contract.|
|`configurationData`|`bytes`|New configuration data (encoded).|

### ContractStatusUpdated
Emitted when a contract's active status is updated.


```solidity
event ContractStatusUpdated(bytes32 indexed chainIdentifier, bytes32 indexed contractIdentifier, bool active);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainIdentifier`|`bytes32`|Unique identifier for the blockchain.|
|`contractIdentifier`|`bytes32`|Unique identifier for the contract.|
|`active`|`bool`|New active status.|

### ZRC20TokenAdded
Emitted when a ZRC20 token is added to the registry.


```solidity
event ZRC20TokenAdded(bytes32 indexed chainIdentifier, bytes32 indexed tokenIdentifier, address tokenAddress);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainIdentifier`|`bytes32`|Unique identifier for the blockchain.|
|`tokenIdentifier`|`bytes32`|Unique identifier for the token.|
|`tokenAddress`|`address`|Address of the ZRC20 token.|

