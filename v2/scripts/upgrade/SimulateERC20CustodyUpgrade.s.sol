// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "forge-std/console.sol";
import "forge-std/Script.sol";
import "contracts/evm/ERC20Custody.sol";

contract UpgradeSimulation is Script {
    function run() external {
        // Forked environment
        address proxyAddress = vm.envAddress("PROXY_ADDRESS");
        address adminAddress = vm.envAddress("ADMIN_ADDRESS");

        // Deploy the new implementation contract
        bytes32 implSalt = keccak256("ERC20Custody");
        address erc20CustodyImpl = address(new ERC20Custody{salt: implSalt}());

        // Simulate the upgrade
        ERC20Custody proxy = ERC20Custody(proxyAddress);
        vm.prank(adminAddress);
        proxy.upgradeToAndCall(erc20CustodyImpl, "");

        // After upgrade, verify that the state is intact
        console.log("Upgraded contract state:");
        console.log("gateway address:", address(proxy.gateway()));
        console.log("tss address:", proxy.tssAddress());
        console.log("supportsLegacy address:", proxy.supportsLegacy());
    }
}