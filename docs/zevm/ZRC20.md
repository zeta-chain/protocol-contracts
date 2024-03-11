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

* [_msgSender()](#ZRC20-_msgSender--)
* [_msgData()](#ZRC20-_msgData--)
* [constructor(name_, symbol_, decimals_, chainid_, coinType_, gasLimit_, systemContractAddress_)](#ZRC20-constructor-string-string-uint8-uint256-enum-CoinType-uint256-address-)
* [name()](#ZRC20-name--)
* [symbol()](#ZRC20-symbol--)
* [decimals()](#ZRC20-decimals--)
* [totalSupply()](#ZRC20-totalSupply--)
* [balanceOf(account)](#ZRC20-balanceOf-address-)
* [transfer(recipient, amount)](#ZRC20-transfer-address-uint256-)
* [allowance(owner, spender)](#ZRC20-allowance-address-address-)
* [approve(spender, amount)](#ZRC20-approve-address-uint256-)
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
_msgSender() → address (internal function)
```

<a name="ZRC20-_msgSender--"></a>

```
_msgData() → bytes (internal function)
```

<a name="ZRC20-_msgData--"></a>

```
constructor(string name_, string symbol_, uint8 decimals_, uint256 chainid_, enum CoinType coinType_, uint256 gasLimit_, address systemContractAddress_) (public function)
```

<a name="ZRC20-constructor-string-string-uint8-uint256-enum-CoinType-uint256-address-"></a>

The only one allowed to deploy new ZRC20 is fungible address.

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

Returns ZRC20 balance of an account.

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

```
_mint(address account, uint256 amount) (internal function)
```

<a name="ZRC20-_mint-address-uint256-"></a>

```
_burn(address account, uint256 amount) (internal function)
```

<a name="ZRC20-_burn-address-uint256-"></a>

```
_approve(address owner, address spender, uint256 amount) (internal function)
```

<a name="ZRC20-_approve-address-address-uint256-"></a>

```
deposit(address to, uint256 amount) → bool (external function)
```

<a name="ZRC20-deposit-address-uint256-"></a>

Deposits corresponding tokens from external chain, only callable by Fungible module.

```
withdrawGasFee() → address, uint256 (public function)
```

<a name="ZRC20-withdrawGasFee--"></a>

Withdraws gas fees.

```
withdraw(bytes to, uint256 amount) → bool (external function)
```

<a name="ZRC20-withdraw-bytes-uint256-"></a>

Withraws ZRC20 tokens to external chains, this function causes cctx module to send out outbound tx to the outbound chain
this contract should be given enough allowance of the gas ZRC20 to pay for outbound tx gas fee.

```
updateSystemContractAddress(address addr) (external function)
```

<a name="ZRC20-updateSystemContractAddress-address-"></a>

Updates system contract address. Can only be updated by the fungible module.

```
updateGasLimit(uint256 gasLimit) (external function)
```

<a name="ZRC20-updateGasLimit-uint256-"></a>

Updates gas limit. Can only be updated by the fungible module.

```
updateProtocolFlatFee(uint256 protocolFlatFee) (external function)
```

<a name="ZRC20-updateProtocolFlatFee-uint256-"></a>

Updates protocol flat fee. Can only be updated by the fungible module.

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

