# IBaseRegistryEvents
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/helpers/interfaces/IBaseRegistry.sol)

Interface for the events emitted by the BaseRegistry contract.


## Events
### ChainStatusChanged
Emitted when a chain status has changed.


```solidity
event ChainStatusChanged(uint256 indexed chainId, bool newStatus);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain.|
|`newStatus`|`bool`|The new chain status (is active or not).|

### ChainMetadataUpdated
Emitted when a chain metadata is set.


```solidity
event ChainMetadataUpdated(uint256 indexed chainId, string key, bytes value);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain.|
|`key`|`string`|The metadata key to update.|
|`value`|`bytes`|The new value for the metadata.|

### ContractRegistered
Emitted when a new contract is registered.


```solidity
event ContractRegistered(uint256 indexed chainId, string indexed contractType, bytes addressBytes);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain where the contract is deployed.|
|`contractType`|`string`|The type of the contract (e.g. "connector", "gateway", "tss").|
|`addressBytes`|`bytes`|The contract address in bytes representation.|

### ContractStatusChanged
Emitted when a contract status has changed.


```solidity
event ContractStatusChanged(bytes addressBytes);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`addressBytes`|`bytes`|The contract address in bytes representation.|

### ContractConfigurationUpdated
Emitted when a contract configuration is updated.


```solidity
event ContractConfigurationUpdated(uint256 indexed chainId, string contractType, string key, bytes value);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain where the contract is deployed.|
|`contractType`|`string`|The type of the contract.|
|`key`|`string`|The configuration key to update.|
|`value`|`bytes`|The new value for the configuration.|

### ZRC20TokenRegistered
Emitted when a ZRC20 token is registered.


```solidity
event ZRC20TokenRegistered(
    bytes indexed originAddress, address indexed address_, uint8 decimals, uint256 originChainId, string symbol
);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`originAddress`|`bytes`|The address of the asset on its native chain.|
|`address_`|`address`|The address of the ZRC20 token on ZetaChain.|
|`decimals`|`uint8`|The number of decimals the token uses.|
|`originChainId`|`uint256`|The ID of the foreign chain where the original asset exists.|
|`symbol`|`string`|The symbol of the token.|

### ZRC20TokenUpdated
Emitted when a ZRC20 token is updated.


```solidity
event ZRC20TokenUpdated(address address_, bool active);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`address_`|`address`|The address of the ZRC20 token.|
|`active`|`bool`|Whether the token should be active.|

