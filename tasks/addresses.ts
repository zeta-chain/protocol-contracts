import { task } from "hardhat/config";
import { HardhatRuntimeEnvironment } from "hardhat/types";
import axios, { AxiosResponse } from "axios";
import { SystemContract__factory } from "../typechain-types/factories/contracts/zevm/SystemContract.sol/SystemContract__factory";
import {
  ZetaConnectorBase__factory,
  ZetaConnectorEth__factory,
  ZetaConnectorNonEth__factory,
  ZetaConnectorZEVM__factory,
} from "@typechain-types";
import { ZetaConnector__factory } from "../typechain-types";
import { zetaConnectorNonEthSol } from "typechain-types/factories/contracts/evm";

const ATHENS_EVM_RPC = "https://zetachain-athens-evm.blockpi.network/v1/rpc/public";
const ATHENS_TENDERMINT_RPC = "http://46.4.15.110:1317";

const fetchChains = async (chainObject: any) => {
  const URL = `${ATHENS_TENDERMINT_RPC}/zeta-chain/observer/supportedChains`;
  try {
    const response: AxiosResponse<any> = await axios.get(URL);

    if (response.status === 200) {
      const chains: string[] = response.data.chains;
      chains.forEach((chain: any) => {
        chainObject[chain.chain_name.toLowerCase()] = {
          chainId: chain.chain_id,
        };
      });
    } else {
      console.error("Error fetching chains:", response.status, response.statusText);
    }
  } catch (error) {
    console.error("Error fetching chains:", error);
  }
};

const fetchTssData = async (chainObject: any) => {
  const URL = `${ATHENS_TENDERMINT_RPC}/zeta-chain/crosschain/get_tss_address`;
  try {
    const tssResponse: AxiosResponse<any> = await axios.get(URL);

    if (tssResponse.status === 200) {
      for (const chain in chainObject) {
        const { btc, eth } = tssResponse.data;
        if (chain.includes("zeta")) continue;
        chainObject[chain]["tss"] = chain.includes("btc") ? btc : eth;
      }
    } else {
      console.error("Error fetching TSS data:", tssResponse.status, tssResponse.statusText);
    }
  } catch (error) {
    console.error("Error fetching TSS data:", error);
  }
};

const fetchSystemContract = async (chainObject: any) => {
  const URL = `${ATHENS_TENDERMINT_RPC}/zeta-chain/fungible/system_contract`;
  try {
    const systemContractResponse: AxiosResponse<any> = await axios.get(URL);

    if (systemContractResponse.status === 200) {
      if (!chainObject["zeta_testnet"]) {
        chainObject["zeta_testnet"] = {};
      }
      chainObject["zeta_testnet"]["systemContract"] = systemContractResponse.data.SystemContract.system_contract;
      chainObject["zeta_testnet"]["connectorZEVM"] = systemContractResponse.data.SystemContract.connector_zevm;
    } else {
      console.error("Error fetching system contract:", systemContractResponse.statusText);
    }
  } catch (error) {
    console.error("Error fetching system contract:", error);
  }
};

const fetchForeignCoinsData = async (chainObject: any) => {
  const URL = `${ATHENS_TENDERMINT_RPC}/zeta-chain/fungible/foreign_coins`;
  try {
    const foreignCoinsResponse: AxiosResponse<any> = await axios.get(URL);
    if (foreignCoinsResponse.status === 200) {
      foreignCoinsResponse.data.foreignCoins.forEach((coin: any) => {
        for (const key in chainObject) {
          if (chainObject.hasOwnProperty(key)) {
            const chain = chainObject[key];
            if (chain.chainId === coin.foreign_chain_id) {
              chain.zrc20ContractAddress = coin.zrc20_contract_address;
            }
          }
        }
      });
    } else {
      console.error("Error fetching foreign coins data:", foreignCoinsResponse.status, foreignCoinsResponse.statusText);
    }
  } catch (error) {
    console.error("Error fetching foreign coins data:", error);
  }
};

