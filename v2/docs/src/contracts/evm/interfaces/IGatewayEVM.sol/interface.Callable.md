# Callable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/40c5aaa5c865ea06658f463587fd9248724b3b38/contracts/evm/interfaces/IGatewayEVM.sol)

Interface implemented by contracts receiving authenticated calls.


## Functions
### onCall


```solidity
function onCall(MessageContext calldata context, bytes calldata message) external payable returns (bytes memory);
```

