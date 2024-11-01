# Deployment

This directory contains script to deploy the protocol contracts on ZetaChain and connected chains.

To execute a deployment script, the general format is:

```
forge script scripts/deploy/<Script>.s.sol --rpc-url <RPC_URL> --private-key <PRIVATE_KEY> --broadcast 
```

A `.env` file should be set up and updated during deployments according to expected env variables in scripts, check `.env.sample` for example on how it should look like.
Currently, `.env.sample` is set with test env variables that can be used to test scripts locally with `anvil` using first account private key.


## Deploying the protocol contracts on connected chains

The GatewayEVM and ERC20Custody contracts must be deployed to setup a new environment on a connected chain.

**GatewayEVM**

The following environment variables must be set in the `.env` file:

- `TSS_ADDRESS`: address of the TSS used on the network
- `GATEWAY_ADMIN_ADDRESS_EVM`: address of the admin
- `ZETA_ERC20_EVM`: address of the ZETA token on the network

Once setup, the contract can be deployed:

```
forge script scripts/deploy/deterministic/DeployGatewayEVM.s.sol \
  --private-key <PRIVATE_KEY> \
  --rpc-url <RPC_URL> \
  --verify \
  --etherscan-api-key <ETHERSCAN_API_KEY> \
  --chain-id <CHAIN_ID> \
  --broadcast
```

Note: `verify` and `etherscan-api-key` options are optional for the contract deployment itself. Some other explorers require different options, for example Blockscout:
```
  --verify \
  --verifier blockscout \
  --verifier-url <VERIFIER_URL> \
```

**ERC20Custody**

In addition to the previous environment variables, the following environment variables must be set:

- `ERC20_CUSTODY_ADMIN_ADDRESS_EVM`: address of the admin
- `GATEWAY_PROXY_EVM`: address of the **proxy** gateway contract deployed in previous step

Once setup, the contract can be deployed:

```
forge script scripts/deploy/deterministic/DeployERC20Custody.s.sol \
  --private-key <PRIVATE_KEY> \
  --rpc-url <RPC_URL> \
  --verify \
  --etherscan-api-key <ETHERSCAN_API_KEY> \
  --chain-id <CHAIN_ID> \
  --broadcast
```

## Deploying the protocol contracts on ZetaChain

Since ZRC20s are deployed by the protocol, only the `GatewayZEVM` contract needs to be deployed manually on ZetaChain.

The following environment variables must be set in the `.env` file:

- `GATEWAY_ADMIN_ADDRESS_ZEVM`: address of the admin 
- `WZETA`: wrapped ZETA contract address

Once setup, the contract can be deployed:

```
forge script scripts/deploy/deterministic/DeployGatewayZEVM.s.sol \
  --private-key <PRIVATE_KEY> \
  --rpc-url <RPC_URL> \
  --verify \
  --verifier blockscout \
  --verifier-url <VERIFIER_URL> \
  --chain-id <CHAIN_ID> \
  --broadcast
```

## Deploying a ZRC20 reference contract

ZRC20 contract is upgradable by the protocol but doesn't follow the `ERC1967Proxy` standard.
The process requires to:
- Deploy a reference for the ZRC20 to upgrade
- Use the `MsgUpdateContractBytecode` message of the blockchain to upgrade the contract using the reference bytecode

The reference must be deployed using the same constructor argument as the base contract.
To obtain the arguments, the following utility script can be used:
```
./scripts/utils/get_zrc20_info.sh <zrc20_address> <rpc>
```

Note: the script requires `cast` CLI (Foundry suite) to be installed.

The command prints the environment variable to put in the `.env` file to deploy the new contract.
Example:
```
SYSTEM_CONTRACT=0xEdf1c3275d13489aCdC6cD6eD246E72458B8795B
ZRC20_NAME=ZetaChain ZRC20 SOL on Solana Devnet
ZRC20_SYMBOL=SOL.SOLANA
ZRC20_DECIMALS=9
ZRC20_CHAIN_ID=901
ZRC20_COIN_TYPE=1
ZRC20_GAS_LIMIT=5000
```

Once the environment variables are set, the ZRC20 can be deployed:
```
forge script scripts/deploy/DeployZRC20.s.sol \
  --private-key <PRIVATE_KEY> \
  --rpc-url <RPC_URL> \
  --verify \
  --verifier blockscout \
  --verifier-url <VERIFIER_URL> \
  --chain-id <CHAIN_ID> \
  --broadcast
```

After deployment, the following utility script allows to verify the ZRC20 has been deployed with the correct constructor arguments:
```
./scripts/utils/compare_zrc20_info.sh <base_contract> <reference_contract> <rpc>
```

## Deterministic deployments

Deployment scripts in `deterministic` uses create2 with Foundry (https://book.getfoundry.sh/tutorials/create2-tutorial) to perform deterministic deployment of contracts.
This ensures that the GatewayEVM contract will have the same address on every EVM chain.

Since UUPS proxy is used for the contracts, both implementation and `ERC1967Proxy` are deployed using above technique:

- calculate expected address
- adding a salt to deployment
- basic assertions to verify that deployed address is same as expected

The contract can be upgraded with the following documentation: https://github.com/OpenZeppelin/openzeppelin-foundry-upgrades. It requires:

- deploy new implementation (doesn't need to be deterministic since proxy address doesn't change)
- use plugin to upgrade proxy


## All deployment scripts

- `deterministic/DeployERC20Custody.s.sol`: deploy the ERC20 custody contract on a connected chain
- `deterministic/DeployERC20CustodyImplementation.s.sol`: deploy the ERC20 custody contract implementation on a connected chain for a contract upgrade
- `deterministic/DeployGatewayEVM.s.sol`: deploy the gateway contract on a connected chain
- `deterministic/DeployGatewayEVMImplementation.s.sol`: deploy the gateway contract implementation on a connected chain for a contract upgrade
- `deterministic/DeployGatewayZEVM.s.sol`: deploy the gateway contract on ZetaChain
- `deterministic/DeployGatewayZEVMImplementation.s.sol`: deploy the gateway contract implementation on ZetaChain for a contract upgrade
- `deterministic/TestERC20.s.sol`: deploy a ERC20 for test purpose
- `deterministic/ZetaConnectorNonNative.s.sol`: deploy the ZETA connector contract on a connected chain, currently not used
- `deterministic/UpgradeGatewayEVM.s.sol`: upgrade the GatewayEVM contract to a test contract implementation, used for test purposes only
- `DeployZRC20.s.sol`: deploy a ZRC20 for test purpose or for the ZRC20 upgrade process