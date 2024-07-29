// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "forge-std/Vm.sol";

import "./utils/GatewayEVMUpgradeTest.sol";

import "./utils/IReceiverEVM.sol";
import "./utils/ReceiverEVM.sol";
import "./utils/TestERC20.sol";
import { ERC1967Proxy } from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {Upgrades} from "openzeppelin-foundry-upgrades/Upgrades.sol";

import "src/evm/interfaces/IGatewayEVM.sol";
import "./utils/IReceiverEVM.sol";
import "src/evm/GatewayEVM.sol";
import "src/evm/ERC20CustodyNew.sol";
import "src/evm/ZetaConnectorNonNative.sol";

contract GatewayEVMUUPSUpgradeTest is Test, IGatewayEVMErrors, IGatewayEVMEvents, IReceiverEVMEvents {
    using SafeERC20 for IERC20;

    event ExecutedV2(address indexed destination, uint256 value, bytes data);

    address proxy;
    GatewayEVM gateway;
    ReceiverEVM receiver;
    ERC20CustodyNew custody;
    ZetaConnectorNonNative zetaConnector;
    TestERC20 token;
    TestERC20 zeta;
    address owner;
    address destination;
    address tssAddress;

    function setUp() public {
        owner = address(this);
        destination = address(0x1234);
        tssAddress = address(0x5678);

        token = new TestERC20("test", "TTK");
        zeta = new TestERC20("zeta", "ZETA");

        proxy = Upgrades.deployUUPSProxy(
            "GatewayEVM.sol", abi.encodeCall(GatewayEVM.initialize, (tssAddress, address(zeta)))
        );
        gateway = GatewayEVM(proxy);

        custody = new ERC20CustodyNew(address(gateway), tssAddress);
        zetaConnector = new ZetaConnectorNonNative(address(gateway), address(zeta), tssAddress);
        receiver = new ReceiverEVM();

        vm.deal(tssAddress, 1 ether);

        vm.startPrank(tssAddress);
        gateway.setCustody(address(custody));
        gateway.setConnector(address(zetaConnector));
        vm.stopPrank();

        token.mint(owner, 1000000);
        token.transfer(address(custody), 500000);

        vm.deal(tssAddress, 1 ether);
    }

    function testUpgradeAndForwardCallToReceivePayable() public {
        address custodyBeforeUpgrade = gateway.custody();
        address tssBeforeUpgrade = gateway.tssAddress();

        string memory str = "Hello, Foundry!";
        uint256 num = 42;
        bool flag = true;
        uint256 value = 1 ether;

        Upgrades.upgradeProxy(proxy, "GatewayEVMUpgradeTest.sol", "", owner);

        bytes memory data = abi.encodeWithSignature("receivePayable(string,uint256,bool)", str, num, flag);
        GatewayEVMUpgradeTest gatewayUpgradeTest = GatewayEVMUpgradeTest(proxy);

        vm.expectCall(address(receiver), value, data);
        vm.expectEmit(true, true, true, true, address(receiver));
        emit ReceivedPayable(address(gateway), value, str, num, flag);
        vm.expectEmit(true, true, true, true, address(gateway));
        emit ExecutedV2(address(receiver), value, data);
        vm.prank(tssAddress);
        gatewayUpgradeTest.execute{value: value}(address(receiver), data);

        assertEq(custodyBeforeUpgrade, gateway.custody());
        assertEq(tssBeforeUpgrade, gateway.tssAddress());
    }
}
