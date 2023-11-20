# ⚠️ Important Notice: Code Freeze for Security Audit ⚠️

Dear Contributors and Users,
We are committed to ensuring the highest standards of security and reliability in our project. To uphold this commitment, we are currently undergoing a thorough security audit conducted by Code4rena. More info: https://code4rena.com/contests/2023-11-zetachain.

During this period, we have instituted a code freeze on our public repository. This means there will be no new commits, merges, or major changes to the codebase until the audit is complete. This process is crucial to maintain the integrity and consistency of the code being audited.

The audit is scheduled from 20 Nov 9:00 PM GMT+1 to 18 Dec 9:00 PM GMT+1. We appreciate your patience and understanding during this vital phase of our project's development.

During this time, we encourage our community to review the current codebase and documentation. While we won't be merging new changes, we welcome your feedback, which we will address post-audit.

For any questions or concerns, feel free to reach out to us on our [Discord](https://discord.com/invite/zetachain).

Thank you for your cooperation and support!

# ZetaChain Protocol Contracts

This repository contains the smart contracts for ZetaChain. The smart contracts
are written in Solidity, and the repository includes scripts to compile the
contracts and generate Go bindings.

## Importing Protocol Contracts

As a dApp developer, you can install the protocol contracts package into your
project:

```
yarn add --dev @zetachain/protocol-contracts
```

Importing
[`ZetaInterfaces`](https://www.zetachain.com/docs/developers/cross-chain-messaging/connector/)
and `ZetaInteractor` for cross-chain messaging:

```solidity
import "@zetachain/protocol-contracts/contracts/evm/interfaces/ZetaInterfaces.sol";
import "@zetachain/protocol-contracts/contracts/evm/tools/ZetaInteractor.sol";
```

Importing [ZRC20](https://www.zetachain.com/docs/developers/concepts/zrc-20/)
and the [system
contract](https://www.zetachain.com/docs/developers/concepts/system-contract/)
for omni-chain smart contracts:

```solidity
import "@zetachain/protocol-contracts/contracts/zevm/interfaces/IZRC20.sol";
import "@zetachain/protocol-contracts/contracts/zevm/interfaces/zContract.sol";
import "@zetachain/protocol-contracts/contracts/zevm/SystemContract.sol";
```

## Prerequisites

Before you can contribute to this project, you must have the following installed:

- [Node.js](https://nodejs.org/)
- [Yarn](https://yarnpkg.com/)
- [jq](https://stedolan.github.io/jq/)
- [abigen](https://geth.ethereum.org/docs/tools/abigen)

## Getting Started

To get started with this project, you should first clone the repository:

```
git clone https://github.com/zeta-chain/protocol-contracts
```

Once you have cloned the repository, you can navigate to the project directory
and run the following command to install the project dependencies:

```
yarn
```

## Compiling Contracts

To compile the contracts, run the following command:

```
yarn compile
```

This will compile the Solidity contracts and output the resulting JSON artifacts
to the `artifacts` directory.

## Generating Go Bindings

To generate Go bindings for the Solidity contracts, run the following command:

```
yarn generate
```

This will use `abigen` to generate Go bindings for the contracts and output the
resulting Go files to the `pkg` directory.

## Contributing

If you would like to contribute to this project, please fork the repository and
submit a pull request. All contributions are welcome!
