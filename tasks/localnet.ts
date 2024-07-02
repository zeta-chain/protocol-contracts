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
