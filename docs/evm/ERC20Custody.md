## ERC20Custody

```solidity
import "@zetachain/protocol-contracts/contracts/evm/ERC20Custody.sol";
```

Source: https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/evm/ERC20Custody.sol

### Modifier List

* [onlyTSS()](#ERC20Custody-onlyTSS--)
* [onlyTSSUpdater()](#ERC20Custody-onlyTSSUpdater--)

### Function List

* [constructor(TSSAddress_, TSSAddressUpdater_, zetaFee_, zetaMaxFee_, zeta_)](#ERC20Custody-constructor-address-address-uint256-uint256-contract-IERC20-)
* [updateTSSAddress(TSSAddress_)](#ERC20Custody-updateTSSAddress-address-)
* [updateZetaFee(zetaFee_)](#ERC20Custody-updateZetaFee-uint256-)
* [renounceTSSAddressUpdater()](#ERC20Custody-renounceTSSAddressUpdater--)
* [pause()](#ERC20Custody-pause--)
* [unpause()](#ERC20Custody-unpause--)
* [whitelist(asset)](#ERC20Custody-whitelist-contract-IERC20-)
* [unwhitelist(asset)](#ERC20Custody-unwhitelist-contract-IERC20-)
* [deposit(recipient, asset, amount, message)](#ERC20Custody-deposit-bytes-contract-IERC20-uint256-bytes-)
* [withdraw(recipient, asset, amount)](#ERC20Custody-withdraw-address-contract-IERC20-uint256-)

### Event List

* [Paused(sender)](#ERC20Custody-Paused-address-)
* [Unpaused(sender)](#ERC20Custody-Unpaused-address-)
* [Whitelisted(asset)](#ERC20Custody-Whitelisted-contract-IERC20-)
* [Unwhitelisted(asset)](#ERC20Custody-Unwhitelisted-contract-IERC20-)
* [Deposited(recipient, asset, amount, message)](#ERC20Custody-Deposited-bytes-contract-IERC20-uint256-bytes-)
* [Withdrawn(recipient, asset, amount)](#ERC20Custody-Withdrawn-address-contract-IERC20-uint256-)
* [RenouncedTSSUpdater(TSSAddressUpdater_)](#ERC20Custody-RenouncedTSSUpdater-address-)
* [UpdatedTSSAddress(TSSAddress_)](#ERC20Custody-UpdatedTSSAddress-address-)
* [UpdatedZetaFee(zetaFee_)](#ERC20Custody-UpdatedZetaFee-uint256-)

### Error List

* [NotWhitelisted()](#ERC20Custody-NotWhitelisted--)
* [NotPaused()](#ERC20Custody-NotPaused--)
* [InvalidSender()](#ERC20Custody-InvalidSender--)
* [InvalidTSSUpdater()](#ERC20Custody-InvalidTSSUpdater--)
* [ZeroAddress()](#ERC20Custody-ZeroAddress--)
* [IsPaused()](#ERC20Custody-IsPaused--)
* [ZetaMaxFeeExceeded()](#ERC20Custody-ZetaMaxFeeExceeded--)
* [ZeroFee()](#ERC20Custody-ZeroFee--)

### Modifiers

```
onlyTSS() (modifier)
```

<a name="ERC20Custody-onlyTSS--"></a>

Only TSS address allowed modifier.

```
onlyTSSUpdater() (modifier)
```

<a name="ERC20Custody-onlyTSSUpdater--"></a>

Only TSS address updater allowed modifier.

### Functions

```
constructor(address TSSAddress_, address TSSAddressUpdater_, uint256 zetaFee_, uint256 zetaMaxFee_, contract IERC20 zeta_) (public function)
```

<a name="ERC20Custody-constructor-address-address-uint256-uint256-contract-IERC20-"></a>

```
updateTSSAddress(address TSSAddress_) (external function)
```

<a name="ERC20Custody-updateTSSAddress-address-"></a>

Update the TSSAddress in case of Zeta blockchain validator nodes churn.

```
updateZetaFee(uint256 zetaFee_) (external function)
```

<a name="ERC20Custody-updateZetaFee-uint256-"></a>

Update zeta fee

```
renounceTSSAddressUpdater() (external function)
```

<a name="ERC20Custody-renounceTSSAddressUpdater--"></a>

Change the ownership of TSSAddressUpdater to the Zeta blockchain TSS nodes.
Effectively, only Zeta blockchain validators collectively can update TSSAddress afterwards.

```
pause() (external function)
```

<a name="ERC20Custody-pause--"></a>

Pause custody operations.

```
unpause() (external function)
```

<a name="ERC20Custody-unpause--"></a>

Unpause custody operations.

```
whitelist(contract IERC20 asset) (external function)
```

<a name="ERC20Custody-whitelist-contract-IERC20-"></a>

Whitelist asset.

```
unwhitelist(contract IERC20 asset) (external function)
```

<a name="ERC20Custody-unwhitelist-contract-IERC20-"></a>

Unwhitelist asset.

```
deposit(bytes recipient, contract IERC20 asset, uint256 amount, bytes message) (external function)
```

<a name="ERC20Custody-deposit-bytes-contract-IERC20-uint256-bytes-"></a>

Deposit asset amount to recipient with message that encodes additional zetachain evm call or message.

```
withdraw(address recipient, contract IERC20 asset, uint256 amount) (external function)
```

<a name="ERC20Custody-withdraw-address-contract-IERC20-uint256-"></a>

Withdraw asset amount to recipient by custody TSS owner.

### Events

```
Paused(address sender) (event)
```

<a name="ERC20Custody-Paused-address-"></a>

```
Unpaused(address sender) (event)
```

<a name="ERC20Custody-Unpaused-address-"></a>

```
Whitelisted(contract IERC20 indexed asset) (event)
```

<a name="ERC20Custody-Whitelisted-contract-IERC20-"></a>

```
Unwhitelisted(contract IERC20 indexed asset) (event)
```

<a name="ERC20Custody-Unwhitelisted-contract-IERC20-"></a>

```
Deposited(bytes recipient, contract IERC20 indexed asset, uint256 amount, bytes message) (event)
```

<a name="ERC20Custody-Deposited-bytes-contract-IERC20-uint256-bytes-"></a>

```
Withdrawn(address indexed recipient, contract IERC20 indexed asset, uint256 amount) (event)
```

<a name="ERC20Custody-Withdrawn-address-contract-IERC20-uint256-"></a>

```
RenouncedTSSUpdater(address TSSAddressUpdater_) (event)
```

<a name="ERC20Custody-RenouncedTSSUpdater-address-"></a>

```
UpdatedTSSAddress(address TSSAddress_) (event)
```

<a name="ERC20Custody-UpdatedTSSAddress-address-"></a>

```
UpdatedZetaFee(uint256 zetaFee_) (event)
```

<a name="ERC20Custody-UpdatedZetaFee-uint256-"></a>

### Errors

```
NotWhitelisted() (error)
```

<a name="ERC20Custody-NotWhitelisted--"></a>

```
NotPaused() (error)
```

<a name="ERC20Custody-NotPaused--"></a>

```
InvalidSender() (error)
```

<a name="ERC20Custody-InvalidSender--"></a>

```
InvalidTSSUpdater() (error)
```

<a name="ERC20Custody-InvalidTSSUpdater--"></a>

```
ZeroAddress() (error)
```

<a name="ERC20Custody-ZeroAddress--"></a>

```
IsPaused() (error)
```

<a name="ERC20Custody-IsPaused--"></a>

```
ZetaMaxFeeExceeded() (error)
```

<a name="ERC20Custody-ZetaMaxFeeExceeded--"></a>

```
ZeroFee() (error)
```

<a name="ERC20Custody-ZeroFee--"></a>

