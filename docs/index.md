# Solidity API

## Ownable

### transferOwnership

```solidity
function transferOwnership(address newOwner) external
```

## ImmutableCreate2Factory

This contract provides a safeCreate2 function that takes a salt value
and a block of initialization code as arguments and passes them into inline
assembly. The contract prevents redeploys by maintaining a mapping of all
contracts that have already been deployed, and prevents frontrunning or other
collisions by requiring that the first 20 bytes of the salt are equal to the
address of the caller (this can be bypassed by setting the first 20 bytes to
the null address). There is also a view function that computes the address of
the contract that will be created when submitting a given salt or nonce along
with a given block of initialization code.

_This contract has not yet been fully tested or audited - proceed with
caution and please share any exploits or optimizations you discover._

### safeCreate2Internal

```solidity
function safeCreate2Internal(bytes32 salt, bytes initializationCode) internal returns (address deploymentAddress)
```

### safeCreate2

```solidity
function safeCreate2(bytes32 salt, bytes initializationCode) public payable returns (address deploymentAddress)
```

_Create a contract using CREATE2 by submitting a given salt or nonce
along with the initialization code for the contract. Note that the first 20
bytes of the salt must match those of the calling address, which prevents
contract creation events from being submitted by unintended parties._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| salt | bytes32 | bytes32 The nonce that will be passed into the CREATE2 call. |
| initializationCode | bytes | bytes The initialization code that will be passed into the CREATE2 call. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| deploymentAddress | address | Address of the contract that will be created, or the null address if a contract has already been deployed to that address. |

### findCreate2Address

```solidity
function findCreate2Address(bytes32 salt, bytes initCode) external view returns (address deploymentAddress)
```

_Compute the address of the contract that will be created when
submitting a given salt or nonce to the contract along with the contract's
initialization code. The CREATE2 address is computed in accordance with
EIP-1014, and adheres to the formula therein of
`keccak256( 0xff ++ address ++ salt ++ keccak256(init_code)))[12:]` when
performing the computation. The computed address is then checked for any
existing contract code - if so, the null address will be returned instead._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| salt | bytes32 | bytes32 The nonce passed into the CREATE2 address calculation. |
| initCode | bytes | bytes The contract initialization code to be used. that will be passed into the CREATE2 address calculation. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| deploymentAddress | address | Address of the contract that will be created, or the null address if a contract has already been deployed to that address. |

### findCreate2AddressViaHash

```solidity
function findCreate2AddressViaHash(bytes32 salt, bytes32 initCodeHash) external view returns (address deploymentAddress)
```

_Compute the address of the contract that will be created when
submitting a given salt or nonce to the contract along with the keccak256
hash of the contract's initialization code. The CREATE2 address is computed
in accordance with EIP-1014, and adheres to the formula therein of
`keccak256( 0xff ++ address ++ salt ++ keccak256(init_code)))[12:]` when
performing the computation. The computed address is then checked for any
existing contract code - if so, the null address will be returned instead._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| salt | bytes32 | bytes32 The nonce passed into the CREATE2 address calculation. |
| initCodeHash | bytes32 | bytes32 The keccak256 hash of the initialization code that will be passed into the CREATE2 address calculation. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| deploymentAddress | address | Address of the contract that will be created, or the null address if a contract has already been deployed to that address. |

### hasBeenDeployed

```solidity
function hasBeenDeployed(address deploymentAddress) external view returns (bool)
```

_Determine if a contract has already been deployed by the factory to a
given address._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| deploymentAddress | address | address The contract address to check. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | bool | True if the contract has been deployed, false otherwise. |

### containsCaller

```solidity
modifier containsCaller(bytes32 salt)
```

_Modifier to ensure that the first 20 bytes of a submitted salt match
those of the calling account. This provides protection against the salt
being stolen by frontrunners or other attackers. The protection can also be
bypassed if desired by setting each of the first 20 bytes to zero._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| salt | bytes32 | bytes32 The salt value to check against the calling address. |

### safeCreate2AndTransfer

```solidity
function safeCreate2AndTransfer(bytes32 salt, bytes initializationCode) external payable returns (address deploymentAddress)
```

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

## ZetaEth

_ZetaEth is an implementation of OpenZeppelin's ERC20_

### constructor

```solidity
constructor(uint256 initialSupply) public
```

## ZetaNonEth

### connectorAddress

```solidity
address connectorAddress
```

### tssAddress

```solidity
address tssAddress
```

_Collectively held by Zeta blockchain validators_

### tssAddressUpdater

```solidity
address tssAddressUpdater
```

_Initially a multi-sig, eventually held by Zeta blockchain validators (via renounceTssAddressUpdater)_

### Minted

```solidity
event Minted(address mintee, uint256 amount, bytes32 internalSendHash)
```

### Burnt

```solidity
event Burnt(address burnee, uint256 amount)
```

### constructor

```solidity
constructor(address tssAddress_, address tssAddressUpdater_) public
```

### updateTssAndConnectorAddresses

```solidity
function updateTssAndConnectorAddresses(address tssAddress_, address connectorAddress_) external
```

### renounceTssAddressUpdater

```solidity
function renounceTssAddressUpdater() external
```

_Sets tssAddressUpdater to be tssAddress_

### mint

```solidity
function mint(address mintee, uint256 value, bytes32 internalSendHash) external
```

### burnFrom

```solidity
function burnFrom(address account, uint256 amount) public
```

## ZetaConnectorBase

_Main abstraction of ZetaConnector.
This contract manages interactions between TSS and different chains.
There's an instance of this contract on each chain supported by ZetaChain._

### zetaToken

```solidity
address zetaToken
```

### pauserAddress

```solidity
address pauserAddress
```

_Multisig contract to pause incoming transactions.
The responsibility of pausing outgoing transactions is left to the protocol for more flexibility._

### tssAddress

```solidity
address tssAddress
```

_Collectively held by ZetaChain validators._

