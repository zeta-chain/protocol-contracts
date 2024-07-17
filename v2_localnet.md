# V2 Contracts - Local Development Environment

Important Notice: The new architecture (V2) is currently in active development. While the high-level interface presented in this document is expected to remain stable, some aspects of the architecture may change.

## Introduction

The new architecture aims to streamline the developer experience for building Universal Apps. Developers will primarily interact with two contracts: GatewayZEVM and GatewayEVM.

* `GatewayEVM`: Deployed on each connected chain to interact with ZetaChain.
* `GatewayZEVM`: Deployed on ZetaChain to interact with connected chains.


## Contract Interfaces

The interface of the gateway contracts is centered around the deposit, withdraw, and call terminology:

* Deposit: Transfer of assets from a connected chain to ZetaChain.
* Withdraw: Transfer of assets from ZetaChain to a connected chain.
* Call: Smart contract call involved in cross-chain transactions.


In consequence, the interface is as follow:
* `GatewayEVM`: `deposit`, `depositAndCall`, and `call`
* `GatewayZEVM`: `withdraw`, `withdrawAndCall`, and `call`

The motivation behind this interface is intuitiveness and simplicity. We support different asset types by using function overloading.

### `GatewayEVM`

* Deposit of native tokens to addresses on ZetaChain:

```
function deposit(address receiver) payable
```

* Deposit of ERC-20 tokens to addresses on ZetaChain:

```
function deposit(address receiver, uint256 amount, address asset)
```

* Deposit of native tokens and smart contract call on ZetaChain:

```
function depositAndCall(address receiver, bytes calldata payload) payable
```

* Deposit of ERC-20 tokens and smart contract call on ZetaChain:

```
depositAndCall(address receiver, uint256 amount, address asset, bytes calldata payload)
```

* Simple universal app contract call

```
function call(address receiver, bytes calldata payload)
```

### `GatewayZEVM`

* Withdraw of ZRC-20 tokens to its native connected chain:

```
function withdraw(bytes memory receiver, uint256 amount, address zrc20)
```

* Withdraw of ZRC-20 tokens and smart contract call on connected chain\:

```
function withdrawAndCall(bytes memory receiver, uint256 amount, address zrc20, bytes calldata message)
```

* Simple call to a contract on a connected chain

```
function call(bytes memory receiver, bytes calldata message) external
```

## Experimenting with the New Architecture

To experiment with the new architecture, you can deploy a local network using Anvil and test the gateways using the following commands:

Clone the repository
```
git clone https://github.com/zeta-chain/protocol-contracts.git
cd protocol-contracts
```

Start the local environment network
Note: `--hide="NODE"` is used to prevent verbose logging
```
yarn
yarn localnet --hide="NODE"
```