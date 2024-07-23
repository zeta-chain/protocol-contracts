// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import "./IReceiverEVM.sol";

// @notice This contract is used just for testing
contract ReceiverEVM is IReceiverEVMEvents, ReentrancyGuard {
    using SafeERC20 for IERC20;
    error ZeroAmount();

    // Payable function
    function receivePayable(string memory str, uint256 num, bool flag) external payable {
        emit ReceivedPayable(msg.sender, msg.value, str, num, flag);
    }

    // Non-payable function
    function receiveNonPayable(string[] memory strs, uint256[] memory nums, bool flag) external {
        emit ReceivedNonPayable(msg.sender, strs, nums, flag);
    }

    // Function using IERC20
    function receiveERC20(uint256 amount, address token, address destination) external nonReentrant {
        // Transfer tokens from the Gateway contract to the destination address
        IERC20(token).safeTransferFrom(msg.sender, destination, amount);

        emit ReceivedERC20(msg.sender, amount, token, destination);
    }

    // Function using IERC20 to partially transfer tokens
    function receiveERC20Partial(uint256 amount, address token, address destination) external nonReentrant {
        uint256 amountToSend = amount / 2;
        if (amountToSend == 0) revert ZeroAmount();

        IERC20(token).safeTransferFrom(msg.sender, destination, amountToSend);

        emit ReceivedERC20(msg.sender, amountToSend, token, destination);
    }

    // Function without parameters
    function receiveNoParams() external {
        emit ReceivedNoParams(msg.sender);
    }

    // onRevertCallback
    function onRevert(bytes calldata data) external {
        emit ReceivedRevert(msg.sender, data);
    }

    receive() external payable {}
    fallback() external payable {}
}