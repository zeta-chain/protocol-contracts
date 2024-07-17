// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./interfaces.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";

// As the current version, ERC20CustodyNew hold the ERC20s deposited on ZetaChain
// This version include a functionality allowing to call a contract
// ERC20Custody doesn't call smart contract directly, it passes through the Gateway contract
contract ERC20CustodyNew is ReentrancyGuard{
    using SafeERC20 for IERC20;
    error ZeroAddress();

    IGatewayEVM public gateway;

    event Withdraw(address indexed token, address indexed to, uint256 amount);
    event WithdrawAndCall(address indexed token, address indexed to, uint256 amount, bytes data);
    event WithdrawAndRevert(address indexed token, address indexed to, uint256 amount, bytes data);

    constructor(address _gateway) {
         if (_gateway == address(0)) {
            revert ZeroAddress();
        }
        gateway = IGatewayEVM(_gateway);
    }
    
    // Withdraw is called by TSS address, it directly transfers the tokens to the destination address without contract call
    // TODO: Finalize access control
    // https://github.com/zeta-chain/protocol-contracts/issues/204
    function withdraw(address token, address to, uint256 amount) external nonReentrant {
        IERC20(token).safeTransfer(to, amount);

        emit Withdraw(token, to, amount);
    }

    // WithdrawAndCall is called by TSS address, it transfers the tokens and call a contract
    // For this, it passes through the Gateway contract, it transfers the tokens to the Gateway contract and then calls the contract
    // TODO: Finalize access control
    // https://github.com/zeta-chain/protocol-contracts/issues/204
    function withdrawAndCall(address token, address to, uint256 amount, bytes calldata data) public nonReentrant {
        // Transfer the tokens to the Gateway contract
        IERC20(token).safeTransfer(address(gateway), amount);

        // Forward the call to the Gateway contract
        gateway.executeWithERC20(token, to, amount, data);

        emit WithdrawAndCall(token, to, amount, data);
    }

    // WithdrawAndRevert is called by TSS address, it transfers the tokens and call a contract
    // For this, it passes through the Gateway contract, it transfers the tokens to the Gateway contract and then calls the contract
    // TODO: Finalize access control
    // https://github.com/zeta-chain/protocol-contracts/issues/204
    function withdrawAndRevert(address token, address to, uint256 amount, bytes calldata data) public nonReentrant {
        // Transfer the tokens to the Gateway contract
        IERC20(token).safeTransfer(address(gateway), amount);

        // Forward the call to the Gateway contract
        gateway.revertWithERC20(token, to, amount, data);

        emit WithdrawAndRevert(token, to, amount, data);
    }
}