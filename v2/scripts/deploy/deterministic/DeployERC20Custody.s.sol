// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "forge-std/Script.sol";
import "src/evm/ERC20Custody.sol";

contract DeployERC20Custody is Script {
    function run() external {
        address payable tss = payable(vm.envAddress("TSS_ADDRESS"));
        address admin = vm.envAddress("ERC20_CUSTODY_ADMIN_ADDRESS_EVM");
        address gateway = vm.envAddress("GATEWAY_PROXY_EVM");

        bytes32 salt = keccak256("ERC20Custody");

        vm.startBroadcast();

        ERC20Custody custody = new ERC20Custody{salt: salt}(gateway, tss, admin);

        address expectedAddr = vm.computeCreate2Address(
            salt,
            hashInitCode(type(ERC20Custody).creationCode, abi.encode(gateway, tss, admin))
        );

        require(expectedAddr == address(custody), "erc20 custody address doesn't match expected address");

        vm.stopBroadcast();
    }
}