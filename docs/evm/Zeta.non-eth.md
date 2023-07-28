## ZetaNonEth

```solidity
import "@zetachain/protocol-contracts/contracts/evm/Zeta.non-eth.sol";
```

Source: https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/evm/Zeta.non-eth.sol

### Function List

* [constructor(tssAddress_, tssAddressUpdater_)](#ZetaNonEth-constructor-address-address-)
* [updateTssAndConnectorAddresses(tssAddress_, connectorAddress_)](#ZetaNonEth-updateTssAndConnectorAddresses-address-address-)
* [renounceTssAddressUpdater()](#ZetaNonEth-renounceTssAddressUpdater--)
* [mint(mintee, value, internalSendHash)](#ZetaNonEth-mint-address-uint256-bytes32-)
* [burnFrom(account, amount)](#ZetaNonEth-burnFrom-address-uint256-)

### Event List

* [Minted(mintee, amount, internalSendHash)](#ZetaNonEth-Minted-address-uint256-bytes32-)
* [Burnt(burnee, amount)](#ZetaNonEth-Burnt-address-uint256-)

* [Transfer(from, to, value)](#IERC20-Transfer-address-address-uint256-)
* [Approval(owner, spender, value)](#IERC20-Approval-address-address-uint256-)

### Error List

* [CallerIsNotTss(caller)](#ZetaErrors-CallerIsNotTss-address-)
* [CallerIsNotConnector(caller)](#ZetaErrors-CallerIsNotConnector-address-)
* [CallerIsNotTssUpdater(caller)](#ZetaErrors-CallerIsNotTssUpdater-address-)
* [CallerIsNotTssOrUpdater(caller)](#ZetaErrors-CallerIsNotTssOrUpdater-address-)
* [InvalidAddress()](#ZetaErrors-InvalidAddress--)
* [ZetaTransferError()](#ZetaErrors-ZetaTransferError--)

### Modifiers

### Functions

```
constructor(address tssAddress_, address tssAddressUpdater_) (public function)
```

<a name="ZetaNonEth-constructor-address-address-"></a>

```
updateTssAndConnectorAddresses(address tssAddress_, address connectorAddress_) (external function)
```

<a name="ZetaNonEth-updateTssAndConnectorAddresses-address-address-"></a>

```
renounceTssAddressUpdater() (external function)
```

<a name="ZetaNonEth-renounceTssAddressUpdater--"></a>

Sets tssAddressUpdater to be tssAddress

```
mint(address mintee, uint256 value, bytes32 internalSendHash) (external function)
```

<a name="ZetaNonEth-mint-address-uint256-bytes32-"></a>

```
burnFrom(address account, uint256 amount) (public function)
```

<a name="ZetaNonEth-burnFrom-address-uint256-"></a>

### Events

```
Minted(address indexed mintee, uint256 amount, bytes32 indexed internalSendHash) (event)
```

<a name="ZetaNonEth-Minted-address-uint256-bytes32-"></a>

```
Burnt(address indexed burnee, uint256 amount) (event)
```

<a name="ZetaNonEth-Burnt-address-uint256-"></a>

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

### Errors

```
CallerIsNotTss(address caller) (error)
```

<a name="ZetaErrors-CallerIsNotTss-address-"></a>

```
CallerIsNotConnector(address caller) (error)
```

<a name="ZetaErrors-CallerIsNotConnector-address-"></a>

```
CallerIsNotTssUpdater(address caller) (error)
```

<a name="ZetaErrors-CallerIsNotTssUpdater-address-"></a>

```
CallerIsNotTssOrUpdater(address caller) (error)
```

<a name="ZetaErrors-CallerIsNotTssOrUpdater-address-"></a>

```
InvalidAddress() (error)
```

<a name="ZetaErrors-InvalidAddress--"></a>

```
ZetaTransferError() (error)
```

<a name="ZetaErrors-ZetaTransferError--"></a>

