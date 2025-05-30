# BaseRegistry
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/helpers/BaseRegistry.sol)

**Inherits:**
Initializable, UUPSUpgradeable, AccessControlUpgradeable, PausableUpgradeable, [IBaseRegistry](/contracts/helpers/interfaces/IBaseRegistry.sol/interface.IBaseRegistry.md)


## State Variables
### PAUSER_ROLE
New role identifier for pauser role.


```solidity
bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");
```


### REGISTRY_MANAGER_ROLE
New role identifier for registry manager role.


```solidity
bytes32 public constant REGISTRY_MANAGER_ROLE = keccak256("REGISTRY_MANAGER_ROLE");
```


### _activeChains
Active chains in the registry.


```solidity
uint256[] internal _activeChains;
```


### _allChains
Array of all chain IDs in the registry (active and inactive).


```solidity
uint256[] internal _allChains;
```


### _allContracts
Array to store all contracts as chainId and contractType pairs.


```solidity
ContractIdentifier[] internal _allContracts;
```


### _allZRC20Addresses
Array of all ZRC20 token addresses.


```solidity
address[] internal _allZRC20Addresses;
```


### _chains
Maps chain IDs to their information.


```solidity
mapping(uint256 => ChainInfo) internal _chains;
```


### _contracts
Maps chain ID -> contract type -> ContractInfo


```solidity
mapping(uint256 => mapping(string => ContractInfo)) internal _contracts;
```


### _zrc20Tokens
Maps ZRC20 token address to their information


```solidity
mapping(address => ZRC20Info) internal _zrc20Tokens;
```


### _zrc20SymbolToAddress
Maps token symbol to ZRC20 address.


```solidity
mapping(string => address) internal _zrc20SymbolToAddress;
```


### _originAssetToZRC20
Maps origin chain ID and origin address to ZRC20 token address.


```solidity
mapping(uint256 => mapping(bytes => address)) internal _originAssetToZRC20;
```


## Functions
### constructor

**Note:**
oz-upgrades-unsafe-allow: constructor


```solidity
constructor();
```

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
function unpause() external onlyRole(DEFAULT_ADMIN_ROLE);
```

### _changeChainStatus

Changes status of the chain to activated/deactivated.


```solidity
function _changeChainStatus(uint256 chainId, address gasZRC20, bytes calldata registry, bool activation) internal;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain to activate.|
|`gasZRC20`|`address`|The address of the ZRC20 token that represents gas token for the chain.|
|`registry`|`bytes`|Address of the Registry contract on the connected chain.|
|`activation`|`bool`|Whether activate or deactivate the chain|


### _updateChainMetadata

Updates chain metadata, only for the active chains.


```solidity
function _updateChainMetadata(uint256 chainId, string calldata key, bytes calldata value) internal;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain.|
|`key`|`string`|The metadata key to update.|
|`value`|`bytes`|The new value for the metadata.|


### _registerContract

Registers a new contract address for a specific chain.


```solidity
function _registerContract(uint256 chainId, string calldata contractType, bytes calldata addressBytes) internal;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain where the contract is deployed.|
|`contractType`|`string`|The type of the contract (e.g., "connector", "gateway").|
|`addressBytes`|`bytes`|The bytes representation of the non-EVM address.|


### _updateContractConfiguration

Updates contract configuration.


```solidity
function _updateContractConfiguration(
    uint256 chainId,
    string calldata contractType,
    string calldata key,
    bytes calldata value
)
    internal;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain where the contract is deployed.|
|`contractType`|`string`|The type of the contract.|
|`key`|`string`|The configuration key to update.|
|`value`|`bytes`|The new value for the configuration.|


### _setContractActive

Sets a contract's active status


```solidity
function _setContractActive(uint256 chainId, string calldata contractType, bool active) internal;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain where the contract is deployed.|
|`contractType`|`string`|The type of the contract.|
|`active`|`bool`|Whether the contract should be active.|


### _registerZRC20Token

Registers a new ZRC20 token in the registry.


```solidity
function _registerZRC20Token(
    address address_,
    string calldata symbol,
    uint256 originChainId,
    bytes calldata originAddress,
    string calldata coinType,
    uint8 decimals
)
    internal;
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


### _setZRC20TokenActive

Updates ZRC20 token active status.


```solidity
function _setZRC20TokenActive(address address_, bool active) internal;
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
    returns (bool active, bytes memory addressBytes);
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
|`addressBytes`|`bytes`|The address of the contract|


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


### getAllChains

Returns information for all chains (active and inactive) in the registry.


```solidity
function getAllChains() external view returns (ChainInfoDTO[] memory chainsInfo);
```
**Returns**

|Name|Type|Description|
|----|----|-----------|
|`chainsInfo`|`ChainInfoDTO[]`|Array of ChainInfoDTO structs containing information about all chains.|


### getAllContracts

Returns information for all contracts in the registry.


```solidity
function getAllContracts() external view returns (ContractInfoDTO[] memory contractsInfo);
```
**Returns**

|Name|Type|Description|
|----|----|-----------|
|`contractsInfo`|`ContractInfoDTO[]`|Array of ContractInfoDTO structs containing information about all contracts.|


### getAllZRC20Tokens

Returns information for all ZRC20 tokens in the registry.


```solidity
function getAllZRC20Tokens() external view returns (ZRC20Info[] memory tokensInfo);
```
**Returns**

|Name|Type|Description|
|----|----|-----------|
|`tokensInfo`|`ZRC20Info[]`|Array of ZRC20Info structs containing information about all ZRC20 tokens.|


### _removeFromActiveChains

Removes a chain ID from the active chains array.


```solidity
function _removeFromActiveChains(uint256 chainId) private;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain to remove.|


