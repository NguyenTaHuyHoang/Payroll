package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (s *PayrollContract) ViewSalary(ctx contractapi.TransactionContextInterface, employeeID string) (uint, error) {
	employeeJSON, err := ctx.GetStub().GetState(employeeID)
	if err != nil {
		return 0, err
	}
	if employeeJSON == nil {
		return 0, fmt.Errorf("Employee %s does not exist", employeeID)
	}

	employee := new(Employee)
	err = json.Unmarshal(employeeJSON, employee)
	if err != nil {
		return 0, err
	}

	return employee.Balance, nil
}
