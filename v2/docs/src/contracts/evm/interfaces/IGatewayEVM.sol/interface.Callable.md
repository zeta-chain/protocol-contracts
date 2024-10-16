# Callable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/dd30c8678b2e1e344e3f2000d2764e34499cd619/contracts/evm/interfaces/IGatewayEVM.sol)

Interface implemented by contracts receiving authenticated calls.


## Functions
### onCall


```solidity
function onCall(MessageContext calldata context, bytes calldata message) external payable returns (bytes memory);
```

