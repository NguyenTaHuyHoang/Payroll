pragma solidity ^0.4.26;

contract Payroll {
    address public owner;
    uint public companyFund;
    mapping(address => uint) public employeeBalances;

    event SalaryPaid(address indexed employee, uint amount);
    event RequestCashConversion(address indexed employee, uint amount);
    event CashTransferred(address indexed employee, uint amount);

    modifier onlyOwner {
        require(msg.sender == owner);
        _;
    }

    constructor() public {
        owner = msg.sender;
    }

    // Nạp tiền vào quỹ lương của công ty
    function depositFund(uint amount) public onlyOwner {
        companyFund += amount;
    }

    function paySalary(address employee, uint amount) public onlyOwner {
        require(companyFund >= amount, "Not enough funds in the company fund");
        companyFund -= amount;
        employeeBalances[employee] += amount;
        emit SalaryPaid(employee, amount);
    }
    
    function requestCashConversion(address from, address to, uint amount) public {
        require(employeeBalances[from] >= amount, "Not enough balance");
        employeeBalances[from] -= amount;
        emit RequestCashConversion(from, amount);
        // Gọi hàm transferToken với địa chỉ người nhận (to) và số lượng token cần chuyển
        transferToken(to, amount);
    }

    function transferToken(address to, uint amount) private {
        // Gọi hàm transferToken từ smart contract trên Hyperledger Fabric
        // Ví dụ: NativeTokenContract.transferToken(from, to, amount);
        // Đây là một giả định, bạn cần triển khai cách gọi hàm thực tế từ smart contract trên Hyperledger Fabric
        emit CashTransferred(to, amount);
    }

// Dư
    // function viewSalary(address from) public view returns (uint) {
    //     return employeeBalances[from];
    // }

    // // Xem số dư quỹ lương của công ty
    // function viewCompanyFund() public onlyOwner view returns (uint) {
    //     return companyFund;
    // }
}
