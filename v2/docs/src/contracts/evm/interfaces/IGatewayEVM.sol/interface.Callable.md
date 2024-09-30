# Callable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/5f09d7eb47b707c65cea167574b26d208e366094/contracts/evm/interfaces/IGatewayEVM.sol)

Interface implemented by contracts receiving authenticated calls.


## Functions
### onCall


```solidity
function onCall(MessageContext calldata context, bytes calldata message) external payable returns (bytes memory);
```