const fetchAthensAddresses = async (chainObject: any, hre: HardhatRuntimeEnvironment) => {
  const zeta = "zeta_testnet";
  if (chainObject[zeta] && chainObject[zeta]["systemContract"]) {
    const athens = chainObject[zeta];
    const systemContract = athens["systemContract"] as string;
    const provider = new hre.ethers.providers.JsonRpcProvider(ATHENS_EVM_RPC);
    const sc = SystemContract__factory.connect(systemContract, provider);
    try {
      athens["uniswapv2FactoryAddress"] = await sc.uniswapv2FactoryAddress();
      athens["wZetaContractAddress"] = await sc.wZetaContractAddress();
      athens["uniswapv2Router02Address"] = await sc.uniswapv2Router02Address();
      athens["zetaConnectorZEVMAddress"] = await sc.zetaConnectorZEVMAddress();
      athens["fungibleModuleAddress"] = await sc.FUNGIBLE_MODULE_ADDRESS();
    } catch (error) {
      console.error("Error fetching addresses from ZetaChain:", error);
    }
  } else {
    console.error("ZetaChain system contract not found.");
  }
};

const fetchChainSpecificAddresses = async (chainObject: any) => {
  const URL = `${ATHENS_TENDERMINT_RPC}/zeta-chain/observer/get_client_params_for_chain`;
  try {
    for (const chain in chainObject) {
      if (chain.includes("zeta") || chain.includes("btc")) continue;
      const id = chainObject[chain].chainId;
      const response: AxiosResponse<any> = await axios.get(`${URL}/${id}`);
      if (response.status === 200) {
        chainObject[chain]["zetaTokenContractAddress"] = response.data.core_params.zeta_token_contract_address;
        chainObject[chain]["connectorContractAddress"] = response.data.core_params.connector_contract_address;
        chainObject[chain]["erc20CustodyContractAddress"] = response.data.core_params.erc20_custody_contract_address;
      } else {
        console.error("Error fetching chain data:", response.status, response.statusText);
      }
    }
  } catch (error) {
    console.error("Error fetching chain data:", error);
  }
};

const fetchPauserUpdater = async (chainObject: any, hre: HardhatRuntimeEnvironment) => {
  try {
    for (const chain in chainObject) {
      const provider = new hre.ethers.providers.JsonRpcProvider(ATHENS_EVM_RPC);
    }
  } catch (error) {
    console.error("Error fetching addresses from ZetaChain:", error);
  }
  const zeta = "zeta_testnet";
  if (chainObject[zeta] && chainObject[zeta]["systemContract"]) {
    const athens = chainObject[zeta];
    const systemContract = athens["systemContract"] as string;
    const provider = new hre.ethers.providers.JsonRpcProvider(ATHENS_EVM_RPC);
    const sc = SystemContract__factory.connect(systemContract, provider);
    try {
      athens["uniswapv2FactoryAddress"] = await sc.uniswapv2FactoryAddress();
      athens["wZetaContractAddress"] = await sc.wZetaContractAddress();
      athens["uniswapv2Router02Address"] = await sc.uniswapv2Router02Address();
      athens["zetaConnectorZEVMAddress"] = await sc.zetaConnectorZEVMAddress();
      athens["fungibleModuleAddress"] = await sc.FUNGIBLE_MODULE_ADDRESS();
    } catch (error) {
      console.error("Error fetching addresses from ZetaChain:", error);
    }
  } else {
    console.error("ZetaChain system contract not found.");
  }
};

const main = async (args: any, hre: HardhatRuntimeEnvironment) => {
  const chainObject: any = {};

  await fetchChains(chainObject);
  await fetchTssData(chainObject);
  await fetchSystemContract(chainObject);
  await fetchForeignCoinsData(chainObject);
  await fetchAthensAddresses(chainObject, hre);
  await fetchChainSpecificAddresses(chainObject);

  console.log(JSON.stringify(chainObject, null, 2));
};

task("addresses", "").setAction(main);
