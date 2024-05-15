package main

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	payrollContract := new(PayrollContract)
	cc, err := contractapi.NewChaincode(payrollContract)
	if err != nil {
		fmt.Printf("Error create payroll chaincode: %s", err)
		return
	}

	if err := cc.Start(); err != nil {
		fmt.Printf("Error starting payroll chaincode: %s", err)
	}
}
