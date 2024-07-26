import { ethers, NonceManager, Signer } from "ethers";
import * as GatewayEVM from "../../out/GatewayEVM.sol/GatewayEVM.json";
import * as Custody from "../../out/ERC20CustodyNew.sol/ERC20CustodyNew.json";
import * as ReceiverEVM from "../../out/ReceiverEVM.sol/ReceiverEVM.json";
import * as SenderZEVM from "../../out/SenderZEVM.sol/SenderZEVM.json";
import * as ERC1967Proxy from "../../out/ERC1967Proxy.sol/ERC1967Proxy.json";
import * as TestERC20 from "../../out/TestERC20.sol/TestERC20.json";
import * as SystemContract from "../../out/SystemContractMock.sol/SystemContractMock.json";
import * as GatewayZEVM from "../../out/GatewayZEVM.sol/GatewayZEVM.json";
import * as TestZContract from "../../out/TestZContract.sol/TestZContract.json";
import * as ZRC20New from "../../out/ZRC20New.sol/ZRC20New.json";
import * as ZetaConnectorNonNative from "../../out/ZetaConnectorNonNative.sol/ZetaConnectorNonNative.json";
import * as WETH9 from "../../out/WZETA.sol/WETH9.json";
/* eslint-disable  @typescript-eslint/no-explicit-any */

const FUNGIBLE_MODULE_ADDRESS = "0x735b14BB79463307AAcBED86DAf3322B1e6226aB";

let protocolContracts: any;
let testContracts: any;
let deployer: Signer;
const deployOpts = {
  gasPrice: 10000000000,
  gasLimit: 6721975,
};
const provider = new ethers.JsonRpcProvider("http://127.0.0.1:8545");
provider.pollingInterval = 100;

const deployProtocolContracts = async (deployer: Signer, fungibleModuleSigner: Signer) => {
  // Prepare EVM
  // Deploy protocol contracts (gateway and custody)
  const testERC20Factory = new ethers.ContractFactory(TestERC20.abi, TestERC20.bytecode, deployer);
  const testEVMZeta = await testERC20Factory.deploy("zeta", "ZETA", deployOpts);
  console.log("TestEVMZeta:", testEVMZeta.target);

  const gatewayEVMFactory = new ethers.ContractFactory(GatewayEVM.abi, GatewayEVM.bytecode, deployer);
  const gatewayEVMImpl = await gatewayEVMFactory.deploy(deployOpts);

  const gatewayEVMInterface = new ethers.Interface(GatewayEVM.abi);
  const gatewayEVMInitFragment = gatewayEVMInterface.getFunction("initialize");
  const gatewayEVMInitdata = gatewayEVMInterface.encodeFunctionData(gatewayEVMInitFragment as ethers.FunctionFragment, [
    await deployer.getAddress(),
    testEVMZeta.target,
  ]);

  const proxyEVMFactory = new ethers.ContractFactory(ERC1967Proxy.abi, ERC1967Proxy.bytecode, deployer);
  const proxyEVM = (await proxyEVMFactory.deploy(gatewayEVMImpl.target, gatewayEVMInitdata, deployOpts)) as any;

  const gatewayEVM = new ethers.Contract(proxyEVM.target, GatewayEVM.abi, deployer);
  console.log("GatewayEVM:", gatewayEVM.target);

  const zetaConnectorFactory = new ethers.ContractFactory(
    ZetaConnectorNonNative.abi,
    ZetaConnectorNonNative.bytecode,
    deployer
  );
  const zetaConnector = await zetaConnectorFactory.deploy(gatewayEVM.target, testEVMZeta.target, deployer, deployOpts);
  console.log("ZetaConnector:", zetaConnector.target);

  const custodyFactory = new ethers.ContractFactory(Custody.abi, Custody.bytecode, deployer);
  const custody = await custodyFactory.deploy(gatewayEVM.target, await deployer.getAddress(), deployOpts);
  console.log("Custody:", custody.target);

  await (gatewayEVM as any).connect(deployer).setCustody(custody.target, deployOpts);
  await (gatewayEVM as any).connect(deployer).setConnector(zetaConnector.target, deployOpts);

  // Prepare ZEVM
  // Deploy protocol contracts (gateway and system)
  const weth9Factory = new ethers.ContractFactory(WETH9.abi, WETH9.bytecode, deployer);
  const wzeta = await weth9Factory.deploy(deployOpts);
  console.log("WZeta:", wzeta.target);

  const systemContractFactory = new ethers.ContractFactory(SystemContract.abi, SystemContract.bytecode, deployer);
  const systemContract = await systemContractFactory.deploy(
    ethers.ZeroAddress,
    ethers.ZeroAddress,
    ethers.ZeroAddress,
    deployOpts
  );

  console.log("SystemContract:", systemContract.target);

  const gatewayZEVMFactory = new ethers.ContractFactory(GatewayZEVM.abi, GatewayZEVM.bytecode, deployer);
  const gatewayZEVMImpl = await gatewayZEVMFactory.deploy(deployOpts);

  const gatewayZEVMInterface = new ethers.Interface(GatewayZEVM.abi);
  const gatewayZEVMInitFragment = gatewayZEVMInterface.getFunction("initialize");
  const gatewayZEVMInitData = gatewayEVMInterface.encodeFunctionData(
    gatewayZEVMInitFragment as ethers.FunctionFragment,
    [wzeta.target]
  );

  const proxyZEVMFactory = new ethers.ContractFactory(ERC1967Proxy.abi, ERC1967Proxy.bytecode, deployer);
  const proxyZEVM = (await proxyZEVMFactory.deploy(gatewayZEVMImpl.target, gatewayZEVMInitData, deployOpts)) as any;

  const gatewayZEVM = new ethers.Contract(proxyZEVM.target, GatewayZEVM.abi, deployer);
  console.log("GatewayZEVM:", gatewayZEVM.target);

  await (wzeta as any).connect(fungibleModuleSigner).deposit({ ...deployOpts, value: ethers.parseEther("10") });
  await (wzeta as any).connect(fungibleModuleSigner).approve(gatewayZEVM.target, ethers.parseEther("10"), deployOpts);
  await (wzeta as any).connect(deployer).deposit({ ...deployOpts, value: ethers.parseEther("10") });
  await (wzeta as any).connect(deployer).approve(gatewayZEVM.target, ethers.parseEther("10"), deployOpts);

  return {
    custody,
    zetaConnector,
    gatewayEVM,
    gatewayZEVM,
    systemContract,
    testEVMZeta,
    wzeta,
  };
};

