// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "forge-std/Script.sol";
import "src/evm/ZetaConnectorNonNative.sol";

contract DeployZetaConnectorNonNative is Script {
    function run() external {
        address payable tss = payable(vm.envAddress("TSS_ADDRESS"));
        address admin = vm.envAddress("ZETA_CONNECTOR_ADMIN_ADDRESS_EVM");
        address gateway = vm.envAddress("GATEWAY_PROXY_EVM");
        address zeta = vm.envAddress("ZETA_ERC20_EVM");

        bytes32 salt = keccak256("ZetaConnectorNonNative");

        vm.startBroadcast();

        ZetaConnectorNonNative connector = new ZetaConnectorNonNative{salt: salt}(gateway, zeta, tss, admin);
        require(address(connector) != address(0), "deployment failed");

        address expectedAddr = vm.computeCreate2Address(
            salt,
            hashInitCode(type(ZetaConnectorNonNative).creationCode, abi.encode(gateway, zeta, tss, admin))
        );

        require(expectedAddr == address(connector), "zeta connector non native address doesn't match expected address");

        vm.stopBroadcast();
    }
}