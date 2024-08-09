// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "./ZetaConnectorBase.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

/// @title ZetaConnectorNative
/// @notice Implementation of ZetaConnectorBase for native token handling.
/// @dev This contract directly transfers Zeta tokens and interacts with the Gateway contract.
contract ZetaConnectorNative is ZetaConnectorBase {
    using SafeERC20 for IERC20;

    constructor(
        address gateway_,
        address zetaToken_,
        address tssAddress_,
        address admin_
    )
        ZetaConnectorBase(gateway_, zetaToken_, tssAddress_, admin_)
    { }

    /// @notice Withdraw tokens to a specified address.
    /// @param to The address to withdraw tokens to.
    /// @param amount The amount of tokens to withdraw.
    /// @param internalSendHash A hash used for internal tracking of the transaction.
    /// @dev This function can only be called by the TSS address.
    function withdraw(
        address to,
        uint256 amount,
        bytes32 internalSendHash
    )
        external
        override
        nonReentrant
        onlyRole(WITHDRAWER_ROLE)
        whenNotPaused
    {
        IERC20(zetaToken).safeTransfer(to, amount);
        emit Withdrawn(to, amount);
    }

    /// @notice Withdraw tokens and call a contract through Gateway.
    /// @param to The address to withdraw tokens to.
    /// @param amount The amount of tokens to withdraw.
    /// @param data The calldata to pass to the contract call.
    /// @param internalSendHash A hash used for internal tracking of the transaction.
    /// @dev This function can only be called by the TSS address.
    function withdrawAndCall(
        address to,
        uint256 amount,
        bytes calldata data,
        bytes32 internalSendHash
    )
        external
        override
        nonReentrant
        onlyRole(WITHDRAWER_ROLE)
        whenNotPaused
    {
        // Transfer zetaToken to the Gateway contract
        IERC20(zetaToken).safeTransfer(address(gateway), amount);

        // Forward the call to the Gateway contract
        gateway.executeWithERC20(address(zetaToken), to, amount, data);

        emit WithdrawnAndCalled(to, amount, data);
    }

    /// @notice Withdraw tokens and call a contract with a revert callback through Gateway.
    /// @param to The address to withdraw tokens to.
    /// @param amount The amount of tokens to withdraw.
    /// @param data The calldata to pass to the contract call.
    /// @param internalSendHash A hash used for internal tracking of the transaction.
    /// @dev This function can only be called by the TSS address.
    function withdrawAndRevert(
        address to,
        uint256 amount,
        bytes calldata data,
        bytes32 internalSendHash
    )
        external
        override
        nonReentrant
        onlyRole(WITHDRAWER_ROLE)
        whenNotPaused
    {
        // Transfer zetaToken to the Gateway contract
        IERC20(zetaToken).safeTransfer(address(gateway), amount);

        // Forward the call to the Gateway contract
        gateway.revertWithERC20(address(zetaToken), to, amount, data);

        emit WithdrawnAndReverted(to, amount, data);
    }

    /// @notice Handle received tokens.
    /// @param amount The amount of tokens received.
    function receiveTokens(uint256 amount) external override whenNotPaused {
        IERC20(zetaToken).safeTransferFrom(msg.sender, address(this), amount);
    }
}
