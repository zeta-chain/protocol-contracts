// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./IGatewayEVM.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";

// As the current version, ERC20CustodyNew hold the ERC20s deposited on ZetaChain
// This version include a functionality allowing to call a contract
// ERC20Custody doesn't call smart contract directly, it passes through the Gateway contract
contract ERC20CustodyNew is ReentrancyGuard {
    using SafeERC20 for IERC20;
    error ZeroAddress();
    error InvalidSender();

    IGatewayEVM public gateway;
    address public tssAddress;

    event Withdraw(address indexed token, address indexed to, uint256 amount);
    event WithdrawAndCall(address indexed token, address indexed to, uint256 amount, bytes data);
    event WithdrawAndRevert(address indexed token, address indexed to, uint256 amount, bytes data);

    // @dev Only TSS address allowed modifier.
    modifier onlyTSS() {
        if (msg.sender != tssAddress) {
            revert InvalidSender();
        }
        _;
    }

    constructor(address _gateway, address _tssAddress) {
         if (_gateway == address(0) || _tssAddress == address(0)) {
            revert ZeroAddress();
        }
        gateway = IGatewayEVM(_gateway);
        tssAddress = _tssAddress;
    }

    // Withdraw is called by TSS address, it directly transfers the tokens to the destination address without contract call
    function withdraw(address token, address to, uint256 amount) external nonReentrant onlyTSS {
        IERC20(token).safeTransfer(to, amount);

        emit Withdraw(token, to, amount);
    }

    // WithdrawAndCall is called by TSS address, it transfers the tokens and call a contract
    // For this, it passes through the Gateway contract, it transfers the tokens to the Gateway contract and then calls the contract
    function withdrawAndCall(address token, address to, uint256 amount, bytes calldata data) public nonReentrant onlyTSS {
        // Transfer the tokens to the Gateway contract
        IERC20(token).safeTransfer(address(gateway), amount);

        // Forward the call to the Gateway contract
        gateway.executeWithERC20(token, to, amount, data);

        emit WithdrawAndCall(token, to, amount, data);
    }

    // WithdrawAndRevert is called by TSS address, it transfers the tokens and call a contract
    // For this, it passes through the Gateway contract, it transfers the tokens to the Gateway contract and then calls the contract
    function withdrawAndRevert(address token, address to, uint256 amount, bytes calldata data) public nonReentrant onlyTSS {
        // Transfer the tokens to the Gateway contract
        IERC20(token).safeTransfer(address(gateway), amount);

        // Forward the call to the Gateway contract
        gateway.revertWithERC20(token, to, amount, data);

        emit WithdrawAndRevert(token, to, amount, data);
    }
}