import { expect } from "chai";
import { Contract } from "ethers";
import { ethers } from "hardhat";
import { UniswapV2Deployer } from "uniswap-v2-deploy-plugin";

import {
  ERC20,
  ERC20CustodyNew,
  Gateway,
  Receiver,
  TestERC20,
  UniswapV2Factory,
  UniswapV2Pair,
  UniswapV2Router02,
} from "../../typechain-types";

describe("Uniswap Integration with Gateway", function () {
  let tokenA: TestERC20;
  let tokenB: TestERC20;
  let factory: UniswapV2Factory;
  let router: UniswapV2Router02;
  let pair: UniswapV2Pair;
  let custody: ERC20CustodyNew;
  let gateway: Gateway;
  let owner, addr1, addr2;

  beforeEach(async function () {
    [owner, addr1, addr2] = await ethers.getSigners();

    // Deploy TestERC20 tokens
    const TestERC20 = await ethers.getContractFactory("TestERC20");
    tokenA = (await TestERC20.deploy("Token A", "TKA")) as TestERC20;
    tokenB = (await TestERC20.deploy("Token B", "TKB")) as TestERC20;
    await tokenA.mint(owner.address, ethers.utils.parseEther("1000"));
    await tokenB.mint(owner.address, ethers.utils.parseEther("1000"));

    const { factory: newFactory, router: newRouter, weth9: weth } = await UniswapV2Deployer.deploy(owner);

    factory = newFactory;
    router = newRouter;

    // Approve Router to move tokens
    await tokenA.approve(router.address, ethers.utils.parseEther("1000"));
    await tokenB.approve(router.address, ethers.utils.parseEther("1000"));

    // Add Liquidity
    await router.addLiquidity(
      tokenA.address,
      tokenB.address,
      ethers.utils.parseEther("500"),
      ethers.utils.parseEther("500"),
      0,
      0,
      owner.address,
      Math.floor(Date.now() / 1000) + 60 * 20
    );

    // Deploy Gateway and Custody Contracts
    const Gateway = await ethers.getContractFactory("Gateway");
    const ERC20CustodyNew = await ethers.getContractFactory("ERC20CustodyNew");
    gateway = (await Gateway.deploy()) as Gateway;
    custody = (await ERC20CustodyNew.deploy(gateway.address)) as ERC20CustodyNew;

    // Transfer some tokens to the custody contract
    await tokenA.transfer(custody.address, ethers.utils.parseEther("100"));
    await tokenB.transfer(custody.address, ethers.utils.parseEther("100"));
  });

  it("should perform a token swap on Uniswap and transfer the output tokens to a destination address", async function () {
    const amountIn = ethers.utils.parseEther("50");

    const data = router.interface.encodeFunctionData("swapExactTokensForTokens", [
      amountIn,
      0,
      [tokenA.address, tokenB.address],
      addr2.address,
      Math.floor(Date.now() / 1000) + 60 * 20,
    ]);

    // Withdraw and call
    await custody.withdrawAndCall(tokenA.address, router.address, amountIn, data);

    // Verify the destination address received the tokens
    const destBalance = await tokenB.balanceOf(addr2.address);
    expect(destBalance).to.be.gt(0);

    // Verify the remaining tokens are refunded to the Custody contract
    const remainingBalance = await tokenA.balanceOf(custody.address);
    expect(remainingBalance).to.equal(ethers.utils.parseEther("50"));

    // Verify the approval was reset
    const allowance = await tokenA.allowance(gateway.address, router.address);
    expect(allowance).to.equal(0);
  });
});
