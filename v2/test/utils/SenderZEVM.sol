// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "src/zevm/interfaces/IGatewayZEVM.sol";
import "src/zevm/interfaces/IZRC20.sol";

// @notice This contract is used just for testing
contract SenderZEVM {
    address public gateway;
    error ApprovalFailed();

    constructor(address _gateway) {
        gateway = _gateway;
    }

    // Call receiver on EVM
    function callReceiver(bytes memory receiver, string memory str, uint256 num, bool flag) external {
        // Encode the function call to the receiver's receivePayable method
        bytes memory message = abi.encodeWithSignature("receivePayable(string,uint256,bool)", str, num, flag);

        // Pass encoded call to gateway
        IGatewayZEVM(gateway).call(receiver, message);
    }

    // Withdraw and call receiver on EVM
    function withdrawAndCallReceiver(bytes memory receiver, uint256 amount, address zrc20, string memory str, uint256 num, bool flag) external {
        // Encode the function call to the receiver's receivePayable method
        bytes memory message = abi.encodeWithSignature("receivePayable(string,uint256,bool)", str, num, flag);

        // Approve gateway to withdraw
        if(!IZRC20(zrc20).approve(gateway, amount)) revert ApprovalFailed();

        // Pass encoded call to gateway
        IGatewayZEVM(gateway).withdrawAndCall(receiver, amount, zrc20, message);
    }
}