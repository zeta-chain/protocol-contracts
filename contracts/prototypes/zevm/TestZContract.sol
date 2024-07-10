// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "../../zevm/interfaces/zContract.sol";

contract TestZContract is UniversalContract {
    event ContextData(bytes origin, address sender, uint256 chainID, address msgSender, string message);

    function onCrossChainCall(
        zContext calldata context,
        address zrc20,
        uint256 amount,
        bytes calldata message
    ) external override {
        string memory decodedMessage = abi.decode(message, (string));
        emit ContextData(context.origin, context.sender, context.chainID, msg.sender, decodedMessage);
    }

    function onRevert(
        revertContext calldata context,
        address zrc20,
        uint256 amount,
        bytes calldata message
    ) external override {
        string memory decodedMessage = abi.decode(message, (string));
        emit ContextData(context.origin, context.sender, context.chainID, msg.sender, decodedMessage);
    }
}