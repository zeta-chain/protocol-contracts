# ZRC20Events
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/b9a2f8db0b39fa25781ead3dfecdaf63731b3d08/contracts/zevm/interfaces/IZRC20.sol)

Interface for the ZRC20 events.


## Events
### Transfer

```solidity
event Transfer(address indexed from, address indexed to, uint256 value);
```

### Approval

```solidity
event Approval(address indexed owner, address indexed spender, uint256 value);
```

### Deposit

```solidity
event Deposit(bytes from, address indexed to, uint256 value);
```

### Withdrawal

```solidity
event Withdrawal(address indexed from, bytes to, uint256 value, uint256 gasFee, uint256 protocolFlatFee);
```

### UpdatedSystemContract

```solidity
event UpdatedSystemContract(address systemContract);
```

### UpdatedGateway

```solidity
event UpdatedGateway(address gateway);
```

### UpdatedGasLimit

```solidity
event UpdatedGasLimit(uint256 gasLimit);
```

### UpdatedProtocolFlatFee

```solidity
event UpdatedProtocolFlatFee(uint256 protocolFlatFee);
```

