package main

import (
    "bytes"
	"encoding/json"
	"fmt"
	//"strconv"
	//"strings"
	"time"
    
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// HeroesServiceChaincode implementation of Chaincode
type UploadDataChaincode struct {
}

type Material struct {
	Name   string `json:"name"`
	Efficacy  string `json:"efficacy"`
	Color string `json:"color"`
	Owner  string `json:"owner"`
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

type TransactionRecord struct{
    Action string `json:'upload'`
    fee int `json:5`
    participant string `json:'Roger'`
    Timestamp string `json:'2018-06-29 09:05:22'`
}

// Init of the chaincode
// This function is called only one when the chaincode is instantiated.
// So the goal is to prepare the ledger to handle future requests.
func (t *UploadDataChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("########### UploadDataChaincode Init ###########")

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
func (t *UploadDataChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("########### HeroesServiceChaincode Invoke ###########")

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
//	if args[0] == "query" {
//		return t.query(stub, args)
//	}
//    if args[0] == "queryHistory" {
//		return t.queryHistory(stub, args)
//	}

	// The update argument will manage all update in the ledger
	if args[0] == "uploadData" {
		return t.uploadData(stub, args)
	}

	// If the arguments given don’t match any function, we return an error
	return shim.Error("Unknown action, check the first argument")
}

// invoke
// Every functions that read and write in the ledger will be here
func (t *UploadDataChaincode) uploadData(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("########### HeroesServiceChaincode invoke ###########")

	if len(args) < 5 {//1個function名稱加上4個值
		return shim.Error("The number of arguments is insufficient.")
	}

	var material = Material{Name: args[1], Efficacy: args[2], Color: args[3], Owner: args[4]}

	materialAsBytes, _ := json.Marshal(material)
	stub.PutState(args[0], materialAsBytes)
    
    //還需要有檢查是否超出一日100額度並儲值10積分
    pocAsBytes, _ := stub.GetState(args[4])
    poc := PoC{}
    json.Unmarshal(pocAsBytes, &poc)
    
    current_time := time.Now().Local()
    queryString := fmt.Sprintf("{\"selector\":{\"timestamp\":\"%s\",\"action\":\"upload\",\"participant\":\"%s\"}}", current_time.Format("2006-01-02"), poc.Name)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
    
    if(queryResults.length<10){//如果當日小於10筆上傳，加10積分
        
        poc.Token += 10

        pocAsBytes, _ = json.Marshal(poc)
        stub.PutState(args[4], pocAsBytes)
        
        //數據銀行token減10
        tmcAsBytes, _ := stub.GetState["TMC0"]
        tmc := TMC{}
        json.Unmarshal(tmcAsBytes, &tmc)
        
        tmc.Token -= 10 
        tmcAsBytes, _ = json.Marshal(tmc)
        stub.PutState("TMC0", tmcAsBytes)
    }
    
	return shim.Success(nil)
}

func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
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

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

func main() {
	// Start the chaincode and make it ready for futures requests
	err := shim.Start(new(UploadDataChaincode))
	if err != nil {
		fmt.Printf("Error starting Heroes Service chaincode: %s", err)
	}
}
