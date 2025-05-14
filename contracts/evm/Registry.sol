// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "./interfaces/IGatewayEVM.sol";
import "./interfaces/IRegistry.sol";

import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";

/// @title Registry
/// @notice Satellite registry contract for connected chains, receiving updates from CoreRegistry.
/// @dev This contract is deployed on every connected chain and maintains a synchronized view of the registry.
contract Registry is Initializable, UUPSUpgradeable, AccessControlUpgradeable, PausableUpgradeable, IRegistry {
    /// @notice New role identifier for pauser role
    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");
    /// @notice Identifier for the gateway role
    bytes32 public constant GATEWAY_ROLE = keccak256("GATEWAY_ROLE");

    /// @notice GatewayEVM contract that will call this contract with messages from CoreRegistry
    IGatewayEVM public gatewayEVM;
    /// @notice Represents the address of the CoreRegistry contract on the ZetaChain
    address public coreRegistry;
    /// @notice Active chains in the registry
    uint256[] public activeChains;
    /// @notice Maps chain IDs to their information
    mapping(uint256 => ChainInfo) private _chains;
    /// @notice Maps chain ID -> contract type -> ContractInfo
    mapping(uint256 => mapping(string => ContractInfo)) private _contracts;
    /// @notice Maps ZRC20 token address to their information
    mapping(address => ZRC20Info) private _zrc20Tokens;
    /// @notice Maps token symbol to ZRC20 address
    mapping(string => address) private _zrc20SymbolToAddress;
    /// @notice Maps origin chain ID and origin address to ZRC20 token address
    mapping(uint256 => mapping(bytes => address)) private _originAssetToZRC20;

    /// @dev Only registry address allowed modifier.
    /// @notice Restricts function calls to only be made by this contract itself
    /// @dev This is used to ensure functions receiving cross-chain messages can only be called through
    /// the onCall function using a self-call pattern, preventing direct external calls to these functions
    modifier onlyRegistry() {
        if (msg.sender != address(this)) {
            revert InvalidSender();
        }
        _;
    }

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    /// @notice Initialize the Registry contract
    /// @param admin_ Address with DEFAULT_ADMIN_ROLE, authorized for upgrades and pausing actions
    /// @param pauserAddress_ Address with PAUSER_ROLE, authorized for pausing actions
    /// @param gatewayEVM_ Address of the GatewayEVM contract for cross-chain messaging
    /// @param coreRegistry_ Address of the CoreRegistry contract deployed on ZetaChain
    function initialize(
        address admin_,
        address pauserAddress_,
        address gatewayEVM_,
        address coreRegistry_
    )
        public
        initializer
    {
        if (
            admin_ == address(0) || gatewayEVM_ == address(0) || coreRegistry_ == address(0)
                || pauserAddress_ == address(0)
        ) {
            revert ZeroAddress();
        }

        __UUPSUpgradeable_init();
        __AccessControl_init_unchained();
        __Pausable_init_unchained();

        _grantRole(DEFAULT_ADMIN_ROLE, admin_);
        _grantRole(PAUSER_ROLE, admin_);
        _grantRole(PAUSER_ROLE, pauserAddress_);
        _grantRole(GATEWAY_ROLE, gatewayEVM_);

        gatewayEVM = IGatewayEVM(gatewayEVM_);
        coreRegistry = coreRegistry_;
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

    /// @notice onCall is called by the GatewayEVM when a cross-chain message is received
    /// @param context Information about the cross-chain message
    /// @param data The encoded function call to execute
    function onCall(
        MessageContext calldata context,
        bytes calldata data
    )
        external
        onlyRole(GATEWAY_ROLE)
        whenNotPaused
        returns (bytes memory)
    {
        if (context.sender != coreRegistry) {
            revert InvalidSender();
        }

        // Execute a self-call with the encoded function data
        // This allows us to reuse existing function signatures and validation logic
        // By using a self-call pattern combined with onlyRegistry modifier, these functions
        // can only be called through this onCall entry point
        (bool success, bytes memory result) = address(this).call(data);
        if (!success) {
            // Extract the error message from the revert
            assembly {
                revert(add(result, 32), mload(result))
            }
        }

        return result;
    }

    /// @notice Changes status of the chain to activated/deactivated
    /// @dev Only callable through onCall from CoreRegistry
    /// @param chainId The ID of the chain being activated/deactivated.
    /// @param gasZRC20 The address of the ZRC20 token that represents gas token for the chain.
    /// @param registry Address of the Registry contract on the connected chain.
    /// @param active Whether activate or deactivate the chain
    function changeChainStatus(
        uint256 chainId,
        address gasZRC20,
        bytes calldata registry,
        bool active
    )
        external
        onlyRegistry
        whenNotPaused
    {
        // Update the chain info
        _chains[chainId].gasZRC20 = gasZRC20;
        _chains[chainId].registry = registry;
        _chains[chainId].active = active;

        // Update active chains array
        if (active) {
            activeChains.push(chainId);
        } else {
            // Remove from active chains
            _removeFromActiveChains(chainId);
        }

        emit ChainStatusChanged(chainId, active);
    }

    /// @notice Updates chain metadata, only for the active chains
    /// @dev Only callable through onCall from CoreRegistry
    /// @param chainId The ID of the chain
    /// @param key The metadata key to update
    /// @param value The new value for the metadata
    function updateChainMetadata(
        uint256 chainId,
        string calldata key,
        bytes calldata value
    )
        external
        onlyRegistry
        whenNotPaused
    {
        // Updates chain metadata
        _chains[chainId].metadata[key] = value;

        emit ChainMetadataUpdated(chainId, key, value);
    }

    /// @notice Registers a new contract address for a specific chain
    /// @dev Only callable through onCall from CoreRegistry
    /// @param chainId The ID of the chain where the contract is deployed
    /// @param address_ The address of the contract
    /// @param contractType The type of the contract (e.g., "connector", "gateway")
    /// @param addressBytes The bytes representation of the non-EVM address
    function registerContract(
        uint256 chainId,
        address address_,
        string calldata contractType,
        bytes calldata addressBytes
    )
        external
        onlyRegistry
        whenNotPaused
    {
        if (bytes(contractType).length == 0) revert InvalidContractType(contractType);
        if (bytes(addressBytes).length == 0) revert ZeroAddress();

        // Store contract info in the storage
        _contracts[chainId][contractType].active = true;
        _contracts[chainId][contractType].address_ = address_;
        _contracts[chainId][contractType].addressBytes = addressBytes;
        _contracts[chainId][contractType].contractType = contractType;

        emit ContractRegistered(chainId, contractType, addressBytes);
    }

    /// @notice Updates contract configuration
    /// @dev Only callable through onCall from CoreRegistry
    /// @param chainId The ID of the chain where the contract is deployed
    /// @param contractType The type of the contract
    /// @param key The configuration key to update
    /// @param value The new value for the configuration
    function updateContractConfiguration(
        uint256 chainId,
        string calldata contractType,
        string calldata key,
        bytes calldata value
    )
        external
        onlyRegistry
        whenNotPaused
    {
        if (bytes(contractType).length == 0) revert InvalidContractType(contractType);

        // Store new configuration in the storage
        _contracts[chainId][contractType].configuration[key] = value;

        emit ContractConfigurationUpdated(chainId, contractType, key, value);
    }

    /// @notice Sets a contract's active status
    /// @dev Only callable through onCall from CoreRegistry
    /// @param chainId The ID of the chain where the contract is deployed
    /// @param contractType The type of the contract
    /// @param active Whether the contract should be active
    function setContractActive(uint256 chainId, string calldata contractType, bool active) external onlyRegistry {
        if (bytes(contractType).length == 0) revert InvalidContractType(contractType);

        // Update the active status
        _contracts[chainId][contractType].active = active;

        emit ContractStatusChanged(_contracts[chainId][contractType].addressBytes);
    }

    /// @notice Registers a new ZRC20 token in the registry
    /// @dev Only callable through onCall from CoreRegistry
    /// @param address_ The address of the ZRC20 token on ZetaChain
    /// @param symbol The symbol of the token
    /// @param originChainId The ID of the foreign chain where the original asset exists
    /// @param originAddress The address or identifier of the asset on its native chain
    /// @param coinType The type of the original coin
    /// @param decimals The number of decimals the token uses
    function registerZRC20Token(
        address address_,
        string calldata symbol,
        uint256 originChainId,
        bytes calldata originAddress,
        string calldata coinType,
        uint8 decimals
    )
        external
        onlyRegistry
        whenNotPaused
    {
        if (address_ == address(0)) revert ZeroAddress();
        if (bytes(symbol).length == 0) revert InvalidContractType("Symbol cannot be empty");
        if (bytes(originAddress).length == 0) revert InvalidContractType("Origin address cannot be empty");

        // Register the ZRC20 token info
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

        emit ZRC20TokenRegistered(originAddress, address_, decimals, originChainId, symbol);
    }

    /// @notice Updates ZRC20 token active status
    /// @dev Only callable through onCall from CoreRegistry
    /// @param address_ The address of the ZRC20 token
    /// @param active Whether the token should be active
    function setZRC20TokenActive(address address_, bool active) external onlyRegistry whenNotPaused {
        if (address_ == address(0)) revert ZeroAddress();

        // Update token status
        _zrc20Tokens[address_].active = active;

        emit ZRC20TokenUpdated(address_, active);
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
    /// @return address_ The address of the contract
    function getContractInfo(
        uint256 chainId,
        string calldata contractType
    )
        external
        view
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

    /// @notice Gets the ZRC20 token address for a specific asset on a foreign chain
    /// @param originChainId The ID of the foreign chain
    /// @param originAddress The address or identifier of the asset on its native chain
    /// @return The address of the corresponding ZRC20 token on ZetaChain
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

    /// @notice Gets all active chains in the registry
    /// @return Array of chain IDs for all active chains
    function getActiveChains() external view returns (uint256[] memory) {
        return activeChains;
    }

    /// @notice Removes a chain ID from the active chains array
    /// @param chainId The ID of the chain to remove
    function _removeFromActiveChains(uint256 chainId) private {
        for (uint256 i = 0; i < activeChains.length; i++) {
            if (activeChains[i] == chainId) {
                // Swap with the last element and pop
                activeChains[i] = activeChains[activeChains.length - 1];
                activeChains.pop();
                break;
            }
        }
    }
}
