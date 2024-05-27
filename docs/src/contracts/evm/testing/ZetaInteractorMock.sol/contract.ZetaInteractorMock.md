# ZetaInteractorMock
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/2e5223462d9ac9dedd79e76ede471832bb2c40e7/contracts/evm/testing/ZetaInteractorMock.sol)

**Inherits:**
Ownable2Step, [ZetaInteractor](/contracts/evm/tools/ZetaInteractor.sol/abstract.ZetaInteractor.md), [ZetaReceiver](/contracts/zevm/ZetaConnectorZEVM.sol/interface.ZetaReceiver.md)


## Functions
### constructor


```solidity
constructor(address zetaConnectorAddress) ZetaInteractor(zetaConnectorAddress);
```

### onZetaMessage


```solidity
function onZetaMessage(ZetaInterfaces.ZetaMessage calldata zetaMessage)
    external
    override
    isValidMessageCall(zetaMessage);
```

### onZetaRevert


```solidity
function onZetaRevert(ZetaInterfaces.ZetaRevert calldata zetaRevert) external override isValidRevertCall(zetaRevert);
```

