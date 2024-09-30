# RevertContext
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/eda1e9957411dee4fdf871d69a9caaa035de8918/contracts/Revert.sol)

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

