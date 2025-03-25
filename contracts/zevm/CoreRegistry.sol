// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "./interfaces/ICoreRegistry.sol";
import "./interfaces/IGatewayZEVM.sol";

import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";

/// @notice CoreRegistry
/// @notice Central registry for ZetaChain, managing chain info, ZRC20 data, and contract addresses across all chains.
/// @dev The contract doesn't hold any funds and should never have active allowances.
contract CoreRegistry is
    Initializable,
    UUPSUpgradeable,
    AccessControlUpgradeable,
    PausableUpgradeable,
    ICoreRegistry
{
    /// @notice New role identifier for registry manager role.
    bytes32 public constant REGISTRY_MANAGER_ROLE = keccak256("REGISTRY_MANAGER_ROLE");
    /// @notice New role identifier for pauser role.
    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");
    /// @notice Cross-chain message gas limit
    uint256 public constant CROSS_CHAIN_GAS_LIMIT = 300_000;

    /// @notice Instance of the GatewayZEVM contract for cross-chain communication
    IGatewayZEVM public gatewayZEVM;
    /// @notice Active chains in the registry
    uint256[] public _activeChains;
    /// @notice Maps chain IDs to their information.
    mapping(uint256 => ChainInfo) private _chains;
    /// @notice Maps chain ID -> contract type -> ContractInfo
    mapping(uint256 => mapping(string => ContractInfo)) private _contracts;
    /// @notice Maps ZRC20 token address to their information
    mapping(address => ZRC20Info) private _zrc20Tokens;
    /// @notice Maps token symbol to ZRC20 address.
    mapping(string => address) private _zrc20SymbolToAddress;
    /// @notice Maps origin chain ID and origin address to ZRC20 token address.
    mapping(uint256 => mapping(bytes => address)) private _originAssetToZRC20;

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    /// @notice Initialize the CoreRegistry contract.
    /// @param admin_ Address with DEFAULT_ADMIN_ROLE, authorized for upgrades and pausing actions.
    /// @param registryManager_ Address with REGISTRY_MANAGER_ROLE, authorized for all registry write actions.
    /// @param gatewayZEVM_ Address of the GatewayZEVM contract for cross-chain messaging
    function initialize(address admin_, address registryManager_, address gatewayZEVM_) public initializer {
        if (admin_ == address(0) || registryManager_ == address(0) || gatewayZEVM_ == address(0)) {
            revert ZeroAddress();
        }

        __UUPSUpgradeable_init();
        __AccessControl_init_unchained();
        __Pausable_init_unchained();

        _grantRole(DEFAULT_ADMIN_ROLE, admin_);
        _grantRole(REGISTRY_MANAGER_ROLE, registryManager_);
        _grantRole(PAUSER_ROLE, registryManager_);
        _grantRole(PAUSER_ROLE, admin_);

        gatewayZEVM = IGatewayZEVM(gatewayZEVM_);
    }

    /// @dev Authorizes the upgrade of the contract, sender must be admin.
    /// @param newImplementation Address of the new implementation,
    function _authorizeUpgrade(address newImplementation) internal override onlyRole(DEFAULT_ADMIN_ROLE) { }

    /// @notice Pause contract.
    function pause() external onlyRole(PAUSER_ROLE) {
        _pause();
    }

    /// @notice Unpause contract.
    function unpause() external onlyRole(PAUSER_ROLE) {
        _unpause();
    }

    /// @notice Changes status of the chain to activated/deactivated.
    /// @param chainId The ID of the chain to activate.
    /// @param activation Whether activate or deactivate the chain
    function chainActivation(
        uint256 chainId,
        bytes calldata registry,
        bool activation
    )
        external
        override
        onlyRole(REGISTRY_MANAGER_ROLE)
        whenNotPaused
    {
        if (registry.length == 0) revert ZeroAddress();
        // In the case chain is already activated
        if (_chains[chainId].active && activation) revert ChainAlreadyActive(chainId);
        // In the case chain is inactive
        if (!_chains[chainId].active && !activation) revert ChainNonActive(chainId);

        // Update the chain info
        _chains[chainId].active = activation;
        _chains[chainId].registry = registry;

        // Update active chains array
        if (activation) {
            _activeChains.push(chainId);
        } else {
            _removeFromActiveChains(chainId);
        }

        // Broadcast update to satellite registries
        _broadcastChainActivation(chainId, activation);

        emit ChainStatusChanged(chainId);
    }

    /// @notice Updates chain metadata.
    /// @param chainId The ID of the chain.
    /// @param key The metadata key to update.
    /// @param value The new value for the metadata.
    function updateChainMetadata(
        uint256 chainId,
        string calldata key,
        bytes calldata value
    )
        external
        override
        onlyRole(REGISTRY_MANAGER_ROLE)
        whenNotPaused
    {
        // Check if the chain is active
        if (!_chains[chainId].active) revert ChainNonActive(chainId);

        // Updates chain metadata
        _chains[chainId].metadata[key] = value;

        // Broadcast update to satellite registries
        _broadcastChainMetadataUpdate(chainId, key, value);

        emit NewChainMetadata(chainId, key, value);
    }

    /// @notice Registers a new contract address for a specific chain.
    /// @param chainId The ID of the chain where the contract is deployed.
    /// @param address_ The address of the contract.
    /// @param contractType The type of the contract (e.g., "connector", "gateway").
    /// @param addressBytes The bytes representation of the non-EVM address.
    function registerContract(
        uint256 chainId,
        address address_,
        string calldata contractType,
        bytes calldata addressBytes
    )
        external
        override
        onlyRole(REGISTRY_MANAGER_ROLE)
        whenNotPaused
    {
        // Validate inputs
        if (!_chains[chainId].active) revert ChainNonActive(chainId);
        if (bytes(contractType).length == 0) revert InvalidContractType(contractType);
        if (address_ == address(0) && bytes(addressBytes).length == 0) {
            revert ZeroAddress();
        }

        // Check if contract already exists in the registry
        if (bytes(_contracts[chainId][contractType].addressBytes).length > 0) {
            revert ContractAlreadyRegistered(chainId, contractType, addressBytes);
        }

        // Store contract info in the storage.
        _contracts[chainId][contractType].active = true;
        _contracts[chainId][contractType].address_ = address_;
        _contracts[chainId][contractType].addressBytes = addressBytes;
        _contracts[chainId][contractType].contractType = contractType;

        // Broadcast update to satellite registries
        _broadcastContractRegistration(chainId, address_, contractType, addressBytes);

        emit ContractRegistered(chainId, contractType, addressBytes);
    }

    /// @notice Updates contract configuration.
    /// @param chainId The ID of the chain where the contract is deployed.
    /// @param contractType The type of the contract.
    /// @param key The configuration key to update.
    /// @param value The new value for the configuration.
    function updateContractConfiguration(
        uint256 chainId,
        string calldata contractType,
        string calldata key,
        bytes calldata value
    )
        external
        override
        onlyRole(REGISTRY_MANAGER_ROLE)
        whenNotPaused
    {
        // Validate inputs
        if (!_chains[chainId].active) revert ChainNonActive(chainId);
        if (bytes(contractType).length == 0) revert InvalidContractType(contractType);

        // Check if contract exists in the registry
        if (
            bytes(_contracts[chainId][contractType].addressBytes).length == 0
                && _contracts[chainId][contractType].active
        ) {
            revert ContractNotFound(chainId, contractType);
        }

        // Store new configuration in the storage.
        _contracts[chainId][contractType].configuration[key] = value;

        // Broadcast update to satellite registries
        _broadcastContractConfigUpdate(chainId, contractType, key, value);

        emit NewContractConfiguration(chainId, contractType, key, value);
    }

    /// @notice Sets a contract's active status
    /// @param chainId The ID of the chain where the contract is deployed.
    /// @param contractType The type of the contract.
    /// @param active Whether the contract should be active.
    function setContractActive(
        uint256 chainId,
        string calldata contractType,
        bool active
    )
        external
        override
        onlyRole(REGISTRY_MANAGER_ROLE)
        whenNotPaused
    {
        // Validate inputs
        if (!_chains[chainId].active) revert ChainNonActive(chainId);
        if (bytes(contractType).length == 0) revert InvalidContractType(contractType);

        // Check if contract exists in the registry
        if (
            bytes(_contracts[chainId][contractType].addressBytes).length == 0
                && _contracts[chainId][contractType].active
        ) {
            revert ContractNotFound(chainId, contractType);
        }

        // Update the active status
        _contracts[chainId][contractType].active = active;

        // Broadcast update to satellite registries
        _broadcastContractStatusUpdate(chainId, contractType, active);

        emit ContractStatusChanged(_contracts[chainId][contractType].addressBytes);
    }

    /// @notice Registers a new ZRC20 token in the registry.
    /// @param address_ The address of the ZRC20 token on ZetaChain.
    /// @param symbol The symbol of the token.
    /// @param originChainId The ID of the foreign chain where the original asset exists.
    /// @param originAddress The address or identifier of the asset on its native chain.
    /// @param coinType The type of the original coin.
    /// @param decimals The number of decimals the token uses.
    function registerZRC20Token(
        address address_,
        string calldata symbol,
        uint256 originChainId,
        bytes calldata originAddress,
        string calldata coinType,
        uint8 decimals
    )
        external
        override
        onlyRole(REGISTRY_MANAGER_ROLE)
        whenNotPaused
    {
        // Validate inputs
        if (address_ == address(0)) revert ZeroAddress();
        if (bytes(symbol).length == 0) revert InvalidContractType("Symbol cannot be empty");
        if (bytes(originAddress).length == 0) revert InvalidContractType("Origin address cannot be empty");

        // Check if token already registered
        if (_zrc20Tokens[address_].address_ != address(0)) revert ZRC20AlreadyRegistered(address_);
        if (_zrc20SymbolToAddress[symbol] != address(0)) revert ZRC20SymbolAlreadyInUse(symbol);

        // Register the ZRC20 token info.
        _zrc20Tokens[address_].active = true;
        _zrc20Tokens[address_].address_ = address_;
        _zrc20Tokens[address_].originAddress = originAddress;
        _zrc20Tokens[address_].originChainId = originChainId;
        _zrc20Tokens[address_].symbol = symbol;
        _zrc20Tokens[address_].decimals = decimals;
        _zrc20Tokens[address_].coinType = coinType;

        // Map foreign asset to ZRC20
        _originAssetToZRC20[originChainId][originAddress] = address_;
        // Map symbol to address
        _zrc20SymbolToAddress[symbol] = address_;

        // Broadcast update to satellite registries
        _broadcastZRC20Registration(address_, symbol, originChainId, originAddress, coinType, decimals);

        emit ZRC20TokenRegistered(originAddress, address_, decimals, originChainId, symbol);
    }

    /// @notice Updates ZRC20 token active status.
    function updateZRC20Token(
        address address_,
        bool active
    )
        external
        override
        onlyRole(REGISTRY_MANAGER_ROLE)
        whenNotPaused
    {
        // Validate inputs
        if (address_ == address(0)) revert ZeroAddress();
        if (_zrc20Tokens[address_].address_ == address(0)) revert InvalidContractType("ZRC20 not registered");

        // Update token status.
        _zrc20Tokens[address_].active = active;

        // Broadcast update to satellite registries
        _broadcastZRC20Update(address_, active);

        emit ZRC20TokenUpdated(address_, active);
    }

    /// @notice Gets chain-specific metadata
    /// @param chainId The ID of the chain
    /// @param key The metadata key to retrieve
    /// @return The value of the requested metadata
    function getChainMetadata(uint256 chainId, string calldata key) external view override returns (bytes memory) {
        return _chains[chainId].metadata[key];
    }

    /// @notice Gets information about a specific contract
    /// @param chainId The ID of the chain where the contract is deployed
    /// @param contractType The type of the contract
    /// @return active Whether the contract is active
    /// @return address_ The address of the contract
    function getContractInfo(
        uint256 chainId,
        string calldata contractType
    )
        external
        view
        override
        returns (bool active, address address_)
    {
        active = _contracts[chainId][contractType].active;
        address_ = _contracts[chainId][contractType].address_;
    }

    /// @notice Gets contract-specific configuration
    /// @param chainId The ID of the chain where the contract is deployed
    /// @param contractType The type of the contract
    /// @param key The configuration key to retrieve
    /// @return The value of the requested configuration
    function getContractConfiguration(
        uint256 chainId,
        string calldata contractType,
        string calldata key
    )
        external
        view
        override
        returns (bytes memory)
    {
        return _contracts[chainId][contractType].configuration[key];
    }

    /// @notice Gets information about a specific ZRC20 token
    /// @param address_ The address of the ZRC20 token
    /// @return active Whether the token is active
    /// @return symbol The symbol of the token
    /// @return originChainId The ID of the foreign chain where the original asset exists
    /// @return originAddress The address or identifier of the asset on its native chain
    /// @return coinType The type of the original coin
    /// @return decimals The number of decimals the token uses
    function getZRC20TokenInfo(address address_)
        external
        view
        override
        returns (
            bool active,
            string memory symbol,
            uint256 originChainId,
            bytes memory originAddress,
            string memory coinType,
            uint8 decimals
        )
    {
        ZRC20Info memory info = _zrc20Tokens[address_];
        active = info.active;
        symbol = info.symbol;
        originChainId = info.originChainId;
        originAddress = info.originAddress;
        coinType = info.coinType;
        decimals = info.decimals;
    }

    /// @notice Gets the ZRC20 token address for a specific asset on a foreign chain.
    /// @param originChainId The ID of the foreign chain
    /// @param originAddress The address or identifier of the asset on its native chain.
    /// @return The address of the corresponding ZRC20 token on ZetaChain.
    function getZRC20AddressByForeignAsset(
        uint256 originChainId,
        bytes calldata originAddress
    )
        external
        view
        override
        returns (address)
    {
        return _originAssetToZRC20[originChainId][originAddress];
    }

    /// @notice Gets all active chains in the registry.
    /// @return Array of chain IDs for all active chains.
    function getActiveChains() external view override returns (uint256[] memory) {
        return _activeChains;
    }

    /// @notice Broadcast chain activation update to all satellite registries.
    /// @param chainId The ID of the chain being activated/deactivated
    /// @param activation Whether the chain is being activated or deactivated
    function _broadcastChainActivation(uint256 chainId, bool activation) internal {
        // Encode the function call for the Registry contract on the target chain
        bytes memory message = abi.encodeWithSignature("updateChainActivation(uint256,bool)", chainId, activation);
        _broadcastToAllChains(message);
    }

    /// @notice Broadcast chain metadata to all satellite registries
    /// @param chainId The ID of the chain whose metadata is being updated
    /// @param key The metadata key being updated
    /// @param value The new value for the metadata
    function _broadcastChainMetadataUpdate(uint256 chainId, string calldata key, bytes calldata value) private {
        // Encode the function call for the Registry contract on the target chain
        bytes memory message = abi.encodeWithSignature("updateChainMetadata(uint256,string,bytes)", chainId, key, value);
        _broadcastToAllChains(message);
    }

    /// @notice Broadcast contract registration to all satellite registries
    /// @param chainId The ID of the chain where the contract is deployed
    /// @notice address_ The address of the contract
    /// @notice contractType The type of the contract
    /// @notice addressBytes The bytes representation of the non-EVM address
    function _broadcastContractRegistration(
        uint256 chainId,
        address address_,
        string calldata contractType,
        bytes calldata addressBytes
    )
        private
    {
        bytes memory message = abi.encodeWithSignature(
            "registerContract(uint256,address,string,bytes)", chainId, address_, contractType, addressBytes
        );
        _broadcastToAllChains(message);
    }

    /// @notice Broadcast contract configuration update to all satellite registries
    /// @param chainId The ID of the chain where the contract is deployed
    /// @notice contractType The type of the contract
    /// @notice key The configuration key being updated
    /// @notice value The new value for the configuration
    function _broadcastContractConfigUpdate(
        uint256 chainId,
        string calldata contractType,
        string calldata key,
        bytes calldata value
    )
        private
    {
        // Encode the function call for the Registry contract on the target chain
        bytes memory message = abi.encodeWithSignature(
            "updateContractConfiguration(uint256,string,string,bytes)", chainId, contractType, key, value
        );
        _broadcastToAllChains(message);
    }

    /// @notice Broadcast contract status update to all satellite registries
    /// @param chainId The ID of the chain where the contract is deployed
    /// @notice contractType The type of the contract
    /// @notice active Whether the contract should be active
    function _broadcastContractStatusUpdate(uint256 chainId, string calldata contractType, bool active) private {
        // Encode the function call for the Registry contract on the target chain
        bytes memory message =
            abi.encodeWithSignature("setContractActive(uint256,string,bool)", chainId, contractType, active);
        _broadcastToAllChains(message);
    }

    /// @notice Broadcast ZRC20 token registration to all satellite registries
    /// @param address_ The address of the ZRC20 token on ZetaChain
    /// @param symbol The symbol of the token
    /// @param originChainId The ID of the foreign chain where the original asset exists
    /// @param originAddress The address or identifier of the asset on its native chain
    /// @param coinType The type of the original coin
    /// @param decimals The number of decimals the token uses
    function _broadcastZRC20Registration(
        address address_,
        string calldata symbol,
        uint256 originChainId,
        bytes calldata originAddress,
        string calldata coinType,
        uint8 decimals
    )
        private
    {
        // Encode the function call for the Registry contract on the target chain
        bytes memory message = abi.encodeWithSignature(
            "registerZRC20Token(address,string,uint256,bytes,string,uint8)",
            address_,
            symbol,
            originChainId,
            originAddress,
            coinType,
            decimals
        );
        _broadcastToAllChains(message);
    }

    /// @notice Broadcast ZRC20 token update to all satellite registries
    /// @param address_ The address of the ZRC20 token
    /// @param active Whether the token should be active
    function _broadcastZRC20Update(address address_, bool active) private {
        // Encode the function call for the Registry contract on the target chain
        bytes memory message = abi.encodeWithSignature("updateZRC20Token(address,bool)", address_, active);
        _broadcastToAllChains(message);
    }

    /// @notice Generic function to broadcast encoded messages to all satellite registries
    /// @param encodedMessage The fully encoded function call to broadcast
    function _broadcastToAllChains(bytes memory encodedMessage) private {
        for (uint256 i = 0; i < _activeChains.length; i++) {
            _sendCrossChainMessage(_activeChains[i], encodedMessage);
        }
    }

    /// @notice Sends a cross-chain message to the Registry contract on a target chain.
    /// @param targetChainId The ID of the chain to send the message to.
    /// @param message The encoded function call to execute on the target chain.
    function _sendCrossChainMessage(uint256 targetChainId, bytes memory message) private {
        // Prepare call options
        CallOptions memory callOptions = CallOptions({ gasLimit: CROSS_CHAIN_GAS_LIMIT, isArbitraryCall: true });

        // Prepare revert options
        // TODO: fill with info
        RevertOptions memory revertOptions;

        gatewayZEVM.call(
            _chains[targetChainId].registry,
            // TODO: Add ZRC20 contract address
            address(1),
            message,
            callOptions,
            revertOptions
        );
    }

    /// @notice Removes a chain ID from the active chains array.
    /// @param chainId The ID of the chain to remove.
    function _removeFromActiveChains(uint256 chainId) private {
        for (uint256 i = 0; i < _activeChains.length; i++) {
            if (_activeChains[i] == chainId) {
                // Swap with the last element and pop
                _activeChains[i] = _activeChains[_activeChains.length - 1];
                _activeChains.pop();
                break;
            }
        }
    }
}
