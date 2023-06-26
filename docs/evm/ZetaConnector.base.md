# evm/ZetaConnector.base.md

## ZetaConnectorBase

_Main abstraction of ZetaConnector.
This contract manages interactions between TSS and different chains.
There's an instance of this contract on each chain supported by ZetaChain._

### zetaToken

```solidity
address zetaToken
```

### pauserAddress

```solidity
address pauserAddress
```

_Multisig contract to pause incoming transactions.
The responsibility of pausing outgoing transactions is left to the protocol for more flexibility._

### tssAddress

```solidity
address tssAddress
```

_Collectively held by ZetaChain validators._

### tssAddressUpdater

```solidity
address tssAddressUpdater
```

_This address will start pointing to a multisig contract, then it will become the TSS address itself._

### ZetaSent

```solidity
event ZetaSent(address sourceTxOriginAddress, address zetaTxSenderAddress, uint256 destinationChainId, bytes destinationAddress, uint256 zetaValueAndGas, uint256 destinationGasLimit, bytes message, bytes zetaParams)
```

### ZetaReceived

```solidity
event ZetaReceived(bytes zetaTxSenderAddress, uint256 sourceChainId, address destinationAddress, uint256 zetaValue, bytes message, bytes32 internalSendHash)
```

### ZetaReverted

```solidity
event ZetaReverted(address zetaTxSenderAddress, uint256 sourceChainId, uint256 destinationChainId, bytes destinationAddress, uint256 remainingZetaValue, bytes message, bytes32 internalSendHash)
```

### TSSAddressUpdated

```solidity
event TSSAddressUpdated(address zetaTxSenderAddress, address newTssAddress)
```

### PauserAddressUpdated

```solidity
event PauserAddressUpdated(address updaterAddress, address newTssAddress)
```

### constructor

```solidity
constructor(address zetaToken_, address tssAddress_, address tssAddressUpdater_, address pauserAddress_) public
```

_Constructor requires initial addresses.
zetaToken address is the only immutable one, while others can be updated._

### onlyPauser

```solidity
modifier onlyPauser()
```

_Modifier to restrict actions to pauser address._

### onlyTssAddress

```solidity
modifier onlyTssAddress()
```

_Modifier to restrict actions to TSS address._

### onlyTssUpdater

```solidity
modifier onlyTssUpdater()
```

_Modifier to restrict actions to TSS updater address._

### updatePauserAddress

```solidity
function updatePauserAddress(address pauserAddress_) external
```

_Update the pauser address. The only address allowed to do that is the current pauser._

### updateTssAddress

```solidity
function updateTssAddress(address tssAddress_) external
```

_Update the TSS address. The address can be updated by the TSS updater or the TSS address itself._

### renounceTssAddressUpdater

```solidity
function renounceTssAddressUpdater() external
```

_Changes the ownership of tssAddressUpdater to be the one held by the ZetaChain TSS Signer nodes._

### pause

```solidity
function pause() external
```

_Pause the input (send) transactions._

### unpause

```solidity
function unpause() external
```

_Unpause the contract to allow transactions again._

### send

```solidity
function send(struct ZetaInterfaces.SendInput input) external virtual
```

_Entrypoint to send data and value through ZetaChain._

### onReceive

```solidity
function onReceive(bytes zetaTxSenderAddress, uint256 sourceChainId, address destinationAddress, uint256 zetaValue, bytes message, bytes32 internalSendHash) external virtual
```

_Handler to receive data from other chain.
This method can be called only by TSS. Access validation is in implementation._

### onRevert

```solidity
function onRevert(address zetaTxSenderAddress, uint256 sourceChainId, bytes destinationAddress, uint256 destinationChainId, uint256 remainingZetaValue, bytes message, bytes32 internalSendHash) external virtual
```

_Handler to receive errors from other chain.
This method can be called only by TSS. Access validation is in implementation._

