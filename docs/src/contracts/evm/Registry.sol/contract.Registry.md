# Registry
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/main/v2/contracts/evm/Registry.sol)

**Inherits:**
Initializable, UUPSUpgradeable, AccessControlUpgradeable, PausableUpgradeable, [IRegistry](/contracts/evm/interfaces/IRegistry.sol/interface.IRegistry.md)

Satellite registry contract for connected chains, receiving updates from CoreRegistry.

*This contract is deployed on every connected chain and maintains a synchronized view of the registry.*


## State Variables
### PAUSER_ROLE
New role identifier for pauser role


```solidity
bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");
```


### RELAY_ROLE
Identifier for the relay role (granted to GatewayEVM)


```solidity
bytes32 public constant RELAY_ROLE = keccak256("RELAY_ROLE");
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


### activeChains
Active chains in the registry


```solidity
uint256[] public activeChains;
```


### _chains
Maps chain IDs to their information


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
Maps token symbol to ZRC20 address


```solidity
mapping(string => address) private _zrc20SymbolToAddress;
```


### _originAssetToZRC20
Maps origin chain ID and origin address to ZRC20 token address


```solidity
mapping(uint256 => mapping(bytes => address)) private _originAssetToZRC20;
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

### constructor

**Note:**
oz-upgrades-unsafe-allow: constructor


```solidity
constructor();
```

### initialize

Initialize the Registry contract


```solidity
function initialize(
    address admin_,
    address pauserAddress_,
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
|`gatewayEVM_`|`address`|Address of the GatewayEVM contract for cross-chain messaging|
|`coreRegistry_`|`address`|Address of the CoreRegistry contract deployed on ZetaChain|


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

### onCall

onCall is called by the GatewayEVM when a cross-chain message is received


```solidity
function onCall(
    MessageContext calldata context,
    bytes calldata data
)
    external
    onlyRole(RELAY_ROLE)
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
    bool active
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
|`active`|`bool`|Whether activate or deactivate the chain|


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
    address address_,
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
|`address_`|`address`|The address of the contract|
|`contractType`|`string`|The type of the contract (e.g., "connector", "gateway")|
|`addressBytes`|`bytes`|The bytes representation of the non-EVM address|


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

Gets the ZRC20 token address for a specific asset on a foreign chain


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
|`originAddress`|`bytes`|The address or identifier of the asset on its native chain|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`address`|The address of the corresponding ZRC20 token on ZetaChain|


### getActiveChains

Gets all active chains in the registry


```solidity
function getActiveChains() external view returns (uint256[] memory);
```
**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`uint256[]`|Array of chain IDs for all active chains|


### _removeFromActiveChains

Removes a chain ID from the active chains array


```solidity
function _removeFromActiveChains(uint256 chainId) private;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain to remove|


