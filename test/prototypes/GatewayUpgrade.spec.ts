import { expect } from "chai";
import { Contract } from "ethers";
import { ethers, upgrades } from "hardhat";

describe("Gateway upgrade", function () {
  let receiver: Contract;
  let gateway: Contract;
  let token: Contract;
  let custody: Contract;
  let owner: any, destination: any;

  beforeEach(async function () {
    const TestERC20 = await ethers.getContractFactory("TestERC20");
    const Receiver = await ethers.getContractFactory("Receiver");
    const Gateway = await ethers.getContractFactory("Gateway");
    const Custody = await ethers.getContractFactory("ERC20CustodyNew");
    [owner, destination] = await ethers.getSigners();

    // Deploy the contracts
    token = await TestERC20.deploy("Test Token", "TTK");
    receiver = await Receiver.deploy();
    gateway = await upgrades.deployProxy(Gateway, [], {
      initializer: "initialize",
      kind: "uups",
    });
    custody = await Custody.deploy(gateway.address);

    gateway.setCustody(custody.address);

    // Mint initial supply to the owner
    await token.mint(owner.address, ethers.utils.parseEther("1000"));

    // Transfer some tokens to the custody contract
    await token.transfer(custody.address, ethers.utils.parseEther("500"));
  });

  it("should upgrade and forward call to Receiver's receiveA function", async function () {
    // Upgrade Gateway contract
    const GatewayV2 = await ethers.getContractFactory("GatewayV2");
    const gatewayV2 = await upgrades.upgradeProxy(gateway.address, GatewayV2);

    // Forward call
    const str = "Hello, Hardhat!";
    const num = 42;
    const flag = true;
    const value = ethers.utils.parseEther("1.0");

    // Encode the function call data
    const data = receiver.interface.encodeFunctionData("receiveA", [str, num, flag]);

    // Call execute on the GatewayV2 contract
    const tx = await gatewayV2.execute(receiver.address, data, { value: value });
    await tx.wait();

    // Listen for the event
    await expect(tx).to.emit(gatewayV2, "ExecutedV2").withArgs(receiver.address, value, data);
    await expect(tx).to.emit(receiver, "ReceivedA").withArgs(gatewayV2.address, value, str, num, flag);
  });
});
