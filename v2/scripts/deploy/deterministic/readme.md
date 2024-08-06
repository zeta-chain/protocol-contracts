## Deterministic GatewayEVM deployments

`DeployGatewayEVMCreate2` script uses create2 with Foundry (https://book.getfoundry.sh/tutorials/create2-tutorial) to perform deterministic deployment of `GatewayEVM` contracts.
This ensures that on every EVM chain `GatewayEVM` contract will be on same address.

Since UUPS proxy is used for `GatewayEVM` contract, both implementation and `ERC1967Proxy` are deployed using above technique:

- calculate expected address
- adding a salt to deployment
- basic assertions to verify that deployed address is same as expected

`UpgradeGatewayEVM` script uses OpenZeppelin's Foundry Upgrades plugin (https://github.com/OpenZeppelin/openzeppelin-foundry-upgrades), to upgrade `GatewayEVM`:

- deploy new implementation (doesn't need to be deterministic since proxy address doesn't change)
- use plugin to upgrade proxy


