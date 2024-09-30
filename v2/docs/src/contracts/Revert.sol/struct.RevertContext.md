# RevertContext
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/90ee283538f7f481339e056dd409e3957361cddd/contracts/Revert.sol)

Struct containing revert context passed to onRevert.


```solidity
struct RevertContext {
    address sender;
    address asset;
    uint64 amount;
    bytes revertMessage;
}
```

**Properties**

|Name|Type|Description|
|----|----|-----------|
|`sender`|`address`|Address of account that initiated smart contract call.|
|`asset`|`address`|Address of asset, empty if it's gas token.|
|`amount`|`uint64`|Amount specified with the transaction.|
|`revertMessage`|`bytes`|Arbitrary data sent back in onRevert.|

