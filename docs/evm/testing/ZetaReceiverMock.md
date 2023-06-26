# evm/testing/ZetaReceiverMock.md

## ZetaReceiverMock

### MockOnZetaMessage

```solidity
event MockOnZetaMessage(address destinationAddress)
```

### MockOnZetaRevert

```solidity
event MockOnZetaRevert(address zetaTxSenderAddress)
```

### onZetaMessage

```solidity
function onZetaMessage(struct ZetaInterfaces.ZetaMessage zetaMessage) external
```

_onZetaMessage is called when a cross-chain message reaches a contract_

### onZetaRevert

```solidity
function onZetaRevert(struct ZetaInterfaces.ZetaRevert zetaRevert) external
```

_onZetaRevert is called when a cross-chain message reverts.
It's useful to rollback to the original state_

