package main

import (
	"encoding/json"
	"strconv"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (s *PayrollContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	employees := []Employee{
		{Address: "John Doe", Balance: 0},
		{Address: "Jane Doe", Balance: 0},
	}

	for i, employee := range employees {
		employeeJSON, err := json.Marshal(employee)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState("EMP"+strconv.Itoa(i), employeeJSON)
		if err != nil {
			return err
		}
	}

	return nil
}
