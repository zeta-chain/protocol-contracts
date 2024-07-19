// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import "../../zevm/interfaces/IZRC20.sol";
import "../../zevm/interfaces/zContract.sol";
import "./IGatewayZEVM.sol";
import "../../zevm/interfaces/IWZETA.sol";


// The GatewayZEVM contract is the endpoint to call smart contracts on omnichain
// The contract doesn't hold any funds and should never have active allowances
contract GatewayZEVM is IGatewayZEVMEvents, IGatewayZEVMErrors, Initializable, OwnableUpgradeable, UUPSUpgradeable, ReentrancyGuard {
    error ZeroAddress();

    address public constant FUNGIBLE_MODULE_ADDRESS = 0x735b14BB79463307AAcBED86DAf3322B1e6226aB;
    address public zetaToken;

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize(address _zetaToken) public initializer {
        if (_zetaToken == address(0)) {
            revert ZeroAddress();
        }
    
        __Ownable_init();
        __UUPSUpgradeable_init();
        zetaToken = _zetaToken;
    }

    function _authorizeUpgrade(address newImplementation) internal override onlyOwner() {}

    function _withdrawZRC20(uint256 amount, address zrc20) internal returns (uint256) {
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

    function _transferZETA(uint256 amount, address to) internal {
        if (!IWETH9(zetaToken).transferFrom(msg.sender, address(this), amount)) revert FailedZetaSent();
        IWETH9(zetaToken).withdraw(amount);
        (bool sent, ) = to.call{value: amount}("");
        if (!sent) revert FailedZetaSent();
    }

    // Withdraw ZRC20 tokens to external chain
    function withdraw(bytes memory receiver, uint256 amount, address zrc20) external nonReentrant {
        uint256 gasFee = _withdrawZRC20(amount, zrc20);
        emit Withdrawal(msg.sender, zrc20, receiver, amount, gasFee, IZRC20(zrc20).PROTOCOL_FLAT_FEE(), "");
    }

    // Withdraw ZRC20 tokens and call smart contract on external chain
    function withdrawAndCall(bytes memory receiver, uint256 amount, address zrc20, bytes calldata message) external nonReentrant {
        uint256 gasFee = _withdrawZRC20(amount, zrc20);
        emit Withdrawal(msg.sender, zrc20, receiver, amount, gasFee, IZRC20(zrc20).PROTOCOL_FLAT_FEE(), message);
    }

    // Withdraw ZETA to external chain
    function withdraw(uint256 amount) external nonReentrant {
        _transferZETA(amount, FUNGIBLE_MODULE_ADDRESS);
        emit Withdrawal(msg.sender, address(0), abi.encodePacked(FUNGIBLE_MODULE_ADDRESS), amount, 0, 0, "");
    }

    // Withdraw ZETA and call smart contract on external chain
    function withdrawAndCall(uint256 amount, bytes calldata message) external nonReentrant {
        _transferZETA(amount, FUNGIBLE_MODULE_ADDRESS);
        emit Withdrawal(msg.sender, address(0), abi.encodePacked(FUNGIBLE_MODULE_ADDRESS), amount, 0, 0, message);
    }

    // Call smart contract on external chain without asset transfer
    function call(bytes memory receiver, bytes calldata message) external nonReentrant {
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

    // Deposit zeta and call user specified contract on ZEVM
    // TODO: Finalize access control
    // https://github.com/zeta-chain/protocol-contracts/issues/204
    function depositAndCall(
        zContext calldata context,
        uint256 amount,
        address target,
        bytes calldata message
    ) external {
        if (msg.sender != FUNGIBLE_MODULE_ADDRESS) revert CallerIsNotFungibleModule();
        if (target == FUNGIBLE_MODULE_ADDRESS || target == address(this)) revert InvalidTarget();

        _transferZETA(amount, target);
        zContract(target).onCrossChainCall(context, zetaToken, amount, message);
    }
}
