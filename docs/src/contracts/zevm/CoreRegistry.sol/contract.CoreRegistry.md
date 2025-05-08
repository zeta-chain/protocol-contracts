# CoreRegistry
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/main/v2/contracts/zevm/CoreRegistry.sol)

**Inherits:**
Initializable, UUPSUpgradeable, AccessControlUpgradeable, PausableUpgradeable, [ICoreRegistry](/contracts/zevm/interfaces/ICoreRegistry.sol/interface.ICoreRegistry.md)

Central registry for ZetaChain, managing chain info, ZRC20 data, and contract addresses across all chains.

*The contract doesn't hold any funds and should never have active allowances.*


## State Variables
### REGISTRY_MANAGER_ROLE
New role identifier for registry manager role.


```solidity
bytes32 public constant REGISTRY_MANAGER_ROLE = keccak256("REGISTRY_MANAGER_ROLE");
```


### PAUSER_ROLE
New role identifier for pauser role.


```solidity
bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");
```


### CROSS_CHAIN_GAS_LIMIT
Cross-chain message gas limit


```solidity
uint256 public constant CROSS_CHAIN_GAS_LIMIT = 300_000;
```


### gatewayZEVM
Instance of the GatewayZEVM contract for cross-chain communication


```solidity
IGatewayZEVM public gatewayZEVM;
```


### _activeChains
Active chains in the registry


```solidity
uint256[] public _activeChains;
```


### _chains
Maps chain IDs to their information.


```solidity
mapping(uint256 => ChainInfo) private _chains;
```


### _contracts
Maps chain ID -> contract type -> ContractInfo


```solidity
mapping(uint256 => mapping(string => ContractInfo)) private _contracts;
```


### _zrc20Tokens
Maps ZRC20 token address to their information


```solidity
mapping(address => ZRC20Info) private _zrc20Tokens;
```


### _zrc20SymbolToAddress
Maps token symbol to ZRC20 address.


```solidity
mapping(string => address) private _zrc20SymbolToAddress;
```


### _originAssetToZRC20
Maps origin chain ID and origin address to ZRC20 token address.


```solidity
mapping(uint256 => mapping(bytes => address)) private _originAssetToZRC20;
```


## Functions
### constructor

**Note:**
oz-upgrades-unsafe-allow: constructor


```solidity
constructor();
```

### initialize

Initialize the CoreRegistry contract.


```solidity
function initialize(address admin_, address registryManager_, address gatewayZEVM_) public initializer;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`admin_`|`address`|Address with DEFAULT_ADMIN_ROLE, authorized for upgrades and pausing actions.|
|`registryManager_`|`address`|Address with REGISTRY_MANAGER_ROLE, authorized for all registry write actions.|
|`gatewayZEVM_`|`address`|Address of the GatewayZEVM contract for cross-chain messaging|


### _authorizeUpgrade

*Authorizes the upgrade of the contract, sender must be admin.*


```solidity
function _authorizeUpgrade(address newImplementation) internal override onlyRole(DEFAULT_ADMIN_ROLE);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`newImplementation`|`address`|Address of the new implementation,|


### pause

Pause contract.


```solidity
function pause() external onlyRole(PAUSER_ROLE);
```

### unpause

Unpause contract.


```solidity
function unpause() external onlyRole(PAUSER_ROLE);
```

### changeChainStatus

Changes status of the chain to activated/deactivated.


```solidity
function changeChainStatus(
    uint256 chainId,
    address gasZRC20,
    bytes calldata registry,
    bool activation
)
    external
    onlyRole(REGISTRY_MANAGER_ROLE)
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain to activate.|
|`gasZRC20`|`address`|The address of the ZRC20 token that represents gas token for the chain.|
|`registry`|`bytes`|Address of the Registry contract on the connected chain.|
|`activation`|`bool`|Whether activate or deactivate the chain|


### updateChainMetadata

Updates chain metadata, only for the active chains.


```solidity
function updateChainMetadata(
    uint256 chainId,
    string calldata key,
    bytes calldata value
)
    external
    onlyRole(REGISTRY_MANAGER_ROLE)
    whenNotPaused;
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
function registerContract(
    uint256 chainId,
    address address_,
    string calldata contractType,
    bytes calldata addressBytes
)
    external
    onlyRole(REGISTRY_MANAGER_ROLE)
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain where the contract is deployed.|
|`address_`|`address`|The address of the contract.|
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
    external
    onlyRole(REGISTRY_MANAGER_ROLE)
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain where the contract is deployed.|
|`contractType`|`string`|The type of the contract.|
|`key`|`string`|The configuration key to update.|
|`value`|`bytes`|The new value for the configuration.|


### setContractActive

Sets a contract's active status


