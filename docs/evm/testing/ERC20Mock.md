## ERC20Mock

```solidity
import "@zetachain/protocol-contracts/contracts/evm/testing/ERC20Mock.sol";
```

Source: https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/evm/testing/ERC20Mock.sol

ZetaEth is an implementation of OpenZeppelin's ERC20

### Function List

* [constructor(name, symbol, creator, initialSupply)](#ERC20Mock-constructor-string-string-address-uint256-)

### Event List

* [Transfer(from, to, value)](#IERC20-Transfer-address-address-uint256-)
* [Approval(owner, spender, value)](#IERC20-Approval-address-address-uint256-)

### Modifiers

### Functions

```
constructor(string name, string symbol, address creator, uint256 initialSupply) (public function)
```

<a name="ERC20Mock-constructor-string-string-address-uint256-"></a>

### Events

```
Transfer(address indexed from, address indexed to, uint256 value) (event)
```

<a name="IERC20-Transfer-address-address-uint256-"></a>

Emitted when `value` tokens are moved from one account (`from`) to
another (`to`).

Note that `value` may be zero.

```
Approval(address indexed owner, address indexed spender, uint256 value) (event)
```

<a name="IERC20-Approval-address-address-uint256-"></a>

Emitted when the allowance of a `spender` for an `owner` is set by
a call to {approve}. `value` is the new allowance.

