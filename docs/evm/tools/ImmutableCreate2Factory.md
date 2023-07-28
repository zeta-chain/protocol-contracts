## Ownable

```solidity
import "@zetachain/protocol-contracts/contracts/evm/tools/ImmutableCreate2Factory.sol";
```

Source: https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/evm/tools/ImmutableCreate2Factory.sol

### Function List

* [transferOwnership(newOwner)](#Ownable-transferOwnership-address-)

### Modifiers

### Functions

```
transferOwnership(address newOwner) (external function)
```

<a name="Ownable-transferOwnership-address-"></a>

## ImmutableCreate2Factory

```solidity
import "@zetachain/protocol-contracts/contracts/evm/tools/ImmutableCreate2Factory.sol";
```

Source: https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/evm/tools/ImmutableCreate2Factory.sol

This contract has not yet been fully tested or audited - proceed with
caution and please share any exploits or optimizations you discover.

### Modifier List

* [containsCaller(salt)](#ImmutableCreate2Factory-containsCaller-bytes32-)

### Function List

* [safeCreate2Internal(salt, initializationCode)](#ImmutableCreate2Factory-safeCreate2Internal-bytes32-bytes-)
* [safeCreate2(salt, initializationCode)](#ImmutableCreate2Factory-safeCreate2-bytes32-bytes-)
* [findCreate2Address(salt, initCode)](#ImmutableCreate2Factory-findCreate2Address-bytes32-bytes-)
* [findCreate2AddressViaHash(salt, initCodeHash)](#ImmutableCreate2Factory-findCreate2AddressViaHash-bytes32-bytes32-)
* [hasBeenDeployed(deploymentAddress)](#ImmutableCreate2Factory-hasBeenDeployed-address-)
* [safeCreate2AndTransfer(salt, initializationCode)](#ImmutableCreate2Factory-safeCreate2AndTransfer-bytes32-bytes-)

### Modifiers

```
containsCaller(bytes32 salt) (modifier)
```

<a name="ImmutableCreate2Factory-containsCaller-bytes32-"></a>

Modifier to ensure that the first 20 bytes of a submitted salt match
those of the calling account. This provides protection against the salt
being stolen by frontrunners or other attackers. The protection can also be
bypassed if desired by setting each of the first 20 bytes to zero.

### Functions

```
safeCreate2Internal(bytes32 salt, bytes initializationCode) → address deploymentAddress (internal function)
```

<a name="ImmutableCreate2Factory-safeCreate2Internal-bytes32-bytes-"></a>

```
safeCreate2(bytes32 salt, bytes initializationCode) → address deploymentAddress (public function)
```

<a name="ImmutableCreate2Factory-safeCreate2-bytes32-bytes-"></a>

Create a contract using CREATE2 by submitting a given salt or nonce
along with the initialization code for the contract. Note that the first 20
bytes of the salt must match those of the calling address, which prevents
contract creation events from being submitted by unintended parties.

```
findCreate2Address(bytes32 salt, bytes initCode) → address deploymentAddress (external function)
```

<a name="ImmutableCreate2Factory-findCreate2Address-bytes32-bytes-"></a>

Compute the address of the contract that will be created when
submitting a given salt or nonce to the contract along with the contract's
initialization code. The CREATE2 address is computed in accordance with
EIP-1014, and adheres to the formula therein of
`keccak256( 0xff ++ address ++ salt ++ keccak256(init_code)))[12:]` when
performing the computation. The computed address is then checked for any
existing contract code - if so, the null address will be returned instead.

```
findCreate2AddressViaHash(bytes32 salt, bytes32 initCodeHash) → address deploymentAddress (external function)
```

<a name="ImmutableCreate2Factory-findCreate2AddressViaHash-bytes32-bytes32-"></a>

Compute the address of the contract that will be created when
submitting a given salt or nonce to the contract along with the keccak256
hash of the contract's initialization code. The CREATE2 address is computed
in accordance with EIP-1014, and adheres to the formula therein of
`keccak256( 0xff ++ address ++ salt ++ keccak256(init_code)))[12:]` when
performing the computation. The computed address is then checked for any
existing contract code - if so, the null address will be returned instead.

```
hasBeenDeployed(address deploymentAddress) → bool (external function)
```

<a name="ImmutableCreate2Factory-hasBeenDeployed-address-"></a>

Determine if a contract has already been deployed by the factory to a
given address.

```
safeCreate2AndTransfer(bytes32 salt, bytes initializationCode) → address deploymentAddress (external function)
```

<a name="ImmutableCreate2Factory-safeCreate2AndTransfer-bytes32-bytes-"></a>

