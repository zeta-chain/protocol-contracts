# Callable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/b0a690824216f461bd292d05ff57810c5c3ecafd/contracts/evm/interfaces/IGatewayEVM.sol)

Interface implemented by contracts receiving authenticated calls.


## Functions
### onCall


```solidity
function onCall(MessageContext calldata context, bytes calldata message) external payable returns (bytes memory);
```

