# ZetaConnectorNative
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/main/v2/contracts/evm/ZetaConnectorNative.sol)

**Inherits:**
[ZetaConnectorBase](/contracts/evm/ZetaConnectorBase.sol/abstract.ZetaConnectorBase.md)

Implementation of ZetaConnectorBase for native token handling.

*This contract directly transfers Zeta tokens and interacts with the Gateway contract.*


## Functions
### withdraw

Withdraw tokens to a specified address.

*This function can only be called by the TSS address.*


```solidity
function withdraw(address to, uint256 amount) external nonReentrant onlyRole(WITHDRAWER_ROLE) whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`to`|`address`|The address to withdraw tokens to.|
|`amount`|`uint256`|The amount of tokens to withdraw.|


### withdrawAndCall

Withdraw tokens and call a contract through Gateway.

*This function can only be called by the TSS address.*


```solidity
function withdrawAndCall(
    MessageContext calldata messageContext,
    address to,
    uint256 amount,
    bytes calldata data
)
    external
    nonReentrant
    onlyRole(WITHDRAWER_ROLE)
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`messageContext`|`MessageContext`|Message context containing sender.|
|`to`|`address`|The address to withdraw tokens to.|
|`amount`|`uint256`|The amount of tokens to withdraw.|
|`data`|`bytes`|The calldata to pass to the contract call.|


### withdrawAndRevert

Withdraw tokens and call a contract with a revert callback through Gateway.

*This function can only be called by the TSS address.*


```solidity
function withdrawAndRevert(
    address to,
    uint256 amount,
    bytes calldata data,
    RevertContext calldata revertContext
)
    external
    nonReentrant
    onlyRole(WITHDRAWER_ROLE)
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`to`|`address`|The address to withdraw tokens to.|
|`amount`|`uint256`|The amount of tokens to withdraw.|
|`data`|`bytes`|The calldata to pass to the contract call.|
|`revertContext`|`RevertContext`|Revert context to pass to onRevert.|


### deposit

Handle received tokens.


```solidity
function deposit(uint256 amount) external override whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`amount`|`uint256`|The amount of tokens received.|


