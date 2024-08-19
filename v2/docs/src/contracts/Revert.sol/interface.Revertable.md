# Revertable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/c36a179b97a1de4e07d52f443c0ba9860e83aa72/contracts/Revert.sol)

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