### tssAddressUpdater

```solidity
address tssAddressUpdater
```

_This address will start pointing to a multisig contract, then it will become the TSS address itself._

### ZetaSent

```solidity
event ZetaSent(address sourceTxOriginAddress, address zetaTxSenderAddress, uint256 destinationChainId, bytes destinationAddress, uint256 zetaValueAndGas, uint256 destinationGasLimit, bytes message, bytes zetaParams)
```

### ZetaReceived

```solidity
event ZetaReceived(bytes zetaTxSenderAddress, uint256 sourceChainId, address destinationAddress, uint256 zetaValue, bytes message, bytes32 internalSendHash)
```

### ZetaReverted

```solidity
event ZetaReverted(address zetaTxSenderAddress, uint256 sourceChainId, uint256 destinationChainId, bytes destinationAddress, uint256 remainingZetaValue, bytes message, bytes32 internalSendHash)
```

### TSSAddressUpdated

```solidity
event TSSAddressUpdated(address zetaTxSenderAddress, address newTssAddress)
```

### PauserAddressUpdated

```solidity
event PauserAddressUpdated(address updaterAddress, address newTssAddress)
```

### constructor

```solidity
constructor(address zetaToken_, address tssAddress_, address tssAddressUpdater_, address pauserAddress_) public
```

_Constructor requires initial addresses.
zetaToken address is the only immutable one, while others can be updated._

### onlyPauser

```solidity
modifier onlyPauser()
```

_Modifier to restrict actions to pauser address._

### onlyTssAddress

```solidity
modifier onlyTssAddress()
```

_Modifier to restrict actions to TSS address._

### onlyTssUpdater

```solidity
modifier onlyTssUpdater()
```

_Modifier to restrict actions to TSS updater address._

### updatePauserAddress

```solidity
function updatePauserAddress(address pauserAddress_) external
```

_Update the pauser address. The only address allowed to do that is the current pauser._

### updateTssAddress

```solidity
function updateTssAddress(address tssAddress_) external
```

_Update the TSS address. The address can be updated by the TSS updater or the TSS address itself._

### renounceTssAddressUpdater

```solidity
function renounceTssAddressUpdater() external
```

_Changes the ownership of tssAddressUpdater to be the one held by the ZetaChain TSS Signer nodes._

### pause

```solidity
function pause() external
```

_Pause the input (send) transactions._

### unpause

```solidity
function unpause() external
```

_Unpause the contract to allow transactions again._

### send

```solidity
function send(struct ZetaInterfaces.SendInput input) external virtual
```

_Entrypoint to send data and value through ZetaChain._

### onReceive

```solidity
function onReceive(bytes zetaTxSenderAddress, uint256 sourceChainId, address destinationAddress, uint256 zetaValue, bytes message, bytes32 internalSendHash) external virtual
```

_Handler to receive data from other chain.
This method can be called only by TSS. Access validation is in implementation._

### onRevert

```solidity
function onRevert(address zetaTxSenderAddress, uint256 sourceChainId, bytes destinationAddress, uint256 destinationChainId, uint256 remainingZetaValue, bytes message, bytes32 internalSendHash) external virtual
```

_Handler to receive errors from other chain.
This method can be called only by TSS. Access validation is in implementation._

## ZetaConnectorEth

_ETH implementation of ZetaConnector.
This contract manages interactions between TSS and different chains.
This version is only for Ethereum network because in the other chains we mint and burn and in this one we lock and unlock._

### constructor

```solidity
constructor(address zetaToken_, address tssAddress_, address tssAddressUpdater_, address pauserAddress_) public
```

### getLockedAmount

```solidity
function getLockedAmount() external view returns (uint256)
```

### send

```solidity
function send(struct ZetaInterfaces.SendInput input) external
```

_Entrypoint to send data through ZetaChain
This call locks the token on the contract and emits an event with all the data needed by the protocol._

### onReceive

```solidity
function onReceive(bytes zetaTxSenderAddress, uint256 sourceChainId, address destinationAddress, uint256 zetaValue, bytes message, bytes32 internalSendHash) external
```

_Handler to receive data from other chain.
This method can be called only by TSS.
Transfers the Zeta tokens to destination and calls onZetaMessage if it's needed._

### onRevert

```solidity
function onRevert(address zetaTxSenderAddress, uint256 sourceChainId, bytes destinationAddress, uint256 destinationChainId, uint256 remainingZetaValue, bytes message, bytes32 internalSendHash) external
```

_Handler to receive errors from other chain.
This method can be called only by TSS.
Transfers the Zeta tokens to destination and calls onZetaRevert if it's needed._

## ZetaConnectorNonEth

_Non ETH implementation of ZetaConnector.
This contract manages interactions between TSS and different chains.
This version is for every chain but Etherum network because in the other chains we mint and burn and in Etherum we lock and unlock_

### maxSupply

```solidity
uint256 maxSupply
```

### constructor

```solidity
constructor(address zetaTokenAddress_, address tssAddress_, address tssAddressUpdater_, address pauserAddress_) public
```

### getLockedAmount

```solidity
function getLockedAmount() external view returns (uint256)
```

### setMaxSupply

```solidity
function setMaxSupply(uint256 maxSupply_) external
```

### send

```solidity
function send(struct ZetaInterfaces.SendInput input) external
```

_Entry point to send data to protocol
This call burn the token and emit an event with all the data needed by the protocol_

### onReceive

```solidity
function onReceive(bytes zetaTxSenderAddress, uint256 sourceChainId, address destinationAddress, uint256 zetaValue, bytes message, bytes32 internalSendHash) external
```

_Handler to receive data from other chain.
This method can be called only by TSS.
Transfer the Zeta tokens to destination and calls onZetaMessage if it's needed.
To perform the transfer mint new tokens, validating first the maxSupply allowed in the current chain._

### onRevert

