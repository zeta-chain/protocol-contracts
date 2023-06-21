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

ZRC-20 is a token standard integrated into ZetaChain's omnichain smart contract
platform. With ZRC-20, developers can build dApps that orchestrate native assets
on any connected chain. This makes building Omnichain DeFi protocols and dApps
such as Omnichain DEXs, Omnichain Lending, Omnichain Portfolio Management, and
anything else that involves fungible tokens on multiple chains from a single
place extremely simple — as if they were all on a single chain.

At a high-level, ZRC-20 tokens are an extension of the standard ERC-20 tokens
found in the Ethereum ecosystem, ZRC-20 tokens have the added ability to manage
assets on all ZetaChain-connected chains. Any fungible token, including Bitcoin,
Dogecoin, ERC-20-equivalents on other chains, gas assets on other chains, and so
on, may be represented on ZetaChain as a ZRC-20 and orchestrated as if it were
any other fungible token (like an ERC-20). [wzeta](./wzeta.md)

### FUNGIBLE_MODULE_ADDRESS

```solidity
address FUNGIBLE_MODULE_ADDRESS
```

The fungible module address, this is maintained at the protocol level and is always constant

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

Only the fungible module is allowed to deploy a new ZRC20 contract.

_Constructor that gives msg.sender all of existing tokens._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| name_ | string | Name of the token |
| symbol_ | string | Symbol of the token |
| decimals_ | uint8 | Number of decimal places the token can be divided into |
| chainid_ | uint256 | Chain ID |
| coinType_ | enum CoinType | Coin Type |
| gasLimit_ | uint256 | Gas limit for transactions |
| systemContractAddress_ | address | Address of the system contract |

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

Transfers a specified amount of tokens to the given recipient.

_This function can be called by the contract owner or any other external address._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| recipient | address | The address of the recipient to whom the tokens will be transferred. |
| amount | uint256 | The amount of tokens to transfer. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | bool | Returns a boolean value indicating whether the transfer was successful or not. |

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

_Internal function to transfer tokens from one address to another.
Throws if either the sender or recipient address is zero.
Throws if the sender's balance is lower than the transfer amount._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| sender | address | The address sending the tokens. |
| recipient | address | The address receiving the tokens. |
| amount | uint256 | The amount of tokens to transfer. |

### _mint

```solidity
function _mint(address account, uint256 amount) internal virtual
```

_Internal function to mint new tokens and assign them to an account.
Throws if the account address is zero._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| account | address | The address to which the minted tokens will be assigned. |
| amount | uint256 | The amount of tokens to be minted. |

### _burn

```solidity
function _burn(address account, uint256 amount) internal virtual
```

_Internal function to burn tokens from an account.
Throws if the account address is zero.
Throws if the account's balance is lower than the burn amount._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| account | address | The address from which tokens will be burned. |
| amount | uint256 | The amount of tokens to be burned. |

### _approve

```solidity
function _approve(address owner, address spender, uint256 amount) internal virtual
```

_Internal function to approve a spender to spend tokens on behalf of the owner.
Throws if the owner address is zero.
Throws if the spender address is zero._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| owner | address | The address that owns the tokens. |
| spender | address | The address that is approved to spend the tokens. |
| amount | uint256 | The maximum amount of tokens that can be spent. |

### deposit

```solidity
function deposit(address to, uint256 amount) external returns (bool)
```

_Deposits corresponding tokens from an external chain.
Only callable by the Fungible module or the System contract.
Throws if called by an invalid sender._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| to | address | The recipient address. |
| amount | uint256 | The amount to deposit. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | bool | A boolean indicating whether the deposit succeeded or failed. |

### withdrawGasFee

```solidity
function withdrawGasFee() public view returns (address, uint256)
```

_Returns the ZRC20 address for gas on the same chain of this ZRC20, and calculates the gas fee for `withdraw()`.
Throws if the gas ZRC20 address is zero.
Throws if the gas price is zero._

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | address | gasZRC20 The address of the gas ZRC20 token on the same chain. |
| [1] | uint256 | gasFee The calculated gas fee for the `withdraw()` function. |

### withdraw

```solidity
function withdraw(bytes to, uint256 amount) external returns (bool)
```

_Withdraws ZRC20 tokens to external chains by triggering the crosschain module to create an outbound transaction.
Requires this contract to have sufficient allowance of the gas ZRC20 token to pay for the outbound transaction gas fee.
Throws if the gas fee transfer fails._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| to | bytes | The recipient address on the external chain. |
| amount | uint256 | The amount of tokens to withdraw. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | bool | A boolean indicating whether the withdrawal succeeded or failed. |

### updateSystemContractAddress

```solidity
function updateSystemContractAddress(address addr) external
```

_Updates the system contract address.
Requires the caller to be the fungible module._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| addr | address | The new system contract address to be set. |

### updateGasLimit

```solidity
function updateGasLimit(uint256 gasLimit) external
```

_Updates the gas limit.
Requires the caller to be the fungible module._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| gasLimit | uint256 | The new gas limit to be set. |

### updateProtocolFlatFee

```solidity
function updateProtocolFlatFee(uint256 protocolFlatFee) external
```

_Updates the protocol flat fee.
Requires the caller to be the fungible module._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| protocolFlatFee | uint256 | The new protocol flat fee to be set. |

