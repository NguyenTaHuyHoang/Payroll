# Payroll Smart Contract and Hyperledger Fabric Chaincode
## Introduction
This project includes a Solidity smart contract and Hyperledger Fabric chaincode for managing a company's payroll system using internal tokens. Employees can receive their salaries in tokens and request to convert these tokens into cash.

## Solidity Smart Contract
The Solidity contract includes functions for:
- Paying salaries in tokens.
- Requesting cash conversion for tokens.
- Viewing the current token balance.

## Hyperledger Fabric Chaincode
The chaincode manages:
- Initializing the ledger with employee data.
- Paying salaries in tokens.
- Handling requests for cash conversion by updating employee balances.
- Viewing the current token balance.

### Chaincode Structure

![image](https://github.com/NguyenTaHuyHoang/Payroll/assets/85854007/171a7fea-70fa-4c36-9130-4276ecc41dbf)


- main.go: Main entry point for the chaincode.
- payrollContract.go: defines the structure and functions of smart contracts
- employee.go: Employee structure definition.
- initLedger.go: Initializes the ledger with default employees.
- paySalary.go: Adds tokens to employee balances.
- requestCashConversion.go: Handles token-to-cash conversion requests.
- transferCash.go: Transfers cash to the employee upon request.
- viewSalary.go: Retrieves the current token balance of an employee.

