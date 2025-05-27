# Registry
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/evm/Registry.sol)

**Inherits:**
[BaseRegistry](/contracts/helpers/BaseRegistry.sol/abstract.BaseRegistry.md), [IRegistry](/contracts/evm/interfaces/IRegistry.sol/interface.IRegistry.md)

Satellite registry contract for connected chains, receiving updates from CoreRegistry.

*This contract is deployed on every connected chain and maintains a synchronized view of the registry.*


## State Variables
### GATEWAY_ROLE
Identifier for the gateway role


```solidity
bytes32 public constant GATEWAY_ROLE = keccak256("GATEWAY_ROLE");
```


### gatewayEVM
GatewayEVM contract that will call this contract with messages from CoreRegistry


```solidity
IGatewayEVM public gatewayEVM;
```


### coreRegistry
Represents the address of the CoreRegistry contract on the ZetaChain


```solidity
address public coreRegistry;
```


## Functions
### onlyRegistry

Restricts function calls to only be made by this contract itself

*Only registry address allowed modifier.*

*This is used to ensure functions receiving cross-chain messages can only be called through
the onCall function using a self-call pattern, preventing direct external calls to these functions*


```solidity
modifier onlyRegistry();
```

### initialize

Initialize the Registry contract


```solidity
function initialize(
    address admin_,
    address pauserAddress_,
    address registryManager_,
    address gatewayEVM_,
    address coreRegistry_
)
    public
    initializer;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`admin_`|`address`|Address with DEFAULT_ADMIN_ROLE, authorized for upgrades and pausing actions|
|`pauserAddress_`|`address`|Address with PAUSER_ROLE, authorized for pausing actions|
|`registryManager_`|`address`|Address with REGISTRY_MANAGER_ROLE, authorized for all registry write actions.|
|`gatewayEVM_`|`address`|Address of the GatewayEVM contract for cross-chain messaging|
|`coreRegistry_`|`address`|Address of the CoreRegistry contract deployed on ZetaChain|


### onCall

onCall is called by the GatewayEVM when a cross-chain message is received


```solidity
function onCall(
    MessageContext calldata context,
    bytes calldata data
)
    external
    onlyRole(GATEWAY_ROLE)
    whenNotPaused
    returns (bytes memory);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`context`|`MessageContext`|Information about the cross-chain message|
|`data`|`bytes`|The encoded function call to execute|


### changeChainStatus

Changes status of the chain to activated/deactivated

*Only callable through onCall from CoreRegistry*


```solidity
function changeChainStatus(
    uint256 chainId,
    address gasZRC20,
    bytes calldata registry,
    bool activation
)
    external
    onlyRegistry
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain being activated/deactivated.|
|`gasZRC20`|`address`|The address of the ZRC20 token that represents gas token for the chain.|
|`registry`|`bytes`|Address of the Registry contract on the connected chain.|
|`activation`|`bool`|Whether activate or deactivate the chain|


### updateChainMetadata

Updates chain metadata, only for the active chains

*Only callable through onCall from CoreRegistry*


```solidity
function updateChainMetadata(
    uint256 chainId,
    string calldata key,
    bytes calldata value
)
    external
    onlyRegistry
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain|
|`key`|`string`|The metadata key to update|
|`value`|`bytes`|The new value for the metadata|


### registerContract

Registers a new contract address for a specific chain

*Only callable through onCall from CoreRegistry*


```solidity
function registerContract(
    uint256 chainId,
    string calldata contractType,
    bytes calldata addressBytes
)
    external
    onlyRegistry
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain where the contract is deployed|
|`contractType`|`string`|The type of the contract (e.g., "connector", "gateway")|
|`addressBytes`|`bytes`|The address of the contract|


### updateContractConfiguration

Updates contract configuration

*Only callable through onCall from CoreRegistry*


```solidity
function updateContractConfiguration(
    uint256 chainId,
    string calldata contractType,
    string calldata key,
    bytes calldata value
)
    external
    onlyRegistry
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain where the contract is deployed|
|`contractType`|`string`|The type of the contract|
|`key`|`string`|The configuration key to update|
|`value`|`bytes`|The new value for the configuration|


### setContractActive

Sets a contract's active status

*Only callable through onCall from CoreRegistry*


```solidity
function setContractActive(uint256 chainId, string calldata contractType, bool active) external onlyRegistry;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain where the contract is deployed|
|`contractType`|`string`|The type of the contract|
|`active`|`bool`|Whether the contract should be active|


### registerZRC20Token

Registers a new ZRC20 token in the registry

*Only callable through onCall from CoreRegistry*


```solidity
function registerZRC20Token(
    address address_,
    string calldata symbol,
    uint256 originChainId,
    bytes calldata originAddress,
    string calldata coinType,
    uint8 decimals
)
    external
    onlyRegistry
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`address_`|`address`|The address of the ZRC20 token on ZetaChain|
|`symbol`|`string`|The symbol of the token|
|`originChainId`|`uint256`|The ID of the foreign chain where the original asset exists|
|`originAddress`|`bytes`|The address or identifier of the asset on its native chain|
|`coinType`|`string`|The type of the original coin|
|`decimals`|`uint8`|The number of decimals the token uses|


### setZRC20TokenActive

Updates ZRC20 token active status

*Only callable through onCall from CoreRegistry*


```solidity
function setZRC20TokenActive(address address_, bool active) external onlyRegistry whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`address_`|`address`|The address of the ZRC20 token|
|`active`|`bool`|Whether the token should be active|


### bootstrapChains

Bootstrap the registry with chain data

*This function can only be called by an address with the REGISTRY_MANAGER_ROLE.*


```solidity
function bootstrapChains(
    ChainInfoDTO[] calldata chains,
    ChainMetadataEntry[] calldata metadataEntries
)
    external
    onlyRole(REGISTRY_MANAGER_ROLE)
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chains`|`ChainInfoDTO[]`|Array of chain data structures to bootstrap|
|`metadataEntries`|`ChainMetadataEntry[]`|Array of chain metadata entries|


### bootstrapContracts

Bootstrap the registry with contract data

*This function can only be called by an address with the REGISTRY_MANAGER_ROLE.*


```solidity
function bootstrapContracts(
    ContractInfoDTO[] calldata contracts,
    ContractConfigEntry[] calldata configEntries
)
    external
    onlyRole(REGISTRY_MANAGER_ROLE)
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`contracts`|`ContractInfoDTO[]`|Array of contract data structures to bootstrap|
|`configEntries`|`ContractConfigEntry[]`|Array of contract configuration entries|


### bootstrapZRC20Tokens

Bootstrap the registry with ZRC20 token data

*This function can only be called by an address with the REGISTRY_MANAGER_ROLE.*


```solidity
function bootstrapZRC20Tokens(ZRC20Info[] calldata tokens) external onlyRole(REGISTRY_MANAGER_ROLE) whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`tokens`|`ZRC20Info[]`|Array of ZRC20 token data structures to bootstrap|


