# Callable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/07bc421f7b5d3de21d96407c91e6a1e2e7289a16/contracts/evm/interfaces/IGatewayEVM.sol)

Interface implemented by contracts receiving authenticated calls.


## Functions
### onCall


```solidity
function onCall(MessageContext calldata context, bytes calldata message) external payable returns (bytes memory);
```

