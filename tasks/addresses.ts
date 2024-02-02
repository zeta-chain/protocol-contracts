import { getEndpoints } from "@zetachain/networks";
import axios, { AxiosResponse } from "axios";
import { task } from "hardhat/config";
import { HardhatRuntimeEnvironment } from "hardhat/types";

import { ERC20Custody__factory } from "../typechain-types/factories/contracts/evm/ERC20Custody__factory";
import { SystemContract__factory } from "../typechain-types/factories/contracts/zevm/SystemContract.sol/SystemContract__factory";

type Network = "zeta_mainnet" | "zeta_testnet";

const api = {
  zeta_mainnet: {
    evm: "https://zetachain-evm.blockpi.network/v1/rpc/public",
    rpc: "https://zetachain.blockpi.network/lcd/v1/public",
  },
  zeta_testnet: {
    evm: "https://zetachain-athens-evm.blockpi.network/v1/rpc/public",
    rpc: "https://zetachain-athens.blockpi.network/lcd/v1/public",
  },
};

const fetchChains = async (network: Network) => {
  const URL = `${api[network].rpc}/zeta-chain/observer/supportedChains`;
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

const fetchTssData = async (chains: any, addresses: any, network: Network) => {
  const URL = `${api[network].rpc}/zeta-chain/observer/get_tss_address/18332`;
  try {
    const tssResponse: AxiosResponse<any> = await axios.get(URL);

    if (tssResponse.status === 200) {
      chains.forEach((chain: any) => {
        const { btc, eth } = tssResponse.data;
        if (chain.chain_name === "zeta_testnet") return;
        addresses.push({
          address: chain.chain_name === "btc_testnet" ? btc : eth,
          category: "omnichain",
          chain_id: chain.chain_id,
          chain_name: chain.chain_name,
          type: "tss",
        });
      });
    } else {
      console.error("Error fetching TSS data:", tssResponse.status, tssResponse.statusText);
    }
  } catch (error) {
    console.error("Error fetching TSS data:", error);
  }
};

const fetchSystemContract = async (addresses: any, network: Network) => {
  const chain_id = network === "zeta_mainnet" ? 7000 : 7001;
  const URL = `${api[network].rpc}/zeta-chain/fungible/system_contract`;
  try {
    const systemContractResponse: AxiosResponse<any> = await axios.get(URL);

    if (systemContractResponse.status === 200) {
      addresses.push({
        address: systemContractResponse.data.SystemContract.system_contract,
        category: "omnichain",
        chain_id,
        chain_name: "zeta_testnet",
        type: "systemContract",
      });
      addresses.push({
        address: systemContractResponse.data.SystemContract.connector_zevm,
        category: "messaging",
        chain_id,
        chain_name: "zeta_testnet",
        type: "connectorZEVM",
      });
    } else {
      console.error("Error fetching system contract:", systemContractResponse.statusText);
    }
  } catch (error) {
    console.error("Error fetching system contract:", error);
  }
};

const fetchForeignCoinsData = async (chains: any, addresses: any, network: Network) => {
  const chain_id = network === "zeta_mainnet" ? 7000 : 7001;
  const URL = `${api[network].rpc}/zeta-chain/fungible/foreign_coins`;
  try {
    const foreignCoinsResponse: AxiosResponse<any> = await axios.get(URL);
    if (foreignCoinsResponse.status === 200) {
      foreignCoinsResponse.data.foreignCoins.forEach((token: any) => {
        addresses.push({
          address: token.asset,
          category: "omnichain",
          chain_id: token.foreign_chain_id,
          chain_name: chains.find((c: any) => c.chain_id === token.foreign_chain_id)?.chain_name,
          coin_type: token.coin_type,
          decimals: token.decimals,
          symbol: token.symbol,
          zrc20: token.zrc20_contract_address, // TODO: dynamically fetch from contract (to verify)
        });
        addresses.push({
          address: token.zrc20_contract_address,
          asset: token.asset,
          category: "omnichain",
          chain_id,
          chain_name: "zeta_testnet",
          coin_type: "ZRC20",
          decimals: 18,
          symbol: token.name, // TODO: dynamically fetch from contract
        });
      });
    } else {
      console.error("Error fetching foreign coins data:", foreignCoinsResponse.status, foreignCoinsResponse.statusText);
    }
  } catch (error) {
    console.error("Error fetching foreign coins data:", error);
  }
};

const fetchAthensAddresses = async (addresses: any, hre: HardhatRuntimeEnvironment, network: Network) => {
  const chain_id = network === "zeta_mainnet" ? 7000 : 7001;
  const systemContract = addresses.find((a: any) => {
    return a.chain_name === "zeta_testnet" && a.type === "systemContract";
  })?.address;
  const provider = new hre.ethers.providers.JsonRpcProvider(api[network].evm);
  const sc = SystemContract__factory.connect(systemContract, provider);
  const common = {
    category: "omnichain",
    chain_id,
    chain_name: "zeta_testnet",
  };
  try {
    addresses.push({ ...common, address: await sc.uniswapv2FactoryAddress(), type: "uniswapv2Factory" });
    addresses.push({ ...common, address: await sc.wZetaContractAddress(), type: "wZetaContract" });
    addresses.push({ ...common, address: await sc.uniswapv2Router02Address(), type: "uniswapv2Router02" });
    addresses.push({ ...common, address: await sc.zetaConnectorZEVMAddress(), type: "zetaConnectorZEVM" });
    addresses.push({ ...common, address: await sc.FUNGIBLE_MODULE_ADDRESS(), type: "fungibleModule" });
  } catch (error) {
    console.error("Error fetching addresses from ZetaChain:", error);
  }
};

const fetchChainSpecificAddresses = async (chains: any, addresses: any, network: Network) => {
  await Promise.all(
    chains.map(async (chain: any) => {
      return axios
        .get(`${api[network].rpc}/zeta-chain/observer/get_chain_params_for_chain/${chain.chain_id}`)
        .then(({ data }) => {
          const zetaToken = data.chain_params.zeta_token_contract_address;
          if (zetaToken && zetaToken != "0x0000000000000000000000000000000000000000") {
            addresses.push({
              address: zetaToken,
              category: "messaging",
              chain_id: chain.chain_id,
              chain_name: chain.chain_name,
              type: "zetaToken",
            });
          }
          const connector = data.chain_params.connector_contract_address;
          if (connector && connector != "0x0000000000000000000000000000000000000000") {
            addresses.push({
              address: connector,
              category: "messaging",
              chain_id: chain.chain_id,
              chain_name: chain.chain_name,
              type: "connector",
            });
          }
          const erc20Custody = data.chain_params.erc20_custody_contract_address;
          if (erc20Custody && erc20Custody != "0x0000000000000000000000000000000000000000") {
            addresses.push({
              address: data.chain_params.erc20_custody_contract_address,
              category: "omnichain",
              chain_id: chain.chain_id,
              chain_name: chain.chain_name,
              type: "erc20Custody",
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
        if (["18332", "8332"].includes(chain.chain_id)) return;
        const rpc = getEndpoints("evm", chain.chain_name)[0]?.url;
        const provider = new hre.ethers.providers.JsonRpcProvider(rpc);
        const custody = ERC20Custody__factory.connect(erc20Custody, provider);
        return custody.TSSAddressUpdater().then((address: string) => {
          addresses.push({
            address,
            category: "omnichain",
            chain_id: chain.chain_id,
            chain_name: chain.chain_name,
            type: "tssUpdater",
          });
        });
      }
    })
  );
};

const main = async (args: any, hre: HardhatRuntimeEnvironment) => {
  const n = hre.network.name;
  if (n !== "zeta_testnet" && n !== "zeta_mainnet") {
    throw new Error(`Unsupported network: ${n}. Must be 'zeta_testnet' or 'zeta_mainnet'.`);
  }

  const network: Network = n;
  const addresses: any = [];

  const chains = await fetchChains(network);
  await fetchTssData(chains, addresses, network);
  await fetchSystemContract(addresses, network);
  await fetchForeignCoinsData(chains, addresses, network);
  await fetchAthensAddresses(addresses, hre, network);
  await fetchChainSpecificAddresses(chains, addresses, network);
  await fetchTSSUpdater(chains, addresses);

  console.log(JSON.stringify(addresses, null, 2));
};

task("addresses", "").setAction(main);
