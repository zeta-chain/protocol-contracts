# Revertable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/317e9a168aa19dc31b1217eef2a50dbf71ae4d80/contracts/Revert.sol)

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


