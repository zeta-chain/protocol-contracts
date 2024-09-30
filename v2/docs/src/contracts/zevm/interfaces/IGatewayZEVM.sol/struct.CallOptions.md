# CallOptions
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/07bc421f7b5d3de21d96407c91e6a1e2e7289a16/contracts/zevm/interfaces/IGatewayZEVM.sol)

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

