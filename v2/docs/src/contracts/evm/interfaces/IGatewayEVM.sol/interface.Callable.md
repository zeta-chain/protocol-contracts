# Callable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/3d536cb237e91172634240b520e138c216b41a29/contracts/evm/interfaces/IGatewayEVM.sol)

Interface implemented by contracts receiving authenticated calls.


## Functions
### onCall


```solidity
function onCall(MessageContext calldata context, bytes calldata message) external payable returns (bytes memory);
```

