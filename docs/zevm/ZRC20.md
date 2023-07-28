## ZRC20Errors

```solidity
import "@zetachain/protocol-contracts/contracts/zevm/ZRC20.sol";
```

Source: https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/zevm/ZRC20.sol

Custom errors for ZRC20

### Error List

* [CallerIsNotFungibleModule()](#ZRC20Errors-CallerIsNotFungibleModule--)
* [InvalidSender()](#ZRC20Errors-InvalidSender--)
* [GasFeeTransferFailed()](#ZRC20Errors-GasFeeTransferFailed--)
* [ZeroGasCoin()](#ZRC20Errors-ZeroGasCoin--)
* [ZeroGasPrice()](#ZRC20Errors-ZeroGasPrice--)
* [LowAllowance()](#ZRC20Errors-LowAllowance--)
* [LowBalance()](#ZRC20Errors-LowBalance--)
* [ZeroAddress()](#ZRC20Errors-ZeroAddress--)

### Modifiers

### Errors

```
CallerIsNotFungibleModule() (error)
```

<a name="ZRC20Errors-CallerIsNotFungibleModule--"></a>

```
InvalidSender() (error)
```

<a name="ZRC20Errors-InvalidSender--"></a>

```
GasFeeTransferFailed() (error)
```

<a name="ZRC20Errors-GasFeeTransferFailed--"></a>

```
ZeroGasCoin() (error)
```

<a name="ZRC20Errors-ZeroGasCoin--"></a>

```
ZeroGasPrice() (error)
```

<a name="ZRC20Errors-ZeroGasPrice--"></a>

```
LowAllowance() (error)
```

<a name="ZRC20Errors-LowAllowance--"></a>

```
LowBalance() (error)
```

<a name="ZRC20Errors-LowBalance--"></a>

```
ZeroAddress() (error)
```

<a name="ZRC20Errors-ZeroAddress--"></a>

## ZRC20

```solidity
import "@zetachain/protocol-contracts/contracts/zevm/ZRC20.sol";
```

Source: https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/zevm/ZRC20.sol

### Modifier List

* [onlyFungible()](#ZRC20-onlyFungible--)

### Function List

* [constructor(name_, symbol_, decimals_, chainid_, coinType_, gasLimit_, systemContractAddress_)](#ZRC20-constructor-string-string-uint8-uint256-enum-CoinType-uint256-address-)
* [name()](#ZRC20-name--)
* [symbol()](#ZRC20-symbol--)
* [decimals()](#ZRC20-decimals--)
* [totalSupply()](#ZRC20-totalSupply--)
* [balanceOf(account)](#ZRC20-balanceOf-address-)
* [transfer(recipient, amount)](#ZRC20-transfer-address-uint256-)
* [allowance(owner, spender)](#ZRC20-allowance-address-address-)
* [approve(spender, amount)](#ZRC20-approve-address-uint256-)
* [increaseAllowance(spender, amount)](#ZRC20-increaseAllowance-address-uint256-)
* [decreaseAllowance(spender, amount)](#ZRC20-decreaseAllowance-address-uint256-)
* [transferFrom(sender, recipient, amount)](#ZRC20-transferFrom-address-address-uint256-)
* [burn(amount)](#ZRC20-burn-uint256-)
* [_transfer(sender, recipient, amount)](#ZRC20-_transfer-address-address-uint256-)
* [_mint(account, amount)](#ZRC20-_mint-address-uint256-)
* [_burn(account, amount)](#ZRC20-_burn-address-uint256-)
* [_approve(owner, spender, amount)](#ZRC20-_approve-address-address-uint256-)
* [deposit(to, amount)](#ZRC20-deposit-address-uint256-)
* [withdrawGasFee()](#ZRC20-withdrawGasFee--)
* [withdraw(to, amount)](#ZRC20-withdraw-bytes-uint256-)
* [updateSystemContractAddress(addr)](#ZRC20-updateSystemContractAddress-address-)
* [updateGasLimit(gasLimit)](#ZRC20-updateGasLimit-uint256-)
* [updateProtocolFlatFee(protocolFlatFee)](#ZRC20-updateProtocolFlatFee-uint256-)

### Event List

* [Transfer(from, to, value)](#IZRC20-Transfer-address-address-uint256-)
* [Approval(owner, spender, value)](#IZRC20-Approval-address-address-uint256-)
* [Deposit(from, to, value)](#IZRC20-Deposit-bytes-address-uint256-)
* [Withdrawal(from, to, value, gasfee, protocolFlatFee)](#IZRC20-Withdrawal-address-bytes-uint256-uint256-uint256-)
* [UpdatedSystemContract(systemContract)](#IZRC20-UpdatedSystemContract-address-)
* [UpdatedGasLimit(gasLimit)](#IZRC20-UpdatedGasLimit-uint256-)
* [UpdatedProtocolFlatFee(protocolFlatFee)](#IZRC20-UpdatedProtocolFlatFee-uint256-)

### Error List

* [CallerIsNotFungibleModule()](#ZRC20Errors-CallerIsNotFungibleModule--)
* [InvalidSender()](#ZRC20Errors-InvalidSender--)
* [GasFeeTransferFailed()](#ZRC20Errors-GasFeeTransferFailed--)
* [ZeroGasCoin()](#ZRC20Errors-ZeroGasCoin--)
* [ZeroGasPrice()](#ZRC20Errors-ZeroGasPrice--)
* [LowAllowance()](#ZRC20Errors-LowAllowance--)
* [LowBalance()](#ZRC20Errors-LowBalance--)
* [ZeroAddress()](#ZRC20Errors-ZeroAddress--)

### Modifiers

```
onlyFungible() (modifier)
```

<a name="ZRC20-onlyFungible--"></a>

Only fungible module modifier.

### Functions

```
constructor(string name_, string symbol_, uint8 decimals_, uint256 chainid_, enum CoinType coinType_, uint256 gasLimit_, address systemContractAddress_) (public function)
```

<a name="ZRC20-constructor-string-string-uint8-uint256-enum-CoinType-uint256-address-"></a>

Constructor that gives msg.sender all of existing tokens.

```
name() → string (public function)
```

<a name="ZRC20-name--"></a>

ZRC20 name

```
symbol() → string (public function)
```

<a name="ZRC20-symbol--"></a>

ZRC20 symbol.

```
decimals() → uint8 (public function)
```

<a name="ZRC20-decimals--"></a>

ZRC20 decimals.

```
totalSupply() → uint256 (public function)
```

<a name="ZRC20-totalSupply--"></a>

ZRC20 total supply.

```
balanceOf(address account) → uint256 (public function)
```

<a name="ZRC20-balanceOf-address-"></a>

Returns ZRC20 balance of an account.

```
transfer(address recipient, uint256 amount) → bool (public function)
```

<a name="ZRC20-transfer-address-uint256-"></a>

This function can be called by the contract owner or any other external address.

```
allowance(address owner, address spender) → uint256 (public function)
```

<a name="ZRC20-allowance-address-address-"></a>

Returns token allowance from owner to spender.

```
approve(address spender, uint256 amount) → bool (public function)
```

<a name="ZRC20-approve-address-uint256-"></a>

Approves amount transferFrom for spender.

```
increaseAllowance(address spender, uint256 amount) → bool (external function)
```

<a name="ZRC20-increaseAllowance-address-uint256-"></a>

Increases allowance by amount for spender.

```
decreaseAllowance(address spender, uint256 amount) → bool (external function)
```

<a name="ZRC20-decreaseAllowance-address-uint256-"></a>

Decreases allowance by amount for spender.

```
transferFrom(address sender, address recipient, uint256 amount) → bool (public function)
```

<a name="ZRC20-transferFrom-address-address-uint256-"></a>

Transfers tokens from sender to recipient.

```
burn(uint256 amount) → bool (external function)
```

<a name="ZRC20-burn-uint256-"></a>

Burns an amount of tokens.

```
_transfer(address sender, address recipient, uint256 amount) (internal function)
```

<a name="ZRC20-_transfer-address-address-uint256-"></a>

Internal function to transfer tokens from one address to another.
Throws if either the sender or recipient address is zero.
Throws if the sender's balance is lower than the transfer amount.

```
_mint(address account, uint256 amount) (internal function)
```

<a name="ZRC20-_mint-address-uint256-"></a>

Internal function to mint new tokens and assign them to an account.
Throws if the account address is zero.

```
_burn(address account, uint256 amount) (internal function)
```

<a name="ZRC20-_burn-address-uint256-"></a>

Internal function to burn tokens from an account.
Throws if the account address is zero.
Throws if the account's balance is lower than the burn amount.

```
_approve(address owner, address spender, uint256 amount) (internal function)
```

<a name="ZRC20-_approve-address-address-uint256-"></a>

Internal function to approve a spender to spend tokens on behalf of the owner.
Throws if the owner address is zero.
Throws if the spender address is zero.

```
deposit(address to, uint256 amount) → bool (external function)
```

<a name="ZRC20-deposit-address-uint256-"></a>

Deposits corresponding tokens from an external chain.
Only callable by the Fungible module or the System contract.
Throws if called by an invalid sender.

```
withdrawGasFee() → address, uint256 (public function)
```

<a name="ZRC20-withdrawGasFee--"></a>

Returns the ZRC20 address for gas on the same chain of this ZRC20, and calculates the gas fee for `withdraw()`.
Throws if the gas ZRC20 address is zero.
Throws if the gas price is zero.

```
withdraw(bytes to, uint256 amount) → bool (external function)
```

<a name="ZRC20-withdraw-bytes-uint256-"></a>

Withdraws ZRC20 tokens to external chains by triggering the crosschain module to create an outbound transaction.
Requires this contract to have sufficient allowance of the gas ZRC20 token to pay for the outbound transaction gas fee.
Throws if the gas fee transfer fails.

```
updateSystemContractAddress(address addr) (external function)
```

<a name="ZRC20-updateSystemContractAddress-address-"></a>

Updates the system contract address.
Requires the caller to be the fungible module.

```
updateGasLimit(uint256 gasLimit) (external function)
```

<a name="ZRC20-updateGasLimit-uint256-"></a>

Updates the gas limit.
Requires the caller to be the fungible module.

```
updateProtocolFlatFee(uint256 protocolFlatFee) (external function)
```

<a name="ZRC20-updateProtocolFlatFee-uint256-"></a>

Updates the protocol flat fee.
Requires the caller to be the fungible module.

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
Withdrawal(address indexed from, bytes to, uint256 value, uint256 gasfee, uint256 protocolFlatFee) (event)
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

### Errors

```
CallerIsNotFungibleModule() (error)
```

<a name="ZRC20Errors-CallerIsNotFungibleModule--"></a>

```
InvalidSender() (error)
```

<a name="ZRC20Errors-InvalidSender--"></a>

```
GasFeeTransferFailed() (error)
```

<a name="ZRC20Errors-GasFeeTransferFailed--"></a>

```
ZeroGasCoin() (error)
```

<a name="ZRC20Errors-ZeroGasCoin--"></a>

```
ZeroGasPrice() (error)
```

<a name="ZRC20Errors-ZeroGasPrice--"></a>

```
LowAllowance() (error)
```

<a name="ZRC20Errors-LowAllowance--"></a>

```
LowBalance() (error)
```

<a name="ZRC20Errors-LowBalance--"></a>

```
ZeroAddress() (error)
```

<a name="ZRC20Errors-ZeroAddress--"></a>

