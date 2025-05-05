# UniversalContract
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/main/v2/contracts/zevm/interfaces/UniversalContract.sol)

Abstract contract for contracts that can receive cross-chain calls on ZetaChain.

*Contracts extending this abstract contract can handle incoming cross-chain messages
and execute logic based on the provided context, token, and message payload.*


## State Variables
### gateway
Reference to the ZetaChain Gateway contract


```solidity
IGatewayZEVM public immutable gateway;
```


## Functions
### onlyGateway

Restricts function access to only the gateway contract

*Used on functions that process cross-chain messages to ensure they're only called through the Gateway,
where message validation occurs.
Important for security in functions like `onCall()` and `onRevert()` that handle incoming cross-chain
operations.*


```solidity
modifier onlyGateway();
```

### constructor

Constructor to initialize the contract with the gateway address


```solidity
constructor(address payable gatewayAddress);
```

### onCall

Function to handle cross-chain calls


```solidity
function onCall(
    MessageContext calldata context,
    address zrc20,
    uint256 amount,
    bytes calldata message
)
    external
    virtual;
```

## Errors
### Unauthorized
Error thrown when a function is called by an unauthorized address


```solidity
error Unauthorized();
```

