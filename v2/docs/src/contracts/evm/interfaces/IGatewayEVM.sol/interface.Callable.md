# Callable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/e55e1d806ff01171e945513bdfc6a523d6a1c116/contracts/evm/interfaces/IGatewayEVM.sol)

Interface implemented by contracts receiving authenticated calls.


## Functions
### onCall


```solidity
function onCall(MessageContext calldata context, bytes calldata message) external payable returns (bytes memory);
```

