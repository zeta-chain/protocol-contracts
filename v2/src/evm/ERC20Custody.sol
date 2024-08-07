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
contract ERC20Custody is IERC20Custody, ReentrancyGuard, AccessControl, Pausable {
    using SafeERC20 for IERC20;

    /// @notice Gateway contract.
    IGatewayEVM public immutable gateway;
    /// @notice Mapping of whitelisted tokens => true/false.
    mapping(address => bool) public whitelisted;
    /// @notice New role identifier for pauser role.
    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");
    /// @notice New role identifier for withdrawer role.
    bytes32 public constant WITHDRAWER_ROLE = keccak256("WITHDRAWER_ROLE");
    /// @notice New role identifier for whitelister role.
    bytes32 public constant WHITELISTER_ROLE = keccak256("WHITELISTER_ROLE");

    /// @notice Constructor for ERC20Custody.
    /// @dev Set admin as default admin and pauser, and tssAddress as tss role.
    constructor(address gateway_, address tssAddress_, address admin_) {
        if (gateway_ == address(0) || tssAddress_ == address(0) || admin_ == address(0)) {
            revert ZeroAddress();
        }
        gateway = IGatewayEVM(gateway_);
        _grantRole(DEFAULT_ADMIN_ROLE, admin_);
        _grantRole(PAUSER_ROLE, admin_);
        _grantRole(WITHDRAWER_ROLE, tssAddress_);
        _grantRole(WHITELISTER_ROLE, admin_);
        _grantRole(WHITELISTER_ROLE, tssAddress_);
    }

    /// @notice Pause contract.
    function pause() external onlyRole(PAUSER_ROLE) {
        _pause();
    }

    /// @notice Unpause contract.
    function unpause() external onlyRole(PAUSER_ROLE) {
        _unpause();
    }

    /// @notice Whitelist ERC20 token.
    /// @param token address of ERC20 token
    function whitelist(address token) external onlyRole(WHITELISTER_ROLE) {
        whitelisted[token] = true;
        emit Whitelisted(token);
    }

    /// @notice Unwhitelist ERC20 token.
    /// @param token address of ERC20 token
    function unwhitelist(address token) external onlyRole(WHITELISTER_ROLE) {
        whitelisted[token] = false;
        emit Unwhitelisted(token);
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
        if (!whitelisted[token]) revert NotWhitelisted();

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
        if (!whitelisted[token]) revert NotWhitelisted();

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
        if (!whitelisted[token]) revert NotWhitelisted();

        // Transfer the tokens to the Gateway contract
        IERC20(token).safeTransfer(address(gateway), amount);

        // Forward the call to the Gateway contract
        gateway.revertWithERC20(token, to, amount, data);

        emit WithdrawAndRevert(token, to, amount, data);
    }
}
