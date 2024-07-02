## Worker script

To start local hardhat node execute:

```
yarn localnode
```

To start localnet using local hardhat:
```
yarn localnet
```
This will run worker script, which will deploy all contracts, and listen and react to events, facilitating communication between contracts.
Tasks to interact with localnet are located in `tasks/localnet`. To make use of default contract addresses on localnet, start localnode and localnet from scratch, so contracts are deployed on same addresses. Otherwise, provide custom addresses as tasks parameters.



