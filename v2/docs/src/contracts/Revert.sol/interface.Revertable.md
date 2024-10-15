# Revertable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/45df03a49b31cc5722a5bb6453b743fc8ac35d1f/contracts/Revert.sol)

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


