# ZetaReceiverMock
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/3bb9d457957aef905a86b30e0813a459014e0a7e/contracts/evm/testing/ZetaReceiverMock.sol)

**Inherits:**
[ZetaReceiver](/contracts/zevm/ZetaConnectorZEVM.sol/interface.ZetaReceiver.md)


## Functions
### onZetaMessage


```solidity
function onZetaMessage(ZetaInterfaces.ZetaMessage calldata zetaMessage) external override;
```

### onZetaRevert


```solidity
function onZetaRevert(ZetaInterfaces.ZetaRevert calldata zetaRevert) external override;
```

## Events
### MockOnZetaMessage

```solidity
event MockOnZetaMessage(address destinationAddress);
```

### MockOnZetaRevert

```solidity
event MockOnZetaRevert(address zetaTxSenderAddress);
```

