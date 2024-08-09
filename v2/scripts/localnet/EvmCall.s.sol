// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "forge-std/Script.sol";
import "src/evm/GatewayEVM.sol";
import "test/utils/TestUniversalContract.sol";

// EvmCallScript executes call method on GatewayEVM and it should be used on localnet.
// It uses anvil private key, and sets default contract addresses deployed on fresh localnet, that can be overriden with env vars.
contract EvmCallScript is Script {
    function run() external {
        address payable gatewayEVMAddress = payable(vm.envOr("GATEWAY_EVM", 0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0));
        address payable universalContractAddress = payable(vm.envOr("UNIVERSAL_CONTRACT", 0x68B1D87F95878fE05B998F19b66F4baba5De1aed));
        string memory mnemonic = "test test test test test test test test test test test junk";
        uint256 privateKey = vm.deriveKey(mnemonic, 0);
        address deployer = vm.rememberKey(privateKey);

        vm.startBroadcast(deployer);

        GatewayEVM gatewayEVM = GatewayEVM(gatewayEVMAddress);
        TestUniversalContract universalContract = TestUniversalContract(universalContractAddress);

        // Encode the message
        bytes memory message = abi.encode("hello");

        // Call the function on GatewayEVM
        try gatewayEVM.call(universalContractAddress, message) {
            console.log("TestUniversalContract called from EVM.");
        } catch (bytes memory err) {
            console.log("Error calling TestUniversalContract:");
            console.logBytes(err);
        }

        vm.stopBroadcast();
    }
}
