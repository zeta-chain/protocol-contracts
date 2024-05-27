# ConnectorErrors
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/2e5223462d9ac9dedd79e76ede471832bb2c40e7/contracts/evm/interfaces/ConnectorErrors.sol)

*Interface with connector custom errors*


## Errors
### CallerIsNotPauser

```solidity
error CallerIsNotPauser(address caller);
```

### CallerIsNotTss

```solidity
error CallerIsNotTss(address caller);
```

### CallerIsNotTssUpdater

```solidity
error CallerIsNotTssUpdater(address caller);
```

### CallerIsNotTssOrUpdater

```solidity
error CallerIsNotTssOrUpdater(address caller);
```

### ZetaTransferError

```solidity
error ZetaTransferError();
```

### ExceedsMaxSupply

```solidity
error ExceedsMaxSupply(uint256 maxSupply);
```

