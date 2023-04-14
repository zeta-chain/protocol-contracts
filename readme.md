# ZetaChain Protocol Contracts

This repository contains the smart contracts for ZetaChain. The smart contracts are written in Solidity, and the repository includes scripts to compile the contracts and generate Go bindings.

## Prerequisites

Before you can work with this project, you must have the following installed:

* Node.js
* Yarn
* jq
* abigen

## Getting Started

To get started with this project, you should first clone the repository:

```
git clone https://github.com/zeta-chain/protocol
```

Once you have cloned the repository, you can navigate to the project directory and run the following command to install the project dependencies:

```
yarn install
```

## Compiling Contracts

To compile the contracts, run the following command:

```
yarn compile
```

This will compile the Solidity contracts and output the resulting JSON artifacts to the build/contracts directory.

## Generating Go Bindings

To generate Go bindings for the Solidity contracts, run the following command:

```
yarn generate:go
```

This will use `abigen` to generate Go bindings for the contracts and output the resulting Go files to the `pkg` directory.

## Contributing

If you would like to contribute to this project, please fork the repository and submit a pull request. All contributions are welcome!