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

