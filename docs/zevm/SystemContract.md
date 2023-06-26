# zevm/SystemContract.md

## SystemContractErrors

_Custom errors for SystemContract_

### CallerIsNotFungibleModule

```solidity
error CallerIsNotFungibleModule()
```

### InvalidTarget

```solidity
error InvalidTarget()
```

### CantBeIdenticalAddresses

```solidity
error CantBeIdenticalAddresses()
```

### CantBeZeroAddress

```solidity
error CantBeZeroAddress()
```

### ZeroAddress

```solidity
error ZeroAddress()
```

## SystemContract

_The system contract it's called by the protocol to interact with the blockchain.
Also includes a lot of tools to make easier to interact with ZetaChain._

### gasPriceByChainId

```solidity
mapping(uint256 => uint256) gasPriceByChainId
```

Map to know the gas price of each chain given a chain id.

### gasCoinZRC20ByChainId

```solidity
mapping(uint256 => address) gasCoinZRC20ByChainId
```

Map to know the ZRC20 address of a token given a chain id, ex zETH, zBNB etc.

### gasZetaPoolByChainId

```solidity
mapping(uint256 => address) gasZetaPoolByChainId
```

### FUNGIBLE_MODULE_ADDRESS

```solidity
address FUNGIBLE_MODULE_ADDRESS
```

Fungible address is always the same, it's on protocol level.

### uniswapv2FactoryAddress

```solidity
address uniswapv2FactoryAddress
```

Uniswap V2 addresses.

### uniswapv2Router02Address

```solidity
address uniswapv2Router02Address
```

### wZetaContractAddress

```solidity
address wZetaContractAddress
```

Address of the wrapped ZETA to interact with Uniswap V2.

### zetaConnectorZEVMAddress

```solidity
address zetaConnectorZEVMAddress
```

Address of ZEVM Zeta Connector.

### SystemContractDeployed

```solidity
event SystemContractDeployed()
```

Custom SystemContract errors.

### SetGasPrice

```solidity
event SetGasPrice(uint256, uint256)
```

### SetGasCoin

```solidity
event SetGasCoin(uint256, address)
```

### SetGasZetaPool

```solidity
event SetGasZetaPool(uint256, address)
```

### SetWZeta

```solidity
event SetWZeta(address)
```

### SetConnectorZEVM

```solidity
event SetConnectorZEVM(address)
```

### constructor

```solidity
constructor(address wzeta_, address uniswapv2Factory_, address uniswapv2Router02_) public
```

_Only fungible module can deploy a system contract._

### depositAndCall

```solidity
function depositAndCall(address zrc20, uint256 amount, address target, bytes message) external
```

_Deposit foreign coins into ZRC20 and call user specified contract on zEVM._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| zrc20 | address | address for deposit. |
| amount | uint256 | to deposit. |
| target | address |  |
| message | bytes |  |

### sortTokens

```solidity
function sortTokens(address tokenA, address tokenB) internal pure returns (address token0, address token1)
```

_Sort token addresses lexicographically. Used to handle return values from pairs sorted in the order._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| tokenA | address | address. |
| tokenB | address | address. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| token0 | address | token1, returns sorted token addresses,. |
| token1 | address |  |

### uniswapv2PairFor

```solidity
function uniswapv2PairFor(address factory, address tokenA, address tokenB) public pure returns (address pair)
```

_Calculates the CREATE2 address for a pair without making any external calls._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| factory | address | address. |
| tokenA | address | address. |
| tokenB | address | address. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| pair | address | tokens pair address. |

### setGasPrice

```solidity
function setGasPrice(uint256 chainID, uint256 price) external
```

_Fungible module updates the gas price oracle periodically._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| chainID | uint256 |  |
| price | uint256 |  |

### setGasCoinZRC20

```solidity
function setGasCoinZRC20(uint256 chainID, address zrc20) external
```

_Setter for gasCoinZRC20ByChainId map._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| chainID | uint256 |  |
| zrc20 | address |  |

### setGasZetaPool

```solidity
function setGasZetaPool(uint256 chainID, address erc20) external
```

_Set the pool wzeta/erc20 address._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| chainID | uint256 |  |
| erc20 | address |  |

### setWZETAContractAddress

```solidity
function setWZETAContractAddress(address addr) external
```

_Setter for wrapped ZETA address._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| addr | address |  |

### setConnectorZEVMAddress

```solidity
function setConnectorZEVMAddress(address addr) external
```

_Setter for zetaConnector ZEVM Address_

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| addr | address |  |

