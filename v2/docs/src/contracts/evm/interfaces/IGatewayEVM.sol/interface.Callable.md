# Callable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/45df03a49b31cc5722a5bb6453b743fc8ac35d1f/contracts/evm/interfaces/IGatewayEVM.sol)

Interface implemented by contracts receiving authenticated calls.


## Functions
### onCall


```solidity
function onCall(MessageContext calldata context, bytes calldata message) external payable returns (bytes memory);
```

