# Callable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/d9d250da078cc3bcf80823d625da2236cadb5515/contracts/evm/interfaces/IGatewayEVM.sol)

Interface implemented by contracts receiving authenticated calls.


## Functions
### onCall


```solidity
function onCall(MessageContext calldata context, bytes calldata message) external payable returns (bytes memory);
```

