## WETH9

```solidity
import "@zetachain/protocol-contracts/contracts/zevm/WZETA.sol";
```

Source: https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/zevm/WZETA.sol

### Function List

* [()](#WETH9-undefined--)
* [deposit()](#WETH9-deposit--)
* [withdraw(wad)](#WETH9-withdraw-uint256-)
* [totalSupply()](#WETH9-totalSupply--)
* [approve(guy, wad)](#WETH9-approve-address-uint256-)
* [transfer(dst, wad)](#WETH9-transfer-address-uint256-)
* [transferFrom(src, dst, wad)](#WETH9-transferFrom-address-address-uint256-)

### Event List

* [Approval(src, guy, wad)](#WETH9-Approval-address-address-uint256-)
* [Transfer(src, dst, wad)](#WETH9-Transfer-address-address-uint256-)
* [Deposit(dst, wad)](#WETH9-Deposit-address-uint256-)
* [Withdrawal(src, wad)](#WETH9-Withdrawal-address-uint256-)

### Modifiers

### Functions

```
() (public function)
```

<a name="WETH9-undefined--"></a>

```
deposit() (public function)
```

<a name="WETH9-deposit--"></a>

```
withdraw(uint256 wad) (public function)
```

<a name="WETH9-withdraw-uint256-"></a>

```
totalSupply() → uint256 (public function)
```

<a name="WETH9-totalSupply--"></a>

```
approve(address guy, uint256 wad) → bool (public function)
```

<a name="WETH9-approve-address-uint256-"></a>

```
transfer(address dst, uint256 wad) → bool (public function)
```

<a name="WETH9-transfer-address-uint256-"></a>

```
transferFrom(address src, address dst, uint256 wad) → bool (public function)
```

<a name="WETH9-transferFrom-address-address-uint256-"></a>

### Events

```
Approval(address indexed src, address indexed guy, uint256 wad) (event)
```

<a name="WETH9-Approval-address-address-uint256-"></a>

```
Transfer(address indexed src, address indexed dst, uint256 wad) (event)
```

<a name="WETH9-Transfer-address-address-uint256-"></a>

```
Deposit(address indexed dst, uint256 wad) (event)
```

<a name="WETH9-Deposit-address-uint256-"></a>

```
Withdrawal(address indexed src, uint256 wad) (event)
```

<a name="WETH9-Withdrawal-address-uint256-"></a>

