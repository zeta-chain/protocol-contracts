import { isNetworkName } from "@zetachain/addresses";
import { saveAddress } from "@zetachain/addresses-tools";
import { BigNumber } from "ethers";
import { ethers, network } from "hardhat";

import { getAddress } from "../../../lib/address.helpers";
import {
  ZETA_INITIAL_SUPPLY,
  ZETA_TOKEN_SALT_NUMBER_ETH,
  ZETA_TOKEN_SALT_NUMBER_NON_ETH,
} from "../../../lib/contracts.constants";
import { isEthNetworkName } from "../../../lib/contracts.helpers";
import {
  deployContractToAddress,
  saltToHex,
} from "../../../lib/ImmutableCreate2Factory/ImmutableCreate2Factory.helpers";
import {
  ZetaEth__factory,
  ZetaNonEth__factory,
} from "../../../typechain-types";

export async function deterministicDeployZetaToken() {
  if (!isNetworkName(network.name)) {
    throw new Error(`network.name: ${network.name} isn't supported.`);
  }

  const accounts = await ethers.getSigners();
  const [signer] = accounts;

  const DEPLOYER_ADDRESS = process.env.DEPLOYER_ADDRESS || signer.address;

  const saltNumber = isEthNetworkName(network.name)
    ? ZETA_TOKEN_SALT_NUMBER_ETH
    : ZETA_TOKEN_SALT_NUMBER_NON_ETH;
  const saltStr = BigNumber.from(saltNumber).toHexString();

  const tss = getAddress("tss");
  const tssUpdater = getAddress("tssUpdater");
  const immutableCreate2Factory = getAddress("immutableCreate2Factory");

  const salthex = saltToHex(saltStr, DEPLOYER_ADDRESS);
  console.log("SaltHex:", salthex);

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

  const { address } = await deployContractToAddress({
    constructorArgs,
    constructorTypes,
    contractBytecode,
    factoryAddress: immutableCreate2Factory,
    salt: salthex,
    signer,
  });

  console.log("Deployed zetaToken. Address:", address);
  console.log("Constructor Args", constructorArgs);
  // saveAddress("zetaToken", address);
}

if (!process.env.EXECUTE_PROGRAMMATICALLY) {
  deterministicDeployZetaToken()
    .then(() => process.exit(0))
    .catch((error) => {
      console.error(error);
      process.exit(1);
    });
}
