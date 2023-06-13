import { isNetworkName } from "@zetachain/addresses";
import { saveAddress } from "@zetachain/addresses-tools";
import { BigNumber } from "ethers";
import { ethers, network } from "hardhat";

import { getAddress } from "../../../lib/address.helpers";
import {
  ERC20_CUSTODY_SALT_NUMBER_ETH,
  ERC20_CUSTODY_SALT_NUMBER_NON_ETH,
  ERC20_CUSTODY_ZETA_FEE,
  ERC20_CUSTODY_ZETA_MAX_FEE,
} from "../../../lib/contracts.constants";
import { isEthNetworkName } from "../../../lib/contracts.helpers";
import {
  deployContractToAddress,
  saltToHex,
} from "../../../lib/ImmutableCreate2Factory/ImmutableCreate2Factory.helpers";
import { ERC20Custody__factory } from "../../../typechain-types";

export const deterministicDeployERC20Custody = async () => {
  if (!isNetworkName(network.name)) {
    throw new Error(`network.name: ${network.name} isn't supported.`);
  }

  const accounts = await ethers.getSigners();
  const [signer] = accounts;

  const DEPLOYER_ADDRESS = process.env.DEPLOYER_ADDRESS || signer.address;

  const saltNumber = isEthNetworkName(network.name)
    ? ERC20_CUSTODY_SALT_NUMBER_ETH
    : ERC20_CUSTODY_SALT_NUMBER_NON_ETH;
  const saltStr = BigNumber.from(saltNumber).toHexString();

  const zetaToken = getAddress("zetaToken");
  const tss = getAddress("tss");
  const tssUpdater = getAddress("tssUpdater");

  const zetaFee = ERC20_CUSTODY_ZETA_FEE;
  const zetaMaxFee = ERC20_CUSTODY_ZETA_MAX_FEE;

  const immutableCreate2Factory = getAddress("immutableCreate2Factory");

  const salthex = saltToHex(saltStr, DEPLOYER_ADDRESS);
  console.log("SaltHex:", salthex);

  const constructorTypes = [
    "address",
    "address",
    "uint256",
    "uint256",
    "address",
  ];
  const constructorArgs = [
    tss,
    tssUpdater,
    zetaFee.toString(),
    zetaMaxFee.toString(),
    zetaToken,
  ];
  const contractBytecode = ERC20Custody__factory.bytecode;

  const { address } = await deployContractToAddress({
    constructorArgs,
    constructorTypes,
    contractBytecode,
    factoryAddress: immutableCreate2Factory,
    salt: salthex,
    signer,
  });

  console.log("Deployed ERC20 Custody. Address:", address);
  console.log("Constructor Args", constructorArgs);
  // saveAddress("zetaToken", address);
};

if (!process.env.EXECUTE_PROGRAMMATICALLY) {
  deterministicDeployERC20Custody()
    .then(() => process.exit(0))
    .catch((error) => {
      console.error(error);
      process.exit(1);
    });
}
