# UniversalContract
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/0d9bd97652a5b48cac02a68a671d223c054a0a52/contracts/zevm/interfaces/UniversalContract.sol)


## Functions
### onCall


```solidity
function onCall(MessageContext calldata context, address zrc20, uint256 amount, bytes calldata message) external;
```

### onRevert


```solidity
function onRevert(RevertContext calldata revertContext) external;
```

