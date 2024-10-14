// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "forge-std/Script.sol";
import "contracts/evm/ERC20Custody.sol";
import { ERC1967Proxy } from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";

contract DeployERC20Custody is Script {
    function run() external {
        address payable tss = payable(vm.envAddress("TSS_ADDRESS"));
        address admin = vm.envAddress("ERC20_CUSTODY_ADMIN_ADDRESS_EVM");
        address gateway = vm.envAddress("GATEWAY_PROXY_EVM");

        address expectedImplAddress;
        address expectedProxyAddress;

        bytes32 implSalt = keccak256("ERC20Custody");
        bytes32 proxySalt = keccak256("ERC20CustodyProxy");

        vm.startBroadcast();

        expectedImplAddress = vm.computeCreate2Address(
            implSalt,
            hashInitCode(type(ERC20Custody).creationCode)
        );

        ERC20Custody erc20CustodyImpl = new ERC20Custody{salt: implSalt}();
        require(address(erc20CustodyImpl) != address(0), "erc20CustodyImpl deployment failed");

        require(expectedImplAddress == address(erc20CustodyImpl), "impl address doesn't match expected address");

        expectedProxyAddress = vm.computeCreate2Address(
            proxySalt,
            hashInitCode(
                type(ERC1967Proxy).creationCode,
                abi.encode(
                    address(erc20CustodyImpl),
                    abi.encodeWithSelector(ERC20Custody.initialize.selector, gateway, tss, admin)
                )
            )
        );

        ERC1967Proxy erc20CustodyProxy = new ERC1967Proxy{salt: proxySalt}(
            address(erc20CustodyImpl),
            abi.encodeWithSelector(ERC20Custody.initialize.selector, gateway, tss, admin)
        );
        require(address(erc20CustodyProxy) != address(0), "erc20CustodyProxy deployment failed");

        require(expectedProxyAddress == address(erc20CustodyProxy), "proxy address doesn't match expected address");

        ERC20Custody erc20Custody = ERC20Custody(address(erc20CustodyProxy));
        require(erc20Custody.tssAddress() == tss, "tss not set");
        require(address(erc20Custody.gateway()) == gateway, "gateway not set");

        vm.stopBroadcast();
    }
}