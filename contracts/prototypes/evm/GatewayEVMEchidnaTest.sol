// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

import "./GatewayEVM.sol";
import "./TestERC20.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

contract GatewayEVMEchidnaTest is GatewayEVM {
    using SafeERC20 for IERC20;
    address e = msg.sender;
    bool initialized = false;

    TestERC20 testERC20;

    // Add a constructor to bypass the initializer
    constructor() {
    // Do nothing here to avoid calling initialize during deployment
    }

    // Function to initialize the contract for testing
    function initializeTest() public {
        __Ownable_init();
        __UUPSUpgradeable_init();

        if (e == address(0)) {
            revert ZeroAddress();
        }
        tssAddress = e;
        initialized = true;

        testERC20 = new TestERC20("test", "TEST");
        testERC20.mint(e, 100000000000);
    }

    // Invariant to check that the contract balance is always zero
    function echidna_balance_is_zero() public view returns (bool) {
        return address(this).balance == 0;
    }

     // Invariant to check that the contract doesn't hold any ERC20 tokens
    function echidna_no_erc20_balance() public view returns (bool) {
         if (initialized == false) {
            return true;
        }
        return testERC20.balanceOf(address(this)) == 0;
    }

    // Invariant to check that the TSS address is never zero after initialization
    function echidna_tss_address_not_zero() public view returns (bool) {
        if (initialized == false) {
            return true;
        }
        return tssAddress != address(0);
    }
}