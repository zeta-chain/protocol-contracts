# Revertable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/85c63250bc8c13c5a3bd590b50ea3b4d1c0d7388/contracts/Revert.sol)

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


