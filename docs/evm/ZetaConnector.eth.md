# evm/ZetaConnector.eth.md

## ZetaConnectorEth

_ETH implementation of ZetaConnector.
This contract manages interactions between TSS and different chains.
This version is only for Ethereum network because in the other chains we mint and burn and in this one we lock and unlock._

### constructor

```solidity
constructor(address zetaToken_, address tssAddress_, address tssAddressUpdater_, address pauserAddress_) public
```

### getLockedAmount

```solidity
function getLockedAmount() external view returns (uint256)
```

### send

```solidity
function send(struct ZetaInterfaces.SendInput input) external
```

_Entrypoint to send data through ZetaChain
This call locks the token on the contract and emits an event with all the data needed by the protocol._

### onReceive

```solidity
function onReceive(bytes zetaTxSenderAddress, uint256 sourceChainId, address destinationAddress, uint256 zetaValue, bytes message, bytes32 internalSendHash) external
```

_Handler to receive data from other chain.
This method can be called only by TSS.
Transfers the Zeta tokens to destination and calls onZetaMessage if it's needed._

### onRevert

```solidity
function onRevert(address zetaTxSenderAddress, uint256 sourceChainId, bytes destinationAddress, uint256 destinationChainId, uint256 remainingZetaValue, bytes message, bytes32 internalSendHash) external
```

_Handler to receive errors from other chain.
This method can be called only by TSS.
Transfers the Zeta tokens to destination and calls onZetaRevert if it's needed._

