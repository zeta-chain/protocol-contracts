// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import { RevertContext } from "../../../contracts/Revert.sol";

/// @custom:deprecated should be removed once v2 SystemContract is not used anymore.
/// MessageContext should be used
struct zContext {
    bytes origin;
    address sender;
    uint256 chainID;
}

/// @custom:deprecated should be removed once v2 SystemContract is not used anymore.
/// UniversalContract should be used
interface zContract {
    function onCrossChainCall(
        zContext calldata context,
        address zrc20,
        uint256 amount,
        bytes calldata message
    )
        external;
}

/// @notice Provides contextual information when executing a cross-chain call on ZetaChain.
/// @dev This struct helps identify the sender of the message across different blockchain environments.
struct MessageContext {
    /// @notice The address of the sender on the connected chain.
    /// @dev This field uses `bytes` to remain chain-agnostic, allowing support for both EVM and non-EVM chains.
    /// If the connected chain is an EVM chain, `senderEVM` will also be populated with the same value.
    bytes sender;
    /// @notice The sender's address in `address` type if the connected chain is an EVM-compatible chain.
    address senderEVM;
    /// @notice The chain ID of the connected chain.
    /// @dev This identifies the origin chain of the message, allowing contract logic to differentiate between sources.
    uint256 chainID;
}

/// @title UniversalContract
/// @notice Interface for contracts that can receive cross-chain calls on ZetaChain.
/// @dev Contracts implementing this interface can handle incoming cross-chain messages
/// and execute logic based on the provided context, token, and message payload.
interface UniversalContract {
    function onCall(MessageContext calldata context, address zrc20, uint256 amount, bytes calldata message) external;
}
