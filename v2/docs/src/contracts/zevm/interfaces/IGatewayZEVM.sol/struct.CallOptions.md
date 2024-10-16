# CallOptions
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/86bca2c09f7eb3b8509097193b2e7504ddcc7cee/contracts/zevm/interfaces/IGatewayZEVM.sol)

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

