# Callable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/874f1cc4ec610cadf0a188ddc14f486915de3671/contracts/evm/interfaces/IGatewayEVM.sol)

Interface implemented by contracts receiving authenticated calls.


## Functions
### onCall


```solidity
function onCall(MessageContext calldata context, bytes calldata message) external payable returns (bytes memory);
```

