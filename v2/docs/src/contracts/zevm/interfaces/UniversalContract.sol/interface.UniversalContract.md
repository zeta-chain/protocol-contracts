# UniversalContract
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/03043003e2b510828e96289d740026d785c81bde/contracts/zevm/interfaces/UniversalContract.sol)


## Functions
### onCall


```solidity
function onCall(MessageContext calldata context, address zrc20, uint256 amount, bytes calldata message) external;
```

### onRevert


```solidity
function onRevert(RevertContext calldata revertContext) external;
```

