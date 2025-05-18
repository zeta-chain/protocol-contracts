# IRegistry
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/main/v2/contracts/evm/interfaces/IRegistry.sol)


## Structs
### ChainBootstrapData
Structure for chain data used during bootstrapping


```solidity
struct ChainBootstrapData {
    uint256 chainId;
    address gasZRC20;
    bytes registry;
    bool active;
}
```

### ChainMetadataEntry
Structure for metadata entries used during bootstrapping


```solidity
struct ChainMetadataEntry {
    uint256 chainId;
    string key;
    bytes value;
}
```

### ContractBootstrapData
Structure for contract data used during bootstrapping


```solidity
struct ContractBootstrapData {
    uint256 chainId;
    string contractType;
    bytes addressBytes;
}
```

### ContractConfigEntry
Structure for contract configuration entries used during bootstrapping


```solidity
struct ContractConfigEntry {
    uint256 chainId;
    string contractType;
    string key;
    bytes value;
}
```

### ZRC20BootstrapData
Structure for ZRC20 token data used during bootstrapping


```solidity
struct ZRC20BootstrapData {
    address address_;
    string symbol;
    uint256 originChainId;
    bytes originAddress;
    string coinType;
    uint8 decimals;
    bool active;
}
```

