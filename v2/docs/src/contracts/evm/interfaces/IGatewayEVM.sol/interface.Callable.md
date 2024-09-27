# Callable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/6accdcc6bd3706c438a6f98bc44ddfca182825fe/contracts/evm/interfaces/IGatewayEVM.sol)

Interface implemented by contracts receiving authenticated calls.


## Functions
### onCall


```solidity
function onCall(MessageContext calldata context, bytes calldata message) external payable returns (bytes memory);
```

