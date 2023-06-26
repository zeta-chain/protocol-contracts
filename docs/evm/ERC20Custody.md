# evm/ERC20Custody.md

## ERC20Custody

ERC20Custody for depositing ERC20 assets into ZetaChain and making operations with them.

### NotWhitelisted

```solidity
error NotWhitelisted()
```

### NotPaused

```solidity
error NotPaused()
```

### InvalidSender

```solidity
error InvalidSender()
```

### InvalidTSSUpdater

```solidity
error InvalidTSSUpdater()
```

### ZeroAddress

```solidity
error ZeroAddress()
```

### IsPaused

```solidity
error IsPaused()
```

### ZetaMaxFeeExceeded

```solidity
error ZetaMaxFeeExceeded()
```

### ZeroFee

```solidity
error ZeroFee()
```

### paused

```solidity
bool paused
```

If custody operations are paused.

### TSSAddress

```solidity
address TSSAddress
```

TSSAddress is the TSS address collectively possessed by Zeta blockchain validators.

### TSSAddressUpdater

```solidity
address TSSAddressUpdater
```

Threshold Signature Scheme (TSS) [GG20] is a multi-sig ECDSA/EdDSA protocol.

### zetaFee

```solidity
uint256 zetaFee
```

Current zeta fee for depositing funds into ZetaChain.

### zetaMaxFee

```solidity
uint256 zetaMaxFee
```

Maximum zeta fee for transaction.

### zeta

```solidity
contract IERC20 zeta
```

Zeta ERC20 token .

### whitelisted

```solidity
mapping(contract IERC20 => bool) whitelisted
```

Mapping of whitelisted token => true/false.

### Paused

```solidity
event Paused(address sender)
```

### Unpaused

```solidity
event Unpaused(address sender)
```

### Whitelisted

```solidity
event Whitelisted(contract IERC20 asset)
```

### Unwhitelisted

```solidity
event Unwhitelisted(contract IERC20 asset)
```

### Deposited

```solidity
event Deposited(bytes recipient, contract IERC20 asset, uint256 amount, bytes message)
```

### Withdrawn

```solidity
event Withdrawn(address recipient, contract IERC20 asset, uint256 amount)
```

### RenouncedTSSUpdater

```solidity
event RenouncedTSSUpdater(address TSSAddressUpdater_)
```

### UpdatedTSSAddress

```solidity
event UpdatedTSSAddress(address TSSAddress_)
```

### UpdatedZetaFee

```solidity
event UpdatedZetaFee(uint256 zetaFee_)
```

### onlyTSS

```solidity
modifier onlyTSS()
```

_Only TSS address allowed modifier._

### onlyTSSUpdater

```solidity
modifier onlyTSSUpdater()
```

_Only TSS address updater allowed modifier._

### constructor

```solidity
constructor(address TSSAddress_, address TSSAddressUpdater_, uint256 zetaFee_, uint256 zetaMaxFee_, contract IERC20 zeta_) public
```

### updateTSSAddress

```solidity
function updateTSSAddress(address TSSAddress_) external
```

_Update the TSSAddress in case of Zeta blockchain validator nodes churn._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| TSSAddress_ | address |  |

### updateZetaFee

```solidity
function updateZetaFee(uint256 zetaFee_) external
```

_Update zeta fee_

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| zetaFee_ | uint256 |  |

### renounceTSSAddressUpdater

```solidity
function renounceTSSAddressUpdater() external
```

_Change the ownership of TSSAddressUpdater to the Zeta blockchain TSS nodes.
Effectively, only Zeta blockchain validators collectively can update TSSAddress afterwards._

### pause

```solidity
function pause() external
```

_Pause custody operations._

### unpause

```solidity
function unpause() external
```

_Unpause custody operations._

### whitelist

```solidity
function whitelist(contract IERC20 asset) external
```

_Whitelist asset._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| asset | contract IERC20 |  |

### unwhitelist

```solidity
function unwhitelist(contract IERC20 asset) external
```

_Unwhitelist asset._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| asset | contract IERC20 |  |

### deposit

```solidity
function deposit(bytes recipient, contract IERC20 asset, uint256 amount, bytes message) external
```

_Deposit asset amount to recipient with message that encodes additional zetachain evm call or message._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| recipient | bytes | address. |
| asset | contract IERC20 | amount. |
| amount | uint256 |  |
| message | bytes |  |

### withdraw

```solidity
function withdraw(address recipient, contract IERC20 asset, uint256 amount) external
```

_Withdraw asset amount to recipient by custody TSS owner._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| recipient | address | address. |
| asset | contract IERC20 | amount. |
| amount | uint256 |  |