const deployTestContracts = async (protocolContracts: any, deployer: Signer, fungibleModuleSigner: Signer) => {
  // Prepare EVM
  // Deploy test contracts (erc20, receiver) and mint funds to test accounts
  const deployOpts = {
    gasPrice: 10000000000,
    gasLimit: 6721975,
  };
  const testERC20Factory = new ethers.ContractFactory(TestERC20.abi, TestERC20.bytecode, deployer);
  const testEVMZeta = await testERC20Factory.deploy("zeta", "ZETA", deployOpts);
  console.log("TestEVMZeta:", testEVMZeta.target);

  const token = await testERC20Factory.deploy("Test Token", "TTK", deployOpts);
  console.log("TestERC20:", token.target);

  const receiverEVMFactory = new ethers.ContractFactory(ReceiverEVM.abi, ReceiverEVM.bytecode, deployer);
  const receiverEVM = await receiverEVMFactory.deploy(deployOpts);
  console.log("ReceiverEVM:", receiverEVM.target);

  await (token as any).connect(deployer).mint(await deployer.getAddress(), ethers.parseEther("1000"), deployOpts);
  await (token as any)
    .connect(deployer)
    .transfer(protocolContracts.custody.target, ethers.parseEther("500"), deployOpts);

  // Prepare ZEVM
  // Deploy test contracts (test zContract, zrc20, sender) and mint funds to test accounts
  const testZContractFactory = new ethers.ContractFactory(TestZContract.abi, TestZContract.bytecode, deployer);
  const testZContract = await testZContractFactory.deploy(deployOpts);
  console.log("TestZContract:", testZContract.target);

  const zrc20Factory = new ethers.ContractFactory(ZRC20New.abi, ZRC20New.bytecode, deployer);
  const zrc20 = await zrc20Factory
    .connect(fungibleModuleSigner)
    .deploy(
      "TOKEN",
      "TKN",
      18,
      1,
      1,
      0,
      protocolContracts.systemContract.target,
      protocolContracts.gatewayZEVM.target,
      deployOpts
    );
  console.log("ZRC20:", zrc20.target);

  await protocolContracts.systemContract.setGasCoinZRC20(1, zrc20.target, deployOpts);
  await protocolContracts.systemContract.setGasPrice(1, zrc20.target, deployOpts);
  await (zrc20 as any)
    .connect(fungibleModuleSigner)
    .deposit(await deployer.getAddress(), ethers.parseEther("100"), deployOpts);
  await (zrc20 as any)
    .connect(deployer)
    .approve(protocolContracts.gatewayZEVM.target, ethers.parseEther("100"), deployOpts);

  // Include abi of gatewayZEVM events, so hardhat can decode them automatically
  const senderABI = [...SenderZEVM.abi, ...GatewayZEVM.abi.filter((f) => f.type === "event")];

  const senderZEVMFactory = new ethers.ContractFactory(senderABI, SenderZEVM.bytecode, deployer);
  const senderZEVM = await senderZEVMFactory.deploy(protocolContracts.gatewayZEVM.target, deployOpts);
  await (zrc20 as any).connect(fungibleModuleSigner).deposit(senderZEVM.target, ethers.parseEther("100"), deployOpts);

  return {
    zrc20,
    receiverEVM,
    senderZEVM,
    testZContract,
  };
};

