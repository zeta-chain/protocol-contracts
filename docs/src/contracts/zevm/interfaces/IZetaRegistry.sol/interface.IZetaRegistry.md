# IZetaRegistry
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/main/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/v2/contracts/zevm/interfaces/IZetaRegistry.sol)

**Inherits:**
[IZetaRegistryEvents](/contracts/zevm/interfaces/IZetaRegistry.sol/interface.IZetaRegistryEvents.md), [IZetaRegistryErrors](/contracts/zevm/interfaces/IZetaRegistry.sol/interface.IZetaRegistryErrors.md)

Interface for the ZetaChain Registry contract which serves as a central
repository for all contract addresses across connected blockchains.


## Functions
### addChain

Adds a new blockchain to the registry.


```solidity
function addChain(bytes32 chainIdentifier, ChainInfo calldata chainInfo) external returns (bool success);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainIdentifier`|`bytes32`|Unique identifier for the blockchain (e.g., bytes32("eth_mainnet")).|
|`chainInfo`|`ChainInfo`|Blockchain information (see ChainInfo struct).|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`success`|`bool`|True if the operation was successful.|


### updateChain

Updates an existing blockchain's information.


```solidity
function updateChain(bytes32 chainIdentifier, ChainInfo calldata chainInfo) external returns (bool success);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainIdentifier`|`bytes32`|Unique identifier for the blockchain.|
|`chainInfo`|`ChainInfo`|Updated blockchain information.|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`success`|`bool`|True if the operation was successful.|


### setChainActive

Sets a blockchain's active status.


```solidity
function setChainActive(bytes32 chainIdentifier, bool active) external returns (bool success);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainIdentifier`|`bytes32`|Unique identifier for the blockchain.|
|`active`|`bool`|New active status.|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`success`|`bool`|True if the operation was successful.|


### removeChain

Removes a blockchain from the registry.


```solidity
function removeChain(bytes32 chainIdentifier) external returns (bool success);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainIdentifier`|`bytes32`|Unique identifier for the blockchain.|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`success`|`bool`|True if the operation was successful.|


### getChainInfo

Retrieves information about a blockchain.


```solidity
function getChainInfo(bytes32 chainIdentifier) external view returns (ChainInfo memory chainInfo);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainIdentifier`|`bytes32`|Unique identifier for the blockchain.|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`chainInfo`|`ChainInfo`|The blockchain information.|


### chainExists

Checks if a blockchain exists in the registry.


```solidity
function chainExists(bytes32 chainIdentifier) external view returns (bool exists);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainIdentifier`|`bytes32`|Unique identifier for the blockchain.|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`exists`|`bool`|True if the blockchain exists.|


### getAllChainIdentifiers

Gets all blockchain identifiers in the registry.


```solidity
function getAllChainIdentifiers() external view returns (bytes32[] memory chainIdentifiers);
```
**Returns**

|Name|Type|Description|
|----|----|-----------|
|`chainIdentifiers`|`bytes32[]`|Array of all blockchain identifiers.|


### getActiveChainIdentifiers

Gets all active blockchain identifiers in the registry.


```solidity
function getActiveChainIdentifiers() external view returns (bytes32[] memory activeChainIdentifiers);
```
**Returns**

|Name|Type|Description|
|----|----|-----------|
|`activeChainIdentifiers`|`bytes32[]`|Array of active blockchain identifiers.|


### getChainCount

Gets the total number of blockchains in the registry.


```solidity
function getChainCount() external view returns (uint256 count);
```
**Returns**

|Name|Type|Description|
|----|----|-----------|
|`count`|`uint256`|Total number of blockchains.|


### addContract

Adds a new contract to a blockchain.


```solidity
function addContract(
    bytes32 chainIdentifier,
    bytes32 contractIdentifier,
    ContractCategory category,
    ContractInfo calldata contractInfo
)
    external
    returns (bool success);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainIdentifier`|`bytes32`|Unique identifier for the blockchain.|
|`contractIdentifier`|`bytes32`|Unique identifier for the contract (e.g., bytes32("ZetaConnector")).|
|`category`|`ContractCategory`|Contract category.|
|`contractInfo`|`ContractInfo`|Contract information (see ContractInfo struct).|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`success`|`bool`|True if the operation was successful.|


### updateContract

