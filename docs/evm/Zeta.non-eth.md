# evm/Zeta.non-eth.md

## ZetaNonEth

### connectorAddress

```solidity
address connectorAddress
```

### tssAddress

```solidity
address tssAddress
```

_Collectively held by Zeta blockchain validators_

### tssAddressUpdater

```solidity
address tssAddressUpdater
```

_Initially a multi-sig, eventually held by Zeta blockchain validators (via renounceTssAddressUpdater)_

### Minted

```solidity
event Minted(address mintee, uint256 amount, bytes32 internalSendHash)
```

### Burnt

```solidity
event Burnt(address burnee, uint256 amount)
```

### constructor

```solidity
constructor(address tssAddress_, address tssAddressUpdater_) public
```

### updateTssAndConnectorAddresses

```solidity
function updateTssAndConnectorAddresses(address tssAddress_, address connectorAddress_) external
```

### renounceTssAddressUpdater

```solidity
function renounceTssAddressUpdater() external
```

_Sets tssAddressUpdater to be tssAddress_

### mint

```solidity
function mint(address mintee, uint256 value, bytes32 internalSendHash) external
```

### burnFrom

```solidity
function burnFrom(address account, uint256 amount) public
```

