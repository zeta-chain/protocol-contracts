# Revertable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/86bca2c09f7eb3b8509097193b2e7504ddcc7cee/contracts/Revert.sol)

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


