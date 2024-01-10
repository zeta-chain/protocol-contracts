import { task } from "hardhat/config";
import { HardhatRuntimeEnvironment } from "hardhat/types";
import axios, { AxiosResponse } from "axios";
import { ERC20Custody__factory } from "../typechain-types/factories/contracts/evm/ERC20Custody__factory";
import { SystemContract__factory } from "../typechain-types/factories/contracts/zevm/SystemContract.sol/SystemContract__factory";
import {
  ZetaConnectorBase__factory,
  ZetaConnectorEth__factory,
  ZetaConnectorNonEth__factory,
  ZetaConnectorZEVM__factory,
} from "@typechain-types";
import { ZetaConnector__factory } from "../typechain-types";
import { zetaConnectorNonEthSol } from "../typechain-types/factories/contracts/evm";
import { getEndpoints } from "@zetachain/networks";

const ATHENS_EVM_RPC = "https://zetachain-athens-evm.blockpi.network/v1/rpc/public";
const ATHENS_TENDERMINT_RPC = "https://zetachain-athens.blockpi.network/lcd/v1/public";

const fetchChains = async () => {
  const URL = `${ATHENS_TENDERMINT_RPC}/zeta-chain/observer/supportedChains`;
  try {
    const response: AxiosResponse<any> = await axios.get(URL);

    if (response.status === 200) {
      return response.data.chains;
    } else {
      console.error("Error fetching chains:", response.status, response.statusText);
    }
  } catch (error) {
    console.error("Error fetching chains:", error);
  }
};

const fetchTssData = async (chains: any, addresses: any) => {
  const URL = `${ATHENS_TENDERMINT_RPC}/zeta-chain/crosschain/get_tss_address`;
  try {
    const tssResponse: AxiosResponse<any> = await axios.get(URL);

    if (tssResponse.status === 200) {
      chains.forEach((chain: any) => {
        const { btc, eth } = tssResponse.data;
        if (chain.chain_name === "zeta_testnet") return;
        addresses.push({
          chain_id: chain.chain_id,
          chain_name: chain.chain_name,
          type: "tss",
          category: "omnichain",
          address: chain.chain_name === "btc_testnet" ? btc : eth,
        });
      });
    } else {
      console.error("Error fetching TSS data:", tssResponse.status, tssResponse.statusText);
    }
  } catch (error) {
    console.error("Error fetching TSS data:", error);
  }
};

const fetchSystemContract = async (addresses: any) => {
  const URL = `${ATHENS_TENDERMINT_RPC}/zeta-chain/fungible/system_contract`;
  try {
    const systemContractResponse: AxiosResponse<any> = await axios.get(URL);

    if (systemContractResponse.status === 200) {
      addresses.push({
        chain_id: 7001,
        chain_name: "zeta_testnet",
        type: "systemContract",
        category: "omnichain",
        address: systemContractResponse.data.SystemContract.system_contract,
      });
      addresses.push({
        chain_id: 7001,
        chain_name: "zeta_testnet",
        type: "connectorZEVM",
        category: "messaging",
        address: systemContractResponse.data.SystemContract.connector_zevm,
      });
    } else {
      console.error("Error fetching system contract:", systemContractResponse.statusText);
    }
  } catch (error) {
    console.error("Error fetching system contract:", error);
  }
};

const fetchForeignCoinsData = async (chains: any, addresses: any) => {
  const URL = `${ATHENS_TENDERMINT_RPC}/zeta-chain/fungible/foreign_coins`;
  try {
    const foreignCoinsResponse: AxiosResponse<any> = await axios.get(URL);
    if (foreignCoinsResponse.status === 200) {
      foreignCoinsResponse.data.foreignCoins.forEach((token: any) => {
        addresses.push({
          chain_id: token.foreign_chain_id,
          chain_name: chains.find((c: any) => c.chain_id === token.foreign_chain_id)?.chain_name,
          coin_type: token.coin_type,
          address: token.asset,
          symbol: token.symbol,
          zrc20: token.zrc20_contract_address,
          category: "omnichain",
          decimals: token.decimals, // TODO: dynamically fetch from contract (to verify)
        });
        addresses.push({
          chain_id: 7001,
          chain_name: "zeta_testnet",
          coin_type: "ZRC20",
          address: token.zrc20_contract_address,
          symbol: token.name,
          asset: token.asset,
          category: "omnichain",
          decimals: 18, // TODO: dynamically fetch from contract
        });
      });
    } else {
      console.error("Error fetching foreign coins data:", foreignCoinsResponse.status, foreignCoinsResponse.statusText);
    }
  } catch (error) {
    console.error("Error fetching foreign coins data:", error);
  }
};