```solidity
function setContractActive(
    uint256 chainId,
    string calldata contractType,
    bool active
)
    external
    onlyRole(REGISTRY_MANAGER_ROLE)
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain where the contract is deployed.|
|`contractType`|`string`|The type of the contract.|
|`active`|`bool`|Whether the contract should be active.|


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
    external
    onlyRole(REGISTRY_MANAGER_ROLE)
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`address_`|`address`|The address of the ZRC20 token on ZetaChain.|
|`symbol`|`string`|The symbol of the token.|
|`originChainId`|`uint256`|The ID of the foreign chain where the original asset exists.|
|`originAddress`|`bytes`|The address or identifier of the asset on its native chain.|
|`coinType`|`string`|The type of the original coin.|
|`decimals`|`uint8`|The number of decimals the token uses.|


### updateZRC20Token

Updates ZRC20 token active status.


```solidity
function updateZRC20Token(address address_, bool active) external onlyRole(REGISTRY_MANAGER_ROLE) whenNotPaused;
```

### getChainMetadata

Gets chain-specific metadata


```solidity
function getChainMetadata(uint256 chainId, string calldata key) external view returns (bytes memory);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain|
|`key`|`string`|The metadata key to retrieve|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`bytes`|The value of the requested metadata|


### getContractInfo

Gets information about a specific contract


```solidity
function getContractInfo(
    uint256 chainId,
    string calldata contractType
)
    external
    view
    returns (bool active, address address_);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain where the contract is deployed|
|`contractType`|`string`|The type of the contract|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`active`|`bool`|Whether the contract is active|
|`address_`|`address`|The address of the contract|


### getContractConfiguration

Gets contract-specific configuration


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
|`chainId`|`uint256`|The ID of the chain where the contract is deployed|
|`contractType`|`string`|The type of the contract|
|`key`|`string`|The configuration key to retrieve|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`bytes`|The value of the requested configuration|


### getZRC20TokenInfo

Gets information about a specific ZRC20 token


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
|`address_`|`address`|The address of the ZRC20 token|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`active`|`bool`|Whether the token is active|
|`symbol`|`string`|The symbol of the token|
|`originChainId`|`uint256`|The ID of the foreign chain where the original asset exists|
|`originAddress`|`bytes`|The address or identifier of the asset on its native chain|
|`coinType`|`string`|The type of the original coin|
|`decimals`|`uint8`|The number of decimals the token uses|


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
|`originChainId`|`uint256`|The ID of the foreign chain|
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


### _broadcastChainActivation

Broadcast chain activation update to all satellite registries.


```solidity
function _broadcastChainActivation(
    uint256 chainId,
    address gasZRC20,
    bytes calldata registry,
    bool activation
)
    internal;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain being activated/deactivated.|
|`gasZRC20`|`address`|The address of the ZRC20 token that represents gas token for the chain.|
|`registry`|`bytes`|Address of the Registry contract on the connected chain.|
|`activation`|`bool`|Whether activate or deactivate the chain|


### _broadcastChainMetadataUpdate

Broadcast chain metadata to all satellite registries


```solidity
function _broadcastChainMetadataUpdate(uint256 chainId, string calldata key, bytes calldata value) private;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain whose metadata is being updated|
|`key`|`string`|The metadata key being updated|
|`value`|`bytes`|The new value for the metadata|


### _broadcastContractRegistration

Broadcast contract registration to all satellite registries

address_ The address of the contract

contractType The type of the contract

addressBytes The bytes representation of the non-EVM address


```solidity
function _broadcastContractRegistration(
    uint256 chainId,
    address address_,
    string calldata contractType,
    bytes calldata addressBytes
)
    private;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain where the contract is deployed|
|`address_`|`address`||
|`contractType`|`string`||
|`addressBytes`|`bytes`||


### _broadcastContractConfigUpdate

Broadcast contract configuration update to all satellite registries

contractType The type of the contract

key The configuration key being updated

value The new value for the configuration


```solidity
function _broadcastContractConfigUpdate(
    uint256 chainId,
    string calldata contractType,
    string calldata key,
    bytes calldata value
)
    private;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain where the contract is deployed|
|`contractType`|`string`||
|`key`|`string`||
|`value`|`bytes`||


### _broadcastContractStatusUpdate

Broadcast contract status update to all satellite registries

contractType The type of the contract

active Whether the contract should be active


```solidity
function _broadcastContractStatusUpdate(uint256 chainId, string calldata contractType, bool active) private;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain where the contract is deployed|
|`contractType`|`string`||
|`active`|`bool`||


### _broadcastZRC20Registration

Broadcast ZRC20 token registration to all satellite registries


```solidity
function _broadcastZRC20Registration(
    address address_,
    string calldata symbol,
    uint256 originChainId,
    bytes calldata originAddress,
    string calldata coinType,
    uint8 decimals
)
    private;
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


### _broadcastZRC20Update

Broadcast ZRC20 token update to all satellite registries


```solidity
function _broadcastZRC20Update(address address_, bool active) private;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`address_`|`address`|The address of the ZRC20 token|
|`active`|`bool`|Whether the token should be active|


### _broadcastToAllChains

Generic function to broadcast encoded messages to all satellite registries


```solidity
function _broadcastToAllChains(bytes memory encodedMessage) private;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`encodedMessage`|`bytes`|The fully encoded function call to broadcast|


### _sendCrossChainMessage

Sends a cross-chain message to the Registry contract on a target chain.


```solidity
function _sendCrossChainMessage(uint256 targetChainId, bytes memory message) private;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`targetChainId`|`uint256`|The ID of the chain to send the message to.|
|`message`|`bytes`|The encoded function call to execute on the target chain.|


### _removeFromActiveChains

Removes a chain ID from the active chains array.


```solidity
function _removeFromActiveChains(uint256 chainId) private;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain to remove.|


