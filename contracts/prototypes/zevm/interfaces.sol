// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

interface IGatewayZEVM {
    function withdraw(bytes memory receiver, uint256 amount, address zrc20) external;

    function withdrawAndCall(bytes memory receiver, uint256 amount, address zrc20, bytes calldata message) external;

    function call(bytes memory receiver, bytes calldata message) external;
}