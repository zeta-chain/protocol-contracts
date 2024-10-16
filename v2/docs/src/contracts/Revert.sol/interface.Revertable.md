# Revertable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/7e13d4407420cd4ff52ed44cc892c54d5f3d02cd/contracts/Revert.sol)

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


