// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

contract Gateway {
    event Forwarded(address indexed destination, uint256 value, bytes data);

    function forwardCall(address destination, bytes calldata data) external payable returns (bytes memory) {
        (bool success, bytes memory result) = destination.call{value: msg.value}(data);
        
        require(success, "Call failed");
        
        emit Forwarded(destination, msg.value, data);

        return result;
    }
}