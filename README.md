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
- Initialize the ledger with some employees
- DepositFund adds funds to the company balance
- Paying salaries in tokens.
- Handling requests for cash conversion by updating employee balances.
- TransferToken transfers tokens to a recipient
- Viewing the current token balance.
