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
    zetachainTestnet: {
      url: process.env.ZETACHAIN_RPC || "https://zetachain-testnet.public.blastapi.io",
      chainId: 7001,
    },
    sepolia: {
      url: process.env.ETHEREUM_RPC || "https://eth-sepolia.public.blastapi.io",
      chainId: 11155111,
    },
    bscTestnet: {
      url: process.env.BSC_RPC || "https://bsc-testnet.public.blastapi.io",
      chainId: 97,
    },
    polygonAmoy: {
      url: process.env.POLYGON_RPC || "https://polygon-amoy.drpc.org",
      chainId: 80002,
    },
    baseSepolia: {
      url: process.env.BASE_RPC || "https://sepolia.base.org",
      chainId: 84532,
    },
    arbitrumSepolia: {
      url: process.env.ARBITRUM_RPC || "https://arbitrum-sepolia-rpc.publicnode.com",
      chainId: 421614,
    },
    avalancheFuji: {
      url: process.env.AVALANCHE_RPC || "https://avalanche-fuji.drpc.org",
      chainId: 43113,
    },
  },
  solidity: "0.8.26",
};

export default config;
