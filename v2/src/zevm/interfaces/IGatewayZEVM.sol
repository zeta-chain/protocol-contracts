// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./zContract.sol";

interface IGatewayZEVM {
    function withdraw(bytes memory receiver, uint256 amount, address zrc20) external;

    function withdrawAndCall(bytes memory receiver, uint256 amount, address zrc20, bytes calldata message) external;

    function call(bytes memory receiver, bytes calldata message) external;

    function deposit(
        address zrc20,
        uint256 amount,
        address target
    ) external;

    function execute(
        zContext calldata context,
        address zrc20,
        uint256 amount,
        address target,
        bytes calldata message
    ) external;

    function depositAndCall(
        zContext calldata context,
        address zrc20,
        uint256 amount,
        address target,
        bytes calldata message
    ) external;
}

interface IGatewayZEVMEvents {
    event Call(address indexed sender, bytes receiver, bytes message);
    event Withdrawal(address indexed from, address zrc20, bytes to, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message);
}

interface IGatewayZEVMErrors {
    error WithdrawalFailed();
    error InsufficientZRC20Amount();
    error ZRC20BurnFailed();
    error ZRC20TransferFailed();
    error GasFeeTransferFailed();
    error CallerIsNotFungibleModule();
    error InvalidTarget();
    error FailedZetaSent();
    error OnlyWZETAOrFungible();
}