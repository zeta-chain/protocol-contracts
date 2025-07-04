import hre from "hardhat";
import path from "path";
import fs from "fs";

// ERC1967 implementation slot: keccak256("eip1967.proxy.implementation") - 1
const ERC1967_IMPLEMENTATION_SLOT = "0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc";

const CHAIN_ID_TO_NETWORK: Record<string, string> = {
  "1": "ethereum",
  "56": "bsc",
  "137": "polygon",
  "8453": "base",
  "42161": "arbitrum",
  "43114": "avalanche",
  "7000": "zetachain"
};

function loadAddresses() {
  try {
    const addressesPath = path.join(process.cwd(), `/data/checksum/mainnet.json`);
    const rawData = fs.readFileSync(addressesPath, 'utf8');
    return JSON.parse(rawData);
  } catch (error) {
    console.error(`❌ Error loading addresses from JSON`);
    process.exit(1);
  }
}

async function checkProxy(
  contractName: string,
  proxyAddress: string,
  localImplementation: string,
  provider: any
) {
  try {
    // Read the implementation slot form the proxy
    const implementationSlot = await provider.getStorage(proxyAddress, ERC1967_IMPLEMENTATION_SLOT);
    // Convert to address
    const remoteImplementation = implementationSlot !== '0x0000000000000000000000000000000000000000000000000000000000000000'
      ? hre.ethers.getAddress('0x' + implementationSlot.slice(-40))
      : '0x0000000000000000000000000000000000000000';

    const matches = remoteImplementation.toLowerCase() === localImplementation.toLowerCase();
    console.log(`    Local Implementation:  ${localImplementation}`);
    console.log(`    Remote Implementation: ${remoteImplementation}`);
    console.log(`    Synced: ${matches ? '✅' : '❌'}`);
    if (!matches) {
      console.log(`      Local: ${localImplementation}`);
      console.log(`      Remote:   ${remoteImplementation}`);
    }

    return matches;
  } catch (error) {
    console.error(`    Proxy Implementation: ❌ (Error)`);
    return false;
  }
}

async function compareBytecode(
  contractName: string,
  contractAddress: string,
  provider: any
) {
  try {
    const remoteBytecode = await provider.getCode(contractAddress);
    const artifact = await hre.artifacts.readArtifact(contractName);
    const localBytecode = artifact.deployedBytecode;
    const matches = remoteBytecode === localBytecode;

    console.log(`    Bytecode Match: ${matches ? '✅' : '❌'}`);
    if (!matches) {
      const lengthMatch = remoteBytecode.length === localBytecode.length;
      console.log(`      Length: ${lengthMatch ? 'Same' : 'Different'} (Remote: ${remoteBytecode.length}, Local: ${localBytecode.length})`);
    }

    return matches;
  } catch (error) {
    console.error(`    Bytecode Match: ❌ (Error)`);
    return false;
  }
}

async function checkGatewayEVM(
  chainId: string,
  contractName: string,
  contractAddress: string,
  allContracts: Record<string, string>,
  provider: any
) {
  console.log(`\n  🏗️  GatewayEVM (Proxy):`);

  if (allContracts.GatewayEVMImplementation) {
    await checkProxy(contractName, contractAddress, allContracts.GatewayEVMImplementation, provider);
    await compareBytecode("GatewayEVM", allContracts.GatewayEVMImplementation, provider);
  } else {
    console.log(`    ⚠️  No implementation found`);
  }
}

async function checkERC20Custody(
  chainId: string,
  contractName: string,
  contractAddress: string,
  allContracts: Record<string, string>,
  provider: any
) {
  console.log(`\n  🏗️  ERC20Custody (Proxy):`);

  if (allContracts.ERC20CustodyImplementation) {
    await checkProxy("ERC20Custody", contractAddress, allContracts.ERC20CustodyImplementation, provider);
    await compareBytecode(contractName, allContracts.ERC20CustodyImplementation, provider);
  } else {
    console.log(`    ⚠️  No implementation found`);
  }
}

async function checkSystemContract(
  chainId: string,
  contractName: string,
  contractAddress: string,
  allContracts: Record<string, string>,
  provider: any
) {
  console.log(`\n  🏗️  SystemContract (Non-upgradeable):`);
  await compareBytecode(contractName, contractAddress, provider);
}

async function checkGatewayZEVM(
  chainId: string,
  contractName: string,
  contractAddress: string,
  allContracts: Record<string, string>,
  provider: any
) {
  console.log(`\n  🏗️  GatewayZEVM (Proxy):`);

  if (allContracts.GatewayZEVMImplementation) {
    await checkProxy("GatewayZEVM", contractAddress, allContracts.GatewayZEVMImplementation, provider);
    await compareBytecode(contractName, allContracts.GatewayZEVMImplementation, provider);
  } else {
    console.log(`    ⚠️  No implementation found`);
  }
}

async function checksumNetworks() {
  const addresses = loadAddresses();

  console.log("🔍 Contract checksum verification:");
  console.log("=".repeat(60));

  for (const [chainId, contracts] of Object.entries(addresses)) {
    const contractsTyped = contracts as Record<string, string>;
    const networkName = CHAIN_ID_TO_NETWORK[chainId];

    console.log(`\nChain ID: ${chainId} (${networkName})`);
    console.log("-".repeat(50));

    if (!networkName) {
      console.log(`  ⚠️  No network configuration found for chain ${chainId}`);
      continue;
    }

    try {
      const networkConfig = hre.config.networks[networkName] as any;
      const provider = new hre.ethers.JsonRpcProvider(networkConfig.url);

      for (const [contractName, contractAddress] of Object.entries(contractsTyped)) {
        switch (contractName) {
          case "GatewayEVM":
            await checkGatewayEVM(chainId, contractName, contractAddress, contractsTyped, provider);
            break;
          case "ERC20Custody":
            await checkERC20Custody(chainId, contractName, contractAddress, contractsTyped, provider);
            break;
          case "SystemContract":
            await checkSystemContract(chainId, contractName, contractAddress, contractsTyped, provider);
            break;
          case "GatewayZEVM":
            await checkGatewayZEVM(chainId, contractName, contractAddress, contractsTyped, provider);
            break;
        }
      }
    } catch (error) {
      console.error(`  ❌ Error connecting to ${networkName}:`);
    }
  }
}

async function main() {
  await checksumNetworks();
}

if (require.main === module) {
  main().catch(console.error);
}