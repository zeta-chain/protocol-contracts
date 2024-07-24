// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "src/zevm/interfaces/zContract.sol";

// @notice This contract is used just for testing
contract TestZContract is UniversalContract {
    event ContextData(bytes origin, address sender, uint256 chainID, address msgSender, string message);
    event ContextDataRevert(bytes origin, address sender, uint256 chainID, address msgSender, string message);

    function onCrossChainCall(
        zContext calldata context,
        address zrc20,
        uint256 amount,
        bytes calldata message
    ) external override {
        string memory decodedMessage;
        if (message.length > 0) {
            decodedMessage = abi.decode(message, (string));
        }
        emit ContextData(context.origin, context.sender, context.chainID, msg.sender, decodedMessage);
    }

    function onRevert(
        revertContext calldata context,
        address zrc20,
        uint256 amount,
        bytes calldata message
    ) external override {
        string memory decodedMessage;
        if (message.length > 0) {
            decodedMessage = abi.decode(message, (string));
        }
        emit ContextDataRevert(context.origin, context.sender, context.chainID, msg.sender, decodedMessage);
    }

    receive() external payable {}
    fallback() external payable {}
}