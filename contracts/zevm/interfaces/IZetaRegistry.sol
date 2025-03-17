// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

/// @title IZetaRegistryEvents
/// @notice Interface for the events emitted by the ZetaRegistry contract.
interface IZetaRegistryEvents {
    /// @notice Emitted when a new blockchain is added to the registry.
    /// @param chainIdentifier Unique identifier for the blockchain.
    /// @param chainId Blockchain network ID.
    /// @param name Human-readable name of the blockchain.
    event ChainAdded(bytes32 indexed chainIdentifier, uint256 chainId, string name);

    /// @notice Emitted when a blockchain's active status is updated.
    /// @param chainIdentifier Unique identifier for the blockchain.
    /// @param active New active status.
    event ChainStatusUpdated(bytes32 indexed chainIdentifier, bool active);

    /// @notice Emitted when a contract address is updated.
    /// @param chainIdentifier Unique identifier for the blockchain.
    /// @param contractIdentifier Unique identifier for the contract.
    /// @param category Contract category.
    /// @param contractAddress New contract address.
    event ContractAddressUpdated(
        bytes32 indexed chainIdentifier,
        bytes32 indexed contractIdentifier,
        uint8 indexed category,
        address contractAddress
    );

    /// @notice Emitted when a contract's configuration is updated.
    /// @param chainIdentifier Unique identifier for the blockchain.
    /// @param contractIdentifier Unique identifier for the contract.
    /// @param configurationData New configuration data (encoded).
    event ContractConfigurationUpdated(
        bytes32 indexed chainIdentifier,
        bytes32 indexed contractIdentifier,
        bytes configurationData
    );

    /// @notice Emitted when a contract's active status is updated.
    /// @param chainIdentifier Unique identifier for the blockchain.
    /// @param contractIdentifier Unique identifier for the contract.
    /// @param active New active status.
    event ContractStatusUpdated(
        bytes32 indexed chainIdentifier,
        bytes32 indexed contractIdentifier,
        bool active
    );

    /// @notice Emitted when a ZRC20 token is added to the registry.
    /// @param chainIdentifier Unique identifier for the blockchain.
    /// @param tokenIdentifier Unique identifier for the token.
    /// @param tokenAddress Address of the ZRC20 token.
    event ZRC20TokenAdded(
        bytes32 indexed chainIdentifier,
        bytes32 indexed tokenIdentifier,
        address tokenAddress
    );
}

/// @title IZetaRegistryErrors
/// @notice Interface for the errors used in the ZetaRegistry contract.
interface IZetaRegistryErrors {
    /// @notice Error indicating the chain already exists.
    /// @param chainIdentifier Unique identifier for the blockchain.
    error ChainAlreadyExists(bytes32 chainIdentifier);

    /// @notice Error indicating the chain does not exist.
    /// @param chainIdentifier Unique identifier for the blockchain.
    error ChainDoesNotExist(bytes32 chainIdentifier);

    /// @notice Error indicating the contract already exists.
    /// @param chainIdentifier Unique identifier for the blockchain.
    /// @param contractIdentifier Unique identifier for the contract.
    error ContractAlreadyExists(bytes32 chainIdentifier, bytes32 contractIdentifier);

    /// @notice Error indicating the contract does not exist.
    /// @param chainIdentifier Unique identifier for the blockchain.
    /// @param contractIdentifier Unique identifier for the contract.
    error ContractDoesNotExist(bytes32 chainIdentifier, bytes32 contractIdentifier);

    /// @notice Error indicating an invalid contract address.
    /// @param contractAddress The invalid contract address.
    error InvalidContractAddress(address contractAddress);

    /// @notice Error indicating the ZRC20 token already exists.
    /// @param chainIdentifier Unique identifier for the blockchain.
    /// @param tokenIdentifier Unique identifier for the token.
    error ZRC20TokenAlreadyExists(bytes32 chainIdentifier, bytes32 tokenIdentifier);

    /// @notice Error indicating the ZRC20 token does not exist.
    /// @param chainIdentifier Unique identifier for the blockchain.
    /// @param tokenIdentifier Unique identifier for the token.
    error ZRC20TokenDoesNotExist(bytes32 chainIdentifier, bytes32 tokenIdentifier);

    /// @notice Error indicating an unauthorized caller.
    /// @param caller The address that attempted to call the function.
    error UnauthorizedCaller(address caller);

    /// @notice Error indicating array length mismatch.
    /// @param expected Expected length.
    /// @param actual Actual length.
    error ArrayLengthMismatch(uint256 expected, uint256 actual);
}

/// @title IZetaRegistry
/// @notice Interface for the ZetaChain Registry contract which serves as a central
/// repository for all contract addresses across connected blockchains.
interface IZetaRegistry is IZetaRegistryEvents, IZetaRegistryErrors {
    /// @notice Enum representing different system contract categories.
    enum ContractCategory {
        ZRC20,
        GATEWAY,
        GOVERNANCE,
        CONNECTOR,
        CUSTODY,
        OTHER
    }

