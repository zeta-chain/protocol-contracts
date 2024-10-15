# Callable
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/2ec55b0a1ac0f37560641afe12becc523babb9ca/contracts/evm/interfaces/IGatewayEVM.sol)

Interface implemented by contracts receiving authenticated calls.


## Functions
### onCall


```solidity
function onCall(MessageContext calldata context, bytes calldata message) external payable returns (bytes memory);
```

