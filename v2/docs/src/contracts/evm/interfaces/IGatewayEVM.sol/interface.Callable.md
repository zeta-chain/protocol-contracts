# Callable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/317e9a168aa19dc31b1217eef2a50dbf71ae4d80/contracts/evm/interfaces/IGatewayEVM.sol)

Interface implemented by contracts receiving authenticated calls.


## Functions
### onCall


```solidity
function onCall(MessageContext calldata context, bytes calldata message) external payable returns (bytes memory);
```

