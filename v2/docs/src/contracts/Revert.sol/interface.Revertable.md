# Revertable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/aef054e72dc168bc0642efb673261c9477c170ae/contracts/Revert.sol)

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