```solidity
function onRevert(address zetaTxSenderAddress, uint256 sourceChainId, bytes destinationAddress, uint256 destinationChainId, uint256 remainingZetaValue, bytes message, bytes32 internalSendHash) external
```

_Handler to receive errors from other chain.
This method can be called only by TSS.
Transfer the Zeta tokens to destination and calls onZetaRevert if it's needed.
To perform the transfer mint new tokens, validating first the maxSupply allowed in the current chain._

## ConnectorErrors

_Interface with connector custom errors_

### CallerIsNotPauser

```solidity
error CallerIsNotPauser(address caller)
```

### CallerIsNotTss

```solidity
error CallerIsNotTss(address caller)
```

### CallerIsNotTssUpdater

```solidity
error CallerIsNotTssUpdater(address caller)
```

### CallerIsNotTssOrUpdater

```solidity
error CallerIsNotTssOrUpdater(address caller)
```

### ZetaTransferError

```solidity
error ZetaTransferError()
```

### ExceedsMaxSupply

```solidity
error ExceedsMaxSupply(uint256 maxSupply)
```

## ZetaErrors

_Common custom errors_

### CallerIsNotTss

```solidity
error CallerIsNotTss(address caller)
```

### CallerIsNotConnector

```solidity
error CallerIsNotConnector(address caller)
```

### CallerIsNotTssUpdater

```solidity
error CallerIsNotTssUpdater(address caller)
```

### CallerIsNotTssOrUpdater

```solidity
error CallerIsNotTssOrUpdater(address caller)
```

### InvalidAddress

```solidity
error InvalidAddress()
```

### ZetaTransferError

```solidity
error ZetaTransferError()
```

## ZetaInteractorErrors

_Interface with Zeta Interactor errors_

### InvalidDestinationChainId

```solidity
error InvalidDestinationChainId()
```

### InvalidCaller

```solidity
error InvalidCaller(address caller)
```

### InvalidZetaMessageCall

```solidity
error InvalidZetaMessageCall()
```

### InvalidZetaRevertCall

```solidity
error InvalidZetaRevertCall()
```

## ZetaInterfaces

### SendInput

```solidity
struct SendInput {
  uint256 destinationChainId;
  bytes destinationAddress;
  uint256 destinationGasLimit;
  bytes message;
  uint256 zetaValueAndGas;
  bytes zetaParams;
}
```

### ZetaMessage

```solidity
struct ZetaMessage {
  bytes zetaTxSenderAddress;
  uint256 sourceChainId;
  address destinationAddress;
  uint256 zetaValue;
  bytes message;
}
```

### ZetaRevert

```solidity
struct ZetaRevert {
  address zetaTxSenderAddress;
  uint256 sourceChainId;
  bytes destinationAddress;
  uint256 destinationChainId;
  uint256 remainingZetaValue;
  bytes message;
}
```

## ZetaConnector

### send

```solidity
function send(struct ZetaInterfaces.SendInput input) external
```

_Sending value and data cross-chain is as easy as calling connector.send(SendInput)_

## ZetaReceiver

### onZetaMessage

```solidity
function onZetaMessage(struct ZetaInterfaces.ZetaMessage zetaMessage) external
```

_onZetaMessage is called when a cross-chain message reaches a contract_

### onZetaRevert

```solidity
function onZetaRevert(struct ZetaInterfaces.ZetaRevert zetaRevert) external
```

_onZetaRevert is called when a cross-chain message reverts.
It's useful to rollback to the original state_

## ZetaTokenConsumer

_ZetaTokenConsumer makes it easier to handle the following situations:
  - Getting Zeta using native coin (to pay for destination gas while using `connector.send`)
  - Getting Zeta using a token (to pay for destination gas while using `connector.send`)
  - Getting native coin using Zeta (to return unused destination gas when `onZetaRevert` is executed)
  - Getting a token using Zeta (to return unused destination gas when `onZetaRevert` is executed)
The interface can be implemented using different strategies, like UniswapV2, UniswapV3, etc_

### EthExchangedForZeta

```solidity
event EthExchangedForZeta(uint256 amountIn, uint256 amountOut)
```

### TokenExchangedForZeta

```solidity
event TokenExchangedForZeta(address token, uint256 amountIn, uint256 amountOut)
```

### ZetaExchangedForEth

```solidity
event ZetaExchangedForEth(uint256 amountIn, uint256 amountOut)
```

### ZetaExchangedForToken

```solidity
event ZetaExchangedForToken(address token, uint256 amountIn, uint256 amountOut)
```

### getZetaFromEth

```solidity
function getZetaFromEth(address destinationAddress, uint256 minAmountOut) external payable returns (uint256)
```

### getZetaFromToken

```solidity
function getZetaFromToken(address destinationAddress, uint256 minAmountOut, address inputToken, uint256 inputTokenAmount) external returns (uint256)
```

### getEthFromZeta

```solidity
function getEthFromZeta(address destinationAddress, uint256 minAmountOut, uint256 zetaTokenAmount) external returns (uint256)
```

### getTokenFromZeta

```solidity
function getTokenFromZeta(address destinationAddress, uint256 minAmountOut, address outputToken, uint256 zetaTokenAmount) external returns (uint256)
```

## ZetaCommonErrors

### InvalidAddress

```solidity
error InvalidAddress()
```

## ZetaNonEthInterface

_ZetaNonEthInterface is a mintable / burnable version of IERC20_

### burnFrom

```solidity
function burnFrom(address account, uint256 amount) external
```

### mint

```solidity
function mint(address mintee, uint256 value, bytes32 internalSendHash) external
```

## INonfungiblePositionManager

Wraps Uniswap V3 positions in a non-fungible token interface which allows for them to be transferred
and authorized.

### IncreaseLiquidity

```solidity
event IncreaseLiquidity(uint256 tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
```

Emitted when liquidity is increased for a position NFT

