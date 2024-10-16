# Revertable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/aed88f337e1bcebb90f013de9e45bbc5d073ba85/contracts/Revert.sol)

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


