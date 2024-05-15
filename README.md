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
![image](https://github.com/NguyenTaHuyHoang/Payroll/assets/85854007/22c7aaf6-6f59-46b8-ba1d-c4d17d4988cc)
- main.go: Main entry point for the chaincode.
- employee.go: Employee structure definition.
- initLedger.go: Initializes the ledger with default employees.
- paySalary.go: Adds tokens to employee balances.
- requestCashConversion.go: Handles token-to-cash conversion requests.
- viewSalary.go: Retrieves the current token balance of an employee.
- transferCash.go: Transfers cash to the employee upon request.
