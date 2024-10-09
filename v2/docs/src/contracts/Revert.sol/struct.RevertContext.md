# RevertContext
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/0f5a289f3179440ee2e4f23b1aa3613d2e644a78/contracts/Revert.sol)

Struct containing revert context passed to onRevert.


```solidity
struct RevertContext {
    address sender;
    address asset;
    uint256 amount;
    bytes revertMessage;
}
```

**Properties**

|Name|Type|Description|
|----|----|-----------|
|`sender`|`address`|Address of account that initiated smart contract call.|
|`asset`|`address`|Address of asset, empty if it's gas token.|
|`amount`|`uint256`|Amount specified with the transaction.|
|`revertMessage`|`bytes`|Arbitrary data sent back in onRevert.|

