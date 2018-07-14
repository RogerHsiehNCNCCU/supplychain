package main

import (
    //"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	//"strings"
	//"time"
    
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// HeroesServiceChaincode implementation of Chaincode
type tokenStoreChaincode struct {
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

// Init of the chaincode
// This function is called only one when the chaincode is instantiated.
// So the goal is to prepare the ledger to handle future requests.
func (t *tokenStoreChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("########### tokenStoreChaincode Init ###########")

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
func (t *tokenStoreChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("########### tokenStoreChaincode Invoke ###########")

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
	if args[0] == "tokenStore" {
		return t.tokenStore(stub, args)
	}

	// If the arguments given don’t match any function, we return an error
	return shim.Error("Unknown action, check the first argument")
}

// invoke
// Every functions that read and write in the ledger will be here
func (t *tokenStoreChaincode) tokenStore(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("########### tokenStoreChaincode invoke ###########")

	if len(args) < 3 {//1個function名稱加上2個值   poc Key以及儲值多少
		return shim.Error("The number of arguments is insufficient.")
	}
    
    //支付後才進入儲值階段，支付是線下處理
    
	pocAsBytes, _ := stub.GetState(args[1])
	poc := PoC{}

	json.Unmarshal(pocAsBytes, &poc)
    tok, err := strconv.Atoi(args[2])
    if err != nil {}
    poc.Token = poc.Token + tok

	pocAsBytes, _ = json.Marshal(poc)
	stub.PutState(args[1], pocAsBytes)

	return shim.Success(nil)
}

func main() {
	// Start the chaincode and make it ready for futures requests
	err := shim.Start(new(tokenStoreChaincode))
	if err != nil {
		fmt.Printf("Error starting Heroes Service chaincode: %s", err)
	}
}
