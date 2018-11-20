package blockchain

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/api/apitxn/chclient"
	"time"
)

// InvokeHello
func (setup *FabricSetup) InvokeHello(value string) (string, error) {

	// Prepare arguments
	var args []string
	args = append(args, "invoke")
	args = append(args, "initLedger")
	//args = append(args, "hello")
	args = append(args, value)

	eventID := "eventInvoke"

	// Add data that will be visible in the proposal, like a description of the invoke request
	transientDataMap := make(map[string][]byte)
	transientDataMap["result"] = []byte("Transient data in hello invoke")

	// Register a notification handler on the client
	notifier := make(chan *chclient.CCEvent)
	rce, err := setup.client.RegisterChaincodeEvent(notifier, setup.ChainCodeID, eventID)
	if err != nil {
		return "", fmt.Errorf("failed to register chaincode evet: %v", err)
	}

	// Create a request (proposal) and send it
	response, err := setup.client.Execute(chclient.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1]), []byte(args[2]), []byte(args[3])}, TransientMap: transientDataMap})
	if err != nil {
		return "", fmt.Errorf("failed to move funds: %v", err)
	}

	// Wait for the result of the submission
	select {
	case ccEvent := <-notifier:
		fmt.Printf("Received CC event: %s\n", ccEvent)
	case <-time.After(time.Second * 20):
		return "", fmt.Errorf("did NOT receive CC event for eventId(%s)", eventID)
	}

	// Unregister the notification handler previously created on the client
	err = setup.client.UnregisterChaincodeEvent(rce)

	return response.TransactionID.ID, nil
}

func (setup *FabricSetup) InvokeInitLedger() (string, error) {

	// Prepare arguments
	var args []string
	args = append(args, "invoke")
	args = append(args, "initLedger")
	//args = append(args, "hello")
	//args = append(args, value)

	eventID := "eventInvoke"

	// Add data that will be visible in the proposal, like a description of the invoke request
	transientDataMap := make(map[string][]byte)
	transientDataMap["result"] = []byte("Transient data in hello invoke")

	// Register a notification handler on the client
	notifier := make(chan *chclient.CCEvent)
	rce, err := setup.client.RegisterChaincodeEvent(notifier, setup.ChainCodeID, eventID)
	if err != nil {
		return "", fmt.Errorf("failed to register chaincode evet: %v", err)
	}

	// Create a request (proposal) and send it
	response, err := setup.client.Execute(chclient.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1])}, TransientMap: transientDataMap})
	if err != nil {
		return "", fmt.Errorf("failed to move funds: %v", err)
	}

	// Wait for the result of the submission
	select {
	case ccEvent := <-notifier:
		fmt.Printf("Received CC event: %s\n", ccEvent)
	case <-time.After(time.Second * 20):
        return "", fmt.Errorf("did NOT receive CC event for eventId(%s) (initLedger)", eventID)
	}

	// Unregister the notification handler previously created on the client
	err = setup.client.UnregisterChaincodeEvent(rce)

	return response.TransactionID.ID, nil
}

func (setup *FabricSetup) UploadData(inputs [4]string) (string, error){
    
    var args []string
    args = append(args, "invoke")
    args = append(args, "uploadData")
    args = append(args, inputs[0])
    args = append(args, inputs[1])
    args = append(args, inputs[2])
    args = append(args, inputs[3])
    
    eventID := "eventInvoke"
    
    transientDataMap := make(map[string][]byte)
    transientDataMap["result"] = []byte("Transient data in uploadData invoke")
    
    notifier := make(chan *chclient.CCEvent)
    rce, err := setup.client.RegisterChaincodeEvent(notifier, setup.ChainCodeID, eventID)
    if err != nil{
        return "", fmt.Errorf("failed to register chaincode event: %v", err)
    }
    
    response, err := setup.client.Execute(chclient.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1]), []byte(args[2]), []byte(args[3]), []byte(args[4]), []byte(args[5])}, TransientMap: transientDataMap})
    if err != nil {
        return "", fmt.Errorf("failed to move funds: %v", err)
    }
    
    select{
        case ccEvent := <-notifier:
		    fmt.Printf("Received CC event: %s\n", ccEvent)
	    case <-time.After(time.Second * 20):
            return "", fmt.Errorf("did NOT receive CC event for eventId(%s) (initLedger)", eventID)
    }
    
    err = setup.client.UnregisterChaincodeEvent(rce)

	return response.TransactionID.ID, nil
}