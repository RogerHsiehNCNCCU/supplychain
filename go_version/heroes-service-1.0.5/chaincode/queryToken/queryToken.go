package main

import (
    "bytes"
	//"encoding/json"
	"fmt"
	"strconv"
	//"strings"
	"time"
    
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// HeroesServiceChaincode implementation of Chaincode
type QueryTokenChaincode struct {
}

type TMC struct {
    Name string `json:TMC`
    Token int `json:30000000`
}

type PoC struct {
    Name string `json:'PoC'`
    Token int `json:10000`
    Role string `farmer`
}

// Init of the chaincode
// This function is called only one when the chaincode is instantiated.
// So the goal is to prepare the ledger to handle future requests.
func (t *QueryTokenChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("########### QueryTokenChaincode Init ###########")

	// Get the function and arguments from the request
	function, _ := stub.GetFunctionAndParameters()

	// Check if the request is the init function
	if function != "init" {
		return shim.Error("Unknown function call")
	}

	// Put in the ledger the key/value hello/world
	/*err := stub.PutState("hello", []byte("world"))
	if err != nil {
		return shim.Error(err.Error())
	}*/

	// Return a successful message
	return shim.Success(nil)
}

// Invoke
// All future requests named invoke will arrive here.
func (t *QueryTokenChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("########### QueryTokenChaincode Invoke ###########")

	// Get the function and arguments from the request
	function, args := stub.GetFunctionAndParameters()

	// Check whether it is an invoke request
	if function != "invoke" {
		return shim.Error("Unknown function call")
	}

	// Check whether the number of arguments is sufficient
	if len(args) < 1 {
		return shim.Error("The number of arguments is insufficient.")
	}

	// In order to manage multiple type of request, we will check the first argument.
	// Here we have one possible argument: query (every query request will read in the ledger without modification)
	if args[0] == "queryTokenByPoc" {
		return t.queryTokenByPoc(stub, args)
	}
    if args[0] == "queryHistoryPoc" {
		return t.queryHistoryPoc(stub, args)
	}
    if args[0] == "queryAllPoCs" {
		return t.queryAllPoCs(stub)
	}

	// If the arguments given don’t match any function, we return an error
	return shim.Error("Unknown action, check the first argument")
}

// query
// Every readonly functions in the ledger will be here
func (t *QueryTokenChaincode) queryTokenByPoc(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("########### QueryTokenChaincode query ###########")

	// Check whether the number of arguments is sufficient
	if len(args) < 2 {
		return shim.Error("The number of arguments is insufficient.")
	}
    
	// Like the Invoke function, we manage multiple type of query requests with the second argument.
	pocAsBytes, _ := stub.GetState(args[1])
	return shim.Success(pocAsBytes)
}

// query History
func (t *QueryTokenChaincode) queryHistoryPoc(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("########### QueryTokenChaincode query History ###########")

	// Check whether the number of arguments is sufficient
	if len(args) < 2 {
		return shim.Error("The number of arguments is insufficient.")
	}

	// Like the Invoke function, we manage multiple type of query requests with the second argument.
	// We also have only one possible argument: hello

    resultsIterator, err := stub.GetHistoryForKey(args[1])
    if err != nil {
        return shim.Error(err.Error())
    }
    defer resultsIterator.Close()

    // buffer is a JSON array containing historic values for the marble
    var buffer bytes.Buffer
    buffer.WriteString("[")

    bArrayMemberAlreadyWritten := false
    for resultsIterator.HasNext() {
        response, err := resultsIterator.Next()
        if err != nil {
            return shim.Error(err.Error())
        }
        // Add a comma before array members, suppress it for the first array member
        if bArrayMemberAlreadyWritten == true {
            buffer.WriteString(",")
        }
        buffer.WriteString("{\"TxId\":")
        buffer.WriteString("\"")
        buffer.WriteString(response.TxId)
        buffer.WriteString("\"")

        buffer.WriteString(", \"Value\":")
        // if it was a delete operation on given key, then we need to set the
        //corresponding value null. Else, we will write the response.Value
        //as-is (as the Value itself a JSON marble)
        if response.IsDelete {
            buffer.WriteString("null")
        } else {
            buffer.WriteString(string(response.Value))
        }

        buffer.WriteString(", \"Timestamp\":")
        buffer.WriteString("\"")
        buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
        buffer.WriteString("\"")

        buffer.WriteString(", \"IsDelete\":")
        buffer.WriteString("\"")
        buffer.WriteString(strconv.FormatBool(response.IsDelete))
        buffer.WriteString("\"")

        buffer.WriteString("}")
        bArrayMemberAlreadyWritten = true
    }
    buffer.WriteString("]")

    fmt.Printf("- getHistoryForPoC returning:\n%s\n", buffer.String())

    return shim.Success(buffer.Bytes())
}

func (s *QueryTokenChaincode) queryAllPoCs(stub shim.ChaincodeStubInterface) pb.Response {

	startKey := "POC0"
	endKey := "POC99999"

	resultsIterator, err := stub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllPoCs:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func main() {
	// Start the chaincode and make it ready for futures requests
	err := shim.Start(new(QueryTokenChaincode))
	if err != nil {
		fmt.Printf("Error starting Heroes Service chaincode: %s", err)
	}
}
