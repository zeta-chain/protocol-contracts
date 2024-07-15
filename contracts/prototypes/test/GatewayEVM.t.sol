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
        receiver = new ReceiverEVM();

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
        vm.expectEmit(true, true, true, true, address(receiver));
        emit ReceivedPayable(address(gateway), value, str, num, flag);
        vm.expectEmit(true, true, true, true, address(gateway));
        emit Executed(address(receiver), value, data);
      
        gateway.execute{value: value}(address(receiver), data);
    }

    function testForwardCallToReceiveNonPayable() public {
        string[] memory str = new string[](1);
        str[0] = "Hello, Foundry!";
        uint256[] memory num = new uint256[](1);
        num[0] = 42;
        bool flag = true;

        bytes memory data = abi.encodeWithSignature("receiveNonPayable(string[],uint256[],bool)", str, num, flag);
        
        vm.expectCall(address(receiver), 0, data);
        vm.expectEmit(true, true, true, true, address(receiver));
        emit ReceivedNonPayable(address(gateway), str, num, flag);
        vm.expectEmit(true, true, true, true, address(gateway));
        emit Executed(address(receiver), 0, data);
        
        gateway.execute(address(receiver), data);
    }

    function testForwardCallToReceiveNoParams() public {
        bytes memory data = abi.encodeWithSignature("receiveNoParams()");
        
        vm.expectCall(address(receiver), 0, data);
        vm.expectEmit(true, true, true, true, address(receiver));
        emit ReceivedNoParams(address(gateway));
        vm.expectEmit(true, true, true, true, address(gateway));
        emit Executed(address(receiver), 0, data);
        
        gateway.execute(address(receiver), data);
    }
}
