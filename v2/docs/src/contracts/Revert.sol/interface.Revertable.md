# Revertable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/2ec55b0a1ac0f37560641afe12becc523babb9ca/contracts/Revert.sol)

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


