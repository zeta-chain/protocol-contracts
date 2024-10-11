# RevertContext
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/053cc6a26755df7c287c7e44aada3142e3eaa263/contracts/Revert.sol)

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

