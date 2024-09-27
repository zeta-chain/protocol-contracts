# Revertable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/5f09d7eb47b707c65cea167574b26d208e366094/contracts/Revert.sol)

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