_Also emitted when a token is minted_

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| tokenId | uint256 | The ID of the token for which liquidity was increased |
| liquidity | uint128 | The amount by which liquidity for the NFT position was increased |
| amount0 | uint256 | The amount of token0 that was paid for the increase in liquidity |
| amount1 | uint256 | The amount of token1 that was paid for the increase in liquidity |

### DecreaseLiquidity

```solidity
event DecreaseLiquidity(uint256 tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
```

Emitted when liquidity is decreased for a position NFT

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| tokenId | uint256 | The ID of the token for which liquidity was decreased |
| liquidity | uint128 | The amount by which liquidity for the NFT position was decreased |
| amount0 | uint256 | The amount of token0 that was accounted for the decrease in liquidity |
| amount1 | uint256 | The amount of token1 that was accounted for the decrease in liquidity |

### Collect

```solidity
event Collect(uint256 tokenId, address recipient, uint256 amount0, uint256 amount1)
```

Emitted when tokens are collected for a position NFT

_The amounts reported may not be exactly equivalent to the amounts transferred, due to rounding behavior_

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| tokenId | uint256 | The ID of the token for which underlying tokens were collected |
| recipient | address | The address of the account that received the collected tokens |
| amount0 | uint256 | The amount of token0 owed to the position that was collected |
| amount1 | uint256 | The amount of token1 owed to the position that was collected |

### positions

```solidity
function positions(uint256 tokenId) external view returns (uint96 nonce, address operator, address token0, address token1, uint24 fee, int24 tickLower, int24 tickUpper, uint128 liquidity, uint256 feeGrowthInside0LastX128, uint256 feeGrowthInside1LastX128, uint128 tokensOwed0, uint128 tokensOwed1)
```

Returns the position information associated with a given token ID.

_Throws if the token ID is not valid._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| tokenId | uint256 | The ID of the token that represents the position |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| nonce | uint96 | The nonce for permits |
| operator | address | The address that is approved for spending |
| token0 | address | The address of the token0 for a specific pool |
| token1 | address | The address of the token1 for a specific pool |
| fee | uint24 | The fee associated with the pool |
| tickLower | int24 | The lower end of the tick range for the position |
| tickUpper | int24 | The higher end of the tick range for the position |
| liquidity | uint128 | The liquidity of the position |
| feeGrowthInside0LastX128 | uint256 | The fee growth of token0 as of the last action on the individual position |
| feeGrowthInside1LastX128 | uint256 | The fee growth of token1 as of the last action on the individual position |
| tokensOwed0 | uint128 | The uncollected amount of token0 owed to the position as of the last computation |
| tokensOwed1 | uint128 | The uncollected amount of token1 owed to the position as of the last computation |

### MintParams

```solidity
struct MintParams {
  address token0;
  address token1;
  uint24 fee;
  int24 tickLower;
  int24 tickUpper;
  uint256 amount0Desired;
  uint256 amount1Desired;
  uint256 amount0Min;
  uint256 amount1Min;
  address recipient;
  uint256 deadline;
}
```

### mint

```solidity
function mint(struct INonfungiblePositionManager.MintParams params) external payable returns (uint256 tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
```

Creates a new position wrapped in a NFT

_Call this when the pool does exist and is initialized. Note that if the pool is created but not initialized
a method does not exist, i.e. the pool is assumed to be initialized._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| params | struct INonfungiblePositionManager.MintParams | The params necessary to mint a position, encoded as `MintParams` in calldata |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| tokenId | uint256 | The ID of the token that represents the minted position |
| liquidity | uint128 | The amount of liquidity for this position |
| amount0 | uint256 | The amount of token0 |
| amount1 | uint256 | The amount of token1 |

### IncreaseLiquidityParams

```solidity
struct IncreaseLiquidityParams {
  uint256 tokenId;
  uint256 amount0Desired;
  uint256 amount1Desired;
  uint256 amount0Min;
  uint256 amount1Min;
  uint256 deadline;
}
```

### increaseLiquidity

```solidity
function increaseLiquidity(struct INonfungiblePositionManager.IncreaseLiquidityParams params) external payable returns (uint128 liquidity, uint256 amount0, uint256 amount1)
```

Increases the amount of liquidity in a position, with tokens paid by the `msg.sender`

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| params | struct INonfungiblePositionManager.IncreaseLiquidityParams | tokenId The ID of the token for which liquidity is being increased, amount0Desired The desired amount of token0 to be spent, amount1Desired The desired amount of token1 to be spent, amount0Min The minimum amount of token0 to spend, which serves as a slippage check, amount1Min The minimum amount of token1 to spend, which serves as a slippage check, deadline The time by which the transaction must be included to effect the change |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| liquidity | uint128 | The new liquidity amount as a result of the increase |
| amount0 | uint256 | The amount of token0 to acheive resulting liquidity |
| amount1 | uint256 | The amount of token1 to acheive resulting liquidity |

### DecreaseLiquidityParams

```solidity
struct DecreaseLiquidityParams {
  uint256 tokenId;
  uint128 liquidity;
  uint256 amount0Min;
  uint256 amount1Min;
  uint256 deadline;
}
```

### decreaseLiquidity

```solidity
function decreaseLiquidity(struct INonfungiblePositionManager.DecreaseLiquidityParams params) external payable returns (uint256 amount0, uint256 amount1)
```

Decreases the amount of liquidity in a position and accounts it to the position

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| params | struct INonfungiblePositionManager.DecreaseLiquidityParams | tokenId The ID of the token for which liquidity is being decreased, amount The amount by which liquidity will be decreased, amount0Min The minimum amount of token0 that should be accounted for the burned liquidity, amount1Min The minimum amount of token1 that should be accounted for the burned liquidity, deadline The time by which the transaction must be included to effect the change |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| amount0 | uint256 | The amount of token0 accounted to the position's tokens owed |
| amount1 | uint256 | The amount of token1 accounted to the position's tokens owed |

### CollectParams

```solidity
struct CollectParams {
  uint256 tokenId;
  address recipient;
  uint128 amount0Max;
  uint128 amount1Max;
}
```

