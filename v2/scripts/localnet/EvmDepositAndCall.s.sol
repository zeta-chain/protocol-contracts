// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "forge-std/Script.sol";
import "src/evm/GatewayEVM.sol";
import "src/evm/ERC20Custody.sol";
import "test/utils/TestERC20.sol";

// EvmDepositAndCallScript executes depositAndCall method on GatewayEVM and it should be used on localnet.
// It uses anvil private key, and sets default contract addresses deployed on fresh localnet, that can be overriden with env vars.
contract EvmDepositAndCallScript is Script {
    function run() external {
        address payable gatewayEVMAddress = payable(vm.envOr("GATEWAY_EVM", 0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0));
        address custodyAddress = vm.envOr("CUSTODY", 0xDc64a140Aa3E981100a9becA4E685f962f0cF6C9);
        address payable universalContractAddress = payable(vm.envOr("Z_CONTRACT", 0x68B1D87F95878fE05B998F19b66F4baba5De1aed));
        address erc20Address = vm.envOr("ERC20", 0x9A676e781A523b5d0C0e43731313A708CB607508);
        uint256 amount = vm.envOr("AMOUNT", uint256(1));
        string memory mnemonic = "test test test test test test test test test test test junk";
        uint256 privateKey = vm.deriveKey(mnemonic, 0);
        address deployer = vm.rememberKey(privateKey);

        vm.startBroadcast(deployer);

        GatewayEVM gatewayEVM = GatewayEVM(gatewayEVMAddress);
        TestERC20 erc20 = TestERC20(erc20Address);
        ERC20Custody custody = ERC20Custody(custodyAddress);

        custody.whitelist(erc20Address);

        // Approve the ERC20 transfer
        erc20.approve(gatewayEVMAddress, amount);

        // Encode the payload
        bytes memory payload = abi.encode("hello");

        // Call the depositAndCall function on GatewayEVM
        try gatewayEVM.depositAndCall(
            universalContractAddress,
            amount,
            erc20Address,
            payload
        ) {
            console.log("TestUniversalContract called from EVM.");
        } catch (bytes memory err) {
            console.log("Error calling TestUniversalContract:");
            console.logBytes(err);
        }

        vm.stopBroadcast();
    }
}
