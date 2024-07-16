// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./interfaces.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";

contract ZetaConnectorNew is ReentrancyGuard{
    using SafeERC20 for IERC20;
    error ZeroAddress();

    IGatewayEVM public immutable gateway;
    IERC20 public immutable zeta;

    event Withdraw(address indexed to, uint256 amount);
    event WithdrawAndCall(address indexed to, uint256 amount, bytes data);

    constructor(address _gateway, address _zeta) {
        if (_gateway == address(0) || _zeta == address(0)) {
            revert ZeroAddress();
        }
        gateway = IGatewayEVM(_gateway);
        zeta = IERC20(_zeta);
    }
    
    // Withdraw is called by TSS address, it directly transfers zeta to the destination address without contract call
    // TODO: Finalize access control
    // https://github.com/zeta-chain/protocol-contracts/issues/204
    function withdraw(address to, uint256 amount) external nonReentrant {
        // TODO: mint?
        zeta.safeTransfer(to, amount);

        emit Withdraw(to, amount);
    }

    // WithdrawAndCall is called by TSS address, it transfers zeta and call a contract
    // For this, it passes through the Gateway contract, it transfers zeta to the Gateway contract and then calls the contract
    // TODO: Finalize access control
    // https://github.com/zeta-chain/protocol-contracts/issues/204
    function withdrawAndCall(address to, uint256 amount, bytes calldata data) public nonReentrant {
        // TODO: mint?
        // Transfer zeta to the Gateway contract
        zeta.safeTransfer(address(gateway), amount);

        // Forward the call to the Gateway contract
        gateway.executeWithERC20(address(zeta), to, amount, data);

        emit WithdrawAndCall(to, amount, data);
    }
}