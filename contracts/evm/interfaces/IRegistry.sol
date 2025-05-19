// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

/// @title IRegistry
interface IRegistry {
    /// @notice Structure for chain data used during bootstrapping
    struct ChainBootstrapData {
        /// @notice Whether the chain is active in the ecosystem.
        bool active;
        /// @notice The unique identifier of the chain.
        uint256 chainId;
        /// @notice The address of the ZRC20 token that represents gas token for the chain.
        address gasZRC20;
        /// @notice The registry address deployed on the chain.
        bytes registry;
    }

    /// @notice Structure for contract data used during bootstrapping
    struct ContractBootstrapData {
        /// @notice Whether the contract is active.
        bool active;
        /// @notice The contract address in bytes representation.
        bytes addressBytes;
        /// @notice The type of the contract (e.g. "connector", "gateway", "tss").
        string contractType;
        /// @notice Represents id of the chain where contract is deployed.
        uint256 chainId;
    }

    /// @notice Structure for ZRC20 token data used during bootstrapping
    struct ZRC20BootstrapData {
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

    /// @notice Structure for metadata entries used during bootstrapping
    struct ChainMetadataEntry {
        /// @notice The unique identifier of the chain.
        uint256 chainId;
        /// @param key The metadata key to update.
        string key;
        /// @param value The new value for the metadata.
        bytes value;
    }

    /// @notice Structure for contract configuration entries used during bootstrapping
    struct ContractConfigEntry {
        /// @notice Represents id of the chain where contract is deployed.
        uint256 chainId;
        /// @notice The type of the contract (e.g. "connector", "gateway", "tss").
        string contractType;
        /// @param key The configuration key to update.
        string key;
        /// @param value The new value for the configuration.
        bytes value;
    }
}
