// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

import "./ZetaConnectorNewBase.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

contract ZetaConnectorNative is ZetaConnectorNewBase {
    using SafeERC20 for IERC20;

    constructor(address _gateway, address _zetaToken)
        ZetaConnectorNewBase(_gateway, _zetaToken)
    {}

    // @dev withdraw is called by TSS address, it directly transfers zetaToken to the destination address without contract call
    function withdraw(address to, uint256 amount, bytes32 internalSendHash) external override nonReentrant {
        IERC20(zetaToken).safeTransfer(to, amount);
        emit Withdraw(to, amount);
    }

    // @dev withdrawAndCall is called by TSS address, it transfers zetaToken to the gateway and calls a contract
    function withdrawAndCall(address to, uint256 amount, bytes calldata data, bytes32 internalSendHash) external override nonReentrant {
        // Transfer zetaToken to the Gateway contract
        IERC20(zetaToken).safeTransfer(address(gateway), amount);

        // Forward the call to the Gateway contract
        gateway.executeWithERC20(address(zetaToken), to, amount, data);

        emit WithdrawAndCall(to, amount, data);
    }

    // @dev withdrawAndRevert is called by TSS address, it transfers zetaToken to the gateway and calls onRevert on a contract
    function withdrawAndRevert(address to, uint256 amount, bytes calldata data, bytes32 internalSendHash) external override nonReentrant {
        // Transfer zetaToken to the Gateway contract
        IERC20(zetaToken).safeTransfer(address(gateway), amount);

        // Forward the call to the Gateway contract
        gateway.revertWithERC20(address(zetaToken), to, amount, data);

        emit WithdrawAndRevert(to, amount, data);
    }

    // @dev receiveTokens handles token transfer and burn them
    function receiveTokens(uint256 amount) external override {
        IERC20(zetaToken).safeTransferFrom(msg.sender, address(this), amount);
    }
}
