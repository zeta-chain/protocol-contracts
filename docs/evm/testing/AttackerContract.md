## Victim

```solidity
import "@zetachain/protocol-contracts/contracts/evm/testing/AttackerContract.sol";
```

Source: https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/evm/testing/AttackerContract.sol

### Function List

* [deposit(recipient, asset, amount, message)](#Victim-deposit-bytes-address-uint256-bytes-)
* [withdraw(recipient, asset, amount)](#Victim-withdraw-address-address-uint256-)

### Modifiers

### Functions

```
deposit(bytes recipient, address asset, uint256 amount, bytes message) (external function)
```

<a name="Victim-deposit-bytes-address-uint256-bytes-"></a>

```
withdraw(address recipient, address asset, uint256 amount) (external function)
```

<a name="Victim-withdraw-address-address-uint256-"></a>

## AttackerContract

```solidity
import "@zetachain/protocol-contracts/contracts/evm/testing/AttackerContract.sol";
```

Source: https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/evm/testing/AttackerContract.sol

### Function List

* [constructor(victimContractAddress_, wzeta, victimMethod)](#AttackerContract-constructor-address-address-uint256-)
* [receive()](#AttackerContract-receive--)
* [attackDeposit()](#AttackerContract-attackDeposit--)
* [attackWidrawal()](#AttackerContract-attackWidrawal--)
* [attack()](#AttackerContract-attack--)
* [balanceOf(account)](#AttackerContract-balanceOf-address-)
* [transferFrom(from, to, amount)](#AttackerContract-transferFrom-address-address-uint256-)
* [transfer(to, amount)](#AttackerContract-transfer-address-uint256-)

### Modifiers

### Functions

```
constructor(address victimContractAddress_, address wzeta, uint256 victimMethod) (public function)
```

<a name="AttackerContract-constructor-address-address-uint256-"></a>

```
receive() (external function)
```

<a name="AttackerContract-receive--"></a>

```
attackDeposit() (internal function)
```

<a name="AttackerContract-attackDeposit--"></a>

```
attackWidrawal() (internal function)
```

<a name="AttackerContract-attackWidrawal--"></a>

```
attack() (internal function)
```

<a name="AttackerContract-attack--"></a>

```
balanceOf(address account) → uint256 (external function)
```

<a name="AttackerContract-balanceOf-address-"></a>

```
transferFrom(address from, address to, uint256 amount) → bool (public function)
```

<a name="AttackerContract-transferFrom-address-address-uint256-"></a>

```
transfer(address to, uint256 amount) → bool (public function)
```

<a name="AttackerContract-transfer-address-uint256-"></a>

