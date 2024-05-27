# ConnectorErrors
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/760564b6e2ea95b8954e5fd40389cee0cb168d35/contracts/evm/interfaces/ConnectorErrors.sol)

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