const fetchAthensAddresses = async (addresses: any, hre: HardhatRuntimeEnvironment) => {
  const systemContract = addresses.find((a: any) => {
    return a.chain_name === "zeta_testnet" && a.type === "systemContract";
  })?.address;
  const provider = new hre.ethers.providers.JsonRpcProvider(ATHENS_EVM_RPC);
  const sc = SystemContract__factory.connect(systemContract, provider);
  const common = {
    chain_id: 7001,
    chain_name: "zeta_testnet",
    category: "omnichain",
  };
  try {
    addresses.push({ ...common, type: "uniswapv2FactoryAddress", address: await sc.uniswapv2FactoryAddress() });
    addresses.push({ ...common, type: "wZetaContractAddress", address: await sc.wZetaContractAddress() });
    addresses.push({ ...common, type: "uniswapv2Router02Address", address: await sc.uniswapv2Router02Address() });
    addresses.push({ ...common, type: "zetaConnectorZEVMAddress", address: await sc.zetaConnectorZEVMAddress() });
    addresses.push({ ...common, type: "fungibleModuleAddress", address: await sc.FUNGIBLE_MODULE_ADDRESS() });
  } catch (error) {
    console.error("Error fetching addresses from ZetaChain:", error);
  }
};

const fetchChainSpecificAddresses = async (chains: any, addresses: any) => {
  await Promise.all(
    chains.map(async (chain: any) => {
      return axios
        .get(`${ATHENS_TENDERMINT_RPC}/zeta-chain/observer/get_client_params_for_chain/${chain.chain_id}`)
        .then(({ data }) => {
          if (data.core_params.zeta_token_contract_address) {
            addresses.push({
              chain_id: chain.chain_id,
              chain_name: chain.chain_name,
              category: "messaging",
              type: "zetaToken",
              address: data.core_params.zeta_token_contract_address,
            });
          }
          if (data.core_params.connector_contract_address) {
            addresses.push({
              chain_id: chain.chain_id,
              chain_name: chain.chain_name,
              category: "messaging",
              type: "connector",
              address: data.core_params.connector_contract_address,
            });
          }
          if (data.core_params.erc20_custody_contract_address) {
            addresses.push({
              chain_id: chain.chain_id,
              category: "omnichain",
              chain_name: chain.chain_name,
              type: "erc20Custody",
              address: data.core_params.erc20_custody_contract_address,
            });
          }
        });
    })
  );
};

const fetchTSSUpdater = async (chains: any, addresses: any) => {
  await Promise.all(
    chains.map(async (chain: any) => {
      const erc20Custody = addresses.find((a: any) => {
        return a.chain_name === chain.chain_name && a.type === "erc20Custody";
      })?.address;
      if (erc20Custody) {
        const rpc = getEndpoints("evm", chain.chain_name)[0]?.url;
        const provider = new hre.ethers.providers.JsonRpcProvider(rpc);
        const custody = ERC20Custody__factory.connect(erc20Custody, provider);
        return custody.TSSAddressUpdater().then((address: string) => {
          addresses.push({
            chain_id: chain.chain_id,
            chain_name: chain.chain_name,
            category: "omnichain",
            type: "tssUpdater",
            address,
          });
        });
      }
    })
  );
};

const main = async (args: any, hre: HardhatRuntimeEnvironment) => {
  const addresses: any = [];

  const chains = await fetchChains();
  await fetchTssData(chains, addresses);
  await fetchSystemContract(addresses);
  await fetchForeignCoinsData(chains, addresses);
  await fetchAthensAddresses(addresses, hre);
  await fetchChainSpecificAddresses(chains, addresses);
  await fetchTSSUpdater(chains, addresses);

  console.log(JSON.stringify(addresses, null, 2));
};

task("addresses", "").setAction(main);
