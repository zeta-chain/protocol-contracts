# Revertable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/874f1cc4ec610cadf0a188ddc14f486915de3671/contracts/Revert.sol)

Interface for contracts that support revertable calls.


## Functions
### onRevert

Called when a revertable call is made.


```solidity
function onRevert(RevertContext calldata revertContext) external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`revertContext`|`RevertContext`|Revert context to pass to onRevert.|


