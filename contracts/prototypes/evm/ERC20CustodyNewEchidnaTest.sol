import "./GatewayEVM.sol";
import "./TestERC20.sol";
import "./ERC20CustodyNew.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

contract ERC20CustodyNewEchidnaTest is ERC20CustodyNew {
    using SafeERC20 for IERC20;

    TestERC20 public testERC20;
    address public echidnaCaller = msg.sender;

    GatewayEVM testGateway = new GatewayEVM();

    constructor() ERC20CustodyNew(address(testGateway)) {
        testGateway.initialize(echidnaCaller);
        testERC20 = new TestERC20("test", "TEST");
        testGateway.setCustody(address(this));
    }

    // Test withdrawAndCall with assertions
    function testWithdrawAndCall(address to, uint256 amount, bytes calldata data) public {
        // mint more than amount
        testERC20.mint(address(this), amount + 5);
        // transfer overhead to gateway
        testERC20.transfer(address(testGateway), 5);

        withdrawAndCall(address(testERC20), to, amount, data);

        // Assertion to ensure no ERC20 tokens are left in the gateway contract after execution
        assert(testERC20.balanceOf(address(gateway)) == 0);
    }
}