# Revertable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/07bc421f7b5d3de21d96407c91e6a1e2e7289a16/contracts/Revert.sol)

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


