# RevertContext
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/7ede96463093bfd534382563222812e5557c84df/contracts/Revert.sol)

Struct containing revert context passed to onRevert.


```solidity
struct RevertContext {
    address asset;
    uint256 amount;
    bytes revertMessage;
}
```

**Properties**

|Name|Type|Description|
|----|----|-----------|
|`asset`|`address`|Address of asset, empty if it's gas token.|
|`amount`|`uint256`|Amount specified with the transaction.|
|`revertMessage`|`bytes`|Arbitrary data sent back in onRevert.|

