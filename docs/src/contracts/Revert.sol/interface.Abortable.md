# Abortable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/Revert.sol)

Interface for contracts that support abortable calls.


## Functions
### onAbort

Called when a revertable call is aborted.


```solidity
function onAbort(AbortContext calldata abortContext) external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`abortContext`|`AbortContext`|Abort context to pass to onAbort.|


