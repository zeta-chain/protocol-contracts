import hre from 'hardhat';
import * as dotenv from 'dotenv';
import * as fs from 'fs';
import * as path from 'path';
import { ethers } from 'ethers';
import SafeApiKit from '@safe-global/api-kit';
import Safe from '@safe-global/protocol-kit';
import {
  MetaTransactionData,
  OperationType
} from '@safe-global/types-kit';

dotenv.config();

const CHAIN_ID_TO_NETWORK: Record<string, string> = {
  "1": "eth_mainnet",
  "56": "bsc_mainnet",
  "137": "polygon_mainnet",
  "8453": "base_mainnet",
  "42161": "arbitrum_mainnet",
  "43114": "avalanche_mainnet",
  "7000": "zeta_mainnet",
  "7001": "zeta_testnet",
  "11155111": "sepolia_testnet",
  "97": "bsc_testnet",
  "80002": "amoy_testnet",
  "84532": "base_sepolia",
  "421614": "arbitrum_sepolia",
  "43113": "avalanche_testnet"
};

function loadAddresses() {
  try {
    const addressesPath = path.join(process.cwd(), 'data/checksum/testnet.json');
    const rawData = fs.readFileSync(addressesPath, 'utf8');
    return JSON.parse(rawData);
  } catch (error) {
    console.error('❌ Error loading addresses from data/checksum/testnet.json:', error);
    process.exit(1);
  }
}

function createUpgradeCalldata(newImplementationAddress: string, upgradeData: string): string {
  const iface = new ethers.Interface([
    'function upgradeToAndCall(address newImplementation, bytes memory data) external payable',
  ]);
  return iface.encodeFunctionData('upgradeToAndCall', [
    newImplementationAddress,
    upgradeData,
  ]);
}

async function createUpgradeProposal(
  chainId: number,
  networkName: string,
  addresses: Record<string, string>,
  upgradeData: string
) {
  const safeAddress = addresses.Admin;
  const gatewayProxy = addresses.GatewayEVM;
  const newImplementationAddress = addresses.GatewayEVMImplementation;
  const networkConfig = hre.config.networks[networkName] as any;

  console.log(`\n🔄 Creating proposal for GatewayEVM upgrade on ${networkName} (Chain ID: ${chainId})`);
  console.log(`  Safe: ${safeAddress}`);
  console.log(`  Gateway: ${gatewayProxy}`);
  console.log(`  New Implementation: ${newImplementationAddress}`);

  const apiKit = new SafeApiKit({
    chainId: BigInt(chainId),
    apiKey: process.env.SAFE_API_KEY,
  });

  const protocolKitOwner1 = await Safe.init({
    provider: networkConfig.url,
    signer: process.env.PROPOSER_PRIVATE_KEY,
    safeAddress: safeAddress
  });

  const upgradeCalldata = createUpgradeCalldata(newImplementationAddress, upgradeData);

  const safeTransactionData: MetaTransactionData = {
    to: gatewayProxy,
    value: '0',
    data: upgradeCalldata,
    operation: OperationType.Call
  }

  const safeTransaction = await protocolKitOwner1.createTransaction({
    transactions: [safeTransactionData]
  })

  const safeTxHash = await protocolKitOwner1.getTransactionHash(safeTransaction)
  const signature = await protocolKitOwner1.signHash(safeTxHash)

  await apiKit.proposeTransaction({
    safeAddress: safeAddress,
    safeTransactionData: safeTransaction.data,
    safeTxHash,
    senderAddress: process.env.PROPOSER_ADDRESS!,
    senderSignature: signature.data
  })

  console.log('  ✅ Proposal submitted to Safe service');
  console.log(`  Safe Transaction Hash: ${safeTxHash}`);
}

async function main() {
  const addresses = loadAddresses();
  const upgradeData = process.env.UPGRADE_DATA || '0x';

  console.log('🚀 Starting upgrade proposals for all networks');
  console.log('='.repeat(60));

  for (const [chainId, contracts] of Object.entries(addresses)) {
    const chainIdNum = parseInt(chainId);
    const networkName = CHAIN_ID_TO_NETWORK[chainId];
    const contractsTyped = contracts as Record<string, string>;
    if (!contractsTyped.GatewayEVM) {
      console.log(`\n⚠️ No GatewayEVM found for ${networkName} (Chain ID: ${chainId})`);
      continue;
    }

    try {
      const newImplementation = contractsTyped.GatewayEVMImplementation;
      if (!newImplementation) {
        console.log(`\n⚠️ No GatewayEVMImplementation found for ${networkName} (Chain ID: ${chainId})`);
        continue;
      }

      await createUpgradeProposal(chainIdNum, networkName, contractsTyped, upgradeData);
    } catch (error) {
      console.error(`\n❌ Error creating proposal for ${networkName} (Chain ID: ${chainId}):`, error);
    }
  }
}

if (require.main === module) {
  main().catch(console.error);
}