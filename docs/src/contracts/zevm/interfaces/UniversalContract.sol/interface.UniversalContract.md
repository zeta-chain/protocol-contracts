# UniversalContract
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/main/v2/v2/v2/v2/contracts/zevm/interfaces/UniversalContract.sol)

Interface for contracts that can receive cross-chain calls on ZetaChain.

*Contracts implementing this interface can handle incoming cross-chain messages
and execute logic based on the provided context, token, and message payload.*


## Functions
### onCall


```solidity
function onCall(MessageContext calldata context, address zrc20, uint256 amount, bytes calldata message) external;
```

