# Callable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/c8047c5cf62b43a480049f1d820da4571e5dcf61/contracts/evm/interfaces/IGatewayEVM.sol)

Interface implemented by contracts receiving authenticated calls.


## Functions
### onCall


```solidity
function onCall(MessageContext calldata context, bytes calldata message) external payable returns (bytes memory);
```

