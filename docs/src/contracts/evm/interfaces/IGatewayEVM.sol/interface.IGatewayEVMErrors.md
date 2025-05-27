# IGatewayEVMErrors
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/evm/interfaces/IGatewayEVM.sol)

Interface for the errors used in the GatewayEVM contract.


## Errors
### ExecutionFailed
Error for failed execution.


```solidity
error ExecutionFailed();
```

### DepositFailed
Error for failed deposit.


```solidity
error DepositFailed();
```

### InsufficientETHAmount
Error for insufficient ETH amount.


```solidity
error InsufficientETHAmount();
```

### InsufficientERC20Amount
Error for insufficient ERC20 token amount.


```solidity
error InsufficientERC20Amount();
```

### ZeroAddress
Error for zero address input.


```solidity
error ZeroAddress();
```

### ApprovalFailed
Error for failed token approval.


```solidity
error ApprovalFailed(address token, address spender);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`token`|`address`|The address of the token for which approval failed.|
|`spender`|`address`|The address that was supposed to be approved to spend the tokens.|

### CustodyInitialized
Error for already initialized custody.


```solidity
error CustodyInitialized();
```

### ConnectorInitialized
Error for already initialized connector.


```solidity
error ConnectorInitialized();
```

### NotWhitelistedInCustody
Error when trying to transfer not whitelisted token to custody.


```solidity
error NotWhitelistedInCustody(address token);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`token`|`address`|The address of the token that is not whitelisted in custody.|

### NotAllowedToCallOnCall
Error when trying to call onCall method using arbitrary call.


```solidity
error NotAllowedToCallOnCall();
```

### NotAllowedToCallOnRevert
Error when trying to call onRevert method using arbitrary call.


```solidity
error NotAllowedToCallOnRevert();
```

### PayloadSizeExceeded
Error indicating payload size exceeded in external functions.


```solidity
error PayloadSizeExceeded(uint256 provided, uint256 maximum);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`provided`|`uint256`|The size of the payload that was provided.|
|`maximum`|`uint256`|The maximum allowed payload size.|

