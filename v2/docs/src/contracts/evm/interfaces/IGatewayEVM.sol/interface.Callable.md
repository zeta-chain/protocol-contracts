# Callable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/9d9f74dd2aed9f09a6d76a6c496c5bc1db204e89/contracts/evm/interfaces/IGatewayEVM.sol)

Interface implemented by contracts receiving authenticated calls.


## Functions
### onCall


```solidity
function onCall(MessageContext calldata context, bytes calldata message) external payable returns (bytes memory);
```

