#!/bin/bash

echo "Generating protocol addresses..."

npx hardhat addresses --network zeta_testnet > ./data/addresses.testnet.json
npx hardhat addresses --network zeta_mainnet > ./data/addresses.mainnet.json