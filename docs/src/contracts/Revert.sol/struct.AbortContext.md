# AbortContext
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/main/v2/contracts/Revert.sol)

Struct containing abort context passed to onAbort.


```solidity
struct AbortContext {
    bytes sender;
    address asset;
    uint256 amount;
    bool outgoing;
    uint256 chainID;
    bytes revertMessage;
}
```

**Properties**

|Name|Type|Description|
|----|----|-----------|
|`sender`|`bytes`|Address of account that initiated smart contract call. bytes is used as the crosschain transaction can be initiated from a non-EVM chain.|
|`asset`|`address`|Address of asset. On a connected chain, it contains the fungible token address or is empty if it's a gas token. On ZetaChain, it contains the address of the ZRC20.|
|`amount`|`uint256`|Amount specified with the transaction.|
|`outgoing`|`bool`|Flag to indicate if the crosschain transaction was outgoing: from ZetaChain to connected chain. if false, the transaction was incoming: from connected chain to ZetaChain.|
|`chainID`|`uint256`|Chain ID of the connected chain.|
|`revertMessage`|`bytes`|Arbitrary data specified in the RevertOptions object when initating the crosschain transaction.|

