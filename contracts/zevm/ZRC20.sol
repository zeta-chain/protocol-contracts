// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;
import "./Interfaces.sol";

/**
 * @dev Custom errors for ZRC20
 */
interface ZRC20Errors {
    // @dev: Error thrown when caller is not the fungible module
    error CallerIsNotFungibleModule();
    error InvalidSender();
    error GasFeeTransferFailed();
    error ZeroGasCoin();
    error ZeroGasPrice();
    error LowAllowance();
    error LowBalance();
    error ZeroAddress();
}

contract ZRC20 is Context, IZRC20, IZRC20Metadata, ZRC20Errors {
    /// @notice The fungible module address, this is maintained at the protocol level and is always constant
    address public constant FUNGIBLE_MODULE_ADDRESS = 0x735b14BB79463307AAcBED86DAf3322B1e6226aB;
    /// @notice Chain id.abi
    uint256 public immutable CHAIN_ID;
    /// @notice Coin type, checkout Interfaces.sol.
    CoinType public immutable COIN_TYPE;
    /// @notice System contract address.
    address public SYSTEM_CONTRACT_ADDRESS;
    /// @notice Gas limit.
    uint256 public GAS_LIMIT;
    /// @notice Protocol flat fee.
    uint256 public PROTOCOL_FLAT_FEE;

    /// @dev Mapping that keeps track of each account's token balance.
    mapping(address => uint256) private _balances;

    /// @dev Nested mapping that tracks how many tokens each account is allowed to spend from each other account.
    mapping(address => mapping(address => uint256)) private _allowances;

    /// @dev Total supply of tokens
    uint256 private _totalSupply;

    /// @notice Name of the token
    string private _name;

    /// @notice Symbol of the token
    string private _symbol;

    /// @notice Number of decimal places the token can be divided into
    uint8 private _decimals;

    /**
     * @dev Only fungible module modifier.
     */
    modifier onlyFungible() {
        if (msg.sender != FUNGIBLE_MODULE_ADDRESS) revert CallerIsNotFungibleModule();
        _;
    }

    /**
     * @dev Constructor that gives msg.sender all of existing tokens.
     * @notice Only the fungible module is allowed to deploy a new ZRC20 contract.
     * @param name_ Name of the token
     * @param symbol_ Symbol of the token
     * @param decimals_ Number of decimal places the token can be divided into
     * @param chainid_ Chain ID
     * @param coinType_ Coin Type
     * @param gasLimit_ Gas limit for transactions
     * @param systemContractAddress_ Address of the system contract
     */
    constructor(
        string memory name_,
        string memory symbol_,
        uint8 decimals_,
        uint256 chainid_,
        CoinType coinType_,
        uint256 gasLimit_,
        address systemContractAddress_
    ) {
        if (msg.sender != FUNGIBLE_MODULE_ADDRESS) revert CallerIsNotFungibleModule();
        _name = name_;
        _symbol = symbol_;
        _decimals = decimals_;
        CHAIN_ID = chainid_;
        COIN_TYPE = coinType_;
        GAS_LIMIT = gasLimit_;
        SYSTEM_CONTRACT_ADDRESS = systemContractAddress_;
    }

    /**
     * @dev ZRC20 name
     * @return name as string
     */
    function name() public view virtual override returns (string memory) {
        return _name;
    }

    /**
     * @dev ZRC20 symbol.
     * @return symbol as string.
     */
    function symbol() public view virtual override returns (string memory) {
        return _symbol;
    }

    /**
     * @dev ZRC20 decimals.
     * @return returns uint8 decimals.
     */
    function decimals() public view virtual override returns (uint8) {
        return _decimals;
    }

    /**
     * @dev ZRC20 total supply.
     * @return returns uint256 total supply.
     */
    function totalSupply() public view virtual override returns (uint256) {
        return _totalSupply;
    }

    /**
     * @dev Returns ZRC20 balance of an account.
     * @param account, account address for which balance is requested.
     * @return uint256 account balance.
     */
    function balanceOf(address account) public view virtual override returns (uint256) {
        return _balances[account];
    }

    /**
     * @notice Transfers a specified amount of tokens to the given recipient.
     * @dev This function can be called by the contract owner or any other external address.
     * @param recipient The address of the recipient to whom the tokens will be transferred.
     * @param amount The amount of tokens to transfer.
     * @return Returns a boolean value indicating whether the transfer was successful or not.
     */
    function transfer(address recipient, uint256 amount) public virtual override returns (bool) {
        _transfer(_msgSender(), recipient, amount);
        return true;
    }

    /**
     * @dev Returns token allowance from owner to spender.
     * @param owner, owner address.
     * @return uint256 allowance.
     */
    function allowance(address owner, address spender) public view virtual override returns (uint256) {
        return _allowances[owner][spender];
    }

    /**
     * @dev Approves amount transferFrom for spender.
     * @param spender, spender address.
     * @param amount, amount to approve.
     * @return true/false if succeeded/failed.
     */
    function approve(address spender, uint256 amount) public virtual override returns (bool) {
        _approve(_msgSender(), spender, amount);
        return true;
    }

    /**
     * @dev Increases allowance by amount for spender.
     * @param spender, spender address.
     * @param amount, amount by which to increase allownace.
     * @return true/false if succeeded/failed.
     */
    function increaseAllowance(address spender, uint256 amount) external virtual returns (bool) {
        _allowances[spender][_msgSender()] += amount;
        return true;
    }

    /**
     * @dev Decreases allowance by amount for spender.
     * @param spender, spender address.
     * @param amount, amount by which to decrease allownace.
     * @return true/false if succeeded/failed.
     */
    function decreaseAllowance(address spender, uint256 amount) external virtual returns (bool) {
        if (_allowances[spender][_msgSender()] < amount) revert LowAllowance();
        _allowances[spender][_msgSender()] -= amount;
        return true;
    }

    /**
     * @dev Transfers tokens from sender to recipient.
     * @param sender, sender address.
     * @param recipient, recipient address.
     * @param amount, amount to transfer.
     * @return true/false if succeeded/failed.
     */
    function transferFrom(address sender, address recipient, uint256 amount) public virtual override returns (bool) {
        _transfer(sender, recipient, amount);

        uint256 currentAllowance = _allowances[sender][_msgSender()];
        if (currentAllowance < amount) revert LowAllowance();

        _approve(sender, _msgSender(), currentAllowance - amount);

        return true;
    }

    /**
     * @dev Burns an amount of tokens.
     * @param amount, amount to burn.
     * @return true/false if succeeded/failed.
     */
    function burn(uint256 amount) external returns (bool) {
        _burn(msg.sender, amount);
        return true;
    }

    /**
     * @dev Internal function to transfer tokens from one address to another.
     * @param sender The address sending the tokens.
     * @param recipient The address receiving the tokens.
     * @param amount The amount of tokens to transfer.
     * @dev Throws if either the sender or recipient address is zero.
     * @dev Throws if the sender's balance is lower than the transfer amount.
     */
    function _transfer(address sender, address recipient, uint256 amount) internal virtual {
        if (sender == address(0)) revert ZeroAddress();
        if (recipient == address(0)) revert ZeroAddress();

        uint256 senderBalance = _balances[sender];
        if (senderBalance < amount) revert LowBalance();

        _balances[sender] = senderBalance - amount;
        _balances[recipient] += amount;

        emit Transfer(sender, recipient, amount);
    }

    /**
     * @dev Internal function to mint new tokens and assign them to an account.
     * @param account The address to which the minted tokens will be assigned.
     * @param amount The amount of tokens to be minted.
     * @dev Throws if the account address is zero.
     */
    function _mint(address account, uint256 amount) internal virtual {
        if (account == address(0)) revert ZeroAddress();

        _totalSupply += amount;
        _balances[account] += amount;
        emit Transfer(address(0), account, amount);
    }

    /**
     * @dev Internal function to burn tokens from an account.
     * @param account The address from which tokens will be burned.
     * @param amount The amount of tokens to be burned.
     * @dev Throws if the account address is zero.
     * @dev Throws if the account's balance is lower than the burn amount.
     */
    function _burn(address account, uint256 amount) internal virtual {
        if (account == address(0)) revert ZeroAddress();

        uint256 accountBalance = _balances[account];
        if (accountBalance < amount) revert LowBalance();

        _balances[account] = accountBalance - amount;
        _totalSupply -= amount;

        emit Transfer(account, address(0), amount);
    }

    /**
     * @dev Internal function to approve a spender to spend tokens on behalf of the owner.
     * @param owner The address that owns the tokens.
     * @param spender The address that is approved to spend the tokens.
     * @param amount The maximum amount of tokens that can be spent.
     * @dev Throws if the owner address is zero.
     * @dev Throws if the spender address is zero.
     */
    function _approve(address owner, address spender, uint256 amount) internal virtual {
        if (owner == address(0)) revert ZeroAddress();
        if (spender == address(0)) revert ZeroAddress();

        _allowances[owner][spender] = amount;
        emit Approval(owner, spender, amount);
    }

    /**
     * @dev Deposits corresponding tokens from an external chain.
     * @param to The recipient address.
     * @param amount The amount to deposit.
     * @return A boolean indicating whether the deposit succeeded or failed.
     * @dev Only callable by the Fungible module or the System contract.
     * @dev Throws if called by an invalid sender.
     */
    function deposit(address to, uint256 amount) external override returns (bool) {
        if (msg.sender != FUNGIBLE_MODULE_ADDRESS && msg.sender != SYSTEM_CONTRACT_ADDRESS) revert InvalidSender();
        _mint(to, amount);
        emit Deposit(abi.encodePacked(FUNGIBLE_MODULE_ADDRESS), to, amount);
        return true;
    }

    /**
     * @dev Returns the ZRC20 address for gas on the same chain of this ZRC20, and calculates the gas fee for `withdraw()`.
     * @return gasZRC20 The address of the gas ZRC20 token on the same chain.
     * @return gasFee The calculated gas fee for the `withdraw()` function.
     * @dev Throws if the gas ZRC20 address is zero.
     * @dev Throws if the gas price is zero.
     */
    function withdrawGasFee() public view override returns (address, uint256) {
        address gasZRC20 = ISystem(SYSTEM_CONTRACT_ADDRESS).gasCoinZRC20ByChainId(CHAIN_ID);
        if (gasZRC20 == address(0)) {
            revert ZeroGasCoin();
        }
        uint256 gasPrice = ISystem(SYSTEM_CONTRACT_ADDRESS).gasPriceByChainId(CHAIN_ID);
        if (gasPrice == 0) {
            revert ZeroGasPrice();
        }
        uint256 gasFee = gasPrice * GAS_LIMIT + PROTOCOL_FLAT_FEE;
        return (gasZRC20, gasFee);
    }

    /**
     * @dev Withdraws ZRC20 tokens to external chains by triggering the crosschain module to create an outbound transaction.
     * @param to The recipient address on the external chain.
     * @param amount The amount of tokens to withdraw.
     * @return A boolean indicating whether the withdrawal succeeded or failed.
     * @dev Requires this contract to have sufficient allowance of the gas ZRC20 token to pay for the outbound transaction gas fee.
     * @dev Throws if the gas fee transfer fails.
     */
    function withdraw(bytes memory to, uint256 amount) external override returns (bool) {
        (address gasZRC20, uint256 gasFee) = withdrawGasFee();
        if (!IZRC20(gasZRC20).transferFrom(msg.sender, FUNGIBLE_MODULE_ADDRESS, gasFee)) {
            revert GasFeeTransferFailed();
        }
        _burn(msg.sender, amount);
        emit Withdrawal(msg.sender, to, amount, gasFee, PROTOCOL_FLAT_FEE);
        return true;
    }

    /**
     * @dev Updates the system contract address.
     * @param addr The new system contract address to be set.
     * @dev Requires the caller to be the fungible module.
     */
    function updateSystemContractAddress(address addr) external onlyFungible {
        SYSTEM_CONTRACT_ADDRESS = addr;
        emit UpdatedSystemContract(addr);
    }

    /**
     * @dev Updates the gas limit.
     * @param gasLimit The new gas limit to be set.
     * @dev Requires the caller to be the fungible module.
     */
    function updateGasLimit(uint256 gasLimit) external onlyFungible {
        GAS_LIMIT = gasLimit;
        emit UpdatedGasLimit(gasLimit);
    }

    /**
     * @dev Updates the protocol flat fee.
     * @param protocolFlatFee The new protocol flat fee to be set.
     * @dev Requires the caller to be the fungible module.
     */
    function updateProtocolFlatFee(uint256 protocolFlatFee) external onlyFungible {
        PROTOCOL_FLAT_FEE = protocolFlatFee;
        emit UpdatedProtocolFlatFee(protocolFlatFee);
    }
}
