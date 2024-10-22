// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "forge-std/Script.sol";
import "contracts/evm/GatewayEVM.sol";
import "test/utils/TestERC20.sol";
import { ERC1967Proxy } from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";

contract DeployGatewayEVM is Script {
    function run() external {
        address payable tss = payable(vm.envAddress("TSS_ADDRESS"));
        address admin = vm.envAddress("GATEWAY_ADMIN_ADDRESS_EVM");
        address zeta = vm.envAddress("ZETA_ERC20_EVM");

        address expectedImplAddress;
        address expectedProxyAddress;

        bytes32 implSalt = keccak256("GatewayEVM");
        bytes32 proxySalt = keccak256("GatewayEVMProxy");

        vm.startBroadcast();

        expectedImplAddress = vm.computeCreate2Address(
            implSalt,
            hashInitCode(type(GatewayEVM).creationCode)
        );

        GatewayEVM gatewayImpl = new GatewayEVM{salt: implSalt}();
        require(address(gatewayImpl) != address(0), "gatewayImpl deployment failed");

        require(expectedImplAddress == address(gatewayImpl), "impl address doesn't match expected address");

        expectedProxyAddress = vm.computeCreate2Address(
            proxySalt,
            hashInitCode(
                type(ERC1967Proxy).creationCode,
                abi.encode(
                    address(gatewayImpl),
                    abi.encodeWithSelector(GatewayEVM.initialize.selector, tss, address(zeta), admin)
                )
            )
        );

        ERC1967Proxy gatewayProxy = new ERC1967Proxy{salt: proxySalt}(
            address(gatewayImpl),
            abi.encodeWithSelector(GatewayEVM.initialize.selector, tss, address(zeta), admin)
        );
        require(address(gatewayProxy) != address(0), "gatewayProxy deployment failed");


        require(expectedProxyAddress == address(gatewayProxy), "proxy address doesn't match expected address");

        GatewayEVM gateway = GatewayEVM(address(gatewayProxy));
        require(gateway.tssAddress() == tss, "tss not set");
        require(gateway.zetaToken() == address(zeta), "zeta token not set");

        vm.stopBroadcast();
    }
}