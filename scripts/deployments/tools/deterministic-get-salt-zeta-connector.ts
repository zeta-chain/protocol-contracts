import { isNetworkName } from "@zetachain/addresses";
import { ZetaConnectorEth__factory, ZetaConnectorNonEth__factory } from "../../../typechain-types";
import { BigNumber } from "ethers";
import { ethers, network } from "hardhat";

import { getAddress } from "../../../lib/address.helpers";
import { isEthNetworkName } from "../../../lib/contracts.helpers";
import { calculateBestSalt } from "../../../lib/deterministic-deploy.helpers";

const MAX_ITERATIONS = BigNumber.from(100000);

export async function deterministicDeployGetSaltZetaConnector() {
  if (!isNetworkName(network.name)) {
    throw new Error(`network.name: ${network.name} isn't supported.`);
  }

  const accounts = await ethers.getSigners();
  const [signer] = accounts;

  const DEPLOYER_ADDRESS = process.env.DEPLOYER_ADDRESS || signer.address;

  const zetaToken = getAddress("zetaToken");
  const tss = getAddress("tss");
  const tssUpdater = getAddress("tssUpdater");

  // @todo: decide which address use as pauser
  const constructorTypes = ["address", "address", "address", "address"];
  const constructorArgs = [zetaToken, tss, tssUpdater, tssUpdater];

  let contractBytecode;

  if (isEthNetworkName(network.name)) {
    contractBytecode = ZetaConnectorEth__factory.bytecode;
  } else {
    contractBytecode = ZetaConnectorNonEth__factory.bytecode;
  }

  calculateBestSalt(MAX_ITERATIONS, DEPLOYER_ADDRESS, constructorTypes, constructorArgs, contractBytecode);
}

if (!process.env.EXECUTE_PROGRAMMATICALLY) {
  deterministicDeployGetSaltZetaConnector()
    .then(() => process.exit(0))
    .catch(error => {
      console.error(error);
      process.exit(1);
    });
}
