## zContext

```solidity
struct zContext {
  bytes origin;
  address sender;
  uint256 chainID;
}
```
## zContract

```solidity
import "@zetachain/protocol-contracts/contracts/zevm/interfaces/zContract.sol";
```

Source: https://github.com/zeta-chain/protocol-contracts/blob/main/contracts/zevm/interfaces/zContract.sol

### Function List

* [onCrossChainCall(context, zrc20, amount, message)](#zContract-onCrossChainCall-struct-zContext-address-uint256-bytes-)

### Modifiers

### Functions

```
onCrossChainCall(struct zContext context, address zrc20, uint256 amount, bytes message) (external function)
```

<a name="zContract-onCrossChainCall-struct-zContext-address-uint256-bytes-"></a>

