## Deterministic deployments

Note: `.env` file should be set up and updated during deployments according to expected env variables in scripts, check `.env.sample` for example on how it should look like.
Currently, `.env.sample` is set with test env variables that can be used to test scripts locally with `anvil` using first account private key.

`DeployGatewayEVM` script uses create2 with Foundry (https://book.getfoundry.sh/tutorials/create2-tutorial) to perform deterministic deployment of `GatewayEVM` contracts.
This ensures that on every EVM chain `GatewayEVM` contract will be on same address.

Since UUPS proxy is used for `GatewayEVM` contract, both implementation and `ERC1967Proxy` are deployed using above technique:

- calculate expected address
- adding a salt to deployment
- basic assertions to verify that deployed address is same as expected

Remaining deployment script work in similar way as `GatewayEVM` but much simpler because there is no proxy.

`UpgradeGatewayEVM` script uses OpenZeppelin's Foundry Upgrades plugin (https://github.com/OpenZeppelin/openzeppelin-foundry-upgrades), to upgrade `GatewayEVM`:

- deploy new implementation (doesn't need to be deterministic since proxy address doesn't change)
- use plugin to upgrade proxy

To execute deployment script, following format is needed:

```
forge script scripts/deploy/deterministic/<Script>.s.sol --rpc-url <RPC_URL> --private-key <PRIVATE_KEY> --broadcast 
```

To automatically verify contracts deployed with above scripts, verifier options are needed. Example for `GatewayEVM` which will deploy and verify both implementation and proxy contracts (it is the same for other scripts):

- etherscan
```
forge script scripts/deploy/deterministic/DeployGatewayEVM.s.sol --private-key <PRIVATE_KEY> --rpc-url <RPC_URL> --verify --etherscan-api-key <ETHERSCAN_API_KEY> --chain-id <CHAIN_ID> --broadcast 
```

- blockscout
```

forge script scripts/deploy/deterministic/DeployGatewayEVM.s.sol --private-key <PRIVATE_KEY> --rpc-url <RPC_URL> --verify --verifier blockscout --verifier-url <VERIFIER_URL> --chain-id <CHAIN_ID> --broadcast
```