# Revertable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/7ede96463093bfd534382563222812e5557c84df/contracts/Revert.sol)

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


