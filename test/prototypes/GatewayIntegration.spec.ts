import { AddressZero } from "@ethersproject/constants";
import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { SystemContract, ZRC20 } from "@typechain-types";
import { expect } from "chai";
import { Contract } from "ethers";
import { parseEther } from "ethers/lib/utils";
import { ethers, upgrades } from "hardhat";

import { FUNGIBLE_MODULE_ADDRESS } from "../test.helpers";
const hre = require("hardhat");

describe("GatewayEVM GatewayZEVM integration", function () {
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

  beforeEach(async function () {
    [ownerEVM, ownerZEVM, destination, tssAddress, ...addrs] = await ethers.getSigners();
    // Prepare EVM
    const TestERC20 = await ethers.getContractFactory("TestERC20");
    const ReceiverEVM = await ethers.getContractFactory("Receiver");
    const GatewayEVM = await ethers.getContractFactory("GatewayEVM");
    const Custody = await ethers.getContractFactory("ERC20CustodyNew");

    // Deploy the contracts
    token = await TestERC20.deploy("Test Token", "TTK");
    receiverEVM = await ReceiverEVM.deploy();
    gatewayEVM = await upgrades.deployProxy(GatewayEVM, [tssAddress.address], {
      initializer: "initialize",
      kind: "uups",
    });
    custody = await Custody.deploy(gatewayEVM.address);

    gatewayEVM.setCustody(custody.address);

    // Mint initial supply to the owner
    await token.mint(ownerEVM.address, ethers.utils.parseEther("1000"));

    // Transfer some tokens to the custody contract
    await token.transfer(custody.address, ethers.utils.parseEther("500"));

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

    const ZRC20Factory = await ethers.getContractFactory("ZRC20");
    ZRC20Contract = (await ZRC20Factory.connect(fungibleModuleSigner).deploy(
      "TOKEN",
      "TKN",
      18,
      1,
      1,
      0,
      systemContract.address
    )) as ZRC20;

    await systemContract.setGasCoinZRC20(1, ZRC20Contract.address);
    await systemContract.setGasPrice(1, ZRC20Contract.address);

    await ZRC20Contract.connect(fungibleModuleSigner).deposit(ownerZEVM.address, parseEther("100"));

    const GatewayZEVM = await ethers.getContractFactory("GatewayZEVM");
    gatewayZEVM = await upgrades.deployProxy(GatewayZEVM, [], {
      initializer: "initialize",
      kind: "uups",
    });

    await ZRC20Contract.connect(ownerZEVM).approve(gatewayZEVM.address, parseEther("100"));

    // including abi of gatewayZEVM events, so hardhat can decode them automatically
    const senderArtifact = await hre.artifacts.readArtifact('Sender');
const gatewayZEVMArtifact = await hre.artifacts.readArtifact('GatewayZEVM');
const senderABI = [
  ...senderArtifact.abi, 
  ...gatewayZEVMArtifact.abi.filter((f: ethers.utils.Fragment) => f.type === 'event')
]

const SenderZEVM = new ethers.ContractFactory(senderABI, senderArtifact.bytecode, ownerZEVM);
    senderZEVM = await SenderZEVM.deploy(gatewayZEVM.address);

    await ZRC20Contract.connect(fungibleModuleSigner).deposit(senderZEVM.address, parseEther("100"));
  });

  it("should call Receiver contract on EVM from ZEVM account", async function () {
    const str = "Hello, Hardhat!";
    const num = 42;
    const flag = true;
    const value = ethers.utils.parseEther("1.0");

    // Encode the function call data and call on zevm
    const message = receiverEVM.interface.encodeFunctionData("receiveA", [str, num, flag]);
    const callTx = await gatewayZEVM.connect(ownerZEVM).call(ethers.utils.arrayify(addrs[0].address), message);
    await expect(callTx).to.emit(gatewayZEVM, "Call").withArgs(ownerZEVM.address, addrs[0].address, message);

    // Get message from events
    const callTxReceipt = await callTx.wait();
    const callEvent = callTxReceipt.events[0];
    const callMessage = callEvent.args[2];

    // Call execute on evm
    const executeTx = await gatewayEVM.execute(receiverEVM.address, callMessage, { value: value });

    // Listen for the event
    await expect(executeTx).to.emit(gatewayEVM, "Executed").withArgs(receiverEVM.address, value, callMessage);
    await expect(executeTx).to.emit(receiverEVM, "ReceivedA").withArgs(gatewayEVM.address, value, str, num, flag);
  });

  it("should withdraw and call Receiver contract on EVM from ZEVM account", async function () {
    const str = "Hello, Hardhat!";
    const num = 42;
    const flag = true;
    const value = ethers.utils.parseEther("1.0");

    // Encode the function call data and call on zevm
    const message = receiverEVM.interface.encodeFunctionData("receiveA", [str, num, flag]);
    const callTx = await gatewayZEVM
      .connect(ownerZEVM)
      .withdrawAndCall(receiverEVM.address, parseEther("1"), ZRC20Contract.address, message);

    await expect(callTx)
      .to.emit(gatewayZEVM, "Withdrawal")
      .withArgs(
        ethers.utils.getAddress(ownerZEVM.address),
        receiverEVM.address.toLowerCase(),
        parseEther("1"),
        0,
        await ZRC20Contract.PROTOCOL_FLAT_FEE(),
        message
      );

    // Get message from events
    const callTxReceipt = await callTx.wait();
    const callEvent = callTxReceipt.events.filter((e) => e.event == "Withdrawal")[0];
    const callMessage = callEvent.args[5];

    // Call execute on evm
    const executeTx = await gatewayEVM.execute(receiverEVM.address, callMessage, { value: value });

    // Listen for the event
    await expect(executeTx).to.emit(gatewayEVM, "Executed").withArgs(receiverEVM.address, value, callMessage);
    await expect(executeTx).to.emit(receiverEVM, "ReceivedA").withArgs(gatewayEVM.address, value, str, num, flag);

    const balanceOfAfterWithdrawal = await ZRC20Contract.balanceOf(ownerZEVM.address);
    expect(balanceOfAfterWithdrawal).to.equal(parseEther("99"));
  });

  it("should call Receiver contract on EVM from Sender contract on ZEVM", async function () {
    const str = "Hello, Hardhat!";
    const num = 42;
    const flag = true;
    const value = ethers.utils.parseEther("1.0");

    // Call sender function
    const callTx = await senderZEVM.connect(ownerZEVM).callReceiver(receiverEVM.address, str, num, flag);

    // Get message from events
    const callTxReceipt = await callTx.wait();
    const expectedMessage = receiverEVM.interface.encodeFunctionData("receiveA", [str, num, flag]);
    await expect(callTx)
      .to.emit(gatewayZEVM, "Call")
      .withArgs(senderZEVM.address, receiverEVM.address, expectedMessage);

    const callEvent = callTxReceipt.events[0];
    const callMessage = callEvent.args[2];

    // Call execute on evm
    const executeTx = await gatewayEVM.execute(receiverEVM.address, callMessage, { value: value });

    // Listen for the event
    await expect(executeTx).to.emit(gatewayEVM, "Executed").withArgs(receiverEVM.address, value, callMessage);
    await expect(executeTx).to.emit(receiverEVM, "ReceivedA").withArgs(gatewayEVM.address, value, str, num, flag);
  });

  it("should withdrawn and call Receiver contract on EVM from Sender contract on ZEVM", async function () {
    const str = "Hello, Hardhat!";
    const num = 42;
    const flag = true;
    const value = ethers.utils.parseEther("1.0");

    // Call sender function
    const callTx = await senderZEVM
      .connect(ownerZEVM)
      .withdrawAndCallReceiver(receiverEVM.address, parseEther("1"), ZRC20Contract.address, str, num, flag);

    // Get message from events
    const callTxReceipt = await callTx.wait();
    const expectedMessage = receiverEVM.interface.encodeFunctionData("receiveA", [str, num, flag]);

    await expect(callTx)
      .to.emit(gatewayZEVM, "Withdrawal")
      .withArgs(
        ethers.utils.getAddress(senderZEVM.address),
        receiverEVM.address.toLowerCase(),
        parseEther("1"),
        0,
        await ZRC20Contract.PROTOCOL_FLAT_FEE(),
        expectedMessage
      );

    const callEvent = callTxReceipt.events.filter((e) => e.event == "Withdrawal")[0];
    const callMessage = callEvent.args[5];

    // Call execute on evm
    const executeTx = await gatewayEVM.execute(receiverEVM.address, callMessage, { value: value });

    // Listen for the event
    await expect(executeTx).to.emit(gatewayEVM, "Executed").withArgs(receiverEVM.address, value, callMessage);
    await expect(executeTx).to.emit(receiverEVM, "ReceivedA").withArgs(gatewayEVM.address, value, str, num, flag);

    const balanceOfAfterWithdrawal = await ZRC20Contract.balanceOf(senderZEVM.address);
    expect(balanceOfAfterWithdrawal).to.equal(parseEther("99"));
  });
});
