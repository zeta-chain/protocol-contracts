// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

interface IERC20CustodyNewEvents {
    event Withdraw(address indexed token, address indexed to, uint256 amount);
    event WithdrawAndCall(address indexed token, address indexed to, uint256 amount, bytes data);
    event WithdrawAndRevert(address indexed token, address indexed to, uint256 amount, bytes data);
}

interface IERC20CustodyNewErrors {
    error ZeroAddress();
    error InvalidSender();
}