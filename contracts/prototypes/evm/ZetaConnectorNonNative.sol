// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./ZetaConnectorNewBase.sol";
import "./IZetaNonEthNew.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";

contract ZetaConnectorNonNative is ZetaConnectorNewBase {
    constructor(address _gateway, address _zetaToken)
        ZetaConnectorNewBase(_gateway, _zetaToken)
    {}

    // Withdraw is called by TSS address, it mints zetaToken to the destination address
    function withdraw(address to, uint256 amount, bytes32 internalSendHash) external override nonReentrant {
        IZetaNonEthNew(zetaToken).mint(to, amount, internalSendHash);
        emit Withdraw(to, amount);
    }

    // WithdrawAndCall is called by TSS address, it mints zetaToken and calls a contract
    function withdrawAndCall(address to, uint256 amount, bytes calldata data, bytes32 internalSendHash) external override nonReentrant {
        // Mint zetaToken to the Gateway contract
        IZetaNonEthNew(zetaToken).mint(address(gateway), amount, internalSendHash);

        // Forward the call to the Gateway contract
        gateway.executeWithERC20(address(zetaToken), to, amount, data);

        emit WithdrawAndCall(to, amount, data);
    }

    // Function to handle token transfer and burn them
    function receiveTokens(uint256 amount) external override {
        // Burn the tokens
        IZetaNonEthNew(zetaToken).burnFrom(msg.sender, amount);
    }
}
