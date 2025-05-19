// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "./interfaces/IBaseRegistry.sol";

import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";

abstract contract BaseRegistry is
    Initializable,
    UUPSUpgradeable,
    AccessControlUpgradeable,
    PausableUpgradeable,
    IBaseRegistry
{
    /// @notice New role identifier for pauser role.
    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");

    /// @notice Active chains in the registry.
    uint256[] internal _activeChains;
    /// @notice Maps chain IDs to their information.
    mapping(uint256 => ChainInfo) internal _chains;
    /// @notice Maps chain ID -> contract type -> ContractInfo
    mapping(uint256 => mapping(string => ContractInfo)) internal _contracts;
    /// @notice Maps ZRC20 token address to their information
    mapping(address => ZRC20Info) internal _zrc20Tokens;
    /// @notice Maps token symbol to ZRC20 address.
    mapping(string => address) internal _zrc20SymbolToAddress;
    /// @notice Maps origin chain ID and origin address to ZRC20 token address.
    mapping(uint256 => mapping(bytes => address)) internal _originAssetToZRC20;

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    /// @dev Authorizes the upgrade of the contract, sender must be admin.
    /// @param newImplementation Address of the new implementation,
    function _authorizeUpgrade(address newImplementation) internal override onlyRole(DEFAULT_ADMIN_ROLE) { }

    /// @notice Pause contract.
    function pause() external onlyRole(PAUSER_ROLE) {
        _pause();
    }

    /// @notice Unpause contract.
    function unpause() external onlyRole(DEFAULT_ADMIN_ROLE) {
        _unpause();
    }

    /// @notice Changes status of the chain to activated/deactivated.
    /// @param chainId The ID of the chain to activate.
    /// @param gasZRC20 The address of the ZRC20 token that represents gas token for the chain.
    /// @param registry Address of the Registry contract on the connected chain.
    /// @param activation Whether activate or deactivate the chain
    function _changeChainStatus(uint256 chainId, address gasZRC20, bytes calldata registry, bool activation) internal {
        if (registry.length == 0) revert ZeroAddress();
        // In the case chain is already activated
        if (_chains[chainId].active && activation) revert ChainActive(chainId);
        // In the case chain is inactive
        if (!_chains[chainId].active && !activation) revert ChainNonActive(chainId);

        // Update the chain info
        _chains[chainId].active = activation;
        _chains[chainId].gasZRC20 = gasZRC20;
        _chains[chainId].registry = registry;

        // Update active chains array
        if (activation) {
            _activeChains.push(chainId);
        } else {
            _removeFromActiveChains(chainId);
        }
    }

    /// @notice Updates chain metadata, only for the active chains.
    /// @param chainId The ID of the chain.
    /// @param key The metadata key to update.
    /// @param value The new value for the metadata.
    function _updateChainMetadata(uint256 chainId, string calldata key, bytes calldata value) internal {
        // Check if the chain is active
        if (!_chains[chainId].active) revert ChainNonActive(chainId);

        // Updates chain metadata
        _chains[chainId].metadata[key] = value;
    }

    /// @notice Registers a new contract address for a specific chain.
    /// @param chainId The ID of the chain where the contract is deployed.
    /// @param contractType The type of the contract (e.g., "connector", "gateway").
    /// @param addressBytes The bytes representation of the non-EVM address.
    function _registerContract(uint256 chainId, string calldata contractType, bytes calldata addressBytes) internal {
        // Validate inputs
        if (!_chains[chainId].active) revert ChainNonActive(chainId);
        if (bytes(contractType).length == 0) revert InvalidContractType(contractType);
        if (bytes(addressBytes).length == 0) revert ZeroAddress();

        // Check if contract already exists in the registry
        if (bytes(_contracts[chainId][contractType].addressBytes).length > 0) {
            revert ContractAlreadyRegistered(chainId, contractType, addressBytes);
        }

        // Store contract info in the storage.
        _contracts[chainId][contractType].active = true;
        _contracts[chainId][contractType].addressBytes = addressBytes;
        _contracts[chainId][contractType].contractType = contractType;

        emit ContractRegistered(chainId, contractType, addressBytes);
    }

    /// @notice Updates contract configuration.
    /// @param chainId The ID of the chain where the contract is deployed.
    /// @param contractType The type of the contract.
    /// @param key The configuration key to update.
    /// @param value The new value for the configuration.
    function _updateContractConfiguration(
        uint256 chainId,
        string calldata contractType,
        string calldata key,
        bytes calldata value
    )
        internal
    {
        // Validate inputs
        if (!_chains[chainId].active) revert ChainNonActive(chainId);
        if (bytes(contractType).length == 0) revert InvalidContractType(contractType);

        // Check if contract exists in the registry
        if (!_contracts[chainId][contractType].active) {
            revert ContractNotFound(chainId, contractType);
        }

        // Store new configuration in the storage.
        _contracts[chainId][contractType].configuration[key] = value;
    }

    /// @notice Sets a contract's active status
    /// @param chainId The ID of the chain where the contract is deployed.
    /// @param contractType The type of the contract.
    /// @param active Whether the contract should be active.
    function _setContractActive(uint256 chainId, string calldata contractType, bool active) internal {
        // Validate inputs
        if (!_chains[chainId].active) revert ChainNonActive(chainId);
        if (bytes(contractType).length == 0) revert InvalidContractType(contractType);

        // Check if contract exists in the registry
        if (bytes(_contracts[chainId][contractType].addressBytes).length == 0) {
            revert ContractNotFound(chainId, contractType);
        }

        // Update the active status
        _contracts[chainId][contractType].active = active;
    }

    /// @notice Registers a new ZRC20 token in the registry.
    /// @param address_ The address of the ZRC20 token on ZetaChain.
    /// @param symbol The symbol of the token.
    /// @param originChainId The ID of the foreign chain where the original asset exists.
    /// @param originAddress The address or identifier of the asset on its native chain.
    /// @param coinType The type of the original coin.
    /// @param decimals The number of decimals the token uses.
    function _registerZRC20Token(
        address address_,
        string calldata symbol,
        uint256 originChainId,
        bytes calldata originAddress,
        string calldata coinType,
        uint8 decimals
    )
        internal
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
    }

    /// @notice Updates ZRC20 token active status.
    function _setZRC20TokenActive(address address_, bool active) internal {
        // Validate inputs
        if (address_ == address(0)) revert ZeroAddress();
        if (_zrc20Tokens[address_].address_ == address(0)) revert InvalidContractType("ZRC20 not registered");

        // Update token status.
        _zrc20Tokens[address_].active = active;
    }

    /// @notice Gets chain-specific metadata
    /// @param chainId The ID of the chain
    /// @param key The metadata key to retrieve
    /// @return The value of the requested metadata
    function getChainMetadata(uint256 chainId, string calldata key) external view returns (bytes memory) {
        return _chains[chainId].metadata[key];
    }

    /// @notice Gets information about a specific contract
    /// @param chainId The ID of the chain where the contract is deployed
    /// @param contractType The type of the contract
    /// @return active Whether the contract is active
    /// @return addressBytes The address of the contract
    function getContractInfo(
        uint256 chainId,
        string calldata contractType
    )
        external
        view
        returns (bool active, bytes memory addressBytes)
    {
        active = _contracts[chainId][contractType].active;
        addressBytes = _contracts[chainId][contractType].addressBytes;
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
        returns (address)
    {
        return _originAssetToZRC20[originChainId][originAddress];
    }

    /// @notice Gets all active chains in the registry.
    /// @return Array of chain IDs for all active chains.
    function getActiveChains() external view returns (uint256[] memory) {
        return _activeChains;
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
