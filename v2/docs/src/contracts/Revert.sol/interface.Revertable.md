# Revertable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/40c5aaa5c865ea06658f463587fd9248724b3b38/contracts/Revert.sol)

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


