// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

interface IZRC20 {
    function totalSupply() external view returns (uint256);

    function balanceOf(address account) external view returns (uint256);

    function transfer(address recipient, uint256 amount) external returns (bool);

    function allowance(address owner, address spender) external view returns (uint256);

    function approve(address spender, uint256 amount) external returns (bool);

    function decreaseAllowance(address spender, uint256 amount) external returns (bool);

    function increaseAllowance(address spender, uint256 amount) external returns (bool);

    function transferFrom(address sender, address recipient, uint256 amount) external returns (bool);

    function deposit(address to, uint256 amount) external returns (bool);

    function burn(address account, uint256 amount) external returns (bool);

    function withdraw(bytes memory to, uint256 amount) external returns (bool);

    function withdrawGasFee() external view returns (address, uint256);

    function PROTOCOL_FEE() external view returns (uint256);

    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(address indexed owner, address indexed spender, uint256 value);
    event Deposit(bytes from, address indexed to, uint256 value);
    event Withdrawal(address indexed from, bytes to, uint256 value, uint256 gasFee, uint256 protocolFlatFee);
    event UpdatedSystemContract(address systemContract);
    event UpdatedGasLimit(uint256 gasLimit);
    event UpdatedProtocolFlatFee(uint256 protocolFlatFee);
}
