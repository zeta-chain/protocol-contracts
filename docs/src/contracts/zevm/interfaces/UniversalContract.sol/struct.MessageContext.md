# MessageContext
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/zevm/interfaces/UniversalContract.sol)

Provides contextual information when executing a cross-chain call on ZetaChain.

*This struct helps identify the sender of the message across different blockchain environments.*


```solidity
struct MessageContext {
    bytes sender;
    address senderEVM;
    uint256 chainID;
}
```

