"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (k !== "default" && Object.prototype.hasOwnProperty.call(mod, k)) __createBinding(result, mod, k);
    __setModuleDefault(result, mod);
    return result;
};
Object.defineProperty(exports, "__esModule", { value: true });
require("@nomiclabs/hardhat-waffle");
require("@nomiclabs/hardhat-etherscan");
require("@typechain/hardhat");
require("tsconfig-paths/register");
require("hardhat-abi-exporter");
const dotenv = __importStar(require("dotenv"));
dotenv.config();
const PRIVATE_KEYS = process.env.PRIVATE_KEY !== undefined ? [`0x${process.env.PRIVATE_KEY}`] : [];
const config = {
    networks: {
        athens: {
            accounts: PRIVATE_KEYS,
            // chainId: 7001,
            gas: 5000000,
            gasPrice: 80000000000,
            url: `https://rpc.ankr.com/zetachain_evm_testnet`,
        },
        "bsc-localnet": {
            gas: 5000000,
            gasPrice: 80000000000,
            url: "http://localhost:8120",
        },
        "bsc-testnet": {
            accounts: PRIVATE_KEYS,
            gas: 5000000,
            gasPrice: 80000000000,
            url: `https://data-seed-prebsc-1-s1.binance.org:8545`,
        },
        "eth-localnet": {
            gas: 2100000,
            gasPrice: 80000000000,
            url: "http://localhost:8100",
        },
        "eth-mainnet": {
            accounts: PRIVATE_KEYS,
            url: "https://rpc.ankr.com/eth",
        },
        goerli: {
            accounts: PRIVATE_KEYS,
            gas: 2100000,
            gasPrice: 38000000000,
            url: `https://rpc.ankr.com/eth_goerli`,
        },
        hardhat: {
            chainId: 1337,
            forking: {
                blockNumber: 14672712,
                url: "https://rpc.ankr.com/eth",
            },
        },
        "klaytn-baobab": {
            accounts: PRIVATE_KEYS,
            chainId: 1001,
            gas: 6000000,
            gasPrice: 54250000000,
            url: "https://api.baobab.klaytn.net:8651",
        },
        "klaytn-cypress": {
            accounts: PRIVATE_KEYS,
            chainId: 8217,
            gas: 2100000,
            gasPrice: 8000000000,
            url: "https://scope.klaytn.com/",
        },
        "polygon-localnet": {
            gas: 5000000,
            gasPrice: 80000000000,
            url: "http://localhost:8140",
        },
        "polygon-mumbai": {
            accounts: PRIVATE_KEYS,
            gas: 5000000,
            gasPrice: 80000000000,
            url: "https://rpc.ankr.com/polygon_mumbai",
        },
        ropsten: {
            accounts: PRIVATE_KEYS,
            gas: 9000000,
            gasPrice: 80000000000,
            url: "https://ropsten.infura.io/v3/",
        },
    },
    solidity: {
        compilers: [
            { version: "0.4.19" /** For zevm/wzeta.sol */ },
            { version: "0.5.10" /** For create2 factory */ },
            { version: "0.5.16", settings: {
                    optimizer: {
                        enabled: true,
                        runs: 999999,
                    }
                } },
            { version: "0.6.6", settings: {
                    optimizer: {
                        enabled: true,
                        runs: 999999,
                    }
                } /** For uniswap v2 */ },
            { version: "0.7.6" /** For uniswap v3 */ },
            { version: "0.8.7" },
        ],
    },
};
exports.default = config;
