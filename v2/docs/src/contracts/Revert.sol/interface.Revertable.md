# Revertable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/e55e1d806ff01171e945513bdfc6a523d6a1c116/contracts/Revert.sol)

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


