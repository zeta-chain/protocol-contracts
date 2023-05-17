import { isNetworkName } from "@zetachain/addresses";
import { ZetaNonEth__factory as ZetaNonEthFactory } from "../../typechain-types";
import { ethers, network } from "hardhat";

import { getAddress } from "../../lib/address.helpers";

export async function setZetaAddresses(connectorAddress:string , zetaTokenAddress: string) {
  const [owner] = await ethers.getSigners();

  if (!isNetworkName(network.name)) {
    throw new Error(`network.name: ${network.name} isn't supported.`);
  }
  if (owner.address !== getAddress("tss") && owner.address !== getAddress("tssUpdater")) {
    console.log("Only TSS or TSS Updater can set Zeta addresses.");
    console.log("Please execute this step with a valid account.");
    return;
  }

  const [tssSigner] = await ethers.getSigners();

  const contract = ZetaNonEthFactory.connect(zetaTokenAddress, tssSigner);

  console.log("Updating");
  console.log("connectorAddress", connectorAddress);
  console.log("zetaTokenAddress", zetaTokenAddress);
  const tx = await contract.updateTssAndConnectorAddresses(getAddress("tss"), connectorAddress)
  await tx.wait();
  console.log("Updated");
}

