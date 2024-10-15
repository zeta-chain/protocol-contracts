# CallOptions
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/2ec55b0a1ac0f37560641afe12becc523babb9ca/contracts/zevm/interfaces/IGatewayZEVM.sol)

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

