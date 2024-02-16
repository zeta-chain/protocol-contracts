import { BigNumber } from "ethers";
import { parseEther } from "ethers/lib/utils";

import {
  BNB_BSC,
  bscConnectorAddress,
  bscERC20CustodyAddress,
  bscTokenAddress,
  BTC_BTC,
  ETH_ETH,
  ethConnectorAddress,
  ethERC20CustodyAddress,
  ethTokenAddress,
} from "./bytecode.constants";
import {
  cleanAddress,
  cleanNumber,
  compareBytecode,
  encodeAddress,
  encodeNumber,
  getDeployedBytecode,
  getEtherscanBytecode,
  getZetaNodeBytecode,
} from "./bytecode.helpers";

const checkEthConnectorBytecode = async () => {
  const remoteBytecode = await getEtherscanBytecode("eth", ethConnectorAddress);
  const cleanRemoteBytecode = cleanAddress(remoteBytecode, encodeAddress(ethTokenAddress));
  const deployedBytecode = await getDeployedBytecode("ZetaConnector.eth.sol/ZetaConnectorEth", "evm");
  compareBytecode(cleanRemoteBytecode, deployedBytecode);
};

const checkEthCustodyBytecode = async () => {
  const remoteBytecode = await getEtherscanBytecode("eth", ethERC20CustodyAddress);
  let cleanRemoteBytecode = cleanAddress(remoteBytecode, encodeAddress(ethTokenAddress));
  cleanRemoteBytecode = cleanNumber(cleanRemoteBytecode, encodeNumber(parseEther("1000"))); // zetaMaxFee
  const deployedBytecode = await getDeployedBytecode("ERC20Custody.sol/ERC20Custody", "evm");
  compareBytecode(cleanRemoteBytecode, deployedBytecode);
};

const checkBscConnectorBytecode = async () => {
  const remoteBytecode = await getEtherscanBytecode("bsc", bscConnectorAddress);
  const cleanRemoteBytecode = cleanAddress(remoteBytecode, encodeAddress(bscTokenAddress));
  const deployedBytecode = await getDeployedBytecode("ZetaConnector.non-eth.sol/ZetaConnectorNonEth", "evm");
  compareBytecode(cleanRemoteBytecode, deployedBytecode);
};

const checkBscCustodyBytecode = async () => {
  const remoteBytecode = await getEtherscanBytecode("bsc", bscERC20CustodyAddress);
  let cleanRemoteBytecode = cleanAddress(remoteBytecode, encodeAddress(bscTokenAddress));
  cleanRemoteBytecode = cleanNumber(cleanRemoteBytecode, encodeNumber(parseEther("1000"))); // zetaMaxFee
  const deployedBytecode = await getDeployedBytecode("ERC20Custody.sol/ERC20Custody", "evm");
  compareBytecode(cleanRemoteBytecode, deployedBytecode);
};

const checkZRC20ETHBytecode = async () => {
  const remoteBytecode = await getZetaNodeBytecode(ETH_ETH);
  let cleanRemoteBytecode = cleanNumber(remoteBytecode, encodeNumber(BigNumber.from("1"))); // ETH CHAIN ID
  cleanRemoteBytecode = cleanNumber(cleanRemoteBytecode, encodeNumber(BigNumber.from("1"))); // Gas COIN TYPE
  const deployedBytecode = await getDeployedBytecode("ZRC20.sol/ZRC20", "zevm");
  compareBytecode(cleanRemoteBytecode, deployedBytecode);
};

const checkZRC20BTCBytecode = async () => {
  const remoteBytecode = await getZetaNodeBytecode(BTC_BTC);
  let cleanRemoteBytecode = cleanNumber(remoteBytecode, encodeNumber(BigNumber.from("8332"))); // BTC CHAIN ID
  cleanRemoteBytecode = cleanNumber(cleanRemoteBytecode, encodeNumber(BigNumber.from("1"))); // Gas COIN TYPE
  const deployedBytecode = await getDeployedBytecode("ZRC20.sol/ZRC20", "zevm");
  compareBytecode(cleanRemoteBytecode, deployedBytecode);
};

const checkZRC20BSCBytecode = async () => {
  const remoteBytecode = await getZetaNodeBytecode(BNB_BSC);
  let cleanRemoteBytecode = cleanNumber(remoteBytecode, encodeNumber(BigNumber.from("56"))); // BSC CHAIN ID
  cleanRemoteBytecode = cleanNumber(cleanRemoteBytecode, encodeNumber(BigNumber.from("1"))); // Gas COIN TYPE
  const deployedBytecode = await getDeployedBytecode("ZRC20.sol/ZRC20", "zevm");
  compareBytecode(cleanRemoteBytecode, deployedBytecode);
};

const checkBytecode = async () => {
  await checkEthConnectorBytecode();
  await checkEthCustodyBytecode();
  await checkBscConnectorBytecode();
  await checkBscCustodyBytecode();
  await checkZRC20ETHBytecode();
  await checkZRC20BTCBytecode();
  await checkZRC20BSCBytecode();
};

checkBytecode()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
