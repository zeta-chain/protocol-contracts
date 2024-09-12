# UniversalContract
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/1ebf02353d4ffe1e3d159fe4887220a0672a2035/contracts/zevm/interfaces/UniversalContract.sol)


## Functions
### onCall


```solidity
function onCall(MessageContext calldata context, address zrc20, uint256 amount, bytes calldata message) external;
```

### onRevert


```solidity
function onRevert(RevertContext calldata revertContext) external;
```

