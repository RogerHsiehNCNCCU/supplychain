package main

import (
	"fmt"
	"github.com/chainHero/heroes-service/blockchain"
	"github.com/chainHero/heroes-service/web"
//	"github.com/chainHero/heroes-service/web/controllers"
	"os"
)

func main() {
	// Definition of the Fabric SDK properties
//	fSetup := blockchain.FabricSetup{
//		// Channel parameters
//		ChannelID:     "chainhero",
//		ChannelConfig: os.Getenv("GOPATH") + "/src/github.com/chainHero/heroes-service/fixtures/artifacts/chainhero.channel.tx",
//
//		// Chaincode parameters
//		ChainCodeID:     "heroes-service",
//		ChaincodeGoPath: os.Getenv("GOPATH"),
//		ChaincodePath:   "github.com/chainHero/heroes-service/chaincode/",
//		OrgAdmin:        "Admin",
//		OrgName:         "Org1",
//		ConfigFile:      "config.yaml",
//
//		// User parameters
//		UserName: "User1",
//	}
//
//    fmt.Printf("第一個chiancode!!!")
//    fmt.Println(fSetup)
//    
//	// Initialization of the Fabric SDK from the previously set properties
//	err := fSetup.Initialize()
//	if err != nil {
//		fmt.Printf("Unable to initialize the Fabric SDK: %v\n", err)
//	}
//
//	// Install and instantiate the chaincode
//	err = fSetup.InstallAndInstantiateCC()
//	if err != nil {
//		fmt.Printf("Unable to install and instantiate the chaincode: %v\n", err)
//	}
//    
    fmt.Printf("第二個chiancode!!!")
    
    
    fSetup := blockchain.FabricSetup{
		// Channel parameters
		ChannelID:     "chainhero",
		ChannelConfig: os.Getenv("GOPATH") + "/src/github.com/chainHero/heroes-service/fixtures/artifacts/chainhero.channel.tx",

		// Chaincode parameters
		ChainCodeID:     "heroes-service",
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath:   "github.com/chainHero/heroes-service/chaincode/unity",
		OrgAdmin:        "Admin",
		OrgName:         "Org1",
		ConfigFile:      "config.yaml",

		// User parameters
		UserName: "User1",
	}

    fmt.Println(fSetup)
    
	// Initialization of the Fabric SDK from the previously set properties
	err := fSetup.Initialize()
	if err != nil {
		fmt.Printf("Unable to initialize the Fabric SDK: %v\n", err)
	}

    
//    fSetup.ChainCodeID = "unity"
//    fSetup.ChaincodePath = "github.com/chainHero/heroes-service/chaincode/unity"
//    
//    fmt.Println(fSetup)
    
    err = fSetup.InstallAndInstantiateCC()
    if err != nil{
       fmt.Printf("Unable to install and instantiate the chaincode: %v\n", err)
    }
    
//	// Install and instantiate the chaincode
//    err2 = fSetup2.InstallAndInstantiateCC2()
//	if err2 != nil {
//		fmt.Printf("Unable to install and instantiate the chaincode: %v\n", err2)
//	}
//
//    
    fmt.Printf("第三個chiancode!!!")
//    fSetup = blockchain.FabricSetup{
//        // Channel parameters
//        ChannelID:     "chainhero",
//        ChannelConfig: os.Getenv("GOPATH") + "/src/github.com/chainHero/heroes-service/fixtures/artifacts/chainhero.channel.tx",
//
//        // Chaincode parameters
//        ChainCodeID:     "queryTransactionRecord",
//        ChaincodeGoPath: os.Getenv("GOPATH"),
//        ChaincodePath:   "github.com/chainHero/heroes-service/chaincode/queryTransactionRecord/",
//        OrgAdmin:        "Admin",
//        OrgName:         "Org1",
//        ConfigFile:      "config.yaml",
//
//        // User parameters
//        UserName: "User1",
//    }
//
//
//
//Initialization of the Fabric SDK from the previously set properties
//    err = fSetup.Initialize()
//    if err != nil {
//        fmt.Printf("Unable to initialize the Fabric SDK: %v\n", err)
//    }

    fSetup.ChainCodeID = "queryTransactionRecord"
    fSetup.ChaincodePath = "github.com/chainHero/heroes-service/chaincode/queryTransactionRecord/"

    fmt.Println(fSetup)

    // Install and instantiate the chaincode
    err = fSetup.InstallAndInstantiateCC()
    if err != nil {
        fmt.Printf("Unable to install and instantiate the chaincode: %v\n", err)
    }
    
//	// Launch the web application listening
//	app := &controllers.Application{
//		Fabric: &fSetup,
//	}
//	web.Serve(app)
    
    app := &web.Application{
        Fabric: &fSetup,
    }
	web.MyServe(app)
}
