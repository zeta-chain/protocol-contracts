// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "forge-std/Script.sol";
import "src/evm/GatewayEVM.sol";
import "test/utils/GatewayEVMUpgradeTest.sol";
import "test/utils/TestERC20.sol";
import { ERC1967Proxy } from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import { Upgrades } from "openzeppelin-foundry-upgrades/Upgrades.sol";

contract UpgradeGatewayEVM is Script {
    function run() external {
        // TODO (https://github.com/zeta-chain/protocol-contracts/issues/251): should be passed as arg
        string memory mnemonic = "test test test test test test test test test test test junk";
        uint256 privateKey = vm.deriveKey(mnemonic, 0);
        address deployer = vm.rememberKey(privateKey);

        // TODO (https://github.com/zeta-chain/protocol-contracts/issues/251): should be passed as arg
        address proxy = vm.envOr("PROXY_ADDRESS", address(0xA7806c719bd377F15bA6CaDf2F94Afb7FfA66256));

        GatewayEVM prevImpl = GatewayEVM(proxy);

        vm.startBroadcast(deployer);

        Upgrades.upgradeProxy(proxy, "GatewayEVMUpgradeTest.sol", "");

        GatewayEVM newImpl = GatewayEVM(proxy);

        require(prevImpl.tssAddress() == newImpl.tssAddress(), "tss addr changed");
        require(prevImpl.zetaToken() == newImpl.zetaToken(), "zeta token addr changed");

        vm.stopBroadcast();
    }
}