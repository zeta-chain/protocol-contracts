# UniversalContract
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/639b8fd295b6e67d3a12f04f00a74c5a31cff469/contracts/zevm/interfaces/UniversalContract.sol)


## Functions
### onCrossChainCall


```solidity
function onCrossChainCall(zContext calldata context, address zrc20, uint256 amount, bytes calldata message) external;
```

### onRevert


```solidity
function onRevert(RevertContext calldata revertContext) external;
```

