# ICoreRegistry
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/main/v2/v2/v2/contracts/zevm/interfaces/ICoreRegistry.sol)

**Inherits:**
[ICoreRegistryErrors](/contracts/zevm/interfaces/ICoreRegistry.sol/interface.ICoreRegistryErrors.md), [ICoreRegistryEvents](/contracts/zevm/interfaces/ICoreRegistry.sol/interface.ICoreRegistryEvents.md)

Primary interface for the CoreRegistry contract, the central registry for ZetaChain ecosystem.


## Functions
### changeChainStatus

Changes status of the chain to activated/deactivated.


```solidity
function changeChainStatus(uint256 chainId, address gasZRC20, bytes calldata registry, bool activation) external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain to activate.|
|`gasZRC20`|`address`|The address of the ZRC20 token that represents gas token for the chain.|
|`registry`|`bytes`||
|`activation`|`bool`|Whether activate or deactivate a chain|


### updateChainMetadata

Updates chain metadata.


```solidity
function updateChainMetadata(uint256 chainId, string calldata key, bytes calldata value) external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain.|
|`key`|`string`|The metadata key to update.|
|`value`|`bytes`|The new value for the metadata.|


### registerContract

Registers a new contract address for a specific chain.


```solidity
function registerContract(uint256 chainId, string calldata contractType, bytes calldata addressBytes) external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain where the contract is deployed.|
|`contractType`|`string`|The type of the contract (e.g., "connector", "gateway").|
|`addressBytes`|`bytes`|The bytes representation of the non-EVM address.|


### updateContractConfiguration

Updates contract configuration.


```solidity
function updateContractConfiguration(
    uint256 chainId,
    string calldata contractType,
    string calldata key,
    bytes calldata value
)
    external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain where the contract is deployed.|
|`contractType`|`string`|The type of the contract.|
|`key`|`string`|The configuration key to update.|
|`value`|`bytes`|The new value for the configuration.|


### setContractActive


```solidity
function setContractActive(uint256 chainId, string calldata contractType, bool active) external;
```

### registerZRC20Token

Registers a new ZRC20 token in the registry.


```solidity
function registerZRC20Token(
    address address_,
    string calldata symbol,
    uint256 originChainId,
    bytes calldata originAddress,
    string calldata coinType,
    uint8 decimals
)
    external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`address_`|`address`|The address of the ZRC20 token on ZetaChain.|
|`symbol`|`string`|The symbol of the token.|
|`originChainId`|`uint256`|The ID of the foreign chain where the original asset exists.|
|`originAddress`|`bytes`|The address or identifier of the asset on its native chain.|
|`coinType`|`string`||
|`decimals`|`uint8`||


### setZRC20TokenActive

Updates ZRC20 token information.


```solidity
function setZRC20TokenActive(address address_, bool active) external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`address_`|`address`|The address of the ZRC20 token.|
|`active`|`bool`|Whether the token should be active.|


### getChainMetadata

Gets chain-specific metadata.


```solidity
function getChainMetadata(uint256 chainId, string calldata key) external view returns (bytes memory);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain.|
|`key`|`string`|The metadata key to retrieve.|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`bytes`|The value of the requested metadata.|


### getContractInfo

Gets information about a specific contract.


```solidity
function getContractInfo(
    uint256 chainId,
    string calldata contractType
)
    external
    view
    returns (bool active, bytes memory addressBytes);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain where the contract is deployed.|
|`contractType`|`string`|The type of the contract.|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`active`|`bool`|Whether the contract is active.|
|`addressBytes`|`bytes`|The address of the contract.|


### getContractConfiguration

Gets contract-specific configuration.


```solidity
function getContractConfiguration(
    uint256 chainId,
    string calldata contractType,
    string calldata key
)
    external
    view
    returns (bytes memory);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain where the contract is deployed.|
|`contractType`|`string`|The type of the contract.|
|`key`|`string`|The configuration key to retrieve.|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`bytes`|The value of the requested configuration.|


### getZRC20TokenInfo

Gets information about a specific ZRC20 token.


```solidity
function getZRC20TokenInfo(address address_)
    external
    view
    returns (
        bool active,
        string memory symbol,
        uint256 originChainId,
        bytes memory originAddress,
        string memory coinType,
        uint8 decimals
    );
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`address_`|`address`|The address of the ZRC20 token.|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`active`|`bool`|Whether the token is active.|
|`symbol`|`string`|The symbol of the token|
|`originChainId`|`uint256`|The ID of the foreign chain where the original asset exists.|
|`originAddress`|`bytes`|The address or identifier of the asset on its native chain.|
|`coinType`|`string`|The type of the original coin.|
|`decimals`|`uint8`|The number of decimals the token uses.|


### getZRC20AddressByForeignAsset

Gets the ZRC20 token address for a specific asset on a foreign chain.


```solidity
function getZRC20AddressByForeignAsset(
    uint256 originChainId,
    bytes calldata originAddress
)
    external
    view
    returns (address);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`originChainId`|`uint256`|The ID of the foreign chain.|
|`originAddress`|`bytes`|The address or identifier of the asset on its native chain.|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`address`|The address of the corresponding ZRC20 token on ZetaChain.|


### getActiveChains

Gets all active chains in the registry.


```solidity
function getActiveChains() external view returns (uint256[] memory);
```
**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`uint256[]`|Array of chain IDs for all active chains.|


