package main

import (
    "bytes"
	"encoding/json"
	"fmt"
	"strconv"
	//"strings"
	"time"
    
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// UnityChaincode implementation of Chaincode
type UnityChaincode struct {
}

type Material struct {
	Name   string `json:"name"`
	Efficacy  string `json:"efficacy"`
	Color string `json:"color"`
	Owner  string `json:"owner"`
}

type TMC struct {
    Name string `json:TMC`
    Token float64 `json:30000000`
}

type PoC struct {
    Name string `json:'PoC'`
    Token float64 `json:10000`
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
func (t *UnityChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("########### InitLedgerChaincode Init ###########")

	// Get the function and arguments from the request
	function, _ := stub.GetFunctionAndParameters()

	// Check if the request is the init function
	if function != "init" {
		return shim.Error("Unknown function call")
	}

	// Put in the ledger the key/value hello/world
	err := stub.PutState("hello", []byte("world"))
	if err != nil {
		return shim.Error(err.Error())
	}

	// Return a successful message
	return shim.Success(nil)
}

// Invoke
// All future requests named invoke will arrive here.
func (t *UnityChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
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

	// The update argument will manage all update in the ledger
	if args[0] == "initLedger" {
		return t.initLedger(stub, args)
    }else if args[0] == "uploadData" {
		return t.uploadData(stub, args)
	}else if args[0] == "queryData" {
		return t.queryData(stub, args)
	}else if args[0] == "tokenStore" {
		return t.tokenStore(stub, args)
	}else if args[0] == "queryTokenByPoc" {
		return t.queryTokenByPoc(stub, args)
	}else if args[0] == "queryHistoryPoc" {
		return t.queryHistoryPoc(stub, args)
	}else if args[0] == "queryAllPoCs" {
		return t.queryAllPoCs(stub)
	}else if args[0] == "queryTransactionByKey" {
		return t.queryTransactionByKey(stub, args)
	}else if args[0] == "queryHistoryByKey" {
		return t.queryHistoryByKey(stub, args)
	}else if args[0] == "queryAllTransactions" {
		return t.queryAllTransactions(stub)
	} 

	// If the arguments given don’t match any function, we return an error
	return shim.Error("Unknown action, check the first argument")
}

// invoke
// Every functions that read and write in the ledger will be here
func (t *UnityChaincode) initLedger(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("########### UnityChaincode invoke ###########")

	if len(args) < 1 {
		return shim.Error("The number of arguments is insufficient.")
	}

	// Check if the ledger key is "hello" and process if it is the case. Otherwise it returns an error.
	//if args[1] == "hello" && len(args) == 3 {

		materials := []Material{
            Material{Name: "gochi", Efficacy: "good for body", Color: "red", Owner: "farmer1"},
            Material{Name: "abc", Efficacy: "abc", Color: "abc", Owner: "farmer2"},
            Material{Name: "qwe", Efficacy: "qwe", Color: "qwe", Owner: "farmer3"},
        }

        i := 0
        for i < len(materials) {
            fmt.Println("i is ", i)
            materialAsBytes, _ := json.Marshal(materials[i])
            stub.PutState("MATERIAL"+strconv.Itoa(i), materialAsBytes)
            fmt.Println("Added", materials[i])
            i = i + 1
        }
    
    
        tmc := []TMC{
            TMC{Name: "TMC", Token:29850000},
        }
    
        i=0
        for i< len(tmc){
            fmt.Println("i is ", i)
            tmcAsBytes, _ := json.Marshal(tmc[i])
            stub.PutState("TMC"+strconv.Itoa(i), tmcAsBytes)
            fmt.Println("Added", tmc[i])
            i = i + 1
        }
    
    
        poc := []PoC{
            PoC{Name: "farmer1", Token:10000, Role:"farmer"},
            PoC{Name: "farmer2", Token:10000, Role:"farmer"},
            PoC{Name: "farmer3", Token:10000, Role:"farmer"},
            PoC{Name: "farmer4", Token:10000, Role:"farmer"},
            PoC{Name: "farmer5", Token:10000, Role:"farmer"},
            PoC{Name: "plantBase1", Token:10000, Role:"plantBase"},
            PoC{Name: "plantBase2", Token:10000, Role:"plantBase"},
            PoC{Name: "wholesaleMarket1", Token:10000, Role:"wholesaleMarket"},
            PoC{Name: "wholesaleMarket2", Token:10000, Role:"wholesaleMarket"},
            PoC{Name: "materialCompony1", Token:10000, Role:"materialCompony"},
            PoC{Name: "wholesaleCompony1", Token:10000, Role:"wholesaleCompony"},
            PoC{Name: "drugstore1", Token:10000, Role:"drugstore"},
            PoC{Name: "drugstore2", Token:10000, Role:"drugstore"},
            PoC{Name: "hospital1", Token:10000, Role:"hospital"},
            PoC{Name: "hospital2", Token:10000, Role:"hospital"},
        }
        
        i=0
        for i< len(poc){
            fmt.Println("i is ", i)
            pocAsBytes, _ := json.Marshal(poc[i])
            stub.PutState("POC"+strconv.Itoa(i), pocAsBytes)
            fmt.Println("Added", poc[i])
            i = i + 1
        }    
    
        err := stub.SetEvent("eventInvoke", []byte{})
		if err != nil {
			return shim.Error(err.Error())
		}
    
        return shim.Success(nil)
	//}

	// If the arguments given don’t match any function, we return an error
	//return shim.Error("Unknown invoke action, check the second argument.")
}

func (t *UnityChaincode) uploadData(stub shim.ChaincodeStubInterface, args []string) pb.Response {
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
    
    if(len(queryResults)<10){//如果當日小於10筆上傳，加10積分
        
        poc.Token += 10

        pocAsBytes, _ = json.Marshal(poc)
        stub.PutState(args[4], pocAsBytes)
        
        //數據銀行token減10
        tmcAsBytes, _ := stub.GetState("TMC0")
        tmc := TMC{}
        json.Unmarshal(tmcAsBytes, &tmc)
        
        tmc.Token -= 10 
        tmcAsBytes, _ = json.Marshal(tmc)
        stub.PutState("TMC0", tmcAsBytes)
    }
    
	return shim.Success(nil)
}

func (t *UnityChaincode) queryData(stub shim.ChaincodeStubInterface, args []string) pb.Response {
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
    
    if poc.Token>=5 {
        poc.Token = poc.Token-5

        pocAsBytes, _ := json.Marshal(poc)
        //使用poc chaincode的 query all poc 然後找length+1為key
 //       queryResults, errq := internalQueryAllPoCs(stub)
 //       if errq!=nil {
 //           return shim.Error(errq.Error())
 //       }
 //       stub.PutState("POC"+strconv.Itoa(len(queryResults)+1), pocAsBytes)
 //應該是直接用原來的Key放入才對
        stub.PutState(args[2],pocAsBytes)
        // Like the Invoke function, we manage multiple type of query requests with the second argument.
        dataAsBytes, _ := stub.GetState(args[1])
        material := Material{}
        json.Unmarshal(dataAsBytes, &material)
        pocAsBytes2, _ := stub.GetState(material.Owner)
        poc2 := PoC{}
        json.Unmarshal(pocAsBytes2, &poc2)
        poc2.Token = poc2.Token+ 5*0.98
        pocAsBytes2, _ = json.Marshal(poc2)
        //使用poc chaincode的 query all poc 然後找length+1為key
        //stub.PutState("POC"+strconv.Itoa(len(queryResults)+1), pocAsBytes2)
		//也是直接放入就可，因為是同個PoC不是新的
		stub.PutState(material.Owner, pocAsBytes2)
        //數據銀行收益2%
        tmcAsBytes, _ := stub.GetState("TMC0")
        tmc := TMC{}
        json.Unmarshal(tmcAsBytes, &tmc)
        
        tmc.Token +=  5*0.02 
        tmcAsBytes, _ = json.Marshal(tmc)
        stub.PutState("TMC0", tmcAsBytes)

        //紀錄transactionRecord     order by 、 limit 還沒實作
        current_time := time.Now().Local()
        timestamp := current_time.Format("2006-01-02 12:05:55")

        /*queryString := fmt.Sprintf("{\"selector\":{\"timestamp\":\"\",\"action\":\"upload\"},order by timestamp desc, limit:1}")

        queryResults, err := getQueryResultForQueryString(stub, queryString)
        if err != nil {
            return shim.Error(err.Error())
        }*/
        queryResults2, errq2 := internalQueryAllTransactions(stub)
        if errq2 != nil{
            return shim.Error(errq2.Error())
        }
        transactionRecord := TransactionRecord{ Action: "get", fee: 5, participant: poc.Name, Timestamp: timestamp }

        transactionRecordAsBytes, _ := json.Marshal(transactionRecord)
        stub.PutState("TRANSACTIONRECORD"+strconv.Itoa(len(queryResults2)+1), transactionRecordAsBytes)
    
        return shim.Success(dataAsBytes)
    }else{ //積分不夠支付5 token
        //引導到積分充值畫面
        return shim.Success(nil)
    }
    

}

func (t *UnityChaincode) tokenStore(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("########### tokenStoreChaincode invoke ###########")

	if len(args) < 3 {//1個function名稱加上2個值   poc Key以及儲值多少
		return shim.Error("The number of arguments is insufficient.")
	}
    
    //支付後才進入儲值階段，支付是線下處理
    
	pocAsBytes, _ := stub.GetState(args[1])
	poc := PoC{}

	json.Unmarshal(pocAsBytes, &poc)
    tok, err := strconv.ParseFloat(args[2],64)
    if err != nil {}
    poc.Token = poc.Token + tok

	pocAsBytes, _ = json.Marshal(poc)
	stub.PutState(args[1], pocAsBytes)

	return shim.Success(nil)
}

func (t *UnityChaincode) queryTokenByPoc(stub shim.ChaincodeStubInterface, args []string) pb.Response {
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
func (t *UnityChaincode) queryHistoryPoc(stub shim.ChaincodeStubInterface, args []string) pb.Response {
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

func (t *UnityChaincode) queryAllPoCs(stub shim.ChaincodeStubInterface) pb.Response {

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

func internalQueryAllPoCs(stub shim.ChaincodeStubInterface) ([]byte, error) {

	startKey := "POC0"
	endKey := "POC99999"

	resultsIterator, err := stub.GetStateByRange(startKey, endKey)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
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

	fmt.Printf("- queryAllPoCs:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

func (t *UnityChaincode) queryTransactionByKey(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("########### QueryTransactionRecordChaincodee query ###########")

	// Check whether the number of arguments is sufficient
	if len(args) < 2 {
		return shim.Error("The number of arguments is insufficient.")
	}
    
	// Like the Invoke function, we manage multiple type of query requests with the second argument.
	transactionAsBytes, _ := stub.GetState(args[1])
	return shim.Success(transactionAsBytes)
}

// query History
func (t *UnityChaincode) queryHistoryByKey(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("########### QueryTransactionRecordChaincode query History ###########")

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

func (s *UnityChaincode) queryAllTransactions(stub shim.ChaincodeStubInterface) pb.Response {

	startKey := "TRANSACTIONRECORD0"
	endKey := "TRANSACTIONRECORD99999"

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

func internalQueryAllTransactions(stub shim.ChaincodeStubInterface) ([]byte, error) {

	startKey := "TRANSACTIONRECORD0"
	endKey := "TRANSACTIONRECORD99999"

	resultsIterator, err := stub.GetStateByRange(startKey, endKey)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
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

	fmt.Printf("- queryAllPoCs:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
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
	err := shim.Start(new(UnityChaincode))
	if err != nil {
		fmt.Printf("Error starting Heroes Service chaincode: %s", err)
	}
}
