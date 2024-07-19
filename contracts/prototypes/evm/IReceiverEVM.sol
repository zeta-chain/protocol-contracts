// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IReceiverEVMEvents {
    event ReceivedPayable(address sender, uint256 value, string str, uint256 num, bool flag);
    event ReceivedNonPayable(address sender, string[] strs, uint256[] nums, bool flag);
    event ReceivedERC20(address sender, uint256 amount, address token, address destination);
    event ReceivedNoParams(address sender);
}
