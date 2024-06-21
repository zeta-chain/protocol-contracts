// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";


// NOTE: Purpose of this contract is to test upgrade process, the only difference should be event names
// The Gateway contract is the endpoint to call smart contracts on external chains
// The contract doesn't hold any funds and should never have active allowances
contract GatewayUpgradeTest is Initializable, OwnableUpgradeable, UUPSUpgradeable {
    error ExecutionFailed();

    address public custody;

    event ExecutedV2(address indexed destination, uint256 value, bytes data);
    event ExecutedWithERC20V2(address indexed token, address indexed to, uint256 amount, bytes data);

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize() public initializer {
        __Ownable_init();
        __UUPSUpgradeable_init();
    }

    function _authorizeUpgrade(address newImplementation) internal override onlyOwner() {}

    function _execute(address destination, bytes calldata data) internal returns (bytes memory) {
        (bool success, bytes memory result) = destination.call{value: msg.value}(data);
    
        if (!success) {
            revert ExecutionFailed();
        }

        return result;
    }

    // Called by the TSS
    // Execution without ERC20 tokens, it is payable and can be used in the case of WithdrawAndCall for Gas ZRC20
    // It can be also used for contract call without asset movement
    function execute(address destination, bytes calldata data) external payable returns (bytes memory) {
        bytes memory result = _execute(destination, data);

        emit ExecutedV2(destination, msg.value, data);

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

        emit ExecutedWithERC20V2(token, to, amount, data);

        return result;
    }

    function setCustody(address _custody) external {
        custody = _custody;
    }
}
