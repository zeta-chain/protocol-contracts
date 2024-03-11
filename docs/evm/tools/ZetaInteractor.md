## ZetaInteractor

```solidity
import "@zetachain/protocol-contracts/contracts/evm/tools/ZetaInteractor.sol";
```

Source: https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/evm/tools/ZetaInteractor.sol

### Modifier List

* [isValidMessageCall(zetaMessage)](#ZetaInteractor-isValidMessageCall-struct-ZetaInterfaces-ZetaMessage-)
* [isValidRevertCall(zetaRevert)](#ZetaInteractor-isValidRevertCall-struct-ZetaInterfaces-ZetaRevert-)

### Function List

* [constructor(zetaConnectorAddress)](#ZetaInteractor-constructor-address-)
* [_isValidChainId(chainId)](#ZetaInteractor-_isValidChainId-uint256-)
* [setInteractorByChainId(destinationChainId, contractAddress)](#ZetaInteractor-setInteractorByChainId-uint256-bytes-)

### Event List

* [OwnershipTransferStarted(previousOwner, newOwner)](#Ownable2Step-OwnershipTransferStarted-address-address-)

* [OwnershipTransferred(previousOwner, newOwner)](#Ownable-OwnershipTransferred-address-address-)

### Error List

* [InvalidDestinationChainId()](#ZetaInteractorErrors-InvalidDestinationChainId--)
* [InvalidCaller(caller)](#ZetaInteractorErrors-InvalidCaller-address-)
* [InvalidZetaMessageCall()](#ZetaInteractorErrors-InvalidZetaMessageCall--)
* [InvalidZetaRevertCall()](#ZetaInteractorErrors-InvalidZetaRevertCall--)

### Modifiers

```
isValidMessageCall(struct ZetaInterfaces.ZetaMessage zetaMessage) (modifier)
```

<a name="ZetaInteractor-isValidMessageCall-struct-ZetaInterfaces-ZetaMessage-"></a>

```
isValidRevertCall(struct ZetaInterfaces.ZetaRevert zetaRevert) (modifier)
```

<a name="ZetaInteractor-isValidRevertCall-struct-ZetaInterfaces-ZetaRevert-"></a>

### Functions

```
constructor(address zetaConnectorAddress) (internal function)
```

<a name="ZetaInteractor-constructor-address-"></a>

```
_isValidChainId(uint256 chainId) → bool (internal function)
```

<a name="ZetaInteractor-_isValidChainId-uint256-"></a>

Useful for contracts that inherit from this one

```
setInteractorByChainId(uint256 destinationChainId, bytes contractAddress) (external function)
```

<a name="ZetaInteractor-setInteractorByChainId-uint256-bytes-"></a>

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

