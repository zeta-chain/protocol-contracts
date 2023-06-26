# evm/ZetaConnector.non-eth.md

## ZetaConnectorNonEth

_Non ETH implementation of ZetaConnector.
This contract manages interactions between TSS and different chains.
This version is for every chain but Etherum network because in the other chains we mint and burn and in Etherum we lock and unlock_

### maxSupply

```solidity
uint256 maxSupply
```

### constructor

```solidity
constructor(address zetaTokenAddress_, address tssAddress_, address tssAddressUpdater_, address pauserAddress_) public
```

### getLockedAmount

```solidity
function getLockedAmount() external view returns (uint256)
```

### setMaxSupply

```solidity
function setMaxSupply(uint256 maxSupply_) external
```

### send

```solidity
function send(struct ZetaInterfaces.SendInput input) external
```

_Entry point to send data to protocol
This call burn the token and emit an event with all the data needed by the protocol_

### onReceive

```solidity
function onReceive(bytes zetaTxSenderAddress, uint256 sourceChainId, address destinationAddress, uint256 zetaValue, bytes message, bytes32 internalSendHash) external
```

_Handler to receive data from other chain.
This method can be called only by TSS.
Transfer the Zeta tokens to destination and calls onZetaMessage if it's needed.
To perform the transfer mint new tokens, validating first the maxSupply allowed in the current chain._

### onRevert

```solidity
function onRevert(address zetaTxSenderAddress, uint256 sourceChainId, bytes destinationAddress, uint256 destinationChainId, uint256 remainingZetaValue, bytes message, bytes32 internalSendHash) external
```

_Handler to receive errors from other chain.
This method can be called only by TSS.
Transfer the Zeta tokens to destination and calls onZetaRevert if it's needed.
To perform the transfer mint new tokens, validating first the maxSupply allowed in the current chain._

