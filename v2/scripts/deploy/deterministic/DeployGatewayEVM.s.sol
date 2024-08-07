// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "forge-std/Script.sol";
import "src/evm/GatewayEVM.sol";
import "test/utils/TestERC20.sol";
import { ERC1967Proxy } from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";

contract DeployGatewayEVMCreate2 is Script {
    function run() external {
        // TODO (https://github.com/zeta-chain/protocol-contracts/issues/251): should be passed as arg
        string memory mnemonic = "test test test test test test test test test test test junk";
        uint256 privateKey = vm.deriveKey(mnemonic, 0);
        address deployer = vm.rememberKey(privateKey);

        // TODO (https://github.com/zeta-chain/protocol-contracts/issues/251): should be passed as arg
        address payable tss = payable(vm.envOr("TSS_ADDRESS", address(0x123)));
        address admin = vm.envOr("ADMIN_ADDRESS", deployer);

        address expectedImplAddress;
        address expectedProxyAddress;

        bytes32 implSalt = bytes32(uint256(10));
        bytes32 proxySalt = bytes32(uint256(11));

        vm.startBroadcast(deployer);

        // TODO (https://github.com/zeta-chain/protocol-contracts/issues/251): should be passed as arg
        TestERC20 zeta = new TestERC20("zeta", "ZETA");

        expectedImplAddress = computeCreate2Address(
            implSalt,
            hashInitCode(type(GatewayEVM).creationCode)
        );

        GatewayEVM gatewayImpl = new GatewayEVM{salt: implSalt}();
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

        require(expectedProxyAddress == address(gatewayProxy), "proxy address doesn't match expected address");

        GatewayEVM gateway = GatewayEVM(address(gatewayProxy));
        require(gateway.tssAddress() == tss, "tss not set");
        require(gateway.zetaToken() == address(zeta), "zeta token not set");

        vm.stopBroadcast();
    }
}