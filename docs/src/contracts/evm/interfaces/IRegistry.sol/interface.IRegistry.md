# IRegistry
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/evm/interfaces/IRegistry.sol)


## Structs
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

