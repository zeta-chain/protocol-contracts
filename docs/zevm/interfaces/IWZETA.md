## IWETH9

```solidity
import "@zetachain/protocol-contracts/contracts/zevm/interfaces/IWZETA.sol";
```

Source: https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/zevm/interfaces/IWZETA.sol

### Function List

* [totalSupply()](#IWETH9-totalSupply--)
* [balanceOf(owner)](#IWETH9-balanceOf-address-)
* [allowance(owner, spender)](#IWETH9-allowance-address-address-)
* [approve(spender, wad)](#IWETH9-approve-address-uint256-)
* [transfer(to, wad)](#IWETH9-transfer-address-uint256-)
* [transferFrom(from, to, wad)](#IWETH9-transferFrom-address-address-uint256-)
* [deposit()](#IWETH9-deposit--)
* [withdraw(wad)](#IWETH9-withdraw-uint256-)

### Event List

* [Approval(owner, spender, value)](#IWETH9-Approval-address-address-uint256-)
* [Transfer(from, to, value)](#IWETH9-Transfer-address-address-uint256-)
* [Deposit(dst, wad)](#IWETH9-Deposit-address-uint256-)
* [Withdrawal(src, wad)](#IWETH9-Withdrawal-address-uint256-)

### Modifiers

### Functions

```
totalSupply() → uint256 (external function)
```

<a name="IWETH9-totalSupply--"></a>

```
balanceOf(address owner) → uint256 (external function)
```

<a name="IWETH9-balanceOf-address-"></a>

```
allowance(address owner, address spender) → uint256 (external function)
```

<a name="IWETH9-allowance-address-address-"></a>

```
approve(address spender, uint256 wad) → bool (external function)
```

<a name="IWETH9-approve-address-uint256-"></a>

```
transfer(address to, uint256 wad) → bool (external function)
```

<a name="IWETH9-transfer-address-uint256-"></a>

```
transferFrom(address from, address to, uint256 wad) → bool (external function)
```

<a name="IWETH9-transferFrom-address-address-uint256-"></a>

```
deposit() (external function)
```

<a name="IWETH9-deposit--"></a>

```
withdraw(uint256 wad) (external function)
```

<a name="IWETH9-withdraw-uint256-"></a>

### Events

```
Approval(address indexed owner, address indexed spender, uint256 value) (event)
```

<a name="IWETH9-Approval-address-address-uint256-"></a>

```
Transfer(address indexed from, address indexed to, uint256 value) (event)
```

<a name="IWETH9-Transfer-address-address-uint256-"></a>

```
Deposit(address indexed dst, uint256 wad) (event)
```

<a name="IWETH9-Deposit-address-uint256-"></a>

```
Withdrawal(address indexed src, uint256 wad) (event)
```

<a name="IWETH9-Withdrawal-address-uint256-"></a>

