# Revertable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/9a2de1f6dc7a74f680c5d9b6496e47b4b63f4957/contracts/Revert.sol)

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


