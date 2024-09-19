# UniversalContract
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/92837ac9178ca835368558d37c2ae9322f290363/contracts/zevm/interfaces/UniversalContract.sol)


## Functions
### onCrossChainCall


```solidity
function onCrossChainCall(zContext calldata context, address zrc20, uint256 amount, bytes calldata message) external;
```

### onRevert


```solidity
function onRevert(RevertContext calldata revertContext) external;
```

