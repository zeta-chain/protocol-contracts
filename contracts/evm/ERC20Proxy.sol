// SPDX-License-Identifier: MIT
// v1.0, 2023-01-10
pragma solidity 0.8.7;
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
/* import "@openzeppelin/contracts/utils/Multicall.sol"; */


interface ZRC20Receiver {
    function onZRC20Message(bytes calldata,uint256,address,address,uint256,bytes calldata) external;
}

interface ERC20Custody {
    function withdraw(address recipient, IERC20 asset, uint256 amount) external; 
}

contract ERC20Proxy  {
    
    function onZRC20WithdrawAndCall(
        bytes calldata zetaTxSenderAddress,
        uint256 sourceChainId,
        address destinationAddress,
        address asset, // ERC20 address
        uint256 amount,
        bytes calldata message,
        bytes32 internalSendHash,
	address ERC20CustodyAddress
    ) external {
	// delegate call:    ERC20Custody(ERC20CustodyAddress).withdraw(address(this), IERC20(asset), amount);
	bytes memory withdrawData = abi.encodeWithSelector(
            ERC20Custody.withdraw.selector,
            destinationAddress,
            IERC20(asset),
            amount
        );
	(bool success, ) = ERC20CustodyAddress.delegatecall(withdrawData);
        require(success, "delegate call to withdraw failed");

        if (message.length > 0) {
            ZRC20Receiver(destinationAddress).onZRC20Message(
                zetaTxSenderAddress, sourceChainId, destinationAddress, asset, amount, message
            );
        }

        /* emit ZetaReceived(zetaTxSenderAddress, sourceChainId, destinationAddress, zetaValue, message, internalSendHash); */
    }

}



// an trivial example app contract that just forwards the received ERC20 contract to the sender on destination chain. 
contract AppContract is ZRC20Receiver {
    function onZRC20Message(bytes calldata sender, uint256 srcChainId, address destAddr, address asset, uint256 amount, bytes calldata message) public override{
	require(sender.length >= 20, "Input bytes length must be at least 20");
        address addr;

        assembly {
            // Calldata is read-only, so use calldatacopy to copy data to memory first
            let memPtr := mload(0x40) // Get a free memory pointer
            calldatacopy(memPtr, sender.offset, 20) // Copy the first 20 bytes to memory
            addr := mload(memPtr) // Load the 20 bytes from memory into the address
        }

        IERC20(asset).transfer(addr, amount); 
    }

}


