import { AddressZero } from "@ethersproject/constants";
import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { SystemContract, ZRC20 } from "@typechain-types";
import { Contract } from "ethers";
import { parseEther } from "ethers/lib/utils";
import { ethers, upgrades } from "hardhat";

const hre = require("hardhat");

export const FUNGIBLE_MODULE_ADDRESS = "0x735b14BB79463307AAcBED86DAf3322B1e6226aB";

export const startLocalnet = async () => {
  console.log("deploying contracts");
  // EVM
  let receiverEVM: Contract;
  let gatewayEVM: Contract;
  let token: Contract;
  let custody: Contract;
  let ownerEVM: any, destination: any, tssAddress: any;

  // ZEVM
  let senderZEVM: Contract;
  let ZRC20Contract: ZRC20;
  let systemContract: SystemContract;
  let gatewayZEVM: Contract;
  let ownerZEVM: SignerWithAddress;
  let addrs: SignerWithAddress[];

  [ownerEVM, ownerZEVM, destination, tssAddress, ...addrs] = await ethers.getSigners();
  // Prepare EVM
  const TestERC20 = await ethers.getContractFactory("TestERC20");
  const ReceiverEVM = await ethers.getContractFactory("ReceiverEVM");
  const GatewayEVM = await ethers.getContractFactory("GatewayEVM");
  const Custody = await ethers.getContractFactory("ERC20CustodyNew");

  // Deploy the contracts
  token = await TestERC20.deploy("Test Token", "TTK");
  console.log("EVM: ERC20(TTK) token deployed to:", token.address);

  receiverEVM = await ReceiverEVM.deploy();
  console.log("EVM: Receiver deployed to:", receiverEVM.address);

  gatewayEVM = await upgrades.deployProxy(GatewayEVM, [tssAddress.address], {
    initializer: "initialize",
    kind: "uups",
  });
  console.log("EVM: GatewayEVM deployed to:", gatewayEVM.address);

  custody = await Custody.deploy(gatewayEVM.address);
  gatewayEVM.setCustody(custody.address);
  console.log("EVM: ERC20Custody deployed to:", custody.address);

  // Mint initial supply to the owner
  await token.mint(ownerEVM.address, ethers.utils.parseEther("1000"));
  console.log("EVM: 1000TTK minted to:", ownerEVM.address);

  // Transfer some tokens to the custody contract
  await token.transfer(custody.address, ethers.utils.parseEther("500"));
  console.log("EVM: 500TTK transfered to custody from:", ownerEVM.address);

  // Prepare ZEVM
  // Impersonate the fungible module account
  await hre.network.provider.request({
    method: "hardhat_impersonateAccount",
    params: [FUNGIBLE_MODULE_ADDRESS],
  });

  // Get a signer for the fungible module account
  const fungibleModuleSigner = await ethers.getSigner(FUNGIBLE_MODULE_ADDRESS);
  hre.network.provider.send("hardhat_setBalance", [FUNGIBLE_MODULE_ADDRESS, parseEther("1000000").toHexString()]);

  const SystemContractFactory = await ethers.getContractFactory("SystemContractMock");
  systemContract = (await SystemContractFactory.deploy(AddressZero, AddressZero, AddressZero)) as SystemContract;
  console.log("ZEVM: SystemContract deployed to:", systemContract.address);

  const GatewayZEVM = await ethers.getContractFactory("GatewayZEVM");
  gatewayZEVM = await upgrades.deployProxy(GatewayZEVM, [], {
    initializer: "initialize",
    kind: "uups",
  });
  console.log("ZEVM: GatewayZEVM deployed to:", gatewayZEVM.address);

  const ZRC20Factory = await ethers.getContractFactory("ZRC20New");
  ZRC20Contract = (await ZRC20Factory.connect(fungibleModuleSigner).deploy(
    "TOKEN",
    "TKN",
    18,
    1,
    1,
    0,
    systemContract.address,
    gatewayZEVM.address
  )) as ZRC20;
  console.log("ZEVM: ZRC20(TKN) contract deployed to:", ZRC20Contract.address);

  await systemContract.setGasCoinZRC20(1, ZRC20Contract.address);
  await systemContract.setGasPrice(1, ZRC20Contract.address);

  await ZRC20Contract.connect(fungibleModuleSigner).deposit(ownerZEVM.address, parseEther("100"));
  console.log("ZEVM: Fungible module deposited 100TKN to:", ownerZEVM.address);

  await ZRC20Contract.connect(ownerZEVM).approve(gatewayZEVM.address, parseEther("100"));
  console.log(`ZEVM: ${ownerZEVM.address} approved GatewayZEVM ${gatewayZEVM.address} 100TKN`);

  // including abi of gatewayZEVM events, so hardhat can decode them automatically
  const senderArtifact = await hre.artifacts.readArtifact("SenderZEVM");
  const gatewayZEVMArtifact = await hre.artifacts.readArtifact("GatewayZEVM");
  const senderABI = [
    ...senderArtifact.abi,
    ...gatewayZEVMArtifact.abi.filter((f: ethers.utils.Fragment) => f.type === "event"),
  ];

  const SenderZEVM = new ethers.ContractFactory(senderABI, senderArtifact.bytecode, ownerZEVM);
  senderZEVM = await SenderZEVM.deploy(gatewayZEVM.address);
  console.log("ZEVM: Sender contract deployed to:", senderZEVM.address);

  await ZRC20Contract.connect(fungibleModuleSigner).deposit(senderZEVM.address, parseEther("100"));
  console.log("ZEVM: Fungible module deposited 100TKN to sender:", senderZEVM.address);

  gatewayZEVM.on("Call", async (...args: Array<any>) => {
    console.log("Worker: Call event on GatewayZEVM.");
    console.log("Worker: Calling ReceiverEVM through GatewayEVM...");
    const executeTx = await gatewayEVM.execute(receiverEVM.address, args[3].args.message, { value: 0 });
    await executeTx.wait();
  });

  receiverEVM.on("ReceivedA", () => {
    console.log("ReceiverEVM: receiveA called!");
  });

  process.stdin.resume();
};

startLocalnet()
  .then(() => {
    console.log("Setup complete, monitoring events. Press CTRL+C to exit.");
  })
  .catch((error) => {
    console.error("Failed to deploy contracts or set up listeners:", error);
    process.exit(1);
  });
