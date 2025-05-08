// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "forge-std/Test.sol";
import "forge-std/Vm.sol";
import { Upgrades } from "openzeppelin-foundry-upgrades/Upgrades.sol";

import { CoreRegistry } from "../contracts/zevm/CoreRegistry.sol";

import "../contracts/zevm/ZRC20.sol";
import "../contracts/zevm/interfaces/ICoreRegistry.sol";
import "../contracts/zevm/interfaces/IGatewayZEVM.sol";

// Mock GatewayZEVM
contract MockGatewayZEVM {
    event CallEmitted(
        bytes receiver, address zrc20, bytes message, CallOptions callOptions, RevertOptions revertOptions
    );

    function call(
        bytes memory receiver,
        address zrc20,
        bytes calldata message,
        CallOptions calldata callOptions,
        RevertOptions calldata revertOptions
    )
        external
    {
        emit CallEmitted(receiver, zrc20, message, callOptions, revertOptions);
    }
}

contract CoreRegistryTest is Test, ICoreRegistryErrors, ICoreRegistryEvents {
    address payable proxy;
    CoreRegistry registry;
    MockGatewayZEVM mockGateway;
    ZRC20 gasZRC20;

    address admin;
    address registryManager;
    address user;

    function setUp() public {
        admin = address(0xABCD);
        registryManager = address(0x1234);
        user = address(0x5678);

        mockGateway = new MockGatewayZEVM();

        proxy = payable(
            Upgrades.deployUUPSProxy(
                "CoreRegistry.sol",
                abi.encodeCall(CoreRegistry.initialize, (admin, registryManager, address(mockGateway)))
            )
        );
        registry = CoreRegistry(proxy);

        gasZRC20 = new ZRC20("TOKEN", "TKN", 18, 1, CoinType.Gas, 0, address(1), address(mockGateway));
    }

    function testInitialize() public view {
        assertTrue(registry.hasRole(registry.DEFAULT_ADMIN_ROLE(), admin));
        assertTrue(registry.hasRole(registry.REGISTRY_MANAGER_ROLE(), registryManager));
        assertTrue(registry.hasRole(registry.PAUSER_ROLE(), registryManager));
        assertTrue(registry.hasRole(registry.PAUSER_ROLE(), admin));
        assertEq(address(registry.gatewayZEVM()), address(mockGateway));
    }

    function testInitializeWithZeroAddress() public {
        address payable uninitializedProxy = payable(Upgrades.deployUUPSProxy("CoreRegistry.sol", ""));

        CoreRegistry uninitializedRegistry = CoreRegistry(uninitializedProxy);

        vm.expectRevert(ZeroAddress.selector);
        uninitializedRegistry.initialize(address(0), registryManager, address(mockGateway));

        vm.expectRevert(ZeroAddress.selector);
        uninitializedRegistry.initialize(admin, address(0), address(mockGateway));

        vm.expectRevert(ZeroAddress.selector);
        uninitializedRegistry.initialize(admin, registryManager, address(0));
    }

    function testPause() public {
        assertFalse(registry.paused());

        vm.prank(admin);
        registry.pause();
        assertTrue(registry.paused());

        vm.prank(registryManager);
        registry.unpause();
        assertFalse(registry.paused());
    }

    function testPauseUnauthorized() public {
        vm.prank(user);
        vm.expectRevert();
        registry.pause();

        vm.prank(user);
        vm.expectRevert();
        registry.unpause();
    }

    function testActivateChain() public {
        uint256 chainId = 101;
        bytes memory registryAddress = abi.encodePacked(address(0x9876));

        // Verify chain not active initially
        uint256[] memory chains = registry.getActiveChains();
        assertEq(chains.length, 1);

        // Activate chain
        vm.prank(registryManager);
        vm.expectEmit(true, true, true, true);
        emit ChainStatusChanged(chainId, false, true);
        registry.changeChainStatus(chainId, address(gasZRC20), registryAddress, true);

        // Verify chain is now active
        chains = registry.getActiveChains();
        assertEq(chains.length, 2);
        assertEq(chains[1], chainId);
    }

    function testActivateChainUnauthorized() public {
        uint256 chainId = 101;
        bytes memory registryAddress = abi.encodePacked(address(0x9876));

        vm.prank(user);
        vm.expectRevert();
        registry.changeChainStatus(chainId, address(gasZRC20), registryAddress, true);
    }

    function testActivateAlreadyActiveChain() public {
        uint256 chainId = 101;
        bytes memory registryAddress = abi.encodePacked(address(0x9876));

        // Activate chain
        vm.prank(registryManager);
        registry.changeChainStatus(chainId, address(gasZRC20), registryAddress, true);

        // Try to activate chain
        vm.prank(registryManager);
        vm.expectRevert(abi.encodeWithSelector(ChainActive.selector, chainId));
        registry.changeChainStatus(chainId, address(gasZRC20), registryAddress, true);
    }

    function testDeactivateNonActiveChain() public {
        uint256 chainId = 101;
        bytes memory registryAddress = abi.encodePacked(address(0x9876));

        // Try to deactivate non-active chain
        vm.prank(registryManager);
        vm.expectRevert(abi.encodeWithSelector(ChainNonActive.selector, chainId));
        registry.changeChainStatus(chainId, address(gasZRC20), registryAddress, false);
    }

    function testDeactivateChain() public {
        uint256 chainId = 101;
        bytes memory registryAddress = abi.encodePacked(address(0x9876));

        // Activate chain
        vm.prank(registryManager);
        registry.changeChainStatus(chainId, address(gasZRC20), registryAddress, true);

        // Verify chain is active
        uint256[] memory chains = registry.getActiveChains();
        assertEq(chains.length, 2);

        // Deactivate chain
        vm.prank(registryManager);
        vm.expectEmit(true, true, true, true);
        emit ChainStatusChanged(chainId, true, false);
        registry.changeChainStatus(chainId, address(gasZRC20), registryAddress, false);

        // Verify chain is no longer active
        chains = registry.getActiveChains();
        assertEq(chains.length, 1);
    }

    function testUpdateChainMetadata() public {
        uint256 chainId = 101;
        bytes memory registryAddress = abi.encodePacked(address(0x9876));
        string memory key = "blockTime";
        bytes memory value = abi.encode(5);

        // Activate chain
        vm.prank(registryManager);
        registry.changeChainStatus(chainId, address(gasZRC20), registryAddress, true);

        // Update metadata
        vm.prank(registryManager);
        vm.expectEmit(true, true, true, true);
        emit ChainMetadataUpdated(chainId, key, value);
        registry.updateChainMetadata(chainId, key, value);

        // Verify metadata
        bytes memory storedValue = registry.getChainMetadata(chainId, key);
        assertEq(keccak256(storedValue), keccak256(value));
    }

    function testUpdateMetadataForNonActiveChain() public {
        uint256 chainId = 101;
        string memory key = "blockTime";
        bytes memory value = abi.encode(5);

        vm.prank(registryManager);
        vm.expectRevert(abi.encodeWithSelector(ChainNonActive.selector, chainId));
        registry.updateChainMetadata(chainId, key, value);
    }

    function testUpdateMetadataUnauthorized() public {
        uint256 chainId = 101;
        bytes memory registryAddress = abi.encodePacked(address(0x9876));
        string memory key = "blockTime";
        bytes memory value = abi.encode(5);

        // Activate chain
        vm.prank(registryManager);
        registry.changeChainStatus(chainId, address(gasZRC20), registryAddress, true);

        // Try to update metadata as unauthorized user
        vm.prank(user);
        vm.expectRevert();
        registry.updateChainMetadata(chainId, key, value);
    }

    function testRegisterContract() public {
        uint256 chainId = 101;
        bytes memory registryAddress = abi.encodePacked(address(0x9876));
        address contractAddress = address(0xAA55);
        string memory contractType = "connector";
        bytes memory addressBytes = abi.encodePacked(contractAddress);

        // Activate chain
        vm.prank(registryManager);
        registry.changeChainStatus(chainId, address(gasZRC20), registryAddress, true);

        // Registry contract
        vm.prank(registryManager);
        vm.expectEmit(true, true, true, true);
        emit ContractRegistered(chainId, contractType, addressBytes);
        registry.registerContract(chainId, contractAddress, contractType, addressBytes);

        // Verify contract info
        (bool active, address storedAddress) = registry.getContractInfo(chainId, contractType);
        assertTrue(active);
        assertEq(storedAddress, contractAddress);
    }

    function testRegisterContractForNonActiveChain() public {
        uint256 chainId = 101;
        address contractAddress = address(0xAA55);
        string memory contractType = "connector";
        bytes memory addressBytes = abi.encodePacked(contractAddress);

        vm.prank(registryManager);
        vm.expectRevert(abi.encodeWithSelector(ChainNonActive.selector, chainId));
        registry.registerContract(chainId, contractAddress, contractType, addressBytes);
    }

    function testRegisterContractWithEmptyType() public {
        uint256 chainId = 101;
        bytes memory registryAddress = abi.encodePacked(address(0x9876));
        address contractAddress = address(0xAA55);
        string memory contractType = "";
        bytes memory addressBytes = abi.encodePacked(contractAddress);

        // Activate chain
        vm.prank(registryManager);
        registry.changeChainStatus(chainId, address(gasZRC20), registryAddress, true);

        // Try to register contract with empty type
        vm.prank(registryManager);
        vm.expectRevert(abi.encodeWithSelector(InvalidContractType.selector, contractType));
        registry.registerContract(chainId, contractAddress, contractType, addressBytes);
    }

    function testRegisterContractWithEmptyAddress() public {
        uint256 chainId = 101;
        bytes memory registryAddress = abi.encodePacked(address(0x9876));
        address contractAddress = address(0xAA55);
        string memory contractType = "connector";
        bytes memory addressBytes = bytes("");

        // Activate chain
        vm.prank(registryManager);
        registry.changeChainStatus(chainId, address(gasZRC20), registryAddress, true);

        // Try to register contract with empty address bytes
        vm.prank(registryManager);
        vm.expectRevert(ZeroAddress.selector);
        registry.registerContract(chainId, contractAddress, contractType, addressBytes);
    }

    function testRegisterContractTwice() public {
        uint256 chainId = 101;
        bytes memory registryAddress = abi.encodePacked(address(0x9876));
        address contractAddress = address(0xAA55);
        string memory contractType = "connector";
        bytes memory addressBytes = abi.encodePacked(contractAddress);

        // Activate chain
        vm.prank(registryManager);
        registry.changeChainStatus(chainId, address(gasZRC20), registryAddress, true);

        // Register contract
        vm.prank(registryManager);
        registry.registerContract(chainId, contractAddress, contractType, addressBytes);

        // Try to register again
        vm.prank(registryManager);
        vm.expectRevert(abi.encodeWithSelector(ContractAlreadyRegistered.selector, chainId, contractType, addressBytes));
        registry.registerContract(chainId, contractAddress, contractType, addressBytes);
    }

    function testUpdateContractConfig() public {
        uint256 chainId = 101;
        bytes memory registryAddress = abi.encodePacked(address(0x9876));
        address contractAddress = address(0xAA55);
        string memory contractType = "connector";
        bytes memory addressBytes = abi.encodePacked(contractAddress);
        string memory key = "gasLimit";
        bytes memory value = abi.encode(300_000);

        // Activate chain
        vm.prank(registryManager);
        registry.changeChainStatus(chainId, address(gasZRC20), registryAddress, true);

        // Register contract
        vm.prank(registryManager);
        registry.registerContract(chainId, contractAddress, contractType, addressBytes);

        // Update contract config
        vm.prank(registryManager);
        vm.expectEmit(true, true, true, true);
        emit NewContractConfiguration(chainId, contractType, key, value);
        registry.updateContractConfiguration(chainId, contractType, key, value);

        // Verify config
        bytes memory storedValue = registry.getContractConfiguration(chainId, contractType, key);
        assertEq(keccak256(storedValue), keccak256(value));
    }

    function testUpdateConfigForNonExistentContract() public {
        uint256 chainId = 101;
        bytes memory registryAddress = abi.encodePacked(address(0x9876));
        string memory contractType = "nonexistent";
        string memory key = "gasLimit";
        bytes memory value = abi.encode(300_000);

        // Activate chain
        vm.prank(registryManager);
        registry.changeChainStatus(chainId, address(gasZRC20), registryAddress, true);

        // Try to update config for non-existent contract
        vm.prank(registryManager);
        vm.expectRevert(abi.encodeWithSelector(ContractNotFound.selector, chainId, contractType));
        registry.updateContractConfiguration(chainId, contractType, key, value);
    }

    function testSetContractActive() public {
        uint256 chainId = 101;
        bytes memory registryAddress = abi.encodePacked(address(0x9876));
        address contractAddress = address(0xAA55);
        string memory contractType = "connector";
        bytes memory addressBytes = abi.encodePacked(contractAddress);

        // Activate chain
        vm.prank(registryManager);
        registry.changeChainStatus(chainId, address(gasZRC20), registryAddress, true);

        // Register contract
        vm.prank(registryManager);
        registry.registerContract(chainId, contractAddress, contractType, addressBytes);

        // Check initial status
        (bool active,) = registry.getContractInfo(chainId, contractType);
        assertTrue(active);

        // Deactivate contract
        vm.prank(registryManager);
        vm.expectEmit(true, true, true, true);
        emit ContractStatusChanged(addressBytes);
        registry.setContractActive(chainId, contractType, false);

        // Verify status
        (active,) = registry.getContractInfo(chainId, contractType);
        assertFalse(active);

        // Reactivate contract
        vm.prank(registryManager);
        vm.expectEmit(true, true, true, true);
        emit ContractStatusChanged(addressBytes);
        registry.setContractActive(chainId, contractType, true);

        // Verify status
        (active,) = registry.getContractInfo(chainId, contractType);
        assertTrue(active);
    }

    function testSetStatusForNonExistenContract() public {
        uint256 chainId = 101;
        bytes memory registryAddress = abi.encodePacked(address(0x9876));
        string memory contractType = "nonexistent";

        // Activate chain
        vm.prank(registryManager);
        registry.changeChainStatus(chainId, address(gasZRC20), registryAddress, true);

        // Try to set status for non-existent contract
        vm.prank(registryManager);
        vm.expectRevert(abi.encodeWithSelector(ContractNotFound.selector, chainId, contractType));
        registry.setContractActive(chainId, contractType, false);
    }

    function testRegisterZRC20Token() public {
        address zrc20Address = address(0xDDDD);
        string memory symbol = "TKN";
        uint256 originChainId = 101;
        bytes memory originAddress = abi.encodePacked(address(0x1357));
        string memory coinType = "ERC20";
        uint8 decimals = 18;

        // Register ZRC20 token
        vm.prank(registryManager);
        vm.expectEmit(true, true, true, true);
        emit ZRC20TokenRegistered(originAddress, zrc20Address, decimals, originChainId, symbol);
        registry.registerZRC20Token(zrc20Address, symbol, originChainId, originAddress, coinType, decimals);

        // Verify token info
        (
            bool active,
            string memory storedSymbol,
            uint256 storedOriginChainId,
            bytes memory storedOriginAddress,
            string memory storedCoinType,
            uint8 storedDecimals
        ) = registry.getZRC20TokenInfo(zrc20Address);

        assertTrue(active);
        assertEq(storedSymbol, symbol);
        assertEq(storedOriginChainId, originChainId);
        assertEq(keccak256(storedOriginAddress), keccak256(originAddress));
        assertEq(storedCoinType, coinType);
        assertEq(storedDecimals, decimals);

        address mappedAddress = registry.getZRC20AddressByForeignAsset(originChainId, originAddress);
        assertEq(mappedAddress, zrc20Address);
    }

    function testRegisterZRC20TokenWithEmptyAddress() public {
        address zrc20Address = address(0);
        string memory symbol = "TKN";
        uint256 originChainId = 101;
        bytes memory originAddress = abi.encodePacked(address(0x1357));
        string memory coinType = "ERC20";
        uint8 decimals = 18;

        vm.prank(registryManager);
        vm.expectRevert(ZeroAddress.selector);
        registry.registerZRC20Token(zrc20Address, symbol, originChainId, originAddress, coinType, decimals);
    }

    function testRegisterZRC20TokenWithEmptySymbol() public {
        address zrc20Address = address(0xDDDD);
        string memory symbol = "";
        uint256 originChainId = 101;
        bytes memory originAddress = abi.encodePacked(address(0x1357));
        string memory coinType = "ERC20";
        uint8 decimals = 18;

        vm.prank(registryManager);
        vm.expectRevert(abi.encodeWithSelector(InvalidContractType.selector, "Symbol cannot be empty"));
        registry.registerZRC20Token(zrc20Address, symbol, originChainId, originAddress, coinType, decimals);
    }

    function testRegisterZRC20TokenWithEmptyOriginAddress() public {
        address zrc20Address = address(0xDDDD);
        string memory symbol = "TKN";
        uint256 originChainId = 101;
        bytes memory originAddress = bytes("");
        string memory coinType = "ERC20";
        uint8 decimals = 18;

        vm.prank(registryManager);
        vm.expectRevert(abi.encodeWithSelector(InvalidContractType.selector, "Origin address cannot be empty"));
        registry.registerZRC20Token(zrc20Address, symbol, originChainId, originAddress, coinType, decimals);
    }

    function testRegisterZRC20TokenTwice() public {
        address zrc20Address = address(0xDDDD);
        string memory symbol = "TKN";
        uint256 originChainId = 101;
        bytes memory originAddress = abi.encodePacked(address(0x1357));
        string memory coinType = "ERC20";
        uint8 decimals = 18;

        // Register ZRC20 token
        vm.prank(registryManager);
        registry.registerZRC20Token(zrc20Address, symbol, originChainId, originAddress, coinType, decimals);

        // Try to register the same address again
        vm.prank(registryManager);
        vm.expectRevert(abi.encodeWithSelector(ZRC20AlreadyRegistered.selector, zrc20Address));
        registry.registerZRC20Token(zrc20Address, "TKN2", originChainId, originAddress, coinType, decimals);
    }

    function testRegisterZRC20TokenWithDuplicateSymbol() public {
        address zrc20Address1 = address(0xDDDD);
        address zrc20Address2 = address(0xCCCC);
        string memory symbol = "TKN";
        uint256 originChainId = 101;
        bytes memory originAddress1 = abi.encodePacked(address(0x1357));
        bytes memory originAddress2 = abi.encodePacked(address(0xAAAA));
        string memory coinType = "ERC20";
        uint8 decimals = 18;

        // Register first ZRC20 token
        vm.prank(registryManager);
        registry.registerZRC20Token(zrc20Address1, symbol, originChainId, originAddress1, coinType, decimals);

        // Try to register another token with the same symbol
        vm.prank(registryManager);
        vm.expectRevert(abi.encodeWithSelector(ZRC20SymbolAlreadyInUse.selector, symbol));
        registry.registerZRC20Token(zrc20Address2, symbol, originChainId, originAddress2, coinType, decimals);
    }

    function testSetZRC20TokenActive() public {
        address zrc20Address = address(0xDDDD);
        string memory symbol = "TKN";
        uint256 originChainId = 101;
        bytes memory originAddress = abi.encodePacked(address(0x1357));
        string memory coinType = "ERC20";
        uint8 decimals = 18;

        // Register ZRC20 token
        vm.prank(registryManager);
        registry.registerZRC20Token(zrc20Address, symbol, originChainId, originAddress, coinType, decimals);

        // Check initial status
        (bool active,,,,,) = registry.getZRC20TokenInfo(zrc20Address);
        assertTrue(active);

        // Deactivate token
        vm.prank(registryManager);
        vm.expectEmit(true, true, true, true);
        emit ZRC20TokenUpdated(zrc20Address, false);
        registry.setZRC20TokenActive(zrc20Address, false);

        // Verify status
        (active,,,,,) = registry.getZRC20TokenInfo(zrc20Address);
        assertFalse(active);

        // Reactivate token
        vm.prank(registryManager);
        vm.expectEmit(true, true, true, true);
        emit ZRC20TokenUpdated(zrc20Address, true);
        registry.setZRC20TokenActive(zrc20Address, true);

        // Verify status
        (active,,,,,) = registry.getZRC20TokenInfo(zrc20Address);
        assertTrue(active);
    }

    function testUpdateNonExistentZRC20Token() public {
        address zrc20Address = address(0xDEAD);

        vm.prank(registryManager);
        vm.expectRevert(abi.encodeWithSelector(InvalidContractType.selector, "ZRC20 not registered"));
        registry.setZRC20TokenActive(zrc20Address, false);
    }

    function testSetZRC20TokenActiveWithZeroAddress() public {
        address zrc20Address = address(0);

        vm.prank(registryManager);
        vm.expectRevert(ZeroAddress.selector);
        registry.setZRC20TokenActive(zrc20Address, false);
    }

    function testWhenPaused() public {
        uint256 chainId = 101;
        bytes memory registryAddress = abi.encodePacked(address(0x9876));

        // Pause the contract
        vm.prank(admin);
        registry.pause();

        // Try to activate chain while paused
        vm.prank(registryManager);
        vm.expectRevert();
        registry.changeChainStatus(chainId, address(gasZRC20), registryAddress, true);

        // Unpause
        vm.prank(admin);
        registry.unpause();

        // Now it should work
        vm.prank(registryManager);
        registry.changeChainStatus(chainId, address(gasZRC20), registryAddress, true);
    }
}
