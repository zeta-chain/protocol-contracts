## ISystem

```solidity
import "@zetachain/protocol-contracts/contracts/zevm/Interfaces.sol";
```

Source: https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/zevm/Interfaces.sol

Interfaces of SystemContract and ZRC20 to make easier to import.

### Function List

* [FUNGIBLE_MODULE_ADDRESS()](#ISystem-FUNGIBLE_MODULE_ADDRESS--)
* [wZetaContractAddress()](#ISystem-wZetaContractAddress--)
* [uniswapv2FactoryAddress()](#ISystem-uniswapv2FactoryAddress--)
* [gasPriceByChainId(chainID)](#ISystem-gasPriceByChainId-uint256-)
* [gasCoinZRC20ByChainId(chainID)](#ISystem-gasCoinZRC20ByChainId-uint256-)
* [gasZetaPoolByChainId(chainID)](#ISystem-gasZetaPoolByChainId-uint256-)

### Modifiers

### Functions

```
FUNGIBLE_MODULE_ADDRESS() → address (external function)
```

<a name="ISystem-FUNGIBLE_MODULE_ADDRESS--"></a>

```
wZetaContractAddress() → address (external function)
```

<a name="ISystem-wZetaContractAddress--"></a>

```
uniswapv2FactoryAddress() → address (external function)
```

<a name="ISystem-uniswapv2FactoryAddress--"></a>

```
gasPriceByChainId(uint256 chainID) → uint256 (external function)
```

<a name="ISystem-gasPriceByChainId-uint256-"></a>

```
gasCoinZRC20ByChainId(uint256 chainID) → address (external function)
```

<a name="ISystem-gasCoinZRC20ByChainId-uint256-"></a>

```
gasZetaPoolByChainId(uint256 chainID) → address (external function)
```

<a name="ISystem-gasZetaPoolByChainId-uint256-"></a>

## IZRC20

```solidity
import "@zetachain/protocol-contracts/contracts/zevm/Interfaces.sol";
```

Source: https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/zevm/Interfaces.sol

### Function List

* [totalSupply()](#IZRC20-totalSupply--)
* [balanceOf(account)](#IZRC20-balanceOf-address-)
* [transfer(recipient, amount)](#IZRC20-transfer-address-uint256-)
* [allowance(owner, spender)](#IZRC20-allowance-address-address-)
* [approve(spender, amount)](#IZRC20-approve-address-uint256-)
* [transferFrom(sender, recipient, amount)](#IZRC20-transferFrom-address-address-uint256-)
* [deposit(to, amount)](#IZRC20-deposit-address-uint256-)
* [withdraw(to, amount)](#IZRC20-withdraw-bytes-uint256-)
* [withdrawGasFee()](#IZRC20-withdrawGasFee--)

### Event List

* [Transfer(from, to, value)](#IZRC20-Transfer-address-address-uint256-)
* [Approval(owner, spender, value)](#IZRC20-Approval-address-address-uint256-)
* [Deposit(from, to, value)](#IZRC20-Deposit-bytes-address-uint256-)
* [Withdrawal(from, to, value, gasfee, protocolFlatFee)](#IZRC20-Withdrawal-address-bytes-uint256-uint256-uint256-)
* [UpdatedSystemContract(systemContract)](#IZRC20-UpdatedSystemContract-address-)
* [UpdatedGasLimit(gasLimit)](#IZRC20-UpdatedGasLimit-uint256-)
* [UpdatedProtocolFlatFee(protocolFlatFee)](#IZRC20-UpdatedProtocolFlatFee-uint256-)

### Modifiers

### Functions

```
totalSupply() → uint256 (external function)
```

<a name="IZRC20-totalSupply--"></a>

```
balanceOf(address account) → uint256 (external function)
```

<a name="IZRC20-balanceOf-address-"></a>

```
transfer(address recipient, uint256 amount) → bool (external function)
```

<a name="IZRC20-transfer-address-uint256-"></a>

```
allowance(address owner, address spender) → uint256 (external function)
```

<a name="IZRC20-allowance-address-address-"></a>

```
approve(address spender, uint256 amount) → bool (external function)
```

<a name="IZRC20-approve-address-uint256-"></a>

```
transferFrom(address sender, address recipient, uint256 amount) → bool (external function)
```

<a name="IZRC20-transferFrom-address-address-uint256-"></a>

```
deposit(address to, uint256 amount) → bool (external function)
```

<a name="IZRC20-deposit-address-uint256-"></a>

```
withdraw(bytes to, uint256 amount) → bool (external function)
```

<a name="IZRC20-withdraw-bytes-uint256-"></a>

```
withdrawGasFee() → address, uint256 (external function)
```

<a name="IZRC20-withdrawGasFee--"></a>

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

## IZRC20Metadata

```solidity
import "@zetachain/protocol-contracts/contracts/zevm/Interfaces.sol";
```

Source: https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/zevm/Interfaces.sol

### Function List

* [name()](#IZRC20Metadata-name--)
* [symbol()](#IZRC20Metadata-symbol--)
* [decimals()](#IZRC20Metadata-decimals--)

### Event List

* [Transfer(from, to, value)](#IZRC20-Transfer-address-address-uint256-)
* [Approval(owner, spender, value)](#IZRC20-Approval-address-address-uint256-)
* [Deposit(from, to, value)](#IZRC20-Deposit-bytes-address-uint256-)
* [Withdrawal(from, to, value, gasfee, protocolFlatFee)](#IZRC20-Withdrawal-address-bytes-uint256-uint256-uint256-)
* [UpdatedSystemContract(systemContract)](#IZRC20-UpdatedSystemContract-address-)
* [UpdatedGasLimit(gasLimit)](#IZRC20-UpdatedGasLimit-uint256-)
* [UpdatedProtocolFlatFee(protocolFlatFee)](#IZRC20-UpdatedProtocolFlatFee-uint256-)

### Modifiers

### Functions

```
name() → string (external function)
```

<a name="IZRC20Metadata-name--"></a>

```
symbol() → string (external function)
```

<a name="IZRC20Metadata-symbol--"></a>

```
decimals() → uint8 (external function)
```

<a name="IZRC20Metadata-decimals--"></a>

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

## CoinType

```solidity
enum CoinType {
  Zeta,
  Gas,
  ERC20
}
```