const startWorker = async () => {
  // anvil test mnemonic
  const mnemonic = "test test test test test test test test test test test junk";

  // impersonate and fund fungible module account
  await provider.send("anvil_impersonateAccount", [FUNGIBLE_MODULE_ADDRESS]);
  await provider.send("anvil_setBalance", [FUNGIBLE_MODULE_ADDRESS, ethers.parseEther("100000").toString()]);
  const fungibleModuleSigner = await provider.getSigner(FUNGIBLE_MODULE_ADDRESS);

  deployer = new NonceManager(ethers.Wallet.fromPhrase(mnemonic, provider));
  deployer.connect(provider);

  protocolContracts = await deployProtocolContracts(deployer, fungibleModuleSigner);
  testContracts = await deployTestContracts(protocolContracts, deployer, fungibleModuleSigner);

  // Listen to contracts events
  // event Call(address indexed sender, bytes receiver, bytes message);
  protocolContracts.gatewayZEVM.on("Call", async (...args: Array<any>) => {
    console.log("Worker: Call event on GatewayZEVM.");
    console.log("Worker: Calling ReceiverEVM through GatewayEVM...");
    try {
      (deployer as NonceManager).reset();

      const receiver = args[1];
      const message = args[2];

      const executeTx = await protocolContracts.gatewayEVM.connect(deployer).execute(receiver, message, deployOpts);
      await executeTx.wait();
    } catch (e) {
      console.error("failed:", e);
    }
  });

  // event Withdrawal(address indexed from, address zrc20, bytes to, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message);
  protocolContracts.gatewayZEVM.on("Withdrawal", async (...args: Array<any>) => {
    console.log("Worker: Withdrawal event on GatewayZEVM.");
    console.log("Worker: Calling ReceiverEVM through GatewayEVM...");
    try {
      const receiver = args[2];
      const message = args[6];
      (deployer as NonceManager).reset();

      if (message != "0x") {
        const executeTx = await protocolContracts.gatewayEVM.connect(deployer).execute(receiver, message, deployOpts);
        await executeTx.wait();
      }
    } catch (e) {
      console.error("failed:", e);
    }
  });

  testContracts.receiverEVM.on("ReceivedPayable", () => {
    console.log("ReceiverEVM: receivePayable called!");
  });

  // event Call(address indexed sender, address indexed receiver, bytes payload);
  protocolContracts.gatewayEVM.on("Call", async (...args: Array<any>) => {
    console.log("Worker: Call event on GatewayEVM.");
    console.log("Worker: Calling TestZContract through GatewayZEVM...");
    try {
      const zContract = args[1];
      const payload = args[2];
      (deployer as NonceManager).reset();
      // Encode the parameters
      const origin = protocolContracts.gatewayZEVM.target;
      const sender = await fungibleModuleSigner.getAddress();
      const chainID = 1;

      // Call the execute function
      const executeTx = await protocolContracts.gatewayZEVM.connect(fungibleModuleSigner).execute(
        {
          origin,
          sender,
          chainID,
        },
        testContracts.zrc20.target,
        0,
        zContract,
        payload,
        deployOpts
      );
      await executeTx.wait();
    } catch (e) {
      console.error("failed:", e);
    }
  });

  // event Deposit(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload);
  protocolContracts.gatewayEVM.on("Deposit", async (...args: Array<any>) => {
    console.log("Worker: Deposit event on GatewayEVM.");
    console.log("Worker: Calling TestZContract through GatewayZEVM...");
    try {
      const receiver = args[1];
      const asset = args[3];
      const payload = args[4];
      if (payload != "0x") {
        const executeTx = await (protocolContracts.gatewayZEVM as any)
          .connect(fungibleModuleSigner)
          .execute(
            [protocolContracts.gatewayZEVM.target, await fungibleModuleSigner.getAddress(), 1],
            asset,
            0,
            receiver,
            payload,
            deployOpts
          );
        await executeTx.wait();
      }
    } catch (e) {
      console.error("failed:", e);
    }
  });

  testContracts.testZContract.on("ContextData", async () => {
    console.log("TestZContract: onCrosschainCall called!");
  });

  process.stdin.resume();
};

startWorker()
  .then(() => {
    console.log("Setup complete, monitoring events. Press CTRL+C to exit.");
  })
  .catch((error) => {
    console.error("Failed to deploy contracts or set up listeners:", error);
    process.exit(1);
  });