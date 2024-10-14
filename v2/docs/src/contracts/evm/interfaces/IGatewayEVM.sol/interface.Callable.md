# Callable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/ec2fd2afc191922ecd1aea1903a837977ec7967e/contracts/evm/interfaces/IGatewayEVM.sol)

Interface implemented by contracts receiving authenticated calls.


## Functions
### onCall


```solidity
function onCall(MessageContext calldata context, bytes calldata message) external payable returns (bytes memory);
```

