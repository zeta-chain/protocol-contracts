# UniversalContract
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/zevm/interfaces/UniversalContract.sol)

Abstract contract for contracts that can receive cross-chain calls on ZetaChain.

*Contracts extending this abstract contract can handle incoming cross-chain messages
and execute logic based on the provided context, token, and message payload.*


## State Variables
### registry
Reference to the ZetaChain Registry contract


```solidity
ICoreRegistry public constant registry = ICoreRegistry(0x7c591652f159496b14e15616F0948a6d63b585E8);
```


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

Initializes the contract by retrieving the gateway address from the registry

*Fetches the gateway contract address for the current chain from the registry.
If the gateway is not active or not found, the gateway will remain uninitialized (address(0)).*


```solidity
constructor();
```

### onCall

Function to handle cross-chain calls (use for ZETA deposits)


```solidity
function onCall(MessageContext calldata context, bytes calldata message) external payable virtual;
```

### onCall

Function to handle cross-chain calls (use for ZRC20 deposits)


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

