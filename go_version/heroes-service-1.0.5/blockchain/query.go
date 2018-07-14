package blockchain

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/api/apitxn/chclient"
    
    //"reflect"
)

// QueryHello query the chaincode to get the state of hello
func (setup *FabricSetup) QueryHello() (string, error) {

	// Prepare arguments
	var args []string
	args = append(args, "invoke")
	args = append(args, "query")
	args = append(args, "hello")

	response, err := setup.client.Query(chclient.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1]), []byte(args[2])}})
	if err != nil {
		return "", fmt.Errorf("failed to query: %v", err)
	}

	return string(response.Payload), nil
}

func (setup *FabricSetup) QueryHelloHistory() (string, error) {

	// Prepare arguments
	var args []string
	args = append(args, "invoke")
	args = append(args, "queryHistory")
	args = append(args, "hello")

	response, err := setup.client.Query(chclient.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1]), []byte(args[2])}})
	if err != nil {
		return "", fmt.Errorf("failed to query: %v", err)
	}

	return string(response.Payload), nil
}

func (setup *FabricSetup) QueryTransactionByKey(key string) (string, error){
    // Prepare arguments
	var args []string
	args = append(args, "invoke")
	args = append(args, "queryTransactionByKey")
	args = append(args, key)

	response, err := setup.client.Query(chclient.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1]), []byte(args[2])}})
	if err != nil {
		return "", fmt.Errorf("failed to query: %v", err)
	}

	return string(response.Payload), nil
}

func (setup *FabricSetup) QueryAllPoCs() (string, error){
    // Prepare arguments
	var args []string
	args = append(args, "invoke")
	args = append(args, "queryAllPoCs")
	//args = append(args, key)

    fmt.Println("Call QueryAllPoCs!")
    
	response, err := setup.client.Query(chclient.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1])}})
    
//    fmt.Println("response")
//    fmt.Println((response)
//    fmt.Println(reflect.TypeOf(response))
//    fmt.Println("response.Payload")
//    fmt.Println(string(response.Payload))
//    fmt.Println(reflect.TypeOf(response.Payload))
    
	if err != nil {
		return "", fmt.Errorf("failed to query: %v", err)
	}
//    if err != nil {
//		return nil, err
//	}

    return string(response.Payload), nil
}