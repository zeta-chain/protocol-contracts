// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract Receiver {
    event ReceivedA(address sender, uint256 value, string str, uint256 num, bool flag);
    event ReceivedB(address sender, string[] strs, uint256[] nums, bool flag);
    event ReceivedC(address sender, uint256 amount, address token, address destination);

    function receiveA(string memory str, uint256 num, bool flag) external payable {
        emit ReceivedA(msg.sender, msg.value, str, num, flag);
    }

    function receiveB(string[] memory strs, uint256[] memory nums, bool flag) external {
        emit ReceivedB(msg.sender, strs, nums, flag);
    }

    function receiveC(uint256 amount, address token, address destination) external {
        // Transfer tokens from the Gateway contract to the destination address
        IERC20(token).transferFrom(msg.sender, destination, amount);

        emit ReceivedC(msg.sender, amount, token, destination);
    }
}