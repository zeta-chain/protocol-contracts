# Deployment

This directory contains script to deploy the protocol contracts on ZetaChain and connected chains.

To execute a deployment script, the general format is:

```
forge script scripts/deploy/<Script>.s.sol --rpc-url <RPC_URL> --private-key <PRIVATE_KEY> --broadcast 
```

A `.env` file should be set up and updated during deployments according to expected env variables in scripts, check `.env.sample` for example on how it should look like.
Currently, `.env.sample` is set with test env variables that can be used to test scripts locally with `anvil` using first account private key.


## Deploying the protocol contracts on connected chains

The GatewayEVM and ERC20Custody contracts must be deploy to setup a new environment on a connected chains.

**GatewayEVM**

The following environment variable must be set in the `.env` file:

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

Note: `verify` and `etherscan-api-key` options are optional for the contract deployment itself.

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

The following environment variable must be set in the `.env` file:

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

## Deterministic deployments

Deployment scripts in `deterministic` uses create2 with Foundry (https://book.getfoundry.sh/tutorials/create2-tutorial) to perform deterministic deployment of  contracts.
This ensures that on every EVM chain `GatewayEVM` contract will be on same address.

Since UUPS proxy is used for the contracts, both implementation and `ERC1967Proxy` are deployed using above technique:

- calculate expected address
- adding a salt to deployment
- basic assertions to verify that deployed address is same as expected

The contract can be upgraded with the following documentation: https://github.com/OpenZeppelin/openzeppelin-foundry-upgrades. It requires:

- deploy new implementation (doesn't need to be deterministic since proxy address doesn't change)
- use plugin to upgrade proxy

