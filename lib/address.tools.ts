import fs from "fs";
import path from "path";

export const addresses = JSON.parse(
  fs.readFileSync(path.resolve(__dirname, "..", "data", "addresses.json")).toString()
);

export declare type ZetaProtocolAddress = "connector" | "immutableCreate2Factory" | "tss" | "tssUpdater" | "zetaToken";
export declare type ZetaProtocolTestNetwork =
  | "baobab_testnet"
  | "bsc_testnet"
  | "goerli_testnet"
  | "mumbai_testnet"
  | "zeta_testnet";

export const zetaProtocolTestNetworks: ZetaProtocolTestNetwork[] = [
  "baobab_testnet",
  "bsc_testnet",
  "goerli_testnet",
  "mumbai_testnet",
  "zeta_testnet",
];

export declare type ZetaProtocolMainNetwork = "etherum_mainnet";
export const zetaProtocolMainNetworks: ZetaProtocolMainNetwork[] = ["etherum_mainnet"];

export declare type ZetaProtocolNetwork = ZetaProtocolMainNetwork | ZetaProtocolTestNetwork;
export const zetaProtocolNetworks: ZetaProtocolNetwork[] = [...zetaProtocolTestNetworks, ...zetaProtocolMainNetworks];

export declare type ZetaProtocolEnviroment = "mainnet" | "testnet";

export const isProtocolNetworkName = (str: string): str is ZetaProtocolNetwork =>
  zetaProtocolNetworks.includes(str as ZetaProtocolNetwork);

export const isTestnetNetwork = (network: ZetaProtocolTestNetwork): boolean => {
  return zetaProtocolTestNetworks.includes(network);
};

export const isMainnetNetwork = (network: ZetaProtocolTestNetwork): boolean => {
  return false;
};

export const getAddress = (address: ZetaProtocolAddress, network: ZetaProtocolNetwork): string => {
  return addresses[network][address];
};
