package main

import (
    "encoding/json"
    "fmt"
    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (s *PayrollContract) PaySalary(ctx contractapi.TransactionContextInterface, employeeID string, amount int) error {
    employeeJSON, err := ctx.GetStub().GetState(employeeID)
    if err != nil {
        return err
    }
    if employeeJSON == nil {
        return fmt.Errorf("Employee %s does not exist", employeeID)
    }

    employee := new(Employee)
    err = json.Unmarshal(employeeJSON, employee)
    if err != nil {
        return err
    }

    employee.Balance += amount

    employeeJSON, err = json.Marshal(employee)
    if err != nil {
        return err
    }

    return ctx.GetStub().PutState(employeeID, employeeJSON)
}
