# Callable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/7e13d4407420cd4ff52ed44cc892c54d5f3d02cd/contracts/evm/interfaces/IGatewayEVM.sol)

Interface implemented by contracts receiving authenticated calls.


## Functions
### onCall


```solidity
function onCall(MessageContext calldata context, bytes calldata message) external payable returns (bytes memory);
```

