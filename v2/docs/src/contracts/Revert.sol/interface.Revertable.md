# Revertable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/379b1dc7eff9cbfc41057f0a5e9977fe1d8e1e93/contracts/Revert.sol)

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


