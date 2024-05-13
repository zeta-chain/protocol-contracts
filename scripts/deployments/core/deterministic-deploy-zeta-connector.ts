import { BigNumber } from "ethers";
import { ethers, network } from "hardhat";
import { getAddress, isProtocolNetworkName } from "lib";

import { getSaltNumber } from "../../../lib/contracts.constants";
import { isEthNetworkName } from "../../../lib/contracts.helpers";
import {
  deployContractToAddress,
  saltToHex,
} from "../../../lib/ImmutableCreate2Factory/ImmutableCreate2Factory.helpers";
import {
  ZetaConnectorEth__factory,
  ZetaConnectorNonEth__factory,
  ZetaConnectorZEVM__factory,
} from "../../../typechain-types";

export const deterministicDeployZetaConnector = async () => {
  if (!isProtocolNetworkName(network.name)) {
    throw new Error(`network.name: ${network.name} isn't supported.`);
  }

  const accounts = await ethers.getSigners();
  const [signer] = accounts;
  const initialBalance = await signer.getBalance();

  const DEPLOYER_ADDRESS = process.env.DEPLOYER_ADDRESS || signer.address;
  // "0x5F0b1a82749cb4E2278EC87F8BF6B618dC71a8bf",
  // "0x8531a5ab847ff5b22d855633c25ed1da3255247e",
  // "0x55122f7590164Ac222504436943FAB17B62F5d7d",
  // "0x55122f7590164Ac222504436943FAB17B62F5d7d",
  const zetaTokenAddress = "0x5F0b1a82749cb4E2278EC87F8BF6B618dC71a8bf";
  const tssAddress = "0x8531a5ab847ff5b22d855633c25ed1da3255247e";
  const tssUpdaterAddress = "0x55122f7590164Ac222504436943FAB17B62F5d7d";

  const factory = new ZetaConnectorZEVM__factory(signer);
  const zetaConnector = await factory.deploy(zetaTokenAddress);
  const rec = await zetaConnector.deployed();
  console.log("ZetaConnector deployed to:", rec.address);

  // const immutableCreate2FactoryAddress = getAddress("immutableCreate2Factory", network.name);

  // const saltNumber = getSaltNumber("zetaConnector", network.name);
  // const saltStr = BigNumber.from(saltNumber).toHexString();

  // const salthex = saltToHex(saltStr, DEPLOYER_ADDRESS);
  // const constructorTypes = ["address", "address", "address", "address"];
  // const constructorArgs = [zetaTokenAddress, tssAddress, tssUpdaterAddress, tssUpdaterAddress];

  // let contractBytecode;
  // if (isEthNetworkName(network.name)) {
  //   contractBytecode = ZetaConnectorEth__factory.bytecode;
  // } else {
  //   contractBytecode = ZetaConnectorNonEth__factory.bytecode;
  // }

  // const { address } = await deployContractToAddress({
  //   constructorArgs,
  //   constructorTypes,
  //   contractBytecode,
  //   factoryAddress: immutableCreate2FactoryAddress,
  //   salt: salthex,
  //   signer,
  // });

  // const finalBalance = await signer.getBalance();
  // console.log("Deployed ZetaConnector. Address:", address);
  // console.log("Constructor Args", constructorArgs);
  // console.log("ETH spent:", initialBalance.sub(finalBalance).toString());

  return address;
};

if (!process.env.EXECUTE_PROGRAMMATICALLY) {
  deterministicDeployZetaConnector()
    .then(() => process.exit(0))
    .catch((error) => {
      console.error(error);
      process.exit(1);
    });
}
