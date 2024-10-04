# Revertable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/c8047c5cf62b43a480049f1d820da4571e5dcf61/contracts/Revert.sol)

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


