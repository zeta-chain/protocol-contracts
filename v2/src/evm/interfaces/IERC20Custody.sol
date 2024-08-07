// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

/// @title IERC20CustodyEvents
/// @notice Interface for the events emitted by the ERC20 custody contract.
interface IERC20CustodyEvents {
    /// @notice Emitted when tokens are withdrawn.
    /// @param token The address of the ERC20 token.
    /// @param to The address receiving the tokens.
    /// @param amount The amount of tokens withdrawn.
    event Withdraw(address indexed token, address indexed to, uint256 amount);

    /// @notice Emitted when tokens are withdrawn and a contract call is made.
    /// @param token The address of the ERC20 token.
    /// @param to The address receiving the tokens.
    /// @param amount The amount of tokens withdrawn.
    /// @param data The calldata passed to the contract call.
    event WithdrawAndCall(address indexed token, address indexed to, uint256 amount, bytes data);

    /// @notice Emitted when tokens are withdrawn and a revertable contract call is made.
    /// @param token The address of the ERC20 token.
    /// @param to The address receiving the tokens.
    /// @param amount The amount of tokens withdrawn.
    /// @param data The calldata passed to the contract call.
    event WithdrawAndRevert(address indexed token, address indexed to, uint256 amount, bytes data);

    /// @notice Emitted when ERC20 token is whitelisted
    /// @param token address of ERC20 token.
    event Whitelisted(address indexed token);

    /// @notice Emitted when ERC20 token is unwhitelisted
    /// @param token address of ERC20 token.
    event Unwhitelisted(address indexed token);
}

/// @title IERC20CustodyErrors
/// @notice Interface for the errors used in the ERC20 custody contract.
interface IERC20CustodyErrors {
    /// @notice Error for zero address input.
    error ZeroAddress();
    /// @notice Error for not whitelisted ERC20 token
    error NotWhitelisted();
}
