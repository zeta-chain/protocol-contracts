# Callable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/053cc6a26755df7c287c7e44aada3142e3eaa263/contracts/evm/interfaces/IGatewayEVM.sol)

Interface implemented by contracts receiving authenticated calls.


## Functions
### onCall


```solidity
function onCall(MessageContext calldata context, bytes calldata message) external payable returns (bytes memory);
```

