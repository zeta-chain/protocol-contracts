// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

/// @title IRegistryEvents
interface IRegistryEvents {
    /// @notice Emitted when a chain's status is changed
    event ChainStatusChanged(uint256 indexed chainId, bool previousState, bool newState);

    /// @notice Emitted when new chain metadata is set
    event NewChainMetadata(uint256 indexed chainId, string key, bytes value);

    /// @notice Emitted when a contract is registered
    event ContractRegistered(uint256 indexed chainId, string contractType, bytes addressBytes);

    /// @notice Emitted when contract configuration is updated
    event NewContractConfiguration(uint256 indexed chainId, string contractType, string key, bytes value);

    /// @notice Emitted when a contract's status is changed
    event ContractStatusChanged(bytes addressBytes);

    /// @notice Emitted when a ZRC20 token is registered
    event ZRC20TokenRegistered(
        bytes originAddress, address indexed zrc20, uint8 decimals, uint256 originChainId, string symbol
    );

    /// @notice Emitted when a ZRC20 token's status is updated
    event ZRC20TokenUpdated(address indexed zrc20, bool active);
}

/// @title IRegistryErrors
interface IRegistryErrors {
    /// @notice Error thrown when an operation requires the caller to be the gateway
    error NotGateway(address caller);

    /// @notice Error thrown when the chain ID is already active
    error ChainActive(uint256 chainId);

    /// @notice Error thrown when the chain ID is not active
    error ChainNonActive(uint256 chainId);

    /// @notice Error thrown when a zero address is provided
    error ZeroAddress();

    /// @notice Error thrown when an invalid contract type is provided
    error InvalidContractType(string contractType);

    /// @notice Error thrown when a contract is already registered
    error ContractAlreadyRegistered(uint256 chainId, string contractType, bytes addressBytes);

    /// @notice Error thrown when a contract is not found
    error ContractNotFound(uint256 chainId, string contractType);

    /// @notice Error thrown when a ZRC20 token is already registered
    error ZRC20AlreadyRegistered(address zrc20);

    /// @notice Error thrown when a ZRC20 symbol is already in use
    error ZRC20SymbolAlreadyInUse(string symbol);

    /// @notice Error thrown when the sender is invalid
    error InvalidSender();

    /// @notice Error thrown when the caller is not authorized
    error NotAuthorized(address caller);
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
    /// @notice The address of the contract (for EVM chains).
    address address_;
    /// @notice Bytes representation of the address (needed for non-EVM chains).
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

/// @title IRegistry
interface IRegistry is IRegistryErrors, IRegistryEvents { }
