// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

/// @notice Struct containing revert options
/// @param revertAddress Address to receive revert.
/// @param callOnRevert Flag if onRevert hook should be called.
/// @pararm abortAddress Address to receive funds if aborted.
/// @param revertMessage Arbitrary data sent back in onRevert.
struct RevertOptions {
    address revertAddress;
    bool callOnRevert;
    address abortAddress;
    bytes revertMessage;
}
