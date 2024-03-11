## ZetaNonEthInterface

```solidity
import "@zetachain/protocol-contracts/contracts/evm/interfaces/ZetaNonEthInterface.sol";
```

Source: https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/evm/interfaces/ZetaNonEthInterface.sol

ZetaNonEthInterface is a mintable / burnable version of IERC20

### Function List

* [burnFrom(account, amount)](#ZetaNonEthInterface-burnFrom-address-uint256-)
* [mint(mintee, value, internalSendHash)](#ZetaNonEthInterface-mint-address-uint256-bytes32-)

### Event List

* [Transfer(from, to, value)](#IERC20-Transfer-address-address-uint256-)
* [Approval(owner, spender, value)](#IERC20-Approval-address-address-uint256-)

### Modifiers

### Functions

```
burnFrom(address account, uint256 amount) (external function)
```

<a name="ZetaNonEthInterface-burnFrom-address-uint256-"></a>

```
mint(address mintee, uint256 value, bytes32 internalSendHash) (external function)
```

<a name="ZetaNonEthInterface-mint-address-uint256-bytes32-"></a>

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

