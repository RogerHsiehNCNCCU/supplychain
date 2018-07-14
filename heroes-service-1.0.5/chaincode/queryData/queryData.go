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
type QueryDataChaincode struct {
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
func (t *QueryDataChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("########### QueryDataChaincode Init ###########")

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
func (t *QueryDataChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("########### QueryDataChaincode Invoke ###########")

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
	if args[0] == "queryData" {
		return t.queryData(stub, args)
	}

	// If the arguments given don’t match any function, we return an error
	return shim.Error("Unknown action, check the first argument")
}

// query
// Every readonly functions in the ledger will be here
func (t *QueryDataChaincode) queryData(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("########### QueryDataChaincode query ###########")

	// Check whether the number of arguments is sufficient
    if len(args) < 3 {//invoke  , dataKey, yourKey(poc)
		return shim.Error("The number of arguments is insufficient.")
	}

    //需要先檢查是否滿足最低積分要求，並有扣除5積分及2%手續費流程
    //如果檢查不足，引導致積分充值畫面
    pocAsBytes, _ := stub.GetState(args[2])
    poc := PoC{}
    json.Unmarshal(pocAsBytes, &poc)
    
    if(poc.Token>5){
        poc.Token = poc.Token-5
        poc.Token = poc.Token - (profit) * 0.02//支付2%手續費完成
        pocAsBytes, _ := json.Marshal(poc)
        使用poc chaincode的 query all poc 然後找length+1為key
	    stub.PutState(POCXXX, pocAsBytes)
        // Like the Invoke function, we manage multiple type of query requests with the second argument.
        dataAsBytes, _ := stub.GetState(args[1])
        material := Material{}
        json.Unmarshal(dataAsBytes, &material)
        pocAsBytes2, _ := stub.GetState(material.owner)
        poc2 := Poc{}
        json.Unmarshal(pocAsBytes2, &poc2)
        poc2.Token = poc2.Token+5
        pocAsBytes2, _ = json.Marshal(poc2)
        使用poc chaincode的 query all poc 然後找length+1為key
	    stub.PutState(POCXXX, pocAsBytes2)
       
    
            //紀錄transactionRecord     order by 、 limit 還沒實作
            current_time := time.Now().Local()
            timestamp := current_time.Format("2006-01-02 12:05:55")

            queryString := fmt.Sprintf("{\"selector\":{\"timestamp\":\"\",\"action\":\"upload\"},order by timestamp desc, limit:1}")

            queryResults, err := getQueryResultForQueryString(stub, queryString)
            if err != nil {
                return shim.Error(err.Error())
            }
            transactionRecord := TransactionRecord{ Action: "get", fee: 5, participant: poc.Name, Timestamp: timestamp }

            transactionRecordAsBytes, _ := json.Marshal(transactionRecord)
            stub.PutState(queryResults[0].Key+1, transactionRecordAsBytes)
    
        return shim.Success(dataAsBytes)
    }else{ //積分不夠支付5 token
        引導到積分充值畫面
    }
    

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
	err := shim.Start(new(QueryDataChaincode))
	if err != nil {
		fmt.Printf("Error starting Heroes Service chaincode: %s", err)
	}
}
