package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (s *PayrollContract) RequestCashConversion(ctx contractapi.TransactionContextInterface, employeeID string, amount uint) error {
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

	if employee.Balance < amount {
		return fmt.Errorf("Insufficient balance")
	}
	employee.Balance -= amount

	employeeJSON, err = json.Marshal(employee)
	if err != nil {
		return err
	}

	err = ctx.GetStub().PutState(employeeID, employeeJSON)
	if err != nil {
		return err
	}

	// Call TransferCash to transfer cash to the employee
	return s.TransferCash(ctx, "ACCOUNTING", amount)
}
