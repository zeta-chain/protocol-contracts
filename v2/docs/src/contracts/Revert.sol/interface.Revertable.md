# Revertable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/b0a690824216f461bd292d05ff57810c5c3ecafd/contracts/Revert.sol)

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


