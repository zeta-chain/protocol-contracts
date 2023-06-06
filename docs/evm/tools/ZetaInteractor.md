# Solidity API

## ZetaInteractor

### ZERO_BYTES

```solidity
bytes32 ZERO_BYTES
```

### currentChainId

```solidity
uint256 currentChainId
```

### connector

```solidity
contract ZetaConnector connector
```

### interactorsByChainId

```solidity
mapping(uint256 => bytes) interactorsByChainId
```

_Maps a chain id to its corresponding address of the MultiChainSwap contract
The address is expressed in bytes to allow non-EVM chains
This mapping is useful, mainly, for two reasons:
 - Given a chain id, the contract is able to route a transaction to its corresponding address
 - To check that the messages (onZetaMessage, onZetaRevert) come from a trusted source_

### isValidMessageCall

```solidity
modifier isValidMessageCall(struct ZetaInterfaces.ZetaMessage zetaMessage)
```

### isValidRevertCall

```solidity
modifier isValidRevertCall(struct ZetaInterfaces.ZetaRevert zetaRevert)
```

### constructor

```solidity
constructor(address zetaConnectorAddress) internal
```

### _isValidChainId

```solidity
function _isValidChainId(uint256 chainId) internal view returns (bool)
```

_Useful for contracts that inherit from this one_

### setInteractorByChainId

```solidity
function setInteractorByChainId(uint256 destinationChainId, bytes contractAddress) external
```

