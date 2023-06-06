# Solidity API

## WETH9

### name

```solidity
string storage pointer name
```

### symbol

```solidity
string storage pointer symbol
```

### decimals

```solidity
uint8 decimals
```

### Approval

```solidity
event Approval(address src, address guy, uint256 wad)
```

### Transfer

```solidity
event Transfer(address src, address dst, uint256 wad)
```

### Deposit

```solidity
event Deposit(address dst, uint256 wad)
```

### Withdrawal

```solidity
event Withdrawal(address src, uint256 wad)
```

### balanceOf

```solidity
mapping(address => uint256) balanceOf
```

### allowance

```solidity
mapping(address => mapping(address => uint256)) allowance
```

### 

```solidity
undefined() public payable
```

### deposit

```solidity
undefined() public payable
```

### withdraw

```solidity
undefined(uint256 wad) public
```

### totalSupply

```solidity
undefined() public view returns (uint256)
```

### approve

```solidity
undefined(address guy, uint256 wad) public returns (bool)
```

### transfer

```solidity
undefined(address dst, uint256 wad) public returns (bool)
```

### transferFrom

```solidity
undefined(address src, address dst, uint256 wad) public returns (bool)
```

