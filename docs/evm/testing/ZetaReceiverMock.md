## ZetaReceiverMock

```solidity
import "@zetachain/protocol-contracts/contracts/evm/testing/ZetaReceiverMock.sol";
```

Source: https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/evm/testing/ZetaReceiverMock.sol

### Function List

* [onZetaMessage(zetaMessage)](#ZetaReceiverMock-onZetaMessage-struct-ZetaInterfaces-ZetaMessage-)
* [onZetaRevert(zetaRevert)](#ZetaReceiverMock-onZetaRevert-struct-ZetaInterfaces-ZetaRevert-)

### Event List

* [MockOnZetaMessage(destinationAddress)](#ZetaReceiverMock-MockOnZetaMessage-address-)
* [MockOnZetaRevert(zetaTxSenderAddress)](#ZetaReceiverMock-MockOnZetaRevert-address-)

### Modifiers

### Functions

```
onZetaMessage(struct ZetaInterfaces.ZetaMessage zetaMessage) (external function)
```

<a name="ZetaReceiverMock-onZetaMessage-struct-ZetaInterfaces-ZetaMessage-"></a>

onZetaMessage is called when a cross-chain message reaches a contract

```
onZetaRevert(struct ZetaInterfaces.ZetaRevert zetaRevert) (external function)
```

<a name="ZetaReceiverMock-onZetaRevert-struct-ZetaInterfaces-ZetaRevert-"></a>

onZetaRevert is called when a cross-chain message reverts.
It's useful to rollback to the original state

### Events

```
MockOnZetaMessage(address destinationAddress) (event)
```

<a name="ZetaReceiverMock-MockOnZetaMessage-address-"></a>

```
MockOnZetaRevert(address zetaTxSenderAddress) (event)
```

<a name="ZetaReceiverMock-MockOnZetaRevert-address-"></a>

