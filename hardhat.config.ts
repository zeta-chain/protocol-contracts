import "@nomiclabs/hardhat-waffle";
import { HardhatUserConfig } from "hardhat/types";

const config: HardhatUserConfig = {
  solidity: {
    compilers: [{ version: "0.8.7" }],
  },
};

export default config;
