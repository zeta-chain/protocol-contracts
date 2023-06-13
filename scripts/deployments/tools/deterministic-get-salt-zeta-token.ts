import { isNetworkName } from "@zetachain/addresses";
import { BigNumber } from "ethers";
import { ethers, network } from "hardhat";

import { getAddress } from "../../../lib/address.helpers";
import { ZETA_INITIAL_SUPPLY } from "../../../lib/contracts.constants";
import { isEthNetworkName } from "../../../lib/contracts.helpers";
import { calculateBestSalt } from "../../../lib/deterministic-deploy.helpers";
import {
  ZetaEth__factory,
  ZetaNonEth__factory,
} from "../../../typechain-types";

const MAX_ITERATIONS = BigNumber.from(1000000);

export async function deterministicDeployGetSaltZetaToken() {
  if (!isNetworkName(network.name)) {
    throw new Error(`network.name: ${network.name} isn't supported.`);
  }

  const accounts = await ethers.getSigners();
  const [signer] = accounts;

  const DEPLOYER_ADDRESS = process.env.DEPLOYER_ADDRESS || signer.address;

  const tss = getAddress("tss");
  const tssUpdater = getAddress("tssUpdater");

  let constructorTypes;
  let constructorArgs;
  let contractBytecode;
  if (isEthNetworkName(network.name)) {
    constructorTypes = ["address", "uint256"];
    constructorArgs = [DEPLOYER_ADDRESS, ZETA_INITIAL_SUPPLY.toString()];
    contractBytecode = ZetaEth__factory.bytecode;
  } else {
    constructorTypes = ["address", "address"];
    constructorArgs = [tss, tssUpdater];
    contractBytecode = ZetaNonEth__factory.bytecode;
  }
  calculateBestSalt(
    MAX_ITERATIONS,
    DEPLOYER_ADDRESS,
    constructorTypes,
    constructorArgs,
    contractBytecode
  );
}

if (!process.env.EXECUTE_PROGRAMMATICALLY) {
  deterministicDeployGetSaltZetaToken()
    .then(() => process.exit(0))
    .catch((error) => {
      console.error(error);
      process.exit(1);
    });
}
