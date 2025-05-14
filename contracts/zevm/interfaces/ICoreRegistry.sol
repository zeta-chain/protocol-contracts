// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

/// @title ICoreRegistryEvents
/// @notice Interface for the events emitted by the CoreRegistry contract.
interface ICoreRegistryEvents {
    /// @notice Emitted when a new contract address is registered.
    /// @param chainId The ID of the chain where the contract is deployed.
    /// @param contractType The type of the contract (e.g. "connector", "tss", "gateway")
    /// @param addressBytes The bytes representation of the registered address
    event ContractRegistered(uint256 indexed chainId, string indexed contractType, bytes addressBytes);

    /// @notice Emitted when a ZRC20 token is registered.
    /// @param originAddress The address of the asset on its native chain.
    /// @param address_ The address of the ZRC20 token on ZetaChain.
    /// @param decimals The number of decimals the token uses.
    /// @param originChainId The ID of the foreign chain where the original asset exists.
    /// @param symbol The symbol of the token.
    event ZRC20TokenRegistered(
        bytes indexed originAddress, address indexed address_, uint8 decimals, uint256 originChainId, string symbol
    );

    /// @notice Emitted when a ZRC20 token is updated.
    /// @param address_ The address of the ZRC20 token.
    /// @param active Whether the token should be active.
    event ZRC20TokenUpdated(address address_, bool active);

    /// @notice Emitted when a chain status has changed
    /// @param chainId The ID of the chain.
    /// @param oldStatus The new chain status (is active or not)
    /// @param newStatus The old chain status (is active or not)
    event ChainStatusChanged(uint256 indexed chainId, bool oldStatus, bool newStatus);

    /// @notice Emitted when a contract status has changed
    /// @param addressBytes The bytes representation of the registered address.
    event ContractStatusChanged(bytes addressBytes);

    /// @notice Emitted when a chain metadata is set
    /// @param chainId The ID of the chain.
    /// @param key The metadata key to update.
    /// @param value The new value for the metadata.
    event ChainMetadataUpdated(uint256 indexed chainId, string key, bytes value);

    /// @notice Emitted when a new contract configuration is updated
    /// @param chainId The ID of the chain where the contract is deployed.
    /// @param contractType The type of the contract.
    /// @param key The configuration key to update.
    /// @param value The new value for the configuration.
    event NewContractConfiguration(uint256 indexed chainId, string contractType, string key, bytes value);
}

/// @title ICoreRegistryErrors
/// @notice Interface for the errors used in the CoreRegistry contract.
interface ICoreRegistryErrors {
    /// @notice Error thrown when a contract type is invalid.
    /// @param message Describes why error happened
    error InvalidContractType(string message);

    /// @notice Error thrown when a zero address is provided where a non-zero address is required.
    error ZeroAddress();

    /// @notice Error thrown when a chain is already active.
    /// @param chainId The ID of the chain that is already active.
    error ChainActive(uint256 chainId);

    /// @notice Error thrown when a chain is not active.
    /// @param chainId The ID of the chain that is not active.
    error ChainNonActive(uint256 chainId);

    /// @notice Error thrown when a contract is already registered.
    /// @param chainId The ID of the chain.
    /// @param contractType The type of the contract.
    /// @param addressBytes The address of the contract.
    error ContractAlreadyRegistered(uint256 chainId, string contractType, bytes addressBytes);

    /// @notice Error thrown when a contract is not found in the registry.
    /// @param chainId The ID of the chain,
    /// @param contractType The type of the contract.
    error ContractNotFound(uint256 chainId, string contractType);

    /// @notice Error thrown when a ZRC20 token is already registered.
    /// @param address_ The address of the ZRC20 token.
    error ZRC20AlreadyRegistered(address address_);

    /// @notice Error thrown when a ZRC20 token symbol is already in use.
    /// @param symbol The symbol that is already in use.
    error ZRC20SymbolAlreadyInUse(string symbol);

    /// @notice Error thrown when a ZRC20 token transfer failed.
    error TransferFailed();
}

/// @title ChainInfo
/// @notice Structure that contains information about a chain.
struct ChainInfo {
    /// @notice Whether the chain is active in the ecosystem.
    bool active;
    /// @notice The unique identifier of the chain.
    uint256 chainId;
    /// @notice The address of the ZRC20 token that represents gas token for the chain.
    address gasZRC20;
    /// @notice The registry address on the target chain
    bytes registry;
    /// @notice Additional chain-specific metadata stored as key-value pairs.
    mapping(string => bytes) metadata;
}

/// @title ContractInfo
/// @notice Structure that contains information about a contract registered in the system.
struct ContractInfo {
    /// @notice Whether the contract is active.
    bool active;
    /// @notice The contract address in bytes representation.
    bytes addressBytes;
    /// @notice The type of the contract (e.g. "connector", "gateway", "tss").
    string contractType;
    /// @notice Additional contract-specific configuration and metadata.
    mapping(string => bytes) configuration;
}

