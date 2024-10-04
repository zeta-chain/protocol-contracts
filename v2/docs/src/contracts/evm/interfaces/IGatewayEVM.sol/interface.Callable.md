# Callable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/e58dcbf9ce300de7ddf02c03c7589608408cb9a0/contracts/evm/interfaces/IGatewayEVM.sol)

Interface implemented by contracts receiving authenticated calls.


## Functions
### onCall


```solidity
function onCall(MessageContext calldata context, bytes calldata message) external payable returns (bytes memory);
```

