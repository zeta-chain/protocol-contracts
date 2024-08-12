// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "forge-std/Script.sol";
import "src/zevm/GatewayZEVM.sol";
import "src/zevm/ZRC20.sol";
import "test/utils/ReceiverEVM.sol";

// ZevmWithdrawAndCallScript executes withdrawAndCall method on GatewayZEVM and it should be used on localnet.
// It uses anvil private key, and sets default contract addresses deployed on fresh localnet, that can be overriden with env vars.
contract ZevmWithdrawAndCallScript is Script {
    function run() external {
        address payable gatewayZEVMAddress = payable(vm.envOr("GATEWAY_ZEVM", 0x610178dA211FEF7D417bC0e6FeD39F05609AD788));
        address payable receiverEVMAddress = payable(vm.envOr("RECEIVER_EVM", 0x0B306BF915C4d645ff596e518fAf3F9669b97016));
        address zrc20Address = vm.envOr("ZRC20", 0x2ca7d64A7EFE2D62A725E2B35Cf7230D6677FfEe);
        uint256 amount = vm.envOr("AMOUNT", uint256(1));
        string memory mnemonic = "test test test test test test test test test test test junk";
        uint256 privateKey = vm.deriveKey(mnemonic, 0);
        address deployer = vm.rememberKey(privateKey);

        vm.startBroadcast(deployer);

        GatewayZEVM gatewayZEVM = GatewayZEVM(gatewayZEVMAddress);
        ReceiverEVM receiverEVM = ReceiverEVM(receiverEVMAddress);
        ZRC20 zrc20 = ZRC20(zrc20Address);

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

        RevertOptions memory revertOptions =
            RevertOptions({ revertAddress: address(0x321), callOnRevert: true, abortAddress: address(0x321), revertMessage: "" });

        try gatewayZEVM.withdrawAndCall(
            abi.encodePacked(receiverEVMAddress),
            amount,
            address(zrc20),
            message,
            revertOptions
        ) {
            console.log("ReceiverEVM called from ZEVM.");
        } catch (bytes memory err) {
            console.log("Error calling ReceiverEVM:");
            console.logBytes(err);
        }

        vm.stopBroadcast();
    }
}
