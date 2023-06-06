# zevm/ZRC20.md

## ZRC20Errors

_Custom errors for ZRC20_

### CallerIsNotFungibleModule

```solidity
error CallerIsNotFungibleModule()
```

### InvalidSender

```solidity
error InvalidSender()
```

### GasFeeTransferFailed

```solidity
error GasFeeTransferFailed()
```

### ZeroGasCoin

```solidity
error ZeroGasCoin()
```

### ZeroGasPrice

```solidity
error ZeroGasPrice()
```

### LowAllowance

```solidity
error LowAllowance()
```

### LowBalance

```solidity
error LowBalance()
```

### ZeroAddress

```solidity
error ZeroAddress()
```

## ZRC20

### FUNGIBLE_MODULE_ADDRESS

```solidity
address FUNGIBLE_MODULE_ADDRESS
```

Fungible address is always the same, maintained at the protocol level

### CHAIN_ID

```solidity
uint256 CHAIN_ID
```

Chain id.abi

### COIN_TYPE

```solidity
enum CoinType COIN_TYPE
```

Coin type, checkout Interfaces.sol.

### SYSTEM_CONTRACT_ADDRESS

```solidity
address SYSTEM_CONTRACT_ADDRESS
```

System contract address.

### GAS_LIMIT

```solidity
uint256 GAS_LIMIT
```

Gas limit.

### PROTOCOL_FLAT_FEE

```solidity
uint256 PROTOCOL_FLAT_FEE
```

Protocol flat fee.

### onlyFungible

```solidity
modifier onlyFungible()
```

_Only fungible module modifier._

### constructor

```solidity
constructor(string name_, string symbol_, uint8 decimals_, uint256 chainid_, enum CoinType coinType_, uint256 gasLimit_, address systemContractAddress_) public
```

_The only one allowed to deploy new ZRC20 is fungible address._

### name

```solidity
function name() public view virtual returns (string)
```

_ZRC20 name_

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | string | name as string |

### symbol

```solidity
function symbol() public view virtual returns (string)
```

_ZRC20 symbol._

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | string | symbol as string. |

### decimals

```solidity
function decimals() public view virtual returns (uint8)
```

_ZRC20 decimals._

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | uint8 | returns uint8 decimals. |

### totalSupply

```solidity
function totalSupply() public view virtual returns (uint256)
```

_ZRC20 total supply._

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | uint256 | returns uint256 total supply. |

### balanceOf

```solidity
function balanceOf(address account) public view virtual returns (uint256)
```

_Returns ZRC20 balance of an account._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| account | address | address for which balance is requested. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | uint256 | uint256 account balance. |

### transfer

```solidity
function transfer(address recipient, uint256 amount) public virtual returns (bool)
```

_Returns ZRC20 balance of an account._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| recipient | address |  |
| amount | uint256 |  |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | bool | true/false if transfer succeeded/failed. |

### allowance

```solidity
function allowance(address owner, address spender) public view virtual returns (uint256)
```

_Returns token allowance from owner to spender._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| owner | address | address. |
| spender | address |  |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | uint256 | uint256 allowance. |

### approve

```solidity
function approve(address spender, uint256 amount) public virtual returns (bool)
```

_Approves amount transferFrom for spender._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| spender | address | address. |
| amount | uint256 | to approve. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | bool | true/false if succeeded/failed. |

### increaseAllowance

```solidity
function increaseAllowance(address spender, uint256 amount) external virtual returns (bool)
```

_Increases allowance by amount for spender._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| spender | address | address. |
| amount | uint256 | by which to increase allownace. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | bool | true/false if succeeded/failed. |

### decreaseAllowance

```solidity
function decreaseAllowance(address spender, uint256 amount) external virtual returns (bool)
```

_Decreases allowance by amount for spender._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| spender | address | address. |
| amount | uint256 | by which to decrease allownace. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | bool | true/false if succeeded/failed. |

### transferFrom

```solidity
function transferFrom(address sender, address recipient, uint256 amount) public virtual returns (bool)
```

_Transfers tokens from sender to recipient._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| sender | address | address. |
| recipient | address | address. |
| amount | uint256 | to transfer. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | bool | true/false if succeeded/failed. |

### burn

```solidity
function burn(uint256 amount) external returns (bool)
```

_Burns an amount of tokens._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| amount | uint256 | to burn. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | bool | true/false if succeeded/failed. |

### _transfer

```solidity
function _transfer(address sender, address recipient, uint256 amount) internal virtual
```

### _mint

```solidity
function _mint(address account, uint256 amount) internal virtual
```

### _burn

```solidity
function _burn(address account, uint256 amount) internal virtual
```

### _approve

```solidity
function _approve(address owner, address spender, uint256 amount) internal virtual
```

### deposit

```solidity
function deposit(address to, uint256 amount) external returns (bool)
```

_Deposits corresponding tokens from external chain, only callable by Fungible module._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| to | address |  |
| amount | uint256 | to deposit. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | bool | true/false if succeeded/failed. |

### withdrawGasFee

```solidity
function withdrawGasFee() public view returns (address, uint256)
```

_Withdraws gas fees._

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | address | returns the ZRC20 address for gas on the same chain of this ZRC20, and calculates the gas fee for withdraw() |
| [1] | uint256 |  |

### withdraw

```solidity
function withdraw(bytes to, uint256 amount) external returns (bool)
```

_Withraws ZRC20 tokens to external chains, this function causes cctx module to send out outbound tx to the outbound chain
this contract should be given enough allowance of the gas ZRC20 to pay for outbound tx gas fee._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| to | bytes |  |
| amount | uint256 | to deposit. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | bool | true/false if succeeded/failed. |

### updateSystemContractAddress

```solidity
function updateSystemContractAddress(address addr) external
```

_Updates system contract address. Can only be updated by the fungible module._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| addr | address |  |

### updateGasLimit

```solidity
function updateGasLimit(uint256 gasLimit) external
```

_Updates gas limit. Can only be updated by the fungible module._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| gasLimit | uint256 |  |

### updateProtocolFlatFee

```solidity
function updateProtocolFlatFee(uint256 protocolFlatFee) external
```

_Updates protocol flat fee. Can only be updated by the fungible module._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| protocolFlatFee | uint256 |  |

