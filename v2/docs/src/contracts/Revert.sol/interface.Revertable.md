# Revertable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/639b8fd295b6e67d3a12f04f00a74c5a31cff469/contracts/Revert.sol)

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


