# UniversalContract
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/100d7659776cae8d2c060c1333655e0cccd1462a/contracts/zevm/interfaces/UniversalContract.sol)


## Functions
### onCrossChainCall


```solidity
function onCrossChainCall(zContext calldata context, address zrc20, uint256 amount, bytes calldata message) external;
```

### onRevert


```solidity
function onRevert(RevertContext calldata revertContext) external;
```

