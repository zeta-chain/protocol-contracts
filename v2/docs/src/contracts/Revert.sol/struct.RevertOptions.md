# RevertOptions
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/9d9f74dd2aed9f09a6d76a6c496c5bc1db204e89/contracts/Revert.sol)

Struct containing revert options


```solidity
struct RevertOptions {
    address revertAddress;
    bool callOnRevert;
    address abortAddress;
    bytes revertMessage;
    uint256 onRevertGasLimit;
}
```

**Properties**

|Name|Type|Description|
|----|----|-----------|
|`revertAddress`|`address`|Address to receive revert.|
|`callOnRevert`|`bool`|Flag if onRevert hook should be called.|
|`abortAddress`|`address`|Address to receive funds if aborted.|
|`revertMessage`|`bytes`|Arbitrary data sent back in onRevert.|
|`onRevertGasLimit`|`uint256`|Gas limit for revert tx, unused on GatewayZEVM methods|

