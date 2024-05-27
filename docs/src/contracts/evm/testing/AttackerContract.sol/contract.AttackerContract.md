# AttackerContract
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/760564b6e2ea95b8954e5fd40389cee0cb168d35/contracts/evm/testing/AttackerContract.sol)


## State Variables
### victimContractAddress

```solidity
address public victimContractAddress;
```


### _victimMethod

```solidity
uint256 private _victimMethod;
```


## Functions
### constructor


```solidity
constructor(address victimContractAddress_, address wzeta, uint256 victimMethod);
```

### receive


```solidity
receive() external payable;
```

### attackDeposit


```solidity
function attackDeposit() internal;
```

### attackWidrawal


```solidity
function attackWidrawal() internal;
```

### attack


```solidity
function attack() internal;
```

### balanceOf


```solidity
function balanceOf(address account) external returns (uint256);
```

### transferFrom


```solidity
function transferFrom(address from, address to, uint256 amount) public returns (bool);
```

### transfer


```solidity
function transfer(address to, uint256 amount) public returns (bool);
```

