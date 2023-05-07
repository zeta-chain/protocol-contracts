import { task } from "hardhat/config";
import { HardhatRuntimeEnvironment } from "hardhat/types";
import axios, { AxiosResponse } from "axios";
import { SystemContract__factory } from "../typechain-types/factories/contracts/zevm/SystemContract.sol/SystemContract__factory";

const ATHENS_EVM_RPC = "https://rpc.ankr.com/zetachain_evm_testnet";
const ATHENS_TENDERMINT_RPC = "https://api.athens2.zetachain.com";

interface ChainObject {
  [chainName: string]: Record<string, unknown>;
}

interface ApiResponse {
  Chains: string[];
}

interface ForeignCoinsResponse {
  foreignCoins: {
    index: string;
    zrc20ContractAddress: string;
    erc20ContractAddress: string;
    foreignChain: string;
    decimals: number;
    name: string;
    symbol: string;
    coinType: string;
  }[];
  pagination: {
    next_key: string | null;
    total: string;
  };
}

interface TssResponse {
  TSS: {
    creator: string;
    index: string;
    chain: string;
    address: string;
    pubkey: string;
    signer: string[];
    finalizedZetaHeight: string;
  }[];
}

interface SystemContractResponse {
  SystemContract: {
    systemContract: string;
  };
}

const fetchChains = async (chainObject: ChainObject) => {
  const URL = `${ATHENS_TENDERMINT_RPC}/zeta-chain/observer/supportedChains`;
  try {
    const response: AxiosResponse<ApiResponse> = await axios.get(URL);

    if (response.status === 200) {
      const chains: string[] = response.data.Chains;
      chains.forEach((chain: string) => {
        chainObject[chain.toLowerCase()] = {};
      });
    } else {
      console.error(
        "Error fetching chains:",
        response.status,
        response.statusText
      );
    }
  } catch (error) {
    console.error("Error fetching chains:", error);
  }
};

const fetchTssData = async (chainObject: ChainObject) => {
  const URL = `${ATHENS_TENDERMINT_RPC}/zeta-chain/crosschain/TSS`;
  try {
    const tssResponse: AxiosResponse<TssResponse> = await axios.get(URL);

    if (tssResponse.status === 200) {
      tssResponse.data.TSS.forEach((tssItem) => {
        if (chainObject[tssItem.chain.toLowerCase()]) {
          chainObject[tssItem.chain.toLowerCase()]["tss"] = tssItem.address;
        }
      });
    } else {
      console.error(
        "Error fetching TSS data:",
        tssResponse.status,
        tssResponse.statusText
      );
    }
  } catch (error) {
    console.error("Error fetching TSS data:", error);
  }
};

const fetchSystemContract = async (chainObject: ChainObject) => {
  const URL = `${ATHENS_TENDERMINT_RPC}/zeta-chain/zetacore/fungible/system_contract`;
  try {
    const systemContractResponse: AxiosResponse<SystemContractResponse> =
      await axios.get(URL);

    if (systemContractResponse.status === 200) {
      if (!chainObject["athens"]) {
        chainObject["athens"] = {};
      }
      chainObject["athens"]["systemContract"] =
        systemContractResponse.data.SystemContract.systemContract;
    } else {
      console.error(
        "Error fetching system contract:",
        systemContractResponse.statusText
      );
    }
  } catch (error) {
    console.error("Error fetching system contract:", error);
  }
};

const fetchForeignCoinsData = async (chainObject: ChainObject) => {
  const URL = `${ATHENS_TENDERMINT_RPC}/zeta-chain/zetacore/fungible/foreign_coins`;
  try {
    const foreignCoinsResponse: AxiosResponse<ForeignCoinsResponse> =
      await axios.get(URL);
    if (foreignCoinsResponse.status === 200) {
      foreignCoinsResponse.data.foreignCoins.forEach((coin) => {
        const chain = coin.foreignChain.toLowerCase();
        if (chainObject[chain]) {
          chainObject[chain]["zrc20ContractAddress"] =
            coin.zrc20ContractAddress;
        } else {
          chainObject[chain] = {
            zrc20ContractAddress: coin.zrc20ContractAddress,
          };
        }
      });
    } else {
      console.error(
        "Error fetching foreign coins data:",
        foreignCoinsResponse.status,
        foreignCoinsResponse.statusText
      );
    }
  } catch (error) {
    console.error("Error fetching foreign coins data:", error);
  }
};

const fetchAthensAddresses = async (
  chainObject: ChainObject,
  hre: HardhatRuntimeEnvironment
) => {
  if (chainObject["athens"] && chainObject["athens"]["systemContract"]) {
    const athens = chainObject["athens"];
    const systemContract = athens["systemContract"] as string;
    const provider = new hre.ethers.providers.JsonRpcProvider(ATHENS_EVM_RPC);
    const sc = SystemContract__factory.connect(systemContract, provider);
    try {
      athens["uniswapv2FactoryAddress"] = await sc.uniswapv2FactoryAddress();
      athens["wZetaContractAddress"] = await sc.wZetaContractAddress();
      athens["uniswapv2Router02Address"] = await sc.uniswapv2Router02Address();
    } catch (error) {
      console.error("Error fetching addresses from Athens:", error);
    }
  } else {
    console.error("Athens system contract not found.");
  }
};

const main = async (args: any, hre: HardhatRuntimeEnvironment) => {
  const chainObject: ChainObject = {};

  await fetchChains(chainObject);
  await fetchTssData(chainObject);
  await fetchSystemContract(chainObject);
  await fetchForeignCoinsData(chainObject);
  await fetchAthensAddresses(chainObject, hre);

  console.log(JSON.stringify(chainObject, null, 2));
};

task("addresses", "").setAction(main);
