## ZetaInterfaces

```solidity
import "@zetachain/protocol-contracts/contracts/zevm/ZetaConnectorZEVM.sol";
```

Source: https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/zevm/ZetaConnectorZEVM.sol

### Modifiers

## WZETA

```solidity
import "@zetachain/protocol-contracts/contracts/zevm/ZetaConnectorZEVM.sol";
```

Source: https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/zevm/ZetaConnectorZEVM.sol

### Function List

* [transferFrom(src, dst, wad)](#WZETA-transferFrom-address-address-uint256-)
* [withdraw(wad)](#WZETA-withdraw-uint256-)

### Modifiers

### Functions

```
transferFrom(address src, address dst, uint256 wad) → bool (external function)
```

<a name="WZETA-transferFrom-address-address-uint256-"></a>

```
withdraw(uint256 wad) (external function)
```

<a name="WZETA-withdraw-uint256-"></a>

## ZetaConnectorZEVM

```solidity
import "@zetachain/protocol-contracts/contracts/zevm/ZetaConnectorZEVM.sol";
```

Source: https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/zevm/ZetaConnectorZEVM.sol

### Function List

* [constructor(wzeta_)](#ZetaConnectorZEVM-constructor-address-)
* [receive()](#ZetaConnectorZEVM-receive--)
* [send(input)](#ZetaConnectorZEVM-send-struct-ZetaInterfaces-SendInput-)
* [setWzetaAddress(wzeta_)](#ZetaConnectorZEVM-setWzetaAddress-address-)

### Event List

* [ZetaSent(sourceTxOriginAddress, zetaTxSenderAddress, destinationChainId, destinationAddress, zetaValueAndGas, destinationGasLimit, message, zetaParams)](#ZetaConnectorZEVM-ZetaSent-address-address-uint256-bytes-uint256-uint256-bytes-bytes-)
* [SetWZETA(wzeta_)](#ZetaConnectorZEVM-SetWZETA-address-)

### Error List

* [OnlyWZETA()](#ZetaConnectorZEVM-OnlyWZETA--)
* [WZETATransferFailed()](#ZetaConnectorZEVM-WZETATransferFailed--)
* [OnlyFungibleModule()](#ZetaConnectorZEVM-OnlyFungibleModule--)
* [FailedZetaSent()](#ZetaConnectorZEVM-FailedZetaSent--)

### Modifiers

### Functions

```
constructor(address wzeta_) (public function)
```

<a name="ZetaConnectorZEVM-constructor-address-"></a>

```
receive() (external function)
```

<a name="ZetaConnectorZEVM-receive--"></a>

Receive function to receive ZETA from WETH9.withdraw().

```
send(struct ZetaInterfaces.SendInput input) (external function)
```

<a name="ZetaConnectorZEVM-send-struct-ZetaInterfaces-SendInput-"></a>

Sends ZETA and bytes messages (to execute it) crosschain.

```
setWzetaAddress(address wzeta_) (external function)
```

<a name="ZetaConnectorZEVM-setWzetaAddress-address-"></a>

Sends ZETA and bytes messages (to execute it) crosschain.

### Events

```
ZetaSent(address sourceTxOriginAddress, address indexed zetaTxSenderAddress, uint256 indexed destinationChainId, bytes destinationAddress, uint256 zetaValueAndGas, uint256 destinationGasLimit, bytes message, bytes zetaParams) (event)
```

<a name="ZetaConnectorZEVM-ZetaSent-address-address-uint256-bytes-uint256-uint256-bytes-bytes-"></a>

```
SetWZETA(address wzeta_) (event)
```

<a name="ZetaConnectorZEVM-SetWZETA-address-"></a>

### Errors

```
OnlyWZETA() (error)
```

<a name="ZetaConnectorZEVM-OnlyWZETA--"></a>

```
WZETATransferFailed() (error)
```

<a name="ZetaConnectorZEVM-WZETATransferFailed--"></a>

```
OnlyFungibleModule() (error)
```

<a name="ZetaConnectorZEVM-OnlyFungibleModule--"></a>

```
FailedZetaSent() (error)
```

<a name="ZetaConnectorZEVM-FailedZetaSent--"></a>

