// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;
import "../evm/interfaces/ZetaInterfaces.sol";
import "../evm/ZetaConnector.base.sol";
import "./interfaces/IWZETA.sol";

contract ZetaConnectorZEVM is ZetaConnectorBase {
    /// @notice Contract custom errors.
    error OnlyWZETA();
    error WZETATransferFailed();
    error OnlyFungibleModule();
    error FailedZetaSent();

    /// @notice Fungible module address.
    address public constant FUNGIBLE_MODULE_ADDRESS = payable(0x735b14BB79463307AAcBED86DAf3322B1e6226aB);

    /**
     * @dev Modifier to restrict actions to fungible module.
     */
    modifier onlyFungibleModule() {
        if (msg.sender != FUNGIBLE_MODULE_ADDRESS) revert OnlyFungibleModule();
        _;
    }

    constructor(
        address zetaTokenAddress_,
        address tssAddress_,
        address tssAddressUpdater_,
        address pauserAddress_
    ) ZetaConnectorBase(zetaTokenAddress_, tssAddress_, tssAddressUpdater_, pauserAddress_) {}

    /// @dev Receive function to receive ZETA from WETH9.withdraw().
    receive() external payable {
        if (msg.sender != zetaToken) revert OnlyWZETA();
    }

    /**
     * @dev Sends ZETA and bytes messages (to execute it) crosschain.
     * @param input, SendInput struct, checkout above.
     */
    function send(ZetaInterfaces.SendInput calldata input) external override whenNotPaused {
        // Transfer wzeta to "fungible" module, which will be burnt by the protocol post processing via hooks.
        if (!IWETH9(zetaToken).transferFrom(msg.sender, address(this), input.zetaValueAndGas))
            revert WZETATransferFailed();
        IWETH9(zetaToken).withdraw(input.zetaValueAndGas);
        (bool sent, ) = FUNGIBLE_MODULE_ADDRESS.call{value: input.zetaValueAndGas}("");
        if (!sent) revert FailedZetaSent();
        emit ZetaSent(
            tx.origin,
            msg.sender,
            input.destinationChainId,
            input.destinationAddress,
            input.zetaValueAndGas,
            input.destinationGasLimit,
            input.message,
            input.zetaParams
        );
    }

    /**
     * @dev Handler to receive data from other chain.
     * This method can be called only by Fungible Module.
     * Transfer the Zeta tokens to destination and calls onZetaMessage if it's needed.
     * To perform the transfer mint new tokens, validating first the maxSupply allowed in the current chain.
     */
    function onReceive(
        bytes calldata zetaTxSenderAddress,
        uint256 sourceChainId,
        address destinationAddress,
        uint256 zetaValue,
        bytes calldata message,
        bytes32 internalSendHash
    ) external override onlyFungibleModule {
        IWETH9(zetaToken).deposit{value: zetaValue}();
        if (!IWETH9(zetaToken).transferFrom(address(this), destinationAddress, zetaValue)) revert WZETATransferFailed();

        if (message.length > 0) {
            ZetaReceiver(destinationAddress).onZetaMessage(
                ZetaInterfaces.ZetaMessage(zetaTxSenderAddress, sourceChainId, destinationAddress, zetaValue, message)
            );
        }

        emit ZetaReceived(zetaTxSenderAddress, sourceChainId, destinationAddress, zetaValue, message, internalSendHash);
    }

    /**
     * @dev Handler to receive errors from other chain.
     * This method can be called only by Fungible Module.
     * Transfer the Zeta tokens to destination and calls onZetaRevert if it's needed.
     */
    function onRevert(
        address zetaTxSenderAddress,
        uint256 sourceChainId,
        bytes calldata destinationAddress,
        uint256 destinationChainId,
        uint256 remainingZetaValue,
        bytes calldata message,
        bytes32 internalSendHash
    ) external override whenNotPaused onlyFungibleModule {
        IWETH9(zetaToken).deposit{value: remainingZetaValue}();
        if (!IWETH9(zetaToken).transferFrom(address(this), zetaTxSenderAddress, remainingZetaValue))
            revert WZETATransferFailed();

        if (message.length > 0) {
            ZetaReceiver(zetaTxSenderAddress).onZetaRevert(
                ZetaInterfaces.ZetaRevert(
                    zetaTxSenderAddress,
                    sourceChainId,
                    destinationAddress,
                    destinationChainId,
                    remainingZetaValue,
                    message
                )
            );
        }

        emit ZetaReverted(
            zetaTxSenderAddress,
            sourceChainId,
            destinationChainId,
            destinationAddress,
            remainingZetaValue,
            message,
            internalSendHash
        );
    }
}
