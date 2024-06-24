// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./interfaces.sol";

// As the current version, ERC20CustodyNew hold the ERC20s deposited on ZetaChain
// This version include a functionality allowing to call a contract
// ERC20Custody doesn't call smart contract directly, it passes through the Gateway contract
contract ERC20CustodyNew {
    IGatewayEVM public gateway;

    event Withdraw(address indexed token, address indexed to, uint256 amount);
    event WithdrawAndCall(address indexed token, address indexed to, uint256 amount, bytes data);

    constructor(address _gateway) {
        gateway = IGatewayEVM(_gateway);
    }
    
    // Withdraw is called by TSS address, it directly transfers the tokens to the destination address without contract call
    function withdraw(address token, address to, uint256 amount) external {
        IERC20(token).transfer(to, amount);

        emit Withdraw(token, to, amount);
    }

    // WithdrawAndCall is called by TSS address, it transfers the tokens and call a contract
    // For this, it passes through the Gateway contract, it transfers the tokens to the Gateway contract and then calls the contract
    function withdrawAndCall(address token, address to, uint256 amount, bytes calldata data) external {
        // Transfer the tokens to the Gateway contract
        IERC20(token).transfer(address(gateway), amount);

        // Forward the call to the Gateway contract
        gateway.executeWithERC20(token, to, amount, data);

        emit WithdrawAndCall(token, to, amount, data);
    }
}