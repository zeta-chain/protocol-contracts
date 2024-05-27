# ISystem
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/2e5223462d9ac9dedd79e76ede471832bb2c40e7/contracts/zevm/Interfaces.sol)

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

