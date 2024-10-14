# Revertable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/ec2fd2afc191922ecd1aea1903a837977ec7967e/contracts/Revert.sol)

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


