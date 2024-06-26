import { expect } from "chai";
import { Contract } from "ethers";
import { ethers, upgrades } from "hardhat";

describe("Gateway and Receiver", function () {
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

  it("should forward call to Receiver's receiveA function", async function () {
    const str = "Hello, Hardhat!";
    const num = 42;
    const flag = true;
    const value = ethers.utils.parseEther("1.0");

    // Encode the function call data
    const data = receiver.interface.encodeFunctionData("receiveA", [str, num, flag]);

    // Call execute on the Gateway contract
    const tx = await gateway.execute(receiver.address, data, { value: value });
    await tx.wait();

    // Listen for the event
    await expect(tx).to.emit(gateway, "Executed").withArgs(receiver.address, value, data);
    await expect(tx).to.emit(receiver, "ReceivedA").withArgs(gateway.address, value, str, num, flag);
  });

  it("should forward call to Receiver's receiveB function", async function () {
    const strs = ["Hello", "Hardhat"];
    const nums = [1, 2, 3];
    const flag = false;
    const data = receiver.interface.encodeFunctionData("receiveB", [strs, nums, flag]);
    const tx = await gateway.execute(receiver.address, data);
    await tx.wait();
    await expect(tx).to.emit(receiver, "ReceivedB").withArgs(gateway.address, strs, nums, flag);
  });

  it("should forward call with withdrawAndCall and give allowance to destination contract", async function () {
    const amount = ethers.utils.parseEther("100");

    // Encode the function call data for receiveC
    const data = receiver.interface.encodeFunctionData("receiveC", [amount, token.address, destination.address]);

    // Withdraw and call
    const tx = await custody.withdrawAndCall(token.address, receiver.address, amount, data);
    await tx.wait();

    // Verify the event was emitted
    await expect(tx)
      .to.emit(receiver, "ReceivedC")
      .withArgs(gateway.address, amount, token.address, destination.address);

    // Verify that the tokens were transferred to the destination address
    const receiverBalance = await token.balanceOf(destination.address);
    expect(receiverBalance).to.equal(amount);

    // Verify that the remaining tokens were refunded to the Custody contract
    const remainingBalance = await token.balanceOf(custody.address);
    expect(remainingBalance).to.equal(ethers.utils.parseEther("400"));

    // Verify that the approval was reset
    const allowance = await token.allowance(gateway.address, receiver.address);
    expect(allowance).to.equal(0);
  });

  it("should forward call to Receiver's receiveD function", async function () {
    const data = receiver.interface.encodeFunctionData("receiveD");

    // Execute the call
    const tx = await gateway.execute(receiver.address, data);
    await tx.wait();

    // Verify the event was emitted
    await expect(tx).to.emit(receiver, "ReceivedD").withArgs(gateway.address);
  });

  it("should forward call to Receiver's receiveD function through withdrawAndCall and return ERC20 tokens to custody", async function () {
    const amount = ethers.utils.parseEther("100");

    // Encode the function call data for receiveD
    const data = receiver.interface.encodeFunctionData("receiveD");

    // Withdraw and call
    await custody.withdrawAndCall(token.address, receiver.address, amount, data);

    // Verify the event was emitted
    await expect(custody.withdrawAndCall(token.address, receiver.address, amount, data))
      .to.emit(receiver, "ReceivedD")
      .withArgs(gateway.address);

    // Verify that the remaining tokens were refunded to the Custody contract
    const remainingBalance = await token.balanceOf(custody.address);
    expect(remainingBalance).to.equal(ethers.utils.parseEther("500"));

    // Verify that the approval was reset
    const allowance = await token.allowance(gateway.address, receiver.address);
    expect(allowance).to.equal(0);
  });
});