### collect

```solidity
function collect(struct INonfungiblePositionManager.CollectParams params) external payable returns (uint256 amount0, uint256 amount1)
```

Collects up to a maximum amount of fees owed to a specific position to the recipient

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| params | struct INonfungiblePositionManager.CollectParams | tokenId The ID of the NFT for which tokens are being collected, recipient The account that should receive the tokens, amount0Max The maximum amount of token0 to collect, amount1Max The maximum amount of token1 to collect |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| amount0 | uint256 | The amount of fees collected in token0 |
| amount1 | uint256 | The amount of fees collected in token1 |

### burn

```solidity
function burn(uint256 tokenId) external payable
```

Burns a token ID, which deletes it from the NFT contract. The token must have 0 liquidity and all tokens
must be collected first.

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| tokenId | uint256 | The ID of the token that is being burned |

## IPoolInitializer

### createAndInitializePoolIfNecessary

```solidity
function createAndInitializePoolIfNecessary(address token0, address token1, uint24 fee, uint160 sqrtPriceX96) external payable returns (address pool)
```

Creates a new pool if it does not exist, then initializes if not initialized

_This method can be bundled with others via IMulticall for the first action (e.g. mint) performed against a pool_

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| token0 | address | The contract address of token0 of the pool |
| token1 | address | The contract address of token1 of the pool |
| fee | uint24 | The fee amount of the v3 pool for the specified token pair |
| sqrtPriceX96 | uint160 | The initial square root price of the pool as a Q64.96 value |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| pool | address | Returns the pool address based on the pair of tokens and fee, will return the newly created pool address if necessary |

## ZetaInteractorMock

### constructor

```solidity
constructor(address zetaConnectorAddress) public
```

### onZetaMessage

```solidity
function onZetaMessage(struct ZetaInterfaces.ZetaMessage zetaMessage) external
```

_onZetaMessage is called when a cross-chain message reaches a contract_

### onZetaRevert

```solidity
function onZetaRevert(struct ZetaInterfaces.ZetaRevert zetaRevert) external
```

_onZetaRevert is called when a cross-chain message reverts.
It's useful to rollback to the original state_

## ZetaReceiverMock

### MockOnZetaMessage

```solidity
event MockOnZetaMessage(address destinationAddress)
```

### MockOnZetaRevert

```solidity
event MockOnZetaRevert(address zetaTxSenderAddress)
```

### onZetaMessage

```solidity
function onZetaMessage(struct ZetaInterfaces.ZetaMessage zetaMessage) external
```

_onZetaMessage is called when a cross-chain message reaches a contract_

### onZetaRevert

```solidity
function onZetaRevert(struct ZetaInterfaces.ZetaRevert zetaRevert) external
```

_onZetaRevert is called when a cross-chain message reverts.
It's useful to rollback to the original state_

## ZetaInteractor

### ZERO_BYTES

```solidity
bytes32 ZERO_BYTES
```

### currentChainId

```solidity
uint256 currentChainId
```

### connector

```solidity
contract ZetaConnector connector
```

### interactorsByChainId

```solidity
mapping(uint256 => bytes) interactorsByChainId
```

_Maps a chain id to its corresponding address of the MultiChainSwap contract
The address is expressed in bytes to allow non-EVM chains
This mapping is useful, mainly, for two reasons:
 - Given a chain id, the contract is able to route a transaction to its corresponding address
 - To check that the messages (onZetaMessage, onZetaRevert) come from a trusted source_

### isValidMessageCall

```solidity
modifier isValidMessageCall(struct ZetaInterfaces.ZetaMessage zetaMessage)
```

### isValidRevertCall

```solidity
modifier isValidRevertCall(struct ZetaInterfaces.ZetaRevert zetaRevert)
```

### constructor

```solidity
constructor(address zetaConnectorAddress) internal
```

### _isValidChainId

```solidity
function _isValidChainId(uint256 chainId) internal view returns (bool)
```

_Useful for contracts that inherit from this one_

### setInteractorByChainId

```solidity
function setInteractorByChainId(uint256 destinationChainId, bytes contractAddress) external
```

## ZetaTokenConsumerTridentErrors

### InputCantBeZero

```solidity
error InputCantBeZero()
```

### ErrorSendingETH

```solidity
error ErrorSendingETH()
```

### ReentrancyError

```solidity
error ReentrancyError()
```

## WETH9

### deposit

```solidity
function deposit() external payable
```

### withdraw

```solidity
function withdraw(uint256 wad) external
```

### depositTo

```solidity
function depositTo(address to) external payable
```

### withdrawTo

```solidity
function withdrawTo(address payable to, uint256 value) external
```

## ZetaTokenConsumerTrident

_Trident strategy for ZetaTokenConsumer_

### MAX_DEADLINE

```solidity
uint256 MAX_DEADLINE
```

### WETH9Address

```solidity
address WETH9Address
```

### zetaToken

```solidity
address zetaToken
```

### tridentRouter

```solidity
contract IPoolRouter tridentRouter
```

### poolFactory

```solidity
contract ConcentratedLiquidityPoolFactory poolFactory
```

### _locked

```solidity
bool _locked
```

### constructor

```solidity
constructor(address zetaToken_, address uniswapV3Router_, address WETH9Address_, address poolFactory_) public
```

### nonReentrant

```solidity
modifier nonReentrant()
```

### receive

```solidity
receive() external payable
```

### getPair

```solidity
function getPair(address token0, address token1) internal pure returns (address, address)
```

### getZetaFromEth

```solidity
function getZetaFromEth(address destinationAddress, uint256 minAmountOut) external payable returns (uint256)
```

### getZetaFromToken

```solidity
function getZetaFromToken(address destinationAddress, uint256 minAmountOut, address inputToken, uint256 inputTokenAmount) external returns (uint256)
```

### getEthFromZeta

```solidity
function getEthFromZeta(address destinationAddress, uint256 minAmountOut, uint256 zetaTokenAmount) external returns (uint256)
```

