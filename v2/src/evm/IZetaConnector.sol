// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

interface IZetaConnectorEvents {
    event Withdraw(address indexed to, uint256 amount);
    event WithdrawAndCall(address indexed to, uint256 amount, bytes data);
    event WithdrawAndRevert(address indexed to, uint256 amount, bytes data);
}