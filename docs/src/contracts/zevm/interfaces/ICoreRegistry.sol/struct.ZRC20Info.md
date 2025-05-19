# ZRC20Info
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/main/v2/contracts/zevm/interfaces/ICoreRegistry.sol)

Structure that contains information about a ZRC20 token.


```solidity
struct ZRC20Info {
    bool active;
    address address_;
    bytes originAddress;
    uint256 originChainId;
    string symbol;
    string coinType;
    uint8 decimals;
}
```

