// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

import "./ZetaConnectorNewBase.sol";
import "./IZetaNonEthNew.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";

contract ZetaConnectorNonNative is ZetaConnectorNewBase {
    uint256 public maxSupply = 2 ** 256 - 1;

    event MaxSupplyUpdated(uint256 maxSupply);
    error ExceedsMaxSupply();

    constructor(address _gateway, address _zetaToken, address _tssAddress)
        ZetaConnectorNewBase(_gateway, _zetaToken, _tssAddress)
    {}

    /// @notice set max supply for minting
    function setMaxSupply(uint256 maxSupply_) external onlyTSS() {
        maxSupply = maxSupply_;
        emit MaxSupplyUpdated(maxSupply_);
    }


    /// @dev withdraw is called by TSS address, it mints zetaToken to the destination address
    function withdraw(address to, uint256 amount, bytes32 internalSendHash) external override nonReentrant onlyTSS {
        if (amount + IERC20(zetaToken).totalSupply() > maxSupply) revert ExceedsMaxSupply();

        IZetaNonEthNew(zetaToken).mint(to, amount, internalSendHash);
        emit Withdraw(to, amount);
    }

    /// @dev withdrawAndCall is called by TSS address, it mints zetaToken and calls a contract
    function withdrawAndCall(address to, uint256 amount, bytes calldata data, bytes32 internalSendHash) external override nonReentrant onlyTSS {
        if (amount + IERC20(zetaToken).totalSupply() > maxSupply) revert ExceedsMaxSupply();

        // Mint zetaToken to the Gateway contract
        IZetaNonEthNew(zetaToken).mint(address(gateway), amount, internalSendHash);

        // Forward the call to the Gateway contract
        gateway.executeWithERC20(address(zetaToken), to, amount, data);

        emit WithdrawAndCall(to, amount, data);
    }

    /// @dev withdrawAndRevert is called by TSS address, it mints zetaToken to the gateway and calls onRevert on a contract
    function withdrawAndRevert(address to, uint256 amount, bytes calldata data, bytes32 internalSendHash) external override nonReentrant onlyTSS {
        if (amount + IERC20(zetaToken).totalSupply() > maxSupply) revert ExceedsMaxSupply();

        // Mint zetaToken to the Gateway contract
        IZetaNonEthNew(zetaToken).mint(address(gateway), amount, internalSendHash);

        // Forward the call to the Gateway contract
        gateway.revertWithERC20(address(zetaToken), to, amount, data);

        emit WithdrawAndRevert(to, amount, data);
    }

    /// @dev receiveTokens handles token transfer and burn them
    function receiveTokens(uint256 amount) external override {
        IZetaNonEthNew(zetaToken).burnFrom(msg.sender, amount);
    }
}
