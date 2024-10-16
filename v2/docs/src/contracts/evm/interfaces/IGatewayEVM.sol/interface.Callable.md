# Callable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/aed88f337e1bcebb90f013de9e45bbc5d073ba85/contracts/evm/interfaces/IGatewayEVM.sol)

Interface implemented by contracts receiving authenticated calls.


## Functions
### onCall


```solidity
function onCall(MessageContext calldata context, bytes calldata message) external payable returns (bytes memory);
```

