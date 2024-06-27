// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

interface IGatewayEVM {
    function executeWithERC20(
        address token,
        address to,
        uint256 amount,
        bytes calldata data
    ) external returns (bytes memory);

    function execute(address destination, bytes calldata data) external payable returns (bytes memory);

    function sendERC20(bytes calldata recipient, address asset, uint256 amount) external;

    function send(bytes calldata recipient, uint256 amount) external payable;
}