### getTokenFromZeta

```solidity
function getTokenFromZeta(address destinationAddress, uint256 minAmountOut, address outputToken, uint256 zetaTokenAmount) external returns (uint256)
```

## ZetaTokenConsumerUniV2Errors

### InputCantBeZero

```solidity
error InputCantBeZero()
```

## ZetaTokenConsumerUniV2

_Uniswap V2 strategy for ZetaTokenConsumer_

### MAX_DEADLINE

```solidity
uint256 MAX_DEADLINE
```

### wETH

```solidity
address wETH
```

### zetaToken

```solidity
address zetaToken
```

### uniswapV2Router

```solidity
contract IUniswapV2Router02 uniswapV2Router
```

### constructor

```solidity
constructor(address zetaToken_, address uniswapV2Router_) public
```

### getZetaFromEth

```solidity
function getZetaFromEth(address destinationAddress, uint256 minAmountOut) external payable returns (uint256)
```

### getZetaFromToken

```solidity
function getZetaFromToken(address destinationAddress, uint256 minAmountOut, address inputToken, uint256 inputTokenAmount) external returns (uint256)
```

### getEthFromZeta

```solidity
function getEthFromZeta(address destinationAddress, uint256 minAmountOut, uint256 zetaTokenAmount) external returns (uint256)
```

### getTokenFromZeta

```solidity
function getTokenFromZeta(address destinationAddress, uint256 minAmountOut, address outputToken, uint256 zetaTokenAmount) external returns (uint256)
```

## ZetaTokenConsumerUniV3Errors

### InputCantBeZero

```solidity
error InputCantBeZero()
```

### ErrorSendingETH

```solidity
error ErrorSendingETH()
```

### ReentrancyError

```solidity
error ReentrancyError()
```

## WETH9

### withdraw

```solidity
function withdraw(uint256 wad) external
```

## ZetaTokenConsumerUniV3

_Uniswap V3 strategy for ZetaTokenConsumer_

### MAX_DEADLINE

```solidity
uint256 MAX_DEADLINE
```

### zetaPoolFee

```solidity
uint24 zetaPoolFee
```

### tokenPoolFee

```solidity
uint24 tokenPoolFee
```

### WETH9Address

```solidity
address WETH9Address
```

### zetaToken

```solidity
address zetaToken
```

### uniswapV3Router

```solidity
contract ISwapRouter uniswapV3Router
```

### quoter

```solidity
contract IQuoter quoter
```

### _locked

```solidity
bool _locked
```

### constructor

```solidity
constructor(address zetaToken_, address uniswapV3Router_, address quoter_, address WETH9Address_, uint24 zetaPoolFee_, uint24 tokenPoolFee_) public
```

### nonReentrant

```solidity
modifier nonReentrant()
```

### receive

```solidity
receive() external payable
```

### getZetaFromEth

```solidity
function getZetaFromEth(address destinationAddress, uint256 minAmountOut) external payable returns (uint256)
```

### getZetaFromToken

```solidity
function getZetaFromToken(address destinationAddress, uint256 minAmountOut, address inputToken, uint256 inputTokenAmount) external returns (uint256)
```

### getEthFromZeta

```solidity
function getEthFromZeta(address destinationAddress, uint256 minAmountOut, uint256 zetaTokenAmount) external returns (uint256)
```

### getTokenFromZeta

```solidity
function getTokenFromZeta(address destinationAddress, uint256 minAmountOut, address outputToken, uint256 zetaTokenAmount) external returns (uint256)
```

## ConcentratedLiquidityPoolFactory

### getPools

```solidity
function getPools(address token0, address token1, uint256 startIndex, uint256 count) external view returns (address[] pairPools)
```

## IPoolRouter

### ExactInputSingleParams

```solidity
struct ExactInputSingleParams {
  address tokenIn;
  uint256 amountIn;
  uint256 amountOutMinimum;
  address pool;
  address to;
  bool unwrap;
}
```

### ExactInputParams

```solidity
struct ExactInputParams {
  address tokenIn;
  uint256 amountIn;
  uint256 amountOutMinimum;
  address[] path;
  address to;
  bool unwrap;
}
```

### ExactOutputSingleParams

```solidity
struct ExactOutputSingleParams {
  address tokenIn;
  uint256 amountOut;
  uint256 amountInMaximum;
  address pool;
  address to;
  bool unwrap;
}
```

### ExactOutputParams

```solidity
struct ExactOutputParams {
  address tokenIn;
  uint256 amountOut;
  uint256 amountInMaximum;
  address[] path;
  address to;
  bool unwrap;
}
```

### exactInputSingle

```solidity
function exactInputSingle(struct IPoolRouter.ExactInputSingleParams params) external payable returns (uint256 amountOut)
```

Swap amountIn of one token for as much as possible of another token

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| params | struct IPoolRouter.ExactInputSingleParams | The parameters necessary for the swap, encoded as ExactInputSingleParams in calldata |

### exactInput

```solidity
function exactInput(struct IPoolRouter.ExactInputParams params) external payable returns (uint256 amountOut)
```

Swap amountIn of one token for as much as possible of another along the specified path

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| params | struct IPoolRouter.ExactInputParams | The parameters necessary for the multi-hop swap, encoded as ExactInputParams in calldata |

### exactOutputSingle

```solidity
function exactOutputSingle(struct IPoolRouter.ExactOutputSingleParams params) external payable returns (uint256 amountIn)
```

Swaps as little as possible of one token for `amountOut` of another token

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| params | struct IPoolRouter.ExactOutputSingleParams | The parameters necessary for the swap, encoded as ExactOutputSingleParams in calldata |

### exactOutput

```solidity
function exactOutput(struct IPoolRouter.ExactOutputParams params) external payable returns (uint256 amountIn)
```

