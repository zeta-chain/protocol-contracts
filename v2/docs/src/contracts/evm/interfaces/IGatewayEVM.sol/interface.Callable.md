# Callable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/0f5a289f3179440ee2e4f23b1aa3613d2e644a78/contracts/evm/interfaces/IGatewayEVM.sol)

Interface implemented by contracts receiving authenticated calls.


## Functions
### onCall


```solidity
function onCall(MessageContext calldata context, bytes calldata message) external payable returns (bytes memory);
```

