// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

contract Receiver {
    event ReceivedA(address sender, uint256 value, string str, uint256 num, bool flag);
    event ReceivedB(address sender, uint256 value, string[] strs, uint256[] nums, bool flag);

    function receiveA(string memory str, uint256 num, bool flag) external payable {
        emit ReceivedA(msg.sender, msg.value, str, num, flag);
    }

    function receiveB(string[] memory strs, uint256[] memory nums, bool flag) external payable {
        emit ReceivedB(msg.sender, msg.value, strs, nums, flag);
    }
}