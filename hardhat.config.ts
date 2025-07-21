import "@nomicfoundation/hardhat-toolbox";
import "./tasks/addresses";
import { getHardhatConfigNetworks } from "@zetachain/networks";
import { HardhatUserConfig } from "hardhat/config";
import { task } from "hardhat/config";

const config: HardhatUserConfig = {
  networks: {
    ...getHardhatConfigNetworks(),
  },
  solidity: "0.8.26",
};

task("protocolChecksum", "Run EVM protocol checksum verification")
  .addFlag("testnet", "Run checksum for testnet networks")
  .addFlag("mainnet", "Run checksum for mainnet networks")
  .setAction(async (taskArgs, hre) => {
    try {
      const isMainnet = taskArgs.mainnet;
      const networkType = isMainnet ? 'mainnet' : 'testnet';
      console.log(`🔍 Running checksum for ${networkType.toUpperCase()} networks`);

      process.env.NETWORK_TYPE = networkType;

      const checksumScript = await import("./scripts/checksum/protocolChecksum");
      await checksumScript.main();
    } catch (error) {
      console.error("❌ Error running protocol checksum script:", error);
      process.exit(1);
    }
  });

export default config;