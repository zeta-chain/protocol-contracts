// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "forge-std/Script.sol";
import "contracts/zevm/GatewayZEVM.sol";
import { ERC1967Proxy } from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";

contract DeployGatewayZEVM is Script {
    function run() external {
        address admin = vm.envAddress("GATEWAY_ADMIN_ADDRESS_ZEVM");
        address zeta = vm.envAddress("WZETA");

        address expectedImplAddress;
        address expectedProxyAddress;

        bytes32 implSalt = keccak256("GatewayZEVM");
        bytes32 proxySalt = keccak256("GatewayZEVMProxy");

        vm.startBroadcast();

        expectedImplAddress = vm.computeCreate2Address(
            implSalt,
            hashInitCode(type(GatewayZEVM).creationCode)
        );

        GatewayZEVM gatewayImpl = new GatewayZEVM{salt: implSalt}();
        require(address(gatewayImpl) != address(0), "gatewayImpl deployment failed");

        require(expectedImplAddress == address(gatewayImpl), "impl address doesn't match expected address");

        expectedProxyAddress = vm.computeCreate2Address(
            proxySalt,
            hashInitCode(
                type(ERC1967Proxy).creationCode,
                abi.encode(
                    address(gatewayImpl),
                    abi.encodeWithSelector(GatewayZEVM.initialize.selector, zeta, admin)
                )
            )
        );

        ERC1967Proxy gatewayProxy = new ERC1967Proxy{salt: proxySalt}(
            address(gatewayImpl),
            abi.encodeWithSelector(GatewayZEVM.initialize.selector, zeta, admin)
        );
        require(address(gatewayProxy) != address(0), "gatewayProxy deployment failed");

        require(expectedProxyAddress == address(gatewayProxy), "proxy address doesn't match expected address");

        GatewayZEVM gateway = GatewayZEVM(payable(address(gatewayProxy)));
        require(gateway.zetaToken() == zeta, "zeta token not set");

        vm.stopBroadcast();
    }
}