    /// @notice Struct representing core blockchain information.
    /// @param chainId The unique identifier for this blockchain.
    /// @param name Human-readable name of the blockchain.
    /// @param active Whether this blockchain is currently active in the ZetaChain network.
    /// @param chainType Type of blockchain (EVM, Cosmos, etc.).
    /// @param blockConfirmations Number of block confirmations required for finality.
    struct ChainInfo {
        uint256 chainId;
        string name;
        bool active;
        uint8 chainType;
        uint64 blockConfirmations;
    }

    /// @notice Struct representing a system contract's configuration data.
    /// @param active Whether this contract is currently active.
    /// @param addr The deployed address of the contract.
    /// @param version Semantic version of the contract.
    /// @param implementation For proxy contracts, the address of the implementation.
    /// @param configurationData Additional contract-specific configuration data (encoded).
    struct ContractInfo {
        bool active;
        address addr;
        string version;
        address implementation;
        bytes configurationData;
    }

    /// @notice Adds a new blockchain to the registry.
    /// @param chainIdentifier Unique identifier for the blockchain (e.g., bytes32("eth_mainnet")).
    /// @param chainInfo Blockchain information (see ChainInfo struct).
    /// @return success True if the operation was successful.
    function addChain(bytes32 chainIdentifier, ChainInfo calldata chainInfo) external returns (bool success);

    /// @notice Updates an existing blockchain's information.
    /// @param chainIdentifier Unique identifier for the blockchain.
    /// @param chainInfo Updated blockchain information.
    /// @return success True if the operation was successful.
    function updateChain(bytes32 chainIdentifier, ChainInfo calldata chainInfo) external returns (bool success);

    /// @notice Sets a blockchain's active status.
    /// @param chainIdentifier Unique identifier for the blockchain.
    /// @param active New active status.
    /// @return success True if the operation was successful.
    function setChainActive(bytes32 chainIdentifier, bool active) external returns (bool success);

    /// @notice Removes a blockchain from the registry.
    /// @param chainIdentifier Unique identifier for the blockchain.
    /// @return success True if the operation was successful.
    function removeChain(bytes32 chainIdentifier) external returns (bool success);

    /// @notice Retrieves information about a blockchain.
    /// @param chainIdentifier Unique identifier for the blockchain.
    /// @return chainInfo The blockchain information.
    function getChainInfo(bytes32 chainIdentifier) external view returns (ChainInfo memory chainInfo);

    /// @notice Checks if a blockchain exists in the registry.
    /// @param chainIdentifier Unique identifier for the blockchain.
    /// @return exists True if the blockchain exists.
    function chainExists(bytes32 chainIdentifier) external view returns (bool exists);

    /// @notice Gets all blockchain identifiers in the registry.
    /// @return chainIdentifiers Array of all blockchain identifiers.
    function getAllChainIdentifiers() external view returns (bytes32[] memory chainIdentifiers);

    /// @notice Gets all active blockchain identifiers in the registry.
    /// @return activeChainIdentifiers Array of active blockchain identifiers.
    function getActiveChainIdentifiers() external view returns (bytes32[] memory activeChainIdentifiers);

    /// @notice Gets the total number of blockchains in the registry.
    /// @return count Total number of blockchains.
    function getChainCount() external view returns (uint256 count);

    /// @notice Adds a new contract to a blockchain.
    /// @param chainIdentifier Unique identifier for the blockchain.
    /// @param contractIdentifier Unique identifier for the contract (e.g., bytes32("ZetaConnector")).
    /// @param category Contract category.
    /// @param contractInfo Contract information (see ContractInfo struct).
    /// @return success True if the operation was successful.
    function addContract(
        bytes32 chainIdentifier,
        bytes32 contractIdentifier,
        ContractCategory category,
        ContractInfo calldata contractInfo
    ) external returns (bool success);

    /// @notice Updates an existing contract's information.
    /// @param chainIdentifier Unique identifier for the blockchain.
    /// @param contractIdentifier Unique identifier for the contract.
    /// @param contractInfo Updated contract information.
    /// @return success True if the operation was successful.
    function updateContract(
        bytes32 chainIdentifier,
        bytes32 contractIdentifier,
        ContractInfo calldata contractInfo
    ) external returns (bool success);

    /// @notice Updates only a contract's address.
    /// @param chainIdentifier Unique identifier for the blockchain.
    /// @param contractIdentifier Unique identifier for the contract.
    /// @param newAddress New contract address.
    /// @return success True if the operation was successful.
    function updateContractAddress(
        bytes32 chainIdentifier,
        bytes32 contractIdentifier,
        address newAddress
    ) external returns (bool success);

    /// @notice Updates a contract's configuration data.
    /// @param chainIdentifier Unique identifier for the blockchain.
    /// @param contractIdentifier Unique identifier for the contract.
    /// @param configurationData New configuration data (encoded).
    /// @return success True if the operation was successful.
    function updateContractConfiguration(
        bytes32 chainIdentifier,
        bytes32 contractIdentifier,
        bytes calldata configurationData
    ) external returns (bool success);

