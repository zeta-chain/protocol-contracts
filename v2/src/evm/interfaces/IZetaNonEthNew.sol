// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/**
 * @dev IZetaNonEthNew is a mintable / burnable version of IERC20
 */
interface IZetaNonEthNew is IERC20 {
    function burnFrom(address account, uint256 amount) external;

    function mint(address mintee, uint256 value, bytes32 internalSendHash) external;
}
