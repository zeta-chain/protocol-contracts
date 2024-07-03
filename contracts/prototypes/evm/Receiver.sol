// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

contract Receiver {
    using SafeERC20 for IERC20;

    event ReceivedPayable(address sender, uint256 value, string str, uint256 num, bool flag);
    event ReceivedNonPayable(address sender, string[] strs, uint256[] nums, bool flag);
    event ReceivedERC20(address sender, uint256 amount, address token, address destination);
    event ReceivedNoParams(address sender);

    // Payable function
    function receivePayable(string memory str, uint256 num, bool flag) external payable {
        emit ReceivedPayable(msg.sender, msg.value, str, num, flag);
    }

    // Non-payable function
    function receiveNonPayable(string[] memory strs, uint256[] memory nums, bool flag) external {
        emit ReceivedNonPayable(msg.sender, strs, nums, flag);
    }

    // Function using IERC20
    function receiveERC20(uint256 amount, address token, address destination) external {
        // Transfer tokens from the Gateway contract to the destination address
        IERC20(token).safeTransferFrom(msg.sender, destination, amount);

        emit ReceivedERC20(msg.sender, amount, token, destination);
    }

    // Function without parameters
    function receiveNoParams() external {
        emit ReceivedNoParams(msg.sender);
    }
}