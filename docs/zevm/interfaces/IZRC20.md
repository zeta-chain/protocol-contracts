## IZRC20

```solidity
import "@zetachain/protocol-contracts/contracts/zevm/interfaces/IZRC20.sol";
```

Source: https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/zevm/interfaces/IZRC20.sol

### Function List

* [totalSupply()](#IZRC20-totalSupply--)
* [balanceOf(account)](#IZRC20-balanceOf-address-)
* [transfer(recipient, amount)](#IZRC20-transfer-address-uint256-)
* [allowance(owner, spender)](#IZRC20-allowance-address-address-)
* [approve(spender, amount)](#IZRC20-approve-address-uint256-)
* [decreaseAllowance(spender, amount)](#IZRC20-decreaseAllowance-address-uint256-)
* [increaseAllowance(spender, amount)](#IZRC20-increaseAllowance-address-uint256-)
* [transferFrom(sender, recipient, amount)](#IZRC20-transferFrom-address-address-uint256-)
* [deposit(to, amount)](#IZRC20-deposit-address-uint256-)
* [burn(account, amount)](#IZRC20-burn-address-uint256-)
* [withdraw(to, amount)](#IZRC20-withdraw-bytes-uint256-)
* [withdrawGasFee()](#IZRC20-withdrawGasFee--)
* [PROTOCOL_FEE()](#IZRC20-PROTOCOL_FEE--)

### Event List

* [Transfer(from, to, value)](#IZRC20-Transfer-address-address-uint256-)
* [Approval(owner, spender, value)](#IZRC20-Approval-address-address-uint256-)
* [Deposit(from, to, value)](#IZRC20-Deposit-bytes-address-uint256-)
* [Withdrawal(from, to, value, gasFee, protocolFlatFee)](#IZRC20-Withdrawal-address-bytes-uint256-uint256-uint256-)
* [UpdatedSystemContract(systemContract)](#IZRC20-UpdatedSystemContract-address-)
* [UpdatedGasLimit(gasLimit)](#IZRC20-UpdatedGasLimit-uint256-)
* [UpdatedProtocolFlatFee(protocolFlatFee)](#IZRC20-UpdatedProtocolFlatFee-uint256-)

### Modifiers

### Functions

```
totalSupply() → uint256 (external function)
```

<a name="IZRC20-totalSupply--"></a>

```
balanceOf(address account) → uint256 (external function)
```

<a name="IZRC20-balanceOf-address-"></a>

```
transfer(address recipient, uint256 amount) → bool (external function)
```

<a name="IZRC20-transfer-address-uint256-"></a>

```
allowance(address owner, address spender) → uint256 (external function)
```

<a name="IZRC20-allowance-address-address-"></a>

```
approve(address spender, uint256 amount) → bool (external function)
```

<a name="IZRC20-approve-address-uint256-"></a>

```
decreaseAllowance(address spender, uint256 amount) → bool (external function)
```

<a name="IZRC20-decreaseAllowance-address-uint256-"></a>

```
increaseAllowance(address spender, uint256 amount) → bool (external function)
```

<a name="IZRC20-increaseAllowance-address-uint256-"></a>

```
transferFrom(address sender, address recipient, uint256 amount) → bool (external function)
```

<a name="IZRC20-transferFrom-address-address-uint256-"></a>

```
deposit(address to, uint256 amount) → bool (external function)
```

<a name="IZRC20-deposit-address-uint256-"></a>

```
burn(address account, uint256 amount) → bool (external function)
```

<a name="IZRC20-burn-address-uint256-"></a>

```
withdraw(bytes to, uint256 amount) → bool (external function)
```

<a name="IZRC20-withdraw-bytes-uint256-"></a>

```
withdrawGasFee() → address, uint256 (external function)
```

<a name="IZRC20-withdrawGasFee--"></a>

```
PROTOCOL_FEE() → uint256 (external function)
```

<a name="IZRC20-PROTOCOL_FEE--"></a>

### Events

```
Transfer(address indexed from, address indexed to, uint256 value) (event)
```

<a name="IZRC20-Transfer-address-address-uint256-"></a>

```
Approval(address indexed owner, address indexed spender, uint256 value) (event)
```

<a name="IZRC20-Approval-address-address-uint256-"></a>

```
Deposit(bytes from, address indexed to, uint256 value) (event)
```

<a name="IZRC20-Deposit-bytes-address-uint256-"></a>

```
Withdrawal(address indexed from, bytes to, uint256 value, uint256 gasFee, uint256 protocolFlatFee) (event)
```

<a name="IZRC20-Withdrawal-address-bytes-uint256-uint256-uint256-"></a>

```
UpdatedSystemContract(address systemContract) (event)
```

<a name="IZRC20-UpdatedSystemContract-address-"></a>

```
UpdatedGasLimit(uint256 gasLimit) (event)
```

<a name="IZRC20-UpdatedGasLimit-uint256-"></a>

```
UpdatedProtocolFlatFee(uint256 protocolFlatFee) (event)
```

<a name="IZRC20-UpdatedProtocolFlatFee-uint256-"></a>

