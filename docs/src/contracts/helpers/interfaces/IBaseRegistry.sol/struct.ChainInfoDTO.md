# ChainInfoDTO
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/main/v2/contracts/helpers/interfaces/IBaseRegistry.sol)

Structure that contains information about a chain, used for data retrieving.


```solidity
struct ChainInfoDTO {
    bool active;
    uint256 chainId;
    address gasZRC20;
    bytes registry;
}
```

