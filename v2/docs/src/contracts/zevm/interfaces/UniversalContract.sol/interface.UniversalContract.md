# UniversalContract
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/7ede96463093bfd534382563222812e5557c84df/contracts/zevm/interfaces/UniversalContract.sol)


## Functions
### onCrossChainCall


```solidity
function onCrossChainCall(zContext calldata context, address zrc20, uint256 amount, bytes calldata message) external;
```

### onRevert


```solidity
function onRevert(RevertContext calldata revertContext) external;
```

