# GatewayZEVMValidations
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/zevm/libraries/GatewayZEVMValidations.sol)

Library containing validation functions for GatewayZEVM contract

*This library provides common validation logic used across GatewayZEVM contract*


## State Variables
### MAX_MESSAGE_SIZE
Maximum message size constant


```solidity
uint256 internal constant MAX_MESSAGE_SIZE = 2048;
```


### MIN_GAS_LIMIT
Minimum gas limit constant


```solidity
uint256 internal constant MIN_GAS_LIMIT = 100_000;
```


## Functions
### validateNonZeroAddress

*Validates that an address is not zero*


```solidity
function validateNonZeroAddress(address addr) internal pure;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`addr`|`address`|The address to validate|


### validateReceiver

*Validates that receiver bytes are not empty*


```solidity
function validateReceiver(bytes memory receiver) internal pure;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`receiver`|`bytes`|The receiver bytes to validate|


### validateAmount

*Validates that amount is not zero*


```solidity
function validateAmount(uint256 amount) internal pure;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`amount`|`uint256`|The amount to validate|


### validateGasLimit

*Validates that gas limit meets minimum requirement*


```solidity
function validateGasLimit(uint256 gasLimit) internal pure;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`gasLimit`|`uint256`|The gas limit to validate|


### validateTarget

*Validates that target address is not restricted*


```solidity
function validateTarget(address target, address protocolAddress, address contractAddress) private pure;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`target`|`address`|The target address to validate|
|`protocolAddress`|`address`|The protocol address to check against|
|`contractAddress`|`address`|The contract address to check against|


### validateMessageSize

*Validates message size constraints*


```solidity
function validateMessageSize(uint256 messageLength, uint256 revertMessageLength) internal pure;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`messageLength`|`uint256`|The length of the main message|
|`revertMessageLength`|`uint256`|The length of the revert message|


### validateRevertOptions

*Validates revert options*


```solidity
function validateRevertOptions(RevertOptions calldata revertOptions) internal pure;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`revertOptions`|`RevertOptions`|The revert options to validate|


### validateCallAndRevertOptions

*Validates call options and revert options together*


```solidity
function validateCallAndRevertOptions(
    CallOptions calldata callOptions,
    RevertOptions calldata revertOptions,
    uint256 messageLength
)
    internal
    pure;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`callOptions`|`CallOptions`|The call options to validate|
|`revertOptions`|`RevertOptions`|The revert options to validate|
|`messageLength`|`uint256`|The length of the message|


### validateWithdrawalParams

*Validates standard withdrawal parameters*


```solidity
function validateWithdrawalParams(
    bytes memory receiver,
    uint256 amount,
    RevertOptions calldata revertOptions
)
    internal
    pure;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`receiver`|`bytes`|The receiver address|
|`amount`|`uint256`|The amount to withdraw|
|`revertOptions`|`RevertOptions`|The revert options|


### validateWithdrawalAndCallParams

*Validates withdrawal and call parameters*


```solidity
function validateWithdrawalAndCallParams(
    bytes memory receiver,
    uint256 amount,
    bytes calldata message,
    CallOptions calldata callOptions,
    RevertOptions calldata revertOptions
)
    internal
    pure;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`receiver`|`bytes`|The receiver address|
|`amount`|`uint256`|The amount to withdraw|
|`message`|`bytes`|The message to send|
|`callOptions`|`CallOptions`|The call options|
|`revertOptions`|`RevertOptions`|The revert options|


### validateDepositParams

*Validates deposit parameters*


```solidity
function validateDepositParams(
    address zrc20,
    uint256 amount,
    address target,
    address protocolAddress,
    address contractAddress
)
    internal
    pure;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`zrc20`|`address`|The ZRC20 token address|
|`amount`|`uint256`|The amount to deposit|
|`target`|`address`|The target address|
|`protocolAddress`|`address`|The protocol address|
|`contractAddress`|`address`|The contract address|


### validateExecuteParams

*Validates execute parameters*


```solidity
function validateExecuteParams(address zrc20, address target) internal pure;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`zrc20`|`address`|The ZRC20 token address|
|`target`|`address`|The target address|


### validateZetaDepositAndCallParams

*Validates ZETA deposit and call parameters*


```solidity
function validateZetaDepositAndCallParams(
    uint256 amount,
    address target,
    address protocolAddress,
    address contractAddress
)
    internal
    pure;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`amount`|`uint256`|The amount to deposit|
|`target`|`address`|The target address|
|`protocolAddress`|`address`|The protocol address|
|`contractAddress`|`address`|The contract address|


## Errors
### EmptyAddress
Error indicating a empty address was provided.


```solidity
error EmptyAddress();
```

