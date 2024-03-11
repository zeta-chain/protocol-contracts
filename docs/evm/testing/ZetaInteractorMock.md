## ZetaInteractorMock

```solidity
import "@zetachain/protocol-contracts/contracts/evm/testing/ZetaInteractorMock.sol";
```

Source: https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/evm/testing/ZetaInteractorMock.sol

### Function List

* [constructor(zetaConnectorAddress)](#ZetaInteractorMock-constructor-address-)
* [onZetaMessage(zetaMessage)](#ZetaInteractorMock-onZetaMessage-struct-ZetaInterfaces-ZetaMessage-)
* [onZetaRevert(zetaRevert)](#ZetaInteractorMock-onZetaRevert-struct-ZetaInterfaces-ZetaRevert-)

### Event List

* [OwnershipTransferStarted(previousOwner, newOwner)](#Ownable2Step-OwnershipTransferStarted-address-address-)

* [OwnershipTransferred(previousOwner, newOwner)](#Ownable-OwnershipTransferred-address-address-)

### Error List

* [InvalidDestinationChainId()](#ZetaInteractorErrors-InvalidDestinationChainId--)
* [InvalidCaller(caller)](#ZetaInteractorErrors-InvalidCaller-address-)
* [InvalidZetaMessageCall()](#ZetaInteractorErrors-InvalidZetaMessageCall--)
* [InvalidZetaRevertCall()](#ZetaInteractorErrors-InvalidZetaRevertCall--)

### Modifiers

### Functions

```
constructor(address zetaConnectorAddress) (public function)
```

<a name="ZetaInteractorMock-constructor-address-"></a>

```
onZetaMessage(struct ZetaInterfaces.ZetaMessage zetaMessage) (external function)
```

<a name="ZetaInteractorMock-onZetaMessage-struct-ZetaInterfaces-ZetaMessage-"></a>

onZetaMessage is called when a cross-chain message reaches a contract

```
onZetaRevert(struct ZetaInterfaces.ZetaRevert zetaRevert) (external function)
```

<a name="ZetaInteractorMock-onZetaRevert-struct-ZetaInterfaces-ZetaRevert-"></a>

onZetaRevert is called when a cross-chain message reverts.
It's useful to rollback to the original state

### Events

```
OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner) (event)
```

<a name="Ownable2Step-OwnershipTransferStarted-address-address-"></a>

```
OwnershipTransferred(address indexed previousOwner, address indexed newOwner) (event)
```

<a name="Ownable-OwnershipTransferred-address-address-"></a>

### Errors

```
InvalidDestinationChainId() (error)
```

<a name="ZetaInteractorErrors-InvalidDestinationChainId--"></a>

```
InvalidCaller(address caller) (error)
```

<a name="ZetaInteractorErrors-InvalidCaller-address-"></a>

```
InvalidZetaMessageCall() (error)
```

<a name="ZetaInteractorErrors-InvalidZetaMessageCall--"></a>

```
InvalidZetaRevertCall() (error)
```

<a name="ZetaInteractorErrors-InvalidZetaRevertCall--"></a>

