# Solidity API

## ISystem

_Interfaces of SystemContract and ZRC20 to make easier to import._

### FUNGIBLE_MODULE_ADDRESS

```solidity
function FUNGIBLE_MODULE_ADDRESS() external view returns (address)
```

### wZetaContractAddress

```solidity
function wZetaContractAddress() external view returns (address)
```

### uniswapv2FactoryAddress

```solidity
function uniswapv2FactoryAddress() external view returns (address)
```

### gasPriceByChainId

```solidity
function gasPriceByChainId(uint256 chainID) external view returns (uint256)
```

### gasCoinZRC20ByChainId

```solidity
function gasCoinZRC20ByChainId(uint256 chainID) external view returns (address)
```

### gasZetaPoolByChainId

```solidity
function gasZetaPoolByChainId(uint256 chainID) external view returns (address)
```

## IZRC20

### totalSupply

```solidity
function totalSupply() external view returns (uint256)
```

### balanceOf

```solidity
function balanceOf(address account) external view returns (uint256)
```

### transfer

```solidity
function transfer(address recipient, uint256 amount) external returns (bool)
```

### allowance

```solidity
function allowance(address owner, address spender) external view returns (uint256)
```

### approve

```solidity
function approve(address spender, uint256 amount) external returns (bool)
```

### transferFrom

```solidity
function transferFrom(address sender, address recipient, uint256 amount) external returns (bool)
```

### deposit

```solidity
function deposit(address to, uint256 amount) external returns (bool)
```

### withdraw

```solidity
function withdraw(bytes to, uint256 amount) external returns (bool)
```

### withdrawGasFee

```solidity
function withdrawGasFee() external view returns (address, uint256)
```

### Transfer

```solidity
event Transfer(address from, address to, uint256 value)
```

### Approval

```solidity
event Approval(address owner, address spender, uint256 value)
```

### Deposit

```solidity
event Deposit(bytes from, address to, uint256 value)
```

### Withdrawal

```solidity
event Withdrawal(address from, bytes to, uint256 value, uint256 gasfee, uint256 protocolFlatFee)
```

### UpdatedSystemContract

```solidity
event UpdatedSystemContract(address systemContract)
```

### UpdatedGasLimit

```solidity
event UpdatedGasLimit(uint256 gasLimit)
```

### UpdatedProtocolFlatFee

```solidity
event UpdatedProtocolFlatFee(uint256 protocolFlatFee)
```

## Context

### _msgSender

```solidity
function _msgSender() internal view virtual returns (address)
```

### _msgData

```solidity
function _msgData() internal view virtual returns (bytes)
```

## IZRC20Metadata

### name

```solidity
function name() external view returns (string)
```

### symbol

```solidity
function symbol() external view returns (string)
```

### decimals

```solidity
function decimals() external view returns (uint8)
```

## CoinType

```solidity
enum CoinType {
  Zeta,
  Gas,
  ERC20
}
```

## zContract

_Any ZetaChain Contract must implement this interface to allow SystemContract to interact with.
This is only required if the contract wants to interact with other chains._

### onCrossChainCall

```solidity
function onCrossChainCall(address zrc20, uint256 amount, bytes message) external
```

