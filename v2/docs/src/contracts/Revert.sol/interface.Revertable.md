# Revertable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/e58dcbf9ce300de7ddf02c03c7589608408cb9a0/contracts/Revert.sol)

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


