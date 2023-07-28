## ZetaConnectorNonEth

```solidity
import "@zetachain/protocol-contracts/contracts/evm/ZetaConnector.non-eth.sol";
```

Source: https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/evm/ZetaConnector.non-eth.sol

Non ETH implementation of ZetaConnector.
This contract manages interactions between TSS and different chains.
This version is for every chain but Etherum network because in the other chains we mint and burn and in Etherum we lock and unlock

### Function List

* [constructor(zetaTokenAddress_, tssAddress_, tssAddressUpdater_, pauserAddress_)](#ZetaConnectorNonEth-constructor-address-address-address-address-)
* [getLockedAmount()](#ZetaConnectorNonEth-getLockedAmount--)
* [setMaxSupply(maxSupply_)](#ZetaConnectorNonEth-setMaxSupply-uint256-)
* [send(input)](#ZetaConnectorNonEth-send-struct-ZetaInterfaces-SendInput-)
* [onReceive(zetaTxSenderAddress, sourceChainId, destinationAddress, zetaValue, message, internalSendHash)](#ZetaConnectorNonEth-onReceive-bytes-uint256-address-uint256-bytes-bytes32-)
* [onRevert(zetaTxSenderAddress, sourceChainId, destinationAddress, destinationChainId, remainingZetaValue, message, internalSendHash)](#ZetaConnectorNonEth-onRevert-address-uint256-bytes-uint256-uint256-bytes-bytes32-)

### Event List

* [ZetaSent(sourceTxOriginAddress, zetaTxSenderAddress, destinationChainId, destinationAddress, zetaValueAndGas, destinationGasLimit, message, zetaParams)](#ZetaConnectorBase-ZetaSent-address-address-uint256-bytes-uint256-uint256-bytes-bytes-)
* [ZetaReceived(zetaTxSenderAddress, sourceChainId, destinationAddress, zetaValue, message, internalSendHash)](#ZetaConnectorBase-ZetaReceived-bytes-uint256-address-uint256-bytes-bytes32-)
* [ZetaReverted(zetaTxSenderAddress, sourceChainId, destinationChainId, destinationAddress, remainingZetaValue, message, internalSendHash)](#ZetaConnectorBase-ZetaReverted-address-uint256-uint256-bytes-uint256-bytes-bytes32-)
* [TSSAddressUpdated(zetaTxSenderAddress, newTssAddress)](#ZetaConnectorBase-TSSAddressUpdated-address-address-)
* [PauserAddressUpdated(updaterAddress, newTssAddress)](#ZetaConnectorBase-PauserAddressUpdated-address-address-)

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

### Functions

```
constructor(address zetaTokenAddress_, address tssAddress_, address tssAddressUpdater_, address pauserAddress_) (public function)
```

<a name="ZetaConnectorNonEth-constructor-address-address-address-address-"></a>

```
getLockedAmount() → uint256 (external function)
```

<a name="ZetaConnectorNonEth-getLockedAmount--"></a>

```
setMaxSupply(uint256 maxSupply_) (external function)
```

<a name="ZetaConnectorNonEth-setMaxSupply-uint256-"></a>

```
send(struct ZetaInterfaces.SendInput input) (external function)
```

<a name="ZetaConnectorNonEth-send-struct-ZetaInterfaces-SendInput-"></a>

Entry point to send data to protocol
This call burn the token and emit an event with all the data needed by the protocol

```
onReceive(bytes zetaTxSenderAddress, uint256 sourceChainId, address destinationAddress, uint256 zetaValue, bytes message, bytes32 internalSendHash) (external function)
```

<a name="ZetaConnectorNonEth-onReceive-bytes-uint256-address-uint256-bytes-bytes32-"></a>

Handler to receive data from other chain.
This method can be called only by TSS.
Transfer the Zeta tokens to destination and calls onZetaMessage if it's needed.
To perform the transfer mint new tokens, validating first the maxSupply allowed in the current chain.

```
onRevert(address zetaTxSenderAddress, uint256 sourceChainId, bytes destinationAddress, uint256 destinationChainId, uint256 remainingZetaValue, bytes message, bytes32 internalSendHash) (external function)
```

<a name="ZetaConnectorNonEth-onRevert-address-uint256-bytes-uint256-uint256-bytes-bytes32-"></a>

Handler to receive errors from other chain.
This method can be called only by TSS.
Transfer the Zeta tokens to destination and calls onZetaRevert if it's needed.
To perform the transfer mint new tokens, validating first the maxSupply allowed in the current chain.

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
TSSAddressUpdated(address zetaTxSenderAddress, address newTssAddress) (event)
```

<a name="ZetaConnectorBase-TSSAddressUpdated-address-address-"></a>

```
PauserAddressUpdated(address updaterAddress, address newTssAddress) (event)
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

