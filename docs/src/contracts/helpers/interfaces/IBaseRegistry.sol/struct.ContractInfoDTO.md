# ContractInfoDTO
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/main/v2/contracts/helpers/interfaces/IBaseRegistry.sol)

Structure that contains information about a contract registered in the system, used for data retrieving.


```solidity
struct ContractInfoDTO {
    bool active;
    bytes addressBytes;
    string contractType;
    uint256 chainId;
}
```

