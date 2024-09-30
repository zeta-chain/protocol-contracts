# Callable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/90ee283538f7f481339e056dd409e3957361cddd/contracts/evm/interfaces/IGatewayEVM.sol)

Interface implemented by contracts receiving authenticated calls.


## Functions
### onCall


```solidity
function onCall(MessageContext calldata context, bytes calldata message) external payable returns (bytes memory);
```

