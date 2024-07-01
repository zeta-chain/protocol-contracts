import { ethers } from "hardhat";

const hre = require("hardhat");

async function main() {
  const gatewayEVMAddress = "0xAD523115cd35a8d4E60B3C0953E0E0ac10418309";
  const gatewayEVM = await hre.ethers.getContractAt("GatewayEVM", gatewayEVMAddress);
  let [, , destination] = await ethers.getSigners();

  const deposit = await gatewayEVM["deposit(address)"](destination.address, { value: ethers.utils.parseEther("1") });

  console.log("deposit hash:", deposit.hash);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
      console.error(error);
      process.exit(1);
  });