Updates an existing contract's information.


```solidity
function updateContract(
    bytes32 chainIdentifier,
    bytes32 contractIdentifier,
    ContractInfo calldata contractInfo
)
    external
    returns (bool success);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainIdentifier`|`bytes32`|Unique identifier for the blockchain.|
|`contractIdentifier`|`bytes32`|Unique identifier for the contract.|
|`contractInfo`|`ContractInfo`|Updated contract information.|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`success`|`bool`|True if the operation was successful.|


### updateContractAddress

Updates only a contract's address.


```solidity
function updateContractAddress(
    bytes32 chainIdentifier,
    bytes32 contractIdentifier,
    address newAddress
)
    external
    returns (bool success);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainIdentifier`|`bytes32`|Unique identifier for the blockchain.|
|`contractIdentifier`|`bytes32`|Unique identifier for the contract.|
|`newAddress`|`address`|New contract address.|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`success`|`bool`|True if the operation was successful.|


### updateContractConfiguration

Updates a contract's configuration data.


```solidity
function updateContractConfiguration(
    bytes32 chainIdentifier,
    bytes32 contractIdentifier,
    bytes calldata configurationData
)
    external
    returns (bool success);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainIdentifier`|`bytes32`|Unique identifier for the blockchain.|
|`contractIdentifier`|`bytes32`|Unique identifier for the contract.|
|`configurationData`|`bytes`|New configuration data (encoded).|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`success`|`bool`|True if the operation was successful.|


### setContractActive

Sets a contract's active status.


```solidity
function setContractActive(
    bytes32 chainIdentifier,
    bytes32 contractIdentifier,
    bool active
)
    external
    returns (bool success);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainIdentifier`|`bytes32`|Unique identifier for the blockchain.|
|`contractIdentifier`|`bytes32`|Unique identifier for the contract.|
|`active`|`bool`|New active status.|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`success`|`bool`|True if the operation was successful.|


### removeContract

Removes a contract from a blockchain.


```solidity
function removeContract(bytes32 chainIdentifier, bytes32 contractIdentifier) external returns (bool success);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainIdentifier`|`bytes32`|Unique identifier for the blockchain.|
|`contractIdentifier`|`bytes32`|Unique identifier for the contract.|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`success`|`bool`|True if the operation was successful.|


### getContractInfo

Gets information about a contract.


```solidity
function getContractInfo(
    bytes32 chainIdentifier,
    bytes32 contractIdentifier
)
    external
    view
    returns (ContractInfo memory contractInfo);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainIdentifier`|`bytes32`|Unique identifier for the blockchain.|
|`contractIdentifier`|`bytes32`|Unique identifier for the contract.|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`contractInfo`|`ContractInfo`|The contract information.|


### getContractAddress

Gets a contract's address.


```solidity
function getContractAddress(bytes32 chainIdentifier, bytes32 contractIdentifier) external view returns (address addr);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainIdentifier`|`bytes32`|Unique identifier for the blockchain.|
|`contractIdentifier`|`bytes32`|Unique identifier for the contract.|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`addr`|`address`|The contract address.|


### getContractIdentifiers

Gets all contract identifiers for a blockchain.


```solidity
function getContractIdentifiers(bytes32 chainIdentifier) external view returns (bytes32[] memory contractIdentifiers);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainIdentifier`|`bytes32`|Unique identifier for the blockchain.|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`contractIdentifiers`|`bytes32[]`|Array of all contract identifiers for the blockchain.|


### getContractIdentifiersByCategory

Gets all contract identifiers for a blockchain by category.


```solidity
function getContractIdentifiersByCategory(
    bytes32 chainIdentifier,
    ContractCategory category
)
    external
    view
    returns (bytes32[] memory contractIdentifiers);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainIdentifier`|`bytes32`|Unique identifier for the blockchain.|
|`category`|`ContractCategory`|Contract category.|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`contractIdentifiers`|`bytes32[]`|Array of contract identifiers in the specified category.|


### contractExists

Checks if a contract exists in the registry.


```solidity
function contractExists(bytes32 chainIdentifier, bytes32 contractIdentifier) external view returns (bool exists);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainIdentifier`|`bytes32`|Unique identifier for the blockchain.|
|`contractIdentifier`|`bytes32`|Unique identifier for the contract.|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`exists`|`bool`|True if the contract exists.|