/// @title ZRC20Info
/// @notice Structure that contains information about a ZRC20 token.
struct ZRC20Info {
    /// @notice Whether the ZRC20 token is active.
    bool active;
    /// @notice The address of the ZRC20 token on ZetaChain.
    address address_;
    /// @notice The address or identifier of the asset on its native chain.
    bytes originAddress;
    /// @notice The ID of the foreign chain where the original asset exists.
    uint256 originChainId;
    /// @notice The symbol of the token.
    string symbol;
    /// @notice The type of the asset gas/erc20.
    string coinType;
    /// @notice The number of decimals the token uses.
    uint8 decimals;
}

/// @title ICoreRegistry
/// @notice Primary interface for the CoreRegistry contract, the central registry for ZetaChain ecosystem.
interface ICoreRegistry is ICoreRegistryErrors, ICoreRegistryEvents {
    //--------------------------------------------------------------------------
    // Chain Management Functions
    //--------------------------------------------------------------------------

    /// @notice Changes status of the chain to activated/deactivated.
    /// @param chainId The ID of the chain to activate.
    /// @param gasZRC20 The address of the ZRC20 token that represents gas token for the chain.
    /// @param activation Whether activate or deactivate a chain
    function changeChainStatus(uint256 chainId, address gasZRC20, bytes calldata registry, bool activation) external;

    /// @notice Updates chain metadata.
    /// @param chainId The ID of the chain.
    /// @param key The metadata key to update.
    /// @param value The new value for the metadata.
    function updateChainMetadata(uint256 chainId, string calldata key, bytes calldata value) external;

    //--------------------------------------------------------------------------
    // Contract Registration and Management Functions
    //--------------------------------------------------------------------------

    /// @notice Registers a new contract address for a specific chain.
    /// @param chainId The ID of the chain where the contract is deployed.
    /// @param contractType The type of the contract (e.g., "connector", "gateway").
    /// @param addressBytes The bytes representation of the non-EVM address.
    function registerContract(uint256 chainId, string calldata contractType, bytes calldata addressBytes) external;

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
        external;

    function setContractActive(uint256 chainId, string calldata contractType, bool active) external;

    //--------------------------------------------------------------------------
    // ZRC20 Token Management Functions
    //--------------------------------------------------------------------------

    /// @notice Registers a new ZRC20 token in the registry.
    /// @param address_ The address of the ZRC20 token on ZetaChain.
    /// @param symbol The symbol of the token.
    /// @param originChainId The ID of the foreign chain where the original asset exists.
    /// @param originAddress The address or identifier of the asset on its native chain.
    function registerZRC20Token(
        address address_,
        string calldata symbol,
        uint256 originChainId,
        bytes calldata originAddress,
        string calldata coinType,
        uint8 decimals
    )
        external;

    /// @notice Updates ZRC20 token information.
    /// @param address_ The address of the ZRC20 token.
    /// @param active Whether the token should be active.
    function setZRC20TokenActive(address address_, bool active) external;

    //--------------------------------------------------------------------------
    // Registry Query Functions
    //--------------------------------------------------------------------------

    /// @notice Gets chain-specific metadata.
    /// @param chainId The ID of the chain.
    /// @param key The metadata key to retrieve.
    /// @return The value of the requested metadata.
    function getChainMetadata(uint256 chainId, string calldata key) external view returns (bytes memory);

    /// @notice Gets information about a specific contract.
    /// @param chainId The ID of the chain where the contract is deployed.
    /// @param contractType The type of the contract.
    /// @return active Whether the contract is active.
    /// @return addressBytes The address of the contract.
    function getContractInfo(
        uint256 chainId,
        string calldata contractType
    )
        external
        view
        returns (bool active, bytes memory addressBytes);

    /// @notice Gets contract-specific configuration.
    /// @param chainId The ID of the chain where the contract is deployed.
    /// @param contractType The type of the contract.
    /// @param key The configuration key to retrieve.
    /// @return The value of the requested configuration.
    function getContractConfiguration(
        uint256 chainId,
        string calldata contractType,
        string calldata key
    )
        external
        view
        returns (bytes memory);

    /// @notice Gets information about a specific ZRC20 token.
    /// @param address_ The address of the ZRC20 token.
    /// @return active Whether the token is active.
    /// @return symbol The symbol of the token
    /// @return originChainId The ID of the foreign chain where the original asset exists.
    /// @return originAddress The address or identifier of the asset on its native chain.
    /// @return coinType The type of the original coin.
    /// @return decimals The number of decimals the token uses.
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
        );

    /// @notice Gets the ZRC20 token address for a specific asset on a foreign chain.
    /// @param originChainId The ID of the foreign chain.
    /// @param originAddress The address or identifier of the asset on its native chain.
    /// @return The address of the corresponding ZRC20 token on ZetaChain.
    function getZRC20AddressByForeignAsset(
        uint256 originChainId,
        bytes calldata originAddress
    )
        external
        view
        returns (address);

    /// @notice Gets all active chains in the registry.
    /// @return Array of chain IDs for all active chains.
    function getActiveChains() external view returns (uint256[] memory);
}