    /// @notice Sets a contract's active status.
    /// @param chainIdentifier Unique identifier for the blockchain.
    /// @param contractIdentifier Unique identifier for the contract.
    /// @param active New active status.
    /// @return success True if the operation was successful.
    function setContractActive(
        bytes32 chainIdentifier,
        bytes32 contractIdentifier,
        bool active
    ) external returns (bool success);

    /// @notice Removes a contract from a blockchain.
    /// @param chainIdentifier Unique identifier for the blockchain.
    /// @param contractIdentifier Unique identifier for the contract.
    /// @return success True if the operation was successful.
    function removeContract(
        bytes32 chainIdentifier,
        bytes32 contractIdentifier
    ) external returns (bool success);

    /// @notice Gets information about a contract.
    /// @param chainIdentifier Unique identifier for the blockchain.
    /// @param contractIdentifier Unique identifier for the contract.
    /// @return contractInfo The contract information.
    function getContractInfo(
        bytes32 chainIdentifier,
        bytes32 contractIdentifier
    ) external view returns (ContractInfo memory contractInfo);

    /// @notice Gets a contract's address.
    /// @param chainIdentifier Unique identifier for the blockchain.
    /// @param contractIdentifier Unique identifier for the contract.
    /// @return addr The contract address.
    function getContractAddress(
        bytes32 chainIdentifier,
        bytes32 contractIdentifier
    ) external view returns (address addr);

    /// @notice Gets all contract identifiers for a blockchain.
    /// @param chainIdentifier Unique identifier for the blockchain.
    /// @return contractIdentifiers Array of all contract identifiers for the blockchain.
    function getContractIdentifiers(
        bytes32 chainIdentifier
    ) external view returns (bytes32[] memory contractIdentifiers);

    /// @notice Gets all contract identifiers for a blockchain by category.
    /// @param chainIdentifier Unique identifier for the blockchain.
    /// @param category Contract category.
    /// @return contractIdentifiers Array of contract identifiers in the specified category.
    function getContractIdentifiersByCategory(
        bytes32 chainIdentifier,
        ContractCategory category
    ) external view returns (bytes32[] memory contractIdentifiers);

    /// @notice Checks if a contract exists in the registry.
    /// @param chainIdentifier Unique identifier for the blockchain.
    /// @param contractIdentifier Unique identifier for the contract.
    /// @return exists True if the contract exists.
    function contractExists(
        bytes32 chainIdentifier,
        bytes32 contractIdentifier
    ) external view returns (bool exists);

    /// @notice Creates a contract identifier from a string.
    /// @param name String name for the contract.
    /// @return identifier bytes32 identifier.
    function createContractIdentifier(string calldata name) external pure returns (bytes32 identifier);

    /// @notice Creates a chain identifier from a string.
    /// @param name String name for the chain.
    /// @return identifier bytes32 identifier.
    function createChainIdentifier(string calldata name) external pure returns (bytes32 identifier);

    /// @notice Helper to get common system contract addresses for a blockchain.
    /// @param chainIdentifier Unique identifier for the blockchain.
    /// @return connector The connector contract address.
    /// @return erc20Custody The ERC20 custody contract address.
    /// @return tss The TSS contract address.
    /// @return tssUpdater The TSS updater contract address.
    /// @return zetaToken The ZETA token contract address.
    function getSystemAddresses(bytes32 chainIdentifier)
    external
    view
    returns (
        address connector,
        address erc20Custody,
        address tss,
        address tssUpdater,
        address zetaToken
    );

    /// @notice Helper to get ZRC-20 token address for a given chain and token.
    /// @param chainIdentifier Unique identifier for the blockchain.
    /// @param tokenIdentifier Unique identifier for the token (e.g., bytes32("USDC")).
    /// @return tokenAddress The ZRC-20 token address.
    function getZRC20Address(
        bytes32 chainIdentifier,
        bytes32 tokenIdentifier
    ) external view returns (address tokenAddress);

    /// @notice Adds a new ZRC-20 token to the registry.
    /// @param chainIdentifier Unique identifier for the blockchain.
    /// @param tokenIdentifier Unique identifier for the token.
    /// @param tokenAddress Address of the ZRC-20 token.
    /// @return success True if the operation was successful.
    function addZRC20Token(
        bytes32 chainIdentifier,
        bytes32 tokenIdentifier,
        address tokenAddress
    ) external returns (bool success);

    /// @notice Batch add contracts to a blockchain.
    /// @param chainIdentifier Unique identifier for the blockchain.
    /// @param contractIdentifiers Array of contract identifiers.
    /// @param categories Array of contract categories.
    /// @param contractInfos Array of contract information.
    /// @return success True if the operation was successful.
    function batchAddContracts(
        bytes32 chainIdentifier,
        bytes32[] calldata contractIdentifiers,
        ContractCategory[] calldata categories,
        ContractInfo[] calldata contractInfos
    ) external returns (bool success);
}