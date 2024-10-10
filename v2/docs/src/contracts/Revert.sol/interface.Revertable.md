# Revertable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/bacc1a1f30b7a6d837b8275e0dfeae0c739ef3ee/contracts/Revert.sol)

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


