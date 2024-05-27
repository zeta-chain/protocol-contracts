# ZetaReceiverMock
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/2e5223462d9ac9dedd79e76ede471832bb2c40e7/contracts/evm/testing/ZetaReceiverMock.sol)

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

