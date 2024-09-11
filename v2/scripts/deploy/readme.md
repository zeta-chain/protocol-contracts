## Deployment scripts

Note: `.env` file should be set up and updated during deployments according to expected env variables in scripts, check `.env.sample` for example on how it should look like.
Currently, `.env.sample` is set with test env variables that can be used to test scripts locally with `anvil` using first account private key.

To execute deployment script, following format is needed:

```
forge script scripts/deploy/<Script>.s.sol --rpc-url <RPC_URL> --private-key <PRIVATE_KEY> --broadcast 
```