// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./Gateway.sol";

contract ERC20Custody {
    Gateway public gateway;

    event Withdraw(address indexed token, address indexed to, uint256 amount);
    event WithdrawAndCall(address indexed token, address indexed to, uint256 amount, bytes data);

    constructor(address _gateway) {
        gateway = Gateway(_gateway);
    }
    
    function withdraw(address token, address to, uint256 amount) external {
        IERC20(token).transfer(to, amount);

        emit Withdraw(token, to, amount);
    }

    function withdrawAndCall(address token, address to, uint256 amount, bytes calldata data) external {
        // Transfer the tokens to the Gateway contract
        IERC20(token).transfer(address(gateway), amount);

        // Forward the call to the Gateway contract
        gateway.executeWithERC20(token, to, amount, data);

        emit WithdrawAndCall(token, to, amount, data);
    }
}