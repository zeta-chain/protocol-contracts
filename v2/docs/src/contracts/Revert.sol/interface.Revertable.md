# Revertable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/80894a563ae5be7526f28c7162bd136554bc5b86/contracts/Revert.sol)

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


