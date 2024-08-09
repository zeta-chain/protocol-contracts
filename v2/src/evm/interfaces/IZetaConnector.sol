// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

/// @title IZetaConnectorEvents
/// @notice Interface for the events emitted by the ZetaConnector contracts.
interface IZetaConnectorEvents {
    /// @notice Emitted when tokens are withdrawn.
    /// @param to The address to which the tokens are withdrawn.
    /// @param amount The amount of tokens withdrawn.
    event Withdrawn(address indexed to, uint256 amount);

    /// @notice Emitted when tokens are withdrawn and a contract is called.
    /// @param to The address to which the tokens are withdrawn.
    /// @param amount The amount of tokens withdrawn.
    /// @param data The calldata passed to the contract call.
    event WithdrawnAndCalled(address indexed to, uint256 amount, bytes data);

    /// @notice Emitted when tokens are withdrawn and a contract is called with a revert callback.
    /// @param to The address to which the tokens are withdrawn.
    /// @param amount The amount of tokens withdrawn.
    /// @param data The calldata passed to the contract call.
    event WithdrawnAndReverted(address indexed to, uint256 amount, bytes data);
}
