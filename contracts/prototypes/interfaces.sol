// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

interface IGateway {
    function executeWithERC20(
        address token,
        address to,
        uint256 amount,
        bytes calldata data
    ) external returns (bytes memory);
}