# RevertContext
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/db433edb21f0084fdbaaabb8e83d79b44b366dd9/contracts/Revert.sol)

Struct containing revert context passed to onRevert.


```solidity
struct RevertContext {
    address asset;
    uint64 amount;
    bytes revertMessage;
}
```

**Properties**

|Name|Type|Description|
|----|----|-----------|
|`asset`|`address`|Address of asset, empty if it's gas token.|
|`amount`|`uint64`|Amount specified with the transaction.|
|`revertMessage`|`bytes`|Arbitrary data sent back in onRevert.|

