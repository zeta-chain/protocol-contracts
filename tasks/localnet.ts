import { task } from "hardhat/config";

declare const hre: any;

// Contains tasks to make it easier to interact with prototype contracts localnet
// To make use of default contract addresses on localnet, start localnode and localnet from scratch, so contracts are deployed on same addresses
// Otherwise, provide custom addresses as parameters

task("call-evm-receiver", "calls evm receiver from zevm account")
  .addOptionalParam("gatewayZEVM", "contract address of gateway on ZEVM", "0x5133BBdfCCa3Eb4F739D599ee4eC45cBCD0E16c5")
  .addOptionalParam("receiverEVM", "contract address of receiver on EVM", "0xB06c856C8eaBd1d8321b687E188204C1018BC4E5")
  .setAction(async (taskArgs) => {
    const gatewayZEVM = await hre.ethers.getContractAt("GatewayZEVM", taskArgs.gatewayZEVM);
    const receiverEVM = await hre.ethers.getContractAt("ReceiverEVM", taskArgs.receiverEVM);

    const str = "Hello!";
    const num = 42;
    const flag = true;

    // Encode the function call data and call on zevm
    const message = receiverEVM.interface.encodeFunctionData("receiveA", [str, num, flag]);
    const callTx = await gatewayZEVM.call(receiverEVM.address, message);

    await callTx.wait();
    console.log("ReceiverEVM called from ZEVM");
  });

task("call-zevm-zcontract", "calls zevm zcontract from evm account")
  .addOptionalParam("gatewayEVM", "contract address of gateway on EVM", "0xAD523115cd35a8d4E60B3C0953E0E0ac10418309")
  .addOptionalParam("zContract", "contract address of zContract on ZEVM", "0x71089Ba41e478702e1904692385Be3972B2cBf9e")
  .setAction(async (taskArgs) => {
    const gatewayEVM = await hre.ethers.getContractAt("GatewayEVM", taskArgs.gatewayEVM);
    const zContract = await hre.ethers.getContractAt("TestZContract", taskArgs.zContract);

    const message = hre.ethers.utils.defaultAbiCoder.encode(["string"], ["hello"]);
    const callTx = await gatewayEVM.call(zContract.address, message);

    await callTx.wait();
    console.log("TestZContract called from EVM");
  });
