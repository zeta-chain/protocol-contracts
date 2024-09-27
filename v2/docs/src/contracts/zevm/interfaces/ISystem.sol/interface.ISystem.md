# ISystem
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/6accdcc6bd3706c438a6f98bc44ddfca182825fe/contracts/zevm/interfaces/ISystem.sol)

Interface for the System contract.

*Defines functions for system contract callable by fungible module.*


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

