// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "./interfaces.sol";

// The GatewayEVM contract is the endpoint to call smart contracts on external chains
// The contract doesn't hold any funds and should never have active allowances
contract GatewayEVM is Initializable, OwnableUpgradeable, UUPSUpgradeable, IGatewayEVMErrors, IGatewayEVMEvents {
    using SafeERC20 for IERC20;

    address public custody;
    address public tssAddress;

    constructor() {}

    function initialize(address _tssAddress) public initializer {
        __Ownable_init();
        __UUPSUpgradeable_init();

        if (_tssAddress == address(0)) revert ZeroAddress();

        tssAddress = _tssAddress;
    }

    function _authorizeUpgrade(address newImplementation) internal override onlyOwner() {}

    function _execute(address destination, bytes calldata data) internal returns (bytes memory) {
        (bool success, bytes memory result) = destination.call{value: msg.value}(data);
    
        if (!success) revert ExecutionFailed();

        return result;
    }

    // Called by the TSS
    // Execution without ERC20 tokens, it is payable and can be used in the case of WithdrawAndCall for Gas ZRC20
    // It can be also used for contract call without asset movement
    function execute(address destination, bytes calldata data) external payable returns (bytes memory) {
        bytes memory result = _execute(destination, data);

        emit Executed(destination, msg.value, data);

        return result;
    }

    // Called by the ERC20Custody contract
    // It call a function using ERC20 transfer
    // Since the goal is to allow calling contract not designed for ZetaChain specifically, it uses ERC20 allowance system
    // It provides allowance to destination contract and call destination contract. In the end, it remove remaining allowance and transfer remaining tokens back to the custody contract for security purposes
    function executeWithERC20(
        address token,
        address to,
        uint256 amount,
        bytes calldata data
    ) public returns (bytes memory) {
        if (amount == 0) revert InsufficientETHAmount();
        // Approve the target contract to spend the tokens
        if(!resetApproval(token, to)) revert ApprovalFailed();
        if(!IERC20(token).approve(to, amount)) revert ApprovalFailed();

        // Execute the call on the target contract
        bytes memory result = _execute(to, data);

        // Reset approval
        if(!resetApproval(token, to)) revert ApprovalFailed();

        // Transfer any remaining tokens back to the custody contract
        uint256 remainingBalance = IERC20(token).balanceOf(address(this));
        if (remainingBalance > 0) {
            IERC20(token).safeTransfer(address(custody), remainingBalance);
        }

        emit ExecutedWithERC20(token, to, amount, data);

        return result;
    }

    // Deposit ETH to tss
    function deposit(address receiver) external payable {
        if (msg.value == 0) revert InsufficientETHAmount();
        (bool deposited, ) = tssAddress.call{value: msg.value}("");

        if (deposited == false) revert DepositFailed();
        
        emit Deposit(msg.sender, receiver, msg.value, address(0), "");
    }

    // Deposit ERC20 tokens to custody
    function deposit(address receiver, uint256 amount, address asset) external {
        if (amount == 0) revert InsufficientERC20Amount();
        IERC20(asset).safeTransferFrom(msg.sender, address(custody), amount);

        emit Deposit(msg.sender, receiver, amount, asset, "");
    }

    // Deposit ETH to tss and call an omnichain smart contract
    function depositAndCall(address receiver, bytes calldata payload) external payable {
        if (msg.value == 0) revert InsufficientETHAmount();
        (bool deposited, ) = tssAddress.call{value: msg.value}("");

        if (deposited == false) revert DepositFailed();
        
        emit Deposit(msg.sender, receiver, msg.value, address(0), payload);
    }

    // Deposit ERC20 tokens to custody and call an omnichain smart contract
    function depositAndCall(address receiver, uint256 amount, address asset, bytes calldata payload) external {
        if (amount == 0) revert InsufficientERC20Amount();
        IERC20(asset).safeTransferFrom(msg.sender, address(custody), amount);

        emit Deposit(msg.sender, receiver, amount, asset, payload);
    }

    // Call an omnichain smart contract without asset transfer
    function call(address receiver, bytes calldata payload) external {
        emit Call(msg.sender, receiver, payload);
    }

    function setCustody(address _custody) external {
        if (custody != address(0)) revert CustodyInitialized();
        custody = _custody;
    }

    function resetApproval(address token, address to) private returns (bool) {
        return IERC20(token).approve(to, 0);
    }
}
