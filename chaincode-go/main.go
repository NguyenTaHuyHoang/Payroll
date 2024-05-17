package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type PayrollContract struct {
	contractapi.Contract
}

type Employee struct {
	Address string `json:"address"`
	Balance uint   `json:"balance"`
}

func GenerateRandom160BitAddress() (string, error) {
	// Tạo một mảng byte có độ dài 20 byte (160 bit)
	bytes := make([]byte, 20)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	// Chuyển đổi mảng byte thành chuỗi hex
	address := hex.EncodeToString(bytes)
	return address, nil
}

// Initialize the ledger with some employees
func (s *PayrollContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	for i := 0; i < 2; i++ {
		address, err := GenerateRandom160BitAddress()
		if err != nil {
			return err
		}

		employee := Employee{
			Address: address,
			Balance: 0,
		}

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

// DepositFund adds funds to the company balance
func (pc *PayrollContract) DepositFund(ctx contractapi.TransactionContextInterface, amount uint) error {
	companyFundBytes, err := ctx.GetStub().GetState("companyFund")
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}

	companyFund := uint(0)
	if companyFundBytes != nil {
		companyFundParsed, parseErr := strconv.ParseUint(string(companyFundBytes), 10, 32)
		if parseErr != nil {
			return fmt.Errorf("failed to parse company fund: %v", parseErr)
		}
		companyFund = uint(companyFundParsed)
	}

	companyFund += amount
	return ctx.GetStub().PutState("companyFund", []byte(strconv.FormatUint(uint64(companyFund), 10)))
}

// PaySalary pays salary to an employee
func (pc *PayrollContract) PaySalary(ctx contractapi.TransactionContextInterface, employeeAddress string, amount uint) error {
	companyFundBytes, err := ctx.GetStub().GetState("companyFund")
	if err != nil {
		return fmt.Errorf("failed to read company fund from world state: %v", err)
	}

	companyFund, parseErr := strconv.ParseUint(string(companyFundBytes), 10, 32)
	if parseErr != nil {
		return fmt.Errorf("failed to parse company fund: %v", parseErr)
	}
	if companyFund < uint64(amount) {
		return fmt.Errorf("not enough funds in the company fund")
	}

	companyFund -= uint64(amount)
	err = ctx.GetStub().PutState("companyFund", []byte(strconv.FormatUint(companyFund, 10)))
	if err != nil {
		return fmt.Errorf("failed to update company fund: %v", err)
	}

	employeeBytes, err := ctx.GetStub().GetState(employeeAddress)
	if err != nil {
		return fmt.Errorf("failed to read employee balance from world state: %v", err)
	}

	employee := Employee{Address: employeeAddress}
	if employeeBytes != nil {
		err = json.Unmarshal(employeeBytes, &employee)
		if err != nil {
			return fmt.Errorf("failed to unmarshal employee balance: %v", err)
		}
	}

	employee.Balance += amount
	employeeBytes, err = json.Marshal(employee)
	if err != nil {
		return fmt.Errorf("failed to marshal employee balance: %v", err)
	}

	return ctx.GetStub().PutState(employeeAddress, employeeBytes)
}

// RequestCashConversion requests a cash conversion
func (pc *PayrollContract) RequestCashConversion(ctx contractapi.TransactionContextInterface, fromAddress string, toAddress string, amount uint) error {
	fromBytes, err := ctx.GetStub().GetState(fromAddress)
	if err != nil {
		return fmt.Errorf("failed to read employee balance from world state: %v", err)
	}

	if fromBytes == nil {
		return fmt.Errorf("employee balance not found")
	}

	fromEmployee := Employee{Address: fromAddress}
	err = json.Unmarshal(fromBytes, &fromEmployee)
	if err != nil {
		return fmt.Errorf("failed to unmarshal employee balance: %v", err)
	}

	if fromEmployee.Balance < amount {
		return fmt.Errorf("not enough balance")
	}

	fromEmployee.Balance -= amount
	fromBytes, err = json.Marshal(fromEmployee)
	if err != nil {
		return fmt.Errorf("failed to marshal employee balance: %v", err)
	}

	err = ctx.GetStub().PutState(fromAddress, fromBytes)
	if err != nil {
		return fmt.Errorf("failed to update employee balance: %v", err)
	}

	// Transfer token to the recipient
	err = pc.TransferToken(ctx, toAddress, amount)
	if err != nil {
		return fmt.Errorf("failed to transfer token: %v", err)
	}

	return nil
}

// TransferToken transfers tokens to a recipient
func (pc *PayrollContract) TransferToken(ctx contractapi.TransactionContextInterface, toAddress string, amount uint) error {
	toBytes, err := ctx.GetStub().GetState(toAddress)
	if err != nil {
		return fmt.Errorf("failed to read recipient balance from world state: %v", err)
	}

	toEmployee := Employee{Address: toAddress}
	if toBytes != nil {
		err = json.Unmarshal(toBytes, &toEmployee)
		if err != nil {
			return fmt.Errorf("failed to unmarshal recipient balance: %v", err)
		}
	}

	toEmployee.Balance += amount
	toBytes, err = json.Marshal(toEmployee)
	if err != nil {
		return fmt.Errorf("failed to marshal recipient balance: %v", err)
	}

	return ctx.GetStub().PutState(toAddress, toBytes)
}

// ViewSalary returns the balance of a specific employee
func (pc *PayrollContract) ViewSalary(ctx contractapi.TransactionContextInterface, employeeAddress string) (uint, error) {
	employeeBytes, err := ctx.GetStub().GetState(employeeAddress)
	if err != nil {
		return 0, fmt.Errorf("failed to read employee balance from world state: %v", err)
	}

	if employeeBytes == nil {
		return 0, fmt.Errorf("employee balance not found")
	}

	employee := Employee{}
	err = json.Unmarshal(employeeBytes, &employee)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal employee balance: %v", err)
	}

	return employee.Balance, nil
}

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
