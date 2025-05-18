// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

/// @title IRegistry
interface IRegistry {
    /// @notice Structure for chain data used during bootstrapping
    struct ChainBootstrapData {
        uint256 chainId;
        address gasZRC20;
        bytes registry;
        bool active;
    }

    /// @notice Structure for metadata entries used during bootstrapping
    struct ChainMetadataEntry {
        uint256 chainId;
        string key;
        bytes value;
    }

    /// @notice Structure for contract data used during bootstrapping
    struct ContractBootstrapData {
        uint256 chainId;
        string contractType;
        bytes addressBytes;
    }

    /// @notice Structure for contract configuration entries used during bootstrapping
    struct ContractConfigEntry {
        uint256 chainId;
        string contractType;
        string key;
        bytes value;
    }

    /// @notice Structure for ZRC20 token data used during bootstrapping
    struct ZRC20BootstrapData {
        address address_;
        string symbol;
        uint256 originChainId;
        bytes originAddress;
        string coinType;
        uint8 decimals;
        bool active;
    }
}
