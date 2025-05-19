# IRegistry
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/main/v2/contracts/evm/interfaces/IRegistry.sol)


## Structs
### ChainBootstrapData
Structure for chain data used during bootstrapping


```solidity
struct ChainBootstrapData {
    bool active;
    uint256 chainId;
    address gasZRC20;
    bytes registry;
}
```

### ContractBootstrapData
Structure for contract data used during bootstrapping


```solidity
struct ContractBootstrapData {
    bool active;
    bytes addressBytes;
    string contractType;
    uint256 chainId;
}
```

### ZRC20BootstrapData
Structure for ZRC20 token data used during bootstrapping


```solidity
struct ZRC20BootstrapData {
    bool active;
    address address_;
    bytes originAddress;
    uint256 originChainId;
    string symbol;
    string coinType;
    uint8 decimals;
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

