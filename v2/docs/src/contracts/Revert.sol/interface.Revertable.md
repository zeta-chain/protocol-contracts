# Revertable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/b9a2f8db0b39fa25781ead3dfecdaf63731b3d08/contracts/Revert.sol)

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


