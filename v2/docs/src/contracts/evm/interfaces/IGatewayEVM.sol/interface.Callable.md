# Callable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/6d255086a2997fe8d79339bbb664b6c1d433f9e9/contracts/evm/interfaces/IGatewayEVM.sol)

Interface implemented by contracts receiving authenticated calls.


## Functions
### onCall


```solidity
function onCall(MessageContext calldata context, bytes calldata message) external payable returns (bytes memory);
```

