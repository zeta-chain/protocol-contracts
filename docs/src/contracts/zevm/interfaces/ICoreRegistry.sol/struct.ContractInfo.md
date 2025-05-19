# ContractInfo
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/main/v2/v2/v2/v2/v2/v2/v2/v2/contracts/zevm/interfaces/ICoreRegistry.sol)

Structure that contains information about a contract registered in the system.


```solidity
struct ContractInfo {
    bool active;
    bytes addressBytes;
    string contractType;
    mapping(string => bytes) configuration;
}
```

