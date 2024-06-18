import { expect } from "chai";
import { ethers } from "hardhat";
import { Contract } from "ethers";

describe("Gateway and Receiver", function () {
  let receiver: Contract;
  let gateway: Contract;
  let owner: any, addr1: any;

  beforeEach(async function () {
    // Get the ContractFactories and Signers
    const Receiver = await ethers.getContractFactory("Receiver");
    const Gateway = await ethers.getContractFactory("Gateway");
    [owner, addr1] = await ethers.getSigners();

    // Deploy the contracts
    receiver = await Receiver.deploy();
    gateway = await Gateway.deploy();
  });

  it("should forward call to Receiver's receiveA function", async function () {
    const str = "Hello, Hardhat!";
    const num = 42;
    const flag = true;
    const value = ethers.utils.parseEther("1.0");

    // Encode the function call data
    const data = receiver.interface.encodeFunctionData("receiveA", [str, num, flag]);

    // Call forwardCall on the Gateway contract
    const tx = await gateway.forwardCall(receiver.address, data, { value: value });
    await tx.wait();

    // Listen for the event
    await expect(tx)
      .to.emit(receiver, "ReceivedA")
      .withArgs(gateway.address, value, str, num, flag);
  });

  it("should forward call to Receiver's receiveB function", async function () {
    const strs = ["Hello", "Hardhat"];
    const nums = [1, 2, 3];
    const flag = false;
    const data = receiver.interface.encodeFunctionData("receiveB", [strs, nums, flag]);
    const tx = await gateway.forwardCall(receiver.address, data);
    await tx.wait();
    await expect(tx)
      .to.emit(receiver, "ReceivedB")
      .withArgs(gateway.address, strs, nums, flag);
  });
});