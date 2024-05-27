# ZetaInterfaces
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/2e5223462d9ac9dedd79e76ede471832bb2c40e7/contracts/evm/interfaces/ZetaInterfaces.sol)


## Structs
### SendInput
*Use SendInput to interact with the Connector: connector.send(SendInput)*


```solidity
struct SendInput {
    uint256 destinationChainId;
    bytes destinationAddress;
    uint256 destinationGasLimit;
    bytes message;
    uint256 zetaValueAndGas;
    bytes zetaParams;
}
```

### ZetaMessage
*Our Connector calls onZetaMessage with this struct as argument*


```solidity
struct ZetaMessage {
    bytes zetaTxSenderAddress;
    uint256 sourceChainId;
    address destinationAddress;
    uint256 zetaValue;
    bytes message;
}
```

### ZetaRevert
*Our Connector calls onZetaRevert with this struct as argument*


```solidity
struct ZetaRevert {
    address zetaTxSenderAddress;
    uint256 sourceChainId;
    bytes destinationAddress;
    uint256 destinationChainId;
    uint256 remainingZetaValue;
    bytes message;
}
```

