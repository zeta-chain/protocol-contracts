# ChainInfo
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/evm/interfaces/IRegistry.sol)

Structure that contains information about a chain.


```solidity
struct ChainInfo {
    bool active;
    uint256 chainId;
    address gasZRC20;
    bytes registry;
    mapping(string => bytes) metadata;
}
```

