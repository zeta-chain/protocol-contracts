# UniversalContract
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/aef054e72dc168bc0642efb673261c9477c170ae/contracts/zevm/interfaces/UniversalContract.sol)


## Functions
### onCall


```solidity
function onCall(MessageContext calldata context, address zrc20, uint256 amount, bytes calldata message) external;
```

### onRevert


```solidity
function onRevert(RevertContext calldata revertContext) external;
```

