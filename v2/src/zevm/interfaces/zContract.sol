// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

struct zContext {
    bytes origin;
    address sender;
    uint256 chainID;
}

interface zContract {
    function onCrossChainCall(
        zContext calldata context,
        address zrc20,
        uint256 amount,
        bytes calldata message
    )
        external;
}

interface UniversalContract {
    function onCrossChainCall(
        zContext calldata context,
        address zrc20,
        uint256 amount,
        bytes calldata message
    )
        external;

    function onRevert(zContext calldata context, address zrc20, uint256 amount, bytes calldata message) external;
}
