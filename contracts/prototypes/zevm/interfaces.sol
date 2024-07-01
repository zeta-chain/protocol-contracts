// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

import "../../zevm/interfaces/zContract.sol";

interface IGatewayZEVM {
    function withdraw(bytes memory receiver, uint256 amount, address zrc20) external;

    function withdrawAndCall(bytes memory receiver, uint256 amount, address zrc20, bytes calldata message) external;

    function call(bytes memory receiver, bytes calldata message) external;

    function deposit(
        address zrc20,
        uint256 amount,
        address target
    ) external;

    function execute(
        zContext calldata context,
        address zrc20,
        uint256 amount,
        address target,
        bytes calldata message
    ) external;

    function depositAndCall(
        zContext calldata context,
        address zrc20,
        uint256 amount,
        address target,
        bytes calldata message
    ) external;
}