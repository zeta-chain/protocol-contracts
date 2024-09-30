# Revertable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/eda1e9957411dee4fdf871d69a9caaa035de8918/contracts/Revert.sol)

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


