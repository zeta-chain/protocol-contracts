# Callable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/86bca2c09f7eb3b8509097193b2e7504ddcc7cee/contracts/evm/interfaces/IGatewayEVM.sol)

Interface implemented by contracts receiving authenticated calls.


## Functions
### onCall


```solidity
function onCall(MessageContext calldata context, bytes calldata message) external payable returns (bytes memory);
```

