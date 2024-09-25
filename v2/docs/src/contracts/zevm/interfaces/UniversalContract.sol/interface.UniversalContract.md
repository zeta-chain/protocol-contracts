# UniversalContract
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/db433edb21f0084fdbaaabb8e83d79b44b366dd9/contracts/zevm/interfaces/UniversalContract.sol)


## Functions
### onCrossChainCall


```solidity
function onCrossChainCall(zContext calldata context, address zrc20, uint256 amount, bytes calldata message) external;
```

### onRevert


```solidity
function onRevert(RevertContext calldata revertContext) external;
```

