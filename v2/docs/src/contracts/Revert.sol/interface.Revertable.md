# Revertable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/6accdcc6bd3706c438a6f98bc44ddfca182825fe/contracts/Revert.sol)

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


