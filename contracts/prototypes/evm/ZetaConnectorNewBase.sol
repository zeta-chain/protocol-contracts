// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./IGatewayEVM.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";

abstract contract ZetaConnectorNewBase is ReentrancyGuard {
    using SafeERC20 for IERC20;

    error ZeroAddress();

    IGatewayEVM public immutable gateway;
    address public immutable zetaToken;

    event Withdraw(address indexed to, uint256 amount);
    event WithdrawAndCall(address indexed to, uint256 amount, bytes data);

    constructor(address _gateway, address _zetaToken) {
        if (_gateway == address(0) || _zetaToken == address(0)) {
            revert ZeroAddress();
        }
        gateway = IGatewayEVM(_gateway);
        zetaToken = _zetaToken;
    }

    function withdraw(address to, uint256 amount, bytes32 internalSendHash) external virtual;

    function withdrawAndCall(address to, uint256 amount, bytes calldata data, bytes32 internalSendHash) external virtual;

    function receiveTokens(uint256 amount) external virtual;
}