import { AddressZero } from "@ethersproject/constants";
import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { SystemContract, ZRC20 } from "@typechain-types";
import { expect } from "chai";
import { BigNumber, Contract } from "ethers";
import { parseEther } from "ethers/lib/utils";
import { ethers, upgrades } from "hardhat";

import { FUNGIBLE_MODULE_ADDRESS } from "../test.helpers";
const hre = require("hardhat");

describe("GatewayZEVM inbound", function () {
  let ZRC20Contract: ZRC20;
  let systemContract: SystemContract;
  let gateway: Contract;
  let owner: SignerWithAddress;
  let addrs: SignerWithAddress[];

  beforeEach(async () => {
    [owner, ...addrs] = await ethers.getSigners();

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

    await ZRC20Contract.connect(fungibleModuleSigner).deposit(owner.address, parseEther("100"));

    const Gateway = await ethers.getContractFactory("GatewayZEVM");
    gateway = await upgrades.deployProxy(Gateway, [], {
      initializer: "initialize",
      kind: "uups",
    });

    await ZRC20Contract.connect(owner).approve(gateway.address, parseEther("100"));
  });

  it("should withdraw zrc20 and emit event", async function () {
    const tx = await gateway
      .connect(owner)
      .withdraw(ethers.utils.arrayify(addrs[0].address), parseEther("1"), ZRC20Contract.address);

    const balanceOfAfterWithdrawal = (await ZRC20Contract.balanceOf(owner.address)) as BigNumber;
    expect(balanceOfAfterWithdrawal).to.equal(parseEther("99"));

    await expect(tx)
      .to.emit(gateway, "Withdrawal")
      .withArgs(
        ethers.utils.getAddress(owner.address),
        addrs[0].address.toLowerCase(),
        parseEther("1"),
        0,
        await ZRC20Contract.PROTOCOL_FLAT_FEE(),
        "0x"
      );
  });

  it("should withdraw zrc20 and call and emit event with message", async function () {
    let ABI = ["function hello(address to)"];
    let iface = new ethers.utils.Interface(ABI);
    const message = iface.encodeFunctionData("hello", ["0x1234567890123456789012345678901234567890"]);

    const tx = await gateway
      .connect(owner)
      .withdrawAndCall(ethers.utils.arrayify(addrs[0].address), parseEther("1"), ZRC20Contract.address, message);

    const balanceOfAfterWithdrawal = (await ZRC20Contract.balanceOf(owner.address)) as BigNumber;
    expect(balanceOfAfterWithdrawal).to.equal(parseEther("99"));

    await expect(tx)
      .to.emit(gateway, "Withdrawal")
      .withArgs(
        ethers.utils.getAddress(owner.address),
        addrs[0].address.toLowerCase(),
        parseEther("1"),
        0,
        await ZRC20Contract.PROTOCOL_FLAT_FEE(),
        message
      );
  });

  it("should call and emit event without asset transfer", async function () {
    let ABI = ["function hello(address to)"];
    let iface = new ethers.utils.Interface(ABI);
    const message = iface.encodeFunctionData("hello", ["0x1234567890123456789012345678901234567890"]);

    const tx = await gateway.connect(owner).call(ethers.utils.arrayify(addrs[0].address), message);
    await expect(tx)
      .to.emit(gateway, "Call")
      .withArgs(ethers.utils.getAddress(owner.address), addrs[0].address.toLowerCase(), message);
  });
});
