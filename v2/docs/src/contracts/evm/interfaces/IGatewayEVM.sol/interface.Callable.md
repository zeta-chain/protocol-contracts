# Callable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/57d1446373e344f7366de3008e0bf2b97aeeabf8/contracts/evm/interfaces/IGatewayEVM.sol)

Interface implemented by contracts receiving authenticated calls.


## Functions
### onCall


```solidity
function onCall(MessageContext calldata context, bytes calldata message) external payable returns (bytes memory);
```