### createContractIdentifier

Creates a contract identifier from a string.


```solidity
function createContractIdentifier(string calldata name) external pure returns (bytes32 identifier);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`name`|`string`|String name for the contract.|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`identifier`|`bytes32`|bytes32 identifier.|


### createChainIdentifier

Creates a chain identifier from a string.


```solidity
function createChainIdentifier(string calldata name) external pure returns (bytes32 identifier);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`name`|`string`|String name for the chain.|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`identifier`|`bytes32`|bytes32 identifier.|


### getSystemAddresses

Helper to get common system contract addresses for a blockchain.


```solidity
function getSystemAddresses(bytes32 chainIdentifier)
    external
    view
    returns (address connector, address erc20Custody, address tss, address tssUpdater, address zetaToken);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainIdentifier`|`bytes32`|Unique identifier for the blockchain.|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`connector`|`address`|The connector contract address.|
|`erc20Custody`|`address`|The ERC20 custody contract address.|
|`tss`|`address`|The TSS contract address.|
|`tssUpdater`|`address`|The TSS updater contract address.|
|`zetaToken`|`address`|The ZETA token contract address.|


### getZRC20Address

Helper to get ZRC-20 token address for a given chain and token.


```solidity
function getZRC20Address(
    bytes32 chainIdentifier,
    bytes32 tokenIdentifier
)
    external
    view
    returns (address tokenAddress);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainIdentifier`|`bytes32`|Unique identifier for the blockchain.|
|`tokenIdentifier`|`bytes32`|Unique identifier for the token (e.g., bytes32("USDC")).|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`tokenAddress`|`address`|The ZRC-20 token address.|


### addZRC20Token

Adds a new ZRC-20 token to the registry.


```solidity
function addZRC20Token(
    bytes32 chainIdentifier,
    bytes32 tokenIdentifier,
    address tokenAddress
)
    external
    returns (bool success);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainIdentifier`|`bytes32`|Unique identifier for the blockchain.|
|`tokenIdentifier`|`bytes32`|Unique identifier for the token.|
|`tokenAddress`|`address`|Address of the ZRC-20 token.|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`success`|`bool`|True if the operation was successful.|


### batchAddContracts

Batch add contracts to a blockchain.


```solidity
function batchAddContracts(
    bytes32 chainIdentifier,
    bytes32[] calldata contractIdentifiers,
    ContractCategory[] calldata categories,
    ContractInfo[] calldata contractInfos
)
    external
    returns (bool success);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainIdentifier`|`bytes32`|Unique identifier for the blockchain.|
|`contractIdentifiers`|`bytes32[]`|Array of contract identifiers.|
|`categories`|`ContractCategory[]`|Array of contract categories.|
|`contractInfos`|`ContractInfo[]`|Array of contract information.|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`success`|`bool`|True if the operation was successful.|


## Structs
### ChainInfo
Struct representing core blockchain information.


```solidity
struct ChainInfo {
    uint256 chainId;
    string name;
    bool active;
    uint8 chainType;
    uint64 blockConfirmations;
}
```

**Properties**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The unique identifier for this blockchain.|
|`name`|`string`|Human-readable name of the blockchain.|
|`active`|`bool`|Whether this blockchain is currently active in the ZetaChain network.|
|`chainType`|`uint8`|Type of blockchain (EVM, Cosmos, etc.).|
|`blockConfirmations`|`uint64`|Number of block confirmations required for finality.|

### ContractInfo
Struct representing a system contract's configuration data.


```solidity
struct ContractInfo {
    bool active;
    address addr;
    string version;
    address implementation;
    bytes configurationData;
}
```

**Properties**

|Name|Type|Description|
|----|----|-----------|
|`active`|`bool`|Whether this contract is currently active.|
|`addr`|`address`|The deployed address of the contract.|
|`version`|`string`|Semantic version of the contract.|
|`implementation`|`address`|For proxy contracts, the address of the implementation.|
|`configurationData`|`bytes`|Additional contract-specific configuration data (encoded).|

## Enums
### ContractCategory
Enum representing different system contract categories.


```solidity
enum ContractCategory {
    ZRC20,
    GATEWAY,
    GOVERNANCE,
    CONNECTOR,
    CUSTODY,
    OTHER
}
```

