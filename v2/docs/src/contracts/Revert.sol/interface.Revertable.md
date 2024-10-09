# Revertable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/0f5a289f3179440ee2e4f23b1aa3613d2e644a78/contracts/Revert.sol)

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


