# UniversalContract
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/379b1dc7eff9cbfc41057f0a5e9977fe1d8e1e93/contracts/zevm/interfaces/UniversalContract.sol)


## Functions
### onCrossChainCall


```solidity
function onCrossChainCall(zContext calldata context, address zrc20, uint256 amount, bytes calldata message) external;
```

### onRevert


```solidity
function onRevert(RevertContext calldata revertContext) external;
```

