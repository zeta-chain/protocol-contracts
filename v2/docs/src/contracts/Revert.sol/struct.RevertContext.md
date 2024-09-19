# RevertContext
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/92837ac9178ca835368558d37c2ae9322f290363/contracts/Revert.sol)

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

