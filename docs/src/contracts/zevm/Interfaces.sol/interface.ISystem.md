# ISystem
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/760564b6e2ea95b8954e5fd40389cee0cb168d35/contracts/zevm/Interfaces.sol)

*Interfaces of SystemContract and ZRC20 to make easier to import.*


## Functions
### FUNGIBLE_MODULE_ADDRESS


```solidity
function FUNGIBLE_MODULE_ADDRESS() external view returns (address);
```

### wZetaContractAddress


```solidity
function wZetaContractAddress() external view returns (address);
```

### uniswapv2FactoryAddress


```solidity
function uniswapv2FactoryAddress() external view returns (address);
```

### gasPriceByChainId


```solidity
function gasPriceByChainId(uint256 chainID) external view returns (uint256);
```

### gasCoinZRC20ByChainId


```solidity
function gasCoinZRC20ByChainId(uint256 chainID) external view returns (address);
```

### gasZetaPoolByChainId


```solidity
function gasZetaPoolByChainId(uint256 chainID) external view returns (address);
```

