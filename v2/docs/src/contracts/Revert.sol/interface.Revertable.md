# Revertable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/1ebf02353d4ffe1e3d159fe4887220a0672a2035/contracts/Revert.sol)

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


