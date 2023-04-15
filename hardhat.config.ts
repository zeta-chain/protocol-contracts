import "@nomiclabs/hardhat-waffle";
import "@nomiclabs/hardhat-etherscan";
import "@typechain/hardhat";
import "tsconfig-paths/register";
import "hardhat-abi-exporter";

import { getNetworks } from "@fadeev/networks";
import type { HardhatUserConfig } from "hardhat/types";

const PRIVATE_KEYS =
  process.env.PRIVATE_KEY !== undefined
    ? [`0x${process.env.PRIVATE_KEY}`, `0x${process.env.TSS_PRIVATE_KEY}`]
    : [];

const config: HardhatUserConfig = {
  networks: {
    ...getNetworks(),
  },
  solidity: {
    compilers: [
      { version: "0.5.10" /** For create2 factory */ },
      { version: "0.6.6" /** For uniswap v2 */ },
      { version: "0.7.6" /** For uniswap v3 */ },
      { version: "0.8.7" },
    ],
  },
};

export default config;
