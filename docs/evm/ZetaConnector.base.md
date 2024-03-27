## ZetaConnectorBase

```solidity
import "@zetachain/protocol-contracts/contracts/evm/ZetaConnector.base.sol";
```

Source: https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/evm/ZetaConnector.base.sol

Main abstraction of ZetaConnector.
This contract manages interactions between TSS and different chains.
There's an instance of this contract on each chain supported by ZetaChain.

### Modifier List

* [onlyPauser()](#ZetaConnectorBase-onlyPauser--)
* [onlyTssAddress()](#ZetaConnectorBase-onlyTssAddress--)
* [onlyTssUpdater()](#ZetaConnectorBase-onlyTssUpdater--)

### Function List

* [constructor(zetaToken_, tssAddress_, tssAddressUpdater_, pauserAddress_)](#ZetaConnectorBase-constructor-address-address-address-address-)
* [updatePauserAddress(pauserAddress_)](#ZetaConnectorBase-updatePauserAddress-address-)
* [updateTssAddress(tssAddress_)](#ZetaConnectorBase-updateTssAddress-address-)
* [renounceTssAddressUpdater()](#ZetaConnectorBase-renounceTssAddressUpdater--)
* [pause()](#ZetaConnectorBase-pause--)
* [unpause()](#ZetaConnectorBase-unpause--)
* [send(input)](#ZetaConnectorBase-send-struct-ZetaInterfaces-SendInput-)
* [onReceive(zetaTxSenderAddress, sourceChainId, destinationAddress, zetaValue, message, internalSendHash)](#ZetaConnectorBase-onReceive-bytes-uint256-address-uint256-bytes-bytes32-)
* [onRevert(zetaTxSenderAddress, sourceChainId, destinationAddress, destinationChainId, remainingZetaValue, message, internalSendHash)](#ZetaConnectorBase-onRevert-address-uint256-bytes-uint256-uint256-bytes-bytes32-)

### Event List

* [ZetaSent(sourceTxOriginAddress, zetaTxSenderAddress, destinationChainId, destinationAddress, zetaValueAndGas, destinationGasLimit, message, zetaParams)](#ZetaConnectorBase-ZetaSent-address-address-uint256-bytes-uint256-uint256-bytes-bytes-)
* [ZetaReceived(zetaTxSenderAddress, sourceChainId, destinationAddress, zetaValue, message, internalSendHash)](#ZetaConnectorBase-ZetaReceived-bytes-uint256-address-uint256-bytes-bytes32-)
* [ZetaReverted(zetaTxSenderAddress, sourceChainId, destinationChainId, destinationAddress, remainingZetaValue, message, internalSendHash)](#ZetaConnectorBase-ZetaReverted-address-uint256-uint256-bytes-uint256-bytes-bytes32-)
* [TSSAddressUpdated(callerAddress, newTssAddress)](#ZetaConnectorBase-TSSAddressUpdated-address-address-)
* [TSSAddressUpdaterUpdated(callerAddress, newTssUpdaterAddress)](#ZetaConnectorBase-TSSAddressUpdaterUpdated-address-address-)
* [PauserAddressUpdated(callerAddress, newTssAddress)](#ZetaConnectorBase-PauserAddressUpdated-address-address-)

* [Paused(account)](#Pausable-Paused-address-)
* [Unpaused(account)](#Pausable-Unpaused-address-)

### Error List

* [CallerIsNotPauser(caller)](#ConnectorErrors-CallerIsNotPauser-address-)
* [CallerIsNotTss(caller)](#ConnectorErrors-CallerIsNotTss-address-)
* [CallerIsNotTssUpdater(caller)](#ConnectorErrors-CallerIsNotTssUpdater-address-)
* [CallerIsNotTssOrUpdater(caller)](#ConnectorErrors-CallerIsNotTssOrUpdater-address-)
* [ZetaTransferError()](#ConnectorErrors-ZetaTransferError--)
* [ExceedsMaxSupply(maxSupply)](#ConnectorErrors-ExceedsMaxSupply-uint256-)

### Modifiers

```
onlyPauser() (modifier)
```

<a name="ZetaConnectorBase-onlyPauser--"></a>

Modifier to restrict actions to pauser address.

```
onlyTssAddress() (modifier)
```

<a name="ZetaConnectorBase-onlyTssAddress--"></a>

Modifier to restrict actions to TSS address.

```
onlyTssUpdater() (modifier)
```

<a name="ZetaConnectorBase-onlyTssUpdater--"></a>

Modifier to restrict actions to TSS updater address.

### Functions

```
constructor(address zetaToken_, address tssAddress_, address tssAddressUpdater_, address pauserAddress_) (public function)
```

<a name="ZetaConnectorBase-constructor-address-address-address-address-"></a>

Constructor requires initial addresses.
zetaToken address is the only immutable one, while others can be updated.

```
updatePauserAddress(address pauserAddress_) (external function)
```

<a name="ZetaConnectorBase-updatePauserAddress-address-"></a>

Update the pauser address. The only address allowed to do that is the current pauser.

```
updateTssAddress(address tssAddress_) (external function)
```

<a name="ZetaConnectorBase-updateTssAddress-address-"></a>

Update the TSS address. The address can be updated by the TSS updater or the TSS address itself.

```
renounceTssAddressUpdater() (external function)
```

<a name="ZetaConnectorBase-renounceTssAddressUpdater--"></a>

Changes the ownership of tssAddressUpdater to be the one held by the ZetaChain TSS Signer nodes.

```
pause() (external function)
```

<a name="ZetaConnectorBase-pause--"></a>

Pause the input (send) transactions.

```
unpause() (external function)
```

<a name="ZetaConnectorBase-unpause--"></a>

Unpause the contract to allow transactions again.

```
send(struct ZetaInterfaces.SendInput input) (external function)
```

<a name="ZetaConnectorBase-send-struct-ZetaInterfaces-SendInput-"></a>

Entrypoint to send data and value through ZetaChain.

```
onReceive(bytes zetaTxSenderAddress, uint256 sourceChainId, address destinationAddress, uint256 zetaValue, bytes message, bytes32 internalSendHash) (external function)
```

<a name="ZetaConnectorBase-onReceive-bytes-uint256-address-uint256-bytes-bytes32-"></a>

Handler to receive data from other chain.
This method can be called only by TSS. Access validation is in implementation.

```
onRevert(address zetaTxSenderAddress, uint256 sourceChainId, bytes destinationAddress, uint256 destinationChainId, uint256 remainingZetaValue, bytes message, bytes32 internalSendHash) (external function)
```

<a name="ZetaConnectorBase-onRevert-address-uint256-bytes-uint256-uint256-bytes-bytes32-"></a>

Handler to receive errors from other chain.
This method can be called only by TSS. Access validation is in implementation.

### Events

```
ZetaSent(address sourceTxOriginAddress, address indexed zetaTxSenderAddress, uint256 indexed destinationChainId, bytes destinationAddress, uint256 zetaValueAndGas, uint256 destinationGasLimit, bytes message, bytes zetaParams) (event)
```

<a name="ZetaConnectorBase-ZetaSent-address-address-uint256-bytes-uint256-uint256-bytes-bytes-"></a>

```
ZetaReceived(bytes zetaTxSenderAddress, uint256 indexed sourceChainId, address indexed destinationAddress, uint256 zetaValue, bytes message, bytes32 indexed internalSendHash) (event)
```

<a name="ZetaConnectorBase-ZetaReceived-bytes-uint256-address-uint256-bytes-bytes32-"></a>

```
ZetaReverted(address zetaTxSenderAddress, uint256 sourceChainId, uint256 indexed destinationChainId, bytes destinationAddress, uint256 remainingZetaValue, bytes message, bytes32 indexed internalSendHash) (event)
```

<a name="ZetaConnectorBase-ZetaReverted-address-uint256-uint256-bytes-uint256-bytes-bytes32-"></a>

```
TSSAddressUpdated(address callerAddress, address newTssAddress) (event)
```

<a name="ZetaConnectorBase-TSSAddressUpdated-address-address-"></a>

```
TSSAddressUpdaterUpdated(address callerAddress, address newTssUpdaterAddress) (event)
```

<a name="ZetaConnectorBase-TSSAddressUpdaterUpdated-address-address-"></a>

```
PauserAddressUpdated(address callerAddress, address newTssAddress) (event)
```

<a name="ZetaConnectorBase-PauserAddressUpdated-address-address-"></a>

```
Paused(address account) (event)
```

<a name="Pausable-Paused-address-"></a>

Emitted when the pause is triggered by `account`.

```
Unpaused(address account) (event)
```

<a name="Pausable-Unpaused-address-"></a>

Emitted when the pause is lifted by `account`.

### Errors

```
CallerIsNotPauser(address caller) (error)
```

<a name="ConnectorErrors-CallerIsNotPauser-address-"></a>

```
CallerIsNotTss(address caller) (error)
```

<a name="ConnectorErrors-CallerIsNotTss-address-"></a>

```
CallerIsNotTssUpdater(address caller) (error)
```

<a name="ConnectorErrors-CallerIsNotTssUpdater-address-"></a>

```
CallerIsNotTssOrUpdater(address caller) (error)
```

<a name="ConnectorErrors-CallerIsNotTssOrUpdater-address-"></a>

```
ZetaTransferError() (error)
```

<a name="ConnectorErrors-ZetaTransferError--"></a>

```
ExceedsMaxSupply(uint256 maxSupply) (error)
```

<a name="ConnectorErrors-ExceedsMaxSupply-uint256-"></a>

