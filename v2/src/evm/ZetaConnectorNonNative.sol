// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "./ZetaConnectorBase.sol";
import "./interfaces/IZetaNonEthNew.sol";

/// @title ZetaConnectorNonNative
/// @notice Implementation of ZetaConnectorBase for non-native token handling.
/// @dev This contract mints and burns Zeta tokens and interacts with the Gateway contract.
contract ZetaConnectorNonNative is ZetaConnectorBase {
    /// @notice Max supply for minting.
    uint256 public maxSupply = type(uint256).max;

    /// @notice Event triggered when max supply is updated.
    /// @param maxSupply New max supply.
    event MaxSupplyUpdated(uint256 maxSupply);

    error ExceedsMaxSupply();

    constructor(
        address _gateway,
        address _zetaToken,
        address _tssAddress,
        address _admin
    )
        ZetaConnectorBase(_gateway, _zetaToken, _tssAddress, _admin)
    { }

    /// @notice Set max supply for minting.
    /// @param _maxSupply New max supply.
    /// @dev This function can only be called by the TSS address.
    function setMaxSupply(uint256 _maxSupply) external onlyRole(WITHDRAWER_ROLE) whenNotPaused {
        maxSupply = _maxSupply;
        emit MaxSupplyUpdated(_maxSupply);
    }

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
        if (amount + IERC20(zetaToken).totalSupply() > maxSupply) revert ExceedsMaxSupply();

        IZetaNonEthNew(zetaToken).mint(to, amount, internalSendHash);
        emit Withdraw(to, amount);
    }

    /// @notice Withdraw tokens and call a contract through Gateway.
    /// @param to The address to withdraw tokens to.
    /// @param amount The amount of tokens to withdraw.
    /// @param data The calldata to pass to the contract call.
    /// @param internalSendHash A hash used for internal tracking of the transaction.
    /// @dev This function can only be called by the TSS address, and mints if supply is not reached.
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
        if (amount + IERC20(zetaToken).totalSupply() > maxSupply) revert ExceedsMaxSupply();

        // Mint zetaToken to the Gateway contract
        IZetaNonEthNew(zetaToken).mint(address(gateway), amount, internalSendHash);

        // Forward the call to the Gateway contract
        gateway.executeWithERC20(address(zetaToken), to, amount, data);

        emit WithdrawAndCall(to, amount, data);
    }

    /// @notice Withdraw tokens and call a contract with a revert callback through Gateway.
    /// @param to The address to withdraw tokens to.
    /// @param amount The amount of tokens to withdraw.
    /// @param data The calldata to pass to the contract call.
    /// @param internalSendHash A hash used for internal tracking of the transaction.
    /// @dev This function can only be called by the TSS address, and mints if supply is not reached.
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
        if (amount + IERC20(zetaToken).totalSupply() > maxSupply) revert ExceedsMaxSupply();

        // Mint zetaToken to the Gateway contract
        IZetaNonEthNew(zetaToken).mint(address(gateway), amount, internalSendHash);

        // Forward the call to the Gateway contract
        gateway.revertWithERC20(address(zetaToken), to, amount, data);

        emit WithdrawAndRevert(to, amount, data);
    }

    /// @notice Handle received tokens and burn them.
    /// @param amount The amount of tokens received.
    function receiveTokens(uint256 amount) external override whenNotPaused {
        IZetaNonEthNew(zetaToken).burnFrom(msg.sender, amount);
    }
}
