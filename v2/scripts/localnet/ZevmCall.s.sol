// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "src/zevm/GatewayZEVM.sol";
import "test/utils/ReceiverEVM.sol";

contract ZevmCallScript is Script {
    function run() external {
        address payable gatewayZEVMAddress = payable(vm.envOr("GATEWAY_ZEVM", 0x610178dA211FEF7D417bC0e6FeD39F05609AD788));
        address payable receiverEVMAddress = payable(vm.envOr("RECEIVER_EVM", 0x0B306BF915C4d645ff596e518fAf3F9669b97016));
        string memory mnemonic = "test test test test test test test test test test test junk";
        uint256 privateKey = vm.deriveKey(mnemonic, 0);
        address deployer = vm.rememberKey(privateKey);

        vm.startBroadcast(deployer);

        GatewayZEVM gatewayZEVM = GatewayZEVM(gatewayZEVMAddress);
        ReceiverEVM receiverEVM = ReceiverEVM(receiverEVMAddress);

        string memory str = "Hello!";
        uint256 num = 42;
        bool flag = true;

        // Encode the function call data
        bytes memory message = abi.encodeWithSelector(
            receiverEVM.receivePayable.selector,
            str,
            num,
            flag
        );

        // Call the function on GatewayZEVM
        try gatewayZEVM.call(abi.encodePacked(address(receiverEVM)), message) {
            console.log("ReceiverEVM called from ZEVM.");
        } catch (bytes memory err) {
            console.log("Error calling ReceiverEVM:");
            console.logBytes(err);
        }

        vm.stopBroadcast();
    }
}