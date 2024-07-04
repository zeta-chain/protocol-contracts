// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "../../zevm/interfaces/IZRC20.sol";
import "../../zevm/interfaces/zContract.sol";

// The GatewayZEVM contract is the endpoint to call smart contracts on omnichain
// The contract doesn't hold any funds and should never have active allowances
contract GatewayZEVM is Initializable, OwnableUpgradeable, UUPSUpgradeable {
    address public constant FUNGIBLE_MODULE_ADDRESS = 0x735b14BB79463307AAcBED86DAf3322B1e6226aB;

    error WithdrawalFailed();
    error InsufficientZRC20Amount();
    error ZRC20BurnFailed();
    error ZRC20TransferFailed();
    error GasFeeTransferFailed();
    error CallerIsNotFungibleModule();
    error InvalidTarget();

    event Call(address indexed sender, bytes receiver, bytes message);
    event Withdrawal(address indexed from, bytes to, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message);

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize() public initializer {
        __Ownable_init();
        __UUPSUpgradeable_init();
    }

    function _authorizeUpgrade(address newImplementation) internal override onlyOwner() {}

    function _withdraw(uint256 amount, address zrc20) internal returns (uint256) {
        (address gasZRC20, uint256 gasFee) = IZRC20(zrc20).withdrawGasFee();
        if (!IZRC20(gasZRC20).transferFrom(msg.sender, FUNGIBLE_MODULE_ADDRESS, gasFee)) {
            revert GasFeeTransferFailed();
        }

        if (!IZRC20(zrc20).transferFrom(msg.sender, address(this), amount)) {
            revert ZRC20TransferFailed();
        }

        if (!IZRC20(zrc20).burn(amount)) revert ZRC20BurnFailed();

        return gasFee;
    }

    // Withdraw ZRC20 tokens to external chain
    function withdraw(bytes memory receiver, uint256 amount, address zrc20) external {
        uint256 gasFee = _withdraw(amount, zrc20);
        emit Withdrawal(msg.sender, receiver, amount, gasFee, IZRC20(zrc20).PROTOCOL_FLAT_FEE(), "");
    }

    // Withdraw ZRC20 tokens and call smart contract on external chain
    function withdrawAndCall(bytes memory receiver, uint256 amount, address zrc20, bytes calldata message) external {
        uint256 gasFee = _withdraw(amount, zrc20);
        emit Withdrawal(msg.sender, receiver, amount, gasFee, IZRC20(zrc20).PROTOCOL_FLAT_FEE(), message);
    }

    // Call smart contract on external chain without asset transfer
    function call(bytes memory receiver, bytes calldata message) external {
        emit Call(msg.sender, receiver, message);
    }

    // Deposit foreign coins into ZRC20
    // TODO: Finalize access control
    // https://github.com/zeta-chain/protocol-contracts/issues/204
    function deposit(
        address zrc20,
        uint256 amount,
        address target
    ) external {
        if (msg.sender != FUNGIBLE_MODULE_ADDRESS) revert CallerIsNotFungibleModule();
        if (target == FUNGIBLE_MODULE_ADDRESS || target == address(this)) revert InvalidTarget();

        IZRC20(zrc20).deposit(target, amount);
    }

    // Execute user specified contract on ZEVM
    // TODO: Finalize access control
    // https://github.com/zeta-chain/protocol-contracts/issues/204
    function execute(
        zContext calldata context,
        address zrc20,
        uint256 amount,
        address target,
        bytes calldata message
    ) external {
        if (msg.sender != FUNGIBLE_MODULE_ADDRESS) revert CallerIsNotFungibleModule();

        zContract(target).onCrossChainCall(context, zrc20, amount, message);
    }

    // Deposit foreign coins into ZRC20 and call user specified contract on ZEVM
    // TODO: Finalize access control
    // https://github.com/zeta-chain/protocol-contracts/issues/204
    function depositAndCall(
        zContext calldata context,
        address zrc20,
        uint256 amount,
        address target,
        bytes calldata message
    ) external {
        if (msg.sender != FUNGIBLE_MODULE_ADDRESS) revert CallerIsNotFungibleModule();
        if (target == FUNGIBLE_MODULE_ADDRESS || target == address(this)) revert InvalidTarget();

        IZRC20(zrc20).deposit(target, amount);
        zContract(target).onCrossChainCall(context, zrc20, amount, message);
    }
}
