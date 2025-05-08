// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "forge-std/Script.sol";
import "contracts/zevm/Registry.sol";
import { ERC1967Proxy } from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";

contract DeployRegistry is Script {
    function run() external {
        address admin = vm.envAddress("REGISTRY_ADMIN_ADDRESS");
        address gatewayEVM = vm.envAddress("GATEWAY_EVM_ADDRESS");
        address coreRegistry = vm.envAddress("CORE_REGISTRY_ADDRESS");

        address expectedImplAddress;
        address expectedProxyAddress;

        bytes32 implSalt = keccak256("Registry");
        bytes32 proxySalt = keccak256("RegistryProxy");

        // Add this specific check to ensure contract is not deployed at deterministic address with wrong params
        require(admin != address(0), "Environment variable REGISTRY_ADMIN_ADDRESS is not set");
        require(gatewayEVM != address(0), "Environment variable GATEWAY_EVM_ADDRESS is not set");
        require(coreRegistry != address(0), "Environment variable CORE_REGISTRY_ADDRESS is not set");

        vm.startBroadcast();

        // Compute expected implementation address
        expectedImplAddress = vm.computeCreate2Address(
            implSalt,
            hashInitCode(type(Registry).creationCode)
        );

        // Deploy the implementation contract using CREATE2
        Registry registryImpl = new Registry{salt: implSalt}();
        require(address(registryImpl) != address(0), "registryImpl deployment failed");
        require(expectedImplAddress == address(registryImpl), "impl address doesn't match expected address");

        // Compute expected proxy address
        expectedProxyAddress = vm.computeCreate2Address(
            proxySalt,
            hashInitCode(
                type(ERC1967Proxy).creationCode,
                abi.encode(
                    address(registryImpl),
                    abi.encodeWithSelector(Registry.initialize.selector, admin, gatewayEVM, coreRegistry)
                )
            )
        );

        // Deploy the proxy contract using CREATE2
        ERC1967Proxy registryProxy = new ERC1967Proxy{salt: proxySalt}(
            address(registryImpl),
            abi.encodeWithSelector(Registry.initialize.selector, admin, gatewayEVM, coreRegistry)
        );
        require(address(registryProxy) != address(0), "registryProxy deployment failed");
        require(expectedProxyAddress == address(registryProxy), "proxy address doesn't match expected address");

        Registry registry = Registry(address(registryProxy));

        // Verify initial configuration
        require(registry.gatewayEVM() == gatewayEVM, "gatewayEVM not set correctly");
        require(registry.coreRegistry() == coreRegistry, "coreRegistry not set correctly");
        require(registry.hasRole(registry.DEFAULT_ADMIN_ROLE(), admin), "admin role not set correctly");
        require(registry.hasRole(registry.RELAY_ROLE(), gatewayEVM), "relay role not set correctly");

        vm.stopBroadcast();

        console.log("Registry Implementation deployed at: %s", address(registryImpl));
        console.log("Registry Proxy deployed at: %s", address(registryProxy));
    }
}