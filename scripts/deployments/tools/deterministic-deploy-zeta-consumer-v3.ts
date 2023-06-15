import { saveAddress } from "@zetachain/addresses-tools";
import { ethers, network } from "hardhat";
import { getAddress, isProtocolNetworkName } from "lib";
import { getExternalAddress } from "lib/address.helpers";

import { ZetaTokenConsumerUniV3__factory } from "../../../typechain-types";

export async function deterministicDeployZetaConsumer() {
  if (!isProtocolNetworkName(network.name)) {
    throw new Error(`network.name: ${network.name} isn't supported.`);
  }
  const accounts = await ethers.getSigners();
  const [signer] = accounts;

  const zetaTokenAddress = getAddress("zetaToken", network.name);

  const uniswapV3Router = getExternalAddress("uniswapV3Router");
  const quoter = getExternalAddress("uniswapV3Quoter");
  const WETH9Address = getExternalAddress("weth9");
  const zetaPoolFee = 500;
  const tokenPoolFee = 3000;

  console.log([zetaTokenAddress, uniswapV3Router, quoter, WETH9Address, zetaPoolFee, tokenPoolFee]);

  const Factory = new ZetaTokenConsumerUniV3__factory(signer);
  const contract = await Factory.deploy(
    zetaTokenAddress,
    uniswapV3Router,
    quoter,
    WETH9Address,
    zetaPoolFee,
    tokenPoolFee
  );
  await contract.deployed();
  const address = contract.address;

  // saveAddress("zetaTokenConsumerUniV3", address);
  console.log("Deployed ZetaConsumer. Address:", address);
}

if (!process.env.EXECUTE_PROGRAMMATICALLY) {
  deterministicDeployZetaConsumer()
    .then(() => process.exit(0))
    .catch((error) => {
      console.error(error);
      process.exit(1);
    });
}
