# ZetaConnectorBase
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/evm/ZetaConnectorBase.sol)

**Inherits:**
Initializable, UUPSUpgradeable, [IZetaConnectorEvents](/contracts/evm/interfaces/IZetaConnector.sol/interface.IZetaConnectorEvents.md), ReentrancyGuardUpgradeable, PausableUpgradeable, AccessControlUpgradeable

Abstract base contract for ZetaConnector.

*This contract implements basic functionality for handling tokens and interacting with the Gateway contract.*


## State Variables
### gateway
The Gateway contract used for executing cross-chain calls.


```solidity
IGatewayEVM public gateway;
```


### zetaToken
The address of the Zeta token.


```solidity
address public zetaToken;
```


### tssAddress
The address of the TSS (Threshold Signature Scheme) contract.


```solidity
address public tssAddress;
```


### WITHDRAWER_ROLE
New role identifier for withdrawer role.


```solidity
bytes32 public constant WITHDRAWER_ROLE = keccak256("WITHDRAWER_ROLE");
```


### PAUSER_ROLE
New role identifier for pauser role.


```solidity
bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");
```


### TSS_ROLE
New role identifier for tss role.


```solidity
bytes32 public constant TSS_ROLE = keccak256("TSS_ROLE");
```


## Functions
### constructor

**Note:**
oz-upgrades-unsafe-allow: constructor


```solidity
constructor();
```

### initialize

Initializer for ZetaConnectors.

*Set admin as default admin and pauser, and tssAddress as tss role.*


```solidity
function initialize(
    address gateway_,
    address zetaToken_,
    address tssAddress_,
    address admin_
)
    public
    virtual
    onlyInitializing;
```

### _authorizeUpgrade

*Authorizes the upgrade of the contract, sender must be owner.*


```solidity
function _authorizeUpgrade(address newImplementation) internal override onlyRole(DEFAULT_ADMIN_ROLE);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`newImplementation`|`address`|Address of the new implementation.|


### updateTSSAddress

Update tss address


```solidity
function updateTSSAddress(address newTSSAddress) external onlyRole(DEFAULT_ADMIN_ROLE);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`newTSSAddress`|`address`|new tss address|


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

### deposit

Handle received tokens.


```solidity
function deposit(uint256 amount) external virtual;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`amount`|`uint256`|The amount of tokens received.|


## Errors
### ZeroAddress
Error indicating that a zero address was provided.


```solidity
error ZeroAddress();
```

