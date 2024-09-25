# UniversalContract
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/85c63250bc8c13c5a3bd590b50ea3b4d1c0d7388/contracts/zevm/interfaces/UniversalContract.sol)


## Functions
### onCrossChainCall


```solidity
function onCrossChainCall(zContext calldata context, address zrc20, uint256 amount, bytes calldata message) external;
```

### onRevert


```solidity
function onRevert(RevertContext calldata revertContext) external;
```