Swaps as little as possible of one token for `amountOut` of another along the specified path (reversed)

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| params | struct IPoolRouter.ExactOutputParams | The parameters necessary for the multi-hop swap, encoded as ExactOutputParams in calldata |

### sweep

```solidity
function sweep(address token, uint256 amount, address recipient) external payable
```

Recover mistakenly sent tokens

## ZetaInterfaces

### SendInput

```solidity
struct SendInput {
  uint256 destinationChainId;
  bytes destinationAddress;
  uint256 destinationGasLimit;
  bytes message;
  uint256 zetaValueAndGas;
  bytes zetaParams;
}
```

### ZetaMessage

```solidity
struct ZetaMessage {
  bytes zetaTxSenderAddress;
  uint256 sourceChainId;
  address destinationAddress;
  uint256 zetaValue;
  bytes message;
}
```

### ZetaRevert

```solidity
struct ZetaRevert {
  address zetaTxSenderAddress;
  uint256 sourceChainId;
  bytes destinationAddress;
  uint256 destinationChainId;
  uint256 remainingZetaValue;
  bytes message;
}
```

## WZETA

### transferFrom

```solidity
function transferFrom(address src, address dst, uint256 wad) external returns (bool)
```

### withdraw

```solidity
function withdraw(uint256 wad) external
```

## ZetaConnectorZEVM

### OnlyWZETA

```solidity
error OnlyWZETA()
```

Contract custom errors.

### WZETATransferFailed

```solidity
error WZETATransferFailed()
```

### OnlyFungibleModule

```solidity
error OnlyFungibleModule()
```

### FailedZetaSent

```solidity
error FailedZetaSent()
```

### FUNGIBLE_MODULE_ADDRESS

```solidity
address FUNGIBLE_MODULE_ADDRESS
```

Fungible module address.

### wzeta

```solidity
address wzeta
```

WZETA token address.

### ZetaSent

```solidity
event ZetaSent(address sourceTxOriginAddress, address zetaTxSenderAddress, uint256 destinationChainId, bytes destinationAddress, uint256 zetaValueAndGas, uint256 destinationGasLimit, bytes message, bytes zetaParams)
```

### SetWZETA

```solidity
event SetWZETA(address wzeta_)
```

### constructor

```solidity
constructor(address _wzeta) public
```

### receive

```solidity
receive() external payable
```

_Receive function to receive ZETA from WETH9.withdraw()._

### send

```solidity
function send(struct ZetaInterfaces.SendInput input) external
```

_Sends ZETA and bytes messages (to execute it) crosschain._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| input | struct ZetaInterfaces.SendInput |  |

### setWzetaAddress

```solidity
function setWzetaAddress(address wzeta_) external
```

_Sends ZETA and bytes messages (to execute it) crosschain._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| wzeta_ | address |  |

## ISystem

_Interfaces of SystemContract and ZRC20 to make easier to import._

### FUNGIBLE_MODULE_ADDRESS

```solidity
function FUNGIBLE_MODULE_ADDRESS() external view returns (address)
```

### wZetaContractAddress

```solidity
function wZetaContractAddress() external view returns (address)
```

### uniswapv2FactoryAddress

```solidity
function uniswapv2FactoryAddress() external view returns (address)
```

### gasPriceByChainId

```solidity
function gasPriceByChainId(uint256 chainID) external view returns (uint256)
```

### gasCoinZRC20ByChainId

```solidity
function gasCoinZRC20ByChainId(uint256 chainID) external view returns (address)
```

### gasZetaPoolByChainId

```solidity
function gasZetaPoolByChainId(uint256 chainID) external view returns (address)
```

## IZRC20

### totalSupply

```solidity
function totalSupply() external view returns (uint256)
```

### balanceOf

```solidity
function balanceOf(address account) external view returns (uint256)
```

### transfer

```solidity
function transfer(address recipient, uint256 amount) external returns (bool)
```

### allowance

```solidity
function allowance(address owner, address spender) external view returns (uint256)
```

### approve

```solidity
function approve(address spender, uint256 amount) external returns (bool)
```

### transferFrom

```solidity
function transferFrom(address sender, address recipient, uint256 amount) external returns (bool)
```

### deposit

```solidity
function deposit(address to, uint256 amount) external returns (bool)
```

### withdraw

```solidity
function withdraw(bytes to, uint256 amount) external returns (bool)
```

### withdrawGasFee

```solidity
function withdrawGasFee() external view returns (address, uint256)
```

### Transfer

```solidity
event Transfer(address from, address to, uint256 value)
```

### Approval

```solidity
event Approval(address owner, address spender, uint256 value)
```

### Deposit

```solidity
event Deposit(bytes from, address to, uint256 value)
```

### Withdrawal

```solidity
event Withdrawal(address from, bytes to, uint256 value, uint256 gasfee, uint256 protocolFlatFee)
```

### UpdatedSystemContract

```solidity
event UpdatedSystemContract(address systemContract)
```

### UpdatedGasLimit

```solidity
event UpdatedGasLimit(uint256 gasLimit)
```

### UpdatedProtocolFlatFee

```solidity
event UpdatedProtocolFlatFee(uint256 protocolFlatFee)
```

## Context

### _msgSender

```solidity
function _msgSender() internal view virtual returns (address)
```

### _msgData

```solidity
function _msgData() internal view virtual returns (bytes)
```

## IZRC20Metadata

### name

```solidity
function name() external view returns (string)
```

### symbol

```solidity
function symbol() external view returns (string)
```

### decimals

```solidity
function decimals() external view returns (uint8)
```

## CoinType

```solidity
enum CoinType {
  Zeta,
  Gas,
  ERC20
}
```

## zContract

_Any ZetaChain Contract must implement this interface to allow SystemContract to interact with.
This is only required if the contract wants to interact with other chains._

### onCrossChainCall

```solidity
function onCrossChainCall(address zrc20, uint256 amount, bytes message) external
```

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

## IUniswapV2Router01

### factory

```solidity
function factory() external pure returns (address)
```

