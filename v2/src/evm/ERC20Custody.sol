// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "./interfaces//IGatewayEVM.sol";
import "./interfaces/IERC20Custody.sol";

import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/utils/Pausable.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

/// @title ERC20Custody
/// @notice Holds the ERC20 tokens deposited on ZetaChain and includes functionality to call a contract.
/// @dev This contract does not call smart contracts directly, it passes through the Gateway contract.
contract ERC20Custody is IERC20CustodyEvents, IERC20CustodyErrors, ReentrancyGuard, AccessControl, Pausable {
    using SafeERC20 for IERC20;

    /// @notice Gateway contract.
    IGatewayEVM public immutable gateway;
    /// @notice New role identifier for pauser role.
    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");
    /// @notice New role identifier for withdrawer role.
    bytes32 public constant WITHDRAWER_ROLE = keccak256("WITHDRAWER_ROLE");

    /// @notice Constructor for ERC20Custody.
    /// @dev Set admin as default admin and pauser, and tssAddress as tss role.
    constructor(address _gateway, address _tssAddress, address _admin) {
        if (_gateway == address(0) || _tssAddress == address(0) || _admin == address(0)) {
            revert ZeroAddress();
        }
        gateway = IGatewayEVM(_gateway);
        _grantRole(DEFAULT_ADMIN_ROLE, _admin);
        _grantRole(PAUSER_ROLE, _admin);
        _grantRole(WITHDRAWER_ROLE, _tssAddress);
    }

    /// @notice Pause contract.
    function pause() external onlyRole(PAUSER_ROLE) {
        _pause();
    }

    /// @notice Unpause contract.
    function unpause() external onlyRole(PAUSER_ROLE) {
        _unpause();
    }

    /// @notice Withdraw directly transfers the tokens to the destination address without contract call.
    /// @dev This function can only be called by the TSS address.
    /// @param token Address of the ERC20 token.
    /// @param to Destination address for the tokens.
    /// @param amount Amount of tokens to withdraw.
    function withdraw(
        address token,
        address to,
        uint256 amount
    )
        external
        nonReentrant
        onlyRole(WITHDRAWER_ROLE)
        whenNotPaused
    {
        IERC20(token).safeTransfer(to, amount);

        emit Withdraw(token, to, amount);
    }

    /// @notice WithdrawAndCall transfers tokens to Gateway and call a contract through the Gateway.
    /// @dev This function can only be called by the TSS address.
    /// @param token Address of the ERC20 token.
    /// @param to Address of the contract to call.
    /// @param amount Amount of tokens to withdraw.
    /// @param data Calldata to pass to the contract call.
    function withdrawAndCall(
        address token,
        address to,
        uint256 amount,
        bytes calldata data
    )
        public
        nonReentrant
        onlyRole(WITHDRAWER_ROLE)
        whenNotPaused
    {
        // Transfer the tokens to the Gateway contract
        IERC20(token).safeTransfer(address(gateway), amount);

        // Forward the call to the Gateway contract
        gateway.executeWithERC20(token, to, amount, data);

        emit WithdrawAndCall(token, to, amount, data);
    }

    /// @notice WithdrawAndRevert transfers tokens to Gateway and call a contract with a revert functionality through
    /// the Gateway.
    /// @dev This function can only be called by the TSS address.
    /// @param token Address of the ERC20 token.
    /// @param to Address of the contract to call.
    /// @param amount Amount of tokens to withdraw.
    /// @param data Calldata to pass to the contract call.
    function withdrawAndRevert(
        address token,
        address to,
        uint256 amount,
        bytes calldata data
    )
        public
        nonReentrant
        onlyRole(WITHDRAWER_ROLE)
        whenNotPaused
    {
        // Transfer the tokens to the Gateway contract
        IERC20(token).safeTransfer(address(gateway), amount);

        // Forward the call to the Gateway contract
        gateway.revertWithERC20(token, to, amount, data);

        emit WithdrawAndRevert(token, to, amount, data);
    }
}
