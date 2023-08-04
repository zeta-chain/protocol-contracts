// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

struct Context {
    bytes origin;
    address sender;
    uint256 chainID;
}

interface zContract {
    function onCrossChainCall(Context calldata context, address zrc20, uint256 amount, bytes calldata message) external;
}


