// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "src/evm/GatewayEVM.sol";
import "test/utils/TestZContract.sol";

// EvmCallScript executes call method on GatewayEVM and it should be used on localnet.
// It uses anvil private key, and sets default contract addresses deployed on fresh localnet, that can be overriden with env vars.
contract EvmCallScript is Script {
    function run() external {
        address payable gatewayEVMAddress = payable(vm.envOr("GATEWAY_EVM", 0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0));
        address payable zContractAddress = payable(vm.envOr("Z_CONTRACT", 0x68B1D87F95878fE05B998F19b66F4baba5De1aed));
        string memory mnemonic = "test test test test test test test test test test test junk";
        uint256 privateKey = vm.deriveKey(mnemonic, 0);
        address deployer = vm.rememberKey(privateKey);

        vm.startBroadcast(deployer);

        GatewayEVM gatewayEVM = GatewayEVM(gatewayEVMAddress);
        TestZContract zContract = TestZContract(zContractAddress);

        // Encode the message
        bytes memory message = abi.encode("hello");

        // Call the function on GatewayEVM
        try gatewayEVM.call(zContractAddress, message) {
            console.log("TestZContract called from EVM.");
        } catch (bytes memory err) {
            console.log("Error calling TestZContract:");
            console.logBytes(err);
        }

        vm.stopBroadcast();
    }
}
