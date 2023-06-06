# Solidity API

## ZetaInterfaces

### SendInput

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

## WZETA

### transferFrom

```solidity
function transferFrom(address src, address dst, uint256 wad) external returns (bool)
```

### withdraw

```solidity
function withdraw(uint256 wad) external
```

## ZetaConnectorZEVM

### OnlyWZETA

```solidity
error OnlyWZETA()
```

Contract custom errors.

### WZETATransferFailed

```solidity
error WZETATransferFailed()
```

### OnlyFungibleModule

```solidity
error OnlyFungibleModule()
```

### FailedZetaSent

```solidity
error FailedZetaSent()
```

### FUNGIBLE_MODULE_ADDRESS

```solidity
address FUNGIBLE_MODULE_ADDRESS
```

Fungible module address.

### wzeta

```solidity
address wzeta
```

WZETA token address.

### ZetaSent

```solidity
event ZetaSent(address sourceTxOriginAddress, address zetaTxSenderAddress, uint256 destinationChainId, bytes destinationAddress, uint256 zetaValueAndGas, uint256 destinationGasLimit, bytes message, bytes zetaParams)
```

### SetWZETA

```solidity
event SetWZETA(address wzeta_)
```

### constructor

```solidity
constructor(address _wzeta) public
```

### receive

```solidity
receive() external payable
```

_Receive function to receive ZETA from WETH9.withdraw()._

### send

```solidity
function send(struct ZetaInterfaces.SendInput input) external
```

_Sends ZETA and bytes messages (to execute it) crosschain._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| input | struct ZetaInterfaces.SendInput |  |

### setWzetaAddress

```solidity
function setWzetaAddress(address wzeta_) external
```

_Sends ZETA and bytes messages (to execute it) crosschain._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| wzeta_ | address |  |

