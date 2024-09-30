# Callable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/eda1e9957411dee4fdf871d69a9caaa035de8918/contracts/evm/interfaces/IGatewayEVM.sol)

Interface implemented by contracts receiving authenticated calls.


## Functions
### onCall


```solidity
function onCall(MessageContext calldata context, bytes calldata message) external payable returns (bytes memory);
```

