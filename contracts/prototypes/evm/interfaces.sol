// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IGatewayEVMEvents {
    event Executed(address indexed destination, uint256 value, bytes data);
    event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data);
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

interface IReceiverEVMEvents {
    event ReceivedPayable(address sender, uint256 value, string str, uint256 num, bool flag);
    event ReceivedNonPayable(address sender, string[] strs, uint256[] nums, bool flag);
    event ReceivedERC20(address sender, uint256 amount, address token, address destination);
    event ReceivedNoParams(address sender);
}

interface IGatewayEVM {
    function executeWithERC20(
        address token,
        address to,
        uint256 amount,
        bytes calldata data
    ) external returns (bytes memory);

    function execute(address destination, bytes calldata data) external payable returns (bytes memory);
}