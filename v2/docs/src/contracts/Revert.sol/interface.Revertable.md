# Revertable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/db433edb21f0084fdbaaabb8e83d79b44b366dd9/contracts/Revert.sol)

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