### WETH

```solidity
function WETH() external pure returns (address)
```

### addLiquidity

```solidity
function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) external returns (uint256 amountA, uint256 amountB, uint256 liquidity)
```

### addLiquidityETH

```solidity
function addLiquidityETH(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) external payable returns (uint256 amountToken, uint256 amountETH, uint256 liquidity)
```

### removeLiquidity

```solidity
function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) external returns (uint256 amountA, uint256 amountB)
```

### removeLiquidityETH

```solidity
function removeLiquidityETH(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) external returns (uint256 amountToken, uint256 amountETH)
```

### removeLiquidityWithPermit

```solidity
function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) external returns (uint256 amountA, uint256 amountB)
```

### removeLiquidityETHWithPermit

```solidity
function removeLiquidityETHWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) external returns (uint256 amountToken, uint256 amountETH)
```

### swapExactTokensForTokens

```solidity
function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) external returns (uint256[] amounts)
```

### swapTokensForExactTokens

```solidity
function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) external returns (uint256[] amounts)
```

### swapExactETHForTokens

```solidity
function swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) external payable returns (uint256[] amounts)
```

### swapTokensForExactETH

```solidity
function swapTokensForExactETH(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) external returns (uint256[] amounts)
```

### swapExactTokensForETH

```solidity
function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) external returns (uint256[] amounts)
```

### swapETHForExactTokens

```solidity
function swapETHForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) external payable returns (uint256[] amounts)
```

### quote

```solidity
function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) external pure returns (uint256 amountB)
```

### getAmountOut

```solidity
function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) external pure returns (uint256 amountOut)
```

### getAmountIn

```solidity
function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) external pure returns (uint256 amountIn)
```

### getAmountsOut

```solidity
function getAmountsOut(uint256 amountIn, address[] path) external view returns (uint256[] amounts)
```

### getAmountsIn

```solidity
function getAmountsIn(uint256 amountOut, address[] path) external view returns (uint256[] amounts)
```

## IUniswapV2Router02

### removeLiquidityETHSupportingFeeOnTransferTokens

```solidity
function removeLiquidityETHSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) external returns (uint256 amountETH)
```

### removeLiquidityETHWithPermitSupportingFeeOnTransferTokens

```solidity
function removeLiquidityETHWithPermitSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) external returns (uint256 amountETH)
```

### swapExactTokensForTokensSupportingFeeOnTransferTokens

```solidity
function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) external
```

### swapExactETHForTokensSupportingFeeOnTransferTokens

```solidity
function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) external payable
```

### swapExactTokensForETHSupportingFeeOnTransferTokens

```solidity
function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) external
```

## IZRC20

### totalSupply

```solidity
function totalSupply() external view returns (uint256)
```

### balanceOf

```solidity
function balanceOf(address account) external view returns (uint256)
```

### transfer

```solidity
function transfer(address recipient, uint256 amount) external returns (bool)
```

### allowance

```solidity
function allowance(address owner, address spender) external view returns (uint256)
```

### approve

```solidity
function approve(address spender, uint256 amount) external returns (bool)
```

### decreaseAllowance

```solidity
function decreaseAllowance(address spender, uint256 amount) external returns (bool)
```

### increaseAllowance

```solidity
function increaseAllowance(address spender, uint256 amount) external returns (bool)
```

### transferFrom

```solidity
function transferFrom(address sender, address recipient, uint256 amount) external returns (bool)
```

### deposit

```solidity
function deposit(address to, uint256 amount) external returns (bool)
```

### burn

```solidity
function burn(address account, uint256 amount) external returns (bool)
```

### withdraw

```solidity
function withdraw(bytes to, uint256 amount) external returns (bool)
```

### withdrawGasFee

```solidity
function withdrawGasFee() external view returns (address, uint256)
```

### PROTOCOL_FEE

```solidity
function PROTOCOL_FEE() external view returns (uint256)
```

### Transfer

```solidity
event Transfer(address from, address to, uint256 value)
```

### Approval

```solidity
event Approval(address owner, address spender, uint256 value)
```

### Deposit

```solidity
event Deposit(bytes from, address to, uint256 value)
```

### Withdrawal

```solidity
event Withdrawal(address from, bytes to, uint256 value, uint256 gasFee, uint256 protocolFlatFee)
```

### UpdatedSystemContract

```solidity
event UpdatedSystemContract(address systemContract)
```

### UpdatedGasLimit

```solidity
event UpdatedGasLimit(uint256 gasLimit)
```

### UpdatedProtocolFlatFee

```solidity
event UpdatedProtocolFlatFee(uint256 protocolFlatFee)
```

## zContract

### onCrossChainCall

```solidity
function onCrossChainCall(address zrc20, uint256 amount, bytes message) external
```

## UniswapImports

## UniswapImports

## WETH9

### name

```solidity
string storage pointer name
```

### symbol

```solidity
string storage pointer symbol
```

### decimals

```solidity
uint8 decimals
```

### Approval

```solidity
event Approval(address src, address guy, uint256 wad)
```

### Transfer

```solidity
event Transfer(address src, address dst, uint256 wad)
```

### Deposit

```solidity
event Deposit(address dst, uint256 wad)
```

### Withdrawal

```solidity
event Withdrawal(address src, uint256 wad)
```

### balanceOf

```solidity
mapping(address => uint256) balanceOf
```

### allowance

```solidity
mapping(address => mapping(address => uint256)) allowance
```

### 

```solidity
undefined() public payable
```

### deposit

```solidity
undefined() public payable
```

### withdraw

```solidity
undefined(uint256 wad) public
```

### totalSupply

```solidity
undefined() public view returns (uint256)
```

### approve

```solidity
undefined(address guy, uint256 wad) public returns (bool)
```

### transfer

```solidity
undefined(address dst, uint256 wad) public returns (bool)
```

### transferFrom

```solidity
undefined(address src, address dst, uint256 wad) public returns (bool)
```

