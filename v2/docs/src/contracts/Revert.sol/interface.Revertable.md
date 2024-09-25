# Revertable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/879ee5a0e19b2c0f863cd2125f3282f628963d5a/contracts/Revert.sol)

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


