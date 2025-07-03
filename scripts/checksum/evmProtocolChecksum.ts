import hre from "hardhat";

async function printNetworks() {
  const networks = hre.config.networks;


}

async function main() {
  await printNetworks();
}

if (require.main === module) {
  main().catch(console.error);
}