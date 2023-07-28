## ConnectorErrors

```solidity
import "@zetachain/protocol-contracts/contracts/evm/interfaces/ConnectorErrors.sol";
```

Source: https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/evm/interfaces/ConnectorErrors.sol

Interface with connector custom errors

### Error List

* [CallerIsNotPauser(caller)](#ConnectorErrors-CallerIsNotPauser-address-)
* [CallerIsNotTss(caller)](#ConnectorErrors-CallerIsNotTss-address-)
* [CallerIsNotTssUpdater(caller)](#ConnectorErrors-CallerIsNotTssUpdater-address-)
* [CallerIsNotTssOrUpdater(caller)](#ConnectorErrors-CallerIsNotTssOrUpdater-address-)
* [ZetaTransferError()](#ConnectorErrors-ZetaTransferError--)
* [ExceedsMaxSupply(maxSupply)](#ConnectorErrors-ExceedsMaxSupply-uint256-)

### Modifiers

### Errors

```
CallerIsNotPauser(address caller) (error)
```

<a name="ConnectorErrors-CallerIsNotPauser-address-"></a>

```
CallerIsNotTss(address caller) (error)
```

<a name="ConnectorErrors-CallerIsNotTss-address-"></a>

```
CallerIsNotTssUpdater(address caller) (error)
```

<a name="ConnectorErrors-CallerIsNotTssUpdater-address-"></a>

```
CallerIsNotTssOrUpdater(address caller) (error)
```

<a name="ConnectorErrors-CallerIsNotTssOrUpdater-address-"></a>

```
ZetaTransferError() (error)
```

<a name="ConnectorErrors-ZetaTransferError--"></a>

```
ExceedsMaxSupply(uint256 maxSupply) (error)
```

<a name="ConnectorErrors-ExceedsMaxSupply-uint256-"></a>

