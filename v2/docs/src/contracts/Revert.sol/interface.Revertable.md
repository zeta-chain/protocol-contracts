# Revertable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/100d7659776cae8d2c060c1333655e0cccd1462a/contracts/Revert.sol)

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


