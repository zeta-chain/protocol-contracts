# CallOptions
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/9d9f74dd2aed9f09a6d76a6c496c5bc1db204e89/contracts/zevm/interfaces/IGatewayZEVM.sol)

CallOptions struct passed to call and withdrawAndCall functions.


```solidity
struct CallOptions {
    uint256 gasLimit;
    bool isArbitraryCall;
}
```

**Properties**

|Name|Type|Description|
|----|----|-----------|
|`gasLimit`|`uint256`|Gas limit.|
|`isArbitraryCall`|`bool`|Indicates if call should be arbitrary or authenticated.|

