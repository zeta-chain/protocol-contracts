# Revertable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/dd30c8678b2e1e344e3f2000d2764e34499cd619/contracts/Revert.sol)

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


