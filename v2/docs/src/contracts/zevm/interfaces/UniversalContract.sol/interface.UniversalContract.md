# UniversalContract
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/9a2de1f6dc7a74f680c5d9b6496e47b4b63f4957/contracts/zevm/interfaces/UniversalContract.sol)


## Functions
### onCrossChainCall


```solidity
function onCrossChainCall(zContext calldata context, address zrc20, uint256 amount, bytes calldata message) external;
```

### onRevert


```solidity
function onRevert(RevertContext calldata revertContext) external;
```

