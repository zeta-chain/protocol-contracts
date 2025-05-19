# IRegistryEvents
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/main/v2/v2/v2/v2/contracts/evm/interfaces/IRegistry.sol)


## Events
### ChainStatusChanged
Emitted when a chain's status is changed


```solidity
event ChainStatusChanged(uint256 indexed chainId, bool newState);
```

### ChainMetadataUpdated
Emitted when new chain metadata is set


```solidity
event ChainMetadataUpdated(uint256 indexed chainId, string key, bytes value);
```

### ContractRegistered
Emitted when a contract is registered


```solidity
event ContractRegistered(uint256 indexed chainId, string contractType, bytes addressBytes);
```

### ContractConfigurationUpdated
Emitted when contract configuration is updated


```solidity
event ContractConfigurationUpdated(uint256 indexed chainId, string contractType, string key, bytes value);
```

### ContractStatusChanged
Emitted when a contract's status is changed


```solidity
event ContractStatusChanged(bytes addressBytes);
```

### ZRC20TokenRegistered
Emitted when a ZRC20 token is registered


```solidity
event ZRC20TokenRegistered(
    bytes originAddress, address indexed zrc20, uint8 decimals, uint256 originChainId, string symbol
);
```

### ZRC20TokenUpdated
Emitted when a ZRC20 token's status is updated


```solidity
event ZRC20TokenUpdated(address indexed zrc20, bool active);
```

