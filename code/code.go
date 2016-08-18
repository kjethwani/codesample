package main

//WARNING - this chaincode's ID is hard-coded in chaincode_example04 to illustrate one way of
//calling chaincode from a chaincode. If this example is modified, chaincode_example04.go has
//to be modified as well with the new ID of chaincode_example02.
//chaincode_example05 show's how chaincode ID can be passed in as a parameter instead of
//hard-coding.

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

func (t *SimpleChaincode) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	var custName, bankName string    // Entities	A,B
	var custMob, bankCode string // Asset holdings		Aval,Bval
	var err error
	var kycCust, kycBank, kycAll bool
	
	if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting 4")
	}

	// Initialize the chaincode
	custName = args[0]
	custMob = args[2]
	// custMob, err = strconv.Atoi(args[1])
	if err != nil {
		return nil, errors.New("Expecting integer value for Customer mobile")	// Expecting integer value for asset holdings
	}
	bankName = args[1]
	bankCode = args[3]
	// bankCode, err = strconv.Atoi(args[3])
	if err != nil {
		return nil, errors.New("Expecting integer value for bank code")			// Expecting integer value for asset holdings 
	}
	fmt.Printf("Customer Name = %s, Customer Mobile = %s\n", custName, custMob)
	fmt.Printf("Bank Name = %s, Bank Code = %s\n", bankName, bankCode)
	fmt.Printf("KYC Completed by Customer = %t, KYC Completed by Bank = %t, KYC Completed by both = %t\n", kycCust, kycBank, kycAll)

	// Write the state to the ledger
	err = stub.PutState(custName, []byte(custName))
	if err != nil {
		return nil, err
	}

	err = stub.PutState(bankName, []byte(bankName))
	if err != nil {
		return nil, err
	}
	
	err = stub.PutState(custMob, []byte(custMob))
	if err != nil {
		return nil, err
	}
	
	err = stub.PutState(bankCode, []byte(bankCode))
	if err != nil {
		return nil, err
	}
	
	return nil, nil
}

// Transaction makes payment of X units from A to B
func (t *SimpleChaincode) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	if function == "delete" {
		// Deletes an entity from its state
		return t.delete(stub, args)
	}

	var custName, bankName string    // Entities	A,B
	var custMob, bankCode string // Asset holdings		Aval,Bval
	var X int          // Transaction value
	var err error

	if len(args) != 3 {
		return nil, errors.New("Incorrect number of arguments. Expecting 3")
	}

	custName = args[0]
	custMob = args[1]

	// Get the state from the ledger
	// TODO: will be nice to have a GetAllState call to ledger
	custNamebytes, err := stub.GetState(custName)
	if err != nil {
		return nil, errors.New("Failed to get state")
	}
	if custNamebytes == nil {
		return nil, errors.New("Entity not found")
	}
	custName, _ = strconv.Atoi(string(custNamebytes))

	bankNamebytes, err := stub.GetState(bankName)
	if err != nil {
		return nil, errors.New("Failed to get state")
	}
	if bankNamebytes == nil {
		return nil, errors.New("Entity not found")
	}
	bankName, _ = strconv.Atoi(string(bankNamebytes))

	custMobbytes, err := stub.GetState(custMob)
	if err != nil {
		return nil, errors.New("Failed to get state")
	}
	if custMobbytes == nil {
		return nil, errors.New("Entity not found")
	}
	custMob, _ = strconv.Atoi(string(custMobbytes))

	bankCodebytes, err := stub.GetState(bankCode)
	if err != nil {
		return nil, errors.New("Failed to get state")
	}
	if bankCodebytes == nil {
		return nil, errors.New("Entity not found")
	}
	bankCode, _ = strconv.Atoi(string(bankCodebytes))

	var kycCust bool = true
	var kycBank bool = false
	var kycAll bool = false
	
	// Perform the execution
	// X, err = strconv.Atoi(args[2])
	// Aval = Aval - X
	// Bval = Bval + X
	// fmt.Printf("Aval = %d, Bval = %d\n", Aval, Bval)

	fmt.Printf("Customer Name = %s, Customer Mobile = %d\n", custName, custMob)
	fmt.Printf("Bank Name = %s, Bank Code = %d\n", bankName, bankCode)
	fmt.Printf("KYC Completed by Customer = %t, KYC Completed by Bank = %t, KYC Completed by both = %t\n", kycCust, kycBank, kycAll)
	
	// Write the state back to the ledger
	// err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
	// if err != nil {
	//	return nil, err
	// }

	// err = stub.PutState(B, []byte(strconv.Itoa(Bval)))
	// if err != nil {
	//	return nil, err
	// }

	// return nil, nil
	
	// Write the state to the ledger
	err = stub.PutState(custName, []byte(custName))
	if err != nil {
		return nil, err
	}

	err = stub.PutState(bankName, []byte(bankName))
	if err != nil {
		return nil, err
	}
	
	err = stub.PutState(custMob, []byte(custMob))
	if err != nil {
		return nil, err
	}
	
	err = stub.PutState(bankCode, []byte(bankCode))
	if err != nil {
		return nil, err
	}
	
	return nil, nil
}

// Deletes an entity from state
func (t *SimpleChaincode) delete(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	A := args[0]

	// Delete the key from the state in ledger
	err := stub.DelState(A)
	if err != nil {
		return nil, errors.New("Failed to delete state")
	}

	return nil, nil
}

// Query callback representing the query of a chaincode
func (t *SimpleChaincode) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	if function != "query" {
		return nil, errors.New("Invalid query function name. Expecting \"query\"")
	}
	var A string // Entities
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the person to query")
	}

	A = args[0]

	// Get the state from the ledger
	custNamebytes, err := stub.GetState(A)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
		return nil, errors.New(jsonResp)
	}

	if custNamebytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + A + "\"}"
		return nil, errors.New(jsonResp)
	}

	jsonResp := "{\"Name\":\"" + A + "\",\"Amount\":\"" + string(custNamebytes) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)
	return custNamebytes, nil
	
	
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
