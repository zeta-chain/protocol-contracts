// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./interfaces.sol";

contract Sender {
    event Call(address indexed sender, bytes indexed receiver, bytes message);

    address public gateway;

    constructor(address _gateway) {
        gateway = _gateway;
    }

    // Call receiver on EVM
    function sendToReceiver(bytes memory receiver, string memory str, uint256 num, bool flag) external {
        // Encode the function call to the receiver's receiveA method
        bytes memory message = abi.encodeWithSignature("receiveA(string,uint256,bool)", str, num, flag);

        // Pass encoded call to gateway
        IGatewayZEVM(gateway).call(receiver, message);
    }
}