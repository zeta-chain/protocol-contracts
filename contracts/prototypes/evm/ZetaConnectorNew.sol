// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./IGatewayEVM.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";

contract ZetaConnectorNew is ReentrancyGuard{
    using SafeERC20 for IERC20;
    error ZeroAddress();

    IGatewayEVM public immutable gateway;
    IERC20 public immutable zetaToken;

    event Withdraw(address indexed to, uint256 amount);
    event WithdrawAndCall(address indexed to, uint256 amount, bytes data);

    constructor(address _gateway, address _zetaToken) {
        if (_gateway == address(0) || _zetaToken == address(0)) {
            revert ZeroAddress();
        }
        gateway = IGatewayEVM(_gateway);
        zetaToken = IERC20(_zetaToken);
    }
    
    // Withdraw is called by TSS address, it directly transfers zetaToken to the destination address without contract call
    // TODO: Finalize access control
    // https://github.com/zeta-chain/protocol-contracts/issues/204
    function withdraw(address to, uint256 amount) external nonReentrant {
        // TODO: mint?
        zetaToken.safeTransfer(to, amount);

        emit Withdraw(to, amount);
    }

    // WithdrawAndCall is called by TSS address, it transfers zetaToken and call a contract
    // For this, it passes through the Gateway contract, it transfers zetaToken to the Gateway contract and then calls the contract
    // TODO: Finalize access control
    // https://github.com/zeta-chain/protocol-contracts/issues/204
    function withdrawAndCall(address to, uint256 amount, bytes calldata data) public nonReentrant {
        // TODO: mint?
        // Transfer zetaToken to the Gateway contract
        zetaToken.safeTransfer(address(gateway), amount);

        // Forward the call to the Gateway contract
        gateway.executeWithERC20(address(zetaToken), to, amount, data);

        emit WithdrawAndCall(to, amount, data);
    }
}