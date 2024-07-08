// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

import "forge-std/Test.sol";
import "forge-std/Vm.sol";

import "contracts/prototypes/evm/GatewayEVM.sol";
import "contracts/prototypes/evm/ReceiverEVM.sol";
import "contracts/prototypes/evm/ERC20CustodyNew.sol";
import "contracts/prototypes/evm/TestERC20.sol";
import "contracts/prototypes/evm/ReceiverEVM.sol";

import "contracts/prototypes/zevm/GatewayZEVM.sol";
import "contracts/prototypes/zevm/SenderZEVM.sol";
import "contracts/zevm/ZRC20New.sol";
import "contracts/zevm/testing/SystemContractMock.sol";

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "../evm/interfaces.sol";
import "../zevm/interfaces.sol";
import "forge-std/console.sol";

contract GatewayEVMZEVMTest is Test, IGatewayEVMErrors, IGatewayEVMEvents, IGatewayZEVMEvents, IGatewayZEVMErrors, IReceiverEVMEvents {
    // evm
    using SafeERC20 for IERC20;

    GatewayEVM gatewayEVM;
    ERC20CustodyNew custody;
    TestERC20 token;
    ReceiverEVM receiverEVM;
    address ownerEVM;
    address destination;
    address tssAddress;

    // zevm
    GatewayZEVM gatewayZEVM;
    SenderZEVM senderZEVM;
    SystemContractMock systemContract;
    ZRC20New zrc20;
    address ownerZEVM;
    
    function setUp() public {
        // evm
        ownerEVM = address(this);
        destination = address(0x1234);
        tssAddress = address(0x5678);
        ownerZEVM = address(0x4321);

        token = new TestERC20("test", "TTK");
        gatewayEVM = new GatewayEVM();
        custody = new ERC20CustodyNew(address(gatewayEVM));

        gatewayEVM.initialize(tssAddress);
        gatewayEVM.setCustody(address(custody));

        // Mint initial supply to the ownerEVM
        token.mint(ownerEVM, 1000000);

        // Transfer some tokens to the custody contract
        token.transfer(address(custody), 500000);

        receiverEVM = new ReceiverEVM();

        // zevm
        gatewayZEVM = new GatewayZEVM();
        senderZEVM = new SenderZEVM(address(gatewayZEVM));
        // Impersonate the fungible module account
        address fungibleModuleAddress = address(0x735b14BB79463307AAcBED86DAf3322B1e6226aB);
        vm.startPrank(fungibleModuleAddress);
        systemContract = new SystemContractMock(address(0), address(0), address(0));
        zrc20 = new ZRC20New("TOKEN", "TKN", 18, 1, CoinType.Zeta, 0, address(systemContract), address(gatewayZEVM));
        systemContract.setGasCoinZRC20(1, address(zrc20));
        systemContract.setGasPrice(1, 1);
        zrc20.deposit(ownerZEVM, 1000000);
        zrc20.deposit(address(senderZEVM), 1000000);
        vm.stopPrank();

        vm.prank(ownerZEVM);
        zrc20.approve(address(gatewayZEVM), 1000000);
    }

      function testCallReceiverEVMFromZEVM() public {
        string memory str = "Hello, Hardhat!";
        uint256 num = 42;
        bool flag = true;
        uint256 value = 1 ether;

        // Encode the function call data and call on zevm
        bytes memory message = abi.encodeWithSelector(receiverEVM.receivePayable.selector, str, num, flag);
        vm.prank(ownerZEVM);
        vm.expectEmit(true, true, true, true, address(gatewayZEVM));
        emit Call(address(ownerZEVM), abi.encodePacked(receiverEVM), message);
        gatewayZEVM.call(abi.encodePacked(receiverEVM), message);

        // Call execute on evm
        vm.deal(address(gatewayEVM), value);
        vm.expectEmit(true, true, true, true, address(gatewayEVM));
        emit Executed(address(receiverEVM), value, message);
        gatewayEVM.execute{value: value}(address(receiverEVM), message);
    }
}