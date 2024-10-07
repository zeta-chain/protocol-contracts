// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import { RevertContext, RevertOptions, Revertable } from "../../contracts/Revert.sol";
import { ZetaConnectorBase } from "./ZetaConnectorBase.sol";
import { IERC20Custody } from "./interfaces/IERC20Custody.sol";
import { IGatewayEVM } from "./interfaces/IGatewayEVM.sol";

import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/ReentrancyGuardUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

/// @title GatewayEVM
/// @notice The GatewayEVM contract is the endpoint to call smart contracts on external chains.
/// @dev The contract doesn't hold any funds and should never have active allowances.
contract GatewayEVM is
    Initializable,
    AccessControlUpgradeable,
    UUPSUpgradeable,
    IGatewayEVM,
    ReentrancyGuardUpgradeable,
    PausableUpgradeable
{
    using SafeERC20 for IERC20;

    /// @notice The address of the custody contract.
    address public custody;
    /// @notice The address of the TSS (Threshold Signature Scheme) contract.
    address public tssAddress;
    /// @notice The address of the ZetaConnector contract.
    address public zetaConnector;
    /// @notice The address of the Zeta token contract.
    address public zetaToken;

    /// @notice New role identifier for tss role.
    bytes32 public constant TSS_ROLE = keccak256("TSS_ROLE");
    /// @notice New role identifier for asset handler role.
    bytes32 public constant ASSET_HANDLER_ROLE = keccak256("ASSET_HANDLER_ROLE");
    /// @notice New role identifier for pauser role.
    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");
    /// @notice Max size of payload + revertOptions revert message.
    uint256 public constant MAX_PAYLOAD_SIZE = 1024;

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    /// @notice Initialize with tss address. address of zeta token and admin account set as DEFAULT_ADMIN_ROLE.
    /// @dev Using admin to authorize upgrades and pause, and tss for tss role.
    function initialize(address tssAddress_, address zetaToken_, address admin_) public initializer {
        if (tssAddress_ == address(0) || zetaToken_ == address(0)) {
            revert ZeroAddress();
        }
        __UUPSUpgradeable_init();
        __ReentrancyGuard_init();
        __AccessControl_init();
        __Pausable_init();

        _grantRole(DEFAULT_ADMIN_ROLE, admin_);
        _grantRole(PAUSER_ROLE, admin_);
        tssAddress = tssAddress_;
        _grantRole(TSS_ROLE, tssAddress_);

        zetaToken = zetaToken_;
    }

    /// @dev Authorizes the upgrade of the contract, sender must be owner.
    /// @param newImplementation Address of the new implementation.
    function _authorizeUpgrade(address newImplementation) internal override onlyRole(DEFAULT_ADMIN_ROLE) { }

    /// @dev Internal function to execute a call to a destination address.
    /// @param destination Address to call.
    /// @param data Calldata to pass to the call.
    /// @return The result of the call.
    function _execute(address destination, bytes calldata data) internal returns (bytes memory) {
        revertIfCallingOnRevert(data);
        (bool success, bytes memory result) = destination.call{ value: msg.value }(data);
        if (!success) revert ExecutionFailed();

        return result;
    }

    /// @notice Update tss address
    /// @param newTSSAddress new tss address
    function updateTSSAddress(address newTSSAddress) external onlyRole(DEFAULT_ADMIN_ROLE) {
        if (newTSSAddress == address(0)) revert ZeroAddress();

        _revokeRole(TSS_ROLE, tssAddress);
        _grantRole(TSS_ROLE, newTSSAddress);

        tssAddress = newTSSAddress;

        emit UpdatedGatewayTSSAddress(newTSSAddress);
    }

    /// @notice Pause contract.
    function pause() external onlyRole(PAUSER_ROLE) {
        _pause();
    }

    /// @notice Unpause contract.
    function unpause() external onlyRole(PAUSER_ROLE) {
        _unpause();
    }

    /// @notice Transfers msg.value to destination contract and executes it's onRevert function.
    /// @dev This function can only be called by the TSS address and it is payable.
    /// @param destination Address to call.
    /// @param data Calldata to pass to the call.
    function executeRevert(
        address destination,
        bytes calldata data,
        RevertContext calldata revertContext
    )
        public
        payable
        onlyRole(TSS_ROLE)
        whenNotPaused
        nonReentrant
    {
        if (destination == address(0)) revert ZeroAddress();
        (bool success,) = destination.call{ value: msg.value }("");
        if (!success) revert ExecutionFailed();
        Revertable(destination).onRevert(revertContext);

        emit Reverted(destination, address(0), msg.value, data, revertContext);
    }

    /// @notice Executes a call to a destination address without ERC20 tokens.
    /// @dev This function can only be called by the TSS address and it is payable.
    /// @param destination Address to call.
    /// @param data Calldata to pass to the call.
    /// @return The result of the call.
    function execute(
        address destination,
        bytes calldata data
    )
        external
        payable
        onlyRole(TSS_ROLE)
        whenNotPaused
        nonReentrant
        returns (bytes memory)
    {
        if (destination == address(0)) revert ZeroAddress();
        bytes memory result = _execute(destination, data);

        emit Executed(destination, msg.value, data);

        return result;
    }

    /// @notice Executes a call to a destination contract using ERC20 tokens.
    /// @dev This function can only be called by the custody or connector address.
    ///      It uses the ERC20 allowance system, resetting gateway allowance at the end.
    /// @param token Address of the ERC20 token.
    /// @param to Address of the contract to call.
    /// @param amount Amount of tokens to transfer.
    /// @param data Calldata to pass to the call.
    function executeWithERC20(
        address token,
        address to,
        uint256 amount,
        bytes calldata data
    )
        public
        onlyRole(ASSET_HANDLER_ROLE)
        whenNotPaused
        nonReentrant
    {
        if (amount == 0) revert InsufficientERC20Amount();
        if (to == address(0)) revert ZeroAddress();
        // Approve the target contract to spend the tokens
        if (!resetApproval(token, to)) revert ApprovalFailed();
        if (!IERC20(token).approve(to, amount)) revert ApprovalFailed();
        // Execute the call on the target contract
        _execute(to, data);

        // Reset approval
        if (!resetApproval(token, to)) revert ApprovalFailed();

        // Transfer any remaining tokens back to the custody/connector contract
        uint256 remainingBalance = IERC20(token).balanceOf(address(this));
        if (remainingBalance > 0) {
            transferToAssetHandler(token, remainingBalance);
        }

        emit ExecutedWithERC20(token, to, amount, data);
    }

    /// @notice Directly transfers ERC20 tokens and calls onRevert.
    /// @dev This function can only be called by the custody or connector address.
    /// @param token Address of the ERC20 token.
    /// @param to Address of the contract to call.
    /// @param amount Amount of tokens to transfer.
    /// @param data Calldata to pass to the call.
    /// @param revertContext Revert context to pass to onRevert.
    function revertWithERC20(
        address token,
        address to,
        uint256 amount,
        bytes calldata data,
        RevertContext calldata revertContext
    )
        external
        onlyRole(ASSET_HANDLER_ROLE)
        whenNotPaused
        nonReentrant
    {
        if (amount == 0) revert InsufficientERC20Amount();
        if (to == address(0)) revert ZeroAddress();

        IERC20(token).safeTransfer(address(to), amount);
        Revertable(to).onRevert(revertContext);

        emit Reverted(to, token, amount, data, revertContext);
    }

    /// @notice Deposits ETH to the TSS address.
    /// @param receiver Address of the receiver.
    /// @param revertOptions Revert options.
    function deposit(
        address receiver,
        RevertOptions calldata revertOptions
    )
        external
        payable
        whenNotPaused
        nonReentrant
    {
        if (msg.value == 0) revert InsufficientETHAmount();
        if (receiver == address(0)) revert ZeroAddress();

        (bool deposited,) = tssAddress.call{ value: msg.value }("");

        if (!deposited) revert DepositFailed();

        emit Deposited(msg.sender, receiver, msg.value, address(0), "", revertOptions);
    }

    /// @notice Deposits ERC20 tokens to the custody or connector contract.
    /// @param receiver Address of the receiver.
    /// @param amount Amount of tokens to deposit.
    /// @param asset Address of the ERC20 token.
    /// @param revertOptions Revert options.
    function deposit(
        address receiver,
        uint256 amount,
        address asset,
        RevertOptions calldata revertOptions
    )
        external
        whenNotPaused
        nonReentrant
    {
        if (amount == 0) revert InsufficientERC20Amount();
        if (receiver == address(0)) revert ZeroAddress();

        transferFromToAssetHandler(msg.sender, asset, amount);

        emit Deposited(msg.sender, receiver, amount, asset, "", revertOptions);
    }

    /// @notice Deposits ETH to the TSS address and calls an omnichain smart contract.
    /// @param receiver Address of the receiver.
    /// @param payload Calldata to pass to the call.
    /// @param revertOptions Revert options.
    function depositAndCall(
        address receiver,
        bytes calldata payload,
        RevertOptions calldata revertOptions
    )
        external
        payable
        whenNotPaused
        nonReentrant
    {
        if (msg.value == 0) revert InsufficientETHAmount();
        if (receiver == address(0)) revert ZeroAddress();
        if (payload.length + revertOptions.revertMessage.length >= MAX_PAYLOAD_SIZE) revert PayloadSizeExceeded();

        (bool deposited,) = tssAddress.call{ value: msg.value }("");

        if (!deposited) revert DepositFailed();

        emit Deposited(msg.sender, receiver, msg.value, address(0), payload, revertOptions);
    }

    /// @notice Deposits ERC20 tokens to the custody or connector contract and calls an omnichain smart contract.
    /// @param receiver Address of the receiver.
    /// @param amount Amount of tokens to deposit.
    /// @param asset Address of the ERC20 token.
    /// @param payload Calldata to pass to the call.
    /// @param revertOptions Revert options.
    function depositAndCall(
        address receiver,
        uint256 amount,
        address asset,
        bytes calldata payload,
        RevertOptions calldata revertOptions
    )
        external
        whenNotPaused
        nonReentrant
    {
        if (amount == 0) revert InsufficientERC20Amount();
        if (receiver == address(0)) revert ZeroAddress();
        if (payload.length + revertOptions.revertMessage.length >= MAX_PAYLOAD_SIZE) revert PayloadSizeExceeded();

        transferFromToAssetHandler(msg.sender, asset, amount);

        emit Deposited(msg.sender, receiver, amount, asset, payload, revertOptions);
    }

    /// @notice Calls an omnichain smart contract without asset transfer.
    /// @param receiver Address of the receiver.
    /// @param payload Calldata to pass to the call.
    /// @param revertOptions Revert options.
    function call(
        address receiver,
        bytes calldata payload,
        RevertOptions calldata revertOptions
    )
        external
        whenNotPaused
        nonReentrant
    {
        if (receiver == address(0)) revert ZeroAddress();
        if (payload.length + revertOptions.revertMessage.length >= MAX_PAYLOAD_SIZE) revert PayloadSizeExceeded();

        emit Called(msg.sender, receiver, payload, revertOptions);
    }

    /// @notice Sets the custody contract address.
    /// @param custody_ Address of the custody contract.
    function setCustody(address custody_) external onlyRole(DEFAULT_ADMIN_ROLE) {
        if (custody_ == address(0)) revert ZeroAddress();
        if (custody != address(0)) revert CustodyInitialized();

        _grantRole(ASSET_HANDLER_ROLE, custody_);
        custody = custody_;
    }

    /// @notice Sets the connector contract address.
    /// @param zetaConnector_ Address of the connector contract.
    function setConnector(address zetaConnector_) external onlyRole(DEFAULT_ADMIN_ROLE) {
        if (zetaConnector_ == address(0)) revert ZeroAddress();
        if (zetaConnector != address(0)) revert ConnectorInitialized();

        _grantRole(ASSET_HANDLER_ROLE, zetaConnector_);
        zetaConnector = zetaConnector_;
    }

    /// @dev Resets the approval of a token for a specified address.
    /// This is used to ensure that the approval is set to zero before setting it to a new value.
    /// @param token Address of the ERC20 token.
    /// @param to Address to reset the approval for.
    /// @return True if the approval reset was successful, false otherwise.
    function resetApproval(address token, address to) private returns (bool) {
        return IERC20(token).approve(to, 0);
    }

    /// @dev Transfers tokens from the sender to the asset handler.
    /// This function handles the transfer of tokens to either the connector or custody contract based on the asset
    /// type.
    /// @param from Address of the sender.
    /// @param token Address of the ERC20 token.
    /// @param amount Amount of tokens to transfer.
    function transferFromToAssetHandler(address from, address token, uint256 amount) private {
        if (token == zetaToken) {
            // transfer to connector
            // transfer amount to gateway
            IERC20(token).safeTransferFrom(from, address(this), amount);
            // approve connector to handle tokens depending on connector version (eg. lock or burn)
            if (!IERC20(token).approve(zetaConnector, amount)) revert ApprovalFailed();
            // send tokens to connector
            ZetaConnectorBase(zetaConnector).receiveTokens(amount);
        } else {
            // transfer to custody
            if (!IERC20Custody(custody).whitelisted(token)) revert NotWhitelistedInCustody();
            IERC20(token).safeTransferFrom(from, custody, amount);
        }
    }

    /// @dev Transfers tokens to the asset handler.
    /// This function handles the transfer of tokens to either the connector or custody contract based on the asset
    /// type.
    /// @param token Address of the ERC20 token.
    /// @param amount Amount of tokens to transfer.
    function transferToAssetHandler(address token, uint256 amount) private {
        if (token == zetaToken) {
            // transfer to connector
            // approve connector to handle tokens depending on connector version (eg. lock or burn)
            if (!IERC20(token).approve(zetaConnector, amount)) revert ApprovalFailed();
            // send tokens to connector
            ZetaConnectorBase(zetaConnector).receiveTokens(amount);
        } else {
            // transfer to custody
            if (!IERC20Custody(custody).whitelisted(token)) revert NotWhitelistedInCustody();
            IERC20(token).safeTransfer(custody, amount);
        }
    }

    // @dev prevent spoofing onRevert functions
    function revertIfCallingOnRevert(bytes calldata data) private pure {
        if (data.length >= 4) {
            bytes4 functionSelector;
            assembly {
                functionSelector := calldataload(data.offset)
            }

            if (functionSelector == Revertable.onRevert.selector) {
                revert NotAllowedToCallOnRevert();
            }
        }
    }
}
