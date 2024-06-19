// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./ERC20CustodyNew.sol";

contract Gateway {
    ERC20CustodyNew public custody;

    event Executed(address indexed destination, uint256 value, bytes data);
    event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data);

    function _execute(address destination, bytes calldata data) internal returns (bytes memory) {
        (bool success, bytes memory result) = destination.call{value: msg.value}(data);
        require(success, "Call failed");
        return result;
    }

    function execute(address destination, bytes calldata data) external payable returns (bytes memory) {
        bytes memory result = _execute(destination, data);

        emit Executed(destination, msg.value, data);

        return result;
    }

    function executeWithERC20(
        address token,
        address to,
        uint256 amount,
        bytes calldata data
    ) external returns (bytes memory) {
        // Approve the target contract to spend the tokens
        IERC20(token).approve(to, amount);

        // Execute the call on the target contract
        bytes memory result = _execute(to, data);

        // Reset approval
        IERC20(token).approve(to, 0);

        // Transfer any remaining tokens back to the custody contract
        uint256 remainingBalance = IERC20(token).balanceOf(address(this));
        if (remainingBalance > 0) {
            IERC20(token).transfer(address(custody), remainingBalance);
        }

        emit ExecutedWithERC20(token, to, amount, data);

        return result;
    }

    function setCustody(address _custody) external {
        custody = ERC20CustodyNew(_custody);
    }
}
