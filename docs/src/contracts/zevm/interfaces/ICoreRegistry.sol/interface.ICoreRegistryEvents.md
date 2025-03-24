# ICoreRegistryEvents
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/main/v2/contracts/zevm/interfaces/ICoreRegistry.sol)

Interface for the events emitted by the CoreRegistry contract.


## Events
### ContractRegistered
Emitted when a new contract address is registered.


```solidity
event ContractRegistered(uint256 indexed chainId, string indexed contractType, string addressString);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain where the contract is deployed.|
|`contractType`|`string`|The type of the contract (e.g. "connector", "tss", "gateway")|
|`addressString`|`string`|The string representation of the registered address|

### ZRC20TokenRegistered
Emitted when a ZRC20 token is registered.


```solidity
event ZRC20TokenRegistered(
    string indexed originAddress, address indexed address_, uint8 decimals, uint256 originChainId, string symbol
);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`originAddress`|`string`|The address of the asset on its native chain.|
|`address_`|`address`|The address of the ZRC20 token on ZetaChain.|
|`decimals`|`uint8`|The number of decimals the token uses.|
|`originChainId`|`uint256`|The ID of the foreign chain where the original asset exists.|
|`symbol`|`string`|The symbol of the token.|

### ChainStatusChanged
Emitted when a chain status has changed


```solidity
event ChainStatusChanged(uint256 indexed chainId);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain.|

### ContractStatusChanged
Emitted when a contract status has changed


```solidity
event ContractStatusChanged(string addressString);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`addressString`|`string`|The string representation of the registered address.|

### NewChainMetadata
Emitted when a chain metadata is set


```solidity
event NewChainMetadata(uint256 indexed chainId, string key, bytes value);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain.|
|`key`|`string`|The metadata key to update.|
|`value`|`bytes`|The new value for the metadata.|

### NewContractConfiguration
Emitted when a new contract configuration is updated


```solidity
event NewContractConfiguration(uint256 indexed chainId, string contractType, string key, bytes value);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain where the contract is deployed.|
|`contractType`|`string`|The type of the contract.|
|`key`|`string`|The configuration key to update.|
|`value`|`bytes`|The new value for the configuration.|

### UpdatedRegistryManager
Emitted when registry manager address is updated


```solidity
event UpdatedRegistryManager(address oldRegistryManager, address newRegistryManager);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`oldRegistryManager`|`address`|old registry manager address|
|`newRegistryManager`|`address`|new registry manager address|

