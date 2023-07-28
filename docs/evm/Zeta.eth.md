## ZetaEth

```solidity
import "@zetachain/protocol-contracts/contracts/evm/Zeta.eth.sol";
```

Source: https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/evm/Zeta.eth.sol

ZetaEth is an implementation of OpenZeppelin's ERC20

### Function List

* [constructor(creator, initialSupply)](#ZetaEth-constructor-address-uint256-)

### Event List

* [Transfer(from, to, value)](#IERC20-Transfer-address-address-uint256-)
* [Approval(owner, spender, value)](#IERC20-Approval-address-address-uint256-)

### Modifiers

### Functions

```
constructor(address creator, uint256 initialSupply) (public function)
```

<a name="ZetaEth-constructor-address-uint256-"></a>

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

