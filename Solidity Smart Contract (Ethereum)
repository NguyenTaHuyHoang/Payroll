pragma solidity ^0.4.0;

contract Payroll {
    address public owner;
    mapping(address => uint) public employeeBalances;

    event SalaryPaid(address indexed employee, uint amount);
    event RequestCashConversion(address indexed employee, uint amount);
    event CashTransferred(address indexed employee, uint amount);

    modifier onlyOwner {
        require(msg.sender == owner);
        _;
    }

    function Payroll() public {
        owner = msg.sender;
    }

    function paySalary(address employee, uint amount) public onlyOwner {
        employeeBalances[employee] += amount;
        emit SalaryPaid(employee, amount);
    }

    function requestCashConversion(uint amount) public {
        require(employeeBalances[msg.sender] >= amount);
        employeeBalances[msg.sender] -= amount;
        emit RequestCashConversion(msg.sender, amount);
        // Here we will call the transferCash function
        transferCash(msg.sender, amount);
    }

    function transferCash(address employee, uint amount) private {
        // This is a stub for transferring cash to the employee.
        // In a real scenario, this function would interact with another contract or external system.
        emit CashTransferred(employee, amount);
    }

    function viewSalary() public view returns (uint) {
        return employeeBalances[msg.sender];
    }
}
