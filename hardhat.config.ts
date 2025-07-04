import "@nomicfoundation/hardhat-toolbox";
import "./tasks/addresses";
import { HardhatUserConfig } from "hardhat/config";

const config: HardhatUserConfig = {
  networks: {
    zetachain: {
      url: process.env.ZETACHAIN_RPC || "https://zetachain-evm.blockpi.network/v1/rpc/public",
      chainId: 7000,
    },
    ethereum: {
      url: process.env.ETHEREUM_RPC || "https://eth-mainnet.public.blastapi.io",
      chainId: 1,
    },
    bsc: {
      url: process.env.BSC_RPC || "https://bsc-mainnet.public.blastapi.io",
      chainId: 56,
    },
    polygon: {
      url: process.env.POLYGON_RPC || "https://polygon.rpc.subquery.network/public",
      chainId: 137,
    },
    base: {
      url: process.env.BASE_RPC || "https://base.rpc.subquery.network/public",
      chainId: 8453,
    },
    arbitrum: {
      url: process.env.ARBITRUM_RPC || "https://arbitrum-one-rpc.publicnode.com",
      chainId: 42161,
    },
    avalanche: {
      url: process.env.AVALANCHE_RPC || "https://avalanche.drpc.org",
      chainId: 43114,
    },
  },
  solidity: "0.8.26",
};

export default config;
