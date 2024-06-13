import { MaxUint256 } from "@ethersproject/constants";
import { parseEther, parseUnits } from "@ethersproject/units";
import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import {
  type ZetaTokenConsumerZEVM,
  IERC20,
  IERC20__factory,
  UniswapV2Router02__factory,
  ZetaTokenConsumerZEVM__factory,
} from "@typechain-types";
import chai, { expect } from "chai";
import { BigNumber } from "ethers";
import { ethers } from "hardhat";

import { WETH9__factory } from "../typechain-types/factories/contracts/zevm/WZETA.sol/WETH9__factory";
import { parseZetaConsumerLog } from "./test.helpers";

chai.should();

const USDC_ADDRESS = "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48";
const WETH_ADDRESS = "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2";
const UNI_V2_ROUTER_ADDRESS = "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D";

describe("ZetaTokenConsumerZEVM tests", () => {
  let zetaTokenConsumerZEVM: ZetaTokenConsumerZEVM;
  let zetaToken: IERC20;

  let accounts: SignerWithAddress[];
  let deployer: SignerWithAddress;
  let randomSigner: SignerWithAddress;

  const getNow = async () => {
    const block = await ethers.provider.getBlock("latest");
    return block.timestamp;
  };

  const swapToken = async (signer: SignerWithAddress, tokenAddress: string, expectedAmount: BigNumber) => {
    const uniswapRouter = UniswapV2Router02__factory.connect(UNI_V2_ROUTER_ADDRESS, signer);
    const WETH = await uniswapRouter.WETH();

    const path = [WETH, tokenAddress];
    const tx = await uniswapRouter
      .connect(signer)
      .swapETHForExactTokens(expectedAmount, path, signer.address, (await getNow()) + 360, { value: parseEther("10") });
    await tx.wait();
  };

  beforeEach(async () => {
    accounts = await ethers.getSigners();
    [deployer, randomSigner] = accounts;

    zetaToken = IERC20__factory.connect(WETH_ADDRESS, deployer);

    const zetaTokenConsumerZEVMFactory = new ZetaTokenConsumerZEVM__factory(deployer);
    zetaTokenConsumerZEVM = await zetaTokenConsumerZEVMFactory.deploy(WETH_ADDRESS, UNI_V2_ROUTER_ADDRESS);
  });

  describe("getZetaFromEth", () => {
    it("Should get zeta from eth", async () => {
      const initialZetaBalance = await zetaToken.balanceOf(randomSigner.address);
      const tx = await zetaTokenConsumerZEVM.getZetaFromEth(randomSigner.address, 1, { value: parseEther("1") });

      const result = await tx.wait();
      const eventNames = parseZetaConsumerLog(result.logs);
      expect(eventNames.filter((e) => e === "EthExchangedForZeta")).to.have.lengthOf(1);

      const finalZetaBalance = await zetaToken.balanceOf(randomSigner.address);
      expect(finalZetaBalance).to.be.gt(initialZetaBalance);
    });
  });

  describe("getZetaFromToken", () => {
    it("Should get zeta from token", async () => {
      const USDCContract = IERC20__factory.connect(USDC_ADDRESS, deployer);
      await swapToken(deployer, USDC_ADDRESS, parseUnits("10000", 6));

      const initialZetaBalance = await zetaToken.balanceOf(randomSigner.address);
      const tx1 = await USDCContract.approve(zetaTokenConsumerZEVM.address, MaxUint256);
      await tx1.wait();

      const tx2 = await zetaTokenConsumerZEVM.getZetaFromToken(
        randomSigner.address,
        1,
        USDC_ADDRESS,
        parseUnits("100", 6)
      );
      const result = await tx2.wait();

      const eventNames = parseZetaConsumerLog(result.logs);
      expect(eventNames.filter((e) => e === "TokenExchangedForZeta")).to.have.lengthOf(1);

      const finalZetaBalance = await zetaToken.balanceOf(randomSigner.address);
      expect(finalZetaBalance).to.be.gt(initialZetaBalance);
    });
  });

  describe("getEthFromZeta", async () => {
    it("Should get eth from zeta", async () => {
      const amount = parseUnits("50", 18);
      const WETH9 = WETH9__factory.connect(WETH_ADDRESS, deployer);

      const tx = await WETH9.deposit({ value: amount });
      await tx.wait();

      const initialEthBalance = await ethers.provider.getBalance(randomSigner.address);
      const tx1 = await WETH9.approve(zetaTokenConsumerZEVM.address, MaxUint256);
      await tx1.wait();

      const zetaBalance = await zetaToken.balanceOf(deployer.address);

      const tx2 = await zetaTokenConsumerZEVM.getEthFromZeta(randomSigner.address, 1, amount);
      const result = await tx2.wait();

      const eventNames = parseZetaConsumerLog(result.logs);
      expect(eventNames.filter((e) => e === "ZetaExchangedForEth")).to.have.lengthOf(1);

      const finalEthBalance = await ethers.provider.getBalance(randomSigner.address);
      expect(finalEthBalance).to.be.gt(initialEthBalance);
    });
  });

  describe("getTokenFromZeta", async () => {
    it("Should get token from zeta", async () => {
      const amount = parseUnits("50", 18);

      {
        const WETH9 = WETH9__factory.connect(WETH_ADDRESS, deployer);
        const tx = await WETH9.deposit({ value: amount });
        await tx.wait();
      }
      const USDCContract = IERC20__factory.connect(USDC_ADDRESS, randomSigner);

      const initialTokenBalance = await USDCContract.balanceOf(randomSigner.address);
      const tx1 = await zetaToken.approve(zetaTokenConsumerZEVM.address, MaxUint256);
      await tx1.wait();

      const tx2 = await zetaTokenConsumerZEVM.getTokenFromZeta(randomSigner.address, 1, USDC_ADDRESS, amount);
      const result = await tx2.wait();

      const eventNames = parseZetaConsumerLog(result.logs);
      expect(eventNames.filter((e) => e === "ZetaExchangedForToken")).to.have.lengthOf(1);

      const finalTokenBalance = await USDCContract.balanceOf(randomSigner.address);
      expect(finalTokenBalance).to.be.gt(initialTokenBalance);
    });
  });
});
