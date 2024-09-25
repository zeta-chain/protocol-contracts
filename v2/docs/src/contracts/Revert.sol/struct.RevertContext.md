# RevertContext
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/7460fd6c98edf7a57303313f277ac3e3bfe89081/contracts/Revert.sol)

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

