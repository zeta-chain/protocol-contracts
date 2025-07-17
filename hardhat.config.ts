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

task("upgradeProposal", "Create Safe upgrade proposals for GatewayEVM")
  .addFlag("testnet", "Create proposals for testnet networks")
  .addFlag("mainnet", "Create proposals for mainnet networks")
  .addOptionalParam("upgradeData", "Upgrade data to include in the proposal", "0x")
  .setAction(async (taskArgs, hre) => {
    try {
      const isMainnet = taskArgs.mainnet;
      const networkType = isMainnet ? 'mainnet' : 'testnet';
      console.log(`🚀 Creating upgrade proposals for ${networkType.toUpperCase()} networks`);

      process.env.NETWORK_TYPE = networkType;
      if (taskArgs.upgradeData) {
        process.env.UPGRADE_DATA = taskArgs.upgradeData;
      }

      const upgradeScript = await import("./scripts/safe/create-upgrade-proposal");
      await upgradeScript.main();
    } catch (error) {
      console.error("❌ Error running upgrade proposal script:", error);
      process.exit(1);
    }
  });


export default config;
