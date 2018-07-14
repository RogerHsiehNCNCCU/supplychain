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
type InitLedgerChaincode struct {
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
func (t *InitLedgerChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("########### InitLedgerChaincode Init ###########")

	// Get the function and arguments from the request
	function, _ := stub.GetFunctionAndParameters()

	// Check if the request is the init function
	if function != "init" {
		return shim.Error("Unknown function call")
	}

	// Put in the ledger the key/value hello/world
//	err := stub.PutState("hello", []byte("world"))
//	if err != nil {
//		return shim.Error(err.Error())
//	}

	// Return a successful message
	return shim.Success(nil)
}

// Invoke
// All future requests named invoke will arrive here.
func (t *InitLedgerChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
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
	if args[0] == "initLedger" {
		return t.initLedger(stub, args)
	}

	// If the arguments given don’t match any function, we return an error
	return shim.Error("Unknown action, check the first argument")
}

// invoke
// Every functions that read and write in the ledger will be here
func (t *InitLedgerChaincode) initLedger(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("########### HeroesServiceChaincode invoke ###########")

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
    
        return shim.Success(nil)
	//}

	// If the arguments given don’t match any function, we return an error
	//return shim.Error("Unknown invoke action, check the second argument.")
}

func main() {
	// Start the chaincode and make it ready for futures requests
	err := shim.Start(new(InitLedgerChaincode))
	if err != nil {
		fmt.Printf("Error starting Heroes Service chaincode: %s", err)
	}
}
