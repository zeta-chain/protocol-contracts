// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

import "forge-std/Test.sol";
import "forge-std/Vm.sol";

import "contracts/prototypes/evm/GatewayEVM.sol";
import "contracts/prototypes/evm/ReceiverEVM.sol";
import "contracts/prototypes/evm/ERC20CustodyNew.sol";
import "contracts/prototypes/evm/TestERC20.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "../evm/interfaces.sol";
import "forge-std/console.sol";

contract GatewayEVMTest is Test, IGatewayEVMErrors, IGatewayEVMEvents, IReceiverEVMEvents {
    using SafeERC20 for IERC20;

    GatewayEVM gateway;
    ReceiverEVM receiver;
    ERC20CustodyNew custody;
    TestERC20 token;
    address owner;
    address destination;
    address tssAddress;

    function setUp() public {
        owner = address(this);
        destination = address(0x1234);
        tssAddress = address(0x5678);

        token = new TestERC20("test", "TTK");
        gateway = new GatewayEVM();
        custody = new ERC20CustodyNew(address(gateway));

        gateway.initialize(tssAddress);
        gateway.setCustody(address(custody));

        // Mint initial supply to the owner
        token.mint(owner, 1000000);

        // Transfer some tokens to the custody contract
        token.transfer(address(custody), 500000);
    }

    function testForwardCallToReceivePayable() public {
        string memory str = "Hello, Foundry!";
        uint256 num = 42;
        bool flag = true;
        uint256 value = 1 ether;

        bytes memory data = abi.encodeWithSignature("receivePayable(string,uint256,bool)", str, num, flag);
        
        vm.expectCall(address(receiver), value, data);
        vm.expectEmit(true, true, true, true, address(gateway));
        emit Executed(address(receiver), value, data);
        // TODO: can not check event emitted in Receiver?
        
        gateway.execute{value: value}(address(receiver), data);
    }
}
