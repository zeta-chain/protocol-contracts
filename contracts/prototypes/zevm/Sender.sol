// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./interfaces.sol";
import "../../zevm/interfaces/IZRC20.sol";

contract Sender {
    event Call(address indexed sender, bytes indexed receiver, bytes message);
    event Withdrawal(address indexed from, bytes to, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message);

    address public gateway;

    constructor(address _gateway) {
        gateway = _gateway;
    }

    // Call receiver on EVM
    function callReceiver(bytes memory receiver, string memory str, uint256 num, bool flag) external {
        // Encode the function call to the receiver's receiveA method
        bytes memory message = abi.encodeWithSignature("receiveA(string,uint256,bool)", str, num, flag);

        // Pass encoded call to gateway
        IGatewayZEVM(gateway).call(receiver, message);
    }

    // Withdraw and call receiver on EVM
    function withdrawAndCallReceiver(bytes memory receiver, uint256 amount, address zrc20, string memory str, uint256 num, bool flag) external {
        // Encode the function call to the receiver's receiveA method
        bytes memory message = abi.encodeWithSignature("receiveA(string,uint256,bool)", str, num, flag);

        // Approve gateway to withdraw
        IZRC20(zrc20).approve(gateway, amount);

        // Pass encoded call to gateway
        IGatewayZEVM(gateway).withdrawAndCall(receiver, amount, zrc20, message);
    }
}