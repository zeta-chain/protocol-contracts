// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IGatewayEVMEvents {
    event Executed(address indexed destination, uint256 value, bytes data);
    event Reverted(address indexed destination, uint256 value, bytes data);
    event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data);
    event RevertedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data);
    event Deposit(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload);
    event Call(address indexed sender, address indexed receiver, bytes payload);
}

interface IGatewayEVMErrors {
    error ExecutionFailed();
    error DepositFailed();
    error InsufficientETHAmount();
    error InsufficientERC20Amount();
    error ZeroAddress();
    error ApprovalFailed();
    error CustodyInitialized();
}

interface IGatewayEVM {
    function executeWithERC20(
        address token,
        address to,
        uint256 amount,
        bytes calldata data
    ) external;

    function execute(address destination, bytes calldata data) external payable returns (bytes memory);

    function revertWithERC20(
        address token,
        address to,
        uint256 amount,
        bytes calldata data
    ) external;
}

interface Revertable {
    function onRevert(bytes calldata data) external;
}
