---
title: "Protocol contracts"
---

# ZetaChain Protocol Contracts

Contracts of official protocol contracts deployed by the core ZetaChain team.

## Learn more about ZetaChain

* Check our [website](https://www.zetachain.com/).
* Read our [docs](https://docs.zetachain.com/).

## Codebase

The protocol contracts codebase is separated into two sections:

- `zevm`: contracts deployed on ZetaChain
- `evm`: contracts deployed on EVM connected chains (Ethereum, Base, etc..)

**ZEVM Contracts**

- `GatewayZEVM`: entrypoints for interaction. The users call the gateway contract to initiate cctx. The gateway contract is also the contract calling the target when a smart contract call is initiated on a connected chain
- `ZRC20`: is a ERC20 compliant contract that represents fungible assets from connected chains
- `WZETA`: wrapped ZETA (fork of WETH)

**EVM Contracts**

- `GatewayEVM`: similar to GatewayZEVM for connected chains. Entrypoint for users.
- `ERC20Custody`: hold ERC20 assets being sent to ZetaChain
- `ZetaConnector`: manage ZETA for connected chains. There are two version of the contract:
    - Native: when the ZETA tokens are native to the chain (ZETA where initially deployed as a ERC20 on Ethereum, not emitted fully on ZetaChain). Use lock/unlock model.
    - Non-Native: when ZETA tokens where never native to the chains but withdrawn from ZetaChain. Use mint/burn model.
- TSS (EOA): Threshold-signature-scheme wallet, this address holds the gas token sent to ZetaChain (like Ethers)

[Learn more about the Gateway contracts](https://www.zetachain.com/docs/developers/evm/gateway/)

## Usage

**Install dependencies**

```shell
$ yarn
$ forge soldeer update
```

**Build**

```shell
$ forge build
```

**Test**

```shell
$ forge test
```

**Format**

```shell
$ forge fmt
```

**Deploy**

```shell
$ forge script script/<DeployScript>.s.sol:<DeployScript> --rpc-url <your_rpc_url> --private-key <your_private_key>
```

To view detailed instructions on how to deploy the contracts, please refer to the [Deployment Guide](./scripts/deploy/readme.md).

This guide covers all steps required to deploy the contracts, including environment setup, commands, and best practices.

## Community

[X](https://x.com/zetablockchain) | [Discord](https://discord.com/invite/zetachain) | [Telegram](https://t.me/zetachainofficial) | [Website](https://zetachain.com)# Summary
- [Home](README.md)
# contracts
  - [❱ evm](protocol/contracts/evm/README.md)
    - [❱ interfaces](protocol/contracts/evm/interfaces/README.md)
      - [IERC20CustodyEvents](protocol/contracts/evm/interfaces/IERC20Custody.sol/interface.IERC20CustodyEvents.md)
      - [IERC20CustodyErrors](protocol/contracts/evm/interfaces/IERC20Custody.sol/interface.IERC20CustodyErrors.md)
      - [IERC20Custody](protocol/contracts/evm/interfaces/IERC20Custody.sol/interface.IERC20Custody.md)
      - [IGatewayEVMEvents](protocol/contracts/evm/interfaces/IGatewayEVM.sol/interface.IGatewayEVMEvents.md)
      - [IGatewayEVMErrors](protocol/contracts/evm/interfaces/IGatewayEVM.sol/interface.IGatewayEVMErrors.md)
      - [IGatewayEVM](protocol/contracts/evm/interfaces/IGatewayEVM.sol/interface.IGatewayEVM.md)
      - [MessageContext](protocol/contracts/evm/interfaces/IGatewayEVM.sol/struct.MessageContext.md)
      - [Callable](protocol/contracts/evm/interfaces/IGatewayEVM.sol/interface.Callable.md)
      - [IRegistryEvents](protocol/contracts/evm/interfaces/IRegistry.sol/interface.IRegistryEvents.md)
      - [IRegistryErrors](protocol/contracts/evm/interfaces/IRegistry.sol/interface.IRegistryErrors.md)
      - [ChainInfo](protocol/contracts/evm/interfaces/IRegistry.sol/struct.ChainInfo.md)
      - [ContractInfo](protocol/contracts/evm/interfaces/IRegistry.sol/struct.ContractInfo.md)
      - [ZRC20Info](protocol/contracts/evm/interfaces/IRegistry.sol/struct.ZRC20Info.md)
      - [IRegistry](protocol/contracts/evm/interfaces/IRegistry.sol/interface.IRegistry.md)
      - [IZetaConnectorEvents](protocol/contracts/evm/interfaces/IZetaConnector.sol/interface.IZetaConnectorEvents.md)
      - [IZetaNonEthNew](protocol/contracts/evm/interfaces/IZetaNonEthNew.sol/interface.IZetaNonEthNew.md)
    - [❱ legacy](protocol/contracts/evm/legacy/README.md)
      - [ConnectorErrors](protocol/contracts/evm/legacy/ConnectorErrors.sol/interface.ConnectorErrors.md)
      - [IZetaNonEthInterface](protocol/contracts/evm/legacy/IZetaNonEthInterface.sol/interface.IZetaNonEthInterface.md)
      - [ZetaConnectorBase](protocol/contracts/evm/legacy/ZetaConnector.base.sol/contract.ZetaConnectorBase.md)
      - [ZetaConnectorEth](protocol/contracts/evm/legacy/ZetaConnector.eth.sol/contract.ZetaConnectorEth.md)
      - [ZetaConnectorNonEth](protocol/contracts/evm/legacy/ZetaConnector.non-eth.sol/contract.ZetaConnectorNonEth.md)
      - [ZetaErrors](protocol/contracts/evm/legacy/ZetaErrors.sol/interface.ZetaErrors.md)
      - [ZetaEth](protocol/contracts/evm/legacy/ZetaEth.sol/contract.ZetaEth.md)
      - [ZetaInterfaces](protocol/contracts/evm/legacy/ZetaInterfaces.sol/interface.ZetaInterfaces.md)
      - [ZetaConnector](protocol/contracts/evm/legacy/ZetaInterfaces.sol/interface.ZetaConnector.md)
      - [ZetaReceiver](protocol/contracts/evm/legacy/ZetaInterfaces.sol/interface.ZetaReceiver.md)
      - [ZetaTokenConsumer](protocol/contracts/evm/legacy/ZetaInterfaces.sol/interface.ZetaTokenConsumer.md)
      - [ZetaCommonErrors](protocol/contracts/evm/legacy/ZetaInterfaces.sol/interface.ZetaCommonErrors.md)
      - [ZetaNonEth](protocol/contracts/evm/legacy/ZetaNonEth.sol/contract.ZetaNonEth.md)
    - [ERC20Custody](protocol/contracts/evm/ERC20Custody.sol/contract.ERC20Custody.md)
    - [GatewayEVM](protocol/contracts/evm/GatewayEVM.sol/contract.GatewayEVM.md)
    - [Registry](protocol/contracts/evm/Registry.sol/contract.Registry.md)
    - [ZetaConnectorBase](protocol/contracts/evm/ZetaConnectorBase.sol/abstract.ZetaConnectorBase.md)
    - [ZetaConnectorNative](protocol/contracts/evm/ZetaConnectorNative.sol/contract.ZetaConnectorNative.md)
    - [ZetaConnectorNonNative](protocol/contracts/evm/ZetaConnectorNonNative.sol/contract.ZetaConnectorNonNative.md)
  - [❱ helpers](protocol/contracts/helpers/README.md)
    - [❱ interfaces](protocol/contracts/helpers/interfaces/README.md)
      - [ChainInfo](protocol/contracts/helpers/interfaces/IBaseRegistry.sol/struct.ChainInfo.md)
      - [ContractInfo](protocol/contracts/helpers/interfaces/IBaseRegistry.sol/struct.ContractInfo.md)
      - [ZRC20Info](protocol/contracts/helpers/interfaces/IBaseRegistry.sol/struct.ZRC20Info.md)
      - [IBaseRegistryEvents](protocol/contracts/helpers/interfaces/IBaseRegistry.sol/interface.IBaseRegistryEvents.md)
      - [IBaseRegistryErrors](protocol/contracts/helpers/interfaces/IBaseRegistry.sol/interface.IBaseRegistryErrors.md)
      - [IBaseRegistry](protocol/contracts/helpers/interfaces/IBaseRegistry.sol/interface.IBaseRegistry.md)
    - [BaseRegistry](protocol/contracts/helpers/BaseRegistry.sol/abstract.BaseRegistry.md)
  - [❱ zevm](protocol/contracts/zevm/README.md)
    - [❱ interfaces](protocol/contracts/zevm/interfaces/README.md)
      - [ICoreRegistryEvents](protocol/contracts/zevm/interfaces/ICoreRegistry.sol/interface.ICoreRegistryEvents.md)
      - [ICoreRegistryErrors](protocol/contracts/zevm/interfaces/ICoreRegistry.sol/interface.ICoreRegistryErrors.md)
      - [ChainInfo](protocol/contracts/zevm/interfaces/ICoreRegistry.sol/struct.ChainInfo.md)
      - [ContractInfo](protocol/contracts/zevm/interfaces/ICoreRegistry.sol/struct.ContractInfo.md)
      - [ZRC20Info](protocol/contracts/zevm/interfaces/ICoreRegistry.sol/struct.ZRC20Info.md)
      - [ICoreRegistry](protocol/contracts/zevm/interfaces/ICoreRegistry.sol/interface.ICoreRegistry.md)
      - [IGatewayZEVMEvents](protocol/contracts/zevm/interfaces/IGatewayZEVM.sol/interface.IGatewayZEVMEvents.md)
      - [IGatewayZEVMErrors](protocol/contracts/zevm/interfaces/IGatewayZEVM.sol/interface.IGatewayZEVMErrors.md)
      - [IGatewayZEVM](protocol/contracts/zevm/interfaces/IGatewayZEVM.sol/interface.IGatewayZEVM.md)
      - [CallOptions](protocol/contracts/zevm/interfaces/IGatewayZEVM.sol/struct.CallOptions.md)
      - [ISystem](protocol/contracts/zevm/interfaces/ISystem.sol/interface.ISystem.md)
      - [IWETH9](protocol/contracts/zevm/interfaces/IWZETA.sol/interface.IWETH9.md)
      - [IZRC20](protocol/contracts/zevm/interfaces/IZRC20.sol/interface.IZRC20.md)
      - [IZRC20Metadata](protocol/contracts/zevm/interfaces/IZRC20.sol/interface.IZRC20Metadata.md)
      - [ZRC20Events](protocol/contracts/zevm/interfaces/IZRC20.sol/interface.ZRC20Events.md)
      - [CoinType](protocol/contracts/zevm/interfaces/IZRC20.sol/enum.CoinType.md)
      - [zContext](protocol/contracts/zevm/interfaces/UniversalContract.sol/struct.zContext.md)
      - [zContract](protocol/contracts/zevm/interfaces/UniversalContract.sol/interface.zContract.md)
      - [MessageContext](protocol/contracts/zevm/interfaces/UniversalContract.sol/struct.MessageContext.md)
      - [UniversalContract](protocol/contracts/zevm/interfaces/UniversalContract.sol/interface.UniversalContract.md)
    - [❱ legacy](protocol/contracts/zevm/legacy/README.md)
      - [ZetaInterfaces](protocol/contracts/zevm/legacy/ZetaConnectorZEVM.sol/interface.ZetaInterfaces.md)
      - [ZetaReceiver](protocol/contracts/zevm/legacy/ZetaConnectorZEVM.sol/interface.ZetaReceiver.md)
      - [ZetaConnectorZEVM](protocol/contracts/zevm/legacy/ZetaConnectorZEVM.sol/contract.ZetaConnectorZEVM.md)
    - [CoreRegistry](protocol/contracts/zevm/CoreRegistry.sol/contract.CoreRegistry.md)
    - [GatewayZEVM](protocol/contracts/zevm/GatewayZEVM.sol/contract.GatewayZEVM.md)
    - [SystemContractErrors](protocol/contracts/zevm/SystemContract.sol/interface.SystemContractErrors.md)
    - [SystemContract](protocol/contracts/zevm/SystemContract.sol/contract.SystemContract.md)
    - [WETH9](protocol/contracts/zevm/WZETA.sol/contract.WETH9.md)
    - [ZRC20Errors](protocol/contracts/zevm/ZRC20.sol/interface.ZRC20Errors.md)
    - [ZRC20](protocol/contracts/zevm/ZRC20.sol/contract.ZRC20.md)
  - [INotSupportedMethods](protocol/contracts/Errors.sol/interface.INotSupportedMethods.md)
  - [RevertGasLimitExceeded](protocol/contracts/Revert.sol/error.RevertGasLimitExceeded.md)
  - [RevertOptions](protocol/contracts/Revert.sol/struct.RevertOptions.md)
  - [RevertContext](protocol/contracts/Revert.sol/struct.RevertContext.md)
  - [AbortContext](protocol/contracts/Revert.sol/struct.AbortContext.md)
  - [Revertable](protocol/contracts/Revert.sol/interface.Revertable.md)
  - [Abortable](protocol/contracts/Revert.sol/interface.Abortable.md)
  - [Revert constants](protocol/contracts/Revert.sol/constants.Revert.md)
